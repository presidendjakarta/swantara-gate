# Panduan Alur Implementasi Swantara Gate

Dokumen ini menjelaskan **urutan langkah** yang harus dilakukan untuk men-setup API Gateway dari nol sampai traffic bisa berjalan.

---

## Apa Itu Swantara Gate?

Swantara Gate adalah **penjaga gerbang** antara pengguna (client) dan server backend Anda.

Bayangkan seperti **satpam komplek perumahan**:
- Semua tamu (request) harus lewat satpam dulu
- Satpam cek identitas (authentication)
- Satpam cek apakah tamu diizinkan masuk (authorization)
- Satpam arahkan tamu ke rumah yang benar (routing)
- Kalau terlalu banyak tamu sekaligus, satpam batasi antrian (rate limiting)

---

## Penjelasan Istilah (Glosarium)

### Host
**Apa:** Nama domain yang dikelola gateway.  
**Analogi:** Nama komplek perumahan.  
**Contoh:** `api.tokoku.com`, `backend.myapp.id`

> **Multi-Host:** Satu Swantara Gate bisa mengelola **banyak domain sekaligus**.  
> Seperti satu perusahaan security yang menjaga banyak komplek perumahan berbeda.

---

### Virtual Host
**Apa:** Konfigurasi routing untuk satu domain, termasuk cara membagi traffic ke beberapa server.  
**Analogi:** Pos satpam di pintu masuk komplek — dia tahu rumah mana saja yang ada di dalam dan bagaimana mengarahkan tamu.  
**Contoh:** Domain `api.tokoku.com` pakai strategi round-robin (giliran) ke 3 server backend.

**Hubungan:** Satu Host bisa punya satu atau lebih Virtual Host.

---

### Upstream Server
**Apa:** Server backend (tujuan akhir) yang benar-benar memproses request.  
**Analogi:** Rumah-rumah di dalam komplek yang menerima tamu.  
**Contoh:** Server di `192.168.1.10:8080` dan `192.168.1.11:8080`

**Hubungan:** Satu Virtual Host bisa punya banyak Upstream Server (untuk load balancing).

---

### Virtual Directory (Route)
**Apa:** Aturan path/URL mana yang diteruskan ke backend.  
**Analogi:** Alamat jalan di dalam komplek. "Kalau tamu mau ke Jl. Mawar, arahkan ke blok A."  
**Contoh:** Path `/api/v1/products` diteruskan ke backend di `/products`

**Hubungan:** Satu Virtual Host bisa punya banyak Virtual Directory (route yang berbeda-beda).

**Match Type (Cara Mencocokkan Path):**

| Match Type | Keterangan | Contoh Pattern | Yang Match |
|---|---|---|---|
| **`exact`** | Harus persis sama | `/api/health` | Hanya `/api/health` |
| **`prefix`** | Path dimulai dengan pattern ini | `/api` | `/api`, `/api/users`, `/api/users/123` |
| **`wildcard`** | Support `*` (1 segment) dan `**` (banyak segment) | `/api/*/detail` | `/api/users/detail` (tapi tidak `/api/users/123/detail`). `/api/**` match semua di bawah `/api/` |
| **`regex`** | Pakai regular expression | `/api/v[0-9]+/.*` | `/api/v1/users`, `/api/v2/orders` |
| **`parameter`** | Path dengan parameter `{nama}` | `/users/{id}/posts/{post_id}` | `/users/42/posts/7` (extract: `id=42`, `post_id=7`) |
| **`rewrite`** | Transformasi path dengan parameter mapping | `/blog/{id}` → `/artikel/{id}` | `/blog/99` → backend: `/artikel/99` |

**Contoh Rewrite:**
```
source_path: /blog/{id}           →  target_path: /artikel/{id}
source_path: /project/{nama}/{id} →  target_path: /project/detail/{id}/{nama}
source_path: /user/{username}     →  target_path: /profile/{username}/view
```

**Perbedaan `parameter` vs `rewrite`:**
- `parameter`: Extract params tapi target path tetap (`/users/{id}` → `/api/users`)
- `rewrite`: Transformasi path lengkap dengan params (`/users/{id}` → `/api/user/{id}`)

**Prioritas matching** (dari yang paling diprioritaskan): `exact` → `regex` → `parameter` → `rewrite` → `prefix` → `wildcard`

