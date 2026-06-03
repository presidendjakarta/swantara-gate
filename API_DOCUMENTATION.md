# Swantara Gate API Documentation

**Base URL:** `http://localhost:8081`

**Credentials:**
- Username: `admin`
- Password: `admin1324`

---

## Health Check

### 1. GET /api/health

**Status Code:** `200`

**Response:**
```json
{
  "status": "ok",
  "message": "Swantara Gate Admin API is running"
}
```

**Duration:** 8ms

---

## Authentication

### 2. POST /api/admin/auth/login

**Request Body:**
```json
{
  "password": "admin1324",
  "username": "testcli"
}
```

**Status Code:** `200`

**Response:**
```json
{
  "success": true,
  "message": "Login berhasil",
  "data": {
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3ODA0NzQyMjksImlhdCI6MTc4MDQ3MjQyOSwicm9sZSI6ImFkbWluIiwidXNlcl9pZCI6MSwidXNlcm5hbWUiOiJ0ZXN0Y2xpIn0.kK5yNNjjiUriIHVmPE-bSt5_caO6o7L9damiaZp06K0",
    "refresh_token": "ffb85a023f89d25de3fe000cf2689728e1c3b4bf0a839a2040aa882ace0e234c",
    "token_type": "Bearer",
    "expires_in": 1800
  }
}

```

**Duration:** 50ms

---

### 3. GET /api/admin/auth/me

**Status Code:** `200`

**Response:**
```json
{
  "success": true,
  "message": "Profil user",
  "data": {
    "role": "admin",
    "user_id": 1,
    "username": "testcli"
  }
}

```

**Duration:** 1ms

---

### 4. POST /api/admin/auth/refresh

**Request Body:**
```json
{
  "refresh_token": "ffb85a023f89d25de3fe000cf2689728e1c3b4bf0a839a2040aa882ace0e234c"
}
```

**Status Code:** `200`

**Response:**
```json
{
  "success": true,
  "message": "Token berhasil diperbarui",
  "data": {
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3ODA0NzQyMjksImlhdCI6MTc4MDQ3MjQyOSwicm9sZSI6ImFkbWluIiwidXNlcl9pZCI6MSwidXNlcm5hbWUiOiJ0ZXN0Y2xpIn0.kK5yNNjjiUriIHVmPE-bSt5_caO6o7L9damiaZp06K0",
    "refresh_token": "426db7048f7c750924b2b8f786a9e5551310800a396069e3d00dca4d080e3950",
    "token_type": "Bearer",
    "expires_in": 1800
  }
}

```

**Duration:** 4ms

---

## Configuration

### 5. POST /api/admin/config/reload

**Status Code:** `200`

**Response:**
```json
{
  "status": "ok",
  "message": "Proxy configuration reloaded successfully"
}
```

**Duration:** 4ms

---

## Users Management

### 6. GET /api/admin/users

**Status Code:** `200`

**Response:**
```json
{
  "success": true,
  "message": "Daftar user berhasil diambil",
  "data": {
    "pagination": {
      "limit": 10,
      "page": 1,
      "total": 5
    },
    "users": [
      {
        "id": 8,
        "username": "apitest_1780472247",
        "full_name": "API Test User",
        "email": "apitest1780472247@example.com",
        "role": "admin",
        "is_active": true,
        "created_at": "2026-06-03T07:37:27Z",
        "updated_at": "2026-06-03T07:37:27Z"
      },
      {
        "id": 7,
        "username": "apitest_1780472241",
        "full_name": "API Test User",
        "email": "apitest1780472241@example.com",
        "role": "admin",
        "is_active": true,
        "created_at": "2026-06-03T07:37:21Z",
        "updated_at": "2026-06-03T07:37:21Z"
      },
      {
        "id": 6,
        "username": "apitest_user",
        "full_name": "API Test User",
        "email": "apitest@example.com",
        "role": "admin",
        "is_active": true,
        "created_at": "2026-06-03T07:34:25Z",
        "updated_at": "2026-06-03T07:34:25Z"
      },
      {
        "id": 4,
        "username": "testuser",
        "full_name": "Test User",
        "email": "test@example.com",
        "role": "admin",
        "is_active": true,
        "created_at": "2026-06-03T07:17:31Z",
        "updated_at": "2026-06-03T07:17:31Z"
      },
      {
        "id": 1,
        "username": "testcli",
        "full_name": "Updated Name",
        "email": "updated@example.com",
        "role": "admin",
        "is_active": true,
        "last_login_at": "2026-06-03T14:40:29.2062808+07:00",
        "created_at": "2026-05-28T14:10:50.0340632+07:00",
        "updated_at": "2026-06-03T14:40:29.2062808+07:00"
      }
    ]
  }
}

```

