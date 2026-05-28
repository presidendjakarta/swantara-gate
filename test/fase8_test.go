package test

import (
	"crypto/hmac"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
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
	"golang.org/x/crypto/bcrypt"
)

// SecurityTestEnv holds the proxy + backend for Phase 8 security tests
type SecurityTestEnv struct {
	ProxyServer   *httptest.Server
	BackendServer *httptest.Server
	DB            *sql.DB
	Proxy         *proxy.ProxyServer
}

func setupSecurityTestEnv(t *testing.T) *SecurityTestEnv {
	t.Helper()

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
			"path":   r.URL.Path,
			"method": r.Method,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	}))

	// Parse backend host and port
	parts := strings.Split(strings.TrimPrefix(backend.URL, "http://"), ":")
	backendHost := parts[0]
	backendPort := 80
	if len(parts) > 1 {
		fmt.Sscanf(parts[1], "%d", &backendPort)
	}

	// Insert base test data
	insertSecurityTestData(t, db, backendHost, backendPort)

	// Create proxy
	proxyServer := proxy.NewProxyServer(db)
	proxyServer.Start()

	proxyTS := httptest.NewServer(proxyServer)

	return &SecurityTestEnv{
		ProxyServer:   proxyTS,
		BackendServer: backend,
		DB:            db,
		Proxy:         proxyServer,
	}
}

func (env *SecurityTestEnv) Close() {
	env.Proxy.Stop()
	env.ProxyServer.Close()
	env.BackendServer.Close()
	env.DB.Close()
}

