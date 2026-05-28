package proxy

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

// ProxyServer is the main reverse proxy handler
type ProxyServer struct {
	db              *sql.DB
	router          *Router
	healthChecker   *HealthChecker
	headerProcessor *HeaderProcessor
	queryRewriter   *QueryRewriter
	cache           *ResponseCache
	transport       *Transport
	wsProxy         *WebSocketProxy

	// Phase 8: Security components
	authenticator   *Authenticator
	rateLimiter     *ProxyRateLimiter
	circuitBreaker  *ProxyCircuitBreaker
	ipFilter        *IPFilter
	cors            *ProxyCORS
	maintenance     *MaintenanceChecker
	accessLogger    *AccessLogger

	// Load balancers per vhost (lazy initialized)
	lbMu          sync.RWMutex
	loadBalancers map[int64]LoadBalancer
}

// NewProxyServer creates a new ProxyServer, loading all config from DB
func NewProxyServer(db *sql.DB) *ProxyServer {
	ps := &ProxyServer{
		db:              db,
		router:          NewRouter(db),
		healthChecker:   NewHealthChecker(),
		headerProcessor: NewHeaderProcessor(db),
		queryRewriter:   NewQueryRewriter(db),
		cache:           NewResponseCache(),
		transport:       NewTransport(),
		wsProxy:         NewWebSocketProxy(),
		authenticator:   NewAuthenticator(db),
		rateLimiter:     NewProxyRateLimiter(db),
		circuitBreaker:  NewProxyCircuitBreaker(db),
		ipFilter:        NewIPFilter(db),
		cors:            NewProxyCORS(db),
		maintenance:     NewMaintenanceChecker(db),
		accessLogger:    NewAccessLogger(db),
		loadBalancers:   make(map[int64]LoadBalancer),
	}
	return ps
}

// Start begins the health checker background goroutines
func (ps *ProxyServer) Start() {
	ps.router.mu.RLock()
	upstreams := ps.router.upstreams
	ps.router.mu.RUnlock()
	ps.healthChecker.Start(upstreams)
	log.Println("[Proxy] Proxy server started with health checks")
}

// Stop gracefully stops the proxy server
func (ps *ProxyServer) Stop() {
	ps.healthChecker.Stop()
	ps.accessLogger.Stop()
}

// ReloadConfig reloads all configuration from the database
func (ps *ProxyServer) ReloadConfig() {
	ps.router.Reload()
	ps.headerProcessor.Reload()
	ps.queryRewriter.Reload()
	ps.cache.Clear()

	// Reload security components
	ps.authenticator.Reload()
	ps.rateLimiter.Reload()
	ps.circuitBreaker.Reload()
	ps.ipFilter.Reload()
	ps.cors.Reload()
	ps.maintenance.Reload()

	// Restart health checks
	ps.healthChecker.Stop()
	ps.router.mu.RLock()
	upstreams := ps.router.upstreams
	ps.router.mu.RUnlock()
	ps.healthChecker.Start(upstreams)

	// Clear LB cache
	ps.lbMu.Lock()
	ps.loadBalancers = make(map[int64]LoadBalancer)
	ps.lbMu.Unlock()

	log.Println("[Proxy] Configuration reloaded")
}

