# Swantara Gate - API Documentation

**Base URL:** `http://localhost:8081`

**Format Response:**
```json
{
  "success": true,
  "message": "Pesan deskriptif",
  "data": {},
  "error": ""
}
```

---

## Health Check

### GET /api/health
Cek status server.

**Response:**
```json
{"status": "ok", "message": "Swantara Gate Admin API is running"}
```

---

## FASE 1 - Foundation CRUD

---

### 1. Users

#### POST /api/admin/users
Membuat user admin baru.

**Request Body:**
```json
{
  "username": "admin",
  "password": "secret123",
  "full_name": "Administrator",
  "email": "admin@example.com",
  "role": "admin",
  "is_active": true
}
```
| Field | Type | Required | Keterangan |
|-------|------|----------|------------|
| username | string | Ya | Username unik |
| password | string | Ya | Password (akan di-hash bcrypt) |
| full_name | string | Tidak | Nama lengkap |
| email | string | Tidak | Email |
| role | string | Tidak | Role: `admin` / `operator` (default: operator) |
| is_active | bool | Tidak | Status aktif |

#### GET /api/admin/users
Mengambil semua users dengan pagination.

**Query Parameters:**
| Param | Default | Keterangan |
|-------|---------|------------|
| page | 1 | Halaman |
| limit | 10 | Jumlah per halaman (max 100) |

**Response:**
```json
{
  "success": true,
  "message": "Daftar user berhasil diambil",
  "data": {
    "users": [...],
    "pagination": {"page": 1, "limit": 10, "total": 5}
  }
}
```

#### GET /api/admin/users/{id}
Mengambil user berdasarkan ID.

#### PUT /api/admin/users/{id}
Update data user.

**Request Body:**
```json
{
  "full_name": "Admin Updated",
  "email": "new@example.com",
  "role": "admin",
  "is_active": true
}
```

#### DELETE /api/admin/users/{id}
Menghapus user.

---

### 2. API Consumers

#### POST /api/admin/consumers
Membuat API consumer baru.

**Request Body:**
```json
{
  "consumer_name": "Mobile App",
  "description": "Aplikasi mobile utama",
  "contact_email": "dev@app.com",
  "is_active": true
}
```
| Field | Type | Required | Keterangan |
|-------|------|----------|------------|
| consumer_name | string | Ya | Nama consumer unik |
| description | string | Tidak | Deskripsi |
| contact_email | string | Tidak | Email kontak |
| is_active | bool | Tidak | Status aktif |

#### GET /api/admin/consumers
Mengambil semua consumers dengan pagination.

**Query Parameters:** `page`, `limit`

#### GET /api/admin/consumers/{id}
Mengambil consumer berdasarkan ID.

#### PUT /api/admin/consumers/{id}
Update consumer.

**Request Body:**
```json
{
  "description": "Updated desc",
  "contact_email": "new@app.com",
  "is_active": true
}
```

#### DELETE /api/admin/consumers/{id}
Menghapus consumer.

---

### 3. Hosts

#### POST /api/admin/hosts
Membuat host baru.

**Request Body:**
```json
{
  "host_name": "api.example.com",
  "description": "Main API host",
  "is_active": true
}
```
| Field | Type | Required | Keterangan |
|-------|------|----------|------------|
| host_name | string | Ya | Nama domain/host |
| description | string | Tidak | Deskripsi |
| is_active | bool | Tidak | Status aktif |

#### GET /api/admin/hosts
Mengambil semua hosts dengan pagination.

#### GET /api/admin/hosts/{id}
Mengambil host berdasarkan ID.

#### PUT /api/admin/hosts/{id}
Update host.

**Request Body:**
```json
{
  "description": "Updated desc",
  "is_active": true
}
```

#### DELETE /api/admin/hosts/{id}
Menghapus host.

---

### 4. Virtual Hosts

#### POST /api/admin/virtual-hosts
Membuat virtual host baru.

