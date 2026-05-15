# API Documentation - Swantara Gate Admin Panel

Base URL: `http://localhost:8080`

## Health Check

### GET /api/health
Cek apakah API berjalan dengan baik.

**Response:**
```json
{
  "status": "ok",
  "message": "Swantara Gate Admin API is running"
}
```

---

## 1. Users Management

### 1.1 Create User
**POST** `/api/admin/users`

Membuat user admin baru.

**Request Body:**
```json
{
  "username": "admin1",
  "password": "password123",
  "full_name": "Admin Utama",
  "email": "admin@example.com",
  "role": "admin",
  "is_active": true
}
```

**Roles yang tersedia:**
- `super_admin`
- `admin`
- `operator`
- `viewer`

**Response (201 Created):**
```json
{
  "success": true,
  "message": "User berhasil dibuat",
  "data": {
    "id": 1,
    "username": "admin1",
    "full_name": "Admin Utama",
    "email": "admin@example.com",
    "role": "admin",
    "is_active": true,
    "created_at": "2026-05-15T15:30:00Z",
    "updated_at": "2026-05-15T15:30:00Z"
  }
}
```

---

### 1.2 Get All Users
**GET** `/api/admin/users`

Mengambil semua user dengan pagination.

**Query Parameters:**
- `page` (optional): Nomor halaman (default: 1)
- `limit` (optional): Jumlah data per halaman (default: 10, max: 100)

**Example:** `GET /api/admin/users?page=1&limit=10`

**Response:**
```json
{
  "success": true,
  "message": "Daftar user berhasil diambil",
  "data": {
    "users": [
      {
        "id": 1,
        "username": "admin1",
        "full_name": "Admin Utama",
        "email": "admin@example.com",
        "role": "admin",
        "is_active": true,
        "last_login_at": null,
        "created_at": "2026-05-15T15:30:00Z",
        "updated_at": "2026-05-15T15:30:00Z"
      }
    ],
    "pagination": {
      "page": 1,
      "limit": 10,
      "total": 1
    }
  }
}
```

---

### 1.3 Get User by ID
**GET** `/api/admin/users/{id}`

Mengambil user berdasarkan ID.

**Example:** `GET /api/admin/users/1`

**Response:**
```json
{
  "success": true,
  "message": "User berhasil diambil",
  "data": {
    "id": 1,
    "username": "admin1",
    "full_name": "Admin Utama",
    "email": "admin@example.com",
    "role": "admin",
    "is_active": true,
    "last_login_at": null,
    "created_at": "2026-05-15T15:30:00Z",
    "updated_at": "2026-05-15T15:30:00Z"
  }
}
```

---

### 1.4 Update User
**PUT** `/api/admin/users/{id}`

Memperbarui data user.

**Request Body:**
```json
{
  "full_name": "Admin Baru",
  "email": "newadmin@example.com",
  "role": "super_admin",
  "is_active": true
}
```

**Response:**
```json
{
  "success": true,
  "message": "User berhasil diupdate",
  "data": null
}
```

---

### 1.5 Delete User
**DELETE** `/api/admin/users/{id}`

Menghapus user.

**Example:** `DELETE /api/admin/users/1`

**Response:**
```json
{
  "success": true,
  "message": "User berhasil dihapus",
  "data": null
}
```

---

## 2. API Consumers Management

### 2.1 Create Consumer
**POST** `/api/admin/consumers`

Membuat consumer/aplikasi baru.

**Request Body:**
```json
{
  "consumer_name": "MyApp",
  "description": "Aplikasi mobile saya",
  "contact_email": "dev@myapp.com",
  "is_active": true
}
```

**Response (201 Created):**
```json
{
  "success": true,
  "message": "Consumer berhasil dibuat",
  "data": {
    "id": 1,
    "consumer_name": "MyApp",
    "description": "Aplikasi mobile saya",
    "contact_email": "dev@myapp.com",
    "is_active": true,
    "created_at": "2026-05-15T15:30:00Z",
    "updated_at": "2026-05-15T15:30:00Z"
  }
}
```

