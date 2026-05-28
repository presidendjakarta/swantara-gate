# Dokumentasi API Swantara Gate

## Informasi Umum

| Item | Nilai |
|------|-------|
| **Nama** | Swantara Gate API Gateway |
| **Base URL (Admin)** | `http://localhost:9090` |
| **Base URL (Proxy)** | `http://localhost:8000` |
| **Autentikasi** | Bearer Token (JWT) |
| **Format** | JSON |

---

## Autentikasi

Semua endpoint admin memerlukan token JWT yang dikirim melalui header `Authorization`:

```
Authorization: Bearer <access_token>
```

Token bisa didapat dari endpoint **Login**. Token otomatis tersimpan di collection variable Postman setelah login berhasil.

### 1. Login

Autentikasi user dan dapatkan access token + refresh token.

```
POST {{base_url}}/api/admin/auth/login
```

**Request Body:**
```json
{
    "username": "admin",
    "password": "admin123"
}
```

**Response:**
```json
{
    "status": "success",
    "message": "Login berhasil",
    "data": {
        "access_token": "eyJhbGci...",
        "refresh_token": "9638eb...",
        "token_type": "Bearer",
        "expires_in": 1800
    }
}
```

---

### 2. Refresh Token

Perbarui access token menggunakan refresh token.

```
POST {{base_url}}/api/admin/auth/refresh
```

**Request Body:**
```json
{
    "refresh_token": "{{refresh_token}}"
}
```

**Response:** Token pair baru (access token + refresh token baru).

---

### 3. Logout (Refresh Token Only)

Logout hanya dengan revoke refresh token di database. Access token masih valid sampai expired.

```
POST {{base_url}}/api/admin/auth/logout
Content-Type: application/json
```

**Request Body:**
```json
{
    "refresh_token": "{{refresh_token}}"
}
```

---

### 4. Logout (Access Token + Refresh Token) ⭐ **Recommended**

Logout dengan **blacklist access token** DAN revoke refresh token. Access token langsung tidak bisa dipakai lagi.

```
POST {{base_url}}/api/admin/auth/logout
Authorization: Bearer {{access_token}}
Content-Type: application/json
```

**Request Body:**
```json
{
    "refresh_token": "{{refresh_token}}"
}
```

**cURL Example:**
```bash
curl --location 'http://localhost:8081/api/admin/auth/logout' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGci...' \
--data '{
    "refresh_token": "9638ebc58f4aa97e1980903af2e8f634fce8629cc1d5289dd2e9aec778086178"
}'
```

---

### 5. Get Current User (Me)

Ambil profil user yang sedang login.

```
GET {{base_url}}/api/admin/auth/me
Authorization: Bearer {{access_token}}
```

**Response:**
```json
{
    "status": "success",
    "message": "Profil user",
    "data": {
        "user_id": 1,
        "username": "admin",
        "role": "super_admin"
    }
}
```

---

## Health Check

```
GET {{base_url}}/api/health
```

**Response:**
```json
{
    "status": "ok",
    "message": "Swantara Gate Admin API is running"
}
```

---

## Users

Manajemen user admin panel.

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| POST | `/api/admin/users` | Buat user baru |
| GET | `/api/admin/users` | List semua user |
| GET | `/api/admin/users/{id}` | Detail user |
| PUT | `/api/admin/users/{id}` | Update user |
| DELETE | `/api/admin/users/{id}` | Hapus user |

### Create User

**Request Body:**
```json
{
    "username": "newuser",
    "password": "password123",
    "full_name": "New User",
    "email": "newuser@example.com",
    "role": "admin",
    "is_active": true
}
```

Role yang valid: `super_admin`, `admin`, `operator`, `viewer`

### Update User

**Request Body:**
```json
{
    "full_name": "Updated Name",
    "email": "updated@example.com",
    "role": "admin",
    "is_active": true
}
```

---

## API Consumers