**Duration:** 0s

---

### 7. POST /api/admin/users

**Request Body:**
```json
{
  "email": "apitest1780472429@example.com",
  "full_name": "API Test User",
  "is_active": true,
  "password": "test123",
  "role": "admin",
  "username": "apitest_1780472429"
}
```

**Status Code:** `201`

**Response:**
```json
{
  "success": true,
  "message": "User berhasil dibuat",
  "data": {
    "id": 9,
    "username": "apitest_1780472429",
    "full_name": "API Test User",
    "email": "apitest1780472429@example.com",
    "role": "admin",
    "is_active": true,
    "created_at": "2026-06-03T07:40:29Z",
    "updated_at": "2026-06-03T07:40:29Z"
  }
}

```

**Duration:** 48ms

---

### 8. GET /api/admin/users/1

**Status Code:** `200`

**Response:**
```json
{
  "success": true,
  "message": "User berhasil diambil",
  "data": {
    "id": 1,
    "username": "testcli",
    "full_name": "Updated Name",
    "email": "updated@example.com",
    "role": "admin",
    "is_active": true,
    "last_login_at": "2026-06-03T14:40:29.2062808+07:00",
    "created_at": "2026-05-28T14:10:50.0340632+07:00",
    "updated_at": "2026-06-03T14:40:29.2062808+07:00"
  }
}

```

**Duration:** 1ms

---

### 9. PUT /api/admin/users/1

**Request Body:**
```json
{
  "email": "updated@example.com",
  "full_name": "Updated Name",
  "is_active": true,
  "role": "admin"
}
```

**Status Code:** `200`

**Response:**
```json
{
  "success": true,
  "message": "User berhasil diupdate"
}

```

**Duration:** 2ms

---

## API Consumers

### 10. POST /api/admin/consumers

**Request Body:**
```json
{
  "consumer_name": "apitest-1780472429",
  "contact_email": "apitest1780472429@app.com",
  "description": "API Test Application",
  "is_active": true
}
```

**Status Code:** `201`

**Response:**
```json
{
  "success": true,
  "message": "Consumer berhasil dibuat",
  "data": {
    "id": 6,
    "consumer_name": "apitest-1780472429",
    "description": "API Test Application",
    "contact_email": "apitest1780472429@app.com",
    "is_active": true,
    "created_at": "2026-06-03T07:40:29Z",
    "updated_at": "2026-06-03T07:40:29Z"
  }
}

```

**Duration:** 2ms

---

### 11. GET /api/admin/consumers

**Status Code:** `200`

**Response:**
```json
{
  "success": true,
  "message": "Daftar consumer berhasil diambil",
  "data": {
    "consumers": [
      {
        "id": 6,
        "consumer_name": "apitest-1780472429",
        "description": "API Test Application",
        "contact_email": "apitest1780472429@app.com",
        "is_active": true,
        "created_at": "2026-06-03T07:40:29Z",
        "updated_at": "2026-06-03T07:40:29Z"
      }
    ],
    "pagination": {
      "limit": 10,
      "page": 1,
      "total": 1
    }
  }
}

```

**Duration:** 1ms

---

### 12. GET /api/admin/consumers/6

**Status Code:** `200`

**Response:**
```json
{
  "success": true,
  "message": "Consumer berhasil diambil",
  "data": {
    "id": 6,
    "consumer_name": "apitest-1780472429",
    "description": "API Test Application",
    "contact_email": "apitest1780472429@app.com",
    "is_active": true,
    "created_at": "2026-06-03T07:40:29Z",
    "updated_at": "2026-06-03T07:40:29Z"
  }
}

```

**Duration:** 0s

---

### 13. PUT /api/admin/consumers/6

