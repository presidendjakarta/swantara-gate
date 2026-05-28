package proxy

import (
	"database/sql"
	"log"
	"sync"
	"time"
)

// CircuitState represents the state of a circuit breaker
type CircuitState int

const (
	CircuitClosed   CircuitState = iota // Normal operation
	CircuitOpen                         // Blocking requests
	CircuitHalfOpen                     // Testing recovery
)

// CircuitBreakerConfig holds config for a route's circuit breaker
type CircuitBreakerConfig struct {
	Enabled                bool
	FailureThreshold       int
	RecoveryTimeoutSeconds int
	HalfOpenMaxRequests    int
}

// circuitBreakerState tracks state of a single circuit breaker
type circuitBreakerState struct {
	State           CircuitState
	FailureCount    int
	SuccessCount    int
	LastFailureAt   time.Time
	HalfOpenCount   int
}

// ProxyCircuitBreaker manages circuit breakers per route
type ProxyCircuitBreaker struct {
	mu      sync.RWMutex
	db      *sql.DB
	configs map[int64]*CircuitBreakerConfig // vdir_id -> config
	states  map[int64]*circuitBreakerState  // vdir_id -> state
}

// NewProxyCircuitBreaker creates a new circuit breaker manager
func NewProxyCircuitBreaker(db *sql.DB) *ProxyCircuitBreaker {
	cb := &ProxyCircuitBreaker{
		db:      db,
		configs: make(map[int64]*CircuitBreakerConfig),
		states:  make(map[int64]*circuitBreakerState),
	}
	cb.Reload()
	return cb
}

// Reload loads circuit breaker configs from database
func (cb *ProxyCircuitBreaker) Reload() {
	configs := make(map[int64]*CircuitBreakerConfig)

	rows, err := cb.db.Query(`
		SELECT virtual_directory_id, enabled, failure_threshold,
		       recovery_timeout_seconds, half_open_max_requests
		FROM circuit_breakers
	`)
	if err != nil {
		log.Printf("[Proxy:CircuitBreaker] Error loading configs: %v", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var vdirID int64
		var cfg CircuitBreakerConfig
		var enabled int
		if err := rows.Scan(&vdirID, &enabled, &cfg.FailureThreshold,
			&cfg.RecoveryTimeoutSeconds, &cfg.HalfOpenMaxRequests); err != nil {
			log.Printf("[Proxy:CircuitBreaker] Error scanning config: %v", err)
			continue
		}
		cfg.Enabled = enabled == 1
		configs[vdirID] = &cfg
	}

	cb.mu.Lock()
	cb.configs = configs
	cb.mu.Unlock()

	log.Printf("[Proxy:CircuitBreaker] Loaded %d circuit breaker configs", len(configs))
}

// Allow checks if a request should be allowed through the circuit breaker
func (cb *ProxyCircuitBreaker) Allow(vdirID int64) (bool, string) {
	cb.mu.RLock()
	cfg, ok := cb.configs[vdirID]
	cb.mu.RUnlock()

	if !ok || !cfg.Enabled {
		return true, ""
	}

	cb.mu.Lock()
	defer cb.mu.Unlock()

	state, exists := cb.states[vdirID]
	if !exists {
		cb.states[vdirID] = &circuitBreakerState{State: CircuitClosed}
		return true, ""
	}

	switch state.State {
	case CircuitClosed:
		return true, ""

	case CircuitOpen:
		// Check if recovery timeout has passed
		recoveryTimeout := time.Duration(cfg.RecoveryTimeoutSeconds) * time.Second
		if time.Since(state.LastFailureAt) > recoveryTimeout {
			// Transition to half-open
			state.State = CircuitHalfOpen
			state.HalfOpenCount = 0
			state.SuccessCount = 0
			return true, ""
		}
		return false, "circuit breaker open: service temporarily unavailable"

	case CircuitHalfOpen:
		if state.HalfOpenCount >= cfg.HalfOpenMaxRequests {
			return false, "circuit breaker half-open: max test requests reached"
		}
		state.HalfOpenCount++
		return true, ""
	}

	return true, ""
}

// RecordSuccess records a successful request for the circuit breaker
func (cb *ProxyCircuitBreaker) RecordSuccess(vdirID int64) {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	state, exists := cb.states[vdirID]
	if !exists {
		return
	}

	cfg := cb.configs[vdirID]
	if cfg == nil || !cfg.Enabled {
		return
	}

	switch state.State {
	case CircuitClosed:
		// Reset failure count on success
		state.FailureCount = 0
	case CircuitHalfOpen:
		state.SuccessCount++
		// If enough successes, close the circuit
		if state.SuccessCount >= cfg.HalfOpenMaxRequests {
			state.State = CircuitClosed
			state.FailureCount = 0
			log.Printf("[Proxy:CircuitBreaker] Route %d: circuit CLOSED (recovered)", vdirID)
		}
	}
}

// RecordFailure records a failed request for the circuit breaker
func (cb *ProxyCircuitBreaker) RecordFailure(vdirID int64) {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	cfg, ok := cb.configs[vdirID]
	if !ok || !cfg.Enabled {
		return
	}

	state, exists := cb.states[vdirID]
	if !exists {
		state = &circuitBreakerState{State: CircuitClosed}
		cb.states[vdirID] = state
	}

	switch state.State {
	case CircuitClosed:
		state.FailureCount++
		state.LastFailureAt = time.Now()
		if state.FailureCount >= cfg.FailureThreshold {
			state.State = CircuitOpen
			log.Printf("[Proxy:CircuitBreaker] Route %d: circuit OPEN (failures: %d/%d)",
				vdirID, state.FailureCount, cfg.FailureThreshold)
		}
	case CircuitHalfOpen:
		// Failure during half-open -> back to open
		state.State = CircuitOpen
		state.LastFailureAt = time.Now()
		log.Printf("[Proxy:CircuitBreaker] Route %d: circuit OPEN (failed during half-open)", vdirID)
	}
}