**Strip Prefix:** Kalau `strip_prefix = true`, maka pattern di Source Path akan dihilangkan sebelum dikirim ke backend.
> Contoh: `source_path=/api`, `target_path=/`, path masuk `/api/users` → dikirim ke backend sebagai `/users`

---

### API Consumer
**Apa:** Aplikasi/pihak ketiga yang menggunakan API Anda.  
**Analogi:** Perusahaan ekspedisi yang sering kirim kurir ke komplek Anda — mereka terdaftar resmi.  
**Contoh:** "Mobile App iOS", "Partner Tokopedia", "Internal Dashboard"

**Hubungan:** Consumer bisa punya API Key atau Credential untuk membuktikan identitasnya.

---

### API Key
**Apa:** Kunci unik yang diberikan ke consumer untuk akses API.  
**Analogi:** Kartu akses/pass yang diberikan ke kurir ekspedisi agar bisa masuk komplek.  
**Contoh:** `sk-abc123xyz789...`

**Hubungan:** Satu Consumer bisa punya banyak API Key (misal: key production, key staging).

---

### Consumer Credential
**Apa:** Username + password untuk consumer (Basic Auth).  
**Analogi:** Kartu identitas dengan foto yang harus ditunjukkan ke satpam.  

---

### Route Consumer Access (ACL)
**Apa:** Aturan "consumer mana boleh akses route mana".  
**Analogi:** Daftar tamu VIP — "Kurir JNE hanya boleh ke blok A, Kurir TIKI boleh ke blok A dan B."  

**Hubungan:** Menghubungkan Consumer dengan Virtual Directory.

---

### JWT Config
**Apa:** Konfigurasi validasi token JWT di route tertentu.  
**Analogi:** Mesin scanner kartu magnetik di pintu — hanya kartu yang valid yang bisa masuk.  

---

### Rate Limit
**Apa:** Batasan jumlah request per waktu.  
**Analogi:** "Maksimal 60 tamu per menit. Lebih dari itu, tunggu di luar."  

---

### Circuit Breaker
**Apa:** Pemutus otomatis jika backend bermasalah.  
**Analogi:** Kalau rumah tujuan sudah 5x tidak buka pintu, satpam langsung bilang "rumah sedang tutup" tanpa perlu ketuk lagi.  

---

### IP Whitelist / Blacklist
**Apa:** Daftar IP yang diizinkan (whitelist) atau diblokir (blacklist).  
**Analogi:** Whitelist = "Hanya mobil plat B1234XY yang boleh masuk." Blacklist = "Mobil plat D5678ZZ dilarang masuk."  

---

### CORS Config
**Apa:** Aturan siapa (domain frontend) yang boleh mengakses API dari browser.  
**Analogi:** "Hanya website resmi tokoku.com yang boleh panggil API ini dari browser pengunjung."  

---

### Maintenance Window
**Apa:** Jadwal maintenance dimana gateway otomatis tolak semua request.  
**Analogi:** "Komplek tutup untuk fumigasi tanggal 1 Juni jam 00:00-04:00."  

---

## Hubungan Antar Komponen (Diagram)

```
┌─────────────────────────────────────────────────────────┐
│                    SWANTARA GATE                         │
│            (1 instance, banyak domain)                   │
├─────────────────────────────────────────────────────────┤
│                                                         │
│  HOST 1 (domain: api.tokoku.com)                        │
│   └── VIRTUAL HOST (round_robin)                        │
│        ├── UPSTREAM: 10.0.0.1:8080                      │
│        ├── UPSTREAM: 10.0.0.2:8080                      │
│        ├── ROUTE: /products (API Key)                   │
│        └── ROUTE: /orders   (JWT)                       │
│                                                         │
│  HOST 2 (domain: admin.tokoku.com)                      │
│   └── VIRTUAL HOST (failover)                           │
│        ├── UPSTREAM: 10.0.1.1:9000                      │
│        ├── ROUTE: /dashboard (JWT + IP Whitelist)       │
│        └── ROUTE: /settings  (JWT)                      │
│                                                         │
│  HOST 3 (domain: partner.example.com)                   │
│   └── VIRTUAL HOST (weighted_round_robin)               │
│        ├── UPSTREAM: 10.0.2.1:8080 (weight: 7)          │
│        ├── UPSTREAM: 10.0.2.2:8080 (weight: 3)          │
│        └── ROUTE: /api/v1 (API Key + Rate Limit)        │
│                                                         │
│  CONSUMER: "Mobile App"                                 │
│   ├── API Key: sk-mobile-xxx                            │
│   └── Access: api.tokoku.com/products                   │
│                                                         │
│  CONSUMER: "Partner Ekspedisi"                          │
│   ├── API Key: sk-partner-yyy                           │
│   └── Access: partner.example.com/api/v1                │
│                                                         │
└─────────────────────────────────────────────────────────┘
```