---

### 2.2 Get All Consumers
**GET** `/api/admin/consumers`

Mengambil semua consumer dengan pagination.

**Query Parameters:**
- `page` (optional): Nomor halaman (default: 1)
- `limit` (optional): Jumlah data per halaman (default: 10)

**Example:** `GET /api/admin/consumers?page=1&limit=10`

**Response:**
```json
{
  "success": true,
  "message": "Daftar consumer berhasil diambil",
  "data": {
    "consumers": [
      {
        "id": 1,
        "consumer_name": "MyApp",
        "description": "Aplikasi mobile saya",
        "contact_email": "dev@myapp.com",
        "is_active": true,
        "created_at": "2026-05-15T15:30:00Z",
        "updated_at": "2026-05-15T15:30:00Z"
      }
    ],
    "pagination": {
      "page": 1,
      "limit": 10,
      "total": 1
    }
  }
}
```

---

### 2.3 Get Consumer by ID
**GET** `/api/admin/consumers/{id}`

**Example:** `GET /api/admin/consumers/1`

**Response:**
```json
{
  "success": true,
  "message": "Consumer berhasil diambil",
  "data": {
    "id": 1,
    "consumer_name": "MyApp",
    "description": "Aplikasi mobile saya",
    "contact_email": "dev@myapp.com",
    "is_active": true,
    "created_at": "2026-05-15T15:30:00Z",
    "updated_at": "2026-05-15T15:30:00Z"
  }
}
```

---

### 2.4 Update Consumer
**PUT** `/api/admin/consumers/{id}`

**Request Body:**
```json
{
  "description": "Deskripsi baru",
  "contact_email": "newemail@myapp.com",
  "is_active": true
}
```

**Response:**
```json
{
  "success": true,
  "message": "Consumer berhasil diupdate",
  "data": null
}
```

---

### 2.5 Delete Consumer
**DELETE** `/api/admin/consumers/{id}`

**Response:**
```json
{
  "success": true,
  "message": "Consumer berhasil dihapus",
  "data": null
}
```

---

## 3. Hosts Management

### 3.1 Create Host
**POST** `/api/admin/hosts`

Membuat host baru.

**Request Body:**
```json
{
  "host_name": "api.example.com",
  "description": "Main API host",
  "is_active": true
}
```

**Response (201 Created):**
```json
{
  "success": true,
  "message": "Host berhasil dibuat",
  "data": {
    "id": 1,
    "host_name": "api.example.com",
    "description": "Main API host",
    "is_active": true,
    "created_at": "2026-05-15T15:30:00Z",
    "updated_at": "2026-05-15T15:30:00Z"
  }
}
```

---

### 3.2 Get All Hosts
**GET** `/api/admin/hosts`

**Query Parameters:**
- `page` (optional)
- `limit` (optional)

**Response:**
```json
{
  "success": true,
  "message": "Daftar host berhasil diambil",
  "data": {
    "hosts": [...],
    "pagination": {
      "page": 1,
      "limit": 10,
      "total": 1
    }
  }
}
```

---

### 3.3 Get Host by ID
**GET** `/api/admin/hosts/{id}`

**Example:** `GET /api/admin/hosts/1`

---

### 3.4 Update Host
**PUT** `/api/admin/hosts/{id}`

**Request Body:**
```json
{
  "description": "Updated description",
  "is_active": true
}
```

---

### 3.5 Delete Host
**DELETE** `/api/admin/hosts/{id}`

---

## 4. Virtual Hosts Management

### 4.1 Create Virtual Host
**POST** `/api/admin/virtual-hosts`

Membuat virtual host baru.

**Request Body:**
```json
{
  "host_id": 1,
  "vhost_name": "v1.api.example.com",
  "lb_algorithm": "round_robin",
  "sticky_session": false,
  "failover_mode": "active-active",
  "is_active": true
}
```

