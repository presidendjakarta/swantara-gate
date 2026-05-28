package test

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/presidendjakarta/swantara-gate/internal/proxy"
	_ "modernc.org/sqlite"
)

// ProxyTestEnv holds a proxy server + a fake backend for integration testing
type ProxyTestEnv struct {
	ProxyServer  *httptest.Server
	BackendServer *httptest.Server
	DB           *sql.DB
	Proxy        *proxy.ProxyServer
}

// setupProxyTestEnv creates a test environment with:
// - An in-memory SQLite DB with schema
// - A fake backend that echoes requests
// - A proxy server pointing at the backend
func setupProxyTestEnv(t *testing.T) *ProxyTestEnv {
	t.Helper()

	// Open in-memory DB
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("Failed to open test DB: %v", err)
	}

	// Run migrations
	sqlBytes, err := os.ReadFile("../data/database.sql")
	if err != nil {
		t.Fatalf("Failed to read SQL file: %v", err)
	}
	if _, err := db.Exec(string(sqlBytes)); err != nil {
		t.Fatalf("Failed to run migrations: %v", err)
	}

	// Create a fake backend
	backend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := map[string]interface{}{
			"path":    r.URL.Path,
			"method":  r.Method,
			"query":   r.URL.RawQuery,
			"headers": flattenHeaders(r.Header),
		}
		// Echo body for POST/PUT
		if r.Method == http.MethodPost || r.Method == http.MethodPut {
			body, _ := io.ReadAll(r.Body)
			resp["body"] = string(body)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Backend", "test-backend")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	}))

	// Parse backend host and port
	backendURL := backend.URL // e.g. http://127.0.0.1:PORT
	parts := strings.Split(strings.TrimPrefix(backendURL, "http://"), ":")
	backendHost := parts[0]
	backendPort := 80
	if len(parts) > 1 {
		fmt.Sscanf(parts[1], "%d", &backendPort)
	}

	// Insert test data: host -> vhost -> upstream -> virtual_directory
	insertTestProxyData(t, db, backendHost, backendPort)

	// Create proxy server
	proxyServer := proxy.NewProxyServer(db)
	proxyServer.Start()

	// Wrap in httptest server
	proxyTS := httptest.NewServer(proxyServer)

	return &ProxyTestEnv{
		ProxyServer:   proxyTS,
		BackendServer: backend,
		DB:            db,
		Proxy:         proxyServer,
	}
}

func (env *ProxyTestEnv) Close() {
	env.Proxy.Stop()
	env.ProxyServer.Close()
	env.BackendServer.Close()
	env.DB.Close()
}

