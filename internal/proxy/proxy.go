package proxy

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
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
}

// ReloadConfig reloads all configuration from the database
func (ps *ProxyServer) ReloadConfig() {
	ps.router.Reload()
	ps.headerProcessor.Reload()
	ps.queryRewriter.Reload()
	ps.cache.Clear()

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
	// Step 1: Match host + route
	match := ps.router.Match(r.Host, r.URL.Path)
	if match == nil {
		http.Error(w, `{"error":"no matching route found"}`, http.StatusBadGateway)
		return
	}

	vhost := match.VirtualHost
	route := match.VirtualDirectory

	// Step 2: Check method allowed
	if !ps.isMethodAllowed(route, r.Method) {
		http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	// Step 3: Enforce max request size
	if route.MaxRequestSizeMB > 0 {
		maxBytes := int64(route.MaxRequestSizeMB) * 1024 * 1024
		if r.ContentLength > maxBytes {
			http.Error(w, `{"error":"request entity too large"}`, http.StatusRequestEntityTooLarge)
			return
		}
		r.Body = http.MaxBytesReader(w, r.Body, maxBytes)
	}

	// Step 4: Check WebSocket upgrade
	if route.WebsocketEnabled && ps.wsProxy.IsWebSocketUpgrade(r) {
		ps.handleWebSocket(w, r, vhost, route)
		return
	}

	// Step 5: Check cache
	if route.CacheEnabled && r.Method == http.MethodGet {
		if cached := ps.cache.Get(r); cached != nil {
			// Serve from cache
			for k, vv := range cached.Header {
				for _, v := range vv {
					w.Header().Add(k, v)
				}
			}
			w.Header().Set("X-Cache", "HIT")
			w.WriteHeader(cached.StatusCode)
			w.Write(cached.Body)
			return
		}
	}

	// Step 6: Apply request header rules + query rewrites
	ps.headerProcessor.ApplyRequestHeaders(route.ID, r)
	ps.queryRewriter.Apply(route.ID, r)

	// Step 7: Select upstream via load balancer
	upstream := ps.selectUpstream(vhost, r)
	if upstream == nil {
		http.Error(w, `{"error":"no healthy upstream servers available"}`, http.StatusServiceUnavailable)
		return
	}

	// Step 8: Forward request via transport (with retries)
	resp, err := ps.transport.Forward(r, upstream, route)
	if err != nil {
		ps.healthChecker.MarkFailed(upstream)
		log.Printf("[Proxy] Forward error to %s:%d: %v", upstream.TargetHost, upstream.TargetPort, err)
		http.Error(w, fmt.Sprintf(`{"error":"upstream error: %s"}`, err.Error()), http.StatusBadGateway)
		return
	}

	// Mark upstream as successful
	ps.healthChecker.MarkSuccess(upstream)

	// Step 9: Apply response header rules
	ps.headerProcessor.ApplyResponseHeaders(route.ID, resp.Header)

	// Step 10: Cache response if enabled
	if route.CacheEnabled {
		ps.cache.Set(r, resp, route.CacheTTLSeconds)
	}

	// Step 11: Write response to client
	for k, vv := range resp.Header {
		for _, v := range vv {
			w.Header().Add(k, v)
		}
	}
	if route.CacheEnabled && r.Method == http.MethodGet {
		w.Header().Set("X-Cache", "MISS")
	}

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
	allUpstreams := ps.router.GetUpstreams(vhost.ID)
	if len(allUpstreams) == 0 {
		return nil
	}

	// Filter to healthy servers
	healthy := ps.healthChecker.GetHealthyServers(allUpstreams)
	if len(healthy) == 0 {
		// If all servers are unhealthy, try all as a fallback
		healthy = allUpstreams
	}

	// Get or create load balancer for this vhost
	lb := ps.getLoadBalancer(vhost)

	// Handle sticky sessions
	if vhost.StickySession {
		sw := NewStickyWrapper(lb)
		selected, needsCookie := sw.Select(healthy, r)
		_ = needsCookie // cookie is set in ServeHTTP
		return selected
	}

	return lb.Select(healthy, r)
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
