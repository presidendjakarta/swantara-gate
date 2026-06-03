# Swantara Gate Admin UI

## Struktur Folder

```
internal/webui/
├── webui.go          # Handler untuk serve web files (embedded + dev mode)
└── static/           # Folder untuk admin template MetroUI + jQuery
    ├── login.html    # Halaman login (contoh)
    ├── dashboard.html # Halaman dashboard (contoh)
    ├── css/          # MetroUI CSS + custom styles
    ├── js/           # MetroUI JS + jQuery + custom scripts
    ├── img/          # Images & icons
    └── ...           # File template lainnya
```

## Cara Pakai

### 1. Copy Template MetroUI Kamu

Copy semua file template MetroUI kamu ke folder:
```
internal/webui/static/
```

Struktur yang diharapkan:
```
internal/webui/static/
├── index.html         # Landing page / redirect ke login
├── login.html         # Halaman login
├── dashboard.html     # Dashboard utama
├── css/
│   ├── metro-all.min.css
│   └── custom.css
├── js/
│   ├── jquery-3.6.0.min.js
│   ├── metro.min.js
│   └── app.js          # Custom JavaScript
└── img/
    └── logo.png
```

### 2. Development Mode

Set `DEV_MODE=true` di `.env`:
```env
DEV_MODE=true
```

Dengan dev mode, web files dibaca dari folder `./web/static/` (bukan dari embedded).
Jadi kamu bisa edit HTML/CSS/JS tanpa recompile binary.

### 3. Production Mode

Set `DEV_MODE=false` di `.env`:
```env
DEV_MODE=false
```

Semua web files akan di-embed ke dalam binary saat build.

### 4. Akses Admin Panel

Setelah server running:
```
http://localhost:9090
```

## Integrasi dengan API

Admin UI otomatis terintegrasi dengan Admin API di port yang sama (9090).

### Contoh API Call dengan jQuery:

```javascript
const API_BASE_URL = 'http://localhost:9090';

// Login
$.ajax({
    url: API_BASE_URL + '/api/admin/auth/login',
    method: 'POST',
    contentType: 'application/json',
    data: JSON.stringify({
        username: 'admin',
        password: 'admin123'
    }),
    success: function(response) {
        localStorage.setItem('access_token', response.data.access_token);
        window.location.href = '/dashboard.html';
    }
});

// Get data dengan auth
$.ajax({
    url: API_BASE_URL + '/api/admin/hosts',
    method: 'GET',
    headers: {
        'Authorization': 'Bearer ' + localStorage.getItem('access_token')
    },
    success: function(response) {
        console.log(response.data);
    }
});
```

## Endpoint API yang Tersedia

| Endpoint | Method | Deskripsi |
|----------|--------|-----------|
| `/api/admin/auth/login` | POST | Login |
| `/api/admin/auth/me` | GET | Get current user |
| `/api/admin/hosts` | GET/POST | CRUD Hosts |
| `/api/admin/virtual-hosts` | GET/POST | CRUD Virtual Hosts |
| `/api/admin/upstream-servers` | GET/POST | CRUD Upstream Servers |
| `/api/admin/virtual-directories` | GET/POST | CRUD Routes |
| `/api/admin/consumers` | GET/POST | CRUD API Consumers |
| `/api/admin/config/reload` | POST | Reload proxy config |
| `/api/health` | GET | Health check |

## Environment Variables

| Variable | Default | Deskripsi |
|----------|---------|-----------|
| `DEV_MODE` | `true` | Development mode (read from disk) |
| `ADMIN_HTTP_PORT` | `9090` | Admin panel port |
| `PROXY_HTTP_PORT` | `8000` | Proxy gateway port |

## Build untuk Production

```bash
# Set dev mode false
export DEV_MODE=false

# Build binary (semua web files akan di-embed)
go build -o swantara-gate.exe ./cmd/server

# Run
./swantara-gate.exe
```