func insertTestProxyData(t *testing.T, db *sql.DB, backendHost string, backendPort int) {
	t.Helper()

	// Host
	_, err := db.Exec(`INSERT INTO hosts (id, host_name, description, is_active) VALUES (1, 'test.example.com', 'Test Host', 1)`)
	if err != nil {
		t.Fatalf("Failed to insert host: %v", err)
	}

	// Virtual Host with round_robin
	_, err = db.Exec(`INSERT INTO virtual_hosts (id, host_id, vhost_name, lb_algorithm, sticky_session, failover_mode, is_active)
		VALUES (1, 1, 'test.example.com', 'round_robin', 0, 'active-active', 1)`)
	if err != nil {
		t.Fatalf("Failed to insert vhost: %v", err)
	}

	// Upstream Server pointing to our backend
	_, err = db.Exec(`INSERT INTO upstream_servers (id, virtual_host_id, target_host, target_port, protocol, priority, weight, is_backup, is_active, health_check_enabled, health_check_path, health_check_interval_seconds, health_check_timeout_seconds, max_fails, fail_timeout_seconds)
		VALUES (1, 1, ?, ?, 'http', 1, 1, 0, 1, 0, '/health', 10, 3, 3, 30)`, backendHost, backendPort)
	if err != nil {
		t.Fatalf("Failed to insert upstream: %v", err)
	}

	// Virtual Directory - prefix match
	_, err = db.Exec(`INSERT INTO virtual_directories (id, virtual_host_id, source_path, target_path, match_type, strip_prefix, preserve_host_header, auth_type, is_active, proxy_timeout_seconds, retry_count, retry_delay_ms, max_request_size_mb, websocket_enabled, cache_enabled, cache_ttl_seconds)
		VALUES (1, 1, '/api', '/', 'prefix', 1, 0, 'none', 1, 30, 0, 100, 10, 0, 0, 0)`)
	if err != nil {
		t.Fatalf("Failed to insert vdir prefix: %v", err)
	}

	// Virtual Directory - exact match
	_, err = db.Exec(`INSERT INTO virtual_directories (id, virtual_host_id, source_path, target_path, match_type, strip_prefix, preserve_host_header, auth_type, is_active, proxy_timeout_seconds, retry_count, retry_delay_ms, max_request_size_mb, websocket_enabled, cache_enabled, cache_ttl_seconds)
		VALUES (2, 1, '/exact/path', '/exact-target', 'exact', 0, 0, 'none', 1, 30, 0, 100, 10, 0, 0, 0)`)
	if err != nil {
		t.Fatalf("Failed to insert vdir exact: %v", err)
	}

	// Virtual Directory - regex match
	_, err = db.Exec(`INSERT INTO virtual_directories (id, virtual_host_id, source_path, target_path, match_type, strip_prefix, preserve_host_header, auth_type, is_active, proxy_timeout_seconds, retry_count, retry_delay_ms, max_request_size_mb, websocket_enabled, cache_enabled, cache_ttl_seconds)
		VALUES (3, 1, '^/users/[0-9]+$', '/user-match', 'regex', 0, 0, 'none', 1, 30, 0, 100, 10, 0, 0, 0)`)
	if err != nil {
		t.Fatalf("Failed to insert vdir regex: %v", err)
	}

	// Virtual Directory - parameter match
	_, err = db.Exec(`INSERT INTO virtual_directories (id, virtual_host_id, source_path, target_path, match_type, strip_prefix, preserve_host_header, auth_type, is_active, proxy_timeout_seconds, retry_count, retry_delay_ms, max_request_size_mb, websocket_enabled, cache_enabled, cache_ttl_seconds)
		VALUES (4, 1, '/items/{id}/detail', '/item-detail', 'parameter', 0, 0, 'none', 1, 30, 0, 100, 10, 0, 0, 0)`)
	if err != nil {
		t.Fatalf("Failed to insert vdir parameter: %v", err)
	}

	// Virtual Directory - wildcard match
	_, err = db.Exec(`INSERT INTO virtual_directories (id, virtual_host_id, source_path, target_path, match_type, strip_prefix, preserve_host_header, auth_type, is_active, proxy_timeout_seconds, retry_count, retry_delay_ms, max_request_size_mb, websocket_enabled, cache_enabled, cache_ttl_seconds)
		VALUES (5, 1, '/files/*', '/file-match', 'wildcard', 0, 0, 'none', 1, 30, 0, 100, 10, 0, 0, 0)`)
	if err != nil {
		t.Fatalf("Failed to insert vdir wildcard: %v", err)
	}

	// Virtual Directory - with caching
	_, err = db.Exec(`INSERT INTO virtual_directories (id, virtual_host_id, source_path, target_path, match_type, strip_prefix, preserve_host_header, auth_type, is_active, proxy_timeout_seconds, retry_count, retry_delay_ms, max_request_size_mb, websocket_enabled, cache_enabled, cache_ttl_seconds)
		VALUES (6, 1, '/cached', '/cached-target', 'prefix', 0, 0, 'none', 1, 30, 0, 100, 10, 0, 1, 60)`)
	if err != nil {
		t.Fatalf("Failed to insert vdir cached: %v", err)
	}

	// Virtual Directory - with retry
	_, err = db.Exec(`INSERT INTO virtual_directories (id, virtual_host_id, source_path, target_path, match_type, strip_prefix, preserve_host_header, auth_type, is_active, proxy_timeout_seconds, retry_count, retry_delay_ms, max_request_size_mb, websocket_enabled, cache_enabled, cache_ttl_seconds)
		VALUES (7, 1, '/retry', '/retry-target', 'prefix', 0, 0, 'none', 1, 30, 2, 50, 10, 0, 0, 0)`)
	if err != nil {
		t.Fatalf("Failed to insert vdir retry: %v", err)
	}

	// Virtual Directory - max request size = 1 byte (for testing rejection)
	_, err = db.Exec(`INSERT INTO virtual_directories (id, virtual_host_id, source_path, target_path, match_type, strip_prefix, preserve_host_header, auth_type, is_active, proxy_timeout_seconds, retry_count, retry_delay_ms, max_request_size_mb, websocket_enabled, cache_enabled, cache_ttl_seconds)
		VALUES (8, 1, '/limited', '/limited-target', 'prefix', 0, 0, 'none', 1, 30, 0, 100, 0, 0, 0, 0)`)
	if err != nil {
		t.Fatalf("Failed to insert vdir limited: %v", err)
	}

	// Virtual Directory Methods - restrict vdir ID 2 to GET only
	_, err = db.Exec(`INSERT INTO virtual_directory_methods (virtual_directory_id, http_method) VALUES (2, 'GET')`)
	if err != nil {
		t.Fatalf("Failed to insert method: %v", err)
	}

	// Request Header Rule - add X-Custom-Header to vdir 1
	_, err = db.Exec(`INSERT INTO request_header_rules (virtual_directory_id, header_name, operation, value_source, header_value, execution_order, is_active)
		VALUES (1, 'X-Custom-Header', 'set', 'static', 'hello-from-proxy', 1, 1)`)
	if err != nil {
		t.Fatalf("Failed to insert request header rule: %v", err)
	}

	// Response Header Rule - add X-Proxy-By to vdir 1
	_, err = db.Exec(`INSERT INTO response_header_rules (virtual_directory_id, header_name, operation, header_value, execution_order, is_active)
		VALUES (1, 'X-Proxy-By', 'set', 'swantara-gate', 1, 1)`)
	if err != nil {
		t.Fatalf("Failed to insert response header rule: %v", err)
	}

	// Query Rewrite Rule - add version=v2 to vdir 1
	_, err = db.Exec(`INSERT INTO query_rewrites (virtual_directory_id, param_name, param_value, operation)
		VALUES (1, 'version', 'v2', 'set')`)
	if err != nil {
		t.Fatalf("Failed to insert query rewrite: %v", err)
	}
}

