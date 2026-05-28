package test

import (
	"net/http"
	"testing"
)

// TestUsersCRUD menguji CRUD lengkap untuk Users
func TestUsersCRUD(t *testing.T) {
	ts := SetupTestServer(t)
	defer ts.Close()

	var userID int64

	// === CREATE ===
	t.Run("Create User", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/users", map[string]interface{}{
			"username":  "admin_test",
			"password":  "password123",
			"full_name": "Admin Test",
			"email":     "admin@test.com",
			"role":      "admin",
			"is_active": true,
		})
		AssertStatus(t, resp, http.StatusCreated)
		res := AssertSuccess(t, body)
		userID = ExtractID(res.Data)
		if userID == 0 {
			t.Fatal("User ID should not be 0")
		}
	})

	// === CREATE - Validasi ===
	t.Run("Create User - Missing Username", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/users", map[string]interface{}{
			"password": "pass",
		})
		AssertStatus(t, resp, http.StatusBadRequest)
		AssertError(t, body)
	})

	// === GET BY ID ===
	t.Run("Get User By ID", func(t *testing.T) {
		resp, body := ts.DoRequest("GET", "/api/admin/users/1", nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	// === GET ALL ===
	t.Run("Get All Users", func(t *testing.T) {
		resp, body := ts.DoRequest("GET", "/api/admin/users?page=1&limit=10", nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	// === UPDATE ===
	t.Run("Update User", func(t *testing.T) {
		resp, body := ts.DoRequest("PUT", "/api/admin/users/1", map[string]interface{}{
			"full_name": "Admin Updated",
			"email":     "updated@test.com",
			"role":      "admin",
			"is_active": true,
		})
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	// === GET NOT FOUND ===
	t.Run("Get User - Not Found", func(t *testing.T) {
		resp, body := ts.DoRequest("GET", "/api/admin/users/999", nil)
		AssertStatus(t, resp, http.StatusNotFound)
		AssertError(t, body)
	})

	// === DELETE ===
	t.Run("Delete User", func(t *testing.T) {
		resp, body := ts.DoRequest("DELETE", "/api/admin/users/1", nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})
}

// TestConsumersCRUD menguji CRUD lengkap untuk API Consumers
func TestConsumersCRUD(t *testing.T) {
	ts := SetupTestServer(t)
	defer ts.Close()

	// === CREATE ===
	t.Run("Create Consumer", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/consumers", map[string]interface{}{
			"consumer_name": "Mobile App",
			"description":   "Aplikasi mobile",
			"contact_email": "dev@app.com",
			"is_active":     true,
		})
		AssertStatus(t, resp, http.StatusCreated)
		res := AssertSuccess(t, body)
		if ExtractID(res.Data) == 0 {
			t.Fatal("Consumer ID should not be 0")
		}
	})

	// === CREATE - Validasi ===
	t.Run("Create Consumer - Missing Name", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/consumers", map[string]interface{}{
			"description": "No name",
		})
		AssertStatus(t, resp, http.StatusBadRequest)
		AssertError(t, body)
	})

	// === GET ALL ===
	t.Run("Get All Consumers", func(t *testing.T) {
		resp, body := ts.DoRequest("GET", "/api/admin/consumers", nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	// === GET BY ID ===
	t.Run("Get Consumer By ID", func(t *testing.T) {
		resp, body := ts.DoRequest("GET", "/api/admin/consumers/1", nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	// === UPDATE ===
	t.Run("Update Consumer", func(t *testing.T) {
		resp, body := ts.DoRequest("PUT", "/api/admin/consumers/1", map[string]interface{}{
			"description":   "Updated description",
			"contact_email": "new@app.com",
			"is_active":     true,
		})
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	// === DELETE ===
	t.Run("Delete Consumer", func(t *testing.T) {
		resp, body := ts.DoRequest("DELETE", "/api/admin/consumers/1", nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})
}

// TestHostsCRUD menguji CRUD lengkap untuk Hosts
func TestHostsCRUD(t *testing.T) {
	ts := SetupTestServer(t)
	defer ts.Close()

	// === CREATE ===
	t.Run("Create Host", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/hosts", map[string]interface{}{
			"host_name":   "api.example.com",
			"description": "Main API host",
			"is_active":   true,
		})
		AssertStatus(t, resp, http.StatusCreated)
		res := AssertSuccess(t, body)
		if ExtractID(res.Data) == 0 {
			t.Fatal("Host ID should not be 0")
		}
	})

	// === CREATE - Validasi ===
	t.Run("Create Host - Missing Name", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/hosts", map[string]interface{}{
			"description": "No name",
		})
		AssertStatus(t, resp, http.StatusBadRequest)
		AssertError(t, body)
	})

	// === GET ALL ===
	t.Run("Get All Hosts", func(t *testing.T) {
		resp, body := ts.DoRequest("GET", "/api/admin/hosts", nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	// === GET BY ID ===
	t.Run("Get Host By ID", func(t *testing.T) {
		resp, body := ts.DoRequest("GET", "/api/admin/hosts/1", nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	// === UPDATE ===
	t.Run("Update Host", func(t *testing.T) {
		resp, body := ts.DoRequest("PUT", "/api/admin/hosts/1", map[string]interface{}{
			"description": "Updated host",
			"is_active":   true,
		})
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	// === DELETE ===
	t.Run("Delete Host", func(t *testing.T) {
		resp, body := ts.DoRequest("DELETE", "/api/admin/hosts/1", nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})
}

// TestVirtualHostsCRUD menguji CRUD lengkap untuk Virtual Hosts
func TestVirtualHostsCRUD(t *testing.T) {
	ts := SetupTestServer(t)
	defer ts.Close()

	// Buat host parent dulu
	hostID := ts.CreateTestHost(t)

	// === CREATE ===
	t.Run("Create Virtual Host", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/virtual-hosts", map[string]interface{}{
			"host_id":        hostID,
			"vhost_name":     "api-v1.example.com",
			"lb_algorithm":   "round_robin",
			"sticky_session": false,
			"failover_mode":  "active_passive",
			"is_active":      true,
		})
		AssertStatus(t, resp, http.StatusCreated)
		res := AssertSuccess(t, body)
		if ExtractID(res.Data) == 0 {
			t.Fatal("Virtual Host ID should not be 0")
		}
	})

	// === GET ALL ===
	t.Run("Get All Virtual Hosts", func(t *testing.T) {
		resp, body := ts.DoRequest("GET", "/api/admin/virtual-hosts", nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	// === GET BY ID ===
	t.Run("Get Virtual Host By ID", func(t *testing.T) {
		resp, body := ts.DoRequest("GET", "/api/admin/virtual-hosts/1", nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	// === UPDATE ===
	t.Run("Update Virtual Host", func(t *testing.T) {
		resp, body := ts.DoRequest("PUT", "/api/admin/virtual-hosts/1", map[string]interface{}{
			"lb_algorithm":   "least_conn",
			"sticky_session": true,
			"failover_mode":  "active_active",
			"is_active":      true,
		})
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	// === DELETE ===
	t.Run("Delete Virtual Host", func(t *testing.T) {
		resp, body := ts.DoRequest("DELETE", "/api/admin/virtual-hosts/1", nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})
}
