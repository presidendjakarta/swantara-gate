
## External Auth

### 61. POST /api/admin/external-auth

#### 📤 Request

**Headers:**

| Header | Value |
|--------|-------|
| Authorization | Bearer *** |

**Request Body:**

```json
{
  "auth_url": "https://auth.example.com/verify",
  "http_method": "POST",
  "is_active": true,
  "message_key": "message",
  "request_timeout_seconds": 5,
  "send_body": false,
  "send_headers": true,
  "success_key": "status",
  "success_value": "true",
  "virtual_directory_id": 9
}
```

#### 📥 Response

**Status Code:** `201` ✅

**Response Body:**

```json
{
  "success": true,
  "message": "External auth berhasil dibuat",
  "data": {
    "id": 2,
    "virtual_directory_id": 9,
    "auth_url": "https://auth.example.com/verify",
    "http_method": "POST",
    "request_timeout_seconds": 5,
    "send_headers": true,
    "send_body": false,
    "success_key": "status",
    "success_value": "true",
    "message_key": "message",
    "token_key": "",
    "is_active": true,
    "created_at": "2026-06-04T07:48:58Z"
  }
}

```

**Duration:** 3ms

---

### 62. GET /api/admin/external-auth

#### 📤 Request

**Headers:**

| Header | Value |
|--------|-------|
| Authorization | Bearer *** |

**Request Body:** None

#### 📥 Response

**Status Code:** `200` ✅

**Response Body:**

```json
{
  "success": true,
  "message": "Daftar external auth berhasil diambil",
  "data": {
    "external_auth": [
      {
        "id": 2,
        "virtual_directory_id": 9,
        "auth_url": "https://auth.example.com/verify",
        "http_method": "POST",
        "request_timeout_seconds": 5,
        "send_headers": true,
        "send_body": false,
        "success_key": "status",
        "success_value": "true",
        "message_key": "message",
        "token_key": "",
        "is_active": true,
        "created_at": "2026-06-04T07:48:58Z",
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

### 63. GET /api/admin/external-auth/2

#### 📤 Request

**Headers:**

| Header | Value |
|--------|-------|
| Authorization | Bearer *** |

**Request Body:** None

#### 📥 Response

**Status Code:** `200` ✅

**Response Body:**

```json
{
  "success": true,
  "message": "External auth berhasil diambil",
  "data": {
    "id": 2,
    "virtual_directory_id": 9,
    "auth_url": "https://auth.example.com/verify",
    "http_method": "POST",
    "request_timeout_seconds": 5,
    "send_headers": true,
    "send_body": false,
    "success_key": "status",
    "success_value": "true",
    "message_key": "message",
    "token_key": "",
    "is_active": true,
    "created_at": "2026-06-04T07:48:58Z",
    "source_path": "/api/v2"
  }
}

```

**Duration:** 0s

---

### 64. PUT /api/admin/external-auth/2

#### 📤 Request

**Headers:**

| Header | Value |
|--------|-------|
| Authorization | Bearer *** |

**Request Body:**

```json
{
  "auth_url": "https://auth.example.com/v2/verify",
  "is_active": true,
  "request_timeout_seconds": 10
}
```

#### 📥 Response

**Status Code:** `200` ✅

**Response Body:**

```json
{
  "success": true,
  "message": "External auth berhasil diupdate"
}

```

**Duration:** 3ms

---

### 65. DELETE /api/admin/external-auth/2

#### 📤 Request

**Headers:**

| Header | Value |
|--------|-------|
| Authorization | Bearer *** |

**Request Body:** None

#### 📥 Response

**Status Code:** `200` ✅

**Response Body:**

```json
{
  "success": true,
  "message": "External auth berhasil dihapus"
}

```

**Duration:** 1ms

---
