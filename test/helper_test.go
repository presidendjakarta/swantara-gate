package test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/presidendjakarta/swantara-gate/internal/handler"
	"github.com/presidendjakarta/swantara-gate/internal/middleware"
	"github.com/presidendjakarta/swantara-gate/internal/repository"
	"github.com/presidendjakarta/swantara-gate/internal/service"

	_ "modernc.org/sqlite"
)

// TestServer menyimpan instance test server dan dependensinya
type TestServer struct {
	Server *httptest.Server
	DB     *sql.DB
}

// APIResponse struktur response standar dari API
type APIResponse struct {
	Success bool                   `json:"success"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
	Error   string                 `json:"error"`
}

// SetupTestServer menginisialisasi server test dengan database in-memory
func SetupTestServer(t *testing.T) *TestServer {
	t.Helper()

	// Buka database SQLite in-memory
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("Gagal membuka database test: %v", err)
	}

	// Jalankan migrasi dari file SQL
	sqlPath := "../data/database.sql"
	sqlBytes, err := os.ReadFile(sqlPath)
	if err != nil {
		t.Fatalf("Gagal membaca file SQL: %v", err)
	}

	if _, err := db.Exec(string(sqlBytes)); err != nil {
		t.Fatalf("Gagal menjalankan migrasi: %v", err)
	}

	// Inisialisasi semua layer
	// Repositories
	userRepo := repository.NewUserRepository(db)
	consumerRepo := repository.NewAPIConsumerRepository(db)
	hostRepo := repository.NewHostRepository(db)
	vhostRepo := repository.NewVirtualHostRepository(db)
	upstreamRepo := repository.NewUpstreamServerRepository(db)
	vdirRepo := repository.NewVirtualDirectoryRepository(db)
	vdirMethodRepo := repository.NewVirtualDirectoryMethodRepository(db)
	credRepo := repository.NewConsumerCredentialRepository(db)
	apiKeyRepo := repository.NewAPIKeyRepository(db)
	accessRepo := repository.NewRouteConsumerAccessRepository(db)
	jwtConfigRepo := repository.NewJWTConfigRepository(db)
	extAuthRepo := repository.NewExternalAuthRepository(db)
	rateLimitRepo := repository.NewRateLimitRepository(db)
	corsConfigRepo := repository.NewCORSConfigRepository(db)
	circuitBreakerRepo := repository.NewCircuitBreakerRepository(db)
	ipWhitelistRepo := repository.NewIPWhitelistRepository(db)
	ipBlacklistRepo := repository.NewIPBlacklistRepository(db)
	reqHeaderRepo := repository.NewRequestHeaderRuleRepository(db)
	resHeaderRepo := repository.NewResponseHeaderRuleRepository(db)
	queryRewriteRepo := repository.NewQueryRewriteRepository(db)
	acmeRepo := repository.NewACMEAccountRepository(db)
	sslCertRepo := repository.NewSSLCertificateRepository(db)
	certDomainRepo := repository.NewCertificateDomainRepository(db)
	sslBindingRepo := repository.NewSSLCertificateBindingRepository(db)
	tlsOptionRepo := repository.NewTLSOptionRepository(db)
	svcDiscoveryRepo := repository.NewServiceDiscoveryRepository(db)
	configVersionRepo := repository.NewConfigVersionRepository(db)
	maintWindowRepo := repository.NewMaintenanceWindowRepository(db)
	authRepo := repository.NewAuthRepository(db)

	// Services
	userService := service.NewUserService(userRepo)
	consumerService := service.NewAPIConsumerService(consumerRepo)
	hostService := service.NewHostService(hostRepo)
	vhostService := service.NewVirtualHostService(vhostRepo)
	upstreamService := service.NewUpstreamServerService(upstreamRepo)
	vdirService := service.NewVirtualDirectoryService(vdirRepo, vdirMethodRepo)
	credService := service.NewConsumerCredentialService(credRepo)
	apiKeyService := service.NewAPIKeyService(apiKeyRepo)
	accessService := service.NewRouteConsumerAccessService(accessRepo)
	jwtConfigService := service.NewJWTConfigService(jwtConfigRepo)
	extAuthService := service.NewExternalAuthService(extAuthRepo)
	rateLimitService := service.NewRateLimitService(rateLimitRepo)
	corsConfigService := service.NewCORSConfigService(corsConfigRepo)
	circuitBreakerService := service.NewCircuitBreakerService(circuitBreakerRepo)
	ipWhitelistService := service.NewIPWhitelistService(ipWhitelistRepo)
	ipBlacklistService := service.NewIPBlacklistService(ipBlacklistRepo)
	reqHeaderService := service.NewRequestHeaderRuleService(reqHeaderRepo)
	resHeaderService := service.NewResponseHeaderRuleService(resHeaderRepo)
	queryRewriteService := service.NewQueryRewriteService(queryRewriteRepo)
	acmeService := service.NewACMEAccountService(acmeRepo)
	sslCertService := service.NewSSLCertificateService(sslCertRepo)
	certDomainService := service.NewCertificateDomainService(certDomainRepo)
	sslBindingService := service.NewSSLCertificateBindingService(sslBindingRepo)
	tlsOptionService := service.NewTLSOptionService(tlsOptionRepo)
	svcDiscoveryService := service.NewServiceDiscoveryService(svcDiscoveryRepo)
	configVersionService := service.NewConfigVersionService(configVersionRepo)
	maintWindowService := service.NewMaintenanceWindowService(maintWindowRepo)
	authService := service.NewAuthService(authRepo, "test-jwt-secret", 30, 7)

	// Handlers
	userHandler := handler.NewUserHandler(userService)
	consumerHandler := handler.NewAPIConsumerHandler(consumerService)
	hostHandler := handler.NewHostHandler(hostService)
	vhostHandler := handler.NewVirtualHostHandler(vhostService)
	upstreamHandler := handler.NewUpstreamServerHandler(upstreamService)
	vdirHandler := handler.NewVirtualDirectoryHandler(vdirService)
	credHandler := handler.NewConsumerCredentialHandler(credService)
	apiKeyHandler := handler.NewAPIKeyHandler(apiKeyService)
	accessHandler := handler.NewRouteConsumerAccessHandler(accessService)
	jwtConfigHandler := handler.NewJWTConfigHandler(jwtConfigService)
	extAuthHandler := handler.NewExternalAuthHandler(extAuthService)
	rateLimitHandler := handler.NewRateLimitHandler(rateLimitService)
	corsConfigHandler := handler.NewCORSConfigHandler(corsConfigService)
	circuitBreakerHandler := handler.NewCircuitBreakerHandler(circuitBreakerService)
	ipWhitelistHandler := handler.NewIPWhitelistHandler(ipWhitelistService)
	ipBlacklistHandler := handler.NewIPBlacklistHandler(ipBlacklistService)
	reqHeaderHandler := handler.NewRequestHeaderRuleHandler(reqHeaderService)
	resHeaderHandler := handler.NewResponseHeaderRuleHandler(resHeaderService)
	queryRewriteHandler := handler.NewQueryRewriteHandler(queryRewriteService)
	acmeHandler := handler.NewACMEAccountHandler(acmeService)
	sslCertHandler := handler.NewSSLCertificateHandler(sslCertService)
	certDomainHandler := handler.NewCertificateDomainHandler(certDomainService)
	sslBindingHandler := handler.NewSSLCertificateBindingHandler(sslBindingService)
	tlsOptionHandler := handler.NewTLSOptionHandler(tlsOptionService)
	svcDiscoveryHandler := handler.NewServiceDiscoveryHandler(svcDiscoveryService)
	configVersionHandler := handler.NewConfigVersionHandler(configVersionService)
	maintWindowHandler := handler.NewMaintenanceWindowHandler(maintWindowService)
	authHandler := handler.NewAuthHandler(authService)
	authMiddleware := middleware.NewAuthMiddleware(authService)
	loginRateLimiter := middleware.NewLoginRateLimiter(50, 60)

	// Setup Router
	mux := http.NewServeMux()

	// Health
	mux.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})

	// Users
	mux.HandleFunc("POST /api/admin/users", userHandler.CreateUser)
	mux.HandleFunc("GET /api/admin/users", userHandler.GetAllUsers)
	mux.HandleFunc("GET /api/admin/users/{id}", userHandler.GetUserByID)
	mux.HandleFunc("PUT /api/admin/users/{id}", userHandler.UpdateUser)
	mux.HandleFunc("DELETE /api/admin/users/{id}", userHandler.DeleteUser)

	// Consumers
	mux.HandleFunc("POST /api/admin/consumers", consumerHandler.CreateConsumer)
	mux.HandleFunc("GET /api/admin/consumers", consumerHandler.GetAllConsumers)
	mux.HandleFunc("GET /api/admin/consumers/{id}", consumerHandler.GetConsumerByID)
	mux.HandleFunc("PUT /api/admin/consumers/{id}", consumerHandler.UpdateConsumer)
	mux.HandleFunc("DELETE /api/admin/consumers/{id}", consumerHandler.DeleteConsumer)

	// Hosts
	mux.HandleFunc("POST /api/admin/hosts", hostHandler.CreateHost)
	mux.HandleFunc("GET /api/admin/hosts", hostHandler.GetAllHosts)
	mux.HandleFunc("GET /api/admin/hosts/{id}", hostHandler.GetHostByID)
	mux.HandleFunc("PUT /api/admin/hosts/{id}", hostHandler.UpdateHost)
	mux.HandleFunc("DELETE /api/admin/hosts/{id}", hostHandler.DeleteHost)

	// Virtual Hosts
	mux.HandleFunc("POST /api/admin/virtual-hosts", vhostHandler.CreateVirtualHost)
	mux.HandleFunc("GET /api/admin/virtual-hosts", vhostHandler.GetAllVirtualHosts)
	mux.HandleFunc("GET /api/admin/virtual-hosts/{id}", vhostHandler.GetVirtualHostByID)
	mux.HandleFunc("PUT /api/admin/virtual-hosts/{id}", vhostHandler.UpdateVirtualHost)
	mux.HandleFunc("DELETE /api/admin/virtual-hosts/{id}", vhostHandler.DeleteVirtualHost)

	// Upstream Servers
	mux.HandleFunc("POST /api/admin/upstream-servers", upstreamHandler.CreateUpstreamServer)
	mux.HandleFunc("GET /api/admin/upstream-servers", upstreamHandler.GetAllUpstreamServers)
	mux.HandleFunc("GET /api/admin/upstream-servers/{id}", upstreamHandler.GetUpstreamServerByID)
	mux.HandleFunc("PUT /api/admin/upstream-servers/{id}", upstreamHandler.UpdateUpstreamServer)
	mux.HandleFunc("DELETE /api/admin/upstream-servers/{id}", upstreamHandler.DeleteUpstreamServer)
	mux.HandleFunc("GET /api/admin/virtual-hosts/{vhost_id}/upstream-servers", upstreamHandler.GetUpstreamServersByVHost)

	// Virtual Directories
	mux.HandleFunc("POST /api/admin/virtual-directories", vdirHandler.CreateVirtualDirectory)
	mux.HandleFunc("GET /api/admin/virtual-directories", vdirHandler.GetAllVirtualDirectories)
	mux.HandleFunc("GET /api/admin/virtual-directories/{id}", vdirHandler.GetVirtualDirectoryByID)
	mux.HandleFunc("PUT /api/admin/virtual-directories/{id}", vdirHandler.UpdateVirtualDirectory)
	mux.HandleFunc("DELETE /api/admin/virtual-directories/{id}", vdirHandler.DeleteVirtualDirectory)
	mux.HandleFunc("GET /api/admin/virtual-hosts/{vhost_id}/virtual-directories", vdirHandler.GetVirtualDirectoriesByVHost)
	mux.HandleFunc("GET /api/admin/virtual-directories/{id}/methods", vdirHandler.GetMethods)
	mux.HandleFunc("PUT /api/admin/virtual-directories/{id}/methods", vdirHandler.SetMethods)

	// Consumer Credentials
	mux.HandleFunc("POST /api/admin/consumer-credentials", credHandler.CreateCredential)
	mux.HandleFunc("GET /api/admin/consumer-credentials", credHandler.GetAllCredentials)
	mux.HandleFunc("GET /api/admin/consumer-credentials/{id}", credHandler.GetCredentialByID)
	mux.HandleFunc("PUT /api/admin/consumer-credentials/{id}", credHandler.UpdateCredential)
	mux.HandleFunc("DELETE /api/admin/consumer-credentials/{id}", credHandler.DeleteCredential)
	mux.HandleFunc("GET /api/admin/consumers/{consumer_id}/credentials", credHandler.GetCredentialsByConsumer)

	// API Keys
	mux.HandleFunc("POST /api/admin/api-keys", apiKeyHandler.CreateAPIKey)
	mux.HandleFunc("GET /api/admin/api-keys", apiKeyHandler.GetAllAPIKeys)
	mux.HandleFunc("GET /api/admin/api-keys/{id}", apiKeyHandler.GetAPIKeyByID)
	mux.HandleFunc("PUT /api/admin/api-keys/{id}", apiKeyHandler.UpdateAPIKey)
	mux.HandleFunc("DELETE /api/admin/api-keys/{id}", apiKeyHandler.DeleteAPIKey)
	mux.HandleFunc("GET /api/admin/consumers/{consumer_id}/api-keys", apiKeyHandler.GetAPIKeysByConsumer)

	// Route Consumer Access
	mux.HandleFunc("POST /api/admin/route-access", accessHandler.CreateAccess)
	mux.HandleFunc("GET /api/admin/route-access", accessHandler.GetAllAccess)
	mux.HandleFunc("GET /api/admin/route-access/{id}", accessHandler.GetAccessByID)
	mux.HandleFunc("PUT /api/admin/route-access/{id}", accessHandler.UpdateAccess)
	mux.HandleFunc("DELETE /api/admin/route-access/{id}", accessHandler.DeleteAccess)
	mux.HandleFunc("GET /api/admin/virtual-directories/{dir_id}/access", accessHandler.GetAccessByDirectory)

	// JWT Configs
	mux.HandleFunc("POST /api/admin/jwt-configs", jwtConfigHandler.Create)
	mux.HandleFunc("GET /api/admin/jwt-configs", jwtConfigHandler.GetAll)
	mux.HandleFunc("GET /api/admin/jwt-configs/{id}", jwtConfigHandler.GetByID)
	mux.HandleFunc("PUT /api/admin/jwt-configs/{id}", jwtConfigHandler.Update)
	mux.HandleFunc("DELETE /api/admin/jwt-configs/{id}", jwtConfigHandler.Delete)

	// External Auth
	mux.HandleFunc("POST /api/admin/external-auth", extAuthHandler.Create)
	mux.HandleFunc("GET /api/admin/external-auth", extAuthHandler.GetAll)
	mux.HandleFunc("GET /api/admin/external-auth/{id}", extAuthHandler.GetByID)
	mux.HandleFunc("PUT /api/admin/external-auth/{id}", extAuthHandler.Update)
	mux.HandleFunc("DELETE /api/admin/external-auth/{id}", extAuthHandler.Delete)

	// Rate Limits
	mux.HandleFunc("POST /api/admin/rate-limits", rateLimitHandler.Create)
	mux.HandleFunc("GET /api/admin/rate-limits", rateLimitHandler.GetAll)
	mux.HandleFunc("GET /api/admin/rate-limits/{id}", rateLimitHandler.GetByID)
	mux.HandleFunc("PUT /api/admin/rate-limits/{id}", rateLimitHandler.Update)
	mux.HandleFunc("DELETE /api/admin/rate-limits/{id}", rateLimitHandler.Delete)

	// CORS Configs
	mux.HandleFunc("POST /api/admin/cors-configs", corsConfigHandler.Create)
	mux.HandleFunc("GET /api/admin/cors-configs", corsConfigHandler.GetAll)
	mux.HandleFunc("GET /api/admin/cors-configs/{id}", corsConfigHandler.GetByID)
	mux.HandleFunc("PUT /api/admin/cors-configs/{id}", corsConfigHandler.Update)
	mux.HandleFunc("DELETE /api/admin/cors-configs/{id}", corsConfigHandler.Delete)

	// Circuit Breakers
	mux.HandleFunc("POST /api/admin/circuit-breakers", circuitBreakerHandler.Create)
	mux.HandleFunc("GET /api/admin/circuit-breakers", circuitBreakerHandler.GetAll)
	mux.HandleFunc("GET /api/admin/circuit-breakers/{id}", circuitBreakerHandler.GetByID)
	mux.HandleFunc("PUT /api/admin/circuit-breakers/{id}", circuitBreakerHandler.Update)
	mux.HandleFunc("DELETE /api/admin/circuit-breakers/{id}", circuitBreakerHandler.Delete)

	// IP Whitelists
	mux.HandleFunc("POST /api/admin/ip-whitelists", ipWhitelistHandler.Create)
	mux.HandleFunc("GET /api/admin/ip-whitelists", ipWhitelistHandler.GetAll)
	mux.HandleFunc("GET /api/admin/ip-whitelists/{id}", ipWhitelistHandler.GetByID)
	mux.HandleFunc("PUT /api/admin/ip-whitelists/{id}", ipWhitelistHandler.Update)
	mux.HandleFunc("DELETE /api/admin/ip-whitelists/{id}", ipWhitelistHandler.Delete)
	mux.HandleFunc("GET /api/admin/virtual-directories/{dir_id}/ip-whitelists", ipWhitelistHandler.GetByDirectory)

	// IP Blacklists
	mux.HandleFunc("POST /api/admin/ip-blacklists", ipBlacklistHandler.Create)
	mux.HandleFunc("GET /api/admin/ip-blacklists", ipBlacklistHandler.GetAll)
	mux.HandleFunc("GET /api/admin/ip-blacklists/{id}", ipBlacklistHandler.GetByID)
	mux.HandleFunc("PUT /api/admin/ip-blacklists/{id}", ipBlacklistHandler.Update)
	mux.HandleFunc("DELETE /api/admin/ip-blacklists/{id}", ipBlacklistHandler.Delete)
	mux.HandleFunc("GET /api/admin/virtual-directories/{dir_id}/ip-blacklists", ipBlacklistHandler.GetByDirectory)

	// Request Header Rules
	mux.HandleFunc("POST /api/admin/request-headers", reqHeaderHandler.Create)
	mux.HandleFunc("GET /api/admin/request-headers", reqHeaderHandler.GetAll)
	mux.HandleFunc("GET /api/admin/request-headers/{id}", reqHeaderHandler.GetByID)
	mux.HandleFunc("PUT /api/admin/request-headers/{id}", reqHeaderHandler.Update)
	mux.HandleFunc("DELETE /api/admin/request-headers/{id}", reqHeaderHandler.Delete)
	mux.HandleFunc("GET /api/admin/virtual-directories/{dir_id}/request-headers", reqHeaderHandler.GetByDirectory)

	// Response Header Rules
	mux.HandleFunc("POST /api/admin/response-headers", resHeaderHandler.Create)
	mux.HandleFunc("GET /api/admin/response-headers", resHeaderHandler.GetAll)
	mux.HandleFunc("GET /api/admin/response-headers/{id}", resHeaderHandler.GetByID)
	mux.HandleFunc("PUT /api/admin/response-headers/{id}", resHeaderHandler.Update)
	mux.HandleFunc("DELETE /api/admin/response-headers/{id}", resHeaderHandler.Delete)
	mux.HandleFunc("GET /api/admin/virtual-directories/{dir_id}/response-headers", resHeaderHandler.GetByDirectory)

	// Query Rewrites
	mux.HandleFunc("POST /api/admin/query-rewrites", queryRewriteHandler.Create)
	mux.HandleFunc("GET /api/admin/query-rewrites", queryRewriteHandler.GetAll)
	mux.HandleFunc("GET /api/admin/query-rewrites/{id}", queryRewriteHandler.GetByID)
	mux.HandleFunc("PUT /api/admin/query-rewrites/{id}", queryRewriteHandler.Update)
	mux.HandleFunc("DELETE /api/admin/query-rewrites/{id}", queryRewriteHandler.Delete)
	mux.HandleFunc("GET /api/admin/virtual-directories/{dir_id}/query-rewrites", queryRewriteHandler.GetByDirectory)

	// ACME Accounts
	mux.HandleFunc("POST /api/admin/acme-accounts", acmeHandler.Create)
	mux.HandleFunc("GET /api/admin/acme-accounts", acmeHandler.GetAll)
	mux.HandleFunc("GET /api/admin/acme-accounts/{id}", acmeHandler.GetByID)
	mux.HandleFunc("PUT /api/admin/acme-accounts/{id}", acmeHandler.Update)
	mux.HandleFunc("DELETE /api/admin/acme-accounts/{id}", acmeHandler.Delete)

	// SSL Certificates
	mux.HandleFunc("POST /api/admin/ssl-certificates", sslCertHandler.Create)
	mux.HandleFunc("GET /api/admin/ssl-certificates", sslCertHandler.GetAll)
	mux.HandleFunc("GET /api/admin/ssl-certificates/{id}", sslCertHandler.GetByID)
	mux.HandleFunc("PUT /api/admin/ssl-certificates/{id}", sslCertHandler.Update)
	mux.HandleFunc("DELETE /api/admin/ssl-certificates/{id}", sslCertHandler.Delete)

	// Certificate Domains
	mux.HandleFunc("POST /api/admin/certificate-domains", certDomainHandler.Create)
	mux.HandleFunc("GET /api/admin/certificate-domains", certDomainHandler.GetAll)
	mux.HandleFunc("GET /api/admin/certificate-domains/{id}", certDomainHandler.GetByID)
	mux.HandleFunc("PUT /api/admin/certificate-domains/{id}", certDomainHandler.Update)
	mux.HandleFunc("DELETE /api/admin/certificate-domains/{id}", certDomainHandler.Delete)
	mux.HandleFunc("GET /api/admin/ssl-certificates/{cert_id}/domains", certDomainHandler.GetByCertificate)

	// SSL Certificate Bindings
	mux.HandleFunc("POST /api/admin/ssl-bindings", sslBindingHandler.Create)
	mux.HandleFunc("GET /api/admin/ssl-bindings", sslBindingHandler.GetAll)
	mux.HandleFunc("GET /api/admin/ssl-bindings/{id}", sslBindingHandler.GetByID)
	mux.HandleFunc("PUT /api/admin/ssl-bindings/{id}", sslBindingHandler.Update)
	mux.HandleFunc("DELETE /api/admin/ssl-bindings/{id}", sslBindingHandler.Delete)

	// TLS Options
	mux.HandleFunc("POST /api/admin/tls-options", tlsOptionHandler.Create)
	mux.HandleFunc("GET /api/admin/tls-options", tlsOptionHandler.GetAll)
	mux.HandleFunc("GET /api/admin/tls-options/{id}", tlsOptionHandler.GetByID)
	mux.HandleFunc("PUT /api/admin/tls-options/{id}", tlsOptionHandler.Update)
	mux.HandleFunc("DELETE /api/admin/tls-options/{id}", tlsOptionHandler.Delete)

	// Service Discovery
	mux.HandleFunc("POST /api/admin/service-discovery", svcDiscoveryHandler.Create)
	mux.HandleFunc("GET /api/admin/service-discovery", svcDiscoveryHandler.GetAll)
	mux.HandleFunc("GET /api/admin/service-discovery/{id}", svcDiscoveryHandler.GetByID)
	mux.HandleFunc("PUT /api/admin/service-discovery/{id}", svcDiscoveryHandler.Update)
	mux.HandleFunc("DELETE /api/admin/service-discovery/{id}", svcDiscoveryHandler.Delete)

	// Config Versions
	mux.HandleFunc("POST /api/admin/config-versions", configVersionHandler.Create)
	mux.HandleFunc("GET /api/admin/config-versions", configVersionHandler.GetAll)
	mux.HandleFunc("GET /api/admin/config-versions/{id}", configVersionHandler.GetByID)
	mux.HandleFunc("PUT /api/admin/config-versions/{id}", configVersionHandler.Update)
	mux.HandleFunc("DELETE /api/admin/config-versions/{id}", configVersionHandler.Delete)

	// Maintenance Windows
	mux.HandleFunc("POST /api/admin/maintenance-windows", maintWindowHandler.Create)
	mux.HandleFunc("GET /api/admin/maintenance-windows", maintWindowHandler.GetAll)
	mux.HandleFunc("GET /api/admin/maintenance-windows/{id}", maintWindowHandler.GetByID)
	mux.HandleFunc("PUT /api/admin/maintenance-windows/{id}", maintWindowHandler.Update)
	mux.HandleFunc("DELETE /api/admin/maintenance-windows/{id}", maintWindowHandler.Delete)

	// Auth routes
	mux.Handle("POST /api/admin/auth/login", loginRateLimiter.RateLimit(http.HandlerFunc(authHandler.Login)))
	mux.HandleFunc("POST /api/admin/auth/refresh", authHandler.Refresh)
	mux.HandleFunc("POST /api/admin/auth/logout", authHandler.Logout)
	mux.Handle("GET /api/admin/auth/me", authMiddleware.RequireAuth(http.HandlerFunc(authHandler.Me)))

	// Bungkus dengan middleware
	server := httptest.NewServer(middleware.LoggingMiddleware(middleware.CORSMiddleware(mux)))

	return &TestServer{Server: server, DB: db}
}

// Close menutup server dan database test
func (ts *TestServer) Close() {
	ts.Server.Close()
	ts.DB.Close()
}

// DoRequest melakukan HTTP request ke test server
func (ts *TestServer) DoRequest(method, path string, body interface{}) (*http.Response, []byte) {
	var reqBody io.Reader
	if body != nil {
		jsonBytes, _ := json.Marshal(body)
		reqBody = bytes.NewBuffer(jsonBytes)
	}

	req, _ := http.NewRequest(method, ts.Server.URL+path, reqBody)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, nil
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	return resp, respBody
}

// ParseResponse mem-parse response body ke APIResponse
func ParseResponse(body []byte) APIResponse {
	var res APIResponse
	json.Unmarshal(body, &res)
	return res
}

// ExtractID mengambil ID dari response data
func ExtractID(data map[string]interface{}) int64 {
	if data == nil {
		return 0
	}
	if id, ok := data["id"]; ok {
		return int64(id.(float64))
	}
	return 0
}

// AssertStatus memvalidasi HTTP status code
func AssertStatus(t *testing.T, resp *http.Response, expected int) {
	t.Helper()
	if resp.StatusCode != expected {
		t.Errorf("Expected status %d, got %d", expected, resp.StatusCode)
	}
}

// AssertSuccess memvalidasi bahwa response sukses
func AssertSuccess(t *testing.T, body []byte) APIResponse {
	t.Helper()
	res := ParseResponse(body)
	if !res.Success {
		t.Errorf("Expected success=true, got false. Message: %s", res.Message)
	}
	return res
}

// AssertError memvalidasi bahwa response error
func AssertError(t *testing.T, body []byte) APIResponse {
	t.Helper()
	res := ParseResponse(body)
	if res.Success {
		t.Errorf("Expected success=false, got true. Message: %s", res.Message)
	}
	return res
}

// CreateTestHost membuat host test dan mengembalikan ID-nya
func (ts *TestServer) CreateTestHost(t *testing.T) int64 {
	t.Helper()
	_, body := ts.DoRequest("POST", "/api/admin/hosts", map[string]interface{}{
		"host_name":   fmt.Sprintf("test-%d.example.com", os.Getpid()),
		"description": "Test host",
		"is_active":   true,
	})
	res := ParseResponse(body)
	return ExtractID(res.Data)
}

// CreateTestVirtualHost membuat virtual host test dan mengembalikan ID-nya
func (ts *TestServer) CreateTestVirtualHost(t *testing.T, hostID int64) int64 {
	t.Helper()
	_, body := ts.DoRequest("POST", "/api/admin/virtual-hosts", map[string]interface{}{
		"host_id":    hostID,
		"vhost_name": fmt.Sprintf("vhost-%d.example.com", os.Getpid()),
		"is_active":  true,
	})
	res := ParseResponse(body)
	return ExtractID(res.Data)
}

// CreateTestVirtualDirectory membuat virtual directory test dan mengembalikan ID-nya
func (ts *TestServer) CreateTestVirtualDirectory(t *testing.T, vhostID int64) int64 {
	t.Helper()
	_, body := ts.DoRequest("POST", "/api/admin/virtual-directories", map[string]interface{}{
		"virtual_host_id": vhostID,
		"source_path":     "/api/test",
		"target_path":     "/test",
		"is_active":       true,
	})
	res := ParseResponse(body)
	return ExtractID(res.Data)
}

// CreateTestConsumer membuat consumer test dan mengembalikan ID-nya
func (ts *TestServer) CreateTestConsumer(t *testing.T) int64 {
	t.Helper()
	_, body := ts.DoRequest("POST", "/api/admin/consumers", map[string]interface{}{
		"consumer_name": fmt.Sprintf("test-consumer-%d", os.Getpid()),
		"description":   "Test consumer",
		"is_active":     true,
	})
	res := ParseResponse(body)
	return ExtractID(res.Data)
}

// DoAuthRequest melakukan HTTP request dengan Authorization header
func (ts *TestServer) DoAuthRequest(method, path string, body interface{}, token string) (*http.Response, []byte) {
	var reqBody io.Reader
	if body != nil {
		jsonBytes, _ := json.Marshal(body)
		reqBody = bytes.NewBuffer(jsonBytes)
	}

	req, _ := http.NewRequest(method, ts.Server.URL+path, reqBody)
	req.Header.Set("Content-Type", "application/json")
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, nil
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	return resp, respBody
}

// CreateTestUser membuat user test dan mengembalikan ID-nya
func (ts *TestServer) CreateTestUser(t *testing.T, username, password, role string) int64 {
	t.Helper()
	_, body := ts.DoRequest("POST", "/api/admin/users", map[string]interface{}{
		"username":  username,
		"password":  password,
		"full_name": "Test User",
		"email":     username + "@test.com",
		"role":      role,
		"is_active": true,
	})
	res := ParseResponse(body)
	return ExtractID(res.Data)
}

// LoginTestUser login dan mengembalikan access token + refresh token
func (ts *TestServer) LoginTestUser(t *testing.T, username, password string) (string, string) {
	t.Helper()
	_, body := ts.DoRequest("POST", "/api/admin/auth/login", map[string]interface{}{
		"username": username,
		"password": password,
	})
	res := ParseResponse(body)
	if !res.Success {
		t.Fatalf("Login gagal: %s", res.Message)
	}
	accessToken, _ := res.Data["access_token"].(string)
	refreshToken, _ := res.Data["refresh_token"].(string)
	return accessToken, refreshToken
}