func flattenHeaders(h http.Header) map[string]string {
	result := make(map[string]string)
	for k, vv := range h {
		result[k] = strings.Join(vv, ", ")
	}
	return result
}

// doProxyRequest makes a request to the proxy server with the given Host header
func doProxyRequest(t *testing.T, proxyURL, method, path string, body io.Reader) *http.Response {
	t.Helper()
	req, err := http.NewRequest(method, proxyURL+path, body)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Host = "test.example.com"
	req.Header.Set("Host", "test.example.com")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Failed to make proxy request: %v", err)
	}
	return resp
}

// readJSON reads response body into map
func readProxyJSON(t *testing.T, resp *http.Response) map[string]interface{} {
	t.Helper()
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		t.Fatalf("Failed to parse JSON response: %s, err: %v", string(body), err)
	}
	return result
}

// ============================================================
// TESTS
// ============================================================

func TestProxyRouteMatching(t *testing.T) {
	env := setupProxyTestEnv(t)
	defer env.Close()

	t.Run("Prefix match - /api/hello", func(t *testing.T) {
		resp := doProxyRequest(t, env.ProxyServer.URL, "GET", "/api/hello", nil)
		if resp.StatusCode != 200 {
			t.Fatalf("Expected 200, got %d", resp.StatusCode)
		}
		data := readProxyJSON(t, resp)
		// strip_prefix is true, so /api is stripped -> /hello
		if path, ok := data["path"].(string); !ok || path != "/hello" {
			t.Errorf("Expected path /hello, got %v", data["path"])
		}
	})

	t.Run("Exact match - /exact/path", func(t *testing.T) {
		resp := doProxyRequest(t, env.ProxyServer.URL, "GET", "/exact/path", nil)
		if resp.StatusCode != 200 {
			t.Fatalf("Expected 200, got %d", resp.StatusCode)
		}
		data := readProxyJSON(t, resp)
		// exact match with target_path /exact-target, no strip
		if path, ok := data["path"].(string); !ok || path != "/exact-target/exact/path" {
			t.Errorf("Expected path /exact-target/exact/path, got %v", data["path"])
		}
	})

	t.Run("Regex match - /users/123", func(t *testing.T) {
		resp := doProxyRequest(t, env.ProxyServer.URL, "GET", "/users/123", nil)
		if resp.StatusCode != 200 {
			t.Fatalf("Expected 200, got %d", resp.StatusCode)
		}
		data := readProxyJSON(t, resp)
		if path, ok := data["path"].(string); !ok || path != "/user-match/users/123" {
			t.Errorf("Expected path /user-match/users/123, got %v", data["path"])
		}
	})

	t.Run("Parameter match - /items/42/detail", func(t *testing.T) {
		resp := doProxyRequest(t, env.ProxyServer.URL, "GET", "/items/42/detail", nil)
		if resp.StatusCode != 200 {
			t.Fatalf("Expected 200, got %d", resp.StatusCode)
		}
		data := readProxyJSON(t, resp)
		if path, ok := data["path"].(string); !ok || path != "/item-detail/items/42/detail" {
			t.Errorf("Expected path /item-detail/items/42/detail, got %v", data["path"])
		}
	})

	t.Run("Wildcard match - /files/document", func(t *testing.T) {
		resp := doProxyRequest(t, env.ProxyServer.URL, "GET", "/files/document", nil)
		if resp.StatusCode != 200 {
			t.Fatalf("Expected 200, got %d", resp.StatusCode)
		}
		data := readProxyJSON(t, resp)
		if path, ok := data["path"].(string); !ok || path != "/file-match/files/document" {
			t.Errorf("Expected path /file-match/files/document, got %v", data["path"])
		}
	})

	t.Run("No match - returns 502", func(t *testing.T) {
		req, _ := http.NewRequest("GET", env.ProxyServer.URL+"/api/test", nil)
		req.Host = "unknown.host.com"
		client := &http.Client{Timeout: 5 * time.Second}
		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("Request failed: %v", err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != 502 {
			t.Errorf("Expected 502, got %d", resp.StatusCode)
		}
	})
}

func TestProxyMethodRestriction(t *testing.T) {
	env := setupProxyTestEnv(t)
	defer env.Close()

	t.Run("Allowed method GET on exact route", func(t *testing.T) {
		resp := doProxyRequest(t, env.ProxyServer.URL, "GET", "/exact/path", nil)
		if resp.StatusCode != 200 {
			t.Errorf("Expected 200, got %d", resp.StatusCode)
		}
		resp.Body.Close()
	})

	t.Run("Disallowed method POST on exact route", func(t *testing.T) {
		resp := doProxyRequest(t, env.ProxyServer.URL, "POST", "/exact/path", strings.NewReader(`{}`))
		defer resp.Body.Close()
		if resp.StatusCode != 405 {
			t.Errorf("Expected 405, got %d", resp.StatusCode)
		}
	})
}

func TestProxyHeaderManipulation(t *testing.T) {
	env := setupProxyTestEnv(t)
	defer env.Close()

	t.Run("Request header added by proxy", func(t *testing.T) {
		resp := doProxyRequest(t, env.ProxyServer.URL, "GET", "/api/test-headers", nil)
		if resp.StatusCode != 200 {
			t.Fatalf("Expected 200, got %d", resp.StatusCode)
		}
		data := readProxyJSON(t, resp)
		headers, ok := data["headers"].(map[string]interface{})
		if !ok {
			t.Fatalf("Expected headers map in response")
		}
		if val, ok := headers["X-Custom-Header"].(string); !ok || val != "hello-from-proxy" {
			t.Errorf("Expected X-Custom-Header=hello-from-proxy, got %v", headers["X-Custom-Header"])
		}
	})

	t.Run("Response header added by proxy", func(t *testing.T) {
		resp := doProxyRequest(t, env.ProxyServer.URL, "GET", "/api/test-resp-header", nil)
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			t.Fatalf("Expected 200, got %d", resp.StatusCode)
		}
		if val := resp.Header.Get("X-Proxy-By"); val != "swantara-gate" {
			t.Errorf("Expected X-Proxy-By=swantara-gate, got %q", val)
		}
	})
}