**Request Body:**
```json
{
  "host_id": 1,
  "vhost_name": "api-v1.example.com",
  "lb_algorithm": "round_robin",
  "sticky_session": false,
  "failover_mode": "active_passive",
  "is_active": true
}
```
| Field | Type | Required | Keterangan |
|-------|------|----------|------------|
| host_id | int | Ya | ID host parent |
| vhost_name | string | Ya | Nama virtual host |
| lb_algorithm | string | Tidak | Algoritma LB: `round_robin` / `least_conn` / `ip_hash` / `weighted` |
| sticky_session | bool | Tidak | Aktifkan sticky session |
| failover_mode | string | Tidak | Mode failover: `active_passive` / `active_active` |
| is_active | bool | Tidak | Status aktif |

#### GET /api/admin/virtual-hosts
Mengambil semua virtual hosts dengan pagination.

#### GET /api/admin/virtual-hosts/{id}
Mengambil virtual host berdasarkan ID.

#### PUT /api/admin/virtual-hosts/{id}
Update virtual host.

**Request Body:**
```json
{
  "lb_algorithm": "least_conn",
  "sticky_session": true,
  "failover_mode": "active_active",
  "is_active": true
}
```

#### DELETE /api/admin/virtual-hosts/{id}
Menghapus virtual host.

---

## FASE 2 - Extended CRUD

---

### 5. Upstream Servers

#### POST /api/admin/upstream-servers
Membuat upstream server baru (backend target).

**Request Body:**
```json
{
  "virtual_host_id": 1,
  "target_host": "192.168.1.10",
  "target_port": 8080,
  "protocol": "http",
  "priority": 1,
  "weight": 10,
  "is_backup": false,
  "is_active": true,
  "health_check_enabled": true,
  "health_check_path": "/health",
  "health_check_interval_seconds": 30,
  "health_check_timeout_seconds": 5,
  "max_fails": 3,
  "fail_timeout_seconds": 30
}
```
| Field | Type | Required | Default | Keterangan |
|-------|------|----------|---------|------------|
| virtual_host_id | int | Ya | - | ID virtual host |
| target_host | string | Ya | - | IP/hostname backend |
| target_port | int | Ya | - | Port backend (1-65535) |
| protocol | string | Tidak | `http` | Protocol: `http` / `https` / `grpc` |
| priority | int | Tidak | 0 | Prioritas (semakin kecil = lebih tinggi) |
| weight | int | Tidak | 1 | Bobot untuk weighted LB |
| is_backup | bool | Tidak | false | Tandai sebagai backup server |
| is_active | bool | Tidak | false | Status aktif |
| health_check_enabled | bool | Tidak | false | Aktifkan health check |
| health_check_path | string | Tidak | "" | Path health check |
| health_check_interval_seconds | int | Tidak | 30 | Interval check (detik) |
| health_check_timeout_seconds | int | Tidak | 5 | Timeout check (detik) |
| max_fails | int | Tidak | 3 | Max gagal sebelum down |
| fail_timeout_seconds | int | Tidak | 30 | Durasi server dianggap down |

#### GET /api/admin/upstream-servers
Mengambil semua upstream servers dengan pagination.

#### GET /api/admin/upstream-servers/{id}
Mengambil upstream server berdasarkan ID.

#### GET /api/admin/virtual-hosts/{vhost_id}/upstream-servers
Mengambil semua upstream servers milik virtual host tertentu.

#### PUT /api/admin/upstream-servers/{id}
Update upstream server.

**Request Body:** (semua field dari create kecuali `virtual_host_id`)

#### DELETE /api/admin/upstream-servers/{id}
Menghapus upstream server.

---

### 6. Virtual Directories (Routes)

#### POST /api/admin/virtual-directories
Membuat route/virtual directory baru.