**Request Body:**
```json
{
  "contact_email": "updated@app.com",
  "description": "Updated App",
  "is_active": true
}
```

**Status Code:** `200`

**Response:**
```json
{
  "success": true,
  "message": "Consumer berhasil diupdate"
}

```

**Duration:** 3ms

---

### 14. DELETE /api/admin/consumers/6

**Status Code:** `200`

**Response:**
```json
{
  "success": true,
  "message": "Consumer berhasil dihapus"
}

```

**Duration:** 2ms

---

## Hosts

### 15. POST /api/admin/hosts

**Request Body:**
```json
{
  "description": "API Test Host",
  "host_name": "api1780472429.apitest.com",
  "is_active": true
}
```

**Status Code:** `201`

**Response:**
```json
{
  "success": true,
  "message": "Host berhasil dibuat",
  "data": {
    "id": 8,
    "host_name": "api1780472429.apitest.com",
    "description": "API Test Host",
    "is_active": true,
    "created_at": "2026-06-03T07:40:29Z",
    "updated_at": "2026-06-03T07:40:29Z"
  }
}

```

**Duration:** 2ms

---

### 16. GET /api/admin/hosts

**Status Code:** `200`

**Response:**
```json
{
  "success": true,
  "message": "Daftar host berhasil diambil",
  "data": {
    "hosts": [
      {
        "id": 8,
        "host_name": "api1780472429.apitest.com",
        "description": "API Test Host",
        "is_active": true,
        "created_at": "2026-06-03T07:40:29Z",
        "updated_at": "2026-06-03T07:40:29Z"
      },
      {
        "id": 3,
        "host_name": "api.test.com",
        "description": "Test Host",
        "is_active": true,
        "created_at": "2026-06-03T07:17:31Z",
        "updated_at": "2026-06-03T07:17:31Z"
      },
      {
        "id": 2,
        "host_name": "api-dua.example.local",
        "description": "Main API Host",
        "is_active": true,
        "created_at": "2026-05-28T10:11:55Z",
        "updated_at": "2026-05-28T10:11:55Z"
      }
    ],
    "pagination": {
      "limit": 10,
      "page": 1,
      "total": 3
    }
  }
}

```

**Duration:** 0s

---

### 17. GET /api/admin/hosts/8

**Status Code:** `200`

**Response:**
```json
{
  "success": true,
  "message": "Host berhasil diambil",
  "data": {
    "id": 8,
    "host_name": "api1780472429.apitest.com",
    "description": "API Test Host",
    "is_active": true,
    "created_at": "2026-06-03T07:40:29Z",
    "updated_at": "2026-06-03T07:40:29Z"
  }
}

```

**Duration:** 0s

---

### 18. PUT /api/admin/hosts/8

**Request Body:**
```json
{
  "description": "Updated Host",
  "is_active": true
}
```

**Status Code:** `200`

**Response:**
```json
{
  "success": true,
  "message": "Host berhasil diupdate"
}

```

**Duration:** 4ms

---

### 19. DELETE /api/admin/hosts/8

**Status Code:** `200`

**Response:**
```json
{
  "success": true,
  "message": "Host berhasil dihapus"
}

```

**Duration:** 3ms

---

## Virtual Hosts

### 20. POST /api/admin/virtual-hosts

**Request Body:**
```json
{
  "failover_mode": "active-active",
  "host_id": 1,
  "is_active": true,
  "lb_algorithm": "round_robin",
  "sticky_session": false,
  "vhost_name": "vhost1780472429.apitest.com"
}
```

**Status Code:** `201`

**Response:**
```json
{
  "success": true,
  "message": "Virtual host berhasil dibuat",
  "data": {
    "id": 6,
    "host_id": 1,
    "vhost_name": "vhost1780472429.apitest.com",
    "lb_algorithm": "round_robin",
    "sticky_session": false,
    "failover_mode": "active-active",
    "is_active": true,
    "created_at": "2026-06-03T07:40:29Z",
    "updated_at": "2026-06-03T07:40:29Z"
  }
}

```

**Duration:** 3ms

---

### 21. GET /api/admin/virtual-hosts

**Status Code:** `200`

