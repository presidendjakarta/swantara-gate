package proxy

import (
	"database/sql"
	"log"
	"regexp"
	"strings"
	"sync"
)

// RouteMatch holds the matched route information for a request
type RouteMatch struct {
	VirtualHost      *VHostConfig
	VirtualDirectory *VDirConfig
	PathParams       map[string]string
}

// VHostConfig represents a virtual host loaded from DB
type VHostConfig struct {
	ID            int64
	HostID        int64
	VHostName     string
	LBAlgorithm   string
	StickySession bool
	FailoverMode  string
}

// VDirConfig represents a virtual directory (route) loaded from DB
type VDirConfig struct {
	ID                  int64
	VirtualHostID       int64
	SourcePath          string
	TargetPath          string
	MatchType           string
	StripPrefix         bool
	PreserveHostHeader  bool
	AuthType            string
	IsActive            bool
	ProxyTimeoutSeconds int
	RetryCount          int
	RetryDelayMs        int
	MaxRequestSizeMB    int
	WebsocketEnabled    bool
	CacheEnabled        bool
	CacheTTLSeconds     int
	AllowedMethods      []string

	// Compiled regex (only for regex match type)
	compiledRegex *regexp.Regexp
}

// UpstreamConfig represents an upstream server loaded from DB
type UpstreamConfig struct {
	ID                         int64
	VirtualHostID              int64
	TargetHost                 string
	TargetPort                 int
	Protocol                   string
	Priority                   int
	Weight                     int
	IsBackup                   bool
	IsActive                   bool
	HealthCheckEnabled         bool
	HealthCheckPath            string
	HealthCheckIntervalSeconds int
	HealthCheckTimeoutSeconds  int
	MaxFails                   int
	FailTimeoutSeconds         int
}

// Router handles route matching for incoming proxy requests
type Router struct {
	mu        sync.RWMutex
	db        *sql.DB
	hostMap   map[string]*VHostConfig   // hostname -> vhost
	routeMap  map[int64][]*VDirConfig   // vhost_id -> sorted routes
	upstreams map[int64][]*UpstreamConfig // vhost_id -> upstream servers
}

// NewRouter creates a new Router and loads configuration from the database
func NewRouter(db *sql.DB) *Router {
	r := &Router{
		db:        db,
		hostMap:   make(map[string]*VHostConfig),
		routeMap:  make(map[int64][]*VDirConfig),
		upstreams: make(map[int64][]*UpstreamConfig),
	}
	r.Reload()
	return r
}