Manajemen konsumen API (aplikasi yang menggunakan gateway).

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| POST | `/api/admin/consumers` | Buat consumer |
| GET | `/api/admin/consumers` | List semua consumer |
| GET | `/api/admin/consumers/{id}` | Detail consumer |
| PUT | `/api/admin/consumers/{id}` | Update consumer |
| DELETE | `/api/admin/consumers/{id}` | Hapus consumer |

### Create Consumer

**Request Body:**
```json
{
    "consumer_name": "my-app",
    "description": "My Application",
    "contact_email": "dev@myapp.com",
    "is_active": true
}
```

---

## Hosts

Manajemen host fisik/server.

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| POST | `/api/admin/hosts` | Buat host |
| GET | `/api/admin/hosts` | List semua host |
| GET | `/api/admin/hosts/{id}` | Detail host |
| PUT | `/api/admin/hosts/{id}` | Update host |
| DELETE | `/api/admin/hosts/{id}` | Hapus host |

### Create Host

**Request Body:**
```json
{
    "host_name": "api.example.com",
    "description": "Main API Host",
    "is_active": true
}
```

---

## Virtual Hosts

Manajemen virtual host (konfigurasi routing per domain).

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| POST | `/api/admin/virtual-hosts` | Buat virtual host |
| GET | `/api/admin/virtual-hosts` | List semua virtual host |
| GET | `/api/admin/virtual-hosts/{id}` | Detail virtual host |
| PUT | `/api/admin/virtual-hosts/{id}` | Update virtual host |
| DELETE | `/api/admin/virtual-hosts/{id}` | Hapus virtual host |

### Create Virtual Host

**Request Body:**
```json
{
    "host_id": 1,
    "vhost_name": "api.example.com",
    "lb_algorithm": "round_robin",
    "sticky_session": false,
    "failover_mode": "active-active",
    "is_active": true
}
```

lb_algorithm yang valid: `round_robin`, `weighted_round_robin`, `least_connections`, `ip_hash`

---

## Upstream Servers

Manajemen server upstream (backend yang diproxy).

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| POST | `/api/admin/upstream-servers` | Buat upstream |
| GET | `/api/admin/upstream-servers` | List semua upstream |
| GET | `/api/admin/upstream-servers/{id}` | Detail upstream |
| GET | `/api/admin/virtual-hosts/{vhost_id}/upstream-servers` | Upstream by VHost |
| PUT | `/api/admin/upstream-servers/{id}` | Update upstream |
| DELETE | `/api/admin/upstream-servers/{id}` | Hapus upstream |

### Create Upstream Server

**Request Body:**
```json
{
    "virtual_host_id": 1,
    "target_host": "192.168.1.10",
    "target_port": 8080,
    "protocol": "http",
    "priority": 1,
    "weight": 1,
    "is_backup": false,
    "is_active": true,
    "health_check_enabled": true,
    "health_check_path": "/health",
    "health_check_interval_seconds": 10,
    "health_check_timeout_seconds": 3,
    "max_fails": 3,
    "fail_timeout_seconds": 30
}
```

---

## Virtual Directories (Routes)

Manajemen route/path yang di-proxy.

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| POST | `/api/admin/virtual-directories` | Buat route |
| GET | `/api/admin/virtual-directories` | List semua route |
| GET | `/api/admin/virtual-directories/{id}` | Detail route |
| GET | `/api/admin/virtual-hosts/{vhost_id}/virtual-directories` | Route by VHost |
| PUT | `/api/admin/virtual-directories/{id}` | Update route |
| DELETE | `/api/admin/virtual-directories/{id}` | Hapus route |
| GET | `/api/admin/virtual-directories/{id}/methods` | Get allowed methods |
| PUT | `/api/admin/virtual-directories/{id}/methods` | Set allowed methods |

### Create Virtual Directory

