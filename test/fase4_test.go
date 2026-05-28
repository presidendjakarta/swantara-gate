package test

import (
	"fmt"
	"net/http"
	"testing"
)

// TestRequestHeaderRulesCRUD menguji CRUD lengkap untuk Request Header Rules
func TestRequestHeaderRulesCRUD(t *testing.T) {
	ts := SetupTestServer(t)
	defer ts.Close()

	// Setup
	hostID := ts.CreateTestHost(t)
	vhostID := ts.CreateTestVirtualHost(t, hostID)
	dirID := ts.CreateTestVirtualDirectory(t, vhostID)

	// === CREATE - Set static header ===
	t.Run("Create Request Header - Set Static", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/request-headers", map[string]interface{}{
			"virtual_directory_id": dirID,
			"header_name":         "X-Request-ID",
			"operation":           "set",
			"value_source":        "static",
			"header_value":        "gateway-123",
			"execution_order":     1,
			"is_active":           true,
		})
		AssertStatus(t, resp, http.StatusCreated)
		res := AssertSuccess(t, body)
		if ExtractID(res.Data) == 0 {
			t.Fatal("Request Header Rule ID should not be 0")
		}
	})

	// === CREATE - Copy from another header ===
	t.Run("Create Request Header - Copy Header", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/request-headers", map[string]interface{}{
			"virtual_directory_id": dirID,
			"header_name":         "X-Forwarded-For",
			"operation":           "copy",
			"value_source":        "header",
			"source_header":       "X-Real-IP",
			"execution_order":     2,
			"is_active":           true,
		})
		AssertStatus(t, resp, http.StatusCreated)
		AssertSuccess(t, body)
	})

	// === CREATE - Remove header ===
	t.Run("Create Request Header - Remove", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/request-headers", map[string]interface{}{
			"virtual_directory_id": dirID,
			"header_name":         "X-Debug",
			"operation":           "remove",
			"execution_order":     3,
			"is_active":           true,
		})
		AssertStatus(t, resp, http.StatusCreated)
		AssertSuccess(t, body)
	})

	// === CREATE - Validasi: missing header_name ===
	t.Run("Create Request Header - Missing Name", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/request-headers", map[string]interface{}{
			"virtual_directory_id": dirID,
			"operation":           "set",
		})
		AssertStatus(t, resp, http.StatusBadRequest)
		AssertError(t, body)
	})

	// === CREATE - Validasi: invalid operation ===
	t.Run("Create Request Header - Invalid Operation", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/request-headers", map[string]interface{}{
			"virtual_directory_id": dirID,
			"header_name":         "X-Test",
			"operation":           "invalid_op",
		})
		AssertStatus(t, resp, http.StatusBadRequest)
		AssertError(t, body)
	})

	// === CREATE - Validasi: invalid value_source ===
	t.Run("Create Request Header - Invalid ValueSource", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/request-headers", map[string]interface{}{
			"virtual_directory_id": dirID,
			"header_name":         "X-Test",
			"operation":           "set",
			"value_source":        "invalid_source",
		})
		AssertStatus(t, resp, http.StatusBadRequest)
		AssertError(t, body)
	})

	// === GET ALL ===
	t.Run("Get All Request Headers", func(t *testing.T) {
		resp, body := ts.DoRequest("GET", "/api/admin/request-headers?page=1&limit=10", nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	// === GET BY ID ===
	t.Run("Get Request Header By ID", func(t *testing.T) {
		resp, body := ts.DoRequest("GET", "/api/admin/request-headers/1", nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	// === GET BY DIRECTORY ===
	t.Run("Get Request Headers By Directory", func(t *testing.T) {
		resp, body := ts.DoRequest("GET", fmt.Sprintf("/api/admin/virtual-directories/%d/request-headers", dirID), nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	// === UPDATE ===
	t.Run("Update Request Header", func(t *testing.T) {
		resp, body := ts.DoRequest("PUT", "/api/admin/request-headers/1", map[string]interface{}{
			"header_name":     "X-Request-ID-V2",
			"operation":       "set",
			"value_source":    "static",
			"header_value":    "gateway-v2",
			"execution_order": 1,
			"is_active":       true,
		})
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	// === UPDATE - Invalid operation ===
	t.Run("Update Request Header - Invalid Op", func(t *testing.T) {
		resp, body := ts.DoRequest("PUT", "/api/admin/request-headers/1", map[string]interface{}{
			"operation": "badop",
		})
		AssertStatus(t, resp, http.StatusBadRequest)
		AssertError(t, body)
	})

	// === DELETE ===
	t.Run("Delete Request Header", func(t *testing.T) {
		resp, body := ts.DoRequest("DELETE", "/api/admin/request-headers/1", nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	// === DELETE - Not Found ===
	t.Run("Delete Request Header - Not Found", func(t *testing.T) {
		resp, body := ts.DoRequest("DELETE", "/api/admin/request-headers/999", nil)
		AssertStatus(t, resp, http.StatusBadRequest)
		AssertError(t, body)
	})
}

// TestResponseHeaderRulesCRUD menguji CRUD lengkap untuk Response Header Rules
func TestResponseHeaderRulesCRUD(t *testing.T) {
	ts := SetupTestServer(t)
	defer ts.Close()

	// Setup
	hostID := ts.CreateTestHost(t)
	vhostID := ts.CreateTestVirtualHost(t, hostID)
	dirID := ts.CreateTestVirtualDirectory(t, vhostID)

	// === CREATE - Set header ===
	t.Run("Create Response Header - Set", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/response-headers", map[string]interface{}{
			"virtual_directory_id": dirID,
			"header_name":         "X-Powered-By",
			"operation":           "set",
			"header_value":        "Swantara-Gate/1.0",
			"execution_order":     1,
			"is_active":           true,
		})
		AssertStatus(t, resp, http.StatusCreated)
		res := AssertSuccess(t, body)
		if ExtractID(res.Data) == 0 {
			t.Fatal("Response Header Rule ID should not be 0")
		}
	})

	// === CREATE - Remove header ===
	t.Run("Create Response Header - Remove Server", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/response-headers", map[string]interface{}{
			"virtual_directory_id": dirID,
			"header_name":         "Server",
			"operation":           "remove",
			"execution_order":     2,
			"is_active":           true,
		})
		AssertStatus(t, resp, http.StatusCreated)
		AssertSuccess(t, body)
	})

	// === CREATE - Add security headers ===
	t.Run("Create Response Header - Security", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/response-headers", map[string]interface{}{
			"virtual_directory_id": dirID,
			"header_name":         "X-Content-Type-Options",
			"operation":           "set",
			"header_value":        "nosniff",
			"execution_order":     3,
			"is_active":           true,
		})
		AssertStatus(t, resp, http.StatusCreated)
		AssertSuccess(t, body)
	})

	// === CREATE - Validasi: missing header_name ===
	t.Run("Create Response Header - Missing Name", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/response-headers", map[string]interface{}{
			"virtual_directory_id": dirID,
			"operation":           "set",
		})
		AssertStatus(t, resp, http.StatusBadRequest)
		AssertError(t, body)
	})

	// === CREATE - Validasi: invalid operation ===
	t.Run("Create Response Header - Invalid Operation", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/response-headers", map[string]interface{}{
			"virtual_directory_id": dirID,
			"header_name":         "X-Test",
			"operation":           "invalid",
		})
		AssertStatus(t, resp, http.StatusBadRequest)
		AssertError(t, body)
	})

	// === GET ALL ===
	t.Run("Get All Response Headers", func(t *testing.T) {
		resp, body := ts.DoRequest("GET", "/api/admin/response-headers", nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	// === GET BY ID ===
	t.Run("Get Response Header By ID", func(t *testing.T) {
		resp, body := ts.DoRequest("GET", "/api/admin/response-headers/1", nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	// === GET BY DIRECTORY ===
	t.Run("Get Response Headers By Directory", func(t *testing.T) {
		resp, body := ts.DoRequest("GET", fmt.Sprintf("/api/admin/virtual-directories/%d/response-headers", dirID), nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	// === UPDATE ===
	t.Run("Update Response Header", func(t *testing.T) {
		resp, body := ts.DoRequest("PUT", "/api/admin/response-headers/1", map[string]interface{}{
			"header_name":     "X-Powered-By",
			"operation":       "set",
			"header_value":    "Swantara-Gate/2.0",
			"execution_order": 1,
			"is_active":       true,
		})
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	// === DELETE ===
	t.Run("Delete Response Header", func(t *testing.T) {
		resp, body := ts.DoRequest("DELETE", "/api/admin/response-headers/1", nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})
}

// TestQueryRewritesCRUD menguji CRUD lengkap untuk Query Rewrites
func TestQueryRewritesCRUD(t *testing.T) {
	ts := SetupTestServer(t)
	defer ts.Close()

	// Setup
	hostID := ts.CreateTestHost(t)
	vhostID := ts.CreateTestVirtualHost(t, hostID)
	dirID := ts.CreateTestVirtualDirectory(t, vhostID)

	// === CREATE - Set param ===
	t.Run("Create Query Rewrite - Set", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/query-rewrites", map[string]interface{}{
			"virtual_directory_id": dirID,
			"param_name":          "version",
			"param_value":         "v2",
			"operation":           "set",
		})
		AssertStatus(t, resp, http.StatusCreated)
		res := AssertSuccess(t, body)
		if ExtractID(res.Data) == 0 {
			t.Fatal("Query Rewrite ID should not be 0")
		}
	})

	// === CREATE - Add param ===
	t.Run("Create Query Rewrite - Add", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/query-rewrites", map[string]interface{}{
			"virtual_directory_id": dirID,
			"param_name":          "source",
			"param_value":         "gateway",
			"operation":           "add",
		})
		AssertStatus(t, resp, http.StatusCreated)
		AssertSuccess(t, body)
	})

	// === CREATE - Remove param ===
	t.Run("Create Query Rewrite - Remove", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/query-rewrites", map[string]interface{}{
			"virtual_directory_id": dirID,
			"param_name":          "debug",
			"operation":           "remove",
		})
		AssertStatus(t, resp, http.StatusCreated)
		AssertSuccess(t, body)
	})

	// === CREATE - Default operation (set) ===
	t.Run("Create Query Rewrite - Default Operation", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/query-rewrites", map[string]interface{}{
			"virtual_directory_id": dirID,
			"param_name":          "format",
			"param_value":         "json",
		})
		AssertStatus(t, resp, http.StatusCreated)
		AssertSuccess(t, body)
	})

	// === CREATE - Validasi: missing param_name ===
	t.Run("Create Query Rewrite - Missing Param", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/query-rewrites", map[string]interface{}{
			"virtual_directory_id": dirID,
			"param_value":         "test",
		})
		AssertStatus(t, resp, http.StatusBadRequest)
		AssertError(t, body)
	})

	// === CREATE - Validasi: invalid operation ===
	t.Run("Create Query Rewrite - Invalid Operation", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/query-rewrites", map[string]interface{}{
			"virtual_directory_id": dirID,
			"param_name":          "test",
			"operation":           "invalid_op",
		})
		AssertStatus(t, resp, http.StatusBadRequest)
		AssertError(t, body)
	})

	// === GET ALL ===
	t.Run("Get All Query Rewrites", func(t *testing.T) {
		resp, body := ts.DoRequest("GET", "/api/admin/query-rewrites?page=1&limit=10", nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	// === GET BY ID ===
	t.Run("Get Query Rewrite By ID", func(t *testing.T) {
		resp, body := ts.DoRequest("GET", "/api/admin/query-rewrites/1", nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	// === GET BY DIRECTORY ===
	t.Run("Get Query Rewrites By Directory", func(t *testing.T) {
		resp, body := ts.DoRequest("GET", fmt.Sprintf("/api/admin/virtual-directories/%d/query-rewrites", dirID), nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	// === UPDATE ===
	t.Run("Update Query Rewrite", func(t *testing.T) {
		resp, body := ts.DoRequest("PUT", "/api/admin/query-rewrites/1", map[string]interface{}{
			"param_name":  "version",
			"param_value": "v3",
			"operation":   "set",
		})
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	// === UPDATE - Invalid operation ===
	t.Run("Update Query Rewrite - Invalid Op", func(t *testing.T) {
		resp, body := ts.DoRequest("PUT", "/api/admin/query-rewrites/1", map[string]interface{}{
			"operation": "badop",
		})
		AssertStatus(t, resp, http.StatusBadRequest)
		AssertError(t, body)
	})

	// === DELETE ===
	t.Run("Delete Query Rewrite", func(t *testing.T) {
		resp, body := ts.DoRequest("DELETE", "/api/admin/query-rewrites/1", nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	// === DELETE - Not Found ===
	t.Run("Delete Query Rewrite - Not Found", func(t *testing.T) {
		resp, body := ts.DoRequest("DELETE", "/api/admin/query-rewrites/999", nil)
		AssertStatus(t, resp, http.StatusBadRequest)
		AssertError(t, body)
	})
}
