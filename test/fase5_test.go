package test

import (
	"fmt"
	"net/http"
	"testing"
)

// =========================================================
// TEST ACME ACCOUNTS
// =========================================================

func TestACMEAccountsCRUD(t *testing.T) {
	ts := SetupTestServer(t)
	defer ts.Close()

	var acmeID int64

	t.Run("Create ACME account", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/acme-accounts", map[string]interface{}{
			"email":            "admin@example.com",
			"provider_url":     "https://acme-v02.api.letsencrypt.org/directory",
			"account_key_path": "/etc/ssl/acme/account.key",
			"is_default":       true,
		})
		AssertStatus(t, resp, http.StatusCreated)
		res := AssertSuccess(t, body)
		acmeID = ExtractID(res.Data)
		if acmeID == 0 {
			t.Fatal("Expected valid ACME account ID")
		}
	})

	t.Run("Create validation - missing email", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/acme-accounts", map[string]interface{}{
			"provider_url":     "https://acme.example.com",
			"account_key_path": "/keys/acme.key",
		})
		AssertStatus(t, resp, http.StatusBadRequest)
		AssertError(t, body)
	})

	t.Run("Get all ACME accounts", func(t *testing.T) {
		resp, body := ts.DoRequest("GET", "/api/admin/acme-accounts", nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	t.Run("Get ACME account by ID", func(t *testing.T) {
		resp, body := ts.DoRequest("GET", fmt.Sprintf("/api/admin/acme-accounts/%d", acmeID), nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	t.Run("Update ACME account", func(t *testing.T) {
		resp, body := ts.DoRequest("PUT", fmt.Sprintf("/api/admin/acme-accounts/%d", acmeID), map[string]interface{}{
			"email":            "updated@example.com",
			"provider_url":     "https://acme-staging-v02.api.letsencrypt.org/directory",
			"account_key_path": "/etc/ssl/acme/staging.key",
			"is_default":       false,
		})
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	t.Run("Delete ACME account", func(t *testing.T) {
		resp, body := ts.DoRequest("DELETE", fmt.Sprintf("/api/admin/acme-accounts/%d", acmeID), nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	t.Run("Delete not found", func(t *testing.T) {
		resp, body := ts.DoRequest("DELETE", "/api/admin/acme-accounts/9999", nil)
		AssertStatus(t, resp, http.StatusBadRequest)
		AssertError(t, body)
	})
}

// =========================================================
// TEST SSL CERTIFICATES
// =========================================================

func TestSSLCertificatesCRUD(t *testing.T) {
	ts := SetupTestServer(t)
	defer ts.Close()

	// Create prerequisite ACME account
	_, acmeBody := ts.DoRequest("POST", "/api/admin/acme-accounts", map[string]interface{}{
		"email":            "ssl@example.com",
		"provider_url":     "https://acme-v02.api.letsencrypt.org/directory",
		"account_key_path": "/keys/acme.key",
		"is_default":       true,
	})
	acmeRes := ParseResponse(acmeBody)
	acmeID := ExtractID(acmeRes.Data)

	var certID int64

	t.Run("Create SSL certificate", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/ssl-certificates", map[string]interface{}{
			"acme_account_id":   acmeID,
			"certificate_path":  "/etc/ssl/certs/domain.crt",
			"private_key_path":  "/etc/ssl/private/domain.key",
			"chain_path":        "/etc/ssl/certs/chain.pem",
			"auto_renew":        true,
			"renew_before_days": 14,
			"is_active":         true,
		})
		AssertStatus(t, resp, http.StatusCreated)
		res := AssertSuccess(t, body)
		certID = ExtractID(res.Data)
		if certID == 0 {
			t.Fatal("Expected valid SSL certificate ID")
		}
	})

	t.Run("Create with defaults", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/ssl-certificates", map[string]interface{}{
			"certificate_path": "/etc/ssl/certs/auto.crt",
			"private_key_path": "/etc/ssl/private/auto.key",
		})
		AssertStatus(t, resp, http.StatusCreated)
		AssertSuccess(t, body)
	})

	t.Run("Create validation - missing cert path", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/ssl-certificates", map[string]interface{}{
			"private_key_path": "/etc/ssl/private/domain.key",
		})
		AssertStatus(t, resp, http.StatusBadRequest)
		AssertError(t, body)
	})

	t.Run("Create validation - missing key path", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/ssl-certificates", map[string]interface{}{
			"certificate_path": "/etc/ssl/certs/domain.crt",
		})
		AssertStatus(t, resp, http.StatusBadRequest)
		AssertError(t, body)
	})

	t.Run("Get all SSL certificates", func(t *testing.T) {
		resp, body := ts.DoRequest("GET", "/api/admin/ssl-certificates", nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	t.Run("Get SSL certificate by ID", func(t *testing.T) {
		resp, body := ts.DoRequest("GET", fmt.Sprintf("/api/admin/ssl-certificates/%d", certID), nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	t.Run("Update SSL certificate", func(t *testing.T) {
		resp, body := ts.DoRequest("PUT", fmt.Sprintf("/api/admin/ssl-certificates/%d", certID), map[string]interface{}{
			"provider":      "custom",
			"renew_status":  "pending",
			"is_active":     false,
		})
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	t.Run("Delete SSL certificate", func(t *testing.T) {
		resp, body := ts.DoRequest("DELETE", fmt.Sprintf("/api/admin/ssl-certificates/%d", certID), nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})
}

// =========================================================
// TEST CERTIFICATE DOMAINS
// =========================================================

func TestCertificateDomainsCRUD(t *testing.T) {
	ts := SetupTestServer(t)
	defer ts.Close()

	// Create prerequisite SSL certificate
	_, certBody := ts.DoRequest("POST", "/api/admin/ssl-certificates", map[string]interface{}{
		"certificate_path": "/etc/ssl/certs/test.crt",
		"private_key_path": "/etc/ssl/private/test.key",
	})
	certRes := ParseResponse(certBody)
	certID := ExtractID(certRes.Data)

	var domainID int64

	t.Run("Create certificate domain", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/certificate-domains", map[string]interface{}{
			"ssl_certificate_id": certID,
			"domain_name":        "example.com",
			"is_wildcard":        false,
		})
		AssertStatus(t, resp, http.StatusCreated)
		res := AssertSuccess(t, body)
		domainID = ExtractID(res.Data)
		if domainID == 0 {
			t.Fatal("Expected valid domain ID")
		}
	})

	t.Run("Create wildcard domain", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/certificate-domains", map[string]interface{}{
			"ssl_certificate_id": certID,
			"domain_name":        "*.example.com",
			"is_wildcard":        true,
		})
		AssertStatus(t, resp, http.StatusCreated)
		AssertSuccess(t, body)
	})

	t.Run("Create validation - missing domain", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/certificate-domains", map[string]interface{}{
			"ssl_certificate_id": certID,
		})
		AssertStatus(t, resp, http.StatusBadRequest)
		AssertError(t, body)
	})

	t.Run("Create validation - missing cert id", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/certificate-domains", map[string]interface{}{
			"domain_name": "orphan.com",
		})
		AssertStatus(t, resp, http.StatusBadRequest)
		AssertError(t, body)
	})

	t.Run("Get all certificate domains", func(t *testing.T) {
		resp, body := ts.DoRequest("GET", "/api/admin/certificate-domains", nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	t.Run("Get domain by ID", func(t *testing.T) {
		resp, body := ts.DoRequest("GET", fmt.Sprintf("/api/admin/certificate-domains/%d", domainID), nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	t.Run("Get domains by certificate", func(t *testing.T) {
		resp, body := ts.DoRequest("GET", fmt.Sprintf("/api/admin/ssl-certificates/%d/domains", certID), nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	t.Run("Update domain", func(t *testing.T) {
		resp, body := ts.DoRequest("PUT", fmt.Sprintf("/api/admin/certificate-domains/%d", domainID), map[string]interface{}{
			"domain_name": "updated.example.com",
			"is_wildcard": true,
		})
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	t.Run("Delete domain", func(t *testing.T) {
		resp, body := ts.DoRequest("DELETE", fmt.Sprintf("/api/admin/certificate-domains/%d", domainID), nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})
}

// =========================================================
// TEST SSL CERTIFICATE BINDINGS
// =========================================================

func TestSSLCertificateBindingsCRUD(t *testing.T) {
	ts := SetupTestServer(t)
	defer ts.Close()

	// Create prerequisites
	hostID := ts.CreateTestHost(t)
	vhostID := ts.CreateTestVirtualHost(t, hostID)

	_, certBody := ts.DoRequest("POST", "/api/admin/ssl-certificates", map[string]interface{}{
		"certificate_path": "/etc/ssl/certs/bind.crt",
		"private_key_path": "/etc/ssl/private/bind.key",
	})
	certRes := ParseResponse(certBody)
	certID := ExtractID(certRes.Data)

	var bindingID int64

	t.Run("Create binding - host type", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/ssl-bindings", map[string]interface{}{
			"ssl_certificate_id": certID,
			"binding_type":       "host",
			"host_id":            hostID,
			"is_default":         true,
			"priority":           1,
		})
		AssertStatus(t, resp, http.StatusCreated)
		res := AssertSuccess(t, body)
		bindingID = ExtractID(res.Data)
		if bindingID == 0 {
			t.Fatal("Expected valid binding ID")
		}
	})

	t.Run("Create binding - virtual_host type", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/ssl-bindings", map[string]interface{}{
			"ssl_certificate_id": certID,
			"binding_type":       "virtual_host",
			"virtual_host_id":    vhostID,
			"priority":           2,
		})
		AssertStatus(t, resp, http.StatusCreated)
		AssertSuccess(t, body)
	})

	t.Run("Create validation - missing cert id", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/ssl-bindings", map[string]interface{}{
			"binding_type": "host",
			"host_id":      hostID,
		})
		AssertStatus(t, resp, http.StatusBadRequest)
		AssertError(t, body)
	})

	t.Run("Create validation - invalid binding type", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/ssl-bindings", map[string]interface{}{
			"ssl_certificate_id": certID,
			"binding_type":       "invalid",
		})
		AssertStatus(t, resp, http.StatusBadRequest)
		AssertError(t, body)
	})

	t.Run("Get all bindings", func(t *testing.T) {
		resp, body := ts.DoRequest("GET", "/api/admin/ssl-bindings", nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	t.Run("Get binding by ID", func(t *testing.T) {
		resp, body := ts.DoRequest("GET", fmt.Sprintf("/api/admin/ssl-bindings/%d", bindingID), nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	t.Run("Update binding", func(t *testing.T) {
		resp, body := ts.DoRequest("PUT", fmt.Sprintf("/api/admin/ssl-bindings/%d", bindingID), map[string]interface{}{
			"binding_type": "global",
			"is_default":   false,
			"priority":     10,
		})
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	t.Run("Delete binding", func(t *testing.T) {
		resp, body := ts.DoRequest("DELETE", fmt.Sprintf("/api/admin/ssl-bindings/%d", bindingID), nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})
}

// =========================================================
// TEST TLS OPTIONS
// =========================================================

func TestTLSOptionsCRUD(t *testing.T) {
	ts := SetupTestServer(t)
	defer ts.Close()

	hostID := ts.CreateTestHost(t)

	var tlsID int64

	t.Run("Create TLS option - host", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/tls-options", map[string]interface{}{
			"binding_type":    "host",
			"host_id":         hostID,
			"min_tls_version": "1.3",
			"http2_enabled":   true,
			"hsts_enabled":    true,
			"hsts_max_age":    63072000,
		})
		AssertStatus(t, resp, http.StatusCreated)
		res := AssertSuccess(t, body)
		tlsID = ExtractID(res.Data)
		if tlsID == 0 {
			t.Fatal("Expected valid TLS option ID")
		}
	})

	t.Run("Create TLS option - global with defaults", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/tls-options", map[string]interface{}{
			"binding_type": "global",
		})
		AssertStatus(t, resp, http.StatusCreated)
		AssertSuccess(t, body)
	})

	t.Run("Create validation - invalid binding type", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/tls-options", map[string]interface{}{
			"binding_type": "invalid",
		})
		AssertStatus(t, resp, http.StatusBadRequest)
		AssertError(t, body)
	})

	t.Run("Create validation - missing binding type", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/tls-options", map[string]interface{}{
			"http2_enabled": true,
		})
		AssertStatus(t, resp, http.StatusBadRequest)
		AssertError(t, body)
	})

	t.Run("Get all TLS options", func(t *testing.T) {
		resp, body := ts.DoRequest("GET", "/api/admin/tls-options", nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	t.Run("Get TLS option by ID", func(t *testing.T) {
		resp, body := ts.DoRequest("GET", fmt.Sprintf("/api/admin/tls-options/%d", tlsID), nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	t.Run("Update TLS option", func(t *testing.T) {
		resp, body := ts.DoRequest("PUT", fmt.Sprintf("/api/admin/tls-options/%d", tlsID), map[string]interface{}{
			"binding_type":    "host",
			"host_id":         hostID,
			"min_tls_version": "1.2",
			"http2_enabled":   false,
			"hsts_enabled":    false,
			"hsts_max_age":    0,
		})
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	t.Run("Delete TLS option", func(t *testing.T) {
		resp, body := ts.DoRequest("DELETE", fmt.Sprintf("/api/admin/tls-options/%d", tlsID), nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})
}

// =========================================================
// TEST SERVICE DISCOVERY
// =========================================================

func TestServiceDiscoveryCRUD(t *testing.T) {
	ts := SetupTestServer(t)
	defer ts.Close()

	hostID := ts.CreateTestHost(t)
	vhostID := ts.CreateTestVirtualHost(t, hostID)

	var sdID int64

	t.Run("Create service discovery", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/service-discovery", map[string]interface{}{
			"virtual_host_id":          vhostID,
			"provider":                 "consul",
			"endpoint_url":             "http://consul:8500/v1/catalog/service/web",
			"refresh_interval_seconds": 60,
			"is_active":                true,
		})
		AssertStatus(t, resp, http.StatusCreated)
		res := AssertSuccess(t, body)
		sdID = ExtractID(res.Data)
		if sdID == 0 {
			t.Fatal("Expected valid service discovery ID")
		}
	})

	t.Run("Create with defaults", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/service-discovery", map[string]interface{}{
			"virtual_host_id": vhostID,
			"provider":        "etcd",
			"endpoint_url":    "http://etcd:2379/services",
		})
		AssertStatus(t, resp, http.StatusCreated)
		AssertSuccess(t, body)
	})

	t.Run("Create validation - missing vhost id", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/service-discovery", map[string]interface{}{
			"provider":     "consul",
			"endpoint_url": "http://consul:8500",
		})
		AssertStatus(t, resp, http.StatusBadRequest)
		AssertError(t, body)
	})

	t.Run("Create validation - missing provider", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/service-discovery", map[string]interface{}{
			"virtual_host_id": vhostID,
			"endpoint_url":    "http://consul:8500",
		})
		AssertStatus(t, resp, http.StatusBadRequest)
		AssertError(t, body)
	})

	t.Run("Create validation - missing endpoint", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/service-discovery", map[string]interface{}{
			"virtual_host_id": vhostID,
			"provider":        "consul",
		})
		AssertStatus(t, resp, http.StatusBadRequest)
		AssertError(t, body)
	})

	t.Run("Get all service discoveries", func(t *testing.T) {
		resp, body := ts.DoRequest("GET", "/api/admin/service-discovery", nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	t.Run("Get service discovery by ID", func(t *testing.T) {
		resp, body := ts.DoRequest("GET", fmt.Sprintf("/api/admin/service-discovery/%d", sdID), nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	t.Run("Update service discovery", func(t *testing.T) {
		resp, body := ts.DoRequest("PUT", fmt.Sprintf("/api/admin/service-discovery/%d", sdID), map[string]interface{}{
			"provider":                 "kubernetes",
			"endpoint_url":             "https://k8s-api:6443/api/v1/endpoints",
			"refresh_interval_seconds": 15,
			"is_active":                false,
		})
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	t.Run("Delete service discovery", func(t *testing.T) {
		resp, body := ts.DoRequest("DELETE", fmt.Sprintf("/api/admin/service-discovery/%d", sdID), nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	t.Run("Delete not found", func(t *testing.T) {
		resp, body := ts.DoRequest("DELETE", "/api/admin/service-discovery/9999", nil)
		AssertStatus(t, resp, http.StatusBadRequest)
		AssertError(t, body)
	})
}

// =========================================================
// TEST CONFIG VERSIONS
// =========================================================

func TestConfigVersionsCRUD(t *testing.T) {
	ts := SetupTestServer(t)
	defer ts.Close()

	var configID int64

	t.Run("Create config version", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/config-versions", map[string]interface{}{
			"config_name":    "gateway_routing",
			"version_number": 1,
			"changed_by":     "admin",
			"notes":          "Initial routing configuration",
		})
		AssertStatus(t, resp, http.StatusCreated)
		res := AssertSuccess(t, body)
		configID = ExtractID(res.Data)
		if configID == 0 {
			t.Fatal("Expected valid config version ID")
		}
	})

	t.Run("Create another version", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/config-versions", map[string]interface{}{
			"config_name":    "gateway_routing",
			"version_number": 2,
			"changed_by":     "devops",
			"notes":          "Added rate limiting rules",
		})
		AssertStatus(t, resp, http.StatusCreated)
		AssertSuccess(t, body)
	})

	t.Run("Create validation - missing config name", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/config-versions", map[string]interface{}{
			"version_number": 1,
		})
		AssertStatus(t, resp, http.StatusBadRequest)
		AssertError(t, body)
	})

	t.Run("Create validation - missing version number", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/config-versions", map[string]interface{}{
			"config_name": "test",
		})
		AssertStatus(t, resp, http.StatusBadRequest)
		AssertError(t, body)
	})

	t.Run("Get all config versions", func(t *testing.T) {
		resp, body := ts.DoRequest("GET", "/api/admin/config-versions", nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	t.Run("Get config version by ID", func(t *testing.T) {
		resp, body := ts.DoRequest("GET", fmt.Sprintf("/api/admin/config-versions/%d", configID), nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	t.Run("Update config version", func(t *testing.T) {
		resp, body := ts.DoRequest("PUT", fmt.Sprintf("/api/admin/config-versions/%d", configID), map[string]interface{}{
			"config_name":    "gateway_routing",
			"version_number": 1,
			"changed_by":     "admin",
			"notes":          "Updated: Initial routing with CORS",
		})
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	t.Run("Delete config version", func(t *testing.T) {
		resp, body := ts.DoRequest("DELETE", fmt.Sprintf("/api/admin/config-versions/%d", configID), nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	t.Run("Delete not found", func(t *testing.T) {
		resp, body := ts.DoRequest("DELETE", "/api/admin/config-versions/9999", nil)
		AssertStatus(t, resp, http.StatusBadRequest)
		AssertError(t, body)
	})
}

// =========================================================
// TEST MAINTENANCE WINDOWS
// =========================================================

func TestMaintenanceWindowsCRUD(t *testing.T) {
	ts := SetupTestServer(t)
	defer ts.Close()

	hostID := ts.CreateTestHost(t)
	vhostID := ts.CreateTestVirtualHost(t, hostID)

	var mwID int64

	t.Run("Create maintenance window", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/maintenance-windows", map[string]interface{}{
			"virtual_host_id":           vhostID,
			"title":                     "Scheduled Maintenance",
			"start_at":                  "2026-06-01 00:00:00",
			"end_at":                    "2026-06-01 04:00:00",
			"maintenance_response_code": 503,
			"maintenance_message":       "System is under maintenance",
			"is_active":                 true,
		})
		AssertStatus(t, resp, http.StatusCreated)
		res := AssertSuccess(t, body)
		mwID = ExtractID(res.Data)
		if mwID == 0 {
			t.Fatal("Expected valid maintenance window ID")
		}
	})

	t.Run("Create global maintenance (no vhost)", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/maintenance-windows", map[string]interface{}{
			"title":    "Global Emergency Maintenance",
			"start_at": "2026-07-01 02:00:00",
			"end_at":   "2026-07-01 06:00:00",
		})
		AssertStatus(t, resp, http.StatusCreated)
		AssertSuccess(t, body)
	})

	t.Run("Create validation - missing title", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/maintenance-windows", map[string]interface{}{
			"start_at": "2026-06-01 00:00:00",
			"end_at":   "2026-06-01 04:00:00",
		})
		AssertStatus(t, resp, http.StatusBadRequest)
		AssertError(t, body)
	})

	t.Run("Create validation - missing start_at", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/maintenance-windows", map[string]interface{}{
			"title":  "Test",
			"end_at": "2026-06-01 04:00:00",
		})
		AssertStatus(t, resp, http.StatusBadRequest)
		AssertError(t, body)
	})

	t.Run("Create validation - missing end_at", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/maintenance-windows", map[string]interface{}{
			"title":    "Test",
			"start_at": "2026-06-01 00:00:00",
		})
		AssertStatus(t, resp, http.StatusBadRequest)
		AssertError(t, body)
	})

	t.Run("Get all maintenance windows", func(t *testing.T) {
		resp, body := ts.DoRequest("GET", "/api/admin/maintenance-windows", nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	t.Run("Get maintenance window by ID", func(t *testing.T) {
		resp, body := ts.DoRequest("GET", fmt.Sprintf("/api/admin/maintenance-windows/%d", mwID), nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	t.Run("Update maintenance window", func(t *testing.T) {
		resp, body := ts.DoRequest("PUT", fmt.Sprintf("/api/admin/maintenance-windows/%d", mwID), map[string]interface{}{
			"title":                     "Extended Maintenance",
			"start_at":                  "2026-06-01 00:00:00",
			"end_at":                    "2026-06-01 08:00:00",
			"maintenance_response_code": 200,
			"maintenance_message":       "We will be back soon!",
			"is_active":                 false,
		})
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	t.Run("Delete maintenance window", func(t *testing.T) {
		resp, body := ts.DoRequest("DELETE", fmt.Sprintf("/api/admin/maintenance-windows/%d", mwID), nil)
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	t.Run("Delete not found", func(t *testing.T) {
		resp, body := ts.DoRequest("DELETE", "/api/admin/maintenance-windows/9999", nil)
		AssertStatus(t, resp, http.StatusBadRequest)
		AssertError(t, body)
	})
}