**Request Body:**
```json
{
    "virtual_host_id": 1,
    "source_path": "/api/v1",
    "target_path": "/",
    "match_type": "prefix",
    "strip_prefix": true,
    "preserve_host_header": false,
    "auth_type": "none",
    "is_active": true,
    "proxy_timeout_seconds": 30,
    "retry_count": 2,
    "retry_delay_ms": 100,
    "max_request_size_mb": 10,
    "websocket_enabled": false,
    "cache_enabled": false,
    "cache_ttl_seconds": 0
}
```

auth_type yang valid: `none`, `api_key`, `jwt`, `basic`, `external`

### Set Methods

**Request Body:**
```json
{
    "methods": ["GET", "POST", "PUT", "DELETE"]
}
```

---

## Consumer Credentials

Manajemen credential (Basic Auth, dll) untuk consumer.

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| POST | `/api/admin/consumer-credentials` | Buat credential |
| GET | `/api/admin/consumer-credentials` | List semua credential |
| GET | `/api/admin/consumer-credentials/{id}` | Detail credential |
| GET | `/api/admin/consumers/{consumer_id}/credentials` | Credential by Consumer |
| PUT | `/api/admin/consumer-credentials/{id}` | Update credential |
| DELETE | `/api/admin/consumer-credentials/{id}` | Hapus credential |

### Create Credential

**Request Body:**
```json
{
    "consumer_id": 1,
    "auth_type": "basic",
    "username": "api-user",
    "password": "secret123",
    "is_active": true
}
```

---

## API Keys

Manajemen API Key untuk consumer.

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| POST | `/api/admin/api-keys` | Buat API Key |
| GET | `/api/admin/api-keys` | List semua API Key |
| GET | `/api/admin/api-keys/{id}` | Detail API Key |
| GET | `/api/admin/consumers/{consumer_id}/api-keys` | API Key by Consumer |
| PUT | `/api/admin/api-keys/{id}` | Update API Key |
| DELETE | `/api/admin/api-keys/{id}` | Hapus API Key |

### Create API Key

**Request Body:**
```json
{
    "consumer_id": 1,
    "description": "Production API Key",
    "expired_at": "2027-12-31 23:59:59",
    "rate_limit_override": 1000,
    "is_active": true
}
```

---

## Route Consumer Access (ACL)

Manajemen akses consumer ke route tertentu.

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| POST | `/api/admin/route-access` | Buat akses |
| GET | `/api/admin/route-access` | List semua akses |
| GET | `/api/admin/route-access/{id}` | Detail akses |
| GET | `/api/admin/virtual-directories/{dir_id}/access` | Akses by Directory |
| PUT | `/api/admin/route-access/{id}` | Update akses |
| DELETE | `/api/admin/route-access/{id}` | Hapus akses |

### Create Access

**Request Body:**
```json
{
    "virtual_directory_id": 1,
    "consumer_id": 1,
    "is_active": true
}
```

---

## JWT Configs

Konfigurasi JWT validation per route.

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| POST | `/api/admin/jwt-configs` | Buat JWT Config |
| GET | `/api/admin/jwt-configs` | List semua JWT Config |
| GET | `/api/admin/jwt-configs/{id}` | Detail JWT Config |
| PUT | `/api/admin/jwt-configs/{id}` | Update JWT Config |
| DELETE | `/api/admin/jwt-configs/{id}` | Hapus JWT Config |

### Create JWT Config

**Request Body:**
```json
{
    "virtual_directory_id": 1,
    "algorithm": "HS256",
    "jwt_secret": "my-secret-key-at-least-32-chars-long",
    "issuer": "my-app",
    "audience": "api-gateway",
    "expired_in_seconds": 3600,
    "clock_skew_seconds": 30,
    "require_exp": true,
    "require_iat": true,
    "is_active": true
}
```

---

## External Auth

