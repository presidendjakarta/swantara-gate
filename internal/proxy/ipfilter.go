package proxy

import (
	"database/sql"
	"log"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"
)

// IPFilterConfig holds IP filtering rules for a route
type IPFilterConfig struct {
	Whitelist []string
	Blacklist []ipBlacklistEntry
}

type ipBlacklistEntry struct {
	IPAddress string
	ExpiredAt *time.Time
}

// IPFilter manages IP whitelist/blacklist enforcement
type IPFilter struct {
	mu      sync.RWMutex
	db      *sql.DB
	configs map[int64]*IPFilterConfig // vdir_id -> config
}

// NewIPFilter creates a new IP filter
func NewIPFilter(db *sql.DB) *IPFilter {
	f := &IPFilter{
		db:      db,
		configs: make(map[int64]*IPFilterConfig),
	}
	f.Reload()
	return f
}

// Reload loads IP filter configs from database
func (f *IPFilter) Reload() {
	configs := make(map[int64]*IPFilterConfig)

	// Load whitelists
	rows, err := f.db.Query(`
		SELECT virtual_directory_id, ip_address
		FROM ip_whitelists WHERE is_active = 1
	`)
	if err != nil {
		log.Printf("[Proxy:IPFilter] Error loading whitelists: %v", err)
	} else {
		defer rows.Close()
		for rows.Next() {
			var vdirID int64
			var ip string
			if err := rows.Scan(&vdirID, &ip); err != nil {
				continue
			}
			if _, ok := configs[vdirID]; !ok {
				configs[vdirID] = &IPFilterConfig{}
			}
			configs[vdirID].Whitelist = append(configs[vdirID].Whitelist, ip)
		}
	}

	// Load blacklists
	rows2, err := f.db.Query(`
		SELECT virtual_directory_id, ip_address, expired_at
		FROM ip_blacklists WHERE is_active = 1
	`)
	if err != nil {
		log.Printf("[Proxy:IPFilter] Error loading blacklists: %v", err)
	} else {
		defer rows2.Close()
		for rows2.Next() {
			var vdirID int64
			var ip string
			var expiredAt *string
			if err := rows2.Scan(&vdirID, &ip, &expiredAt); err != nil {
				continue
			}
			if _, ok := configs[vdirID]; !ok {
				configs[vdirID] = &IPFilterConfig{}
			}
			entry := ipBlacklistEntry{IPAddress: ip}
			if expiredAt != nil && *expiredAt != "" {
				if t, err := time.Parse("2006-01-02 15:04:05", *expiredAt); err == nil {
					entry.ExpiredAt = &t
				}
			}
			configs[vdirID].Blacklist = append(configs[vdirID].Blacklist, entry)
		}
	}

	f.mu.Lock()
	f.configs = configs
	f.mu.Unlock()

	log.Printf("[Proxy:IPFilter] Loaded IP filters for %d routes", len(configs))
}

// Allow checks if the client IP is allowed for the route
func (f *IPFilter) Allow(r *http.Request, route *VDirConfig) (bool, string) {
	f.mu.RLock()
	cfg, ok := f.configs[route.ID]
	f.mu.RUnlock()

	if !ok {
		return true, "" // No IP filter configured
	}

	clientIP := extractClientIP(r)

	// Check whitelist first (if whitelist exists, only whitelisted IPs are allowed)
	if len(cfg.Whitelist) > 0 {
		if !f.ipMatchesAny(clientIP, cfg.Whitelist) {
			return false, "IP not whitelisted"
		}
		return true, ""
	}

	// Check blacklist
	if len(cfg.Blacklist) > 0 {
		now := time.Now()
		for _, entry := range cfg.Blacklist {
			if f.ipMatches(clientIP, entry.IPAddress) {
				// Check if blacklist entry has expired
				if entry.ExpiredAt != nil && now.After(*entry.ExpiredAt) {
					continue // Expired, skip
				}
				return false, "IP blacklisted"
			}
		}
	}

	return true, ""
}

// ipMatchesAny checks if clientIP matches any in the list (supports CIDR)
func (f *IPFilter) ipMatchesAny(clientIP string, list []string) bool {
	for _, entry := range list {
		if f.ipMatches(clientIP, entry) {
			return true
		}
	}
	return false
}

// ipMatches checks if clientIP matches an entry (exact or CIDR)
func (f *IPFilter) ipMatches(clientIP, entry string) bool {
	// Exact match
	if clientIP == entry {
		return true
	}

	// CIDR match
	if strings.Contains(entry, "/") {
		_, network, err := net.ParseCIDR(entry)
		if err != nil {
			return false
		}
		ip := net.ParseIP(clientIP)
		if ip != nil && network.Contains(ip) {
			return true
		}
	}

	return false
}
