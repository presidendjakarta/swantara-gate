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
	CreatedRateLimitID    int
	CreatedCORSID         int
	CreatedCircuitBreakerID int
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
					if strings.Contains(endpoint, "/consumers") && !strings.Contains(endpoint, "credentials") {
						suite.CreatedConsumerID = int(id)
						suite.Logger.Log("🆔 Created Consumer ID: %d", suite.CreatedConsumerID)
					} else if strings.Contains(endpoint, "/hosts") && !strings.Contains(endpoint, "virtual") {
						suite.CreatedHostID = int(id)
						suite.Logger.Log("🆔 Created Host ID: %d", suite.CreatedHostID)
					} else if strings.Contains(endpoint, "virtual-hosts") {
						suite.CreatedVHostID = int(id)
						suite.Logger.Log("🆔 Created VHost ID: %d", suite.CreatedVHostID)
					} else if strings.Contains(endpoint, "virtual-directories") {
						suite.CreatedVDirID = int(id)
						suite.Logger.Log("🆔 Created VDir ID: %d", suite.CreatedVDirID)
					} else if strings.Contains(endpoint, "/rate-limits") {
						suite.CreatedRateLimitID = int(id)
						suite.Logger.Log("🆔 Created RateLimit ID: %d", suite.CreatedRateLimitID)
					} else if strings.Contains(endpoint, "/cors-configs") {
						suite.CreatedCORSID = int(id)
						suite.Logger.Log("🆔 Created CORS ID: %d", suite.CreatedCORSID)
					} else if strings.Contains(endpoint, "/circuit-breakers") {
						suite.CreatedCircuitBreakerID = int(id)
						suite.Logger.Log("🆔 Created CircuitBreaker ID: %d", suite.CreatedCircuitBreakerID)
					}
				}
			}
		}
	}

	suite.Results = append(suite.Results, result)
	return result
}

