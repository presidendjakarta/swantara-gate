package proxy

import (
	"hash/fnv"
	"math/rand"
	"net/http"
	"sync"
	"sync/atomic"
)

// LoadBalancer selects an upstream server based on a given algorithm
type LoadBalancer interface {
	// Select picks an upstream from the available healthy servers
	Select(servers []*UpstreamConfig, r *http.Request) *UpstreamConfig
}

// NewLoadBalancer creates a load balancer based on algorithm name
func NewLoadBalancer(algorithm string) LoadBalancer {
	switch algorithm {
	case "round_robin":
		return &RoundRobinLB{}
	case "weighted_round_robin":
		return &WeightedRoundRobinLB{}
	case "least_conn":
		return &LeastConnLB{}
	case "ip_hash":
		return &IPHashLB{}
	case "random":
		return &RandomLB{}
	case "failover":
		return &FailoverLB{}
	default:
		return &RoundRobinLB{}
	}
}

// ===========================
// Round Robin
// ===========================

// RoundRobinLB cycles through servers sequentially
type RoundRobinLB struct {
	counter uint64
}

func (lb *RoundRobinLB) Select(servers []*UpstreamConfig, r *http.Request) *UpstreamConfig {
	if len(servers) == 0 {
		return nil
	}
	idx := atomic.AddUint64(&lb.counter, 1) - 1
	return servers[idx%uint64(len(servers))]
}

// ===========================
// Weighted Round Robin
// ===========================

// WeightedRoundRobinLB distributes requests based on server weights
type WeightedRoundRobinLB struct {
	mu      sync.Mutex
	current int
	cw      int // current weight
	maxW    int
	gcdW    int
}