Konfigurasi external auth service (forward auth).

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| POST | `/api/admin/external-auth` | Buat External Auth |
| GET | `/api/admin/external-auth` | List semua External Auth |
| GET | `/api/admin/external-auth/{id}` | Detail External Auth |
| PUT | `/api/admin/external-auth/{id}` | Update External Auth |
| DELETE | `/api/admin/external-auth/{id}` | Hapus External Auth |

### Create External Auth

**Request Body:**
```json
{
    "virtual_directory_id": 1,
    "auth_url": "https://auth.example.com/verify",
    "http_method": "POST",
    "request_timeout_seconds": 5,
    "send_headers": true,
    "send_body": false,
    "success_key": "status",
    "success_value": "true",
    "message_key": "message",
    "is_active": true
}
```

---

## Rate Limits

Konfigurasi rate limiting per route.

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| POST | `/api/admin/rate-limits` | Buat Rate Limit |
| GET | `/api/admin/rate-limits` | List semua Rate Limit |
| GET | `/api/admin/rate-limits/{id}` | Detail Rate Limit |
| PUT | `/api/admin/rate-limits/{id}` | Update Rate Limit |
| DELETE | `/api/admin/rate-limits/{id}` | Hapus Rate Limit |

### Create Rate Limit

**Request Body:**
```json
{
    "virtual_directory_id": 1,
    "limit_by": "ip",
    "requests_per_minute": 60,
    "burst": 10,
    "block_duration_seconds": 60,
    "is_active": true
}
```

limit_by yang valid: `ip`, `api_key`

---

## CORS Configs

Konfigurasi CORS per route.

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| POST | `/api/admin/cors-configs` | Buat CORS Config |
| GET | `/api/admin/cors-configs` | List semua CORS Config |
| GET | `/api/admin/cors-configs/{id}` | Detail CORS Config |
| PUT | `/api/admin/cors-configs/{id}` | Update CORS Config |
| DELETE | `/api/admin/cors-configs/{id}` | Hapus CORS Config |

### Create CORS Config

**Request Body:**
```json
{
    "virtual_directory_id": 1,
    "allowed_origins": "https://app.example.com,https://admin.example.com",
    "allowed_methods": "GET,POST,PUT,DELETE,OPTIONS",
    "allowed_headers": "Content-Type,Authorization,X-API-Key",
    "exposed_headers": "X-Request-Id",
    "allow_credentials": true,
    "max_age_seconds": 3600,
    "is_active": true
}
```

---

## Circuit Breakers

Konfigurasi circuit breaker per route.

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| POST | `/api/admin/circuit-breakers` | Buat Circuit Breaker |
| GET | `/api/admin/circuit-breakers` | List semua Circuit Breaker |
| GET | `/api/admin/circuit-breakers/{id}` | Detail Circuit Breaker |
| PUT | `/api/admin/circuit-breakers/{id}` | Update Circuit Breaker |
| DELETE | `/api/admin/circuit-breakers/{id}` | Hapus Circuit Breaker |

### Create Circuit Breaker

**Request Body:**
```json
{
    "virtual_directory_id": 1,
    "enabled": true,
    "failure_threshold": 5,
    "recovery_timeout_seconds": 30,
    "half_open_max_requests": 3
}
```

---

## IP Whitelists

Whitelist IP address per route.

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| POST | `/api/admin/ip-whitelists` | Buat IP Whitelist |
| GET | `/api/admin/ip-whitelists` | List semua IP Whitelist |
| GET | `/api/admin/ip-whitelists/{id}` | Detail IP Whitelist |
| GET | `/api/admin/virtual-directories/{dir_id}/ip-whitelists` | Whitelist by Directory |
| PUT | `/api/admin/ip-whitelists/{id}` | Update IP Whitelist |
| DELETE | `/api/admin/ip-whitelists/{id}` | Hapus IP Whitelist |

### Create IP Whitelist

