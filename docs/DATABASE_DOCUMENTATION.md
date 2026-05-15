# Dokumentasi Database - Swantara Gate API Gateway

Dokumentasi lengkap untuk struktur database API Gateway Proxy menggunakan SQLite.

---

## Daftar Isi

1. [Tabel Users](#1-tabel-users)
2. [Tabel API Consumers](#2-tabel-api_consumers)
3. [Tabel Consumer Credentials](#3-tabel-consumer_credentials)
4. [Tabel Hosts](#4-tabel-hosts)
5. [Tabel Virtual Hosts](#5-tabel-virtual_hosts)
6. [Tabel Upstream Servers](#6-tabel-upstream_servers)
7. [Tabel Virtual Directories](#7-tabel-virtual_directories)
8. [Tabel Virtual Directory Methods](#8-tabel-virtual_directory_methods)
9. [Tabel Route Consumer Access](#9-tabel-route_consumer_access)
10. [Tabel External Auth](#10-tabel-external_auth)
11. [Tabel JWT Configs](#11-tabel-jwt_configs)
12. [Tabel JWT Tokens](#12-tabel-jwt_tokens)
13. [Tabel Rate Limits](#13-tabel-rate_limits)
14. [Tabel CORS Configs](#14-tabel-cors_configs)
15. [Tabel Circuit Breakers](#15-tabel-circuit_breakers)
16. [Tabel Request Header Rules](#16-tabel-request_header_rules)
17. [Tabel Response Header Rules](#17-tabel-response_header_rules)
18. [Tabel Query Rewrites](#18-tabel-query_rewrites)
19. [Tabel ACME Accounts](#19-tabel-acme_accounts)
20. [Tabel SSL Certificates](#20-tabel-ssl_certificates)
21. [Tabel Certificate Domains](#21-tabel-certificate_domains)
22. [Tabel SSL Certificate Bindings](#22-tabel-ssl_certificate_bindings)
23. [Tabel TLS Options](#23-tabel-tls_options)
24. [Tabel IP Whitelists](#24-tabel-ip_whitelists)
25. [Tabel IP Blacklists](#25-tabel-ip_blacklists)
26. [Tabel Service Discovery](#26-tabel-service_discovery)
27. [Tabel Config Versions](#27-tabel-config_versions)
28. [Tabel API Keys](#28-tabel-api_keys)
29. [Tabel Maintenance Windows](#29-tabel-maintenance_windows)

---

## 1. Tabel Users

**Deskripsi:** Menyimpan data pengguna admin panel untuk mengelola API gateway.

| Kolom | Tipe Data | Constraint | Default | Deskripsi |
|-------|-----------|------------|---------|-----------|
| id | INTEGER | PRIMARY KEY AUTOINCREMENT | - | ID unik pengguna |
| username | TEXT | NOT NULL, UNIQUE | - | Username untuk login |
| password_hash | TEXT | NOT NULL | - | Hash password pengguna |
| full_name | TEXT | - | - | Nama lengkap pengguna |
| email | TEXT | - | - | Alamat email pengguna |
| role | TEXT | - | 'admin' | Role pengguna: `super_admin`, `admin`, `operator`, `viewer` |
| is_active | INTEGER | - | 1 | Status aktif (1 = aktif, 0 = tidak aktif) |
| last_login_at | DATETIME | - | - | Waktu terakhir login |
| created_at | DATETIME | - | CURRENT_TIMESTAMP | Waktu pembuatan record |
| updated_at | DATETIME | - | CURRENT_TIMESTAMP | Waktu terakhir update |

**Index:**
- `idx_users_username` pada kolom `username`

---

## 2. Tabel api_consumers

**Deskripsi:** Menyimpan data konsumen/aplikasi yang terdaftar untuk menggunakan API gateway.

| Kolom | Tipe Data | Constraint | Default | Deskripsi |
|-------|-----------|------------|---------|-----------|
| id | INTEGER | PRIMARY KEY AUTOINCREMENT | - | ID unik konsumen |
| consumer_name | TEXT | NOT NULL, UNIQUE | - | Nama konsumen/aplikasi |
| description | TEXT | - | - | Deskripsi konsumen |
| contact_email | TEXT | - | - | Email kontak konsumen |
| is_active | INTEGER | - | 1 | Status aktif (1 = aktif, 0 = tidak aktif) |
| created_at | DATETIME | - | CURRENT_TIMESTAMP | Waktu pembuatan record |
| updated_at | DATETIME | - | CURRENT_TIMESTAMP | Waktu terakhir update |

---

## 3. Tabel consumer_credentials

**Deskripsi:** Menyimpan kredensial autentikasi untuk setiap konsumen API.

| Kolom | Tipe Data | Constraint | Default | Deskripsi |
|-------|-----------|------------|---------|-----------|
| id | INTEGER | PRIMARY KEY AUTOINCREMENT | - | ID unik kredensial |
| consumer_id | INTEGER | NOT NULL, FOREIGN KEY | - | ID konsumen (ref: api_consumers) |
| auth_type | TEXT | NOT NULL | - | Tipe autentikasi: `basic`, `api_key`, `jwt` |
| username | TEXT | - | - | Username untuk basic auth |
| password_hash | TEXT | - | - | Hash password untuk basic auth |
| api_key | TEXT | - | - | API key untuk api_key auth |
| jwt_secret | TEXT | - | - | Secret key untuk JWT auth |
| expired_at | DATETIME | - | - | Waktu kedaluwarsa kredensial |
| is_active | INTEGER | - | 1 | Status aktif (1 = aktif, 0 = tidak aktif) |
| created_at | DATETIME | - | CURRENT_TIMESTAMP | Waktu pembuatan record |

**Foreign Key:**
- `consumer_id` → `api_consumers(id)` ON DELETE CASCADE

---

## 4. Tabel hosts

**Deskripsi:** Menyimpan data host utama sebagai grouping root untuk virtual hosts.

| Kolom | Tipe Data | Constraint | Default | Deskripsi |
|-------|-----------|------------|---------|-----------|
| id | INTEGER | PRIMARY KEY AUTOINCREMENT | - | ID unik host |
| host_name | TEXT | NOT NULL, UNIQUE | - | Nama host |
| description | TEXT | - | - | Deskripsi host |
| is_active | INTEGER | - | 1 | Status aktif (1 = aktif, 0 = tidak aktif) |
| created_at | DATETIME | - | CURRENT_TIMESTAMP | Waktu pembuatan record |
| updated_at | DATETIME | - | CURRENT_TIMESTAMP | Waktu terakhir update |

---

## 5. Tabel virtual_hosts

**Deskripsi:** Menyimpan konfigurasi virtual domain dengan pengaturan load balancing.

| Kolom | Tipe Data | Constraint | Default | Deskripsi |
|-------|-----------|------------|---------|-----------|
| id | INTEGER | PRIMARY KEY AUTOINCREMENT | - | ID unik virtual host |
| host_id | INTEGER | NOT NULL, FOREIGN KEY | - | ID host utama (ref: hosts) |
| vhost_name | TEXT | NOT NULL, UNIQUE | - | Nama virtual host/domain |
| lb_algorithm | TEXT | - | 'round_robin' | Algoritma load balancing: `round_robin`, `weighted_round_robin`, `least_conn`, `ip_hash`, `random`, `failover` |
| sticky_session | INTEGER | - | 0 | Sticky session enabled (1 = ya, 0 = tidak) |
| failover_mode | TEXT | - | 'active-active' | Mode failover: `active-active`, `active-passive` |
| is_active | INTEGER | - | 1 | Status aktif (1 = aktif, 0 = tidak aktif) |
| created_at | DATETIME | - | CURRENT_TIMESTAMP | Waktu pembuatan record |
| updated_at | DATETIME | - | CURRENT_TIMESTAMP | Waktu terakhir update |

**Foreign Key:**
- `host_id` → `hosts(id)` ON DELETE CASCADE

---

## 6. Tabel upstream_servers

**Deskripsi:** Menyimpan data server backend dengan konfigurasi health check.

| Kolom | Tipe Data | Constraint | Default | Deskripsi |
|-------|-----------|------------|---------|-----------|
| id | INTEGER | PRIMARY KEY AUTOINCREMENT | - | ID unik upstream server |
| virtual_host_id | INTEGER | NOT NULL, FOREIGN KEY | - | ID virtual host (ref: virtual_hosts) |
| target_host | TEXT | NOT NULL | - | Host/IP server backend |
| target_port | INTEGER | NOT NULL | - | Port server backend |
| protocol | TEXT | - | 'http' | Protokol: `http`, `https` |
| priority | INTEGER | - | 1 | Prioritas server |
| weight | INTEGER | - | 1 | Bobot untuk weighted load balancing |
| is_backup | INTEGER | - | 0 | Server backup (1 = ya, 0 = tidak) |
| is_active | INTEGER | - | 1 | Status aktif (1 = aktif, 0 = tidak aktif) |
| health_check_enabled | INTEGER | - | 1 | Health check enabled (1 = ya, 0 = tidak) |
| health_check_path | TEXT | - | '/health' | Path untuk health check |
| health_check_interval_seconds | INTEGER | - | 10 | Interval health check (detik) |
| health_check_timeout_seconds | INTEGER | - | 3 | Timeout health check (detik) |
| max_fails | INTEGER | - | 3 | Maksimal kegagalan sebelum dianggap down |
| fail_timeout_seconds | INTEGER | - | 30 | Timeout sebelum server dianggap failed |
| created_at | DATETIME | - | CURRENT_TIMESTAMP | Waktu pembuatan record |
| updated_at | DATETIME | - | CURRENT_TIMESTAMP | Waktu terakhir update |

**Foreign Key:**
- `virtual_host_id` → `virtual_hosts(id)` ON DELETE CASCADE

---

## 7. Tabel virtual_directories

**Deskripsi:** Menyimpan konfigurasi route/API endpoint dengan berbagai pengaturan proxy.

| Kolom | Tipe Data | Constraint | Default | Deskripsi |
|-------|-----------|------------|---------|-----------|
| id | INTEGER | PRIMARY KEY AUTOINCREMENT | - | ID unik virtual directory |
| virtual_host_id | INTEGER | NOT NULL, FOREIGN KEY | - | ID virtual host (ref: virtual_hosts) |
| source_path | TEXT | NOT NULL | - | Path sumber (route yang diakses client) |
| target_path | TEXT | NOT NULL | - | Path target (path di server backend) |
| match_type | TEXT | - | 'prefix' | Tipe matching: `exact`, `prefix`, `wildcard`, `regex`, `parameter` |
| strip_prefix | INTEGER | - | 1 | Strip prefix dari path (1 = ya, 0 = tidak) |
| preserve_host_header | INTEGER | - | 0 | Preserve host header (1 = ya, 0 = tidak) |
| auth_type | TEXT | - | 'none' | Tipe autentikasi: `none`, `basic`, `jwt`, `external`, `api_key` |
| is_active | INTEGER | - | 1 | Status aktif (1 = aktif, 0 = tidak aktif) |
| proxy_timeout_seconds | INTEGER | - | 30 | Timeout proxy (detik) |
| retry_count | INTEGER | - | 0 | Jumlah retry jika gagal |
| retry_delay_ms | INTEGER | - | 100 | Delay antar retry (milidetik) |
| max_request_size_mb | INTEGER | - | 10 | Maksimal ukuran request (MB) |
| websocket_enabled | INTEGER | - | 0 | WebSocket enabled (1 = ya, 0 = tidak) |
| cache_enabled | INTEGER | - | 0 | Cache enabled (1 = ya, 0 = tidak) |
| cache_ttl_seconds | INTEGER | - | 60 | Cache TTL (detik) |
| created_at | DATETIME | - | CURRENT_TIMESTAMP | Waktu pembuatan record |
| updated_at | DATETIME | - | CURRENT_TIMESTAMP | Waktu terakhir update |

**Foreign Key:**
- `virtual_host_id` → `virtual_hosts(id)` ON DELETE CASCADE

---

## 8. Tabel virtual_directory_methods

**Deskripsi:** Menyimpan metode HTTP yang diizinkan untuk setiap virtual directory.

| Kolom | Tipe Data | Constraint | Default | Deskripsi |
|-------|-----------|------------|---------|-----------|
| id | INTEGER | PRIMARY KEY AUTOINCREMENT | - | ID unik record |
| virtual_directory_id | INTEGER | NOT NULL, FOREIGN KEY | - | ID virtual directory (ref: virtual_directories) |
| http_method | TEXT | NOT NULL | - | Metode HTTP: `GET`, `POST`, `PUT`, `DELETE`, `PATCH`, `OPTIONS`, dll |

**Foreign Key:**
- `virtual_directory_id` → `virtual_directories(id)` ON DELETE CASCADE

---

## 9. Tabel route_consumer_access

**Deskripsi:** Mengontrol akses konsumen ke route tertentu.

| Kolom | Tipe Data | Constraint | Default | Deskripsi |
|-------|-----------|------------|---------|-----------|
| id | INTEGER | PRIMARY KEY AUTOINCREMENT | - | ID unik record |
| virtual_directory_id | INTEGER | NOT NULL, FOREIGN KEY | - | ID virtual directory (ref: virtual_directories) |
| consumer_id | INTEGER | NOT NULL, FOREIGN KEY | - | ID konsumen (ref: api_consumers) |
| is_active | INTEGER | - | 1 | Status akses (1 = diizinkan, 0 = tidak) |
| created_at | DATETIME | - | CURRENT_TIMESTAMP | Waktu pembuatan record |

**Foreign Keys:**
- `virtual_directory_id` → `virtual_directories(id)` ON DELETE CASCADE
- `consumer_id` → `api_consumers(id)` ON DELETE CASCADE

---

## 10. Tabel external_auth

**Deskripsi:** Konfigurasi autentikasi eksternal untuk validasi request.

| Kolom | Tipe Data | Constraint | Default | Deskripsi |
|-------|-----------|------------|---------|-----------|
| id | INTEGER | PRIMARY KEY AUTOINCREMENT | - | ID unik record |
| virtual_directory_id | INTEGER | NOT NULL, FOREIGN KEY | - | ID virtual directory (ref: virtual_directories) |
| auth_url | TEXT | NOT NULL | - | URL endpoint autentikasi eksternal |
| http_method | TEXT | - | 'POST' | Metode HTTP untuk auth request |
| request_timeout_seconds | INTEGER | - | 5 | Timeout request auth (detik) |
| send_headers | INTEGER | - | 1 | Kirim headers ke auth endpoint (1 = ya, 0 = tidak) |
| send_body | INTEGER | - | 0 | Kirim body ke auth endpoint (1 = ya, 0 = tidak) |
| success_key | TEXT | - | 'status' | Key JSON untuk cek sukses |
| success_value | TEXT | - | 'true' | Value yang menandakan sukses |
| message_key | TEXT | - | 'message' | Key JSON untuk pesan |
| token_key | TEXT | - | - | Key JSON untuk token (jika ada) |
| is_active | INTEGER | - | 1 | Status aktif (1 = aktif, 0 = tidak aktif) |
| created_at | DATETIME | - | CURRENT_TIMESTAMP | Waktu pembuatan record |

**Foreign Key:**
- `virtual_directory_id` → `virtual_directories(id)` ON DELETE CASCADE

---

## 11. Tabel jwt_configs

**Deskripsi:** Konfigurasi JWT untuk validasi token.

| Kolom | Tipe Data | Constraint | Default | Deskripsi |
|-------|-----------|------------|---------|-----------|
| id | INTEGER | PRIMARY KEY AUTOINCREMENT | - | ID unik record |
| virtual_directory_id | INTEGER | NOT NULL, FOREIGN KEY | - | ID virtual directory (ref: virtual_directories) |
| algorithm | TEXT | - | 'HS256' | Algoritma JWT: `HS256`, `HS384`, `HS512`, `RS256`, dll |
| jwt_secret | TEXT | NOT NULL | - | Secret key untuk validasi JWT |
| issuer | TEXT | - | - | Issuer yang diharapkan (claim `iss`) |
| audience | TEXT | - | - | Audience yang diharapkan (claim `aud`) |
| expired_in_seconds | INTEGER | - | 3600 | Masa berlaku token (detik) |
| clock_skew_seconds | INTEGER | - | 30 | Toleransi perbedaan waktu (detik) |
| require_exp | INTEGER | - | 1 | Wajibkan claim `exp` (1 = ya, 0 = tidak) |
| require_iat | INTEGER | - | 1 | Wajibkan claim `iat` (1 = ya, 0 = tidak) |
| is_active | INTEGER | - | 1 | Status aktif (1 = aktif, 0 = tidak aktif) |
| created_at | DATETIME | - | CURRENT_TIMESTAMP | Waktu pembuatan record |

**Foreign Key:**
- `virtual_directory_id` → `virtual_directories(id)` ON DELETE CASCADE

---

## 12. Tabel jwt_tokens

**Deskripsi:** Manajemen token JWT untuk blacklist/session tracking.

| Kolom | Tipe Data | Constraint | Default | Deskripsi |
|-------|-----------|------------|---------|-----------|
| id | INTEGER | PRIMARY KEY AUTOINCREMENT | - | ID unik record |
| consumer_id | INTEGER | FOREIGN KEY | - | ID konsumen (ref: api_consumers) |
| token | TEXT | NOT NULL | - | Token JWT |
| issued_at | DATETIME | - | CURRENT_TIMESTAMP | Waktu token diterbitkan |
| expired_at | DATETIME | - | - | Waktu token kedaluwarsa |
| is_revoked | INTEGER | - | 0 | Status revoke (1 = revoked, 0 = aktif) |
| ip_address | TEXT | - | - | IP address saat token diterbitkan |
| user_agent | TEXT | - | - | User agent saat token diterbitkan |

**Foreign Key:**
- `consumer_id` → `api_consumers(id)` ON DELETE CASCADE

---

## 13. Tabel rate_limits

**Deskripsi:** Konfigurasi rate limiting untuk mencegah abuse.

| Kolom | Tipe Data | Constraint | Default | Deskripsi |
|-------|-----------|------------|---------|-----------|
| id | INTEGER | PRIMARY KEY AUTOINCREMENT | - | ID unik record |
| virtual_directory_id | INTEGER | NOT NULL, FOREIGN KEY | - | ID virtual directory (ref: virtual_directories) |
| limit_by | TEXT | - | 'ip' | Limit berdasarkan: `ip`, `consumer`, `api_key` |
| requests_per_minute | INTEGER | - | 60 | Maksimal request per menit |
| burst | INTEGER | - | 10 | Maksimal burst request |
| block_duration_seconds | INTEGER | - | 60 | Durasi block jika melebihi limit (detik) |
| is_active | INTEGER | - | 1 | Status aktif (1 = aktif, 0 = tidak aktif) |
| created_at | DATETIME | - | CURRENT_TIMESTAMP | Waktu pembuatan record |

**Foreign Key:**
- `virtual_directory_id` → `virtual_directories(id)` ON DELETE CASCADE

---

## 14. Tabel cors_configs

**Deskripsi:** Konfigurasi CORS (Cross-Origin Resource Sharing).

| Kolom | Tipe Data | Constraint | Default | Deskripsi |
|-------|-----------|------------|---------|-----------|
| id | INTEGER | PRIMARY KEY AUTOINCREMENT | - | ID unik record |
| virtual_directory_id | INTEGER | NOT NULL, FOREIGN KEY | - | ID virtual directory (ref: virtual_directories) |
| allowed_origins | TEXT | - | '*' | Origin yang diizinkan (comma-separated) |
| allowed_methods | TEXT | - | 'GET,POST,PUT,DELETE,OPTIONS' | Metode HTTP yang diizinkan |
| allowed_headers | TEXT | - | '*' | Header yang diizinkan |
| exposed_headers | TEXT | - | - | Header yang di-expose ke client |
| allow_credentials | INTEGER | - | 0 | Allow credentials (1 = ya, 0 = tidak) |
| max_age_seconds | INTEGER | - | 3600 | Max age untuk preflight cache (detik) |
| is_active | INTEGER | - | 1 | Status aktif (1 = aktif, 0 = tidak aktif) |
| created_at | DATETIME | - | CURRENT_TIMESTAMP | Waktu pembuatan record |

**Foreign Key:**
- `virtual_directory_id` → `virtual_directories(id)` ON DELETE CASCADE

---

## 15. Tabel circuit_breakers

**Deskripsi:** Circuit breaker protection untuk mencegah cascade failure.

| Kolom | Tipe Data | Constraint | Default | Deskripsi |
|-------|-----------|------------|---------|-----------|
| id | INTEGER | PRIMARY KEY AUTOINCREMENT | - | ID unik record |
| virtual_directory_id | INTEGER | NOT NULL, FOREIGN KEY | - | ID virtual directory (ref: virtual_directories) |
| enabled | INTEGER | - | 1 | Circuit breaker enabled (1 = ya, 0 = tidak) |
| failure_threshold | INTEGER | - | 5 | Jumlah failure sebelum circuit terbuka |
| recovery_timeout_seconds | INTEGER | - | 30 | Timeout sebelum mencoba recovery (detik) |
| half_open_max_requests | INTEGER | - | 3 | Maksimal request di state half-open |
| created_at | DATETIME | - | CURRENT_TIMESTAMP | Waktu pembuatan record |

**Foreign Key:**
- `virtual_directory_id` → `virtual_directories(id)` ON DELETE CASCADE

---

## 16. Tabel request_header_rules

**Deskripsi:** Aturan manipulasi header request sebelum diteruskan ke backend.

| Kolom | Tipe Data | Constraint | Default | Deskripsi |
|-------|-----------|------------|---------|-----------|
| id | INTEGER | PRIMARY KEY AUTOINCREMENT | - | ID unik record |
| virtual_directory_id | INTEGER | NOT NULL, FOREIGN KEY | - | ID virtual directory (ref: virtual_directories) |
| header_name | TEXT | NOT NULL | - | Nama header |
| operation | TEXT | NOT NULL | - | Operasi: `add`, `set`, `remove`, `rename` |
| value_source | TEXT | - | 'static' | Sumber value: `static`, `header`, `variable` |
| header_value | TEXT | - | - | Nilai header (jika static) |
| source_header | TEXT | - | - | Header sumber (jika value_source = header) |
| variable_name | TEXT | - | - | Nama variabel (jika value_source = variable) |
| execution_order | INTEGER | - | 1 | Urutan eksekusi |
| is_active | INTEGER | - | 1 | Status aktif (1 = aktif, 0 = tidak aktif) |
| created_at | DATETIME | - | CURRENT_TIMESTAMP | Waktu pembuatan record |

**Foreign Key:**
- `virtual_directory_id` → `virtual_directories(id)` ON DELETE CASCADE

---

## 17. Tabel response_header_rules

**Deskripsi:** Aturan manipulasi header response sebelum dikirim ke client.

| Kolom | Tipe Data | Constraint | Default | Deskripsi |
|-------|-----------|------------|---------|-----------|
| id | INTEGER | PRIMARY KEY AUTOINCREMENT | - | ID unik record |
| virtual_directory_id | INTEGER | NOT NULL, FOREIGN KEY | - | ID virtual directory (ref: virtual_directories) |
| header_name | TEXT | NOT NULL | - | Nama header |
| operation | TEXT | NOT NULL | - | Operasi: `add`, `set`, `remove`, `rename` |
| header_value | TEXT | - | - | Nilai header |
| execution_order | INTEGER | - | 1 | Urutan eksekusi |
| is_active | INTEGER | - | 1 | Status aktif (1 = aktif, 0 = tidak aktif) |
| created_at | DATETIME | - | CURRENT_TIMESTAMP | Waktu pembuatan record |

**Foreign Key:**
- `virtual_directory_id` → `virtual_directories(id)` ON DELETE CASCADE

---

## 18. Tabel query_rewrites

**Deskripsi:** Aturan rewrite parameter query string.

| Kolom | Tipe Data | Constraint | Default | Deskripsi |
|-------|-----------|------------|---------|-----------|
| id | INTEGER | PRIMARY KEY AUTOINCREMENT | - | ID unik record |
| virtual_directory_id | INTEGER | NOT NULL, FOREIGN KEY | - | ID virtual directory (ref: virtual_directories) |
| param_name | TEXT | NOT NULL | - | Nama parameter |
| param_value | TEXT | - | - | Nilai parameter |
| operation | TEXT | - | 'set' | Operasi: `set`, `add`, `remove`, `rename` |
| created_at | DATETIME | - | CURRENT_TIMESTAMP | Waktu pembuatan record |

**Foreign Key:**
- `virtual_directory_id` → `virtual_directories(id)` ON DELETE CASCADE

---

## 19. Tabel acme_accounts

**Deskripsi:** Akun ACME untuk Let's Encrypt SSL certificate.

| Kolom | Tipe Data | Constraint | Default | Deskripsi |
|-------|-----------|------------|---------|-----------|
| id | INTEGER | PRIMARY KEY AUTOINCREMENT | - | ID unik record |
| email | TEXT | NOT NULL | - | Email untuk registrasi ACME |
| provider_url | TEXT | NOT NULL | - | URL provider ACME |
| account_key_path | TEXT | NOT NULL | - | Path ke file account key |
| is_default | INTEGER | - | 0 | Akun default (1 = ya, 0 = tidak) |
| created_at | DATETIME | - | CURRENT_TIMESTAMP | Waktu pembuatan record |

---

## 20. Tabel ssl_certificates

**Deskripsi:** Penyimpanan sertifikat SSL/TLS.

| Kolom | Tipe Data | Constraint | Default | Deskripsi |
|-------|-----------|------------|---------|-----------|
| id | INTEGER | PRIMARY KEY AUTOINCREMENT | - | ID unik record |
| acme_account_id | INTEGER | FOREIGN KEY | - | ID akun ACME (ref: acme_accounts) |
| provider | TEXT | - | 'lets_encrypt' | Provider sertifikat |
| challenge_type | TEXT | - | 'http01' | Tipe challenge: `http01`, `dns01` |
| certificate_path | TEXT | NOT NULL | - | Path ke file sertifikat |
| private_key_path | TEXT | NOT NULL | - | Path ke file private key |
| chain_path | TEXT | - | - | Path ke file certificate chain |
| auto_renew | INTEGER | - | 1 | Auto renew enabled (1 = ya, 0 = tidak) |
| renew_before_days | INTEGER | - | 30 | Renew sebelum expired (hari) |
| last_renew_at | DATETIME | - | - | Waktu renew terakhir |
| expired_at | DATETIME | - | - | Waktu kedaluwarsa sertifikat |
| renew_status | TEXT | - | 'pending' | Status renew: `pending`, `success`, `failed` |
| last_error | TEXT | - | - | Error terakhir (jika ada) |
| is_active | INTEGER | - | 1 | Status aktif (1 = aktif, 0 = tidak aktif) |
| created_at | DATETIME | - | CURRENT_TIMESTAMP | Waktu pembuatan record |
| updated_at | DATETIME | - | CURRENT_TIMESTAMP | Waktu terakhir update |

**Foreign Key:**
- `acme_account_id` → `acme_accounts(id)` ON DELETE SET NULL

---

## 21. Tabel certificate_domains

**Deskripsi:** Domain yang terasosiasi dengan sertifikat SSL.

| Kolom | Tipe Data | Constraint | Default | Deskripsi |
|-------|-----------|------------|---------|-----------|
| id | INTEGER | PRIMARY KEY AUTOINCREMENT | - | ID unik record |
| ssl_certificate_id | INTEGER | NOT NULL, FOREIGN KEY | - | ID sertifikat SSL (ref: ssl_certificates) |
| domain_name | TEXT | NOT NULL | - | Nama domain |
| is_wildcard | INTEGER | - | 0 | Wildcard domain (1 = ya, 0 = tidak) |
| created_at | DATETIME | - | CURRENT_TIMESTAMP | Waktu pembuatan record |

**Foreign Key:**
- `ssl_certificate_id` → `ssl_certificates(id)` ON DELETE CASCADE

---

## 22. Tabel ssl_certificate_bindings

**Deskripsi:** Binding sertifikat SSL ke host atau virtual host.

| Kolom | Tipe Data | Constraint | Default | Deskripsi |
|-------|-----------|------------|---------|-----------|
| id | INTEGER | PRIMARY KEY AUTOINCREMENT | - | ID unik record |
| ssl_certificate_id | INTEGER | NOT NULL, FOREIGN KEY | - | ID sertifikat SSL (ref: ssl_certificates) |
| binding_type | TEXT | NOT NULL | - | Tipe binding: `host`, `virtual_host`, `global` |
| host_id | INTEGER | FOREIGN KEY | - | ID host (ref: hosts) |
| virtual_host_id | INTEGER | FOREIGN KEY | - | ID virtual host (ref: virtual_hosts) |
| is_default | INTEGER | - | 0 | Binding default (1 = ya, 0 = tidak) |
| priority | INTEGER | - | 1 | Prioritas binding |
| created_at | DATETIME | - | CURRENT_TIMESTAMP | Waktu pembuatan record |

**Foreign Keys:**
- `ssl_certificate_id` → `ssl_certificates(id)` ON DELETE CASCADE
- `host_id` → `hosts(id)` ON DELETE CASCADE
- `virtual_host_id` → `virtual_hosts(id)` ON DELETE CASCADE

---

## 23. Tabel tls_options

**Deskripsi:** Konfigurasi opsi TLS/SSL.

| Kolom | Tipe Data | Constraint | Default | Deskripsi |
|-------|-----------|------------|---------|-----------|
| id | INTEGER | PRIMARY KEY AUTOINCREMENT | - | ID unik record |
| binding_type | TEXT | NOT NULL | - | Tipe binding: `host`, `virtual_host`, `global` |
| host_id | INTEGER | FOREIGN KEY | - | ID host (ref: hosts) |
| virtual_host_id | INTEGER | FOREIGN KEY | - | ID virtual host (ref: virtual_hosts) |
| min_tls_version | TEXT | - | '1.2' | Versi TLS minimum: `1.0`, `1.1`, `1.2`, `1.3` |
| http2_enabled | INTEGER | - | 1 | HTTP/2 enabled (1 = ya, 0 = tidak) |
| hsts_enabled | INTEGER | - | 1 | HSTS enabled (1 = ya, 0 = tidak) |
| hsts_max_age | INTEGER | - | 31536000 | HSTS max age (detik) |
| created_at | DATETIME | - | CURRENT_TIMESTAMP | Waktu pembuatan record |

**Foreign Keys:**
- `host_id` → `hosts(id)` ON DELETE CASCADE
- `virtual_host_id` → `virtual_hosts(id)` ON DELETE CASCADE

---

## 24. Tabel ip_whitelists

**Deskripsi:** Daftar IP yang diizinkan mengakses route tertentu.

| Kolom | Tipe Data | Constraint | Default | Deskripsi |
|-------|-----------|------------|---------|-----------|
| id | INTEGER | PRIMARY KEY AUTOINCREMENT | - | ID unik record |
| virtual_directory_id | INTEGER | NOT NULL, FOREIGN KEY | - | ID virtual directory (ref: virtual_directories) |
| ip_address | TEXT | NOT NULL | - | Alamat IP yang diizinkan |
| description | TEXT | - | - | Deskripsi |
| is_active | INTEGER | - | 1 | Status aktif (1 = aktif, 0 = tidak aktif) |
| created_at | DATETIME | - | CURRENT_TIMESTAMP | Waktu pembuatan record |

**Foreign Key:**
- `virtual_directory_id` → `virtual_directories(id)` ON DELETE CASCADE

---

## 25. Tabel ip_blacklists

**Deskripsi:** Daftar IP yang diblokir dari route tertentu.

| Kolom | Tipe Data | Constraint | Default | Deskripsi |
|-------|-----------|------------|---------|-----------|
| id | INTEGER | PRIMARY KEY AUTOINCREMENT | - | ID unik record |
| virtual_directory_id | INTEGER | NOT NULL, FOREIGN KEY | - | ID virtual directory (ref: virtual_directories) |
| ip_address | TEXT | NOT NULL | - | Alamat IP yang diblokir |
| reason | TEXT | - | - | Alasan pemblokiran |
| expired_at | DATETIME | - | - | Waktu kedaluwarsa blacklist |
| is_active | INTEGER | - | 1 | Status aktif (1 = aktif, 0 = tidak aktif) |
| created_at | DATETIME | - | CURRENT_TIMESTAMP | Waktu pembuatan record |

**Foreign Key:**
- `virtual_directory_id` → `virtual_directories(id)` ON DELETE CASCADE

---

## 26. Tabel service_discovery

**Deskripsi:** Konfigurasi service discovery untuk backend otomatis.

| Kolom | Tipe Data | Constraint | Default | Deskripsi |
|-------|-----------|------------|---------|-----------|
| id | INTEGER | PRIMARY KEY AUTOINCREMENT | - | ID unik record |
| virtual_host_id | INTEGER | NOT NULL, FOREIGN KEY | - | ID virtual host (ref: virtual_hosts) |
| provider | TEXT | NOT NULL | - | Provider service discovery |
| endpoint_url | TEXT | NOT NULL | - | URL endpoint service discovery |
| refresh_interval_seconds | INTEGER | - | 30 | Interval refresh (detik) |
| is_active | INTEGER | - | 1 | Status aktif (1 = aktif, 0 = tidak aktif) |
| created_at | DATETIME | - | CURRENT_TIMESTAMP | Waktu pembuatan record |

**Foreign Key:**
- `virtual_host_id` → `virtual_hosts(id)` ON DELETE CASCADE

---

## 27. Tabel config_versions

**Deskripsi:** Versioning konfigurasi untuk tracking perubahan.

| Kolom | Tipe Data | Constraint | Default | Deskripsi |
|-------|-----------|------------|---------|-----------|
| id | INTEGER | PRIMARY KEY AUTOINCREMENT | - | ID unik record |
| config_name | TEXT | NOT NULL | - | Nama konfigurasi |
| version_number | INTEGER | NOT NULL | - | Nomor versi |
| changed_by | TEXT | - | - | Diubah oleh |
| notes | TEXT | - | - | Catatan perubahan |
| created_at | DATETIME | - | CURRENT_TIMESTAMP | Waktu pembuatan record |

---

## 28. Tabel api_keys

**Deskripsi:** API keys dedicated untuk konsumen dengan override settings.

| Kolom | Tipe Data | Constraint | Default | Deskripsi |
|-------|-----------|------------|---------|-----------|
| id | INTEGER | PRIMARY KEY AUTOINCREMENT | - | ID unik record |
| consumer_id | INTEGER | NOT NULL, FOREIGN KEY | - | ID konsumen (ref: api_consumers) |
| api_key | TEXT | NOT NULL, UNIQUE | - | API key unik |
| description | TEXT | - | - | Deskripsi API key |
| expired_at | DATETIME | - | - | Waktu kedaluwarsa |
| rate_limit_override | INTEGER | - | - | Override rate limit (jika ada) |
| is_active | INTEGER | - | 1 | Status aktif (1 = aktif, 0 = tidak aktif) |
| created_at | DATETIME | - | CURRENT_TIMESTAMP | Waktu pembuatan record |

**Foreign Key:**
- `consumer_id` → `api_consumers(id)` ON DELETE CASCADE

---

## 29. Tabel maintenance_windows

**Deskripsi:** Konfigurasi mode maintenance untuk virtual host.

| Kolom | Tipe Data | Constraint | Default | Deskripsi |
|-------|-----------|------------|---------|-----------|
| id | INTEGER | PRIMARY KEY AUTOINCREMENT | - | ID unik record |
| virtual_host_id | INTEGER | FOREIGN KEY | - | ID virtual host (ref: virtual_hosts) |
| title | TEXT | - | - | Judul maintenance |
| start_at | DATETIME | - | - | Waktu mulai maintenance |
| end_at | DATETIME | - | - | Waktu selesai maintenance |
| maintenance_response_code | INTEGER | - | 503 | Response code saat maintenance |
| maintenance_message | TEXT | - | - | Pesan maintenance |
| is_active | INTEGER | - | 1 | Status aktif (1 = aktif, 0 = tidak aktif) |
| created_at | DATETIME | - | CURRENT_TIMESTAMP | Waktu pembuatan record |

**Foreign Key:**
- `virtual_host_id` → `virtual_hosts(id)` ON DELETE CASCADE

---

## Relasi Antar Tabel

```
users
  └── (independen)

hosts
  └── virtual_hosts
       ├── upstream_servers
       ├── virtual_directories
       │    ├── virtual_directory_methods
       │    ├── route_consumer_access ← api_consumers
       │    ├── external_auth
       │    ├── jwt_configs
       │    ├── rate_limits
       │    ├── cors_configs
       │    ├── circuit_breakers
       │    ├── request_header_rules
       │    ├── response_header_rules
       │    ├── query_rewrites
       │    ├── ip_whitelists
       │    └── ip_blacklists
       ├── service_discovery
       └── maintenance_windows

api_consumers
  ├── consumer_credentials
  ├── api_keys
  ├── route_consumer_access → virtual_directories
  └── jwt_tokens

acme_accounts
  └── ssl_certificates
       ├── certificate_domains
       └── ssl_certificate_bindings → hosts, virtual_hosts

tls_options → hosts, virtual_hosts

config_versions
  └── (independen)
```

---

## Catatan Penting

### Tipe Data
- **INTEGER**: SQLite integer (0 atau 1 untuk boolean)
- **TEXT**: String/text
- **DATETIME**: Timestamp dalam format `YYYY-MM-DD HH:MM:SS`

### Konvensi
- Semua tabel menggunakan `id` sebagai PRIMARY KEY dengan AUTOINCREMENT
- Foreign keys menggunakan `ON DELETE CASCADE` atau `ON DELETE SET NULL`
- Timestamp `created_at` dan `updated_at` menggunakan `CURRENT_TIMESTAMP`
- Boolean diwakili oleh INTEGER: `1` = true/aktif, `0` = false/tidak aktif

### Index
- `idx_users_username` pada tabel `users(username)` untuk query yang lebih cepat

---

**Dibuat:** 2026-05-15  
**Versi Database:** 1.0  
**Database Engine:** SQLite 3