// ServeHTTP implements the http.Handler interface - main proxy pipeline
func (ps *ProxyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	// Step 1: Match host + route
	match := ps.router.Match(r.Host, r.URL.Path)
	if match == nil {
		http.Error(w, `{"error":"no matching route found"}`, http.StatusBadGateway)
		return
	}

	vhost := match.VirtualHost
	route := match.VirtualDirectory

	// Step 2: Maintenance mode check
	if inMaint, cfg := ps.maintenance.IsInMaintenance(vhost.ID); inMaint {
		http.Error(w, fmt.Sprintf(`{"error":"%s"}`, cfg.Message), cfg.ResponseCode)
		ps.accessLogger.LogRequest(r, cfg.ResponseCode, 0, time.Since(start), vhost.ID, route.ID, nil)
		return
	}

	// Step 3: CORS preflight handling
	if ps.cors.HandlePreflight(w, r, route) {
		ps.accessLogger.LogRequest(r, http.StatusNoContent, 0, time.Since(start), vhost.ID, route.ID, nil)
		return
	}

	// Step 4: IP filter
	if allowed, msg := ps.ipFilter.Allow(r, route); !allowed {
		http.Error(w, fmt.Sprintf(`{"error":"%s"}`, msg), http.StatusForbidden)
		ps.accessLogger.LogRequest(r, http.StatusForbidden, 0, time.Since(start), vhost.ID, route.ID, nil)
		return
	}

	// Step 5: Rate limiting
	if allowed, msg := ps.rateLimiter.Allow(r, route); !allowed {
		http.Error(w, fmt.Sprintf(`{"error":"%s"}`, msg), http.StatusTooManyRequests)
		ps.accessLogger.LogRequest(r, http.StatusTooManyRequests, 0, time.Since(start), vhost.ID, route.ID, nil)
		return
	}

	// Step 6: Authentication
	if ok, msg := ps.authenticator.Authenticate(r, route); !ok {
		http.Error(w, fmt.Sprintf(`{"error":"%s"}`, msg), http.StatusUnauthorized)
		ps.accessLogger.LogRequest(r, http.StatusUnauthorized, 0, time.Since(start), vhost.ID, route.ID, nil)
		return
	}

	// Step 7: Circuit breaker
	if allowed, msg := ps.circuitBreaker.Allow(route.ID); !allowed {
		http.Error(w, fmt.Sprintf(`{"error":"%s"}`, msg), http.StatusServiceUnavailable)
		ps.accessLogger.LogRequest(r, http.StatusServiceUnavailable, 0, time.Since(start), vhost.ID, route.ID, nil)
		return
	}

	// Step 8: Check method allowed
	if !ps.isMethodAllowed(route, r.Method) {
		http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
		ps.accessLogger.LogRequest(r, http.StatusMethodNotAllowed, 0, time.Since(start), vhost.ID, route.ID, nil)
		return
	}

	// Step 9: Enforce max request size
	if route.MaxRequestSizeMB > 0 {
		maxBytes := int64(route.MaxRequestSizeMB) * 1024 * 1024
		if r.ContentLength > maxBytes {
			http.Error(w, `{"error":"request entity too large"}`, http.StatusRequestEntityTooLarge)
			ps.accessLogger.LogRequest(r, http.StatusRequestEntityTooLarge, 0, time.Since(start), vhost.ID, route.ID, nil)
			return
		}
		r.Body = http.MaxBytesReader(w, r.Body, maxBytes)
	}

	// Step 10: Check WebSocket upgrade
	if route.WebsocketEnabled && ps.wsProxy.IsWebSocketUpgrade(r) {
		ps.handleWebSocket(w, r, vhost, route)
		return
	}

	// Step 11: Check cache
	if route.CacheEnabled && r.Method == http.MethodGet {
		if cached := ps.cache.Get(r); cached != nil {
			// Serve from cache
			for k, vv := range cached.Header {
				for _, v := range vv {
					w.Header().Add(k, v)
				}
			}
			w.Header().Set("X-Cache", "HIT")
			ps.cors.ApplyCORSHeaders(w, r, route)
			w.WriteHeader(cached.StatusCode)
			w.Write(cached.Body)
			ps.accessLogger.LogRequest(r, cached.StatusCode, len(cached.Body), time.Since(start), vhost.ID, route.ID, nil)
			return
		}
	}

	// Step 12: Apply request header rules + query rewrites
	ps.headerProcessor.ApplyRequestHeaders(route.ID, r)
	ps.queryRewriter.Apply(route.ID, r)

	// Step 13: Select upstream and forward with failover
	// Try each upstream one by one. If one fails, move to the next.
	upstreamList := ps.getUpstreamsOrdered(vhost, r)
	if len(upstreamList) == 0 {
		http.Error(w, `{"error":"no upstream servers available"}`, http.StatusServiceUnavailable)
		ps.accessLogger.LogRequest(r, http.StatusServiceUnavailable, 0, time.Since(start), vhost.ID, route.ID, nil)
		return
	}

	var resp *ProxyResponse
	var upstream *UpstreamConfig
	var lastErr error

	for i, candidate := range upstreamList {
		resp, lastErr = ps.transport.ForwardOnce(r, candidate, route)
		if lastErr == nil && resp.StatusCode < 502 {
			// Success
			upstream = candidate
			ps.healthChecker.MarkSuccess(candidate)
			ps.circuitBreaker.RecordSuccess(route.ID)
			break
		}

		// Failure — mark and try next upstream
		ps.healthChecker.MarkFailed(candidate)
		if lastErr == nil {
			lastErr = fmt.Errorf("upstream returned %d", resp.StatusCode)
		}
		log.Printf("[Proxy] Upstream %s:%d failed (%d/%d): %v",
			candidate.TargetHost, candidate.TargetPort, i+1, len(upstreamList), lastErr)
		resp = nil

		// Brief delay before trying next
		if i < len(upstreamList)-1 && route.RetryDelayMs > 0 {
			time.Sleep(time.Duration(route.RetryDelayMs) * time.Millisecond)
		}
	}

	if resp == nil {
		ps.circuitBreaker.RecordFailure(route.ID)
		errMsg := "all upstream servers failed"
		if lastErr != nil {
			errMsg = fmt.Sprintf("all %d upstreams failed: %s", len(upstreamList), lastErr.Error())
		}
		http.Error(w, fmt.Sprintf(`{"error":"%s"}`, errMsg), http.StatusBadGateway)
		ps.accessLogger.LogRequest(r, http.StatusBadGateway, 0, time.Since(start), vhost.ID, route.ID, upstream)
		return
	}

	// Step 15: Apply response header rules
	ps.headerProcessor.ApplyResponseHeaders(route.ID, resp.Header)

	// Step 16: Cache response if enabled
	if route.CacheEnabled {
		ps.cache.Set(r, resp, route.CacheTTLSeconds)
	}

	// Step 17: Write response to client
	for k, vv := range resp.Header {
		for _, v := range vv {
			w.Header().Add(k, v)
		}
	}
	if route.CacheEnabled && r.Method == http.MethodGet {
		w.Header().Set("X-Cache", "MISS")
	}

	// Apply CORS headers to response
	ps.cors.ApplyCORSHeaders(w, r, route)

	// Set sticky session cookie if needed
	if vhost.StickySession {
		http.SetCookie(w, &http.Cookie{
			Name:     stickyCookieName,
			Value:    formatServerID(upstream.ID),
			Path:     "/",
			HttpOnly: true,
		})
	}

	w.WriteHeader(resp.StatusCode)
	w.Write(resp.Body)

	// Step 18: Access log
	ps.accessLogger.LogRequest(r, resp.StatusCode, len(resp.Body), time.Since(start), vhost.ID, route.ID, upstream)
}

