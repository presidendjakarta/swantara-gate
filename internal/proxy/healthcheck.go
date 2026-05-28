package proxy

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

// HealthStatus tracks the health of an upstream server
type HealthStatus struct {
	Healthy       bool
	FailCount     int
	LastCheckAt   time.Time
	LastFailAt    time.Time
	LastSuccessAt time.Time
}

// HealthChecker performs background health checks on upstream servers
type HealthChecker struct {
	mu       sync.RWMutex
	statuses map[int64]*HealthStatus // server_id -> status
	client   *http.Client
	cancel   context.CancelFunc
	wg       sync.WaitGroup
}

// NewHealthChecker creates a new HealthChecker
func NewHealthChecker() *HealthChecker {
	return &HealthChecker{
		statuses: make(map[int64]*HealthStatus),
		client: &http.Client{
			Timeout: 5 * time.Second,
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		},
	}
}

// Start begins background health check goroutines for the given upstreams
func (hc *HealthChecker) Start(upstreams map[int64][]*UpstreamConfig) {
	ctx, cancel := context.WithCancel(context.Background())
	hc.cancel = cancel

	// Initialize all statuses as healthy
	hc.mu.Lock()
	for _, servers := range upstreams {
		for _, s := range servers {
			if _, exists := hc.statuses[s.ID]; !exists {
				hc.statuses[s.ID] = &HealthStatus{Healthy: true}
			}
		}
	}
	hc.mu.Unlock()

	// Start a goroutine for each server that has health checks enabled
	for _, servers := range upstreams {
		for _, server := range servers {
			if !server.HealthCheckEnabled {
				// Mark as always healthy
				hc.mu.Lock()
				hc.statuses[server.ID] = &HealthStatus{Healthy: true}
				hc.mu.Unlock()
				continue
			}

			s := server // capture
			hc.wg.Add(1)
			go hc.checkLoop(ctx, s)
		}
	}
}

// Stop stops all health check goroutines
func (hc *HealthChecker) Stop() {
	if hc.cancel != nil {
		hc.cancel()
	}
	hc.wg.Wait()
}

// IsHealthy returns whether a server is currently healthy
func (hc *HealthChecker) IsHealthy(serverID int64) bool {
	hc.mu.RLock()
	defer hc.mu.RUnlock()
	status, ok := hc.statuses[serverID]
	if !ok {
		return true // unknown servers assumed healthy
	}
	return status.Healthy
}

// GetHealthyServers filters a list to only healthy servers
func (hc *HealthChecker) GetHealthyServers(servers []*UpstreamConfig) []*UpstreamConfig {
	var healthy []*UpstreamConfig
	hc.mu.RLock()
	defer hc.mu.RUnlock()

	for _, s := range servers {
		status, ok := hc.statuses[s.ID]
		if !ok || status.Healthy {
			healthy = append(healthy, s)
		}
	}
	return healthy
}

// MarkFailed manually marks a server as failed (e.g., after proxy error)
func (hc *HealthChecker) MarkFailed(server *UpstreamConfig) {
	hc.mu.Lock()
	defer hc.mu.Unlock()

	status, ok := hc.statuses[server.ID]
	if !ok {
		status = &HealthStatus{Healthy: true}
		hc.statuses[server.ID] = status
	}

	status.FailCount++
	status.LastFailAt = time.Now()

	if status.FailCount >= server.MaxFails {
		status.Healthy = false
		log.Printf("[HealthCheck] Server %s:%d marked UNHEALTHY (fails: %d/%d)",
			server.TargetHost, server.TargetPort, status.FailCount, server.MaxFails)
	}
}

// MarkSuccess manually marks a server as successful
func (hc *HealthChecker) MarkSuccess(server *UpstreamConfig) {
	hc.mu.Lock()
	defer hc.mu.Unlock()

	status, ok := hc.statuses[server.ID]
	if !ok {
		status = &HealthStatus{Healthy: true}
		hc.statuses[server.ID] = status
	}

	status.FailCount = 0
	status.Healthy = true
	status.LastSuccessAt = time.Now()
}

// checkLoop is the main health check loop for a single server
func (hc *HealthChecker) checkLoop(ctx context.Context, server *UpstreamConfig) {
	defer hc.wg.Done()

	interval := time.Duration(server.HealthCheckIntervalSeconds) * time.Second
	if interval <= 0 {
		interval = 10 * time.Second
	}

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			hc.doCheck(server)
		}
	}
}

// doCheck performs a single health check against a server
func (hc *HealthChecker) doCheck(server *UpstreamConfig) {
	url := fmt.Sprintf("%s://%s:%d%s", server.Protocol, server.TargetHost, server.TargetPort, server.HealthCheckPath)

	timeout := time.Duration(server.HealthCheckTimeoutSeconds) * time.Second
	if timeout <= 0 {
		timeout = 3 * time.Second
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		hc.MarkFailed(server)
		return
	}
	req.Header.Set("User-Agent", "SwantaraGate-HealthCheck/1.0")

	resp, err := hc.client.Do(req)
	if err != nil {
		hc.MarkFailed(server)
		return
	}
	resp.Body.Close()

	hc.mu.Lock()
	status, ok := hc.statuses[server.ID]
	if !ok {
		status = &HealthStatus{Healthy: true}
		hc.statuses[server.ID] = status
	}
	status.LastCheckAt = time.Now()
	hc.mu.Unlock()

	if resp.StatusCode >= 200 && resp.StatusCode < 400 {
		hc.MarkSuccess(server)
	} else {
		hc.MarkFailed(server)
	}
}

// RecoverAfterTimeout checks if unhealthy servers should be recovered
// (called periodically or can be incorporated into checkLoop)
func (hc *HealthChecker) RecoverAfterTimeout(servers []*UpstreamConfig) {
	hc.mu.Lock()
	defer hc.mu.Unlock()

	now := time.Now()
	for _, s := range servers {
		status, ok := hc.statuses[s.ID]
		if !ok || status.Healthy {
			continue
		}
		failTimeout := time.Duration(s.FailTimeoutSeconds) * time.Second
		if failTimeout > 0 && now.Sub(status.LastFailAt) > failTimeout {
			status.Healthy = true
			status.FailCount = 0
			log.Printf("[HealthCheck] Server %s:%d recovered after fail_timeout",
				s.TargetHost, s.TargetPort)
		}
	}
}
