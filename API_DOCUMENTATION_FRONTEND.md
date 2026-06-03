# 📘 Swantara Gate API - Frontend Integration Guide

> **Generated from real API tests** - 42/42 endpoints tested successfully ✅  
> **Last Updated:** 2026-06-03  
> **Base URL:** `http://localhost:8081`  
> **Test Credentials:** `testcli` / `admin1324`

---

## 📋 Table of Contents

1. [Authentication](#1-authentication)
2. [Health Check](#2-health-check)
3. [Configuration](#3-configuration)
4. [Users Management](#4-users-management)
5. [API Consumers](#5-api-consumers)
6. [Hosts](#6-hosts)
7. [Virtual Hosts](#7-virtual-hosts)
8. [Virtual Directories (Routes)](#8-virtual-directories-routes)
9. [Rate Limits](#9-rate-limits)
10. [CORS Configs](#10-cors-configs)
11. [Circuit Breakers](#11-circuit-breakers)
12. [JavaScript/fetch Examples](#javascriptfetch-examples)
13. [Error Handling](#error-handling)
14. [Response Format Standard](#response-format-standard)

---

## 🔑 Authentication

### 1. Login

**Endpoint:** `POST /api/admin/auth/login`

**Request:**
```http
POST http://localhost:8081/api/admin/auth/login
Content-Type: application/json

{
  "username": "testcli",
  "password": "admin1324"
}
```

**Real Response (200 OK):**
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

**Frontend Usage:**
```javascript
// Store tokens after login
localStorage.setItem('access_token', response.data.access_token);
localStorage.setItem('refresh_token', response.data.refresh_token);
localStorage.setItem('token_expires', Date.now() + (response.data.expires_in * 1000));
```

---

### 2. Get Current User (Me)

**Endpoint:** `GET /api/admin/auth/me`

**Headers Required:**
```http
Authorization: Bearer {access_token}
```

**Real Response (200 OK):**
```json
{
  "success": true,
  "message": "Profil user",
  "data": {
    "user_id": 1,
    "username": "testcli",
    "role": "admin"
  }
}
```

---

### 3. Refresh Token

**Endpoint:** `POST /api/admin/auth/refresh`

**Request:**
```http
POST http://localhost:8081/api/admin/auth/refresh
Content-Type: application/json

{
  "refresh_token": "ffb85a023f89d25de3fe000cf2689728e1c3b4bf0a839a2040aa882ace0e234c"
}
```

**Real Response (200 OK):**
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

---

## 📊 1. Health Check

### Check API Status

**Endpoint:** `GET /api/health`

**Real Response (200 OK):**
```json
{
  "status": "ok",
  "message": "Swantara Gate Admin API is running"
}
```

---

## ⚙️ 2. Configuration

### Reload Proxy Configuration

**Endpoint:** `POST /api/admin/config/reload`

**Headers Required:**
```http
Authorization: Bearer {access_token}
```

**Real Response (200 OK):**
```json
{
  "status": "ok",
  "message": "Proxy configuration reloaded successfully"
}
```

---

## 👥 3. Users Management

### 1. Get All Users

**Endpoint:** `GET /api/admin/users`

**Query Parameters (optional):**
- `page` (default: 1)
- `limit` (default: 10)

**Real Response (200 OK):**
```json
{
  "success": true,
  "message": "Daftar user berhasil diambil",
  "data": {
    "users": [
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
    ],
    "pagination": {
      "page": 1,
      "limit": 10,
      "total": 5
    }
  }
}
```

---

### 2. Create User

**Endpoint:** `POST /api/admin/users`

**Request:**
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

**Real Response (201 Created):**
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

---

### 3. Get User by ID

**Endpoint:** `GET /api/admin/users/1`

**Real Response (200 OK):**
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

---

### 4. Update User

**Endpoint:** `PUT /api/admin/users/1`

**Request:**
```json
{
  "full_name": "Updated Name",
  "email": "updated@example.com",
  "role": "admin",
  "is_active": true
}
```

**Real Response (200 OK):**
```json
{
  "success": true,
  "message": "User berhasil diupdate"
}
```

---

## 🏢 4. API Consumers

### 1. Create Consumer

**Endpoint:** `POST /api/admin/consumers`

**Request:**
```json
{
  "consumer_name": "my-app",
  "description": "My Application",
  "contact_email": "dev@myapp.com",
  "is_active": true
}
```

**Real Response (201 Created):**
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

---

### 2. Get All Consumers

**Endpoint:** `GET /api/admin/consumers`

**Real Response (200 OK):**
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
      "page": 1,
      "limit": 10,
      "total": 1
    }
  }
}
```

---

### 3. Get Consumer by ID

**Endpoint:** `GET /api/admin/consumers/6`

**Real Response (200 OK):**
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

---

### 4. Update Consumer

**Endpoint:** `PUT /api/admin/consumers/6`

**Request:**
```json
{
  "description": "Updated App",
  "contact_email": "updated@app.com",
  "is_active": true
}
```

**Real Response (200 OK):**
```json
{
  "success": true,
  "message": "Consumer berhasil diupdate"
}
```

---

### 5. Delete Consumer

**Endpoint:** `DELETE /api/admin/consumers/6`

**Real Response (200 OK):**
```json
{
  "success": true,
  "message": "Consumer berhasil dihapus"
}
```

---

## 🌐 5. Hosts

### 1. Create Host

**Endpoint:** `POST /api/admin/hosts`

**Request:**
```json
{
  "host_name": "api.example.com",
  "description": "Main API Host",
  "is_active": true
}
```

**Real Response (201 Created):**
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

---

### 2. Get All Hosts

**Endpoint:** `GET /api/admin/hosts`

**Real Response (200 OK):**
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
      }
    ],
    "pagination": {
      "page": 1,
      "limit": 10,
      "total": 3
    }
  }
}
```

---

### 3. Update Host

**Endpoint:** `PUT /api/admin/hosts/8`

**Request:**
```json
{
  "description": "Updated Host",
  "is_active": true
}
```

**Real Response (200 OK):**
```json
{
  "success": true,
  "message": "Host berhasil diupdate"
}
```

---

### 4. Delete Host

**Endpoint:** `DELETE /api/admin/hosts/8`

**Real Response (200 OK):**
```json
{
  "success": true,
  "message": "Host berhasil dihapus"
}
```

---

## 🔀 6. Virtual Hosts

### 1. Create Virtual Host

**Endpoint:** `POST /api/admin/virtual-hosts`

**Request:**
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

**Real Response (201 Created):**
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

---

### 2. Get All Virtual Hosts

**Endpoint:** `GET /api/admin/virtual-hosts`

**Real Response (200 OK):**
```json
{
  "success": true,
  "message": "Daftar virtual host berhasil diambil",
  "data": {
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
      }
    ],
    "pagination": {
      "page": 1,
      "limit": 10,
      "total": 6
    }
  }
}
```

---

### 3. Update Virtual Host

**Endpoint:** `PUT /api/admin/virtual-hosts/6`

**Request:**
```json
{
  "lb_algorithm": "weighted_round_robin",
  "sticky_session": true,
  "failover_mode": "active-passive",
  "is_active": true
}
```

**Real Response (200 OK):**
```json
{
  "success": true,
  "message": "Virtual host berhasil diupdate"
}
```

---

## 🛣️ 7. Virtual Directories (Routes)

### 1. Create Virtual Directory

**Endpoint:** `POST /api/admin/virtual-directories`

**Request:**
```json
{
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
  "cache_ttl_seconds": 0
}
```

**Real Response (201 Created):**
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

---

### 2. Get All Virtual Directories

**Endpoint:** `GET /api/admin/virtual-directories`

**Real Response (200 OK):**
```json
{
  "success": true,
  "message": "Daftar virtual directory berhasil diambil",
  "data": {
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
      }
    ],
    "pagination": {
      "page": 1,
      "limit": 10,
      "total": 5
    }
  }
}
```

---

### 3. Update Virtual Directory

**Endpoint:** `PUT /api/admin/virtual-directories/6`

**Request:**
```json
{
  "source_path": "/api/v2",
  "target_path": "/v2",
  "match_type": "prefix",
  "strip_prefix": true,
  "is_active": true
}
```

**Real Response (200 OK):**
```json
{
  "success": true,
  "message": "Virtual directory berhasil diupdate"
}
```

---

## 📊 8. Rate Limits

### 1. Create Rate Limit

**Endpoint:** `POST /api/admin/rate-limits`

**Request:**
```json
{
  "virtual_directory_id": 6,
  "limit_by": "ip",
  "requests_per_minute": 60,
  "burst": 10,
  "block_duration_seconds": 60,
  "is_active": true
}
```

**Real Response (201 Created):**
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

---

### 2. Update Rate Limit

**Endpoint:** `PUT /api/admin/rate-limits/4`

**Request:**
```json
{
  "requests_per_minute": 120,
  "burst": 20,
  "is_active": true
}
```

**Real Response (200 OK):**
```json
{
  "success": true,
  "message": "Rate limit berhasil diupdate"
}
```

---

### 3. Delete Rate Limit

**Endpoint:** `DELETE /api/admin/rate-limits/4`

**Real Response (200 OK):**
```json
{
  "success": true,
  "message": "Rate limit berhasil dihapus"
}
```

---

## 🔗 9. CORS Configs

### 1. Create CORS Config

**Endpoint:** `POST /api/admin/cors-configs`

**Request:**
```json
{
  "virtual_directory_id": 6,
  "allowed_origins": "https://app.example.com",
  "allowed_methods": "GET,POST,PUT,DELETE,OPTIONS",
  "allowed_headers": "Content-Type,Authorization",
  "allow_credentials": true,
  "max_age_seconds": 3600,
  "is_active": true
}
```

**Real Response (201 Created):**
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

---

### 2. Update CORS Config

**Endpoint:** `PUT /api/admin/cors-configs/4`

**Request:**
```json
{
  "allowed_origins": "*",
  "allow_credentials": false,
  "is_active": true
}
```

**Real Response (200 OK):**
```json
{
  "success": true,
  "message": "CORS config berhasil diupdate"
}
```

---

## ⚡ 10. Circuit Breakers

### 1. Create Circuit Breaker

**Endpoint:** `POST /api/admin/circuit-breakers`

**Request:**
```json
{
  "virtual_directory_id": 6,
  "enabled": true,
  "failure_threshold": 5,
  "recovery_timeout_seconds": 30,
  "half_open_max_requests": 3
}
```

**Real Response (201 Created):**
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

---

### 2. Update Circuit Breaker

**Endpoint:** `PUT /api/admin/circuit-breakers/4`

**Request:**
```json
{
  "enabled": true,
  "failure_threshold": 10,
  "recovery_timeout_seconds": 60,
  "half_open_max_requests": 5
}
```

**Real Response (200 OK):**
```json
{
  "success": true,
  "message": "Circuit breaker berhasil diupdate"
}
```

---

## JavaScript/fetch Examples

### Complete Authentication Flow

```javascript
class SwantaraGateAPI {
  constructor(baseURL = 'http://localhost:8081') {
    this.baseURL = baseURL;
    this.accessToken = localStorage.getItem('access_token');
    this.refreshToken = localStorage.getItem('refresh_token');
  }

  // Login and store tokens
  async login(username, password) {
    const response = await fetch(`${this.baseURL}/api/admin/auth/login`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ username, password })
    });
    
    const data = await response.json();
    
    if (data.success) {
      this.accessToken = data.data.access_token;
      this.refreshToken = data.data.refresh_token;
      
      localStorage.setItem('access_token', this.accessToken);
      localStorage.setItem('refresh_token', this.refreshToken);
      localStorage.setItem('token_expires', Date.now() + (data.data.expires_in * 1000));
    }
    
    return data;
  }

  // Make authenticated request
  async request(method, endpoint, body = null) {
    const headers = {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${this.accessToken}`
    };

    let response = await fetch(`${this.baseURL}${endpoint}`, {
      method,
      headers,
      body: body ? JSON.stringify(body) : null
    });

    // If token expired, try refresh
    if (response.status === 401) {
      await this.refreshAccessToken();
      
      // Retry with new token
      headers['Authorization'] = `Bearer ${this.accessToken}`;
      response = await fetch(`${this.baseURL}${endpoint}`, {
        method,
        headers,
        body: body ? JSON.stringify(body) : null
      });
    }

    return await response.json();
  }

  // Refresh access token
  async refreshAccessToken() {
    const response = await fetch(`${this.baseURL}/api/admin/auth/refresh`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ refresh_token: this.refreshToken })
    });

    const data = await response.json();
    
    if (data.success) {
      this.accessToken = data.data.access_token;
      this.refreshToken = data.data.refresh_token;
      
      localStorage.setItem('access_token', this.accessToken);
      localStorage.setItem('refresh_token', this.refreshToken);
    }
  }

  // CRUD Examples
  async getUsers() {
    return this.request('GET', '/api/admin/users');
  }

  async createUser(userData) {
    return this.request('POST', '/api/admin/users', userData);
  }

  async updateUser(userId, userData) {
    return this.request('PUT', `/api/admin/users/${userId}`, userData);
  }
}

// Usage
const api = new SwantaraGateAPI();

// Login
await api.login('testcli', 'admin1324');

// Get users
const users = await api.getUsers();
console.log(users);

// Create user
const newUser = await api.createUser({
  username: 'john',
  password: 'secret123',
  full_name: 'John Doe',
  email: 'john@example.com',
  role: 'admin',
  is_active: true
});
```

---

## Error Handling

### Common Error Responses

**401 Unauthorized:**
```json
{
  "success": false,
  "message": "Token tidak ditemukan",
  "error": "Token tidak ditemukan"
}
```

**400 Bad Request:**
```json
{
  "success": false,
  "message": "username sudah digunakan",
  "error": "username sudah digunakan"
}
```

**404 Not Found:**
```json
{
  "success": false,
  "message": "User tidak ditemukan",
  "error": "User tidak ditemukan"
}
```

### Error Handling in JavaScript

```javascript
async function safeRequest(method, endpoint, body = null) {
  try {
    const response = await api.request(method, endpoint, body);
    
    if (!response.success) {
      console.error('API Error:', response.message);
      // Show user-friendly error message
      alert(response.message);
      return null;
    }
    
    return response;
  } catch (error) {
    console.error('Request failed:', error);
    alert('Network error. Please check your connection.');
    return null;
  }
}
```

---

## Response Format Standard

All API responses follow this structure:

### Success Response
```json
{
  "success": true,
  "message": "Deskripsi operasi berhasil",
  "data": {
    // Response data here
  }
}
```

### Error Response
```json
{
  "success": false,
  "message": "Deskripsi error",
  "error": "Detail error teknis"
}
```

### Pagination Response
```json
{
  "success": true,
  "message": "Daftar data berhasil diambil",
  "data": {
    "items": [],
    "pagination": {
      "page": 1,
      "limit": 10,
      "total": 100
    }
  }
}
```

---

## 📝 Quick Reference

| Endpoint | Method | Auth Required | Description |
|----------|--------|---------------|-------------|
| `/api/health` | GET | No | Health check |
| `/api/admin/auth/login` | POST | No | Login |
| `/api/admin/auth/me` | GET | Yes | Get current user |
| `/api/admin/auth/refresh` | POST | No | Refresh token |
| `/api/admin/config/reload` | POST | Yes | Reload config |
| `/api/admin/users` | GET | Yes | List users |
| `/api/admin/users` | POST | Yes | Create user |
| `/api/admin/users/:id` | GET | Yes | Get user |
| `/api/admin/users/:id` | PUT | Yes | Update user |
| `/api/admin/consumers` | GET/POST | Yes | List/Create consumers |
| `/api/admin/consumers/:id` | GET/PUT/DELETE | Yes | Manage consumer |
| `/api/admin/hosts` | GET/POST | Yes | List/Create hosts |
| `/api/admin/hosts/:id` | GET/PUT/DELETE | Yes | Manage host |
| `/api/admin/virtual-hosts` | GET/POST | Yes | List/Create vhosts |
| `/api/admin/virtual-hosts/:id` | GET/PUT | Yes | Manage vhost |
| `/api/admin/virtual-directories` | GET/POST | Yes | List/Create routes |
| `/api/admin/virtual-directories/:id` | GET/PUT | Yes | Manage route |
| `/api/admin/rate-limits` | GET/POST | Yes | List/Create rate limits |
| `/api/admin/rate-limits/:id` | GET/PUT/DELETE | Yes | Manage rate limit |
| `/api/admin/cors-configs` | GET/POST | Yes | List/Create CORS |
| `/api/admin/cors-configs/:id` | GET/PUT/DELETE | Yes | Manage CORS |
| `/api/admin/circuit-breakers` | GET/POST | Yes | List/Create circuit breakers |
| `/api/admin/circuit-breakers/:id` | GET/PUT/DELETE | Yes | Manage circuit breaker |

---

## 🎯 Frontend Integration Checklist

- [ ] Store `access_token` and `refresh_token` in localStorage
- [ ] Add `Authorization: Bearer {token}` header to all requests
- [ ] Implement token refresh on 401 errors
- [ ] Handle pagination with `page` and `limit` parameters
- [ ] Show loading states during API calls
- [ ] Display error messages from `response.message`
- [ ] Validate form data before sending to API
- [ ] Use proper HTTP methods (GET/POST/PUT/DELETE)
- [ ] Set `Content-Type: application/json` header
- [ ] Implement proper error handling with try/catch

---

**Generated from real API test results** - All examples are from actual API responses ✅
