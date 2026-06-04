package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

// Logger for detailed output
type Logger struct {
	logFile *os.File
}

func NewLogger(filename string) (*Logger, error) {
	f, err := os.Create(filename)
	if err != nil {
		return nil, err
	}
	return &Logger{logFile: f}, nil
}

func (l *Logger) Log(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	fmt.Print(msg)
	if l.logFile != nil {
		l.logFile.WriteString(msg + "\n")
	}
}

func (l *Logger) Close() {
	if l.logFile != nil {
		l.logFile.Close()
	}
}

type TestResult struct {
	Endpoint     string
	Method       string
	Status       int
	Success      bool
	Duration     time.Duration
	RequestBody  string
	ResponseBody string
	RequestHeaders map[string]string
	Error        string
}

type APITestSuite struct {
	BaseURL      string
	AccessToken  string
	RefreshToken string
	Results      []TestResult
	Logger       *Logger
	
	// Track created resource IDs for proper CRUD flow
	CreatedConsumerID     int
	CreatedHostID         int
	CreatedVHostID        int
	CreatedVDirID         int
	CreatedUpstreamID     int
	CreatedRateLimitID    int
	CreatedCORSID         int
	CreatedCircuitBreakerID int
	CreatedExternalAuthID int
	CreatedJWTConfigID    int
	CreatedCredentialID   int
	CreatedAPIKeyID       int
	CreatedRouteAccessID  int
	CreatedIPWhitelistID  int
	CreatedIPBlacklistID  int
	CreatedRequestHeaderID int
	CreatedResponseHeaderID int
	CreatedQueryRewriteID int
	CreatedACMEAccountID  int
	CreatedSSLCertID      int
	CreatedCertDomainID   int
	CreatedSSLBindingID   int
	CreatedTLSOptionID    int
	CreatedServiceDiscoveryID int
	CreatedConfigVersionID int
	CreatedMaintenanceWindowID int
}

func NewAPITestSuite(baseURL string) *APITestSuite {
	return &APITestSuite{
		BaseURL: baseURL,
		Results: make([]TestResult, 0),
	}
}

func (suite *APITestSuite) makeRequest(method, endpoint string, body interface{}, headers map[string]string) TestResult {
	result := TestResult{
		Endpoint:       endpoint,
		Method:         method,
		RequestHeaders: make(map[string]string),
	}

	start := time.Now()

	// Log request
	suite.Logger.Log("\n%s %s", method, endpoint)
	if body != nil {
		jsonBytes, _ := json.MarshalIndent(body, "  ", "  ")
		suite.Logger.Log("Request Body: %s", string(jsonBytes))
	}

	// Prepare request body
	var reqBody io.Reader
	if body != nil {
		jsonBytes, _ := json.Marshal(body)
		reqBody = bytes.NewBuffer(jsonBytes)
		result.RequestBody = string(jsonBytes)
	}

	// Create request
	req, err := http.NewRequest(method, suite.BaseURL+endpoint, reqBody)
	if err != nil {
		result.Error = err.Error()
		result.Duration = time.Since(start)
		suite.Results = append(suite.Results, result)
		suite.Logger.Log("❌ Error: %s", err.Error())
		return result
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	for k, v := range headers {
		req.Header.Set(k, v)
		result.RequestHeaders[k] = v
	}

	// Add auth token if available
	if suite.AccessToken != "" {
		req.Header.Set("Authorization", "Bearer "+suite.AccessToken)
		result.RequestHeaders["Authorization"] = "Bearer ***"
	}

	// Execute request
	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		result.Error = err.Error()
		result.Duration = time.Since(start)
		suite.Results = append(suite.Results, result)
		suite.Logger.Log("❌ HTTP Error: %s", err.Error())
		return result
	}
	defer resp.Body.Close()

	// Read response
	respBody, _ := io.ReadAll(resp.Body)
	result.Status = resp.StatusCode
	result.ResponseBody = string(respBody)
	result.Duration = time.Since(start)

	// Log response
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		suite.Logger.Log("✅ Status: %d (%.2fms)", resp.StatusCode, float64(result.Duration.Microseconds())/1000.0)
	} else {
		suite.Logger.Log("❌ Status: %d (%.2fms)", resp.StatusCode, float64(result.Duration.Microseconds())/1000.0)
		suite.Logger.Log("Response: %s", string(respBody))
	}

	// Check success
	result.Success = resp.StatusCode >= 200 && resp.StatusCode < 300

	// Extract tokens from login/refresh
	if strings.Contains(endpoint, "/auth/login") || strings.Contains(endpoint, "/auth/refresh") {
		if result.Success {
			var respData map[string]interface{}
			if err := json.Unmarshal(respBody, &respData); err == nil {
				if data, ok := respData["data"].(map[string]interface{}); ok {
					if token, ok := data["access_token"].(string); ok {
						suite.AccessToken = token
						suite.Logger.Log("🔑 Access Token extracted")
					}
					if token, ok := data["refresh_token"].(string); ok {
						suite.RefreshToken = token
						suite.Logger.Log("🔑 Refresh Token extracted")
					}
				}
			}
		}
	}
	
	// Extract created resource ID from POST responses
	if method == "POST" && result.Success {
		var respData map[string]interface{}
		if err := json.Unmarshal(respBody, &respData); err == nil {
			if data, ok := respData["data"].(map[string]interface{}); ok {
				if id, ok := data["id"].(float64); ok {
					// Auto-track IDs based on endpoint
					if strings.Contains(endpoint, "/consumers") && !strings.Contains(endpoint, "credentials") && !strings.Contains(endpoint, "api-keys") {
						suite.CreatedConsumerID = int(id)
						suite.Logger.Log("🆔 Created Consumer ID: %d", suite.CreatedConsumerID)
					} else if strings.Contains(endpoint, "/hosts") && !strings.Contains(endpoint, "virtual") && !strings.Contains(endpoint, "certificate") {
						suite.CreatedHostID = int(id)
						suite.Logger.Log("🆔 Created Host ID: %d", suite.CreatedHostID)
					} else if strings.Contains(endpoint, "virtual-hosts") {
						suite.CreatedVHostID = int(id)
						suite.Logger.Log("🆔 Created VHost ID: %d", suite.CreatedVHostID)
					} else if strings.Contains(endpoint, "virtual-directories") {
						suite.CreatedVDirID = int(id)
						suite.Logger.Log("🆔 Created VDir ID: %d", suite.CreatedVDirID)
					} else if strings.Contains(endpoint, "upstream-servers") {
						suite.CreatedUpstreamID = int(id)
						suite.Logger.Log("🆔 Created Upstream ID: %d", suite.CreatedUpstreamID)
					} else if strings.Contains(endpoint, "/rate-limits") {
						suite.CreatedRateLimitID = int(id)
						suite.Logger.Log("🆔 Created RateLimit ID: %d", suite.CreatedRateLimitID)
					} else if strings.Contains(endpoint, "/cors-configs") {
						suite.CreatedCORSID = int(id)
						suite.Logger.Log("🆔 Created CORS ID: %d", suite.CreatedCORSID)
					} else if strings.Contains(endpoint, "/circuit-breakers") {
						suite.CreatedCircuitBreakerID = int(id)
						suite.Logger.Log("🆔 Created CircuitBreaker ID: %d", suite.CreatedCircuitBreakerID)
					} else if strings.Contains(endpoint, "/external-auth") {
						suite.CreatedExternalAuthID = int(id)
						suite.Logger.Log("🆔 Created ExternalAuth ID: %d", suite.CreatedExternalAuthID)
					} else if strings.Contains(endpoint, "/jwt-configs") {
						suite.CreatedJWTConfigID = int(id)
						suite.Logger.Log("🆔 Created JWTConfig ID: %d", suite.CreatedJWTConfigID)
					} else if strings.Contains(endpoint, "/consumer-credentials") {
						suite.CreatedCredentialID = int(id)
						suite.Logger.Log("🆔 Created Credential ID: %d", suite.CreatedCredentialID)
					} else if strings.Contains(endpoint, "/api-keys") {
						suite.CreatedAPIKeyID = int(id)
						suite.Logger.Log("🆔 Created APIKey ID: %d", suite.CreatedAPIKeyID)
					} else if strings.Contains(endpoint, "/route-access") {
						suite.CreatedRouteAccessID = int(id)
						suite.Logger.Log("🆔 Created RouteAccess ID: %d", suite.CreatedRouteAccessID)
					} else if strings.Contains(endpoint, "/ip-whitelists") {
						suite.CreatedIPWhitelistID = int(id)
						suite.Logger.Log("🆔 Created IPWhitelist ID: %d", suite.CreatedIPWhitelistID)
					} else if strings.Contains(endpoint, "/ip-blacklists") {
						suite.CreatedIPBlacklistID = int(id)
						suite.Logger.Log("🆔 Created IPBlacklist ID: %d", suite.CreatedIPBlacklistID)
					} else if strings.Contains(endpoint, "/request-headers") {
						suite.CreatedRequestHeaderID = int(id)
						suite.Logger.Log("🆔 Created RequestHeader ID: %d", suite.CreatedRequestHeaderID)
					} else if strings.Contains(endpoint, "/response-headers") {
						suite.CreatedResponseHeaderID = int(id)
						suite.Logger.Log("🆔 Created ResponseHeader ID: %d", suite.CreatedResponseHeaderID)
					} else if strings.Contains(endpoint, "/query-rewrites") {
						suite.CreatedQueryRewriteID = int(id)
						suite.Logger.Log("🆔 Created QueryRewrite ID: %d", suite.CreatedQueryRewriteID)
					} else if strings.Contains(endpoint, "/acme-accounts") {
						suite.CreatedACMEAccountID = int(id)
						suite.Logger.Log("🆔 Created ACMEAccount ID: %d", suite.CreatedACMEAccountID)
					} else if strings.Contains(endpoint, "/ssl-certificates") {
						suite.CreatedSSLCertID = int(id)
						suite.Logger.Log("🆔 Created SSLCert ID: %d", suite.CreatedSSLCertID)
					} else if strings.Contains(endpoint, "/certificate-domains") {
						suite.CreatedCertDomainID = int(id)
						suite.Logger.Log("🆔 Created CertDomain ID: %d", suite.CreatedCertDomainID)
					} else if strings.Contains(endpoint, "/ssl-bindings") {
						suite.CreatedSSLBindingID = int(id)
						suite.Logger.Log("🆔 Created SSLBinding ID: %d", suite.CreatedSSLBindingID)
					} else if strings.Contains(endpoint, "/tls-options") {
						suite.CreatedTLSOptionID = int(id)
						suite.Logger.Log("🆔 Created TLSOption ID: %d", suite.CreatedTLSOptionID)
					} else if strings.Contains(endpoint, "/service-discovery") {
						suite.CreatedServiceDiscoveryID = int(id)
						suite.Logger.Log("🆔 Created ServiceDiscovery ID: %d", suite.CreatedServiceDiscoveryID)
					} else if strings.Contains(endpoint, "/config-versions") {
						suite.CreatedConfigVersionID = int(id)
						suite.Logger.Log("🆔 Created ConfigVersion ID: %d", suite.CreatedConfigVersionID)
					} else if strings.Contains(endpoint, "/maintenance-windows") {
						suite.CreatedMaintenanceWindowID = int(id)
						suite.Logger.Log("🆔 Created MaintenanceWindow ID: %d", suite.CreatedMaintenanceWindowID)
					}
				}
			}
		}
	}

	suite.Results = append(suite.Results, result)
	return result
}