**Request Body:**
```json
{
  "virtual_host_id": 1,
  "source_path": "/api/v1/users",
  "target_path": "/users",
  "match_type": "prefix",
  "strip_prefix": true,
  "preserve_host_header": false,
  "auth_type": "api_key",
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
| Field | Type | Required | Default | Keterangan |
|-------|------|----------|---------|------------|
| virtual_host_id | int | Ya | - | ID virtual host |
| source_path | string | Ya | - | Path masuk dari client |
| target_path | string | Ya | - | Path tujuan di backend |
| match_type | string | Tidak | `prefix` | Tipe match: `prefix` / `exact` / `regex` |
| strip_prefix | bool | Tidak | false | Hapus prefix saat forward |
| preserve_host_header | bool | Tidak | false | Pertahankan Host header |
| auth_type | string | Tidak | `none` | Auth: `none` / `api_key` / `jwt` / `basic` |
| is_active | bool | Tidak | false | Status aktif |
| proxy_timeout_seconds | int | Tidak | 30 | Timeout proxy (detik) |
| retry_count | int | Tidak | 0 | Jumlah retry saat gagal |
| retry_delay_ms | int | Tidak | 0 | Delay antar retry (ms) |
| max_request_size_mb | int | Tidak | 10 | Max ukuran request (MB) |
| websocket_enabled | bool | Tidak | false | Dukung websocket |
| cache_enabled | bool | Tidak | false | Aktifkan cache |
| cache_ttl_seconds | int | Tidak | 0 | TTL cache (detik) |

#### GET /api/admin/virtual-directories
Mengambil semua virtual directories dengan pagination.

#### GET /api/admin/virtual-directories/{id}
Mengambil virtual directory berdasarkan ID.

#### GET /api/admin/virtual-hosts/{vhost_id}/virtual-directories
Mengambil semua directories milik virtual host tertentu.

#### PUT /api/admin/virtual-directories/{id}
Update virtual directory.

#### DELETE /api/admin/virtual-directories/{id}
Menghapus virtual directory (dan methods-nya).

---

### 7. Virtual Directory Methods

#### GET /api/admin/virtual-directories/{id}/methods
Mengambil HTTP methods yang diizinkan untuk directory.

**Response:**
```json
{
  "success": true,
  "data": [
    {"id": 1, "virtual_directory_id": 1, "http_method": "GET"},
    {"id": 2, "virtual_directory_id": 1, "http_method": "POST"}
  ]
}
```

#### PUT /api/admin/virtual-directories/{id}/methods
Mengatur HTTP methods (replace semua).

**Request Body:**
```json
{
  "methods": ["GET", "POST", "PUT", "DELETE"]
}
```
| Nilai Valid | Keterangan |
|-------------|------------|
| GET | HTTP GET |
| POST | HTTP POST |
| PUT | HTTP PUT |
| PATCH | HTTP PATCH |
| DELETE | HTTP DELETE |
| HEAD | HTTP HEAD |
| OPTIONS | HTTP OPTIONS |

---

### 8. Consumer Credentials

#### POST /api/admin/consumer-credentials
Membuat credential autentikasi untuk consumer.

**Request Body (Basic Auth):**
```json
{
  "consumer_id": 1,
  "auth_type": "basic",
  "username": "app_user",
  "password": "secret123",
  "is_active": true
}
```

**Request Body (API Key):**
```json
{
  "consumer_id": 1,
  "auth_type": "api_key",
  "api_key": "my-custom-key",
  "expired_at": "2027-12-31 23:59:59",
  "is_active": true
}
```

**Request Body (JWT):**
```json
{
  "consumer_id": 1,
  "auth_type": "jwt",
  "jwt_secret": "my-jwt-secret-key",
  "is_active": true
}
```

| Field | Type | Required | Keterangan |
|-------|------|----------|------------|
| consumer_id | int | Ya | ID consumer |
| auth_type | string | Ya | Tipe: `basic` / `api_key` / `jwt` |
| username | string | Kondisional | Wajib jika auth_type = basic |
| password | string | Kondisional | Wajib jika auth_type = basic (di-hash bcrypt) |
| api_key | string | Tidak | Custom API key |
| jwt_secret | string | Tidak | JWT signing secret |
| expired_at | string | Tidak | Format: `YYYY-MM-DD HH:MM:SS` |
| is_active | bool | Tidak | Status aktif |

#### GET /api/admin/consumer-credentials
Mengambil semua credentials dengan pagination.

#### GET /api/admin/consumer-credentials/{id}
Mengambil credential berdasarkan ID.

#### GET /api/admin/consumers/{consumer_id}/credentials
Mengambil semua credentials milik consumer tertentu.

#### PUT /api/admin/consumer-credentials/{id}
Update credential.

**Request Body:**
```json
{
  "username": "new_user",
  "password": "new_pass",
  "api_key": "new-key",
  "jwt_secret": "new-secret",
  "expired_at": "2028-12-31 23:59:59",
  "is_active": true
}
```

#### DELETE /api/admin/consumer-credentials/{id}
Menghapus credential.

---

### 9. API Keys

#### POST /api/admin/api-keys
Membuat API key baru (key di-generate otomatis dengan prefix `sgk_`).

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
| Field | Type | Required | Keterangan |
|-------|------|----------|------------|
| consumer_id | int | Ya | ID consumer |
| description | string | Tidak | Deskripsi key |
| expired_at | string | Tidak | Format: `YYYY-MM-DD HH:MM:SS` |
| rate_limit_override | int | Tidak | Override rate limit (req/menit) |
| is_active | bool | Tidak | Status aktif |

**Response:**
```json
{
  "success": true,
  "message": "API key berhasil dibuat",
  "data": {
    "id": 1,
    "consumer_id": 1,
    "api_key": "sgk_a1b2c3d4e5f6...",
    "description": "Production API Key",
    "expired_at": "2027-12-31T23:59:59Z",
    "rate_limit_override": 1000,
    "is_active": true,
    "created_at": "2026-05-15T10:00:00Z"
  }
}
```

#### GET /api/admin/api-keys
Mengambil semua API keys dengan pagination.

#### GET /api/admin/api-keys/{id}
Mengambil API key berdasarkan ID.

#### GET /api/admin/consumers/{consumer_id}/api-keys
Mengambil semua API keys milik consumer tertentu.

#### PUT /api/admin/api-keys/{id}
Update API key (key tidak berubah, hanya metadata).

**Request Body:**
```json
{
  "description": "Updated description",
  "expired_at": "2028-06-30 00:00:00",
  "rate_limit_override": 2000,
  "is_active": true
}
```

#### DELETE /api/admin/api-keys/{id}
Menghapus API key.

---

### 10. Route Consumer Access (ACL)

#### POST /api/admin/route-access
Memberikan akses consumer ke route tertentu.

**Request Body:**
```json
{
  "virtual_directory_id": 1,
  "consumer_id": 1,
  "is_active": true
}
```
| Field | Type | Required | Keterangan |
|-------|------|----------|------------|
| virtual_directory_id | int | Ya | ID virtual directory/route |
| consumer_id | int | Ya | ID consumer |
| is_active | bool | Tidak | Status akses aktif |

#### GET /api/admin/route-access
Mengambil semua access dengan pagination.

#### GET /api/admin/route-access/{id}
Mengambil access berdasarkan ID.

#### GET /api/admin/virtual-directories/{dir_id}/access
Mengambil semua access untuk directory tertentu.

#### PUT /api/admin/route-access/{id}
Update status akses.

**Request Body:**
```json
{
  "is_active": false
}
```

#### DELETE /api/admin/route-access/{id}
Menghapus akses.

---

## Catatan Umum

### Pagination
Semua endpoint GET list mendukung pagination:
```
GET /api/admin/users?page=2&limit=20
```

### Error Response
```json
{
  "success": false,
  "message": "Deskripsi error",
  "data": null,
  "error": "Detail error"
}
```

### HTTP Status Codes
| Code | Keterangan |
|------|------------|
| 200 | OK - Request berhasil |
| 201 | Created - Resource berhasil dibuat |
| 400 | Bad Request - Input tidak valid |
| 404 | Not Found - Resource tidak ditemukan |
| 409 | Conflict - Resource sudah ada (duplikat) |
| 500 | Internal Server Error - Kesalahan server |

### Alur Konfigurasi Gateway
```
1. Buat Host → 2. Buat Virtual Host (link ke Host)
→ 3. Buat Upstream Server (link ke VHost)
→ 4. Buat Virtual Directory/Route (link ke VHost)
→ 5. Set Methods pada Route
→ 6. Buat Consumer → 7. Buat Credential/API Key
→ 8. Berikan Route Access ke Consumer
```