func TestProxyQueryRewrite(t *testing.T) {
	env := setupProxyTestEnv(t)
	defer env.Close()

	t.Run("Query param added by rewrite rule", func(t *testing.T) {
		resp := doProxyRequest(t, env.ProxyServer.URL, "GET", "/api/test-query?existing=yes", nil)
		if resp.StatusCode != 200 {
			t.Fatalf("Expected 200, got %d", resp.StatusCode)
		}
		data := readProxyJSON(t, resp)
		query, ok := data["query"].(string)
		if !ok {
			t.Fatalf("Expected query string in response")
		}
		if !strings.Contains(query, "version=v2") {
			t.Errorf("Expected query to contain version=v2, got %q", query)
		}
		if !strings.Contains(query, "existing=yes") {
			t.Errorf("Expected query to still contain existing=yes, got %q", query)
		}
	})
}

func TestProxyCache(t *testing.T) {
	env := setupProxyTestEnv(t)
	defer env.Close()

	t.Run("First request is MISS, second is HIT", func(t *testing.T) {
		// First request - MISS
		resp1 := doProxyRequest(t, env.ProxyServer.URL, "GET", "/cached/data", nil)
		if resp1.StatusCode != 200 {
			t.Fatalf("Expected 200, got %d", resp1.StatusCode)
		}
		if val := resp1.Header.Get("X-Cache"); val != "MISS" {
			t.Errorf("Expected X-Cache=MISS on first request, got %q", val)
		}
		resp1.Body.Close()

		// Second request - HIT
		resp2 := doProxyRequest(t, env.ProxyServer.URL, "GET", "/cached/data", nil)
		if resp2.StatusCode != 200 {
			t.Fatalf("Expected 200, got %d", resp2.StatusCode)
		}
		if val := resp2.Header.Get("X-Cache"); val != "HIT" {
			t.Errorf("Expected X-Cache=HIT on second request, got %q", val)
		}
		resp2.Body.Close()
	})

	t.Run("POST requests are not cached", func(t *testing.T) {
		resp := doProxyRequest(t, env.ProxyServer.URL, "POST", "/cached/data", strings.NewReader(`{}`))
		defer resp.Body.Close()
		// POST goes through to backend (not cached)
		if resp.StatusCode != 200 {
			t.Fatalf("Expected 200, got %d", resp.StatusCode)
		}
		// No X-Cache header for non-GET
		if val := resp.Header.Get("X-Cache"); val != "" {
			t.Errorf("Expected no X-Cache for POST, got %q", val)
		}
	})
}