func (suite *APITestSuite) RunAllTests() {
	suite.Logger.Log("🚀 Starting Swantara Gate API Test Suite (Full Postman Collection)")
	suite.Logger.Log(strings.Repeat("=", 60))
	suite.Logger.Log("Base URL: %s\n", suite.BaseURL)
	
	// Generate unique suffix for test data
	testSuffix := fmt.Sprintf("%d", time.Now().Unix())
	suite.Logger.Log("🏷️  Test suffix: %s\n", testSuffix)

	// 1. Health Check
	suite.Logger.Log("📋 1. Health Check")
	suite.makeRequest("GET", "/api/health", nil, nil)

	// 2. Auth Tests
	suite.Logger.Log("\n🔐 2. Authentication")
	
	// Login with existing user
	suite.makeRequest("POST", "/api/admin/auth/login", map[string]interface{}{
		"username": "testcli",
		"password": "admin1324",
	}, nil)

	// Get Current User (Me)
	suite.makeRequest("GET", "/api/admin/auth/me", nil, nil)

	// Refresh Token
	suite.makeRequest("POST", "/api/admin/auth/refresh", map[string]interface{}{
		"refresh_token": suite.RefreshToken,
	}, nil)

	// Note: Logout is SKIPPED to prevent token blacklisting during tests
	
	suite.Logger.Log("\n✅ Authentication tests completed, continuing with other endpoints...")

	// 3. Configuration
	suite.Logger.Log("\n⚙️  3. Configuration")
	suite.makeRequest("POST", "/api/admin/config/reload", nil, nil)

	// 4. Users CRUD
	suite.Logger.Log("\n👥 4. Users")
	suite.makeRequest("GET", "/api/admin/users", nil, nil)
	suite.makeRequest("POST", "/api/admin/users", map[string]interface{}{
		"username":  "apitest_" + testSuffix,
		"password":  "test123",
		"full_name": "API Test User",
		"email":     "apitest" + testSuffix + "@example.com",
		"role":      "admin",
		"is_active": true,
	}, nil)
	suite.makeRequest("GET", "/api/admin/users/1", nil, nil)
	suite.makeRequest("PUT", "/api/admin/users/1", map[string]interface{}{
		"full_name": "Updated Name",
		"email":     "updated@example.com",
		"role":      "admin",
		"is_active": true,
	}, nil)

	// 5. Consumers CRUD
	suite.Logger.Log("\n🏢 5. API Consumers")
	suite.makeRequest("POST", "/api/admin/consumers", map[string]interface{}{
		"consumer_name": "apitest-" + testSuffix,
		"description":   "API Test Application",
		"contact_email": "apitest" + testSuffix + "@app.com",
		"is_active":     true,
	}, nil)
	suite.makeRequest("GET", "/api/admin/consumers", nil, nil)
	if suite.CreatedConsumerID > 0 {
		suite.makeRequest("GET", "/api/admin/consumers/"+fmt.Sprintf("%d", suite.CreatedConsumerID), nil, nil)
		suite.makeRequest("PUT", "/api/admin/consumers/"+fmt.Sprintf("%d", suite.CreatedConsumerID), map[string]interface{}{
			"description":   "Updated App",
			"contact_email": "updated@app.com",
			"is_active":     true,
		}, nil)
		suite.makeRequest("DELETE", "/api/admin/consumers/"+fmt.Sprintf("%d", suite.CreatedConsumerID), nil, nil)
	}

	// 6. Hosts CRUD
	suite.Logger.Log("\n🌐 6. Hosts")
	suite.makeRequest("POST", "/api/admin/hosts", map[string]interface{}{
		"host_name":   "api" + testSuffix + ".apitest.com",
		"description": "API Test Host",
		"is_active":   true,
	}, nil)
	suite.makeRequest("GET", "/api/admin/hosts", nil, nil)
	if suite.CreatedHostID > 0 {
		suite.makeRequest("GET", "/api/admin/hosts/"+fmt.Sprintf("%d", suite.CreatedHostID), nil, nil)
		suite.makeRequest("PUT", "/api/admin/hosts/"+fmt.Sprintf("%d", suite.CreatedHostID), map[string]interface{}{
			"description": "Updated Host",
			"is_active":   true,
		}, nil)
		suite.makeRequest("DELETE", "/api/admin/hosts/"+fmt.Sprintf("%d", suite.CreatedHostID), nil, nil)
	}

	// 7. Virtual Hosts CRUD
	suite.Logger.Log("\n🔀 7. Virtual Hosts")
	suite.makeRequest("POST", "/api/admin/virtual-hosts", map[string]interface{}{
		"host_id":         1,
		"vhost_name":      "vhost" + testSuffix + ".apitest.com",
		"lb_algorithm":    "round_robin",
		"sticky_session":  false,
		"failover_mode":   "active-active",
		"is_active":       true,
	}, nil)
	suite.makeRequest("GET", "/api/admin/virtual-hosts", nil, nil)
	if suite.CreatedVHostID > 0 {
		suite.makeRequest("GET", "/api/admin/virtual-hosts/"+fmt.Sprintf("%d", suite.CreatedVHostID), nil, nil)
		suite.makeRequest("PUT", "/api/admin/virtual-hosts/"+fmt.Sprintf("%d", suite.CreatedVHostID), map[string]interface{}{
			"lb_algorithm":  "weighted_round_robin",
			"sticky_session": true,
			"failover_mode": "active-passive",
			"is_active":     true,
		}, nil)
	} else {
		suite.Logger.Log("\n⚠️  VHost creation failed, using existing VHost ID 1")
		suite.CreatedVHostID = 1
	}

	// 8. Upstream Servers CRUD
	suite.Logger.Log("\n🖥️  8. Upstream Servers")
	suite.makeRequest("POST", "/api/admin/upstream-servers", map[string]interface{}{
		"virtual_host_id":            suite.CreatedVHostID,
		"target_host":                "192.168.1.10",
		"target_port":                8080,
		"protocol":                   "http",
		"priority":                   1,
		"weight":                     1,
		"is_backup":                  false,
		"is_active":                  true,
		"health_check_enabled":       true,
		"health_check_path":          "/health",
		"health_check_interval_seconds": 10,
		"health_check_timeout_seconds": 3,
		"max_fails":                  3,
		"fail_timeout_seconds":       30,
	}, nil)
	suite.makeRequest("GET", "/api/admin/upstream-servers", nil, nil)
	if suite.CreatedUpstreamID > 0 {
		suite.makeRequest("GET", "/api/admin/upstream-servers/"+fmt.Sprintf("%d", suite.CreatedUpstreamID), nil, nil)
		suite.makeRequest("GET", "/api/admin/virtual-hosts/"+fmt.Sprintf("%d", suite.CreatedVHostID)+"/upstream-servers", nil, nil)
		suite.makeRequest("PUT", "/api/admin/upstream-servers/"+fmt.Sprintf("%d", suite.CreatedUpstreamID), map[string]interface{}{
			"target_host":          "192.168.1.11",
			"target_port":          8081,
			"weight":               2,
			"health_check_interval_seconds": 15,
			"is_active":            true,
		}, nil)
		suite.makeRequest("DELETE", "/api/admin/upstream-servers/"+fmt.Sprintf("%d", suite.CreatedUpstreamID), nil, nil)
	}

	// 9. Virtual Directories (Routes) CRUD
	suite.Logger.Log("\n🛣️  9. Virtual Directories (Routes)")
	suite.makeRequest("POST", "/api/admin/virtual-directories", map[string]interface{}{
		"virtual_host_id":     suite.CreatedVHostID,
		"source_path":         "/api/v1",
		"target_path":         "/",
		"match_type":          "prefix",
		"strip_prefix":        true,
		"preserve_host_header": false,
		"auth_type":           "none",
		"is_active":           true,
		"proxy_timeout_seconds": 30,
		"retry_count":         2,
		"retry_delay_ms":      100,
		"max_request_size_mb": 10,
		"websocket_enabled":   false,
		"cache_enabled":       false,
		"cache_ttl_seconds":   0,
	}, nil)
	suite.makeRequest("GET", "/api/admin/virtual-directories", nil, nil)
	if suite.CreatedVDirID > 0 {
		suite.makeRequest("GET", "/api/admin/virtual-directories/"+fmt.Sprintf("%d", suite.CreatedVDirID), nil, nil)
		suite.makeRequest("GET", "/api/admin/virtual-hosts/"+fmt.Sprintf("%d", suite.CreatedVHostID)+"/virtual-directories", nil, nil)
		suite.makeRequest("GET", "/api/admin/hosts/1/virtual-directories", nil, nil)
		suite.makeRequest("PUT", "/api/admin/virtual-directories/"+fmt.Sprintf("%d", suite.CreatedVDirID), map[string]interface{}{
			"source_path":         "/api/v2",
			"target_path":         "/v2",
			"match_type":          "prefix",
			"strip_prefix":        true,
			"auth_type":           "api_key",
			"is_active":           true,
			"proxy_timeout_seconds": 60,
		}, nil)
		// Methods
		suite.makeRequest("GET", "/api/admin/virtual-directories/"+fmt.Sprintf("%d", suite.CreatedVDirID)+"/methods", nil, nil)
		suite.makeRequest("PUT", "/api/admin/virtual-directories/"+fmt.Sprintf("%d", suite.CreatedVDirID)+"/methods", map[string]interface{}{
			"methods": []string{"GET", "POST", "PUT", "DELETE"},
		}, nil)
	}

	// 10. Consumer Credentials CRUD
	suite.Logger.Log("\n🔑 10. Consumer Credentials")
	if suite.CreatedConsumerID > 0 {
		suite.makeRequest("POST", "/api/admin/consumer-credentials", map[string]interface{}{
			"consumer_id": suite.CreatedConsumerID,
			"auth_type":   "basic",
			"username":    "api-user-" + testSuffix,
			"password":    "secret123",
			"is_active":   true,
		}, nil)
	} else {
		suite.makeRequest("POST", "/api/admin/consumer-credentials", map[string]interface{}{
			"consumer_id": 1,
			"auth_type":   "basic",
			"username":    "api-user-" + testSuffix,
			"password":    "secret123",
			"is_active":   true,
		}, nil)
	}
	suite.makeRequest("GET", "/api/admin/consumer-credentials", nil, nil)
	if suite.CreatedCredentialID > 0 {
		suite.makeRequest("GET", "/api/admin/consumer-credentials/"+fmt.Sprintf("%d", suite.CreatedCredentialID), nil, nil)
		if suite.CreatedConsumerID > 0 {
			suite.makeRequest("GET", "/api/admin/consumers/"+fmt.Sprintf("%d", suite.CreatedConsumerID)+"/credentials", nil, nil)
		}
		suite.makeRequest("PUT", "/api/admin/consumer-credentials/"+fmt.Sprintf("%d", suite.CreatedCredentialID), map[string]interface{}{
			"username":  "new-user",
			"password":  "new-secret",
			"is_active": true,
		}, nil)
		suite.makeRequest("DELETE", "/api/admin/consumer-credentials/"+fmt.Sprintf("%d", suite.CreatedCredentialID), nil, nil)
	}

	// 11. API Keys CRUD
	suite.Logger.Log("\n🗝️  11. API Keys")
	if suite.CreatedConsumerID > 0 {
		suite.makeRequest("POST", "/api/admin/api-keys", map[string]interface{}{
			"consumer_id":         suite.CreatedConsumerID,
			"description":         "Production API Key",
			"expired_at":          "2027-12-31 23:59:59",
			"rate_limit_override": 1000,
			"is_active":           true,
		}, nil)
	} else {
		suite.makeRequest("POST", "/api/admin/api-keys", map[string]interface{}{
			"consumer_id":         1,
			"description":         "Production API Key",
			"expired_at":          "2027-12-31 23:59:59",
			"rate_limit_override": 1000,
			"is_active":           true,
		}, nil)
	}
	suite.makeRequest("GET", "/api/admin/api-keys", nil, nil)
	if suite.CreatedAPIKeyID > 0 {
		suite.makeRequest("GET", "/api/admin/api-keys/"+fmt.Sprintf("%d", suite.CreatedAPIKeyID), nil, nil)
		if suite.CreatedConsumerID > 0 {
			suite.makeRequest("GET", "/api/admin/consumers/"+fmt.Sprintf("%d", suite.CreatedConsumerID)+"/api-keys", nil, nil)
		}
		suite.makeRequest("PUT", "/api/admin/api-keys/"+fmt.Sprintf("%d", suite.CreatedAPIKeyID), map[string]interface{}{
			"description": "Updated Key",
			"is_active":   true,
		}, nil)
		suite.makeRequest("DELETE", "/api/admin/api-keys/"+fmt.Sprintf("%d", suite.CreatedAPIKeyID), nil, nil)
	}

	// 12. Route Consumer Access (ACL) CRUD
	suite.Logger.Log("\n🚦 12. Route Consumer Access (ACL)")
	if suite.CreatedVDirID > 0 && suite.CreatedConsumerID > 0 {
		suite.makeRequest("POST", "/api/admin/route-access", map[string]interface{}{
			"virtual_directory_id": suite.CreatedVDirID,
			"consumer_id":          suite.CreatedConsumerID,
			"is_active":            true,
		}, nil)
	} else {
		suite.makeRequest("POST", "/api/admin/route-access", map[string]interface{}{
			"virtual_directory_id": 1,
			"consumer_id":          1,
			"is_active":            true,
		}, nil)
	}
	suite.makeRequest("GET", "/api/admin/route-access", nil, nil)
	if suite.CreatedRouteAccessID > 0 {
		suite.makeRequest("GET", "/api/admin/route-access/"+fmt.Sprintf("%d", suite.CreatedRouteAccessID), nil, nil)
		if suite.CreatedVDirID > 0 {
			suite.makeRequest("GET", "/api/admin/virtual-directories/"+fmt.Sprintf("%d", suite.CreatedVDirID)+"/access", nil, nil)
		}
		suite.makeRequest("PUT", "/api/admin/route-access/"+fmt.Sprintf("%d", suite.CreatedRouteAccessID), map[string]interface{}{
			"is_active": false,
		}, nil)
		suite.makeRequest("DELETE", "/api/admin/route-access/"+fmt.Sprintf("%d", suite.CreatedRouteAccessID), nil, nil)
	}

	// 13. JWT Configs CRUD
	suite.Logger.Log("\n🔐 13. JWT Configs")
	suite.makeRequest("POST", "/api/admin/jwt-configs", map[string]interface{}{
		"virtual_directory_id": suite.CreatedVDirID,
		"algorithm":            "HS256",
		"jwt_secret":           "my-secret-key-at-least-32-chars-long",
		"issuer":               "my-app",
		"audience":             "api-gateway",
		"expired_in_seconds":   3600,
		"clock_skew_seconds":   30,
		"require_exp":          true,
		"require_iat":          true,
		"is_active":            true,
	}, nil)
	suite.makeRequest("GET", "/api/admin/jwt-configs", nil, nil)
	if suite.CreatedJWTConfigID > 0 {
		suite.makeRequest("GET", "/api/admin/jwt-configs/"+fmt.Sprintf("%d", suite.CreatedJWTConfigID), nil, nil)
		suite.makeRequest("PUT", "/api/admin/jwt-configs/"+fmt.Sprintf("%d", suite.CreatedJWTConfigID), map[string]interface{}{
			"algorithm":  "HS256",
			"jwt_secret": "updated-secret-key",
			"issuer":     "updated-issuer",
			"is_active":  true,
		}, nil)
		suite.makeRequest("DELETE", "/api/admin/jwt-configs/"+fmt.Sprintf("%d", suite.CreatedJWTConfigID), nil, nil)
	}

	// 14. External Auth CRUD
	suite.Logger.Log("\n🌐 14. External Auth Providers")
	suite.makeRequest("POST", "/api/admin/external-auth", map[string]interface{}{
		"virtual_directory_id":  suite.CreatedVDirID,
		"auth_url":              "https://auth.example.com/verify",
		"http_method":           "POST",
		"request_timeout_seconds": 5,
		"send_headers":          true,
		"send_body":             false,
		"success_key":           "status",
		"success_value":         "true",
		"message_key":           "message",
		"is_active":             true,
	}, nil)
	suite.makeRequest("GET", "/api/admin/external-auth", nil, nil)
	if suite.CreatedExternalAuthID > 0 {
		suite.makeRequest("GET", "/api/admin/external-auth/"+fmt.Sprintf("%d", suite.CreatedExternalAuthID), nil, nil)
		suite.makeRequest("PUT", "/api/admin/external-auth/"+fmt.Sprintf("%d", suite.CreatedExternalAuthID), map[string]interface{}{
			"auth_url":              "https://auth.example.com/v2/verify",
			"request_timeout_seconds": 10,
			"is_active":             true,
		}, nil)
		suite.makeRequest("DELETE", "/api/admin/external-auth/"+fmt.Sprintf("%d", suite.CreatedExternalAuthID), nil, nil)
	}

	// 15. Rate Limits CRUD
	suite.Logger.Log("\n📊 15. Rate Limits")
	suite.makeRequest("POST", "/api/admin/rate-limits", map[string]interface{}{
		"virtual_directory_id":   suite.CreatedVDirID,
		"limit_by":               "ip",
		"requests_per_minute":    60,
		"burst":                  10,
		"block_duration_seconds": 60,
		"is_active":              true,
	}, nil)
	suite.makeRequest("GET", "/api/admin/rate-limits", nil, nil)
	if suite.CreatedRateLimitID > 0 {
		suite.makeRequest("GET", "/api/admin/rate-limits/"+fmt.Sprintf("%d", suite.CreatedRateLimitID), nil, nil)
		suite.makeRequest("PUT", "/api/admin/rate-limits/"+fmt.Sprintf("%d", suite.CreatedRateLimitID), map[string]interface{}{
			"requests_per_minute": 120,
			"burst":               20,
			"is_active":           true,
		}, nil)
		suite.makeRequest("DELETE", "/api/admin/rate-limits/"+fmt.Sprintf("%d", suite.CreatedRateLimitID), nil, nil)
	}

	// 16. CORS Configs CRUD
	suite.Logger.Log("\n🔗 16. CORS Configs")
	suite.makeRequest("POST", "/api/admin/cors-configs", map[string]interface{}{
		"virtual_directory_id": suite.CreatedVDirID,
		"allowed_origins":      "https://app.example.com,https://admin.example.com",
		"allowed_methods":      "GET,POST,PUT,DELETE,OPTIONS",
		"allowed_headers":      "Content-Type,Authorization,X-API-Key",
		"exposed_headers":      "X-Request-Id",
		"allow_credentials":    true,
		"max_age_seconds":      3600,
		"is_active":            true,
	}, nil)
	suite.makeRequest("GET", "/api/admin/cors-configs", nil, nil)
	if suite.CreatedCORSID > 0 {
		suite.makeRequest("GET", "/api/admin/cors-configs/"+fmt.Sprintf("%d", suite.CreatedCORSID), nil, nil)
		suite.makeRequest("PUT", "/api/admin/cors-configs/"+fmt.Sprintf("%d", suite.CreatedCORSID), map[string]interface{}{
			"allowed_origins":   "*",
			"allow_credentials": false,
			"is_active":         true,
		}, nil)
		suite.makeRequest("DELETE", "/api/admin/cors-configs/"+fmt.Sprintf("%d", suite.CreatedCORSID), nil, nil)
	}

	// 17. Circuit Breakers CRUD
	suite.Logger.Log("\n⚡ 17. Circuit Breakers")
	suite.makeRequest("POST", "/api/admin/circuit-breakers", map[string]interface{}{
		"virtual_directory_id":     suite.CreatedVDirID,
		"enabled":                  true,
		"failure_threshold":        5,
		"recovery_timeout_seconds": 30,
		"half_open_max_requests":   3,
	}, nil)
	suite.makeRequest("GET", "/api/admin/circuit-breakers", nil, nil)
	if suite.CreatedCircuitBreakerID > 0 {
		suite.makeRequest("GET", "/api/admin/circuit-breakers/"+fmt.Sprintf("%d", suite.CreatedCircuitBreakerID), nil, nil)
		suite.makeRequest("PUT", "/api/admin/circuit-breakers/"+fmt.Sprintf("%d", suite.CreatedCircuitBreakerID), map[string]interface{}{
			"enabled":                  true,
			"failure_threshold":        10,
			"recovery_timeout_seconds": 60,
			"half_open_max_requests":   5,
		}, nil)
		suite.makeRequest("DELETE", "/api/admin/circuit-breakers/"+fmt.Sprintf("%d", suite.CreatedCircuitBreakerID), nil, nil)
	}

	// 18. IP Whitelists CRUD
	suite.Logger.Log("\n✅ 18. IP Whitelists")
	suite.makeRequest("POST", "/api/admin/ip-whitelists", map[string]interface{}{
		"virtual_directory_id": suite.CreatedVDirID,
		"ip_address":           "192.168.1.0/24",
		"description":          "Office network",
		"is_active":            true,
	}, nil)
	suite.makeRequest("GET", "/api/admin/ip-whitelists", nil, nil)
	if suite.CreatedIPWhitelistID > 0 {
		suite.makeRequest("GET", "/api/admin/ip-whitelists/"+fmt.Sprintf("%d", suite.CreatedIPWhitelistID), nil, nil)
		if suite.CreatedVDirID > 0 {
			suite.makeRequest("GET", "/api/admin/virtual-directories/"+fmt.Sprintf("%d", suite.CreatedVDirID)+"/ip-whitelists", nil, nil)
		}
		suite.makeRequest("PUT", "/api/admin/ip-whitelists/"+fmt.Sprintf("%d", suite.CreatedIPWhitelistID), map[string]interface{}{
			"ip_address":  "10.0.0.0/8",
			"description": "VPN network",
			"is_active":   true,
		}, nil)
		suite.makeRequest("DELETE", "/api/admin/ip-whitelists/"+fmt.Sprintf("%d", suite.CreatedIPWhitelistID), nil, nil)
	}

	// 19. IP Blacklists CRUD
	suite.Logger.Log("\n❌ 19. IP Blacklists")
	suite.makeRequest("POST", "/api/admin/ip-blacklists", map[string]interface{}{
		"virtual_directory_id": suite.CreatedVDirID,
		"ip_address":           "203.0.113.50",
		"reason":               "Suspicious activity",
		"expired_at":           "2027-01-01 00:00:00",
		"is_active":            true,
	}, nil)
	suite.makeRequest("GET", "/api/admin/ip-blacklists", nil, nil)
	if suite.CreatedIPBlacklistID > 0 {
		suite.makeRequest("GET", "/api/admin/ip-blacklists/"+fmt.Sprintf("%d", suite.CreatedIPBlacklistID), nil, nil)
		if suite.CreatedVDirID > 0 {
			suite.makeRequest("GET", "/api/admin/virtual-directories/"+fmt.Sprintf("%d", suite.CreatedVDirID)+"/ip-blacklists", nil, nil)
		}
		suite.makeRequest("PUT", "/api/admin/ip-blacklists/"+fmt.Sprintf("%d", suite.CreatedIPBlacklistID), map[string]interface{}{
			"ip_address": "203.0.113.50",
			"reason":     "Confirmed attacker",
			"is_active":  true,
		}, nil)
		suite.makeRequest("DELETE", "/api/admin/ip-blacklists/"+fmt.Sprintf("%d", suite.CreatedIPBlacklistID), nil, nil)
	}

	// 20. Request Header Rules CRUD
	suite.Logger.Log("\n📝 20. Request Header Rules")
	suite.makeRequest("POST", "/api/admin/request-headers", map[string]interface{}{
		"virtual_directory_id": suite.CreatedVDirID,
		"header_name":          "X-Request-ID",
		"operation":            "set",
		"value_source":         "static",
		"header_value":         "gateway-request",
		"execution_order":      1,
		"is_active":            true,
	}, nil)
	suite.makeRequest("GET", "/api/admin/request-headers", nil, nil)
	if suite.CreatedRequestHeaderID > 0 {
		suite.makeRequest("GET", "/api/admin/request-headers/"+fmt.Sprintf("%d", suite.CreatedRequestHeaderID), nil, nil)
		if suite.CreatedVDirID > 0 {
			suite.makeRequest("GET", "/api/admin/virtual-directories/"+fmt.Sprintf("%d", suite.CreatedVDirID)+"/request-headers", nil, nil)
		}
		suite.makeRequest("PUT", "/api/admin/request-headers/"+fmt.Sprintf("%d", suite.CreatedRequestHeaderID), map[string]interface{}{
			"header_name":     "X-Forwarded-For",
			"operation":       "add",
			"value_source":    "variable",
			"variable_name":   "client_ip",
			"execution_order": 2,
			"is_active":       true,
		}, nil)
		suite.makeRequest("DELETE", "/api/admin/request-headers/"+fmt.Sprintf("%d", suite.CreatedRequestHeaderID), nil, nil)
	}

	// 21. Response Header Rules CRUD
	suite.Logger.Log("\n📤 21. Response Header Rules")
	suite.makeRequest("POST", "/api/admin/response-headers", map[string]interface{}{
		"virtual_directory_id": suite.CreatedVDirID,
		"header_name":          "X-Powered-By",
		"operation":            "delete",
		"header_value":         "",
		"execution_order":      1,
		"is_active":            true,
	}, nil)
	suite.makeRequest("GET", "/api/admin/response-headers", nil, nil)
	if suite.CreatedResponseHeaderID > 0 {
		suite.makeRequest("GET", "/api/admin/response-headers/"+fmt.Sprintf("%d", suite.CreatedResponseHeaderID), nil, nil)
		if suite.CreatedVDirID > 0 {
			suite.makeRequest("GET", "/api/admin/virtual-directories/"+fmt.Sprintf("%d", suite.CreatedVDirID)+"/response-headers", nil, nil)
		}
		suite.makeRequest("PUT", "/api/admin/response-headers/"+fmt.Sprintf("%d", suite.CreatedResponseHeaderID), map[string]interface{}{
			"header_name":     "X-Gateway",
			"operation":       "set",
			"header_value":    "swantara-gate",
			"execution_order": 1,
			"is_active":       true,
		}, nil)
		suite.makeRequest("DELETE", "/api/admin/response-headers/"+fmt.Sprintf("%d", suite.CreatedResponseHeaderID), nil, nil)
	}

	// 22. Query Rewrites CRUD
	suite.Logger.Log("\n🔄 22. Query Rewrites")
	suite.makeRequest("POST", "/api/admin/query-rewrites", map[string]interface{}{
		"virtual_directory_id": suite.CreatedVDirID,
		"param_name":           "version",
		"param_value":          "v2",
		"operation":            "set",
	}, nil)
	suite.makeRequest("GET", "/api/admin/query-rewrites", nil, nil)
	if suite.CreatedQueryRewriteID > 0 {
		suite.makeRequest("GET", "/api/admin/query-rewrites/"+fmt.Sprintf("%d", suite.CreatedQueryRewriteID), nil, nil)
		if suite.CreatedVDirID > 0 {
			suite.makeRequest("GET", "/api/admin/virtual-directories/"+fmt.Sprintf("%d", suite.CreatedVDirID)+"/query-rewrites", nil, nil)
		}
		suite.makeRequest("PUT", "/api/admin/query-rewrites/"+fmt.Sprintf("%d", suite.CreatedQueryRewriteID), map[string]interface{}{
			"param_name":  "api_version",
			"param_value": "v3",
			"operation":   "set",
		}, nil)
		suite.makeRequest("DELETE", "/api/admin/query-rewrites/"+fmt.Sprintf("%d", suite.CreatedQueryRewriteID), nil, nil)
	}

	// 23. ACME Accounts CRUD
	suite.Logger.Log("\n🔒 23. ACME Accounts")
	suite.makeRequest("POST", "/api/admin/acme-accounts", map[string]interface{}{
		"email":            "admin@example.com",
		"provider_url":     "https://acme-v02.api.letsencrypt.org/directory",
		"account_key_path": "/etc/ssl/acme/account.key",
		"is_default":       true,
	}, nil)
	suite.makeRequest("GET", "/api/admin/acme-accounts", nil, nil)
	if suite.CreatedACMEAccountID > 0 {
		suite.makeRequest("GET", "/api/admin/acme-accounts/"+fmt.Sprintf("%d", suite.CreatedACMEAccountID), nil, nil)
		suite.makeRequest("PUT", "/api/admin/acme-accounts/"+fmt.Sprintf("%d", suite.CreatedACMEAccountID), map[string]interface{}{
			"email":      "newemail@example.com",
			"is_default": true,
		}, nil)
		suite.makeRequest("DELETE", "/api/admin/acme-accounts/"+fmt.Sprintf("%d", suite.CreatedACMEAccountID), nil, nil)
	}

	// 24. SSL Certificates CRUD
	suite.Logger.Log("\n🔐 24. SSL Certificates")
	suite.makeRequest("POST", "/api/admin/ssl-certificates", map[string]interface{}{
		"acme_account_id":     suite.CreatedACMEAccountID,
		"provider":            "letsencrypt",
		"challenge_type":      "http-01",
		"certificate_path":    "/etc/ssl/certs/example.crt",
		"private_key_path":    "/etc/ssl/private/example.key",
		"chain_path":          "/etc/ssl/certs/chain.pem",
		"auto_renew":          true,
		"renew_before_days":   30,
		"expired_at":          "2027-12-31 23:59:59",
		"is_active":           true,
	}, nil)
	suite.makeRequest("GET", "/api/admin/ssl-certificates", nil, nil)
	if suite.CreatedSSLCertID > 0 {
		suite.makeRequest("GET", "/api/admin/ssl-certificates/"+fmt.Sprintf("%d", suite.CreatedSSLCertID), nil, nil)
		suite.makeRequest("PUT", "/api/admin/ssl-certificates/"+fmt.Sprintf("%d", suite.CreatedSSLCertID), map[string]interface{}{
			"auto_renew":        true,
			"renew_before_days": 14,
			"is_active":         true,
		}, nil)
		suite.makeRequest("DELETE", "/api/admin/ssl-certificates/"+fmt.Sprintf("%d", suite.CreatedSSLCertID), nil, nil)
	}

	// 25. Certificate Domains CRUD
	suite.Logger.Log("\n🌍 25. Certificate Domains")
	if suite.CreatedSSLCertID > 0 {
		suite.makeRequest("POST", "/api/admin/certificate-domains", map[string]interface{}{
			"ssl_certificate_id": suite.CreatedSSLCertID,
			"domain_name":        "example.com",
			"is_wildcard":        false,
		}, nil)
	} else {
		suite.makeRequest("POST", "/api/admin/certificate-domains", map[string]interface{}{
			"ssl_certificate_id": 1,
			"domain_name":        "example.com",
			"is_wildcard":        false,
		}, nil)
	}
	suite.makeRequest("GET", "/api/admin/certificate-domains", nil, nil)
	if suite.CreatedCertDomainID > 0 {
		suite.makeRequest("GET", "/api/admin/certificate-domains/"+fmt.Sprintf("%d", suite.CreatedCertDomainID), nil, nil)
		if suite.CreatedSSLCertID > 0 {
			suite.makeRequest("GET", "/api/admin/ssl-certificates/"+fmt.Sprintf("%d", suite.CreatedSSLCertID)+"/domains", nil, nil)
		}
		suite.makeRequest("PUT", "/api/admin/certificate-domains/"+fmt.Sprintf("%d", suite.CreatedCertDomainID), map[string]interface{}{
			"domain_name":  "*.example.com",
			"is_wildcard":  true,
		}, nil)
		suite.makeRequest("DELETE", "/api/admin/certificate-domains/"+fmt.Sprintf("%d", suite.CreatedCertDomainID), nil, nil)
	}

	// 26. SSL Bindings CRUD
	suite.Logger.Log("\n🔗 26. SSL Bindings")
	if suite.CreatedSSLCertID > 0 {
		suite.makeRequest("POST", "/api/admin/ssl-bindings", map[string]interface{}{
			"ssl_certificate_id": suite.CreatedSSLCertID,
			"binding_type":       "host",
			"host_id":            1,
			"is_default":         true,
			"priority":           1,
		}, nil)
	} else {
		suite.makeRequest("POST", "/api/admin/ssl-bindings", map[string]interface{}{
			"ssl_certificate_id": 1,
			"binding_type":       "host",
			"host_id":            1,
			"is_default":         true,
			"priority":           1,
		}, nil)
	}
	suite.makeRequest("GET", "/api/admin/ssl-bindings", nil, nil)
	if suite.CreatedSSLBindingID > 0 {
		suite.makeRequest("GET", "/api/admin/ssl-bindings/"+fmt.Sprintf("%d", suite.CreatedSSLBindingID), nil, nil)
		suite.makeRequest("PUT", "/api/admin/ssl-bindings/"+fmt.Sprintf("%d", suite.CreatedSSLBindingID), map[string]interface{}{
			"binding_type":  "virtual_host",
			"virtual_host_id": 1,
			"is_default":    false,
			"priority":      2,
		}, nil)
		suite.makeRequest("DELETE", "/api/admin/ssl-bindings/"+fmt.Sprintf("%d", suite.CreatedSSLBindingID), nil, nil)
	}

	// 27. TLS Options CRUD
	suite.Logger.Log("\n🔏 27. TLS Options")
	suite.makeRequest("POST", "/api/admin/tls-options", map[string]interface{}{
		"binding_type":     "host",
		"host_id":          1,
		"min_tls_version":  "1.2",
		"http2_enabled":    true,
		"hsts_enabled":     true,
		"hsts_max_age":     31536000,
	}, nil)
	suite.makeRequest("GET", "/api/admin/tls-options", nil, nil)
	if suite.CreatedTLSOptionID > 0 {
		suite.makeRequest("GET", "/api/admin/tls-options/"+fmt.Sprintf("%d", suite.CreatedTLSOptionID), nil, nil)
		suite.makeRequest("PUT", "/api/admin/tls-options/"+fmt.Sprintf("%d", suite.CreatedTLSOptionID), map[string]interface{}{
			"min_tls_version": "1.3",
			"http2_enabled":   true,
			"hsts_enabled":    true,
			"hsts_max_age":    63072000,
		}, nil)
		suite.makeRequest("DELETE", "/api/admin/tls-options/"+fmt.Sprintf("%d", suite.CreatedTLSOptionID), nil, nil)
	}

	// 28. Service Discovery CRUD
	suite.Logger.Log("\n🔍 28. Service Discovery")
	suite.makeRequest("POST", "/api/admin/service-discovery", map[string]interface{}{
		"virtual_host_id":          suite.CreatedVHostID,
		"provider":                 "consul",
		"endpoint_url":             "http://consul:8500/v1/catalog/service/my-api",
		"refresh_interval_seconds": 30,
		"is_active":                true,
	}, nil)
	suite.makeRequest("GET", "/api/admin/service-discovery", nil, nil)
	if suite.CreatedServiceDiscoveryID > 0 {
		suite.makeRequest("GET", "/api/admin/service-discovery/"+fmt.Sprintf("%d", suite.CreatedServiceDiscoveryID), nil, nil)
		suite.makeRequest("PUT", "/api/admin/service-discovery/"+fmt.Sprintf("%d", suite.CreatedServiceDiscoveryID), map[string]interface{}{
			"provider":                 "consul",
			"endpoint_url":             "http://consul:8500/v1/catalog/service/updated",
			"refresh_interval_seconds": 60,
			"is_active":                true,
		}, nil)
		suite.makeRequest("DELETE", "/api/admin/service-discovery/"+fmt.Sprintf("%d", suite.CreatedServiceDiscoveryID), nil, nil)
	}

	// 29. Config Versions CRUD
	suite.Logger.Log("\n📚 29. Config Versions")
	suite.makeRequest("POST", "/api/admin/config-versions", map[string]interface{}{
		"config_name":    "proxy-routes",
		"version_number": 1,
		"changed_by":     "admin",
		"notes":          "Initial configuration",
	}, nil)
	suite.makeRequest("GET", "/api/admin/config-versions", nil, nil)
	if suite.CreatedConfigVersionID > 0 {
		suite.makeRequest("GET", "/api/admin/config-versions/"+fmt.Sprintf("%d", suite.CreatedConfigVersionID), nil, nil)
		suite.makeRequest("PUT", "/api/admin/config-versions/"+fmt.Sprintf("%d", suite.CreatedConfigVersionID), map[string]interface{}{
			"config_name":    "proxy-routes",
			"version_number": 2,
			"changed_by":     "admin",
			"notes":          "Added new routes",
		}, nil)
		suite.makeRequest("DELETE", "/api/admin/config-versions/"+fmt.Sprintf("%d", suite.CreatedConfigVersionID), nil, nil)
	}

	// 30. Maintenance Windows CRUD
	suite.Logger.Log("\n🔧 30. Maintenance Windows")
	suite.makeRequest("POST", "/api/admin/maintenance-windows", map[string]interface{}{
		"virtual_host_id":          suite.CreatedVHostID,
		"title":                    "Scheduled Maintenance",
		"start_at":                 "2026-06-01 00:00:00",
		"end_at":                   "2026-06-01 04:00:00",
		"maintenance_response_code": 503,
		"maintenance_message":      "Service is under scheduled maintenance. Please try again later.",
		"is_active":                true,
	}, nil)
	suite.makeRequest("GET", "/api/admin/maintenance-windows", nil, nil)
	if suite.CreatedMaintenanceWindowID > 0 {
		suite.makeRequest("GET", "/api/admin/maintenance-windows/"+fmt.Sprintf("%d", suite.CreatedMaintenanceWindowID), nil, nil)
		suite.makeRequest("PUT", "/api/admin/maintenance-windows/"+fmt.Sprintf("%d", suite.CreatedMaintenanceWindowID), map[string]interface{}{
			"title":               "Extended Maintenance",
			"end_at":              "2026-06-01 06:00:00",
			"maintenance_message": "Extended maintenance in progress",
			"is_active":           true,
		}, nil)
		suite.makeRequest("DELETE", "/api/admin/maintenance-windows/"+fmt.Sprintf("%d", suite.CreatedMaintenanceWindowID), nil, nil)
	}

	// Print Summary
	suite.PrintSummary()
}