**Request Body:**
```json
{
    "virtual_directory_id": 1,
    "ip_address": "192.168.1.0/24",
    "description": "Office network",
    "is_active": true
}
```

---

## IP Blacklists

Blacklist IP address per route.

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| POST | `/api/admin/ip-blacklists` | Buat IP Blacklist |
| GET | `/api/admin/ip-blacklists` | List semua IP Blacklist |
| GET | `/api/admin/ip-blacklists/{id}` | Detail IP Blacklist |
| GET | `/api/admin/virtual-directories/{dir_id}/ip-blacklists` | Blacklist by Directory |
| PUT | `/api/admin/ip-blacklists/{id}` | Update IP Blacklist |
| DELETE | `/api/admin/ip-blacklists/{id}` | Hapus IP Blacklist |

### Create IP Blacklist

**Request Body:**
```json
{
    "virtual_directory_id": 1,
    "ip_address": "203.0.113.50",
    "reason": "Suspicious activity",
    "expired_at": "2027-01-01 00:00:00",
    "is_active": true
}
```

---

## Request Header Rules

Atur manipulasi header request sebelum diteruskan ke upstream.

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| POST | `/api/admin/request-headers` | Buat Header Rule |
| GET | `/api/admin/request-headers` | List semua Header Rule |
| GET | `/api/admin/request-headers/{id}` | Detail Header Rule |
| GET | `/api/admin/virtual-directories/{dir_id}/request-headers` | Header by Directory |
| PUT | `/api/admin/request-headers/{id}` | Update Header Rule |
| DELETE | `/api/admin/request-headers/{id}` | Hapus Header Rule |

### Create Request Header Rule

**Request Body:**
```json
{
    "virtual_directory_id": 1,
    "header_name": "X-Request-ID",
    "operation": "set",
    "value_source": "static",
    "header_value": "gateway-request",
    "execution_order": 1,
    "is_active": true
}
```

operation yang valid: `set`, `add`, `remove`

value_source yang valid: `static`, `variable`, `header`, `query`

---

## Response Header Rules

Atur manipulasi header response dari upstream sebelum dikirim ke client.

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| POST | `/api/admin/response-headers` | Buat Header Rule |
| GET | `/api/admin/response-headers` | List semua Header Rule |
| GET | `/api/admin/response-headers/{id}` | Detail Header Rule |
| GET | `/api/admin/virtual-directories/{dir_id}/response-headers` | Header by Directory |
| PUT | `/api/admin/response-headers/{id}` | Update Header Rule |
| DELETE | `/api/admin/response-headers/{id}` | Hapus Header Rule |

### Create Response Header Rule

**Request Body:**
```json
{
    "virtual_directory_id": 1,
    "header_name": "X-Powered-By",
    "operation": "delete",
    "header_value": "",
    "execution_order": 1,
    "is_active": true
}
```

---

## Query Rewrites

Manipulasi query string sebelum diteruskan ke upstream.

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| POST | `/api/admin/query-rewrites` | Buat Query Rewrite |
| GET | `/api/admin/query-rewrites` | List semua Query Rewrite |
| GET | `/api/admin/query-rewrites/{id}` | Detail Query Rewrite |
| GET | `/api/admin/virtual-directories/{dir_id}/query-rewrites` | Query by Directory |
| PUT | `/api/admin/query-rewrites/{id}` | Update Query Rewrite |
| DELETE | `/api/admin/query-rewrites/{id}` | Hapus Query Rewrite |

### Create Query Rewrite

**Request Body:**
```json
{
    "virtual_directory_id": 1,
    "param_name": "version",
    "param_value": "v2",
    "operation": "set"
}
```

operation yang valid: `set`, `add`, `delete`

---

## ACME Accounts