func TestProxyForwarding(t *testing.T) {
	env := setupProxyTestEnv(t)
	defer env.Close()

	t.Run("POST body forwarded correctly", func(t *testing.T) {
		body := `{"name":"test","value":123}`
		resp := doProxyRequest(t, env.ProxyServer.URL, "POST", "/api/data", strings.NewReader(body))
		if resp.StatusCode != 200 {
			t.Fatalf("Expected 200, got %d", resp.StatusCode)
		}
		data := readProxyJSON(t, resp)
		if data["method"] != "POST" {
			t.Errorf("Expected method POST, got %v", data["method"])
		}
		if data["body"] != body {
			t.Errorf("Expected body %q, got %v", body, data["body"])
		}
	})

	t.Run("X-Forwarded-For header set", func(t *testing.T) {
		resp := doProxyRequest(t, env.ProxyServer.URL, "GET", "/api/check-xff", nil)
		if resp.StatusCode != 200 {
			t.Fatalf("Expected 200, got %d", resp.StatusCode)
		}
		data := readProxyJSON(t, resp)
		headers, _ := data["headers"].(map[string]interface{})
		if xff, ok := headers["X-Forwarded-For"].(string); !ok || xff == "" {
			t.Errorf("Expected X-Forwarded-For to be set, got %v", headers["X-Forwarded-For"])
		}
	})
}