**Response:**
```json
{
  "success": true,
  "message": "Daftar virtual host berhasil diambil",
  "data": {
    "pagination": {
      "limit": 10,
      "page": 1,
      "total": 6
    },
    "virtual_hosts": [
      {
        "id": 6,
        "host_id": 1,
        "vhost_name": "vhost1780472429.apitest.com",
        "lb_algorithm": "round_robin",
        "sticky_session": false,
        "failover_mode": "active-active",
        "is_active": true,
        "created_at": "2026-06-03T07:40:29Z",
        "updated_at": "2026-06-03T07:40:29Z"
      },
      {
        "id": 5,
        "host_id": 1,
        "vhost_name": "vhost1780472247.apitest.com",
        "lb_algorithm": "weighted_round_robin",
        "sticky_session": true,
        "failover_mode": "active-passive",
        "is_active": true,
        "created_at": "2026-06-03T07:37:27Z",
        "updated_at": "2026-06-03T14:37:27.6350168+07:00"
      },
      {
        "id": 4,
        "host_id": 1,
        "vhost_name": "vhost1780472241.apitest.com",
        "lb_algorithm": "weighted_round_robin",
        "sticky_session": true,
        "failover_mode": "active-passive",
        "is_active": true,
        "created_at": "2026-06-03T07:37:21Z",
        "updated_at": "2026-06-03T14:37:21.1726446+07:00"
      },
      {
        "id": 3,
        "host_id": 1,
        "vhost_name": "api.apitest.com",
        "lb_algorithm": "round_robin",
        "sticky_session": false,
        "failover_mode": "active-active",
        "is_active": true,
        "created_at": "2026-06-03T07:34:25Z",
        "updated_at": "2026-06-03T07:34:25Z"
      },
      {
        "id": 2,
        "host_id": 1,
        "vhost_name": "server-3-4.local",
        "lb_algorithm": "round_robin",
        "sticky_session": false,
        "failover_mode": "active-active",
        "is_active": true,
        "created_at": "2026-05-28T10:35:13Z",
        "updated_at": "2026-05-28T10:35:13Z"
      },
      {
        "id": 1,
        "host_id": 1,
        "vhost_name": "server-1-2.local",
        "lb_algorithm": "round_robin",
        "sticky_session": false,
        "failover_mode": "active-active",
        "is_active": true,
        "created_at": "2026-05-28T10:35:07Z",
        "updated_at": "2026-05-28T10:35:07Z"
      }
    ]
  }
}

```

**Duration:** 1ms

---

### 22. GET /api/admin/virtual-hosts/6

**Status Code:** `200`

**Response:**
```json
{
  "success": true,
  "message": "Virtual host berhasil diambil",
  "data": {
    "id": 6,
    "host_id": 1,
    "vhost_name": "vhost1780472429.apitest.com",
    "lb_algorithm": "round_robin",
    "sticky_session": false,
    "failover_mode": "active-active",
    "is_active": true,
    "created_at": "2026-06-03T07:40:29Z",
    "updated_at": "2026-06-03T07:40:29Z"
  }
}

```

**Duration:** 0s

---

### 23. PUT /api/admin/virtual-hosts/6

**Request Body:**
```json
{
  "failover_mode": "active-passive",
  "is_active": true,
  "lb_algorithm": "weighted_round_robin",
  "sticky_session": true
}
```

**Status Code:** `200`

**Response:**
```json
{
  "success": true,
  "message": "Virtual host berhasil diupdate"
}

```

**Duration:** 2ms

---

## Routes (Virtual Directories)

### 24. POST /api/admin/virtual-directories

**Request Body:**
```json
{
  "auth_type": "none",
  "cache_enabled": false,
  "cache_ttl_seconds": 0,
  "is_active": true,
  "match_type": "prefix",
  "max_request_size_mb": 10,
  "preserve_host_header": false,
  "proxy_timeout_seconds": 30,
  "retry_count": 2,
  "retry_delay_ms": 100,
  "source_path": "/api/v1",
  "strip_prefix": true,
  "target_path": "/",
  "virtual_host_id": 6,
  "websocket_enabled": false
}
```

**Status Code:** `201`

