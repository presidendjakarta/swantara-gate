# 📋 TODO - Swantara Gate API Gateway

## Ringkasan Proyek
Membangun API Gateway Proxy dengan Go + SQLite.
**Strategi:** API management dulu → validasi semua → baru Web GUI.

---

## ✅ FASE 1 - Foundation (SELESAI)

- [x] Setup project structure (clean architecture)
- [x] Konfigurasi .env (multi-port: admin HTTP/HTTPS, proxy HTTP/HTTPS)
- [x] Koneksi database SQLite (pure Go, tanpa CGO)
- [x] Smart migration (cek tabel sebelum migrasi)
- [x] Response helper (standardisasi JSON response)
- [x] Middleware (CORS, Logging)
- [x] CRUD Users (model, repository, service, handler)
- [x] CRUD API Consumers
- [x] CRUD Hosts
- [x] CRUD Virtual Hosts
- [x] Air hot reload untuk development

---

## 🔧 FASE 2 - Admin API (Lengkapi CRUD Semua Entity)

### 2.1 Upstream Servers
- [ ] Model upstream_server.go
- [ ] Repository upstream_server_repository.go
- [ ] Service upstream_server_service.go
- [ ] Handler + Routes (POST/GET/PUT/DELETE /api/admin/upstream-servers)
- [ ] Test di Postman

### 2.2 Virtual Directories (API Routes)
- [ ] Model virtual_directory.go
- [ ] Repository virtual_directory_repository.go
- [ ] Service virtual_directory_service.go
- [ ] Handler + Routes (POST/GET/PUT/DELETE /api/admin/virtual-directories)
- [ ] Test di Postman

### 2.3 Virtual Directory Methods
- [ ] Model virtual_directory_method.go
- [ ] Repository + Service + Handler
- [ ] Routes (CRUD /api/admin/virtual-directories/{id}/methods)
- [ ] Test di Postman

### 2.4 Consumer Credentials
- [ ] Model consumer_credential.go
- [ ] Repository + Service + Handler
- [ ] Routes (CRUD /api/admin/consumers/{id}/credentials)
- [ ] Test di Postman

### 2.5 API Keys
- [ ] Model api_key.go (generate key otomatis)
- [ ] Repository + Service + Handler
- [ ] Routes (CRUD /api/admin/api-keys)
- [ ] Test di Postman

### 2.6 Route Consumer Access (ACL)
- [ ] Model route_consumer_access.go
- [ ] Repository + Service + Handler
- [ ] Routes (CRUD /api/admin/route-access)
- [ ] Test di Postman

---

## 🔐 FASE 3 - Admin API (Keamanan & Proteksi)

### 3.1 JWT Configs
- [ ] Model jwt_config.go
- [ ] Repository + Service + Handler
- [ ] Routes (CRUD /api/admin/jwt-configs)
- [ ] Test di Postman

### 3.2 External Auth
- [ ] Model external_auth.go
- [ ] Repository + Service + Handler
- [ ] Routes (CRUD /api/admin/external-auth)
- [ ] Test di Postman

### 3.3 Rate Limits
- [ ] Model rate_limit.go
- [ ] Repository + Service + Handler
- [ ] Routes (CRUD /api/admin/rate-limits)
- [ ] Test di Postman

### 3.4 CORS Configs
- [ ] Model cors_config.go
- [ ] Repository + Service + Handler
- [ ] Routes (CRUD /api/admin/cors-configs)
- [ ] Test di Postman

### 3.5 Circuit Breakers
- [ ] Model circuit_breaker.go
- [ ] Repository + Service + Handler
- [ ] Routes (CRUD /api/admin/circuit-breakers)
- [ ] Test di Postman

### 3.6 IP Whitelists
- [ ] Model ip_whitelist.go
- [ ] Repository + Service + Handler
- [ ] Routes (CRUD /api/admin/ip-whitelists)
- [ ] Test di Postman

### 3.7 IP Blacklists
- [ ] Model ip_blacklist.go
- [ ] Repository + Service + Handler
- [ ] Routes (CRUD /api/admin/ip-blacklists)
- [ ] Test di Postman

---

## 🔀 FASE 4 - Admin API (Header & Rewrite)

### 4.1 Request Header Rules
- [ ] Model request_header_rule.go
- [ ] Repository + Service + Handler
- [ ] Routes (CRUD /api/admin/request-headers)
- [ ] Test di Postman

### 4.2 Response Header Rules
- [ ] Model response_header_rule.go
- [ ] Repository + Service + Handler
- [ ] Routes (CRUD /api/admin/response-headers)
- [ ] Test di Postman

### 4.3 Query Rewrites
- [ ] Model query_rewrite.go
- [ ] Repository + Service + Handler
- [ ] Routes (CRUD /api/admin/query-rewrites)
- [ ] Test di Postman

---

## 🔒 FASE 5 - Admin API (SSL/TLS & Maintenance)

### 5.1 ACME Accounts
- [ ] Model acme_account.go
- [ ] Repository + Service + Handler
- [ ] Routes (CRUD /api/admin/acme-accounts)
- [ ] Test di Postman