func (suite *APITestSuite) PrintSummary() {
	suite.Logger.Log("\n\n" + strings.Repeat("=", 60))
	suite.Logger.Log("📊 TEST SUMMARY")
	suite.Logger.Log(strings.Repeat("=", 60))
	
	passed := 0
	failed := 0
	
	for _, r := range suite.Results {
		if r.Success {
			passed++
		} else {
			failed++
		}
	}
	
	suite.Logger.Log("✅ Passed: %d", passed)
	suite.Logger.Log("❌ Failed: %d", failed)
	suite.Logger.Log("📝 Total:  %d", len(suite.Results))
	suite.Logger.Log("")
	
	// Generate documentation
	suite.GenerateDocumentation()
}

func (suite *APITestSuite) GenerateDocumentation() {
	suite.Logger.Log("📄 Generating API Documentation...")
	
	// Generate timestamp for filename
	now := time.Now()
	timestamp := now.Format("2006-01-02 15_04")
	filename := fmt.Sprintf("RESULT-%s.md", timestamp)
	
	doc := "# Swantara Gate API Test Results\n\n"
	doc += fmt.Sprintf("**Test Date:** %s\n\n", now.Format("January 2, 2006 at 15:04:05"))
	doc += fmt.Sprintf("**Base URL:** `%s`\n\n", suite.BaseURL)
	doc += "**Credentials:**\n- Username: `testcli`\n- Password: `admin1324`\n\n"
	doc += "---\n\n"
	
	// Summary section
	totalTests := len(suite.Results)
	passed := 0
	failed := 0
	totalDuration := time.Duration(0)
	
	for _, r := range suite.Results {
		if r.Success {
			passed++
		} else {
			failed++
		}
		totalDuration += r.Duration
	}
	
	doc += "## 📊 Summary\n\n"
	doc += fmt.Sprintf("- ✅ **Passed:** %d\n", passed)
	doc += fmt.Sprintf("- ❌ **Failed:** %d\n", failed)
	doc += fmt.Sprintf("- 📝 **Total Tests:** %d\n", totalTests)
	doc += fmt.Sprintf("- ⏱️ **Total Duration:** %v\n\n", totalDuration.Round(time.Millisecond))
	doc += fmt.Sprintf("- 📈 **Success Rate:** %.1f%%\n\n", float64(passed)/float64(totalTests)*100)
	doc += "---\n\n"
	
	// Detailed test results grouped by category
	currentCategory := ""
	
	for i, r := range suite.Results {
		// Determine category
		category := getCategory(r.Endpoint)
		if category != currentCategory {
			currentCategory = category
			doc += fmt.Sprintf("## %s\n\n", category)
		}
		
		doc += fmt.Sprintf("### %d. %s %s\n\n", i+1, r.Method, r.Endpoint)
		
		// Request details
		doc += "#### 📤 Request\n\n"
		
		// Headers
		if len(r.RequestHeaders) > 0 {
			doc += "**Headers:**\n\n"
			doc += "| Header | Value |\n|--------|-------|\n"
			for k, v := range r.RequestHeaders {
				doc += fmt.Sprintf("| %s | %s |\n", k, v)
			}
			doc += "\n"
		}
		
		// Body
		if r.RequestBody != "" {
			doc += "**Request Body:**\n\n```json\n" + formatJSON(r.RequestBody) + "\n```\n\n"
		} else {
			doc += "**Request Body:** None\n\n"
		}
		
		// Response details
		doc += "#### 📥 Response\n\n"
		doc += fmt.Sprintf("**Status Code:** `%d`", r.Status)
		
		if r.Success {
			doc += " ✅\n\n"
		} else {
			doc += " ❌\n\n"
		}
		
		if r.ResponseBody != "" {
			doc += "**Response Body:**\n\n```json\n" + formatJSON(r.ResponseBody) + "\n```\n\n"
		}
		
		if r.Error != "" {
			doc += fmt.Sprintf("**Error:** `%s`\n\n", r.Error)
		}
		
		doc += fmt.Sprintf("**Duration:** %v\n\n---\n\n", r.Duration.Round(time.Millisecond))
	}
	
	// Write to file
	err := os.WriteFile(filename, []byte(doc), 0644)
	if err != nil {
		suite.Logger.Log("❌ Error writing documentation: %v", err)
		return
	}
	
	suite.Logger.Log("✅ Documentation saved to %s", filename)
	
	// Generate failed tests report
	suite.generateFailedReport()
}