> **Catatan:** Satu gateway bisa handle puluhan bahkan ratusan domain.
> Proxy engine otomatis mengarahkan request berdasarkan hostname yang masuk.

---

## Analogi Lengkap: Komplek Perumahan

| Komponen | Analogi | Fungsi |
|----------|---------|--------|
| **Host** | Nama komplek | Identitas domain |
| **Virtual Host** | Pos satpam | Terima & arahkan request |
| **Upstream Server** | Rumah di dalam | Server yang memproses |
| **Virtual Directory** | Alamat jalan | URL path routing |
| **Consumer** | Perusahaan ekspedisi | Pihak yang akses API |
| **API Key** | Kartu akses | Bukti izin masuk |
| **Rate Limit** | Aturan antrian | Batasi jumlah per waktu |
| **Circuit Breaker** | "Rumah tutup" otomatis | Stop jika backend error |
| **IP Whitelist** | Plat nomor diizinkan | Hanya IP tertentu |
| **IP Blacklist** | Plat nomor dilarang | Blokir IP tertentu |
| **CORS** | Daftar website resmi | Browser mana yang boleh |
| **Maintenance** | Komplek tutup sementara | Jadwal downtime |

---

## Arsitektur Singkat

```
Client Request
      │
      ▼
┌─────────────────────────────────────────────┐
│           SWANTARA GATE (Proxy)              │
│                                             │
│  Host ─► Virtual Host ─► Virtual Directory  │
│                              │              │
│              ┌───────────────┼──────────┐   │
│              │               │          │   │
│          IP Filter     Rate Limit   Auth    │
│          CORS          Circuit Brk  ACL     │
│              │               │          │   │
│              └───────────────┼──────────┘   │
│                              ▼              │
│                      Upstream Servers       │
│                     (Load Balanced)         │
└─────────────────────────────────────────────┘
      │
      ▼
Backend Service (target_host:target_port)
```

---

## Hirarki Data (Wajib Dipahami)

```
Host (domain induk)
 └── Virtual Host (domain + config LB)
      ├── Upstream Server 1 (backend server)
      ├── Upstream Server 2 (backend server)
      └── Virtual Directory (route/path)
           ├── Methods (GET, POST, dll)
           ├── JWT Config / External Auth
           ├── API Key + Consumer Access
           ├── Rate Limit
           ├── CORS Config
           ├── Circuit Breaker
           ├── IP Whitelist / Blacklist
           ├── Request Header Rules
           ├── Response Header Rules
           └── Query Rewrites
```

**Aturan penting:** Buat dari ATAS ke BAWAH. Tidak bisa buat Virtual Host tanpa Host, tidak bisa buat Route tanpa Virtual Host.

---

## Use Case 1: Setup Proxy Sederhana (Tanpa Auth)

> **Skenario:** Saya punya backend di `localhost:3000`, ingin diakses melalui gateway di `api.myapp.com/v1/*`

### Langkah:

```
1. Login
2. Buat Host           → "api.myapp.com"
3. Buat Virtual Host   → link ke Host, set LB algorithm
4. Buat Upstream Server → target: localhost:3000
5. Buat Virtual Directory → source: /v1, target: /, auth: none
6. Set Methods         → GET, POST, PUT, DELETE
```

### Detail Postman:

| # | Endpoint | Body Penting |
|---|----------|-------------|
| 1 | POST `/api/admin/auth/login` | username, password |
| 2 | POST `/api/admin/hosts` | `host_name: "api.myapp.com"` |
| 3 | POST `/api/admin/virtual-hosts` | `host_id: 1, vhost_name: "api.myapp.com", lb_algorithm: "round_robin"` |
| 4 | POST `/api/admin/upstream-servers` | `virtual_host_id: 1, target_host: "127.0.0.1", target_port: 3000` |
| 5 | POST `/api/admin/virtual-directories` | `virtual_host_id: 1, source_path: "/v1", target_path: "/", auth_type: "none"` |
| 6 | PUT `/api/admin/virtual-directories/1/methods` | `methods: ["GET","POST","PUT","DELETE"]` |

**Hasil:** Request ke `http://api.myapp.com:8000/v1/users` → diproxy ke `http://localhost:3000/users`

---

## Use Case 2: Proxy dengan API Key Authentication

> **Skenario:** Route `/v1/private/*` hanya bisa diakses dengan API Key valid.

### Langkah:

```
1. (Host & VHost sudah ada dari Use Case 1)
2. Buat API Consumer    → "mobile-app"
3. Buat API Key         → untuk consumer "mobile-app"
4. Buat Virtual Directory → source: /v1/private, auth_type: "api_key"
5. Buat Route Consumer Access → link directory + consumer
```

### Detail Postman:

| # | Endpoint | Body Penting |
|---|----------|-------------|
| 1 | POST `/api/admin/consumers` | `consumer_name: "mobile-app"` |
| 2 | POST `/api/admin/api-keys` | `consumer_id: 1, description: "Mobile App Key"` |
| 3 | POST `/api/admin/virtual-directories` | `source_path: "/v1/private", auth_type: "api_key"` |
| 4 | POST `/api/admin/route-access` | `virtual_directory_id: 2, consumer_id: 1` |

**Hasil:** Client harus kirim header `X-API-Key: <key_value>` untuk akses `/v1/private/*`

---

## Use Case 3: Proxy dengan JWT Authentication

> **Skenario:** Route `/v1/secure/*` hanya bisa diakses dengan JWT token yang valid.

### Langkah:

```
1. Buat Virtual Directory → source: /v1/secure, auth_type: "jwt"
2. Buat JWT Config       → link ke directory, set secret & algorithm
```

### Detail Postman:

| # | Endpoint | Body Penting |
|---|----------|-------------|
| 1 | POST `/api/admin/virtual-directories` | `source_path: "/v1/secure", auth_type: "jwt"` |
| 2 | POST `/api/admin/jwt-configs` | `virtual_directory_id: X, algorithm: "HS256", jwt_secret: "your-secret"` |

**Hasil:** Client harus kirim `Authorization: Bearer <jwt_token>` yang ditandatangani dengan secret yang sama.

---

## Use Case 4: Proxy dengan Basic Auth

> **Skenario:** Route tertentu dilindungi Basic Authentication.

### Langkah:

```
1. Buat Consumer
2. Buat Consumer Credential → auth_type: "basic", username, password
3. Buat Virtual Directory    → auth_type: "basic"
4. Buat Route Consumer Access
```

### Detail Postman:

| # | Endpoint | Body Penting |
|---|----------|-------------|
| 1 | POST `/api/admin/consumers` | `consumer_name: "partner-app"` |
| 2 | POST `/api/admin/consumer-credentials` | `consumer_id: X, auth_type: "basic", username: "user", password: "pass"` |
| 3 | POST `/api/admin/virtual-directories` | `source_path: "/v1/partner", auth_type: "basic"` |
| 4 | POST `/api/admin/route-access` | `virtual_directory_id: X, consumer_id: X` |

**Hasil:** Client harus kirim `Authorization: Basic <base64(user:pass)>`

---

## Use Case 5: Rate Limiting

> **Skenario:** Batasi `/v1/public/*` maksimal 60 request/menit per IP.

### Langkah:

```
1. (Virtual Directory sudah ada)
2. Buat Rate Limit → link ke directory
```

### Detail Postman:

| # | Endpoint | Body Penting |
|---|----------|-------------|
| 1 | POST `/api/admin/rate-limits` | `virtual_directory_id: X, limit_by: "ip", requests_per_minute: 60, burst: 10, block_duration_seconds: 60` |

**Hasil:** IP yang melebihi 60 req/menit akan diblokir selama 60 detik (return 429).

---

## Use Case 6: CORS Configuration

> **Skenario:** Frontend di `https://app.mysite.com` butuh akses ke API.

### Langkah:

```
1. (Virtual Directory sudah ada)
2. Buat CORS Config → link ke directory
```

### Detail Postman:

| # | Endpoint | Body Penting |
|---|----------|-------------|
| 1 | POST `/api/admin/cors-configs` | `virtual_directory_id: X, allowed_origins: "https://app.mysite.com", allowed_methods: "GET,POST,PUT,DELETE", allow_credentials: true` |

**Hasil:** Preflight request (OPTIONS) dijawab dengan CORS headers yang benar.

---

## Use Case 7: IP Whitelist (Hanya IP Tertentu)

> **Skenario:** Route admin internal hanya boleh diakses dari jaringan kantor.

### Langkah:

```
1. Buat Virtual Directory → /internal/admin
2. Buat IP Whitelist → CIDR 192.168.1.0/24
```

### Detail Postman:

| # | Endpoint | Body Penting |
|---|----------|-------------|
| 1 | POST `/api/admin/virtual-directories` | `source_path: "/internal/admin"` |
| 2 | POST `/api/admin/ip-whitelists` | `virtual_directory_id: X, ip_address: "192.168.1.0/24", description: "Office network"` |

**Hasil:** Hanya IP dari range 192.168.1.x yang bisa akses. Lainnya kena 403.

---

## Use Case 8: IP Blacklist (Blokir IP Tertentu)

> **Skenario:** Blokir IP yang terdeteksi melakukan serangan.

### Detail Postman:

| # | Endpoint | Body Penting |
|---|----------|-------------|
| 1 | POST `/api/admin/ip-blacklists` | `virtual_directory_id: X, ip_address: "203.0.113.50", reason: "DDoS attack", expired_at: "2027-01-01 00:00:00"` |

**Hasil:** IP tersebut diblokir sampai tanggal expired.

---

## Use Case 9: Circuit Breaker

> **Skenario:** Jika backend error 5x berturut-turut, stop forwarding selama 30 detik.

### Detail Postman:

| # | Endpoint | Body Penting |
|---|----------|-------------|
| 1 | POST `/api/admin/circuit-breakers` | `virtual_directory_id: X, enabled: true, failure_threshold: 5, recovery_timeout_seconds: 30, half_open_max_requests: 3` |

**Hasil:** Setelah 5 error, circuit "open" → return 503 langsung tanpa hit backend.

---

## Use Case 10: Load Balancing Multi-Server

> **Skenario:** 3 backend server, distribusi traffic merata.

### Langkah:

```
1. Buat Virtual Host → lb_algorithm: "round_robin"
2. Buat 3 Upstream Server → masing-masing dengan target berbeda
```

### Detail Postman:

| # | Endpoint | Body Penting |
|---|----------|-------------|
| 1 | POST `/api/admin/virtual-hosts` | `lb_algorithm: "round_robin"` |
| 2 | POST `/api/admin/upstream-servers` | `target_host: "10.0.0.1", target_port: 8080, weight: 1` |
| 3 | POST `/api/admin/upstream-servers` | `target_host: "10.0.0.2", target_port: 8080, weight: 1` |
| 4 | POST `/api/admin/upstream-servers` | `target_host: "10.0.0.3", target_port: 8080, weight: 1` |

**Hasil:** Traffic didistribusi: Server1 → Server2 → Server3 → Server1 → ...

---

## Use Case 11: Weighted Load Balancing

> **Skenario:** Server A lebih powerful, beri 70% traffic. Server B backup 30%.

### Detail Postman:

| # | Endpoint | Body Penting |
|---|----------|-------------|
| 1 | POST `/api/admin/virtual-hosts` | `lb_algorithm: "weighted_round_robin"` |
| 2 | POST `/api/admin/upstream-servers` | `target_host: "10.0.0.1", weight: 7` |
| 3 | POST `/api/admin/upstream-servers` | `target_host: "10.0.0.2", weight: 3` |

---

## Use Case 12: Header Manipulation

> **Skenario:** Tambah header `X-Request-ID` sebelum dikirim ke backend, hapus `X-Powered-By` dari response.

### Detail Postman:

| # | Endpoint | Body Penting |
|---|----------|-------------|
| 1 | POST `/api/admin/request-headers` | `virtual_directory_id: X, header_name: "X-Request-ID", operation: "set", value_source: "static", header_value: "gateway-123"` |
| 2 | POST `/api/admin/response-headers` | `virtual_directory_id: X, header_name: "X-Powered-By", operation: "delete"` |