### 5.2 SSL Certificates
- [ ] Model ssl_certificate.go
- [ ] Repository + Service + Handler
- [ ] Routes (CRUD /api/admin/ssl-certificates)
- [ ] Test di Postman

### 5.3 Certificate Domains
- [ ] Model certificate_domain.go
- [ ] Repository + Service + Handler
- [ ] Routes (CRUD /api/admin/certificate-domains)
- [ ] Test di Postman

### 5.4 SSL Certificate Bindings
- [ ] Model ssl_certificate_binding.go
- [ ] Repository + Service + Handler
- [ ] Routes (CRUD /api/admin/ssl-bindings)
- [ ] Test di Postman

### 5.5 TLS Options
- [ ] Model tls_option.go
- [ ] Repository + Service + Handler
- [ ] Routes (CRUD /api/admin/tls-options)
- [ ] Test di Postman

### 5.6 Service Discovery
- [ ] Model service_discovery.go
- [ ] Repository + Service + Handler
- [ ] Routes (CRUD /api/admin/service-discovery)
- [ ] Test di Postman

### 5.7 Maintenance Windows
- [ ] Model maintenance_window.go
- [ ] Repository + Service + Handler
- [ ] Routes (CRUD /api/admin/maintenance-windows)
- [ ] Test di Postman

### 5.8 Config Versions
- [ ] Model config_version.go
- [ ] Repository + Service + Handler
- [ ] Routes (CRUD /api/admin/config-versions)
- [ ] Test di Postman

---

## 🛡️ FASE 6 - Admin Auth & Security

- [ ] Login endpoint (POST /api/admin/auth/login)
- [ ] JWT token untuk admin panel
- [ ] Auth middleware (proteksi semua /api/admin/*)
- [ ] Refresh token
- [ ] Logout / revoke token
- [ ] Role-based access control (RBAC) di middleware
- [ ] Rate limiting untuk login endpoint

---

## 🚀 FASE 7 - Proxy Server (Core Engine)

- [ ] Reverse proxy handler
- [ ] Route matching (exact, prefix, wildcard, regex, parameter)
- [ ] Load balancing implementation (round_robin, weighted, least_conn, ip_hash, random, failover)
- [ ] Sticky session support
- [ ] Health check worker (background goroutine)
- [ ] Failover logic (active-active, active-passive)
- [ ] Request forwarding ke upstream servers
- [ ] Header manipulation (request & response)
- [ ] Query parameter rewrite
- [ ] WebSocket proxy support
- [ ] Response caching
- [ ] Max request size enforcement
- [ ] Proxy timeout handling
- [ ] Retry logic

---

## 🛡️ FASE 8 - Proxy Security & Features

- [ ] API Key authentication di proxy
- [ ] JWT validation di proxy
- [ ] Basic auth validation di proxy
- [ ] External auth (forward auth)
- [ ] Rate limiting enforcement
- [ ] Circuit breaker implementation
- [ ] IP whitelist/blacklist enforcement
- [ ] CORS handling per route
- [ ] Maintenance mode enforcement
- [ ] Request/response logging (access log)

---

## 🔒 FASE 9 - SSL/TLS & HTTPS

- [ ] Admin HTTPS server
- [ ] Proxy HTTPS server
- [ ] Let's Encrypt ACME integration
- [ ] Auto certificate renewal
- [ ] TLS version enforcement
- [ ] HTTP/2 support
- [ ] HSTS header
- [ ] HTTP → HTTPS redirect

---

## 🖥️ FASE 10 - Web GUI (Setelah API Valid)

- [ ] Pilih framework frontend (React/Vue/Svelte)
- [ ] Dashboard overview (statistik gateway)
- [ ] Halaman manajemen Users
- [ ] Halaman manajemen Consumers & API Keys
- [ ] Halaman manajemen Hosts & Virtual Hosts
- [ ] Halaman manajemen Routes (Virtual Directories)
- [ ] Halaman manajemen Upstream Servers
- [ ] Halaman konfigurasi Security (Rate limit, CORS, IP filter)
- [ ] Halaman konfigurasi SSL/TLS
- [ ] Halaman monitoring & health check
- [ ] Halaman access logs
- [ ] Halaman maintenance mode
- [ ] Login page & session management

---

## 📊 FASE 11 - Monitoring & Logging

- [ ] Access log table di database
- [ ] Endpoint statistik (request count, latency, error rate)
- [ ] Health check status dashboard data
- [ ] Alert system (email/webhook saat server down)
- [ ] Log rotation

---

## 🧪 FASE 12 - Testing & Production Ready

- [ ] Unit test untuk semua service
- [ ] Integration test untuk API
- [ ] Load testing
- [ ] Build script (Makefile/script)
- [ ] Docker support
- [ ] CI/CD pipeline
- [ ] Production deployment guide

---

## 📌 Status Saat Ini

**Posisi:** Awal FASE 2 - Melengkapi CRUD semua entity
**Yang jalan:** Admin HTTP server di port 8081
**Yang sudah bisa di-test:** Users, Consumers, Hosts, Virtual Hosts

## 🎯 Next Action

→ Lanjut FASE 2.1: Buat CRUD Upstream Servers