func (suite *APITestSuite) generateFailedReport() {
	// Filter failed tests
	var failedTests []TestResult
	for _, r := range suite.Results {
		if !r.Success {
			failedTests = append(failedTests, r)
		}
	}
	
	if len(failedTests) == 0 {
		suite.Logger.Log("✅ No failed tests, skipping failed report")
		return
	}
	
	suite.Logger.Log("📄 Generating Failed Tests Report...")
	
	// Generate timestamp for filename
	now := time.Now()
	timestamp := now.Format("2006-01-02 15_04")
	filename := fmt.Sprintf("RESULT-%s-failed.md", timestamp)
	
	doc := "# Swantara Gate API - Failed Tests Report\n\n"
	doc += fmt.Sprintf("**Test Date:** %s\n\n", now.Format("January 2, 2006 at 15:04:05"))
	doc += fmt.Sprintf("**Base URL:** `%s`\n\n", suite.BaseURL)
	doc += "---\n\n"
	
	// Summary
	doc += "## 📊 Failed Tests Summary\n\n"
	doc += fmt.Sprintf("- ❌ **Total Failed:** %d\n", len(failedTests))
	doc += fmt.Sprintf("- 📝 **Total Tests:** %d\n", len(suite.Results))
	doc += fmt.Sprintf("- 📈 **Failure Rate:** %.1f%%\n\n", float64(len(failedTests))/float64(len(suite.Results))*100)
	doc += "---\n\n"
	
	// Group by category
	currentCategory := ""
	
	for _, r := range failedTests {
		// Determine category
		category := getCategory(r.Endpoint)
		if category != currentCategory {
			currentCategory = category
			doc += fmt.Sprintf("## %s\n\n", category)
		}
		
		// Find original index
		originalIndex := 0
		for j, original := range suite.Results {
			if original.Endpoint == r.Endpoint && original.Method == r.Method && original.Status == r.Status {
				originalIndex = j + 1
				break
			}
		}
		
		doc += fmt.Sprintf("### %d. %s %s\n\n", originalIndex, r.Method, r.Endpoint)
		
		// Request details
		doc += "#### 📤 Request\n\n"
		
		// Headers
		if len(r.RequestHeaders) > 0 {
			doc += "**Headers:**\n\n"
			doc += "| Header | Value |\n|--------|-------|\n"
			for k, v := range r.RequestHeaders {
				doc += fmt.Sprintf("| %s | %s |\n", k, v)
			}
			doc += "\n"
		}
		
		// Body
		if r.RequestBody != "" {
			doc += "**Request Body:**\n\n```json\n" + formatJSON(r.RequestBody) + "\n```\n\n"
		} else {
			doc += "**Request Body:** None\n\n"
		}
		
		// Response details
		doc += "#### 📥 Response\n\n"
		doc += fmt.Sprintf("**Status Code:** `%d` ❌\n\n", r.Status)
		
		if r.ResponseBody != "" {
			doc += "**Response Body:**\n\n```json\n" + formatJSON(r.ResponseBody) + "\n```\n\n"
		}
		
		if r.Error != "" {
			doc += fmt.Sprintf("**Error:** `%s`\n\n", r.Error)
		}
		
		doc += fmt.Sprintf("**Duration:** %v\n\n", r.Duration.Round(time.Millisecond))
		
		// Analysis
		doc += "#### 🔍 Analysis\n\n"
		if r.Status == 400 {
			doc += "**Type:** Bad Request - Invalid request body or parameters\n\n"
		} else if r.Status == 404 {
			doc += "**Type:** Not Found - Resource doesn't exist or endpoint incorrect\n\n"
		} else if r.Status == 500 {
			doc += "**Type:** Internal Server Error - Server-side issue\n\n"
		} else if r.Status == 401 || r.Status == 403 {
			doc += "**Type:** Authentication/Authorization Error\n\n"
		} else if r.Error != "" {
			doc += fmt.Sprintf("**Type:** Connection/Network Error\n\n")
		}
		
		doc += "---\n\n"
	}
	
	// Recommendations
	doc += "## 💡 Recommendations\n\n"
	doc += "### Common Issues:\n\n"
	
	// Count error types
	badRequest := 0
	notFound := 0
	serverError := 0
	otherError := 0
	
	for _, r := range failedTests {
		if r.Status == 400 {
			badRequest++
		} else if r.Status == 404 {
			notFound++
		} else if r.Status >= 500 {
			serverError++
		} else {
			otherError++
		}
	}
	
	if badRequest > 0 {
		doc += fmt.Sprintf("1. **Bad Request Errors (%d):** Check request body format and required fields\n", badRequest)
	}
	if notFound > 0 {
		doc += fmt.Sprintf("2. **Not Found Errors (%d):** Verify resource IDs exist before making requests\n", notFound)
	}
	if serverError > 0 {
		doc += fmt.Sprintf("3. **Server Errors (%d):** Check server logs for internal errors\n", serverError)
	}
	if otherError > 0 {
		doc += fmt.Sprintf("4. **Other Errors (%d):** Review authentication and network connectivity\n", otherError)
	}
	
	doc += "\n### Next Steps:\n\n"
	doc += "1. Review the failed test details above\n"
	doc += "2. Check server logs at `logs/` directory\n"
	doc += "3. Verify database state and existing resources\n"
	doc += "4. Re-run tests after fixing issues\n"
	
	// Write to file
	err := os.WriteFile(filename, []byte(doc), 0644)
	if err != nil {
		suite.Logger.Log("❌ Error writing failed report: %v", err)
		return
	}
	
	suite.Logger.Log("✅ Failed tests report saved to %s", filename)
}