func insertSecurityTestData(t *testing.T, db *sql.DB, backendHost string, backendPort int) {
	t.Helper()

	// Host
	db.Exec(`INSERT INTO hosts (id, host_name, is_active) VALUES (1, 'secure.example.com', 1)`)

	// Virtual Host
	db.Exec(`INSERT INTO virtual_hosts (id, host_id, vhost_name, lb_algorithm, sticky_session, failover_mode, is_active)
		VALUES (1, 1, 'secure.example.com', 'round_robin', 0, 'active-active', 1)`)

	// Upstream
	db.Exec(`INSERT INTO upstream_servers (id, virtual_host_id, target_host, target_port, protocol, priority, weight, is_backup, is_active, health_check_enabled, health_check_path, health_check_interval_seconds, health_check_timeout_seconds, max_fails, fail_timeout_seconds)
		VALUES (1, 1, ?, ?, 'http', 1, 1, 0, 1, 0, '/health', 10, 3, 3, 30)`, backendHost, backendPort)

	// === Routes ===

	// Route 1: Open route (no auth, no limits)
	db.Exec(`INSERT INTO virtual_directories (id, virtual_host_id, source_path, target_path, match_type, strip_prefix, preserve_host_header, auth_type, is_active, proxy_timeout_seconds, retry_count, retry_delay_ms, max_request_size_mb, websocket_enabled, cache_enabled, cache_ttl_seconds)
		VALUES (1, 1, '/open', '/', 'prefix', 1, 0, 'none', 1, 30, 0, 100, 10, 0, 0, 0)`)

	// Route 2: API Key protected route
	db.Exec(`INSERT INTO virtual_directories (id, virtual_host_id, source_path, target_path, match_type, strip_prefix, preserve_host_header, auth_type, is_active, proxy_timeout_seconds, retry_count, retry_delay_ms, max_request_size_mb, websocket_enabled, cache_enabled, cache_ttl_seconds)
		VALUES (2, 1, '/api-key-protected', '/', 'prefix', 1, 0, 'api_key', 1, 30, 0, 100, 10, 0, 0, 0)`)

	// Route 3: JWT protected route
	db.Exec(`INSERT INTO virtual_directories (id, virtual_host_id, source_path, target_path, match_type, strip_prefix, preserve_host_header, auth_type, is_active, proxy_timeout_seconds, retry_count, retry_delay_ms, max_request_size_mb, websocket_enabled, cache_enabled, cache_ttl_seconds)
		VALUES (3, 1, '/jwt-protected', '/', 'prefix', 1, 0, 'jwt', 1, 30, 0, 100, 10, 0, 0, 0)`)

	// Route 4: Basic auth protected route
	db.Exec(`INSERT INTO virtual_directories (id, virtual_host_id, source_path, target_path, match_type, strip_prefix, preserve_host_header, auth_type, is_active, proxy_timeout_seconds, retry_count, retry_delay_ms, max_request_size_mb, websocket_enabled, cache_enabled, cache_ttl_seconds)
		VALUES (4, 1, '/basic-protected', '/', 'prefix', 1, 0, 'basic', 1, 30, 0, 100, 10, 0, 0, 0)`)

	// Route 5: Rate-limited route
	db.Exec(`INSERT INTO virtual_directories (id, virtual_host_id, source_path, target_path, match_type, strip_prefix, preserve_host_header, auth_type, is_active, proxy_timeout_seconds, retry_count, retry_delay_ms, max_request_size_mb, websocket_enabled, cache_enabled, cache_ttl_seconds)
		VALUES (5, 1, '/rate-limited', '/', 'prefix', 1, 0, 'none', 1, 30, 0, 100, 10, 0, 0, 0)`)

	// Route 6: IP-whitelisted route
	db.Exec(`INSERT INTO virtual_directories (id, virtual_host_id, source_path, target_path, match_type, strip_prefix, preserve_host_header, auth_type, is_active, proxy_timeout_seconds, retry_count, retry_delay_ms, max_request_size_mb, websocket_enabled, cache_enabled, cache_ttl_seconds)
		VALUES (6, 1, '/ip-whitelist', '/', 'prefix', 1, 0, 'none', 1, 30, 0, 100, 10, 0, 0, 0)`)

	// Route 7: IP-blacklisted route
	db.Exec(`INSERT INTO virtual_directories (id, virtual_host_id, source_path, target_path, match_type, strip_prefix, preserve_host_header, auth_type, is_active, proxy_timeout_seconds, retry_count, retry_delay_ms, max_request_size_mb, websocket_enabled, cache_enabled, cache_ttl_seconds)
		VALUES (7, 1, '/ip-blacklist', '/', 'prefix', 1, 0, 'none', 1, 30, 0, 100, 10, 0, 0, 0)`)

	// Route 8: CORS-enabled route
	db.Exec(`INSERT INTO virtual_directories (id, virtual_host_id, source_path, target_path, match_type, strip_prefix, preserve_host_header, auth_type, is_active, proxy_timeout_seconds, retry_count, retry_delay_ms, max_request_size_mb, websocket_enabled, cache_enabled, cache_ttl_seconds)
		VALUES (8, 1, '/cors', '/', 'prefix', 1, 0, 'none', 1, 30, 0, 100, 10, 0, 0, 0)`)

	// Route 9: Circuit-breaker-protected route
	db.Exec(`INSERT INTO virtual_directories (id, virtual_host_id, source_path, target_path, match_type, strip_prefix, preserve_host_header, auth_type, is_active, proxy_timeout_seconds, retry_count, retry_delay_ms, max_request_size_mb, websocket_enabled, cache_enabled, cache_ttl_seconds)
		VALUES (9, 1, '/circuit', '/', 'prefix', 1, 0, 'none', 1, 30, 0, 100, 10, 0, 0, 0)`)

	// === Security Configs ===

	// API Key setup: consumer + api_key
	db.Exec(`INSERT INTO api_consumers (id, consumer_name, is_active) VALUES (1, 'test-consumer', 1)`)
	db.Exec(`INSERT INTO api_keys (id, consumer_id, api_key, is_active) VALUES (1, 1, 'test-api-key-123', 1)`)

	// JWT config for route 3
	db.Exec(`INSERT INTO jwt_configs (virtual_directory_id, algorithm, jwt_secret, issuer, audience, clock_skew_seconds, require_exp, require_iat, is_active)
		VALUES (3, 'HS256', 'test-secret-key', 'test-issuer', 'test-audience', 30, 1, 0, 1)`)

	// Basic auth: consumer credential
	hash, _ := bcrypt.GenerateFromPassword([]byte("test-password"), bcrypt.DefaultCost)
	db.Exec(`INSERT INTO consumer_credentials (id, consumer_id, auth_type, username, password_hash, is_active)
		VALUES (1, 1, 'basic', 'testuser', ?, 1)`, string(hash))

	// Rate limit for route 5: 3 requests/minute, burst 1
	db.Exec(`INSERT INTO rate_limits (virtual_directory_id, limit_by, requests_per_minute, burst, block_duration_seconds, is_active)
		VALUES (5, 'ip', 3, 1, 60, 1)`)

	// IP whitelist for route 6: only allow 10.0.0.1
	db.Exec(`INSERT INTO ip_whitelists (virtual_directory_id, ip_address, is_active) VALUES (6, '10.0.0.1', 1)`)

	// IP blacklist for route 7: block 192.168.1.100
	db.Exec(`INSERT INTO ip_blacklists (virtual_directory_id, ip_address, is_active) VALUES (7, '127.0.0.1', 1)`)

	// CORS config for route 8
	db.Exec(`INSERT INTO cors_configs (virtual_directory_id, allowed_origins, allowed_methods, allowed_headers, exposed_headers, allow_credentials, max_age_seconds, is_active)
		VALUES (8, 'https://allowed.example.com', 'GET,POST', 'Content-Type,Authorization', 'X-Custom', 1, 3600, 1)`)

	// Circuit breaker for route 9: threshold 2, recovery 1s
	db.Exec(`INSERT INTO circuit_breakers (virtual_directory_id, enabled, failure_threshold, recovery_timeout_seconds, half_open_max_requests)
		VALUES (9, 1, 2, 1, 1)`)
}

