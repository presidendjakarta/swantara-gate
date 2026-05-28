package test

import (
	"net/http"
	"testing"
)

// TestUpstreamServersCRUD menguji CRUD lengkap untuk Upstream Servers
func TestUpstreamServersCRUD(t *testing.T) {
	ts := SetupTestServer(t)
	defer ts.Close()

	// Setup: buat host → virtual host
	hostID := ts.CreateTestHost(t)
	vhostID := ts.CreateTestVirtualHost(t, hostID)

	// === CREATE ===
	t.Run("Create Upstream Server", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/upstream-servers", map[string]interface{}{
			"virtual_host_id":              vhostID,
			"target_host":                  "192.168.1.10",
			"target_port":                  8080,
			"protocol":                     "http",
			"priority":                     1,
			"weight":                       10,
			"is_backup":                    false,
			"is_active":                    true,
			"health_check_enabled":         true,
			"health_check_path":            "/health",
			"health_check_interval_seconds": 30,
			"health_check_timeout_seconds":  5,
			"max_fails":                    3,
			"fail_timeout_seconds":         30,
		})
		AssertStatus(t, resp, http.StatusCreated)
		res := AssertSuccess(t, body)
		if ExtractID(res.Data) == 0 {
			t.Fatal("Upstream Server ID should not be 0")
		}
	})

	// === CREATE - Validasi ===
	t.Run("Create Upstream - Missing host", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/upstream-servers", map[string]interface{}{
			"virtual_host_id": vhostID,
			"target_port":     8080,
		})
		AssertStatus(t, resp, http.StatusBadRequest)
		AssertError(t, body)
	})

	// === GET ALL ===
	t.Run("Get All Upstream Servers", func(t *testing.T) {
		resp, body := ts.DoRequest("GET", "/api/admin/upstream-servers", nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	// === GET BY ID ===
	t.Run("Get Upstream Server By ID", func(t *testing.T) {
		resp, body := ts.DoRequest("GET", "/api/admin/upstream-servers/1", nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	// === GET BY VHOST ===
	t.Run("Get Upstream Servers By VHost", func(t *testing.T) {
		resp, body := ts.DoRequest("GET", "/api/admin/virtual-hosts/1/upstream-servers", nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	// === UPDATE ===
	t.Run("Update Upstream Server", func(t *testing.T) {
		resp, body := ts.DoRequest("PUT", "/api/admin/upstream-servers/1", map[string]interface{}{
			"target_host": "192.168.1.20",
			"target_port": 9090,
			"protocol":    "https",
			"is_active":   true,
			"weight":      20,
		})
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	// === DELETE ===
	t.Run("Delete Upstream Server", func(t *testing.T) {
		resp, body := ts.DoRequest("DELETE", "/api/admin/upstream-servers/1", nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})
}

// TestVirtualDirectoriesCRUD menguji CRUD lengkap untuk Virtual Directories
func TestVirtualDirectoriesCRUD(t *testing.T) {
	ts := SetupTestServer(t)
	defer ts.Close()

	// Setup
	hostID := ts.CreateTestHost(t)
	vhostID := ts.CreateTestVirtualHost(t, hostID)

	// === CREATE ===
	t.Run("Create Virtual Directory", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/virtual-directories", map[string]interface{}{
			"virtual_host_id":      vhostID,
			"source_path":          "/api/v1/users",
			"target_path":          "/users",
			"match_type":           "prefix",
			"strip_prefix":         true,
			"preserve_host_header": false,
			"auth_type":            "api_key",
			"is_active":            true,
			"proxy_timeout_seconds": 30,
			"retry_count":          2,
			"max_request_size_mb":  10,
		})
		AssertStatus(t, resp, http.StatusCreated)
		res := AssertSuccess(t, body)
		if ExtractID(res.Data) == 0 {
			t.Fatal("Virtual Directory ID should not be 0")
		}
	})

	// === GET ALL ===
	t.Run("Get All Virtual Directories", func(t *testing.T) {
		resp, body := ts.DoRequest("GET", "/api/admin/virtual-directories", nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	// === GET BY ID ===
	t.Run("Get Virtual Directory By ID", func(t *testing.T) {
		resp, body := ts.DoRequest("GET", "/api/admin/virtual-directories/1", nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	// === GET BY VHOST ===
	t.Run("Get VDirs By VHost", func(t *testing.T) {
		resp, body := ts.DoRequest("GET", "/api/admin/virtual-hosts/1/virtual-directories", nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	// === UPDATE ===
	t.Run("Update Virtual Directory", func(t *testing.T) {
		resp, body := ts.DoRequest("PUT", "/api/admin/virtual-directories/1", map[string]interface{}{
			"source_path": "/api/v2/users",
			"target_path": "/v2/users",
			"is_active":   true,
		})
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	// === DELETE ===
	t.Run("Delete Virtual Directory", func(t *testing.T) {
		resp, body := ts.DoRequest("DELETE", "/api/admin/virtual-directories/1", nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})
}

// TestVirtualDirectoryMethodsCRUD menguji methods pada virtual directory
func TestVirtualDirectoryMethodsCRUD(t *testing.T) {
	ts := SetupTestServer(t)
	defer ts.Close()

	// Setup
	hostID := ts.CreateTestHost(t)
	vhostID := ts.CreateTestVirtualHost(t, hostID)
	dirID := ts.CreateTestVirtualDirectory(t, vhostID)
	_ = dirID

	// === SET METHODS ===
	t.Run("Set Methods", func(t *testing.T) {
		resp, body := ts.DoRequest("PUT", "/api/admin/virtual-directories/1/methods", map[string]interface{}{
			"methods": []string{"GET", "POST", "PUT", "DELETE"},
		})
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	// === GET METHODS ===
	t.Run("Get Methods", func(t *testing.T) {
		resp, body := ts.DoRequest("GET", "/api/admin/virtual-directories/1/methods", nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	// === SET METHODS - Invalid ===
	t.Run("Set Methods - Invalid Method", func(t *testing.T) {
		resp, body := ts.DoRequest("PUT", "/api/admin/virtual-directories/1/methods", map[string]interface{}{
			"methods": []string{"GET", "INVALID"},
		})
		AssertStatus(t, resp, http.StatusBadRequest)
		AssertError(t, body)
	})
}

// TestConsumerCredentialsCRUD menguji CRUD untuk Consumer Credentials
func TestConsumerCredentialsCRUD(t *testing.T) {
	ts := SetupTestServer(t)
	defer ts.Close()

	// Setup: buat consumer dulu
	consumerID := ts.CreateTestConsumer(t)

	// === CREATE - Basic Auth ===
	t.Run("Create Credential - Basic Auth", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/consumer-credentials", map[string]interface{}{
			"consumer_id": consumerID,
			"auth_type":   "basic",
			"username":    "app_user",
			"password":    "secret123",
			"is_active":   true,
		})
		AssertStatus(t, resp, http.StatusCreated)
		AssertSuccess(t, body)
	})

	// === CREATE - Validasi ===
	t.Run("Create Credential - Invalid Type", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/consumer-credentials", map[string]interface{}{
			"consumer_id": consumerID,
			"auth_type":   "invalid_type",
			"is_active":   true,
		})
		AssertStatus(t, resp, http.StatusBadRequest)
		AssertError(t, body)
	})

	// === GET ALL ===
	t.Run("Get All Credentials", func(t *testing.T) {
		resp, body := ts.DoRequest("GET", "/api/admin/consumer-credentials", nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	// === GET BY CONSUMER ===
	t.Run("Get Credentials By Consumer", func(t *testing.T) {
		resp, body := ts.DoRequest("GET", "/api/admin/consumers/1/credentials", nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	// === UPDATE ===
	t.Run("Update Credential", func(t *testing.T) {
		resp, body := ts.DoRequest("PUT", "/api/admin/consumer-credentials/1", map[string]interface{}{
			"username":  "new_user",
			"password":  "new_pass",
			"is_active": true,
		})
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	// === DELETE ===
	t.Run("Delete Credential", func(t *testing.T) {
		resp, body := ts.DoRequest("DELETE", "/api/admin/consumer-credentials/1", nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})
}

// TestAPIKeysCRUD menguji CRUD untuk API Keys
func TestAPIKeysCRUD(t *testing.T) {
	ts := SetupTestServer(t)
	defer ts.Close()

	consumerID := ts.CreateTestConsumer(t)

	// === CREATE ===
	t.Run("Create API Key", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/api-keys", map[string]interface{}{
			"consumer_id":        consumerID,
			"description":        "Production Key",
			"rate_limit_override": 1000,
			"is_active":          true,
		})
		AssertStatus(t, resp, http.StatusCreated)
		AssertSuccess(t, body)
	})

	// === GET ALL ===
	t.Run("Get All API Keys", func(t *testing.T) {
		resp, body := ts.DoRequest("GET", "/api/admin/api-keys", nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	// === GET BY CONSUMER ===
	t.Run("Get API Keys By Consumer", func(t *testing.T) {
		resp, body := ts.DoRequest("GET", "/api/admin/consumers/1/api-keys", nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	// === UPDATE ===
	t.Run("Update API Key", func(t *testing.T) {
		resp, body := ts.DoRequest("PUT", "/api/admin/api-keys/1", map[string]interface{}{
			"description": "Updated Key",
			"is_active":   true,
		})
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	// === DELETE ===
	t.Run("Delete API Key", func(t *testing.T) {
		resp, body := ts.DoRequest("DELETE", "/api/admin/api-keys/1", nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})
}

// TestRouteConsumerAccessCRUD menguji CRUD untuk Route Consumer Access (ACL)
func TestRouteConsumerAccessCRUD(t *testing.T) {
	ts := SetupTestServer(t)
	defer ts.Close()

	// Setup
	hostID := ts.CreateTestHost(t)
	vhostID := ts.CreateTestVirtualHost(t, hostID)
	dirID := ts.CreateTestVirtualDirectory(t, vhostID)
	consumerID := ts.CreateTestConsumer(t)

	// === CREATE ===
	t.Run("Create Route Access", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/route-access", map[string]interface{}{
			"virtual_directory_id": dirID,
			"consumer_id":         consumerID,
			"is_active":           true,
		})
		AssertStatus(t, resp, http.StatusCreated)
		AssertSuccess(t, body)
	})

	// === GET ALL ===
	t.Run("Get All Route Access", func(t *testing.T) {
		resp, body := ts.DoRequest("GET", "/api/admin/route-access", nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	// === GET BY ID ===
	t.Run("Get Route Access By ID", func(t *testing.T) {
		resp, body := ts.DoRequest("GET", "/api/admin/route-access/1", nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	// === UPDATE ===
	t.Run("Update Route Access", func(t *testing.T) {
		resp, body := ts.DoRequest("PUT", "/api/admin/route-access/1", map[string]interface{}{
			"is_active": false,
		})
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	// === DELETE ===
	t.Run("Delete Route Access", func(t *testing.T) {
		resp, body := ts.DoRequest("DELETE", "/api/admin/route-access/1", nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})
}