func (lb *WeightedRoundRobinLB) Select(servers []*UpstreamConfig, r *http.Request) *UpstreamConfig {
	if len(servers) == 0 {
		return nil
	}
	if len(servers) == 1 {
		return servers[0]
	}

	lb.mu.Lock()
	defer lb.mu.Unlock()

	// Compute max weight and GCD of weights
	maxW := 0
	gcdW := 0
	for _, s := range servers {
		w := s.Weight
		if w <= 0 {
			w = 1
		}
		if w > maxW {
			maxW = w
		}
		if gcdW == 0 {
			gcdW = w
		} else {
			gcdW = gcd(gcdW, w)
		}
	}

	for {
		lb.current = (lb.current + 1) % len(servers)
		if lb.current == 0 {
			lb.cw -= gcdW
			if lb.cw <= 0 {
				lb.cw = maxW
			}
		}
		w := servers[lb.current].Weight
		if w <= 0 {
			w = 1
		}
		if w >= lb.cw {
			return servers[lb.current]
		}
	}
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// ===========================
// Least Connections
// ===========================

// LeastConnLB selects the server with fewest active connections
type LeastConnLB struct {
	mu          sync.Mutex
	connections map[int64]*int64 // server_id -> active connections
}

func (lb *LeastConnLB) Select(servers []*UpstreamConfig, r *http.Request) *UpstreamConfig {
	if len(servers) == 0 {
		return nil
	}

	lb.mu.Lock()
	if lb.connections == nil {
		lb.connections = make(map[int64]*int64)
	}
	// Ensure all servers have a counter
	for _, s := range servers {
		if _, ok := lb.connections[s.ID]; !ok {
			var zero int64
			lb.connections[s.ID] = &zero
		}
	}
	lb.mu.Unlock()

	// Find server with minimum connections
	var selected *UpstreamConfig
	var minConn int64 = -1

	for _, s := range servers {
		lb.mu.Lock()
		conn := atomic.LoadInt64(lb.connections[s.ID])
		lb.mu.Unlock()
		if minConn == -1 || conn < minConn {
			minConn = conn
			selected = s
		}
	}

	return selected
}

// IncrementConn increments connection count for a server
func (lb *LeastConnLB) IncrementConn(serverID int64) {
	lb.mu.Lock()
	if lb.connections == nil {
		lb.connections = make(map[int64]*int64)
	}
	if _, ok := lb.connections[serverID]; !ok {
		var zero int64
		lb.connections[serverID] = &zero
	}
	counter := lb.connections[serverID]
	lb.mu.Unlock()
	atomic.AddInt64(counter, 1)
}

// DecrementConn decrements connection count for a server
func (lb *LeastConnLB) DecrementConn(serverID int64) {
	lb.mu.Lock()
	if lb.connections == nil {
		return
	}
	counter, ok := lb.connections[serverID]
	lb.mu.Unlock()
	if ok {
		atomic.AddInt64(counter, -1)
	}
}

// ===========================
// IP Hash
// ===========================

// IPHashLB consistently maps client IPs to the same server
type IPHashLB struct{}

func (lb *IPHashLB) Select(servers []*UpstreamConfig, r *http.Request) *UpstreamConfig {
	if len(servers) == 0 {
		return nil
	}

	clientIP := extractClientIP(r)
	h := fnv.New32a()
	h.Write([]byte(clientIP))
	idx := h.Sum32() % uint32(len(servers))
	return servers[idx]
}

// ===========================
// Random
// ===========================

// RandomLB selects a random server
type RandomLB struct{}

func (lb *RandomLB) Select(servers []*UpstreamConfig, r *http.Request) *UpstreamConfig {
	if len(servers) == 0 {
		return nil
	}
	return servers[rand.Intn(len(servers))]
}

// ===========================
// Failover
// ===========================

// FailoverLB selects the highest priority (lowest number) non-backup server;
// falls back to backup servers only when all primary servers are provided
type FailoverLB struct{}

func (lb *FailoverLB) Select(servers []*UpstreamConfig, r *http.Request) *UpstreamConfig {
	if len(servers) == 0 {
		return nil
	}

	// Try non-backup servers first (already sorted by priority)
	for _, s := range servers {
		if !s.IsBackup {
			return s
		}
	}
	// All are backup; pick first
	return servers[0]
}

// ===========================
// Sticky Session support
// ===========================

const stickyCookieName = "_sg_sticky"

// StickyWrapper wraps a load balancer with sticky session support
type StickyWrapper struct {
	inner LoadBalancer
}

func NewStickyWrapper(inner LoadBalancer) *StickyWrapper {
	return &StickyWrapper{inner: inner}
}

func (sw *StickyWrapper) Select(servers []*UpstreamConfig, r *http.Request) (*UpstreamConfig, bool) {
	// Check for sticky cookie
	cookie, err := r.Cookie(stickyCookieName)
	if err == nil && cookie.Value != "" {
		// Try to find the server with this ID
		for _, s := range servers {
			if formatServerID(s.ID) == cookie.Value {
				return s, false // no need to set cookie again
			}
		}
	}

	// No sticky match, use inner LB
	selected := sw.inner.Select(servers, r)
	return selected, true // need to set cookie
}

func formatServerID(id int64) string {
	return string(rune('A'+id%26)) + string(rune('0'+id/26%10))
}

// ===========================
// Helpers
// ===========================

func extractClientIP(r *http.Request) string {
	// Check X-Forwarded-For first
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		parts := splitFirst(xff, ",")
		return trimSpace(parts)
	}
	// Check X-Real-IP
	if xri := r.Header.Get("X-Real-IP"); xri != "" {
		return xri
	}
	// Use RemoteAddr
	addr := r.RemoteAddr
	if idx := lastIndexByte(addr, ':'); idx != -1 {
		return addr[:idx]
	}
	return addr
}

func splitFirst(s, sep string) string {
	idx := 0
	for idx < len(s) && string(s[idx]) != sep {
		idx++
	}
	return s[:idx]
}

func trimSpace(s string) string {
	start := 0
	for start < len(s) && s[start] == ' ' {
		start++
	}
	end := len(s)
	for end > start && s[end-1] == ' ' {
		end--
	}
	return s[start:end]
}

func lastIndexByte(s string, b byte) int {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == b {
			return i
		}
	}
	return -1
}