func doSecurityRequest(t *testing.T, proxyURL, path, method string, headers map[string]string) *http.Response {
	t.Helper()
	req, err := http.NewRequest(method, proxyURL+path, nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Host = "secure.example.com"
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Request failed: %v", err)
	}
	return resp
}

// === Tests ===

func TestSecurityAPIKeyAuth(t *testing.T) {
	env := setupSecurityTestEnv(t)
	defer env.Close()

	t.Run("Valid API key in header", func(t *testing.T) {
		resp := doSecurityRequest(t, env.ProxyServer.URL, "/api-key-protected/test", "GET",
			map[string]string{"X-API-Key": "test-api-key-123"})
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			body, _ := io.ReadAll(resp.Body)
			t.Errorf("Expected 200, got %d: %s", resp.StatusCode, string(body))
		}
	})

	t.Run("Valid API key in query param", func(t *testing.T) {
		resp := doSecurityRequest(t, env.ProxyServer.URL, "/api-key-protected/test?api_key=test-api-key-123", "GET", nil)
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			body, _ := io.ReadAll(resp.Body)
			t.Errorf("Expected 200, got %d: %s", resp.StatusCode, string(body))
		}
	})

	t.Run("Missing API key", func(t *testing.T) {
		resp := doSecurityRequest(t, env.ProxyServer.URL, "/api-key-protected/test", "GET", nil)
		defer resp.Body.Close()
		if resp.StatusCode != 401 {
			t.Errorf("Expected 401, got %d", resp.StatusCode)
		}
	})

	t.Run("Invalid API key", func(t *testing.T) {
		resp := doSecurityRequest(t, env.ProxyServer.URL, "/api-key-protected/test", "GET",
			map[string]string{"X-API-Key": "wrong-key"})
		defer resp.Body.Close()
		if resp.StatusCode != 401 {
			t.Errorf("Expected 401, got %d", resp.StatusCode)
		}
	})
}