func (suite *APITestSuite) saveTestResults() {
	type TestReport struct {
		TotalTests int           `json:"total_tests"`
		Passed     int           `json:"passed"`
		Failed     int           `json:"failed"`
		Duration   time.Duration `json:"total_duration"`
		Results    []TestResult  `json:"results"`
	}
	
	report := TestReport{
		TotalTests: len(suite.Results),
		Results:    suite.Results,
	}
	
	for _, r := range suite.Results {
		if r.Success {
			report.Passed++
		} else {
			report.Failed++
		}
		report.Duration += r.Duration
	}
	
	jsonData, _ := json.MarshalIndent(report, "", "  ")
	os.WriteFile("test_results.json", jsonData, 0644)
	suite.Logger.Log("✅ Test results saved to test_results.json")
	
	// Print failed endpoints detail
	suite.Logger.Log("\n📋 Failed Endpoints Detail:")
	suite.Logger.Log(strings.Repeat("-", 60))
	for _, r := range suite.Results {
		if !r.Success {
			suite.Logger.Log("❌ %s %s - Status: %d", r.Method, r.Endpoint, r.Status)
			if r.ResponseBody != "" {
				suite.Logger.Log("   Response: %s", r.ResponseBody)
			}
			if r.Error != "" {
				suite.Logger.Log("   Error: %s", r.Error)
			}
		}
	}
}