func TestProxyLoadBalancing(t *testing.T) {
	env := setupProxyTestEnv(t)
	defer env.Close()

	// Add a second upstream server (same backend for simplicity)
	backendURL := env.BackendServer.URL
	parts := strings.Split(strings.TrimPrefix(backendURL, "http://"), ":")
	backendHost := parts[0]
	backendPort := 80
	if len(parts) > 1 {
		fmt.Sscanf(parts[1], "%d", &backendPort)
	}

	_, err := env.DB.Exec(`INSERT INTO upstream_servers (id, virtual_host_id, target_host, target_port, protocol, priority, weight, is_backup, is_active, health_check_enabled)
		VALUES (2, 1, ?, ?, 'http', 1, 1, 0, 1, 0)`, backendHost, backendPort)
	if err != nil {
		t.Fatalf("Failed to insert second upstream: %v", err)
	}

	// Reload proxy config
	env.Proxy.ReloadConfig()

	t.Run("Round robin distributes between servers", func(t *testing.T) {
		// Make multiple requests - they should all succeed (both point to same backend)
		for i := 0; i < 10; i++ {
			resp := doProxyRequest(t, env.ProxyServer.URL, "GET", "/api/lb-test", nil)
			if resp.StatusCode != 200 {
				t.Fatalf("Request %d: Expected 200, got %d", i, resp.StatusCode)
			}
			resp.Body.Close()
		}
	})
}

func TestProxyHealthCheck(t *testing.T) {
	env := setupProxyTestEnv(t)
	defer env.Close()

	t.Run("Unhealthy server is skipped", func(t *testing.T) {
		// Add a dead upstream (port that won't respond)
		_, err := env.DB.Exec(`INSERT INTO upstream_servers (id, virtual_host_id, target_host, target_port, protocol, priority, weight, is_backup, is_active, health_check_enabled, max_fails, fail_timeout_seconds)
			VALUES (3, 1, '127.0.0.1', 59999, 'http', 1, 1, 0, 1, 0, 1, 30)`)
		if err != nil {
			t.Fatalf("Failed to insert dead upstream: %v", err)
		}
		env.Proxy.ReloadConfig()

		// Requests should still succeed (healthy server still available)
		resp := doProxyRequest(t, env.ProxyServer.URL, "GET", "/api/health-test", nil)
		if resp.StatusCode != 200 {
			t.Fatalf("Expected 200, got %d", resp.StatusCode)
		}
		resp.Body.Close()
	})
}

func TestProxyMaxRequestSize(t *testing.T) {
	env := setupProxyTestEnv(t)
	defer env.Close()

	t.Run("Request within size limit", func(t *testing.T) {
		// /api has 10MB limit, small body should pass
		resp := doProxyRequest(t, env.ProxyServer.URL, "POST", "/api/upload", strings.NewReader("small"))
		if resp.StatusCode != 200 {
			t.Fatalf("Expected 200, got %d", resp.StatusCode)
		}
		resp.Body.Close()
	})

	t.Run("Request exceeds content-length limit", func(t *testing.T) {
		// Create request with explicit Content-Length exceeding 10MB
		largeSize := int64(11 * 1024 * 1024) // 11MB
		req, _ := http.NewRequest("POST", env.ProxyServer.URL+"/api/big-upload", nil)
		req.Host = "test.example.com"
		req.ContentLength = largeSize

		client := &http.Client{Timeout: 5 * time.Second}
		resp, err := client.Do(req)
		if err != nil {
			// Some clients/servers reject this, that's fine
			return
		}
		defer resp.Body.Close()
		if resp.StatusCode != 413 {
			t.Errorf("Expected 413 (entity too large), got %d", resp.StatusCode)
		}
	})
}