**Load Balancing Algorithms:**
- `round_robin`
- `weighted_round_robin`
- `least_conn`
- `ip_hash`
- `random`
- `failover`

**Failover Modes:**
- `active-active`
- `active-passive`

**Response (201 Created):**
```json
{
  "success": true,
  "message": "Virtual host berhasil dibuat",
  "data": {
    "id": 1,
    "host_id": 1,
    "vhost_name": "v1.api.example.com",
    "lb_algorithm": "round_robin",
    "sticky_session": false,
    "failover_mode": "active-active",
    "is_active": true,
    "created_at": "2026-05-15T15:30:00Z",
    "updated_at": "2026-05-15T15:30:00Z",
    "host_name": "api.example.com"
  }
}
```

---

### 4.2 Get All Virtual Hosts
**GET** `/api/admin/virtual-hosts`

**Query Parameters:**
- `page` (optional)
- `limit` (optional)

---

### 4.3 Get Virtual Host by ID
**GET** `/api/admin/virtual-hosts/{id}`

**Example:** `GET /api/admin/virtual-hosts/1`

---

### 4.4 Update Virtual Host
**PUT** `/api/admin/virtual-hosts/{id}`

**Request Body:**
```json
{
  "lb_algorithm": "least_conn",
  "sticky_session": true,
  "failover_mode": "active-passive",
  "is_active": true
}
```

---

### 4.5 Delete Virtual Host
**DELETE** `/api/admin/virtual-hosts/{id}`

---

## Error Responses

### Bad Request (400)
```json
{
  "success": false,
  "message": "Username wajib diisi",
  "error": "Username wajib diisi"
}
```

### Not Found (404)
```json
{
  "success": false,
  "message": "User tidak ditemukan",
  "error": "User tidak ditemukan"
}
```

### Conflict (409)
```json
{
  "success": false,
  "message": "Username sudah digunakan",
  "error": "Username sudah digunakan"
}
```

### Internal Server Error (500)
```json
{
  "success": false,
  "message": "Gagal mengambil daftar user",
  "error": "Gagal mengambil daftar user"
}
```

---

## Testing dengan Postman

### 1. Import Collection
Buat collection baru di Postman dengan nama "Swantara Gate API"

### 2. Setup Environment Variable
- Base URL: `http://localhost:8080`

### 3. Test Sequence

**Test 1: Health Check**
```
GET http://localhost:8080/api/health
```

**Test 2: Create User**
```
POST http://localhost:8080/api/admin/users
Body (JSON):
{
  "username": "admin",
  "password": "admin123",
  "full_name": "Super Admin",
  "email": "admin@swantara.com",
  "role": "super_admin",
  "is_active": true
}
```

**Test 3: Get All Users**
```
GET http://localhost:8080/api/admin/users
```

**Test 4: Create Consumer**
```
POST http://localhost:8080/api/admin/consumers
Body (JSON):
{
  "consumer_name": "TestApp",
  "description": "Test application",
  "contact_email": "test@example.com",
  "is_active": true
}
```

**Test 5: Create Host**
```
POST http://localhost:8080/api/admin/hosts
Body (JSON):
{
  "host_name": "api.test.com",
  "description": "Test API host",
  "is_active": true
}
```

**Test 6: Create Virtual Host**
```
POST http://localhost:8080/api/admin/virtual-hosts
Body (JSON):
{
  "host_id": 1,
  "vhost_name": "v1.api.test.com",
  "lb_algorithm": "round_robin",
  "sticky_session": false,
  "failover_mode": "active-active",
  "is_active": true
}
```

---

## Catatan

- Semua response menggunakan format JSON standar
- Password otomatis di-hash menggunakan bcrypt
- Field `password_hash` tidak akan pernah dikirim di response
- Pagination tersedia untuk semua endpoint GET list
- Soft delete menggunakan field `is_active`
- Timestamp menggunakan format RFC3339
