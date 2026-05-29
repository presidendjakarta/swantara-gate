package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/presidendjakarta/swantara-gate/internal/config"
	"github.com/presidendjakarta/swantara-gate/internal/database"
	"github.com/presidendjakarta/swantara-gate/internal/handler"
	"github.com/presidendjakarta/swantara-gate/internal/middleware"
	"github.com/presidendjakarta/swantara-gate/internal/proxy"
	"github.com/presidendjakarta/swantara-gate/internal/repository"
	"github.com/presidendjakarta/swantara-gate/internal/service"
	"github.com/joho/godotenv"
)

func main() {
	// Memuat environment variables dari file .env
	if err := godotenv.Load(); err != nil {
		log.Println("⚠ File .env tidak ditemukan, menggunakan environment default")
	}

	// Memuat konfigurasi
	cfg := config.LoadConfig()

	log.Println("🚀 Memulai Swantara Gate API Gateway...")

	// Inisialisasi database
	if err := database.InitDatabase(cfg.DatabasePath, cfg.DatabaseSQLPath); err != nil {
		log.Fatalf("❌ Gagal menginisialisasi database: %v", err)
	}

	// Memastikan database ditutup saat aplikasi berhenti
	defer database.CloseDatabase()

	db := database.GetDB()

	// Inisialisasi repositories
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

	// Inisialisasi services
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
	authService := service.NewAuthService(authRepo, cfg.JWTSecret, cfg.JWTAccessExpireMinutes, cfg.JWTRefreshExpireDays)

	// Inisialisasi handlers
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

	// Auth middleware & rate limiter
	authMiddleware := middleware.NewAuthMiddleware(authService)
	loginRateLimiter := middleware.NewLoginRateLimiter(5, 60) // 5 attempts per 60 seconds

	// Inisialisasi Proxy Server (perlu sebelum admin router untuk endpoint reload)
	proxyServer := proxy.NewProxyServer(db)
	proxyServer.Start()
	defer proxyServer.Stop()

	// Setup Admin Router
	adminMux := setupAdminRouter(
		userHandler, consumerHandler, hostHandler, vhostHandler,
		upstreamHandler, vdirHandler, credHandler, apiKeyHandler, accessHandler,
		jwtConfigHandler, extAuthHandler, rateLimitHandler, corsConfigHandler,
		circuitBreakerHandler, ipWhitelistHandler, ipBlacklistHandler,
		reqHeaderHandler, resHeaderHandler, queryRewriteHandler,
		acmeHandler, sslCertHandler, certDomainHandler, sslBindingHandler,
		tlsOptionHandler, svcDiscoveryHandler, configVersionHandler, maintWindowHandler,
		authHandler, authMiddleware, loginRateLimiter,
		proxyServer,
	)

	// Menjalankan Admin HTTP Server
	go func() {
		addr := ":" + toString(cfg.AdminHTTPPort)
		log.Printf("🌐 Admin HTTP Server berjalan di %s", addr)
		if err := http.ListenAndServe(addr, adminMux); err != nil {
			log.Fatalf("❌ Admin HTTP Server error: %v", err)
		}
	}()

	// TODO: Jalankan Admin HTTPS Server (perlu sertifikat SSL)
	// go func() {
	// 	addr := ":" + toString(cfg.AdminHTTPSPort)
	// 	log.Printf("🔒 Admin HTTPS Server berjalan di %s", addr)
	// 	if err := http.ListenAndServeTLS(addr, cfg.AdminSSLCertPath, cfg.AdminSSLKeyPath, adminMux); err != nil {
	// 		log.Fatalf("❌ Admin HTTPS Server error: %v", err)
	// 	}
	// }()

	// Menjalankan Proxy HTTP Server
	go func() {
		addr := ":" + toString(cfg.ProxyHTTPPort)
		log.Printf("🔀 Proxy HTTP Server berjalan di %s", addr)
		if err := http.ListenAndServe(addr, proxyServer); err != nil {
			log.Fatalf("❌ Proxy HTTP Server error: %v", err)
		}
	}()

	// TODO: Jalankan Proxy HTTPS Server (perlu sertifikat SSL - Phase 9)

	log.Println("✅ Swantara Gate API Gateway siap digunakan!")
	log.Println("📍 Admin Panel: http://localhost:" + toString(cfg.AdminHTTPPort))
	log.Println("📍 Proxy Gateway: http://localhost:" + toString(cfg.ProxyHTTPPort))
	
	// Blocking agar program tidak selesai
	select {}
}