**Response:**
```json
{
  "success": true,
  "message": "Virtual directory berhasil dibuat",
  "data": {
    "id": 6,
    "virtual_host_id": 6,
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
    "cache_ttl_seconds": 0,
    "created_at": "2026-06-03T07:40:29Z",
    "updated_at": "2026-06-03T07:40:29Z"
  }
}

```

**Duration:** 3ms

---

### 25. GET /api/admin/virtual-directories

**Status Code:** `200`

**Response:**
```json
{
  "success": true,
  "message": "Daftar virtual directory berhasil diambil",
  "data": {
    "pagination": {
      "limit": 10,
      "page": 1,
      "total": 5
    },
    "virtual_directories": [
      {
        "id": 6,
        "virtual_host_id": 6,
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
        "cache_ttl_seconds": 0,
        "created_at": "2026-06-03T07:40:29Z",
        "updated_at": "2026-06-03T07:40:29Z",
        "vhost_name": "vhost1780472429.apitest.com"
      },
      {
        "id": 5,
        "virtual_host_id": 5,
        "source_path": "/api/v2",
        "target_path": "/v2",
        "match_type": "prefix",
        "strip_prefix": true,
        "preserve_host_header": false,
        "auth_type": "",
        "is_active": true,
        "proxy_timeout_seconds": 0,
        "retry_count": 0,
        "retry_delay_ms": 0,
        "max_request_size_mb": 0,
        "websocket_enabled": false,
        "cache_enabled": false,
        "cache_ttl_seconds": 0,
        "created_at": "2026-06-03T07:37:27Z",
        "updated_at": "2026-06-03T14:37:27.6424963+07:00",
        "vhost_name": "vhost1780472247.apitest.com"
      },
      {
        "id": 4,
        "virtual_host_id": 4,
        "source_path": "/api/v2",
        "target_path": "/v2",
        "match_type": "prefix",
        "strip_prefix": true,
        "preserve_host_header": false,
        "auth_type": "",
        "is_active": true,
        "proxy_timeout_seconds": 0,
        "retry_count": 0,
        "retry_delay_ms": 0,
        "max_request_size_mb": 0,
        "websocket_enabled": false,
        "cache_enabled": false,
        "cache_ttl_seconds": 0,
        "created_at": "2026-06-03T07:37:21Z",
        "updated_at": "2026-06-03T14:37:21.1793762+07:00",
        "vhost_name": "vhost1780472241.apitest.com"
      },
      {
        "id": 3,
        "virtual_host_id": 3,
        "source_path": "/api/v2",
        "target_path": "/v2",
        "match_type": "prefix",
        "strip_prefix": true,
        "preserve_host_header": false,
        "auth_type": "",
        "is_active": true,
        "proxy_timeout_seconds": 0,
        "retry_count": 0,
        "retry_delay_ms": 0,
        "max_request_size_mb": 0,
        "websocket_enabled": false,
        "cache_enabled": false,
        "cache_ttl_seconds": 0,
        "created_at": "2026-06-03T07:34:25Z",
        "updated_at": "2026-06-03T14:34:25.4943936+07:00",
        "vhost_name": "api.apitest.com"
      },
      {
        "id": 2,
        "virtual_host_id": 2,
        "source_path": "/jamet",
        "target_path": "/",
        "match_type": "rewrite",
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
        "cache_ttl_seconds": 0,
        "created_at": "2026-05-28T10:41:05Z",
        "updated_at": "2026-05-28T10:41:05Z",
        "vhost_name": "server-3-4.local"
      }
    ]
  }
}

```

**Duration:** 1ms

---

### 26. GET /api/admin/virtual-directories/6

**Status Code:** `200`

**Response:**
```json
{
  "success": true,
  "message": "Virtual directory berhasil diambil",
  "data": {
    "id": 6,
    "virtual_host_id": 6,
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
    "cache_ttl_seconds": 0,
    "created_at": "2026-06-03T07:40:29Z",
    "updated_at": "2026-06-03T07:40:29Z",
    "vhost_name": "vhost1780472429.apitest.com"
  }
}

```

**Duration:** 0s

---

### 27. PUT /api/admin/virtual-directories/6

**Request Body:**
```json
{
  "is_active": true,
  "match_type": "prefix",
  "source_path": "/api/v2",
  "strip_prefix": true,
  "target_path": "/v2"
}
```