Manajemen akun ACME untuk sertifikat SSL otomatis (Let's Encrypt).

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| POST | `/api/admin/acme-accounts` | Buat ACME Account |
| GET | `/api/admin/acme-accounts` | List semua ACME Account |
| GET | `/api/admin/acme-accounts/{id}` | Detail ACME Account |
| PUT | `/api/admin/acme-accounts/{id}` | Update ACME Account |
| DELETE | `/api/admin/acme-accounts/{id}` | Hapus ACME Account |

### Create ACME Account

**Request Body:**
```json
{
    "email": "admin@example.com",
    "provider_url": "https://acme-v02.api.letsencrypt.org/directory",
    "account_key_path": "/etc/ssl/acme/account.key",
    "is_default": true
}
```

---

## SSL Certificates

Manajemen sertifikat SSL.

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| POST | `/api/admin/ssl-certificates` | Buat SSL Certificate |
| GET | `/api/admin/ssl-certificates` | List semua SSL Certificate |
| GET | `/api/admin/ssl-certificates/{id}` | Detail SSL Certificate |
| PUT | `/api/admin/ssl-certificates/{id}` | Update SSL Certificate |
| DELETE | `/api/admin/ssl-certificates/{id}` | Hapus SSL Certificate |

### Create SSL Certificate

**Request Body:**
```json
{
    "acme_account_id": 1,
    "provider": "letsencrypt",
    "challenge_type": "http-01",
    "certificate_path": "/etc/ssl/certs/example.crt",
    "private_key_path": "/etc/ssl/private/example.key",
    "chain_path": "/etc/ssl/certs/chain.pem",
    "auto_renew": true,
    "renew_before_days": 30,
    "expired_at": "2027-12-31 23:59:59",
    "is_active": true
}
```

---

## Certificate Domains

Manajemen domain yang terikat ke sertifikat SSL.

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| POST | `/api/admin/certificate-domains` | Buat Certificate Domain |
| GET | `/api/admin/certificate-domains` | List semua Certificate Domain |
| GET | `/api/admin/certificate-domains/{id}` | Detail Certificate Domain |
| GET | `/api/admin/ssl-certificates/{cert_id}/domains` | Domain by Certificate |
| PUT | `/api/admin/certificate-domains/{id}` | Update Certificate Domain |
| DELETE | `/api/admin/certificate-domains/{id}` | Hapus Certificate Domain |

### Create Certificate Domain

**Request Body:**
```json
{
    "ssl_certificate_id": 1,
    "domain_name": "example.com",
    "is_wildcard": false
}
```

---

## SSL Bindings

Binding sertifikat SSL ke host/virtual host.

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| POST | `/api/admin/ssl-bindings` | Buat SSL Binding |
| GET | `/api/admin/ssl-bindings` | List semua SSL Binding |
| GET | `/api/admin/ssl-bindings/{id}` | Detail SSL Binding |
| PUT | `/api/admin/ssl-bindings/{id}` | Update SSL Binding |
| DELETE | `/api/admin/ssl-bindings/{id}` | Hapus SSL Binding |

### Create SSL Binding

**Request Body:**
```json
{
    "ssl_certificate_id": 1,
    "binding_type": "host",
    "host_id": 1,
    "is_default": true,
    "priority": 1
}
```

binding_type yang valid: `host`, `virtual_host`

---

## TLS Options

Konfigurasi opsi TLS per host/virtual host.

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| POST | `/api/admin/tls-options` | Buat TLS Option |
| GET | `/api/admin/tls-options` | List semua TLS Option |
| GET | `/api/admin/tls-options/{id}` | Detail TLS Option |
| PUT | `/api/admin/tls-options/{id}` | Update TLS Option |
| DELETE | `/api/admin/tls-options/{id}` | Hapus TLS Option |

### Create TLS Option

**Request Body:**
```json
{
    "binding_type": "host",
    "host_id": 1,
    "min_tls_version": "1.2",
    "http2_enabled": true,
    "hsts_enabled": true,
    "hsts_max_age": 31536000
}
```

min_tls_version yang valid: `1.0`, `1.1`, `1.2`, `1.3`

---

## Service Discovery

Konfigurasi service discovery untuk upstream otomatis.

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| POST | `/api/admin/service-discovery` | Buat Service Discovery |
| GET | `/api/admin/service-discovery` | List semua Service Discovery |
| GET | `/api/admin/service-discovery/{id}` | Detail Service Discovery |
| PUT | `/api/admin/service-discovery/{id}` | Update Service Discovery |
| DELETE | `/api/admin/service-discovery/{id}` | Hapus Service Discovery |

### Create Service Discovery

**Request Body:**
```json
{
    "virtual_host_id": 1,
    "provider": "consul",
    "endpoint_url": "http://consul:8500/v1/catalog/service/my-api",
    "refresh_interval_seconds": 30,
    "is_active": true
}
```

---

## Config Versions

Versioning konfigurasi untuk tracking perubahan.

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| POST | `/api/admin/config-versions` | Buat Config Version |
| GET | `/api/admin/config-versions` | List semua Config Version |
| GET | `/api/admin/config-versions/{id}` | Detail Config Version |
| PUT | `/api/admin/config-versions/{id}` | Update Config Version |
| DELETE | `/api/admin/config-versions/{id}` | Hapus Config Version |

### Create Config Version

**Request Body:**
```json
{
    "config_name": "proxy-routes",
    "version_number": 1,
    "changed_by": "admin",
    "notes": "Initial configuration"
}
```

---

## Maintenance Windows

Jadwal maintenance per virtual host.

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| POST | `/api/admin/maintenance-windows` | Buat Maintenance Window |
| GET | `/api/admin/maintenance-windows` | List semua Maintenance Window |
| GET | `/api/admin/maintenance-windows/{id}` | Detail Maintenance Window |
| PUT | `/api/admin/maintenance-windows/{id}` | Update Maintenance Window |
| DELETE | `/api/admin/maintenance-windows/{id}` | Hapus Maintenance Window |

### Create Maintenance Window

**Request Body:**
```json
{
    "virtual_host_id": 1,
    "title": "Scheduled Maintenance",
    "start_at": "2026-06-01 00:00:00",
    "end_at": "2026-06-01 04:00:00",
    "maintenance_response_code": 503,
    "maintenance_message": "Service is under scheduled maintenance. Please try again later.",
    "is_active": true
}
```

---

## Response Format

Semua response menggunakan format standar:

**Success:**
```json
{
    "status": "success",
    "message": "Deskripsi pesan",
    "data": { ... }
}
```

**Error:**
```json
{
    "status": "error",
    "message": "Deskripsi error"
}
```

**HTTP Status Codes:**
| Code | Deskripsi |
|------|-----------|
| 200 | OK - Request berhasil |
| 201 | Created - Resource berhasil dibuat |
| 400 | Bad Request - Input tidak valid |
| 401 | Unauthorized - Token tidak valid atau expired |
| 403 | Forbidden - Role tidak memiliki akses |
| 404 | Not Found - Resource tidak ditemukan |
| 500 | Internal Server Error |

---

## CLI Commands

Swantara Gate menyediakan CLI tool untuk manajemen user:

### Buat User Baru
```bash
./swantara-cli create-user \
  -username=admin \
  -password=secret123 \
  -fullname="Super Admin" \
  -email=admin@example.com \
  -role=super_admin \
  -active=true
```

### Reset Password
```bash
./swantara-cli reset-password \
  -username=admin \
  -password=newpassword123
```

### Build CLI
```bash
go build -o swantara-cli.exe ./cmd/cli/
```

---

## Cara Import ke Postman

1. Buka Postman
2. Klik **Import** di kiri atas
3. Pilih file `Swantara_Gate_API.postman_collection.json`
4. Collection akan muncul di sidebar dengan semua endpoint siap pakai
5. Login dulu di endpoint **Auth > Login**, token akan otomatis tersimpan
