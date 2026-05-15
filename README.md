# Swantara Gate - API Gateway Proxy

Aplikasi API Gateway Proxy yang dibangun menggunakan Go (Golang) dengan SQLite sebagai database.

**Repository:** https://github.com/presidendjakarta/swantara-gate

## 🚀 Fitur Utama

### Admin Panel
- ✅ Manajemen User (CRUD dengan role-based access)
- ✅ Manajemen API Consumers
- ✅ Manajemen Hosts & Virtual Hosts
- ✅ Load Balancing Configuration
- ✅ Multi-port support (HTTP & HTTPS)

### Coming Soon
- ⏳ Virtual Directories (API Routes)
- ⏳ Authentication (JWT, API Key, Basic Auth)
- ⏳ Rate Limiting
- ⏳ Circuit Breaker
- ⏳ SSL/TLS Management
- ⏳ Proxy Server
- ⏳ Health Check
- ⏳ CORS Configuration

## 📋 Requirements

- Go 1.21 atau lebih tinggi
- SQLite (pure Go, tidak perlu install SQLite)
- Postman (untuk testing API)

## 🛠️ Instalasi

### 1. Clone Repository
```bash
cd x:\laragon\go-apps\swantara-gate
```

### 2. Install Dependencies
```bash
go mod tidy
```

### 3. Konfigurasi Environment
File `.env` sudah tersedia dengan konfigurasi default:

```env
# Admin Panel Ports
ADMIN_HTTP_PORT=8080
ADMIN_HTTPS_PORT=8443

# Proxy Gateway Ports
PROXY_HTTP_PORT=8000
PROXY_HTTPS_PORT=8440

# Database
DATABASE_PATH=./data/database.db
DATABASE_SQL_PATH=./data/database.sql
```

### 4. Jalankan Aplikasi
```bash
go run cmd/server/main.go
```

Aplikasi akan berjalan di:
- Admin HTTP: `http://localhost:8080`
- Admin HTTPS: `https://localhost:8443` (butuh SSL cert)
- Proxy HTTP: `http://localhost:8000` (coming soon)
- Proxy HTTPS: `https://localhost:8440` (coming soon)

## 📚 Struktur Project

```
swantara-gate/
├── cmd/
│   └── server/
│       └── main.go              # Entry point aplikasi
├── internal/
│   ├── config/
│   │   └── config.go            # Konfigurasi aplikasi
│   ├── database/
│   │   └── database.go          # Database connection & migration
│   ├── model/
│   │   ├── user.go              # User model
│   │   ├── api_consumer.go      # API Consumer model
│   │   ├── host.go              # Host model
│   │   └── virtual_host.go      # Virtual Host model
│   ├── repository/
│   │   ├── user_repository.go
│   │   ├── api_consumer_repository.go
│   │   ├── host_repository.go
│   │   └── virtual_host_repository.go
│   ├── service/
│   │   ├── user_service.go
│   │   ├── api_consumer_service.go
│   │   └── host_service.go
│   ├── handler/
│   │   └── admin_handler.go     # HTTP handlers
│   ├── middleware/
│   │   └── middleware.go        # Logging & CORS middleware
│   └── response/
│       └── response.go          # Response helpers
├── data/
│   ├── database.db              # SQLite database file
│   └── database.sql             # Database schema
├── docs/
│   ├── DATABASE_DOCUMENTATION.md
│   └── API_DOCUMENTATION.md
├── .env                         # Environment variables
├── go.mod
├── go.sum
└── README.md
```

## 🧪 Testing dengan Postman

### 1. Health Check
```bash
GET http://localhost:8080/api/health
```

**Expected Response:**
```json
{
  "status": "ok",
  "message": "Swantara Gate Admin API is running"
}
```

### 2. Create User
```bash
POST http://localhost:8080/api/admin/users
Content-Type: application/json

{
  "username": "admin",
  "password": "admin123",
  "full_name": "Super Admin",
  "email": "admin@swantara.com",
  "role": "super_admin",
  "is_active": true
}
```

### 3. Get All Users
```bash
GET http://localhost:8080/api/admin/users?page=1&limit=10
```

### 4. Create API Consumer
```bash
POST http://localhost:8080/api/admin/consumers
Content-Type: application/json

{
  "consumer_name": "MyApp",
  "description": "My mobile application",
  "contact_email": "dev@myapp.com",
  "is_active": true
}
```

### 5. Create Host
```bash
POST http://localhost:8080/api/admin/hosts
Content-Type: application/json

{
  "host_name": "api.example.com",
  "description": "Main API host",
  "is_active": true
}
```

### 6. Create Virtual Host
```bash
POST http://localhost:8080/api/admin/virtual-hosts
Content-Type: application/json

{
  "host_id": 1,
  "vhost_name": "v1.api.example.com",
  "lb_algorithm": "round_robin",
  "sticky_session": false,
  "failover_mode": "active-active",
  "is_active": true
}
```

## 📖 API Documentation

Dokumentasi lengkap API tersedia di:
- [API Documentation](docs/API_DOCUMENTATION.md)
- [Database Documentation](docs/DATABASE_DOCUMENTATION.md)

## 🔐 Security Features

- ✅ Password hashing dengan bcrypt
- ✅ SQLite pure Go (no CGO/gcc required)
- ✅ CORS middleware
- ✅ Request logging
- ✅ Input validation

## 🏗️ Architecture

Aplikasi menggunakan **Clean Architecture** dengan lapisan:

1. **Handler/Controller** - Menangani HTTP request/response
2. **Service** - Business logic dan validasi
3. **Repository** - Database operations
4. **Model** - Data structures

## 📝 Database

Menggunakan **SQLite** dengan driver pure Go (`modernc.org/sqlite`):
- ✅ Tidak perlu CGO
- ✅ Tidak perlu install GCC
- ✅ Cross-platform compatible
- ✅ Zero configuration

## 🔧 Configuration

Semua konfigurasi dapat diubah melalui file `.env`:

| Variable | Default | Description |
|----------|---------|-------------|
| ADMIN_HTTP_PORT | 8080 | Port Admin HTTP |
| ADMIN_HTTPS_PORT | 8443 | Port Admin HTTPS |
| PROXY_HTTP_PORT | 8000 | Port Proxy HTTP |
| PROXY_HTTPS_PORT | 8440 | Port Proxy HTTPS |
| DATABASE_PATH | ./data/database.db | Path database |
| APP_ENV | development | Environment (development/production) |
| LOG_LEVEL | info | Log level (debug/info/warn/error) |

## 🚧 Next Steps

Yang akan diimplementasikan selanjutnya:

1. **Virtual Directories CRUD** - API routes management
2. **Authentication** - JWT, API Key, Basic Auth
3. **Proxy Server** - Reverse proxy dengan load balancing
4. **Rate Limiting** - Request throttling
5. **Circuit Breaker** - Fault tolerance
6. **SSL/TLS** - HTTPS dengan Let's Encrypt
7. **Health Check** - Backend monitoring
8. **Dashboard UI** - Admin panel frontend

## 🤝 Contributing

Untuk menambahkan fitur baru:

1. Buat branch fitur baru
2. Commit changes
3. Push ke branch
4. Buat Pull Request

## 📄 License

Project ini dibuat untuk tujuan pembelajaran.

## 👥 Author

Swantara Gate API Gateway

---

**Status Development:** ✅ Admin API CRUD已基本完成，可以测试！