// setupAdminRouter mengatur routes untuk Admin API
func setupAdminRouter(
	userHandler *handler.UserHandler,
	consumerHandler *handler.APIConsumerHandler,
	hostHandler *handler.HostHandler,
	vhostHandler *handler.VirtualHostHandler,
	upstreamHandler *handler.UpstreamServerHandler,
	vdirHandler *handler.VirtualDirectoryHandler,
	credHandler *handler.ConsumerCredentialHandler,
	apiKeyHandler *handler.APIKeyHandler,
	accessHandler *handler.RouteConsumerAccessHandler,
	jwtConfigHandler *handler.JWTConfigHandler,
	extAuthHandler *handler.ExternalAuthHandler,
	rateLimitHandler *handler.RateLimitHandler,
	corsConfigHandler *handler.CORSConfigHandler,
	circuitBreakerHandler *handler.CircuitBreakerHandler,
	ipWhitelistHandler *handler.IPWhitelistHandler,
	ipBlacklistHandler *handler.IPBlacklistHandler,
	reqHeaderHandler *handler.RequestHeaderRuleHandler,
	resHeaderHandler *handler.ResponseHeaderRuleHandler,
	queryRewriteHandler *handler.QueryRewriteHandler,
	acmeHandler *handler.ACMEAccountHandler,
	sslCertHandler *handler.SSLCertificateHandler,
	certDomainHandler *handler.CertificateDomainHandler,
	sslBindingHandler *handler.SSLCertificateBindingHandler,
	tlsOptionHandler *handler.TLSOptionHandler,
	svcDiscoveryHandler *handler.ServiceDiscoveryHandler,
	configVersionHandler *handler.ConfigVersionHandler,
	maintWindowHandler *handler.MaintenanceWindowHandler,
	authHandler *handler.AuthHandler,
	authMiddleware *middleware.AuthMiddleware,
	loginRateLimiter *middleware.LoginRateLimiter,
	proxyServer *proxy.ProxyServer,
) http.Handler {
	mux := http.NewServeMux()

	// Health check endpoint
	mux.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok","message":"Swantara Gate Admin API is running"}`))
	})

	// Routes untuk Users
	mux.HandleFunc("POST /api/admin/users", userHandler.CreateUser)
	mux.HandleFunc("GET /api/admin/users", userHandler.GetAllUsers)
	mux.HandleFunc("GET /api/admin/users/{id}", userHandler.GetUserByID)
	mux.HandleFunc("PUT /api/admin/users/{id}", userHandler.UpdateUser)
	mux.HandleFunc("DELETE /api/admin/users/{id}", userHandler.DeleteUser)

	// Routes untuk API Consumers
	mux.HandleFunc("POST /api/admin/consumers", consumerHandler.CreateConsumer)
	mux.HandleFunc("GET /api/admin/consumers", consumerHandler.GetAllConsumers)
	mux.HandleFunc("GET /api/admin/consumers/{id}", consumerHandler.GetConsumerByID)
	mux.HandleFunc("PUT /api/admin/consumers/{id}", consumerHandler.UpdateConsumer)
	mux.HandleFunc("DELETE /api/admin/consumers/{id}", consumerHandler.DeleteConsumer)

	// Routes untuk Hosts
	mux.HandleFunc("POST /api/admin/hosts", hostHandler.CreateHost)
	mux.HandleFunc("GET /api/admin/hosts", hostHandler.GetAllHosts)
	mux.HandleFunc("GET /api/admin/hosts/{id}", hostHandler.GetHostByID)
	mux.HandleFunc("PUT /api/admin/hosts/{id}", hostHandler.UpdateHost)
	mux.HandleFunc("DELETE /api/admin/hosts/{id}", hostHandler.DeleteHost)

	// Routes untuk Virtual Hosts
	mux.HandleFunc("POST /api/admin/virtual-hosts", vhostHandler.CreateVirtualHost)
	mux.HandleFunc("GET /api/admin/virtual-hosts", vhostHandler.GetAllVirtualHosts)
	mux.HandleFunc("GET /api/admin/virtual-hosts/{id}", vhostHandler.GetVirtualHostByID)
	mux.HandleFunc("PUT /api/admin/virtual-hosts/{id}", vhostHandler.UpdateVirtualHost)
	mux.HandleFunc("DELETE /api/admin/virtual-hosts/{id}", vhostHandler.DeleteVirtualHost)

	// Routes untuk Upstream Servers
	mux.HandleFunc("POST /api/admin/upstream-servers", upstreamHandler.CreateUpstreamServer)
	mux.HandleFunc("GET /api/admin/upstream-servers", upstreamHandler.GetAllUpstreamServers)
	mux.HandleFunc("GET /api/admin/upstream-servers/{id}", upstreamHandler.GetUpstreamServerByID)
	mux.HandleFunc("PUT /api/admin/upstream-servers/{id}", upstreamHandler.UpdateUpstreamServer)
	mux.HandleFunc("DELETE /api/admin/upstream-servers/{id}", upstreamHandler.DeleteUpstreamServer)
	mux.HandleFunc("GET /api/admin/virtual-hosts/{vhost_id}/upstream-servers", upstreamHandler.GetUpstreamServersByVHost)

	// Routes untuk Virtual Directories
	mux.HandleFunc("POST /api/admin/virtual-directories", vdirHandler.CreateVirtualDirectory)
	mux.HandleFunc("GET /api/admin/virtual-directories", vdirHandler.GetAllVirtualDirectories)
	mux.HandleFunc("GET /api/admin/virtual-directories/{id}", vdirHandler.GetVirtualDirectoryByID)
	mux.HandleFunc("PUT /api/admin/virtual-directories/{id}", vdirHandler.UpdateVirtualDirectory)
	mux.HandleFunc("DELETE /api/admin/virtual-directories/{id}", vdirHandler.DeleteVirtualDirectory)
	mux.HandleFunc("GET /api/admin/virtual-hosts/{vhost_id}/virtual-directories", vdirHandler.GetVirtualDirectoriesByVHost)

	// Routes untuk Virtual Directory Methods
	mux.HandleFunc("GET /api/admin/virtual-directories/{id}/methods", vdirHandler.GetMethods)
	mux.HandleFunc("PUT /api/admin/virtual-directories/{id}/methods", vdirHandler.SetMethods)

	// Routes untuk Consumer Credentials
	mux.HandleFunc("POST /api/admin/consumer-credentials", credHandler.CreateCredential)
	mux.HandleFunc("GET /api/admin/consumer-credentials", credHandler.GetAllCredentials)
	mux.HandleFunc("GET /api/admin/consumer-credentials/{id}", credHandler.GetCredentialByID)
	mux.HandleFunc("PUT /api/admin/consumer-credentials/{id}", credHandler.UpdateCredential)
	mux.HandleFunc("DELETE /api/admin/consumer-credentials/{id}", credHandler.DeleteCredential)
	mux.HandleFunc("GET /api/admin/consumers/{consumer_id}/credentials", credHandler.GetCredentialsByConsumer)

	// Routes untuk API Keys
	mux.HandleFunc("POST /api/admin/api-keys", apiKeyHandler.CreateAPIKey)
	mux.HandleFunc("GET /api/admin/api-keys", apiKeyHandler.GetAllAPIKeys)
	mux.HandleFunc("GET /api/admin/api-keys/{id}", apiKeyHandler.GetAPIKeyByID)
	mux.HandleFunc("PUT /api/admin/api-keys/{id}", apiKeyHandler.UpdateAPIKey)
	mux.HandleFunc("DELETE /api/admin/api-keys/{id}", apiKeyHandler.DeleteAPIKey)
	mux.HandleFunc("GET /api/admin/consumers/{consumer_id}/api-keys", apiKeyHandler.GetAPIKeysByConsumer)

	// Routes untuk Route Consumer Access (ACL)
	mux.HandleFunc("POST /api/admin/route-access", accessHandler.CreateAccess)
	mux.HandleFunc("GET /api/admin/route-access", accessHandler.GetAllAccess)
	mux.HandleFunc("GET /api/admin/route-access/{id}", accessHandler.GetAccessByID)
	mux.HandleFunc("PUT /api/admin/route-access/{id}", accessHandler.UpdateAccess)
	mux.HandleFunc("DELETE /api/admin/route-access/{id}", accessHandler.DeleteAccess)
	mux.HandleFunc("GET /api/admin/virtual-directories/{dir_id}/access", accessHandler.GetAccessByDirectory)

	// Routes untuk JWT Configs
	mux.HandleFunc("POST /api/admin/jwt-configs", jwtConfigHandler.Create)
	mux.HandleFunc("GET /api/admin/jwt-configs", jwtConfigHandler.GetAll)
	mux.HandleFunc("GET /api/admin/jwt-configs/{id}", jwtConfigHandler.GetByID)
	mux.HandleFunc("PUT /api/admin/jwt-configs/{id}", jwtConfigHandler.Update)
	mux.HandleFunc("DELETE /api/admin/jwt-configs/{id}", jwtConfigHandler.Delete)

	// Routes untuk External Auth
	mux.HandleFunc("POST /api/admin/external-auth", extAuthHandler.Create)
	mux.HandleFunc("GET /api/admin/external-auth", extAuthHandler.GetAll)
	mux.HandleFunc("GET /api/admin/external-auth/{id}", extAuthHandler.GetByID)
	mux.HandleFunc("PUT /api/admin/external-auth/{id}", extAuthHandler.Update)
	mux.HandleFunc("DELETE /api/admin/external-auth/{id}", extAuthHandler.Delete)

	// Routes untuk Rate Limits
	mux.HandleFunc("POST /api/admin/rate-limits", rateLimitHandler.Create)
	mux.HandleFunc("GET /api/admin/rate-limits", rateLimitHandler.GetAll)
	mux.HandleFunc("GET /api/admin/rate-limits/{id}", rateLimitHandler.GetByID)
	mux.HandleFunc("PUT /api/admin/rate-limits/{id}", rateLimitHandler.Update)
	mux.HandleFunc("DELETE /api/admin/rate-limits/{id}", rateLimitHandler.Delete)

	// Routes untuk CORS Configs
	mux.HandleFunc("POST /api/admin/cors-configs", corsConfigHandler.Create)
	mux.HandleFunc("GET /api/admin/cors-configs", corsConfigHandler.GetAll)
	mux.HandleFunc("GET /api/admin/cors-configs/{id}", corsConfigHandler.GetByID)
	mux.HandleFunc("PUT /api/admin/cors-configs/{id}", corsConfigHandler.Update)
	mux.HandleFunc("DELETE /api/admin/cors-configs/{id}", corsConfigHandler.Delete)

	// Routes untuk Circuit Breakers
	mux.HandleFunc("POST /api/admin/circuit-breakers", circuitBreakerHandler.Create)
	mux.HandleFunc("GET /api/admin/circuit-breakers", circuitBreakerHandler.GetAll)
	mux.HandleFunc("GET /api/admin/circuit-breakers/{id}", circuitBreakerHandler.GetByID)
	mux.HandleFunc("PUT /api/admin/circuit-breakers/{id}", circuitBreakerHandler.Update)
	mux.HandleFunc("DELETE /api/admin/circuit-breakers/{id}", circuitBreakerHandler.Delete)

	// Routes untuk IP Whitelists
	mux.HandleFunc("POST /api/admin/ip-whitelists", ipWhitelistHandler.Create)
	mux.HandleFunc("GET /api/admin/ip-whitelists", ipWhitelistHandler.GetAll)
	mux.HandleFunc("GET /api/admin/ip-whitelists/{id}", ipWhitelistHandler.GetByID)
	mux.HandleFunc("PUT /api/admin/ip-whitelists/{id}", ipWhitelistHandler.Update)
	mux.HandleFunc("DELETE /api/admin/ip-whitelists/{id}", ipWhitelistHandler.Delete)
	mux.HandleFunc("GET /api/admin/virtual-directories/{dir_id}/ip-whitelists", ipWhitelistHandler.GetByDirectory)

	// Routes untuk IP Blacklists
	mux.HandleFunc("POST /api/admin/ip-blacklists", ipBlacklistHandler.Create)
	mux.HandleFunc("GET /api/admin/ip-blacklists", ipBlacklistHandler.GetAll)
	mux.HandleFunc("GET /api/admin/ip-blacklists/{id}", ipBlacklistHandler.GetByID)
	mux.HandleFunc("PUT /api/admin/ip-blacklists/{id}", ipBlacklistHandler.Update)
	mux.HandleFunc("DELETE /api/admin/ip-blacklists/{id}", ipBlacklistHandler.Delete)
	mux.HandleFunc("GET /api/admin/virtual-directories/{dir_id}/ip-blacklists", ipBlacklistHandler.GetByDirectory)

	// Routes untuk Request Header Rules
	mux.HandleFunc("POST /api/admin/request-headers", reqHeaderHandler.Create)
	mux.HandleFunc("GET /api/admin/request-headers", reqHeaderHandler.GetAll)
	mux.HandleFunc("GET /api/admin/request-headers/{id}", reqHeaderHandler.GetByID)
	mux.HandleFunc("PUT /api/admin/request-headers/{id}", reqHeaderHandler.Update)
	mux.HandleFunc("DELETE /api/admin/request-headers/{id}", reqHeaderHandler.Delete)
	mux.HandleFunc("GET /api/admin/virtual-directories/{dir_id}/request-headers", reqHeaderHandler.GetByDirectory)

	// Routes untuk Response Header Rules
	mux.HandleFunc("POST /api/admin/response-headers", resHeaderHandler.Create)
	mux.HandleFunc("GET /api/admin/response-headers", resHeaderHandler.GetAll)
	mux.HandleFunc("GET /api/admin/response-headers/{id}", resHeaderHandler.GetByID)
	mux.HandleFunc("PUT /api/admin/response-headers/{id}", resHeaderHandler.Update)
	mux.HandleFunc("DELETE /api/admin/response-headers/{id}", resHeaderHandler.Delete)
	mux.HandleFunc("GET /api/admin/virtual-directories/{dir_id}/response-headers", resHeaderHandler.GetByDirectory)

	// Routes untuk Query Rewrites
	mux.HandleFunc("POST /api/admin/query-rewrites", queryRewriteHandler.Create)
	mux.HandleFunc("GET /api/admin/query-rewrites", queryRewriteHandler.GetAll)
	mux.HandleFunc("GET /api/admin/query-rewrites/{id}", queryRewriteHandler.GetByID)
	mux.HandleFunc("PUT /api/admin/query-rewrites/{id}", queryRewriteHandler.Update)
	mux.HandleFunc("DELETE /api/admin/query-rewrites/{id}", queryRewriteHandler.Delete)
	mux.HandleFunc("GET /api/admin/virtual-directories/{dir_id}/query-rewrites", queryRewriteHandler.GetByDirectory)

	// Routes untuk ACME Accounts
	mux.HandleFunc("POST /api/admin/acme-accounts", acmeHandler.Create)
	mux.HandleFunc("GET /api/admin/acme-accounts", acmeHandler.GetAll)
	mux.HandleFunc("GET /api/admin/acme-accounts/{id}", acmeHandler.GetByID)
	mux.HandleFunc("PUT /api/admin/acme-accounts/{id}", acmeHandler.Update)
	mux.HandleFunc("DELETE /api/admin/acme-accounts/{id}", acmeHandler.Delete)

	// Routes untuk SSL Certificates
	mux.HandleFunc("POST /api/admin/ssl-certificates", sslCertHandler.Create)
	mux.HandleFunc("GET /api/admin/ssl-certificates", sslCertHandler.GetAll)
	mux.HandleFunc("GET /api/admin/ssl-certificates/{id}", sslCertHandler.GetByID)
	mux.HandleFunc("PUT /api/admin/ssl-certificates/{id}", sslCertHandler.Update)
	mux.HandleFunc("DELETE /api/admin/ssl-certificates/{id}", sslCertHandler.Delete)

	// Routes untuk Certificate Domains
	mux.HandleFunc("POST /api/admin/certificate-domains", certDomainHandler.Create)
	mux.HandleFunc("GET /api/admin/certificate-domains", certDomainHandler.GetAll)
	mux.HandleFunc("GET /api/admin/certificate-domains/{id}", certDomainHandler.GetByID)
	mux.HandleFunc("PUT /api/admin/certificate-domains/{id}", certDomainHandler.Update)
	mux.HandleFunc("DELETE /api/admin/certificate-domains/{id}", certDomainHandler.Delete)
	mux.HandleFunc("GET /api/admin/ssl-certificates/{cert_id}/domains", certDomainHandler.GetByCertificate)

	// Routes untuk SSL Certificate Bindings
	mux.HandleFunc("POST /api/admin/ssl-bindings", sslBindingHandler.Create)
	mux.HandleFunc("GET /api/admin/ssl-bindings", sslBindingHandler.GetAll)
	mux.HandleFunc("GET /api/admin/ssl-bindings/{id}", sslBindingHandler.GetByID)
	mux.HandleFunc("PUT /api/admin/ssl-bindings/{id}", sslBindingHandler.Update)
	mux.HandleFunc("DELETE /api/admin/ssl-bindings/{id}", sslBindingHandler.Delete)

	// Routes untuk TLS Options
	mux.HandleFunc("POST /api/admin/tls-options", tlsOptionHandler.Create)
	mux.HandleFunc("GET /api/admin/tls-options", tlsOptionHandler.GetAll)
	mux.HandleFunc("GET /api/admin/tls-options/{id}", tlsOptionHandler.GetByID)
	mux.HandleFunc("PUT /api/admin/tls-options/{id}", tlsOptionHandler.Update)
	mux.HandleFunc("DELETE /api/admin/tls-options/{id}", tlsOptionHandler.Delete)

	// Routes untuk Service Discovery
	mux.HandleFunc("POST /api/admin/service-discovery", svcDiscoveryHandler.Create)
	mux.HandleFunc("GET /api/admin/service-discovery", svcDiscoveryHandler.GetAll)
	mux.HandleFunc("GET /api/admin/service-discovery/{id}", svcDiscoveryHandler.GetByID)
	mux.HandleFunc("PUT /api/admin/service-discovery/{id}", svcDiscoveryHandler.Update)
	mux.HandleFunc("DELETE /api/admin/service-discovery/{id}", svcDiscoveryHandler.Delete)

	// Routes untuk Config Versions
	mux.HandleFunc("POST /api/admin/config-versions", configVersionHandler.Create)
	mux.HandleFunc("GET /api/admin/config-versions", configVersionHandler.GetAll)
	mux.HandleFunc("GET /api/admin/config-versions/{id}", configVersionHandler.GetByID)
	mux.HandleFunc("PUT /api/admin/config-versions/{id}", configVersionHandler.Update)
	mux.HandleFunc("DELETE /api/admin/config-versions/{id}", configVersionHandler.Delete)

	// Routes untuk Maintenance Windows
	mux.HandleFunc("POST /api/admin/maintenance-windows", maintWindowHandler.Create)
	mux.HandleFunc("GET /api/admin/maintenance-windows", maintWindowHandler.GetAll)
	mux.HandleFunc("GET /api/admin/maintenance-windows/{id}", maintWindowHandler.GetByID)
	mux.HandleFunc("PUT /api/admin/maintenance-windows/{id}", maintWindowHandler.Update)
	mux.HandleFunc("DELETE /api/admin/maintenance-windows/{id}", maintWindowHandler.Delete)

	// Reload proxy configuration (routes, upstreams, security, etc.)
	mux.HandleFunc("POST /api/admin/config/reload", func(w http.ResponseWriter, r *http.Request) {
		if proxyServer == nil {
			http.Error(w, `{"error":"proxy server not initialized"}`, http.StatusServiceUnavailable)
			return
		}
		log.Println("[Admin] Reloading proxy configuration...")
		proxyServer.ReloadConfig()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok","message":"Proxy configuration reloaded successfully"}`))
	})

	// Routes untuk Auth (public - tidak perlu token)
	mux.Handle("POST /api/admin/auth/login", loginRateLimiter.RateLimit(http.HandlerFunc(authHandler.Login)))
	mux.HandleFunc("POST /api/admin/auth/refresh", authHandler.Refresh)
	mux.HandleFunc("POST /api/admin/auth/logout", authHandler.Logout)

	// Routes yang dilindungi auth middleware
	mux.Handle("GET /api/admin/auth/me", authMiddleware.RequireAuth(http.HandlerFunc(authHandler.Me)))

	// Wrap semua route dengan auth middleware kecuali public routes
	protectedHandler := authMiddleware.RequireAuthExcept(
		mux,
		"/api/health",
		"/api/admin/auth/login",
		"/api/admin/auth/refresh",
		"/api/admin/auth/logout",
	)

	return middleware.LoggingMiddleware(middleware.CORSMiddleware(protectedHandler))
}

// toString mengubah int ke string
func toString(i int) string {
	return strconv.Itoa(i)
}