**Status Code:** `200`

**Response:**
```json
{
  "success": true,
  "message": "Virtual directory berhasil diupdate"
}

```

**Duration:** 3ms

---

## Rate Limits

### 28. POST /api/admin/rate-limits

**Request Body:**
```json
{
  "block_duration_seconds": 60,
  "burst": 10,
  "is_active": true,
  "limit_by": "ip",
  "requests_per_minute": 60,
  "virtual_directory_id": 6
}
```

**Status Code:** `201`

**Response:**
```json
{
  "success": true,
  "message": "Rate limit berhasil dibuat",
  "data": {
    "id": 4,
    "virtual_directory_id": 6,
    "limit_by": "ip",
    "requests_per_minute": 60,
    "burst": 10,
    "block_duration_seconds": 60,
    "is_active": true,
    "created_at": "2026-06-03T07:40:29Z"
  }
}

```

**Duration:** 2ms

---

### 29. GET /api/admin/rate-limits

**Status Code:** `200`

**Response:**
```json
{
  "success": true,
  "message": "Daftar rate limit berhasil diambil",
  "data": {
    "pagination": {
      "limit": 10,
      "page": 1,
      "total": 1
    },
    "rate_limits": [
      {
        "id": 4,
        "virtual_directory_id": 6,
        "limit_by": "ip",
        "requests_per_minute": 60,
        "burst": 10,
        "block_duration_seconds": 60,
        "is_active": true,
        "created_at": "2026-06-03T07:40:29Z",
        "source_path": "/api/v2"
      }
    ]
  }
}

```

**Duration:** 1ms

---

### 30. GET /api/admin/rate-limits/4

**Status Code:** `200`

**Response:**
```json
{
  "success": true,
  "message": "Rate limit berhasil diambil",
  "data": {
    "id": 4,
    "virtual_directory_id": 6,
    "limit_by": "ip",
    "requests_per_minute": 60,
    "burst": 10,
    "block_duration_seconds": 60,
    "is_active": true,
    "created_at": "2026-06-03T07:40:29Z",
    "source_path": "/api/v2"
  }
}

```

**Duration:** 1ms

---

### 31. PUT /api/admin/rate-limits/4

**Request Body:**
```json
{
  "burst": 20,
  "is_active": true,
  "requests_per_minute": 120
}
```

**Status Code:** `200`

**Response:**
```json
{
  "success": true,
  "message": "Rate limit berhasil diupdate"
}

```

**Duration:** 2ms

---

### 32. DELETE /api/admin/rate-limits/4

**Status Code:** `200`

**Response:**
```json
{
  "success": true,
  "message": "Rate limit berhasil dihapus"
}

```

**Duration:** 3ms

---

## CORS Configuration

### 33. POST /api/admin/cors-configs

**Request Body:**
```json
{
  "allow_credentials": true,
  "allowed_headers": "Content-Type,Authorization",
  "allowed_methods": "GET,POST,PUT,DELETE,OPTIONS",
  "allowed_origins": "https://app.example.com",
  "is_active": true,
  "max_age_seconds": 3600,
  "virtual_directory_id": 6
}
```

**Status Code:** `201`

**Response:**
```json
{
  "success": true,
  "message": "CORS config berhasil dibuat",
  "data": {
    "id": 4,
    "virtual_directory_id": 6,
    "allowed_origins": "https://app.example.com",
    "allowed_methods": "GET,POST,PUT,DELETE,OPTIONS",
    "allowed_headers": "Content-Type,Authorization",
    "exposed_headers": "",
    "allow_credentials": true,
    "max_age_seconds": 3600,
    "is_active": true,
    "created_at": "2026-06-03T07:40:29Z"
  }
}

```

**Duration:** 2ms

---

### 34. GET /api/admin/cors-configs

**Status Code:** `200`

**Response:**
```json
{
  "success": true,
  "message": "Daftar CORS config berhasil diambil",
  "data": {
    "cors_configs": [
      {
        "id": 4,
        "virtual_directory_id": 6,
        "allowed_origins": "https://app.example.com",
        "allowed_methods": "GET,POST,PUT,DELETE,OPTIONS",
        "allowed_headers": "Content-Type,Authorization",
        "exposed_headers": "",
        "allow_credentials": true,
        "max_age_seconds": 3600,
        "is_active": true,
        "created_at": "2026-06-03T07:40:29Z",
        "source_path": "/api/v2"
      }
    ],
    "pagination": {
      "limit": 10,
      "page": 1,
      "total": 1
    }
  }
}

```

