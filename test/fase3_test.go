package test

import (
	"fmt"
	"net/http"
	"testing"
)

// TestJWTConfigsCRUD menguji CRUD lengkap untuk JWT Configs
func TestJWTConfigsCRUD(t *testing.T) {
	ts := SetupTestServer(t)
	defer ts.Close()

	// Setup
	hostID := ts.CreateTestHost(t)
	vhostID := ts.CreateTestVirtualHost(t, hostID)
	dirID := ts.CreateTestVirtualDirectory(t, vhostID)

	// === CREATE ===
	t.Run("Create JWT Config", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/jwt-configs", map[string]interface{}{
			"virtual_directory_id": dirID,
			"jwt_secret":           "my-super-secret-key-2024",
			"algorithm":            "HS256",
			"issuer":               "swantara-gate",
			"audience":             "mobile-app",
			"expired_in_seconds":   7200,
			"clock_skew_seconds":   60,
			"is_active":            true,
		})
		AssertStatus(t, resp, http.StatusCreated)
		res := AssertSuccess(t, body)
		if ExtractID(res.Data) == 0 {
			t.Fatal("JWT Config ID should not be 0")
		}
	})

	// === CREATE - Validasi ===
	t.Run("Create JWT Config - Missing Secret", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/jwt-configs", map[string]interface{}{
			"virtual_directory_id": dirID,
			"algorithm":            "HS256",
		})
		AssertStatus(t, resp, http.StatusBadRequest)
		AssertError(t, body)
	})

	// === GET ALL ===
	t.Run("Get All JWT Configs", func(t *testing.T) {
		resp, body := ts.DoRequest("GET", "/api/admin/jwt-configs?page=1&limit=10", nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	// === GET BY ID ===
	t.Run("Get JWT Config By ID", func(t *testing.T) {
		resp, body := ts.DoRequest("GET", "/api/admin/jwt-configs/1", nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	// === UPDATE ===
	t.Run("Update JWT Config", func(t *testing.T) {
		resp, body := ts.DoRequest("PUT", "/api/admin/jwt-configs/1", map[string]interface{}{
			"jwt_secret":         "new-secret-key-2024",
			"expired_in_seconds": 3600,
			"clock_skew_seconds": 30,
			"is_active":          true,
		})
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	// === DELETE ===
	t.Run("Delete JWT Config", func(t *testing.T) {
		resp, body := ts.DoRequest("DELETE", "/api/admin/jwt-configs/1", nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})
}

// TestExternalAuthCRUD menguji CRUD lengkap untuk External Auth
func TestExternalAuthCRUD(t *testing.T) {
	ts := SetupTestServer(t)
	defer ts.Close()

	hostID := ts.CreateTestHost(t)
	vhostID := ts.CreateTestVirtualHost(t, hostID)
	dirID := ts.CreateTestVirtualDirectory(t, vhostID)

	// === CREATE ===
	t.Run("Create External Auth", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/external-auth", map[string]interface{}{
			"virtual_directory_id":  dirID,
			"auth_url":             "https://auth.example.com/validate",
			"method":               "POST",
			"timeout_seconds":      10,
			"success_status_key":   "status",
			"success_status_value": "valid",
			"forward_headers":      "Authorization,X-Request-ID",
			"cache_ttl_seconds":    300,
			"is_active":            true,
		})
		AssertStatus(t, resp, http.StatusCreated)
		res := AssertSuccess(t, body)
		if ExtractID(res.Data) == 0 {
			t.Fatal("External Auth ID should not be 0")
		}
	})

	// === CREATE - Validasi ===
	t.Run("Create External Auth - Missing URL", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/external-auth", map[string]interface{}{
			"virtual_directory_id": dirID,
			"method":              "POST",
		})
		AssertStatus(t, resp, http.StatusBadRequest)
		AssertError(t, body)
	})

	// === GET ALL ===
	t.Run("Get All External Auth", func(t *testing.T) {
		resp, body := ts.DoRequest("GET", "/api/admin/external-auth", nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	// === GET BY ID ===
	t.Run("Get External Auth By ID", func(t *testing.T) {
		resp, body := ts.DoRequest("GET", "/api/admin/external-auth/1", nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	// === UPDATE ===
	t.Run("Update External Auth", func(t *testing.T) {
		resp, body := ts.DoRequest("PUT", "/api/admin/external-auth/1", map[string]interface{}{
			"auth_url":        "https://auth2.example.com/check",
			"timeout_seconds": 15,
			"is_active":       true,
		})
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	// === DELETE ===
	t.Run("Delete External Auth", func(t *testing.T) {
		resp, body := ts.DoRequest("DELETE", "/api/admin/external-auth/1", nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})
}

// TestRateLimitsCRUD menguji CRUD lengkap untuk Rate Limits
func TestRateLimitsCRUD(t *testing.T) {
	ts := SetupTestServer(t)
	defer ts.Close()

	hostID := ts.CreateTestHost(t)
	vhostID := ts.CreateTestVirtualHost(t, hostID)
	dirID := ts.CreateTestVirtualDirectory(t, vhostID)

	// === CREATE ===
	t.Run("Create Rate Limit", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/rate-limits", map[string]interface{}{
			"virtual_directory_id": dirID,
			"limit_by":            "ip",
			"requests_per_period":  100,
			"period_seconds":       60,
			"burst_size":           20,
			"block_duration_seconds": 120,
			"is_active":           true,
		})
		AssertStatus(t, resp, http.StatusCreated)
		res := AssertSuccess(t, body)
		if ExtractID(res.Data) == 0 {
			t.Fatal("Rate Limit ID should not be 0")
		}
	})

	// === CREATE - Validasi ===
	t.Run("Create Rate Limit - Missing VDir", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/rate-limits", map[string]interface{}{
			"limit_by": "ip",
		})
		AssertStatus(t, resp, http.StatusBadRequest)
		AssertError(t, body)
	})

	// === GET ALL ===
	t.Run("Get All Rate Limits", func(t *testing.T) {
		resp, body := ts.DoRequest("GET", "/api/admin/rate-limits", nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	// === GET BY ID ===
	t.Run("Get Rate Limit By ID", func(t *testing.T) {
		resp, body := ts.DoRequest("GET", "/api/admin/rate-limits/1", nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	// === UPDATE ===
	t.Run("Update Rate Limit", func(t *testing.T) {
		resp, body := ts.DoRequest("PUT", "/api/admin/rate-limits/1", map[string]interface{}{
			"requests_per_period": 200,
			"burst_size":          50,
			"is_active":           true,
		})
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	// === DELETE ===
	t.Run("Delete Rate Limit", func(t *testing.T) {
		resp, body := ts.DoRequest("DELETE", "/api/admin/rate-limits/1", nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})
}

// TestCORSConfigsCRUD menguji CRUD lengkap untuk CORS Configs
func TestCORSConfigsCRUD(t *testing.T) {
	ts := SetupTestServer(t)
	defer ts.Close()

	hostID := ts.CreateTestHost(t)
	vhostID := ts.CreateTestVirtualHost(t, hostID)
	dirID := ts.CreateTestVirtualDirectory(t, vhostID)

	// === CREATE ===
	t.Run("Create CORS Config", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/cors-configs", map[string]interface{}{
			"virtual_directory_id":   dirID,
			"allowed_origins":        "https://example.com,https://app.example.com",
			"allowed_methods":        "GET,POST,PUT,DELETE",
			"allowed_headers":        "Authorization,Content-Type,X-Request-ID",
			"exposed_headers":        "X-Total-Count",
			"allow_credentials":      true,
			"max_age_seconds":        7200,
			"is_active":             true,
		})
		AssertStatus(t, resp, http.StatusCreated)
		res := AssertSuccess(t, body)
		if ExtractID(res.Data) == 0 {
			t.Fatal("CORS Config ID should not be 0")
		}
	})

	// === CREATE - Validasi ===
	t.Run("Create CORS Config - Missing VDir", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/cors-configs", map[string]interface{}{
			"allowed_origins": "*",
		})
		AssertStatus(t, resp, http.StatusBadRequest)
		AssertError(t, body)
	})

	// === GET ALL ===
	t.Run("Get All CORS Configs", func(t *testing.T) {
		resp, body := ts.DoRequest("GET", "/api/admin/cors-configs", nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	// === GET BY ID ===
	t.Run("Get CORS Config By ID", func(t *testing.T) {
		resp, body := ts.DoRequest("GET", "/api/admin/cors-configs/1", nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	// === UPDATE ===
	t.Run("Update CORS Config", func(t *testing.T) {
		resp, body := ts.DoRequest("PUT", "/api/admin/cors-configs/1", map[string]interface{}{
			"allowed_origins":   "*",
			"allow_credentials": false,
			"max_age_seconds":   3600,
			"is_active":         true,
		})
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	// === DELETE ===
	t.Run("Delete CORS Config", func(t *testing.T) {
		resp, body := ts.DoRequest("DELETE", "/api/admin/cors-configs/1", nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})
}

// TestCircuitBreakersCRUD menguji CRUD lengkap untuk Circuit Breakers
func TestCircuitBreakersCRUD(t *testing.T) {
	ts := SetupTestServer(t)
	defer ts.Close()

	hostID := ts.CreateTestHost(t)
	vhostID := ts.CreateTestVirtualHost(t, hostID)
	dirID := ts.CreateTestVirtualDirectory(t, vhostID)

	// === CREATE ===
	t.Run("Create Circuit Breaker", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/circuit-breakers", map[string]interface{}{
			"virtual_directory_id":    dirID,
			"failure_threshold":       5,
			"recovery_timeout_seconds": 30,
			"half_open_max_requests":   3,
			"is_active":               true,
		})
		AssertStatus(t, resp, http.StatusCreated)
		res := AssertSuccess(t, body)
		if ExtractID(res.Data) == 0 {
			t.Fatal("Circuit Breaker ID should not be 0")
		}
	})

	// === CREATE - Validasi ===
	t.Run("Create Circuit Breaker - Missing VDir", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/circuit-breakers", map[string]interface{}{
			"failure_threshold": 5,
		})
		AssertStatus(t, resp, http.StatusBadRequest)
		AssertError(t, body)
	})

	// === GET ALL ===
	t.Run("Get All Circuit Breakers", func(t *testing.T) {
		resp, body := ts.DoRequest("GET", "/api/admin/circuit-breakers", nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	// === GET BY ID ===
	t.Run("Get Circuit Breaker By ID", func(t *testing.T) {
		resp, body := ts.DoRequest("GET", "/api/admin/circuit-breakers/1", nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	// === UPDATE ===
	t.Run("Update Circuit Breaker", func(t *testing.T) {
		resp, body := ts.DoRequest("PUT", "/api/admin/circuit-breakers/1", map[string]interface{}{
			"failure_threshold":        10,
			"recovery_timeout_seconds": 60,
			"is_active":               true,
		})
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	// === DELETE ===
	t.Run("Delete Circuit Breaker", func(t *testing.T) {
		resp, body := ts.DoRequest("DELETE", "/api/admin/circuit-breakers/1", nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})
}

// TestIPWhitelistsCRUD menguji CRUD lengkap untuk IP Whitelists
func TestIPWhitelistsCRUD(t *testing.T) {
	ts := SetupTestServer(t)
	defer ts.Close()

	hostID := ts.CreateTestHost(t)
	vhostID := ts.CreateTestVirtualHost(t, hostID)
	dirID := ts.CreateTestVirtualDirectory(t, vhostID)

	// === CREATE ===
	t.Run("Create IP Whitelist", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/ip-whitelists", map[string]interface{}{
			"virtual_directory_id": dirID,
			"ip_address":          "192.168.1.100",
			"description":         "Office IP",
			"is_active":           true,
		})
		AssertStatus(t, resp, http.StatusCreated)
		res := AssertSuccess(t, body)
		if ExtractID(res.Data) == 0 {
			t.Fatal("IP Whitelist ID should not be 0")
		}
	})

	// === CREATE - CIDR ===
	t.Run("Create IP Whitelist - CIDR", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/ip-whitelists", map[string]interface{}{
			"virtual_directory_id": dirID,
			"ip_address":          "10.0.0.0/24",
			"description":         "Internal network",
			"is_active":           true,
		})
		AssertStatus(t, resp, http.StatusCreated)
		AssertSuccess(t, body)
	})

	// === CREATE - Validasi ===
	t.Run("Create IP Whitelist - Missing IP", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/ip-whitelists", map[string]interface{}{
			"virtual_directory_id": dirID,
			"description":         "No IP",
		})
		AssertStatus(t, resp, http.StatusBadRequest)
		AssertError(t, body)
	})

	// === GET ALL ===
	t.Run("Get All IP Whitelists", func(t *testing.T) {
		resp, body := ts.DoRequest("GET", "/api/admin/ip-whitelists", nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	// === GET BY DIRECTORY ===
	t.Run("Get IP Whitelists By Directory", func(t *testing.T) {
		resp, body := ts.DoRequest("GET", fmt.Sprintf("/api/admin/virtual-directories/%d/ip-whitelists", dirID), nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	// === UPDATE ===
	t.Run("Update IP Whitelist", func(t *testing.T) {
		resp, body := ts.DoRequest("PUT", "/api/admin/ip-whitelists/1", map[string]interface{}{
			"ip_address":  "192.168.1.200",
			"description": "Updated Office IP",
			"is_active":   true,
		})
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	// === DELETE ===
	t.Run("Delete IP Whitelist", func(t *testing.T) {
		resp, body := ts.DoRequest("DELETE", "/api/admin/ip-whitelists/1", nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})
}

// TestIPBlacklistsCRUD menguji CRUD lengkap untuk IP Blacklists
func TestIPBlacklistsCRUD(t *testing.T) {
	ts := SetupTestServer(t)
	defer ts.Close()

	hostID := ts.CreateTestHost(t)
	vhostID := ts.CreateTestVirtualHost(t, hostID)
	dirID := ts.CreateTestVirtualDirectory(t, vhostID)

	// === CREATE ===
	t.Run("Create IP Blacklist", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/ip-blacklists", map[string]interface{}{
			"virtual_directory_id": dirID,
			"ip_address":          "203.0.113.1",
			"reason":              "Brute force attempt",
			"is_active":           true,
		})
		AssertStatus(t, resp, http.StatusCreated)
		res := AssertSuccess(t, body)
		if ExtractID(res.Data) == 0 {
			t.Fatal("IP Blacklist ID should not be 0")
		}
	})

	// === CREATE - With Expiry ===
	t.Run("Create IP Blacklist - With Expiry", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/ip-blacklists", map[string]interface{}{
			"virtual_directory_id": dirID,
			"ip_address":          "203.0.113.50",
			"reason":              "Temporary ban",
			"expired_at":          "2030-12-31 23:59:59",
			"is_active":           true,
		})
		AssertStatus(t, resp, http.StatusCreated)
		AssertSuccess(t, body)
	})

	// === CREATE - Validasi ===
	t.Run("Create IP Blacklist - Missing IP", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/ip-blacklists", map[string]interface{}{
			"virtual_directory_id": dirID,
			"reason":              "No IP provided",
		})
		AssertStatus(t, resp, http.StatusBadRequest)
		AssertError(t, body)
	})

	// === GET ALL ===
	t.Run("Get All IP Blacklists", func(t *testing.T) {
		resp, body := ts.DoRequest("GET", "/api/admin/ip-blacklists", nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	// === GET BY DIRECTORY ===
	t.Run("Get IP Blacklists By Directory", func(t *testing.T) {
		resp, body := ts.DoRequest("GET", fmt.Sprintf("/api/admin/virtual-directories/%d/ip-blacklists", dirID), nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	// === UPDATE ===
	t.Run("Update IP Blacklist", func(t *testing.T) {
		resp, body := ts.DoRequest("PUT", "/api/admin/ip-blacklists/1", map[string]interface{}{
			"ip_address": "203.0.113.5",
			"reason":     "Updated reason",
			"is_active":  true,
		})
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	// === DELETE ===
	t.Run("Delete IP Blacklist", func(t *testing.T) {
		resp, body := ts.DoRequest("DELETE", "/api/admin/ip-blacklists/1", nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})
}