func getCategory(endpoint string) string {
	if strings.Contains(endpoint, "/health") {
		return "Health Check"
	} else if strings.Contains(endpoint, "/auth/") {
		return "Authentication"
	} else if strings.Contains(endpoint, "/config/") {
		return "Configuration"
	} else if strings.Contains(endpoint, "/users") {
		return "Users Management"
	} else if strings.Contains(endpoint, "/consumers") && !strings.Contains(endpoint, "credentials") {
		return "API Consumers"
	} else if strings.Contains(endpoint, "/consumer-credentials") {
		return "Consumer Credentials"
	} else if strings.Contains(endpoint, "/api-keys") {
		return "API Keys"
	} else if strings.Contains(endpoint, "/route-access") {
		return "Route Consumer Access (ACL)"
	} else if strings.Contains(endpoint, "/hosts") && !strings.Contains(endpoint, "virtual") && !strings.Contains(endpoint, "certificate") {
		return "Hosts"
	} else if strings.Contains(endpoint, "virtual-hosts") {
		return "Virtual Hosts"
	} else if strings.Contains(endpoint, "/upstream-servers") {
		return "Upstream Servers"
	} else if strings.Contains(endpoint, "virtual-directories") {
		return "Routes (Virtual Directories)"
	} else if strings.Contains(endpoint, "/jwt-configs") {
		return "JWT Configs"
	} else if strings.Contains(endpoint, "/external-auth") {
		return "External Auth"
	} else if strings.Contains(endpoint, "/rate-limits") {
		return "Rate Limits"
	} else if strings.Contains(endpoint, "/cors-configs") {
		return "CORS Configuration"
	} else if strings.Contains(endpoint, "/circuit-breakers") {
		return "Circuit Breakers"
	} else if strings.Contains(endpoint, "/ip-whitelists") {
		return "IP Whitelists"
	} else if strings.Contains(endpoint, "/ip-blacklists") {
		return "IP Blacklists"
	} else if strings.Contains(endpoint, "/request-headers") {
		return "Request Header Rules"
	} else if strings.Contains(endpoint, "/response-headers") {
		return "Response Header Rules"
	} else if strings.Contains(endpoint, "/query-rewrites") {
		return "Query Rewrites"
	} else if strings.Contains(endpoint, "/acme-accounts") {
		return "ACME Accounts"
	} else if strings.Contains(endpoint, "/ssl-certificates") {
		return "SSL Certificates"
	} else if strings.Contains(endpoint, "/certificate-domains") {
		return "Certificate Domains"
	} else if strings.Contains(endpoint, "/ssl-bindings") {
		return "SSL Bindings"
	} else if strings.Contains(endpoint, "/tls-options") {
		return "TLS Options"
	} else if strings.Contains(endpoint, "/service-discovery") {
		return "Service Discovery"
	} else if strings.Contains(endpoint, "/config-versions") {
		return "Config Versions"
	} else if strings.Contains(endpoint, "/maintenance-windows") {
		return "Maintenance Windows"
	}
	return "Other"
}

func formatJSON(jsonStr string) string {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(jsonStr), "", "  "); err != nil {
		return jsonStr
	}
	return prettyJSON.String()
}

func main() {
	baseURL := "http://localhost:8081"
	
	if len(os.Args) > 1 {
		baseURL = os.Args[1]
	}
	
	// Create logger
	logger, err := NewLogger("api_test.log")
	if err != nil {
		fmt.Printf("❌ Error creating log file: %v\n", err)
		os.Exit(1)
	}
	defer logger.Close()
	
	suite := NewAPITestSuite(baseURL)
	suite.Logger = logger
	suite.RunAllTests()
}
