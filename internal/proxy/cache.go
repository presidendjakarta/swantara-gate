package proxy

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"sync"
	"time"
)

// CacheEntry represents a cached response
type CacheEntry struct {
	StatusCode int
	Header     http.Header
	Body       []byte
	ExpiresAt  time.Time
}

// ResponseCache provides in-memory response caching with TTL
type ResponseCache struct {
	mu      sync.RWMutex
	entries map[string]*CacheEntry
}

// NewResponseCache creates a new ResponseCache
func NewResponseCache() *ResponseCache {
	rc := &ResponseCache{
		entries: make(map[string]*CacheEntry),
	}
	// Start background cleanup goroutine
	go rc.cleanup()
	return rc
}

// Get retrieves a cached response if it exists and hasn't expired
func (rc *ResponseCache) Get(r *http.Request) *CacheEntry {
	// Only cache GET requests
	if r.Method != http.MethodGet {
		return nil
	}

	key := rc.buildKey(r)

	rc.mu.RLock()
	entry, ok := rc.entries[key]
	rc.mu.RUnlock()

	if !ok {
		return nil
	}

	// Check if expired
	if time.Now().After(entry.ExpiresAt) {
		rc.mu.Lock()
		delete(rc.entries, key)
		rc.mu.Unlock()
		return nil
	}

	return entry
}

// Set stores a response in the cache with the given TTL
func (rc *ResponseCache) Set(r *http.Request, resp *ProxyResponse, ttlSeconds int) {
	// Only cache GET requests with 2xx responses
	if r.Method != http.MethodGet {
		return
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return
	}
	if ttlSeconds <= 0 {
		return
	}

	key := rc.buildKey(r)

	// Copy headers to avoid reference issues
	header := make(http.Header)
	for k, vv := range resp.Header {
		header[k] = append([]string{}, vv...)
	}

	// Copy body
	body := make([]byte, len(resp.Body))
	copy(body, resp.Body)

	entry := &CacheEntry{
		StatusCode: resp.StatusCode,
		Header:     header,
		Body:       body,
		ExpiresAt:  time.Now().Add(time.Duration(ttlSeconds) * time.Second),
	}

	rc.mu.Lock()
	rc.entries[key] = entry
	rc.mu.Unlock()
}

// Invalidate removes a specific cache entry
func (rc *ResponseCache) Invalidate(r *http.Request) {
	key := rc.buildKey(r)
	rc.mu.Lock()
	delete(rc.entries, key)
	rc.mu.Unlock()
}

// Clear removes all cache entries
func (rc *ResponseCache) Clear() {
	rc.mu.Lock()
	rc.entries = make(map[string]*CacheEntry)
	rc.mu.Unlock()
}

// buildKey creates a cache key from the request
func (rc *ResponseCache) buildKey(r *http.Request) string {
	raw := r.Method + "|" + r.Host + "|" + r.URL.Path + "|" + r.URL.RawQuery
	h := sha256.Sum256([]byte(raw))
	return hex.EncodeToString(h[:])
}

// cleanup periodically removes expired entries
func (rc *ResponseCache) cleanup() {
	ticker := time.NewTicker(60 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		rc.mu.Lock()
		now := time.Now()
		for key, entry := range rc.entries {
			if now.After(entry.ExpiresAt) {
				delete(rc.entries, key)
			}
		}
		rc.mu.Unlock()
	}
}
