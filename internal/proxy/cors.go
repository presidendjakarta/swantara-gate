package proxy

import (
	"database/sql"
	"log"
	"net/http"
	"strings"
	"sync"
)

// CORSRouteConfig holds CORS settings for a route
type CORSRouteConfig struct {
	AllowedOrigins   []string
	AllowedMethods   []string
	AllowedHeaders   []string
	ExposedHeaders   []string
	AllowCredentials bool
	MaxAgeSeconds    int
}

// ProxyCORS handles per-route CORS configuration
type ProxyCORS struct {
	mu      sync.RWMutex
	db      *sql.DB
	configs map[int64]*CORSRouteConfig // vdir_id -> config
}

// NewProxyCORS creates a new CORS handler
func NewProxyCORS(db *sql.DB) *ProxyCORS {
	c := &ProxyCORS{
		db:      db,
		configs: make(map[int64]*CORSRouteConfig),
	}
	c.Reload()
	return c
}

// Reload loads CORS configs from database
func (c *ProxyCORS) Reload() {
	configs := make(map[int64]*CORSRouteConfig)

	rows, err := c.db.Query(`
		SELECT virtual_directory_id, COALESCE(allowed_origins, '*'),
		       COALESCE(allowed_methods, 'GET,POST,PUT,DELETE,OPTIONS'),
		       COALESCE(allowed_headers, '*'), COALESCE(exposed_headers, ''),
		       allow_credentials, max_age_seconds
		FROM cors_configs WHERE is_active = 1
	`)
	if err != nil {
		log.Printf("[Proxy:CORS] Error loading configs: %v", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var vdirID int64
		var origins, methods, headers, exposed string
		var allowCreds int
		var maxAge int
		if err := rows.Scan(&vdirID, &origins, &methods, &headers,
			&exposed, &allowCreds, &maxAge); err != nil {
			log.Printf("[Proxy:CORS] Error scanning config: %v", err)
			continue
		}
		cfg := &CORSRouteConfig{
			AllowedOrigins:   splitCSV(origins),
			AllowedMethods:   splitCSV(methods),
			AllowedHeaders:   splitCSV(headers),
			ExposedHeaders:   splitCSV(exposed),
			AllowCredentials: allowCreds == 1,
			MaxAgeSeconds:    maxAge,
		}
		configs[vdirID] = cfg
	}

	c.mu.Lock()
	c.configs = configs
	c.mu.Unlock()

	log.Printf("[Proxy:CORS] Loaded %d CORS configs", len(configs))
}

// HandlePreflight checks if the request is a CORS preflight and handles it
// Returns true if the request was handled (caller should return)
func (c *ProxyCORS) HandlePreflight(w http.ResponseWriter, r *http.Request, route *VDirConfig) bool {
	c.mu.RLock()
	cfg, ok := c.configs[route.ID]
	c.mu.RUnlock()

	if !ok {
		return false // No CORS config, let it through
	}

	origin := r.Header.Get("Origin")
	if origin == "" {
		return false // Not a CORS request
	}

	// Check if origin is allowed
	if !c.isOriginAllowed(origin, cfg.AllowedOrigins) {
		return false
	}

	// Set CORS headers
	c.setCORSHeaders(w, origin, cfg)

	// Handle preflight OPTIONS request
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return true
	}

	return false
}

// ApplyCORSHeaders adds CORS headers to the response for non-preflight requests
func (c *ProxyCORS) ApplyCORSHeaders(w http.ResponseWriter, r *http.Request, route *VDirConfig) {
	c.mu.RLock()
	cfg, ok := c.configs[route.ID]
	c.mu.RUnlock()

	if !ok {
		return
	}

	origin := r.Header.Get("Origin")
	if origin == "" {
		return
	}

	if !c.isOriginAllowed(origin, cfg.AllowedOrigins) {
		return
	}

	c.setCORSHeaders(w, origin, cfg)
}

// setCORSHeaders sets the actual CORS response headers
func (c *ProxyCORS) setCORSHeaders(w http.ResponseWriter, origin string, cfg *CORSRouteConfig) {
	// Allow-Origin
	if containsStar(cfg.AllowedOrigins) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
	} else {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Vary", "Origin")
	}

	// Allow-Methods
	if len(cfg.AllowedMethods) > 0 {
		w.Header().Set("Access-Control-Allow-Methods", strings.Join(cfg.AllowedMethods, ", "))
	}

	// Allow-Headers
	if len(cfg.AllowedHeaders) > 0 {
		w.Header().Set("Access-Control-Allow-Headers", strings.Join(cfg.AllowedHeaders, ", "))
	}

	// Exposed-Headers
	if len(cfg.ExposedHeaders) > 0 {
		joined := strings.Join(cfg.ExposedHeaders, ", ")
		if joined != "" {
			w.Header().Set("Access-Control-Expose-Headers", joined)
		}
	}

	// Allow-Credentials
	if cfg.AllowCredentials {
		w.Header().Set("Access-Control-Allow-Credentials", "true")
	}

	// Max-Age
	if cfg.MaxAgeSeconds > 0 {
		w.Header().Set("Access-Control-Max-Age", strings.TrimSpace(
			strings.Replace(strings.Replace(
				string(rune(cfg.MaxAgeSeconds+'0')), "\x00", "", -1), " ", "", -1)))
		// Use fmt for proper integer conversion
		w.Header().Set("Access-Control-Max-Age", formatInt(cfg.MaxAgeSeconds))
	}
}

// isOriginAllowed checks if the origin is in the allowed list
func (c *ProxyCORS) isOriginAllowed(origin string, allowed []string) bool {
	if containsStar(allowed) {
		return true
	}
	for _, o := range allowed {
		if strings.EqualFold(strings.TrimSpace(o), origin) {
			return true
		}
	}
	return false
}

func containsStar(list []string) bool {
	for _, item := range list {
		if strings.TrimSpace(item) == "*" {
			return true
		}
	}
	return false
}

func splitCSV(s string) []string {
	if s == "" {
		return nil
	}
	parts := strings.Split(s, ",")
	result := make([]string, 0, len(parts))
	for _, p := range parts {
		trimmed := strings.TrimSpace(p)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
}

func formatInt(n int) string {
	if n == 0 {
		return "0"
	}
	result := ""
	negative := n < 0
	if negative {
		n = -n
	}
	for n > 0 {
		result = string(rune('0'+n%10)) + result
		n /= 10
	}
	if negative {
		result = "-" + result
	}
	return result
}