// Reload reloads all route configuration from the database
func (r *Router) Reload() {
	hostMap := make(map[string]*VHostConfig)
	routeMap := make(map[int64][]*VDirConfig)
	upstreams := make(map[int64][]*UpstreamConfig)

	// Load active virtual hosts with their host names
	rows, err := r.db.Query(`
		SELECT vh.id, vh.host_id, h.host_name, vh.vhost_name, vh.lb_algorithm, 
		       vh.sticky_session, vh.failover_mode
		FROM virtual_hosts vh
		JOIN hosts h ON h.id = vh.host_id
		WHERE vh.is_active = 1 AND h.is_active = 1
	`)
	if err != nil {
		log.Printf("[Proxy] Error loading virtual hosts: %v", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var vh VHostConfig
		var hostName, vhostName string
		var stickySession int
		if err := rows.Scan(&vh.ID, &vh.HostID, &hostName, &vhostName,
			&vh.LBAlgorithm, &stickySession, &vh.FailoverMode); err != nil {
			log.Printf("[Proxy] Error scanning virtual host: %v", err)
			continue
		}
		vh.VHostName = vhostName
		vh.StickySession = stickySession == 1

		// Map both host_name and vhost_name to this vhost config
		hostMap[strings.ToLower(hostName)] = &vh
		hostMap[strings.ToLower(vhostName)] = &vh
	}

	// Load active virtual directories
	rows2, err := r.db.Query(`
		SELECT id, virtual_host_id, source_path, target_path, match_type,
		       strip_prefix, preserve_host_header, auth_type, proxy_timeout_seconds,
		       retry_count, retry_delay_ms, max_request_size_mb,
		       websocket_enabled, cache_enabled, cache_ttl_seconds
		FROM virtual_directories
		WHERE is_active = 1
		ORDER BY 
			CASE match_type 
				WHEN 'exact' THEN 1 
				WHEN 'regex' THEN 2
				WHEN 'parameter' THEN 3
				WHEN 'prefix' THEN 4 
				WHEN 'wildcard' THEN 5
			END,
			LENGTH(source_path) DESC
	`)
	if err != nil {
		log.Printf("[Proxy] Error loading virtual directories: %v", err)
		return
	}
	defer rows2.Close()

	for rows2.Next() {
		var vd VDirConfig
		var stripPrefix, preserveHost, wsEnabled, cacheEnabled int
		if err := rows2.Scan(&vd.ID, &vd.VirtualHostID, &vd.SourcePath, &vd.TargetPath,
			&vd.MatchType, &stripPrefix, &preserveHost, &vd.AuthType,
			&vd.ProxyTimeoutSeconds, &vd.RetryCount, &vd.RetryDelayMs,
			&vd.MaxRequestSizeMB, &wsEnabled, &cacheEnabled, &vd.CacheTTLSeconds); err != nil {
			log.Printf("[Proxy] Error scanning virtual directory: %v", err)
			continue
		}
		vd.StripPrefix = stripPrefix == 1
		vd.PreserveHostHeader = preserveHost == 1
		vd.WebsocketEnabled = wsEnabled == 1
		vd.CacheEnabled = cacheEnabled == 1
		vd.IsActive = true

		// Compile regex if needed
		if vd.MatchType == "regex" {
			compiled, err := regexp.Compile(vd.SourcePath)
			if err != nil {
				log.Printf("[Proxy] Error compiling regex for route %s: %v", vd.SourcePath, err)
				continue
			}
			vd.compiledRegex = compiled
		}

		routeMap[vd.VirtualHostID] = append(routeMap[vd.VirtualHostID], &vd)
	}

	// Load allowed methods for each virtual directory
	rows3, err := r.db.Query(`SELECT virtual_directory_id, http_method FROM virtual_directory_methods`)
	if err != nil {
		log.Printf("[Proxy] Error loading methods: %v", err)
	} else {
		defer rows3.Close()
		methodMap := make(map[int64][]string)
		for rows3.Next() {
			var vdID int64
			var method string
			if err := rows3.Scan(&vdID, &method); err == nil {
				methodMap[vdID] = append(methodMap[vdID], strings.ToUpper(method))
			}
		}
		// Assign methods to routes
		for _, routes := range routeMap {
			for _, route := range routes {
				if methods, ok := methodMap[route.ID]; ok {
					route.AllowedMethods = methods
				}
			}
		}
	}

	// Load active upstream servers
	rows4, err := r.db.Query(`
		SELECT id, virtual_host_id, target_host, target_port, protocol,
		       priority, weight, is_backup, is_active,
		       health_check_enabled, health_check_path,
		       health_check_interval_seconds, health_check_timeout_seconds,
		       max_fails, fail_timeout_seconds
		FROM upstream_servers
		WHERE is_active = 1
		ORDER BY priority ASC, weight DESC
	`)
	if err != nil {
		log.Printf("[Proxy] Error loading upstream servers: %v", err)
		return
	}
	defer rows4.Close()

	for rows4.Next() {
		var us UpstreamConfig
		var isBackup, isActive, hcEnabled int
		if err := rows4.Scan(&us.ID, &us.VirtualHostID, &us.TargetHost, &us.TargetPort,
			&us.Protocol, &us.Priority, &us.Weight, &isBackup, &isActive,
			&hcEnabled, &us.HealthCheckPath, &us.HealthCheckIntervalSeconds,
			&us.HealthCheckTimeoutSeconds, &us.MaxFails, &us.FailTimeoutSeconds); err != nil {
			log.Printf("[Proxy] Error scanning upstream server: %v", err)
			continue
		}
		us.IsBackup = isBackup == 1
		us.IsActive = isActive == 1
		us.HealthCheckEnabled = hcEnabled == 1

		upstreams[us.VirtualHostID] = append(upstreams[us.VirtualHostID], &us)
	}

	// Swap in new config atomically
	r.mu.Lock()
	r.hostMap = hostMap
	r.routeMap = routeMap
	r.upstreams = upstreams
	r.mu.Unlock()

	log.Printf("[Proxy] Config loaded: %d hosts, %d route groups, %d upstream groups",
		len(hostMap), len(routeMap), len(upstreams))
}

// Match finds the matching virtual host and route for a given hostname and path
func (r *Router) Match(hostname, path string) *RouteMatch {
	r.mu.RLock()
	defer r.mu.RUnlock()

	// Strip port from hostname
	if idx := strings.LastIndex(hostname, ":"); idx != -1 {
		hostname = hostname[:idx]
	}
	hostname = strings.ToLower(hostname)

	vhost, ok := r.hostMap[hostname]
	if !ok {
		return nil
	}

	routes, ok := r.routeMap[vhost.ID]
	if !ok {
		return nil
	}

	for _, route := range routes {
		if params, matched := matchRoute(route, path); matched {
			return &RouteMatch{
				VirtualHost:      vhost,
				VirtualDirectory: route,
				PathParams:       params,
			}
		}
	}

	return nil
}

// GetUpstreams returns upstream servers for a virtual host
func (r *Router) GetUpstreams(vhostID int64) []*UpstreamConfig {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.upstreams[vhostID]
}

// matchRoute checks if a path matches a route based on its match type
func matchRoute(route *VDirConfig, path string) (map[string]string, bool) {
	switch route.MatchType {
	case "exact":
		if path == route.SourcePath {
			return nil, true
		}
	case "prefix":
		if strings.HasPrefix(path, route.SourcePath) {
			return nil, true
		}
	case "wildcard":
		if matchWildcard(route.SourcePath, path) {
			return nil, true
		}
	case "regex":
		if route.compiledRegex != nil && route.compiledRegex.MatchString(path) {
			return nil, true
		}
	case "parameter":
		if params, ok := matchParameterPath(route.SourcePath, path); ok {
			return params, true
		}
	default:
		// Default to prefix matching
		if strings.HasPrefix(path, route.SourcePath) {
			return nil, true
		}
	}
	return nil, false
}

// matchWildcard matches a path against a wildcard pattern
// Supports * (any segment) and ** (any number of segments)
func matchWildcard(pattern, path string) bool {
	patternParts := strings.Split(strings.Trim(pattern, "/"), "/")
	pathParts := strings.Split(strings.Trim(path, "/"), "/")

	return matchWildcardParts(patternParts, pathParts)
}

func matchWildcardParts(pattern, path []string) bool {
	pi, pj := 0, 0
	for pi < len(pattern) && pj < len(path) {
		if pattern[pi] == "**" {
			// ** matches zero or more segments
			if pi == len(pattern)-1 {
				return true
			}
			// Try matching the rest of pattern against remaining path
			for k := pj; k <= len(path); k++ {
				if matchWildcardParts(pattern[pi+1:], path[k:]) {
					return true
				}
			}
			return false
		} else if pattern[pi] == "*" {
			// * matches exactly one segment
			pi++
			pj++
		} else if pattern[pi] == path[pj] {
			pi++
			pj++
		} else {
			return false
		}
	}
	// Handle trailing **
	for pi < len(pattern) && pattern[pi] == "**" {
		pi++
	}
	return pi == len(pattern) && pj == len(path)
}

// matchParameterPath matches a path against a parameterized pattern like /api/users/{id}/posts/{post_id}
func matchParameterPath(pattern, path string) (map[string]string, bool) {
	patternParts := strings.Split(strings.Trim(pattern, "/"), "/")
	pathParts := strings.Split(strings.Trim(path, "/"), "/")

	if len(patternParts) != len(pathParts) {
		return nil, false
	}

	params := make(map[string]string)
	for i, pp := range patternParts {
		if strings.HasPrefix(pp, "{") && strings.HasSuffix(pp, "}") {
			paramName := pp[1 : len(pp)-1]
			params[paramName] = pathParts[i]
		} else if pp != pathParts[i] {
			return nil, false
		}
	}
	return params, true
}