---

## Use Case 13: Maintenance Mode

> **Skenario:** Jadwalkan maintenance 1 Juni 2026 pukul 00:00-04:00.

### Detail Postman:

| # | Endpoint | Body Penting |
|---|----------|-------------|
| 1 | POST `/api/admin/maintenance-windows` | `virtual_host_id: 1, title: "Scheduled Maintenance", start_at: "2026-06-01 00:00:00", end_at: "2026-06-01 04:00:00", maintenance_response_code: 503, maintenance_message: "Sedang maintenance"` |

**Hasil:** Selama jam tersebut, semua request ke virtual host return 503 + pesan maintenance.

---

## Use Case 14: External Auth (Forward Auth)

> **Skenario:** Validasi token ke auth service eksternal sebelum forward.

### Detail Postman:

| # | Endpoint | Body Penting |
|---|----------|-------------|
| 1 | POST `/api/admin/virtual-directories` | `auth_type: "external"` |
| 2 | POST `/api/admin/external-auth` | `virtual_directory_id: X, auth_url: "https://auth.myapp.com/verify", http_method: "POST", success_key: "valid", success_value: "true"` |

**Hasil:** Setiap request, gateway POST ke auth service dulu. Kalau response `{"valid": "true"}`, request diteruskan.

---

## Use Case 15: Multi-Host (Banyak Domain dalam 1 Gateway)

> **Skenario:** Perusahaan punya 3 domain berbeda: `api.tokoku.com`, `admin.tokoku.com`, dan `partner.example.com`. Semua mau dikelola 1 Swantara Gate.

### Langkah:

```
1. Login
2. Buat Host 1 → "api.tokoku.com"
3. Buat Host 2 → "admin.tokoku.com"
4. Buat Host 3 → "partner.example.com"
5. Buat Virtual Host per domain (masing-masing bisa LB algorithm berbeda)
6. Buat Upstream Server per virtual host (bisa server berbeda)
7. Buat Virtual Directory per virtual host (route masing-masing)
```

### Detail Postman:

| # | Endpoint | Body Penting |
|---|----------|-------------|
| 1 | POST `/api/admin/hosts` | `host_name: "api.tokoku.com"` |
| 2 | POST `/api/admin/hosts` | `host_name: "admin.tokoku.com"` |
| 3 | POST `/api/admin/hosts` | `host_name: "partner.example.com"` |
| 4 | POST `/api/admin/virtual-hosts` | `host_id: 1, vhost_name: "api.tokoku.com", lb_algorithm: "round_robin"` |
| 5 | POST `/api/admin/virtual-hosts` | `host_id: 2, vhost_name: "admin.tokoku.com", lb_algorithm: "failover"` |
| 6 | POST `/api/admin/virtual-hosts` | `host_id: 3, vhost_name: "partner.example.com", lb_algorithm: "weighted_round_robin"` |
| 7 | POST `/api/admin/upstream-servers` | `virtual_host_id: 1, target_host: "10.0.0.1", target_port: 8080` |
| 8 | POST `/api/admin/upstream-servers` | `virtual_host_id: 2, target_host: "10.0.1.1", target_port: 9000` |
| 9 | POST `/api/admin/upstream-servers` | `virtual_host_id: 3, target_host: "10.0.2.1", target_port: 8080, weight: 7` |
| 10 | POST `/api/admin/upstream-servers` | `virtual_host_id: 3, target_host: "10.0.2.2", target_port: 8080, weight: 3` |

**Hasil:**
- Request ke `api.tokoku.com:8000/*` → diarahkan ke server 10.0.0.1
- Request ke `admin.tokoku.com:8000/*` → diarahkan ke server 10.0.1.1
- Request ke `partner.example.com:8000/*` → dibagi 70/30 ke 2 server

> **Penting:** Setiap domain punya konfigurasi keamanan, rate limit, dan upstream sendiri-sendiri. Tidak saling pengaruh.

---

## Urutan Setup Lengkap (Full Stack)

Jika ingin setup gateway lengkap dengan semua fitur, ikuti urutan ini:

```
┌─────────────────────────────────────────────────┐
│ TAHAP 1: INFRASTRUKTUR DASAR                    │
├─────────────────────────────────────────────────┤
│ 1. Login / Create User (CLI atau API)           │
│ 2. Buat Host                                    │
│ 3. Buat Virtual Host (pilih LB algorithm)       │
│ 4. Buat Upstream Server(s)                      │
│ 5. Buat Virtual Directory (route)               │
│ 6. Set Allowed Methods                          │
└─────────────────────────────────────────────────┘
            │
            ▼ (Traffic sudah bisa jalan tanpa auth)
┌─────────────────────────────────────────────────┐
│ TAHAP 2: SECURITY (Opsional per route)          │
├─────────────────────────────────────────────────┤
│ 7. Buat Consumer (jika pakai API Key/Basic)     │
│ 8. Buat API Key / Credential                    │
│ 9. Buat Route Consumer Access (ACL)             │
│ 10. Set auth_type di Virtual Directory          │
│ 11. Buat JWT Config / External Auth             │
└─────────────────────────────────────────────────┘
            │
            ▼
┌─────────────────────────────────────────────────┐
│ TAHAP 3: PROTECTION (Opsional per route)        │
├─────────────────────────────────────────────────┤
│ 12. Buat Rate Limit                             │
│ 13. Buat IP Whitelist / Blacklist               │
│ 14. Buat Circuit Breaker                        │
│ 15. Buat CORS Config                            │
└─────────────────────────────────────────────────┘
            │
            ▼
┌─────────────────────────────────────────────────┐
│ TAHAP 4: FINE-TUNING (Opsional)                 │
├─────────────────────────────────────────────────┤
│ 16. Request Header Rules                        │
│ 17. Response Header Rules                       │
│ 18. Query Rewrites                              │
│ 19. Maintenance Windows                         │
└─────────────────────────────────────────────────┘
```

---

## Quick Start (Minimum Viable Proxy)

Kalau buru-buru dan cuma mau proxy basic:

```bash
# 1. Buat user admin (dari terminal)
./swantara-cli create-user -username=admin -password=admin123 -role=super_admin

# 2. Jalankan server
go run ./cmd/server/
```

Lalu di Postman, jalankan 6 request ini berurutan:

1. **Login** → dapat token
2. **Create Host** → `api.myapp.com`
3. **Create Virtual Host** → `api.myapp.com`, round_robin
4. **Create Upstream** → `localhost:3000`
5. **Create Virtual Directory** → `/`, auth: none
6. **Set Methods** → GET, POST, PUT, DELETE

**Selesai!** Gateway sudah aktif di port 8000.

---

## Diagram Alur Request (Runtime)

Ketika client mengirim request ke proxy, ini yang terjadi:

```
Client Request
      │
      ▼
1. Route Matching (host + path)
      │
      ▼
2. Maintenance Check ──── aktif? ──► 503 Maintenance
      │
      ▼
3. CORS Preflight ──── OPTIONS? ──► Return CORS Headers
      │
      ▼
4. IP Filter ──── blacklisted? ──► 403 Forbidden
             ──── not whitelisted? ──► 403 Forbidden
      │
      ▼
5. Rate Limit ──── exceeded? ──► 429 Too Many Requests
      │
      ▼
6. Authentication ──── invalid? ──► 401 Unauthorized
   (api_key/jwt/basic/external)
      │
      ▼
7. Method Check ──── not allowed? ──► 405 Method Not Allowed
      │
      ▼
8. Circuit Breaker ──── open? ──► 503 Service Unavailable
      │
      ▼
9. Forward to Upstream (Load Balanced)
      │
      ▼
10. Return Response + CORS Headers
```

---

## Tips

- **Mulai dari Use Case 1** — setup proxy tanpa auth dulu, pastikan traffic bisa jalan.
- **Tambah security bertahap** — jangan langsung aktifkan semua fitur.
- **Test per langkah** — setelah buat route, langsung test di browser/curl.
- **Perhatikan urutan** — Host → VHost → Upstream → Directory (wajib berurutan).
- **Health check** — `GET http://localhost:9090/api/health` untuk cek admin API hidup.
- **Proxy port** — Traffic masuk melalui port `8000` (default proxy port).
- **Admin port** — Konfigurasi dilakukan di port `9090` (default admin port).