**Duration:** 1ms

---

### 35. GET /api/admin/cors-configs/4

**Status Code:** `200`

**Response:**
```json
{
  "success": true,
  "message": "CORS config berhasil diambil",
  "data": {
    "id": 4,
    "virtual_directory_id": 6,
    "allowed_origins": "https://app.example.com",
    "allowed_methods": "GET,POST,PUT,DELETE,OPTIONS",
    "allowed_headers": "Content-Type,Authorization",
    "exposed_headers": "",
    "allow_credentials": true,
    "max_age_seconds": 3600,
    "is_active": true,
    "created_at": "2026-06-03T07:40:29Z",
    "source_path": "/api/v2"
  }
}

```

**Duration:** 0s

---

### 36. PUT /api/admin/cors-configs/4

**Request Body:**
```json
{
  "allow_credentials": false,
  "allowed_origins": "*",
  "is_active": true
}
```

**Status Code:** `200`

**Response:**
```json
{
  "success": true,
  "message": "CORS config berhasil diupdate"
}

```

**Duration:** 2ms

---

### 37. DELETE /api/admin/cors-configs/4

**Status Code:** `200`

**Response:**
```json
{
  "success": true,
  "message": "CORS config berhasil dihapus"
}

```

**Duration:** 3ms

---

## Circuit Breakers

### 38. POST /api/admin/circuit-breakers

**Request Body:**
```json
{
  "enabled": true,
  "failure_threshold": 5,
  "half_open_max_requests": 3,
  "recovery_timeout_seconds": 30,
  "virtual_directory_id": 6
}
```

**Status Code:** `201`

**Response:**
```json
{
  "success": true,
  "message": "Circuit breaker berhasil dibuat",
  "data": {
    "id": 4,
    "virtual_directory_id": 6,
    "enabled": true,
    "failure_threshold": 5,
    "recovery_timeout_seconds": 30,
    "half_open_max_requests": 3,
    "created_at": "2026-06-03T07:40:29Z"
  }
}

```

**Duration:** 2ms

---

### 39. GET /api/admin/circuit-breakers

**Status Code:** `200`

**Response:**
```json
{
  "success": true,
  "message": "Daftar circuit breaker berhasil diambil",
  "data": {
    "circuit_breakers": [
      {
        "id": 4,
        "virtual_directory_id": 6,
        "enabled": true,
        "failure_threshold": 5,
        "recovery_timeout_seconds": 30,
        "half_open_max_requests": 3,
        "created_at": "2026-06-03T07:40:29Z",
        "source_path": "/api/v2"
      }
    ],
    "pagination": {
      "limit": 10,
      "page": 1,
      "total": 1
    }
  }
}

```

**Duration:** 1ms

---

### 40. GET /api/admin/circuit-breakers/4

**Status Code:** `200`

**Response:**
```json
{
  "success": true,
  "message": "Circuit breaker berhasil diambil",
  "data": {
    "id": 4,
    "virtual_directory_id": 6,
    "enabled": true,
    "failure_threshold": 5,
    "recovery_timeout_seconds": 30,
    "half_open_max_requests": 3,
    "created_at": "2026-06-03T07:40:29Z",
    "source_path": "/api/v2"
  }
}

```

**Duration:** 0s

---

### 41. PUT /api/admin/circuit-breakers/4

**Request Body:**
```json
{
  "enabled": true,
  "failure_threshold": 10,
  "half_open_max_requests": 5,
  "recovery_timeout_seconds": 60
}
```

**Status Code:** `200`

**Response:**
```json
{
  "success": true,
  "message": "Circuit breaker berhasil diupdate"
}

```

**Duration:** 2ms

---

### 42. DELETE /api/admin/circuit-breakers/4

**Status Code:** `200`

**Response:**
```json
{
  "success": true,
  "message": "Circuit breaker berhasil dihapus"
}

```

**Duration:** 2ms

---