func TestSecurityJWTAuth(t *testing.T) {
	env := setupSecurityTestEnv(t)
	defer env.Close()

	t.Run("Valid JWT token", func(t *testing.T) {
		token := createTestJWT(t, "test-secret-key", "test-issuer", "test-audience", time.Now().Add(1*time.Hour))
		resp := doSecurityRequest(t, env.ProxyServer.URL, "/jwt-protected/test", "GET",
			map[string]string{"Authorization": "Bearer " + token})
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			body, _ := io.ReadAll(resp.Body)
			t.Errorf("Expected 200, got %d: %s", resp.StatusCode, string(body))
		}
	})

	t.Run("Expired JWT token", func(t *testing.T) {
		token := createTestJWT(t, "test-secret-key", "test-issuer", "test-audience", time.Now().Add(-1*time.Hour))
		resp := doSecurityRequest(t, env.ProxyServer.URL, "/jwt-protected/test", "GET",
			map[string]string{"Authorization": "Bearer " + token})
		defer resp.Body.Close()
		if resp.StatusCode != 401 {
			t.Errorf("Expected 401, got %d", resp.StatusCode)
		}
	})

	t.Run("Invalid signature", func(t *testing.T) {
		token := createTestJWT(t, "wrong-secret", "test-issuer", "test-audience", time.Now().Add(1*time.Hour))
		resp := doSecurityRequest(t, env.ProxyServer.URL, "/jwt-protected/test", "GET",
			map[string]string{"Authorization": "Bearer " + token})
		defer resp.Body.Close()
		if resp.StatusCode != 401 {
			t.Errorf("Expected 401, got %d", resp.StatusCode)
		}
	})

	t.Run("Missing token", func(t *testing.T) {
		resp := doSecurityRequest(t, env.ProxyServer.URL, "/jwt-protected/test", "GET", nil)
		defer resp.Body.Close()
		if resp.StatusCode != 401 {
			t.Errorf("Expected 401, got %d", resp.StatusCode)
		}
	})

	t.Run("Invalid issuer", func(t *testing.T) {
		token := createTestJWT(t, "test-secret-key", "wrong-issuer", "test-audience", time.Now().Add(1*time.Hour))
		resp := doSecurityRequest(t, env.ProxyServer.URL, "/jwt-protected/test", "GET",
			map[string]string{"Authorization": "Bearer " + token})
		defer resp.Body.Close()
		if resp.StatusCode != 401 {
			t.Errorf("Expected 401, got %d", resp.StatusCode)
		}
	})
}

func TestSecurityBasicAuth(t *testing.T) {
	env := setupSecurityTestEnv(t)
	defer env.Close()

	t.Run("Valid credentials", func(t *testing.T) {
		creds := base64.StdEncoding.EncodeToString([]byte("testuser:test-password"))
		resp := doSecurityRequest(t, env.ProxyServer.URL, "/basic-protected/test", "GET",
			map[string]string{"Authorization": "Basic " + creds})
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			body, _ := io.ReadAll(resp.Body)
			t.Errorf("Expected 200, got %d: %s", resp.StatusCode, string(body))
		}
	})

	t.Run("Wrong password", func(t *testing.T) {
		creds := base64.StdEncoding.EncodeToString([]byte("testuser:wrong-password"))
		resp := doSecurityRequest(t, env.ProxyServer.URL, "/basic-protected/test", "GET",
			map[string]string{"Authorization": "Basic " + creds})
		defer resp.Body.Close()
		if resp.StatusCode != 401 {
			t.Errorf("Expected 401, got %d", resp.StatusCode)
		}
	})

	t.Run("Missing auth header", func(t *testing.T) {
		resp := doSecurityRequest(t, env.ProxyServer.URL, "/basic-protected/test", "GET", nil)
		defer resp.Body.Close()
		if resp.StatusCode != 401 {
			t.Errorf("Expected 401, got %d", resp.StatusCode)
		}
	})

	t.Run("Invalid username", func(t *testing.T) {
		creds := base64.StdEncoding.EncodeToString([]byte("wronguser:test-password"))
		resp := doSecurityRequest(t, env.ProxyServer.URL, "/basic-protected/test", "GET",
			map[string]string{"Authorization": "Basic " + creds})
		defer resp.Body.Close()
		if resp.StatusCode != 401 {
			t.Errorf("Expected 401, got %d", resp.StatusCode)
		}
	})
}