// handleWebSocket handles WebSocket upgrade requests
func (ps *ProxyServer) handleWebSocket(w http.ResponseWriter, r *http.Request, vhost *VHostConfig, route *VDirConfig) {
	upstream := ps.selectUpstream(vhost, r)
	if upstream == nil {
		http.Error(w, `{"error":"no healthy upstream servers available"}`, http.StatusServiceUnavailable)
		return
	}

	err := ps.wsProxy.Proxy(w, r, upstream, route)
	if err != nil {
		log.Printf("[Proxy] WebSocket proxy error: %v", err)
		// Connection might already be hijacked, can't send HTTP error
	}
}

// selectUpstream selects an upstream server using load balancing
func (ps *ProxyServer) selectUpstream(vhost *VHostConfig, r *http.Request) *UpstreamConfig {
	upstreams := ps.getUpstreamsOrdered(vhost, r)
	if len(upstreams) == 0 {
		return nil
	}
	return upstreams[0]
}

// getUpstreamsOrdered returns all upstream servers for a vhost, ordered by load balancer preference.
// Healthy servers come first, then unhealthy ones as fallback.
func (ps *ProxyServer) getUpstreamsOrdered(vhost *VHostConfig, r *http.Request) []*UpstreamConfig {
	allUpstreams := ps.router.GetUpstreams(vhost.ID)
	if len(allUpstreams) == 0 {
		return nil
	}

	// Separate healthy and unhealthy
	var healthy, unhealthy []*UpstreamConfig
	for _, u := range allUpstreams {
		if ps.healthChecker.IsHealthy(u.ID) {
			healthy = append(healthy, u)
		} else {
			unhealthy = append(unhealthy, u)
		}
	}

	// If no healthy, use all as fallback
	if len(healthy) == 0 {
		healthy = allUpstreams
		unhealthy = nil
	}

	// Use load balancer to determine starting order for healthy servers
	lb := ps.getLoadBalancer(vhost)
	if vhost.StickySession {
		sw := NewStickyWrapper(lb)
		preferred, _ := sw.Select(healthy, r)
		if preferred != nil {
			// Put preferred first, then the rest
			ordered := []*UpstreamConfig{preferred}
			for _, u := range healthy {
				if u.ID != preferred.ID {
					ordered = append(ordered, u)
				}
			}
			return append(ordered, unhealthy...)
		}
	} else {
		preferred := lb.Select(healthy, r)
		if preferred != nil {
			ordered := []*UpstreamConfig{preferred}
			for _, u := range healthy {
				if u.ID != preferred.ID {
					ordered = append(ordered, u)
				}
			}
			return append(ordered, unhealthy...)
		}
	}

	// Fallback: healthy first, then unhealthy
	return append(healthy, unhealthy...)
}

// getLoadBalancer returns the load balancer for a vhost, creating if needed
func (ps *ProxyServer) getLoadBalancer(vhost *VHostConfig) LoadBalancer {
	ps.lbMu.RLock()
	lb, ok := ps.loadBalancers[vhost.ID]
	ps.lbMu.RUnlock()

	if ok {
		return lb
	}

	// Create new load balancer
	lb = NewLoadBalancer(vhost.LBAlgorithm)

	ps.lbMu.Lock()
	ps.loadBalancers[vhost.ID] = lb
	ps.lbMu.Unlock()

	return lb
}

// isMethodAllowed checks if the request method is allowed for the route
func (ps *ProxyServer) isMethodAllowed(route *VDirConfig, method string) bool {
	// If no methods specified, allow all
	if len(route.AllowedMethods) == 0 {
		return true
	}

	method = strings.ToUpper(method)
	for _, m := range route.AllowedMethods {
		if m == method {
			return true
		}
	}
	return false
}