func (suite *APITestSuite) RunAllTests() {
	suite.Logger.Log("🚀 Starting Swantara Gate API Test Suite")
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
		"password": "admin1324", // Try common password
	}, nil)
	
	// If login failed, try creating admin user first via CLI or use testuser
	// For now, continue - if auth failed, all tests will fail with 401

	// Get Current User (Me)
	suite.makeRequest("GET", "/api/admin/auth/me", nil, nil)

	// Refresh Token
	suite.makeRequest("POST", "/api/admin/auth/refresh", map[string]interface{}{
		"refresh_token": suite.RefreshToken,
	}, nil)

	// Note: Logout is SKIPPED to prevent token blacklisting during tests
	// suite.makeRequest("POST", "/api/admin/auth/logout", ...)
	
	suite.Logger.Log("\n✅ Authentication tests completed, continuing with other endpoints...")

	// 3. Configuration
	suite.Logger.Log("\n⚙️  3. Configuration")
	suite.makeRequest("POST", "/api/admin/config/reload", nil, nil)

	// 4. Users CRUD
	suite.Logger.Log("\n👥 4. Users")
	
	// Get All Users first (to see current state)
	suite.makeRequest("GET", "/api/admin/users", nil, nil)
	
	// Create User with UNIQUE username
	suite.makeRequest("POST", "/api/admin/users", map[string]interface{}{
		"username":  "apitest_" + testSuffix,
		"password":  "test123",
		"full_name": "API Test User",
		"email":     "apitest" + testSuffix + "@example.com",
		"role":      "admin",
		"is_active": true,
	}, nil)

	// Get User by ID (use ID 1 - should exist)
	suite.makeRequest("GET", "/api/admin/users/1", nil, nil)

	// Update User (use ID 1)
	suite.makeRequest("PUT", "/api/admin/users/1", map[string]interface{}{
		"full_name": "Updated Name",
		"email":     "updated@example.com",
		"role":      "admin",
		"is_active": true,
	}, nil)

	// Delete the user we created (ID 4 = testuser)
	// Don't delete ID 1 - it might be needed
	// suite.makeRequest("DELETE", "/api/admin/users/4", nil, nil)

	// 5. Consumers CRUD
	suite.Logger.Log("\n🏢 5. API Consumers")
	
	suite.makeRequest("POST", "/api/admin/consumers", map[string]interface{}{
		"consumer_name": "apitest-" + testSuffix,
		"description":   "API Test Application",
		"contact_email": "apitest" + testSuffix + "@app.com",
		"is_active":     true,
	}, nil)

	suite.makeRequest("GET", "/api/admin/consumers", nil, nil)
	
	// Use dynamically created ID
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
		"host_id":         1, // Reference existing host
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
			"lb_algorithm":    "weighted_round_robin",
			"sticky_session":  true,
			"failover_mode":   "active-passive",
			"is_active":       true,
		}, nil)

		// Don't delete VHost - needed for downstream tests
		// suite.makeRequest("DELETE", "/api/admin/virtual-hosts/"+fmt.Sprintf("%d", suite.CreatedVHostID), nil, nil)
	} else {
		// If VHost creation failed, try to use an existing one
		suite.Logger.Log("\n⚠️  VHost creation failed, using existing VHost ID 1")
		suite.CreatedVHostID = 1
	}

	// 8. Virtual Directories (Routes) CRUD
	suite.Logger.Log("\n🛣️  8. Virtual Directories (Routes)")
	
	suite.makeRequest("POST", "/api/admin/virtual-directories", map[string]interface{}{
		"virtual_host_id":      suite.CreatedVHostID, // Use created VHost
		"source_path":            "/api/v1",
		"target_path":            "/",
		"match_type":             "prefix",
		"strip_prefix":           true,
		"preserve_host_header":   false,
		"auth_type":              "none",
		"is_active":              true,
		"proxy_timeout_seconds":  30,
		"retry_count":            2,
		"retry_delay_ms":         100,
		"max_request_size_mb":    10,
		"websocket_enabled":      false,
		"cache_enabled":          false,
		"cache_ttl_seconds":      0,
	}, nil)

	suite.makeRequest("GET", "/api/admin/virtual-directories", nil, nil)
	
	if suite.CreatedVDirID > 0 {
		suite.makeRequest("GET", "/api/admin/virtual-directories/"+fmt.Sprintf("%d", suite.CreatedVDirID), nil, nil)

		suite.makeRequest("PUT", "/api/admin/virtual-directories/"+fmt.Sprintf("%d", suite.CreatedVDirID), map[string]interface{}{
			"source_path":            "/api/v2",
			"target_path":            "/v2",
			"match_type":             "prefix",
			"strip_prefix":           true,
			"is_active":              true,
		}, nil)

		// Don't delete VDir - needed for downstream tests
		// suite.makeRequest("DELETE", "/api/admin/virtual-directories/"+fmt.Sprintf("%d", suite.CreatedVDirID), nil, nil)
	}

	// 9. Rate Limits CRUD
	suite.Logger.Log("\n📊 9. Rate Limits")
	
	suite.makeRequest("POST", "/api/admin/rate-limits", map[string]interface{}{
		"virtual_directory_id": suite.CreatedVDirID, // Use created VDir
		"limit_by":             "ip",
		"requests_per_minute":  60,
		"burst":                10,
		"block_duration_seconds": 60,
		"is_active":            true,
	}, nil)

	suite.makeRequest("GET", "/api/admin/rate-limits", nil, nil)
	
	if suite.CreatedRateLimitID > 0 {
		suite.makeRequest("GET", "/api/admin/rate-limits/"+fmt.Sprintf("%d", suite.CreatedRateLimitID), nil, nil)

		suite.makeRequest("PUT", "/api/admin/rate-limits/"+fmt.Sprintf("%d", suite.CreatedRateLimitID), map[string]interface{}{
			"requests_per_minute":  120,
			"burst":                20,
			"is_active":            true,
		}, nil)

		suite.makeRequest("DELETE", "/api/admin/rate-limits/"+fmt.Sprintf("%d", suite.CreatedRateLimitID), nil, nil)
	}

	// 10. CORS Configs CRUD
	suite.Logger.Log("\n🔗 10. CORS Configs")
	
	suite.makeRequest("POST", "/api/admin/cors-configs", map[string]interface{}{
		"virtual_directory_id": suite.CreatedVDirID, // Use created VDir
		"allowed_origins":   "https://app.example.com",
		"allowed_methods":   "GET,POST,PUT,DELETE,OPTIONS",
		"allowed_headers":   "Content-Type,Authorization",
		"allow_credentials": true,
		"max_age_seconds":   3600,
		"is_active":         true,
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

	// 11. Circuit Breakers CRUD
	suite.Logger.Log("\n⚡ 11. Circuit Breakers")
	
	suite.makeRequest("POST", "/api/admin/circuit-breakers", map[string]interface{}{
		"virtual_directory_id":    suite.CreatedVDirID, // Use created VDir
		"enabled":                 true,
		"failure_threshold":       5,
		"recovery_timeout_seconds": 30,
		"half_open_max_requests":  3,
	}, nil)

	suite.makeRequest("GET", "/api/admin/circuit-breakers", nil, nil)
	
	if suite.CreatedCircuitBreakerID > 0 {
		suite.makeRequest("GET", "/api/admin/circuit-breakers/"+fmt.Sprintf("%d", suite.CreatedCircuitBreakerID), nil, nil)

		suite.makeRequest("PUT", "/api/admin/circuit-breakers/"+fmt.Sprintf("%d", suite.CreatedCircuitBreakerID), map[string]interface{}{
			"enabled":                 true,
			"failure_threshold":       10,
			"recovery_timeout_seconds": 60,
			"half_open_max_requests":  5,
		}, nil)

		suite.makeRequest("DELETE", "/api/admin/circuit-breakers/"+fmt.Sprintf("%d", suite.CreatedCircuitBreakerID), nil, nil)
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
	
	doc := "# Swantara Gate API Documentation\n\n"
	doc += fmt.Sprintf("**Base URL:** `%s`\n\n", suite.BaseURL)
	doc += "**Credentials:**\n- Username: `admin`\n- Password: `admin1324`\n\n"
	doc += "---\n\n"
	
	currentCategory := ""
	
	for i, r := range suite.Results {
		// Determine category
		category := getCategory(r.Endpoint)
		if category != currentCategory {
			currentCategory = category
			doc += fmt.Sprintf("## %s\n\n", category)
		}
		
		doc += fmt.Sprintf("### %d. %s %s\n\n", i+1, r.Method, r.Endpoint)
		
		if r.RequestBody != "" {
			doc += "**Request Body:**\n```json\n" + formatJSON(r.RequestBody) + "\n```\n\n"
		}
		
		doc += fmt.Sprintf("**Status Code:** `%d`\n\n", r.Status)
		
		if r.ResponseBody != "" {
			doc += "**Response:**\n```json\n" + formatJSON(r.ResponseBody) + "\n```\n\n"
		}
		
		if r.Error != "" {
			doc += fmt.Sprintf("**Error:** `%s`\n\n", r.Error)
		}
		
		doc += fmt.Sprintf("**Duration:** %v\n\n---\n\n", r.Duration.Round(time.Millisecond))
	}
	
	// Write to file
	err := os.WriteFile("API_DOCUMENTATION.md", []byte(doc), 0644)
	if err != nil {
		suite.Logger.Log("❌ Error writing documentation: %v", err)
		return
	}
	
	suite.Logger.Log("✅ Documentation saved to API_DOCUMENTATION.md")
	
	// Also save detailed test results
	suite.saveTestResults()
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