func TestSecurityRateLimit(t *testing.T) {
	env := setupSecurityTestEnv(t)
	defer env.Close()

	t.Run("Rate limit enforcement", func(t *testing.T) {
		// Config: 3 requests/min + 1 burst = 4 max before block
		for i := 0; i < 4; i++ {
			resp := doSecurityRequest(t, env.ProxyServer.URL, "/rate-limited/test", "GET", nil)
			resp.Body.Close()
			if resp.StatusCode != 200 {
				t.Errorf("Request %d: expected 200, got %d", i+1, resp.StatusCode)
			}
		}

		// 5th request should be rate limited
		resp := doSecurityRequest(t, env.ProxyServer.URL, "/rate-limited/test", "GET", nil)
		defer resp.Body.Close()
		if resp.StatusCode != 429 {
			t.Errorf("Expected 429 (rate limited), got %d", resp.StatusCode)
		}
	})
}

func TestSecurityIPWhitelist(t *testing.T) {
	env := setupSecurityTestEnv(t)
	defer env.Close()

	t.Run("Non-whitelisted IP blocked", func(t *testing.T) {
		// Test client IP is 127.0.0.1, whitelist only allows 10.0.0.1
		resp := doSecurityRequest(t, env.ProxyServer.URL, "/ip-whitelist/test", "GET", nil)
		defer resp.Body.Close()
		if resp.StatusCode != 403 {
			t.Errorf("Expected 403, got %d", resp.StatusCode)
		}
	})
}

func TestSecurityIPBlacklist(t *testing.T) {
	env := setupSecurityTestEnv(t)
	defer env.Close()

	t.Run("Blacklisted IP blocked", func(t *testing.T) {
		// Test client IP is 127.0.0.1, which is blacklisted for route 7
		resp := doSecurityRequest(t, env.ProxyServer.URL, "/ip-blacklist/test", "GET", nil)
		defer resp.Body.Close()
		if resp.StatusCode != 403 {
			t.Errorf("Expected 403, got %d", resp.StatusCode)
		}
	})
}

func TestSecurityCORS(t *testing.T) {
	env := setupSecurityTestEnv(t)
	defer env.Close()

	t.Run("Preflight request from allowed origin", func(t *testing.T) {
		resp := doSecurityRequest(t, env.ProxyServer.URL, "/cors/test", "OPTIONS",
			map[string]string{
				"Origin":                        "https://allowed.example.com",
				"Access-Control-Request-Method":  "POST",
				"Access-Control-Request-Headers": "Content-Type",
			})
		defer resp.Body.Close()
		if resp.StatusCode != 204 {
			t.Errorf("Expected 204, got %d", resp.StatusCode)
		}
		acao := resp.Header.Get("Access-Control-Allow-Origin")
		if acao != "https://allowed.example.com" {
			t.Errorf("Expected Allow-Origin 'https://allowed.example.com', got '%s'", acao)
		}
		acac := resp.Header.Get("Access-Control-Allow-Credentials")
		if acac != "true" {
			t.Errorf("Expected Allow-Credentials 'true', got '%s'", acac)
		}
	})

	t.Run("Preflight from disallowed origin", func(t *testing.T) {
		resp := doSecurityRequest(t, env.ProxyServer.URL, "/cors/test", "OPTIONS",
			map[string]string{
				"Origin":                       "https://evil.example.com",
				"Access-Control-Request-Method": "POST",
			})
		defer resp.Body.Close()
		// Should not get CORS headers (passes through without CORS handling)
		acao := resp.Header.Get("Access-Control-Allow-Origin")
		if acao != "" {
			t.Errorf("Expected no Allow-Origin header, got '%s'", acao)
		}
	})

	t.Run("Normal request with CORS headers", func(t *testing.T) {
		resp := doSecurityRequest(t, env.ProxyServer.URL, "/cors/test", "GET",
			map[string]string{"Origin": "https://allowed.example.com"})
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			t.Errorf("Expected 200, got %d", resp.StatusCode)
		}
		acao := resp.Header.Get("Access-Control-Allow-Origin")
		if acao != "https://allowed.example.com" {
			t.Errorf("Expected Allow-Origin, got '%s'", acao)
		}
	})
}

