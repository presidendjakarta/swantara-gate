package proxy

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

// RateLimitConfig holds rate limit settings for a route
type RateLimitConfig struct {
	LimitBy              string // ip, api_key, consumer
	RequestsPerMinute    int
	Burst                int
	BlockDurationSeconds int
}

// rateLimitEntry tracks request counts per key
type rateLimitEntry struct {
	Count     int
	WindowEnd time.Time
	BlockedAt *time.Time
}

// ProxyRateLimiter enforces rate limits per route
type ProxyRateLimiter struct {
	mu      sync.RWMutex
	db      *sql.DB
	configs map[int64]*RateLimitConfig // vdir_id -> config
	entries map[string]*rateLimitEntry // "vdir_id:key" -> entry
}

// NewProxyRateLimiter creates a new rate limiter
func NewProxyRateLimiter(db *sql.DB) *ProxyRateLimiter {
	rl := &ProxyRateLimiter{
		db:      db,
		configs: make(map[int64]*RateLimitConfig),
		entries: make(map[string]*rateLimitEntry),
	}
	rl.Reload()
	go rl.cleanup()
	return rl
}

// Reload loads rate limit configs from database
func (rl *ProxyRateLimiter) Reload() {
	configs := make(map[int64]*RateLimitConfig)

	rows, err := rl.db.Query(`
		SELECT virtual_directory_id, COALESCE(limit_by, 'ip'), 
		       requests_per_minute, burst, block_duration_seconds
		FROM rate_limits WHERE is_active = 1
	`)
	if err != nil {
		log.Printf("[Proxy:RateLimit] Error loading configs: %v", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var vdirID int64
		var cfg RateLimitConfig
		if err := rows.Scan(&vdirID, &cfg.LimitBy, &cfg.RequestsPerMinute,
			&cfg.Burst, &cfg.BlockDurationSeconds); err != nil {
			log.Printf("[Proxy:RateLimit] Error scanning config: %v", err)
			continue
		}
		configs[vdirID] = &cfg
	}

	rl.mu.Lock()
	rl.configs = configs
	rl.mu.Unlock()

	log.Printf("[Proxy:RateLimit] Loaded %d rate limit configs", len(configs))
}

// Allow checks if a request is allowed or should be rate limited
func (rl *ProxyRateLimiter) Allow(r *http.Request, route *VDirConfig) (bool, string) {
	rl.mu.RLock()
	cfg, ok := rl.configs[route.ID]
	rl.mu.RUnlock()

	if !ok {
		return true, "" // No rate limit configured
	}

	// Determine the key based on limit_by strategy
	key := rl.getKey(r, route.ID, cfg.LimitBy)

	rl.mu.Lock()
	defer rl.mu.Unlock()

	entry, exists := rl.entries[key]
	now := time.Now()

	// Check if currently blocked
	if exists && entry.BlockedAt != nil {
		blockEnd := entry.BlockedAt.Add(time.Duration(cfg.BlockDurationSeconds) * time.Second)
		if now.Before(blockEnd) {
			return false, "rate limit exceeded, try again later"
		}
		// Block expired, reset
		entry.BlockedAt = nil
		entry.Count = 0
		entry.WindowEnd = now.Add(time.Minute)
	}

	if !exists || now.After(entry.WindowEnd) {
		// New window
		rl.entries[key] = &rateLimitEntry{
			Count:     1,
			WindowEnd: now.Add(time.Minute),
		}
		return true, ""
	}

	entry.Count++

	// Check burst limit (immediate overflow)
	maxAllowed := cfg.RequestsPerMinute + cfg.Burst
	if entry.Count > maxAllowed {
		blockedAt := now
		entry.BlockedAt = &blockedAt
		return false, "rate limit exceeded"
	}

	// Check requests per minute
	if entry.Count > cfg.RequestsPerMinute {
		// Over the base limit but within burst
		return true, ""
	}

	return true, ""
}

// getKey builds the rate limit key
func (rl *ProxyRateLimiter) getKey(r *http.Request, vdirID int64, limitBy string) string {
	var identifier string
	switch limitBy {
	case "api_key":
		identifier = r.Header.Get("X-API-Key")
		if identifier == "" {
			identifier = r.URL.Query().Get("api_key")
		}
		if identifier == "" {
			identifier = extractClientIP(r)
		}
	case "consumer":
		// Fall back to IP if no consumer info
		identifier = extractClientIP(r)
	default: // "ip"
		identifier = extractClientIP(r)
	}
	return fmt.Sprintf("%d:%s", vdirID, identifier)
}

// cleanup periodically removes expired entries
func (rl *ProxyRateLimiter) cleanup() {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		rl.mu.Lock()
		now := time.Now()
		for key, entry := range rl.entries {
			if now.After(entry.WindowEnd) && entry.BlockedAt == nil {
				delete(rl.entries, key)
			}
		}
		rl.mu.Unlock()
	}
}