func TestSecurityCircuitBreaker(t *testing.T) {
	env := setupSecurityTestEnv(t)
	defer env.Close()

	t.Run("Circuit opens after failures", func(t *testing.T) {
		// Stop the backend to cause failures
		env.BackendServer.Close()

		// Circuit breaker config: threshold=2, recovery=1s
		// Send requests that will fail (backend is down)
		for i := 0; i < 3; i++ {
			resp := doSecurityRequest(t, env.ProxyServer.URL, "/circuit/test", "GET", nil)
			resp.Body.Close()
		}

		// After 2+ failures, circuit should be open
		resp := doSecurityRequest(t, env.ProxyServer.URL, "/circuit/test", "GET", nil)
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)

		if resp.StatusCode != 503 {
			t.Errorf("Expected 503 (circuit open), got %d: %s", resp.StatusCode, string(body))
		}
		if !strings.Contains(string(body), "circuit breaker") {
			t.Errorf("Expected circuit breaker message, got: %s", string(body))
		}
	})
}

func TestSecurityMaintenance(t *testing.T) {
	env := setupSecurityTestEnv(t)
	defer env.Close()

	t.Run("Maintenance mode blocks traffic", func(t *testing.T) {
		// Insert an active maintenance window
		start := "2020-01-01 00:00:00"
		end := "2030-12-31 23:59:59"
		_, err := env.DB.Exec(`INSERT INTO maintenance_windows (virtual_host_id, title, start_at, end_at, maintenance_response_code, maintenance_message, is_active)
			VALUES (1, 'Test Maintenance', ?, ?, 503, 'Under maintenance', 1)`, start, end)
		if err != nil {
			t.Fatalf("Failed to insert maintenance window: %v", err)
		}

		// Reload proxy config
		env.Proxy.ReloadConfig()

		resp := doSecurityRequest(t, env.ProxyServer.URL, "/open/test", "GET", nil)
		defer resp.Body.Close()
		if resp.StatusCode != 503 {
			body, _ := io.ReadAll(resp.Body)
			t.Errorf("Expected 503, got %d: %s", resp.StatusCode, string(body))
		}
	})
}

func TestSecurityOpenRoute(t *testing.T) {
	env := setupSecurityTestEnv(t)
	defer env.Close()

	t.Run("Open route passes without auth", func(t *testing.T) {
		resp := doSecurityRequest(t, env.ProxyServer.URL, "/open/test", "GET", nil)
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			body, _ := io.ReadAll(resp.Body)
			t.Errorf("Expected 200, got %d: %s", resp.StatusCode, string(body))
		}
	})
}

// === Helpers ===

func createTestJWT(t *testing.T, secret, issuer, audience string, expiry time.Time) string {
	t.Helper()

	// Header
	header := map[string]string{"alg": "HS256", "typ": "JWT"}
	headerJSON, _ := json.Marshal(header)
	headerB64 := base64URLEncode(headerJSON)

	// Payload
	payload := map[string]interface{}{
		"sub": "test-user",
		"iss": issuer,
		"aud": audience,
		"exp": expiry.Unix(),
		"iat": time.Now().Unix(),
	}
	payloadJSON, _ := json.Marshal(payload)
	payloadB64 := base64URLEncode(payloadJSON)

	// Signature
	signingInput := headerB64 + "." + payloadB64
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(signingInput))
	sig := mac.Sum(nil)
	sigB64 := base64URLEncode(sig)

	return headerB64 + "." + payloadB64 + "." + sigB64
}

func base64URLEncode(data []byte) string {
	return strings.TrimRight(base64.URLEncoding.EncodeToString(data), "=")
}
