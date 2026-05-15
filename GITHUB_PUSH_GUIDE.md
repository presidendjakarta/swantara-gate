# Panduan Push ke GitHub

## Repository
https://github.com/presidendjakarta/swantara-gate

## Langkah-langkah Push

### 1. Inisialisasi Git (jika belum)
```bash
cd x:\laragon\go-apps\swantara-gate
git init
```

### 2. Tambahkan Remote Repository
```bash
git remote add origin https://github.com/presidendjakarta/swantara-gate.git
```

### 3. Tambahkan Semua File
```bash
git add .
```

### 4. Commit Changes
```bash
git commit -m "feat: Initial commit - Admin API CRUD implementation

- User management (CRUD)
- API Consumer management (CRUD)
- Host management (CRUD)
- Virtual Host management (CRUD)
- SQLite pure Go (no CGO)
- Clean architecture
- RESTful API
- Multi-port support (admin & proxy)
- Documentation"
```

### 5. Push ke GitHub
```bash
git branch -M main
git push -u origin main
```

## Jika Ada Error Authentication

### Opsi 1: Menggunakan GitHub CLI
```bash
# Login ke GitHub
gh auth login

# Push ulang
git push -u origin main
```

### Opsi 2: Menggunakan Personal Access Token
1. Buat token di: https://github.com/settings/tokens
2. Push dengan token:
```bash
git push https://TOKEN@github.com/presidendjakarta/swantara-gate.git main
```

### Opsi 3: Menggunakan SSH
```bash
# Generate SSH key (jika belum)
ssh-keygen -t ed25519 -C "your_email@example.com"

# Tambahkan ke GitHub
# Copy public key: cat ~/.ssh/id_ed25519.pub
# Paste di: https://github.com/settings/keys

# Ubah remote ke SSH
git remote set-url origin git@github.com:presidendjakarta/swantara-gate.git

# Push
git push -u origin main
```

## Verifikasi

Setelah push, cek repository di:
https://github.com/presidendjakarta/swantara-gate

## File yang Di-push

✅ Source code Go
✅ Database schema
✅ Documentation
✅ .env.example (template)
✅ README.md
✅ .gitignore

## File yang TIDAK Di-push (sesuai .gitignore)

❌ *.exe (binary files)
❌ .env (berisi sensitive data)
❌ *.db (database files)
❌ go.work.sum
❌ .DS_Store, Thumbs.db, dll

## Struktur Repository

```
swantara-gate/
├── cmd/server/main.go          # Entry point
├── internal/                    # Private application code
│   ├── config/                 # Configuration
│   ├── database/               # Database connection
│   ├── handler/                # HTTP handlers
│   ├── middleware/             # HTTP middleware
│   ├── model/                  # Data models
│   ├── repository/             # Database repositories
│   ├── response/               # Response helpers
│   └── service/                # Business logic
├── data/
│   └── database.sql            # Database schema
├── docs/                        # Documentation
│   ├── API_DOCUMENTATION.md
│   └── DATABASE_DOCUMENTATION.md
├── .env.example                # Environment template
├── .gitignore
├── go.mod
├── go.sum
└── README.md
```

## Testing Setelah Clone

```bash
# Clone repository
git clone https://github.com/presidendjakarta/swantara-gate.git
cd swantara-gate

# Copy env file
cp .env.example .env

# Install dependencies
go mod tidy

# Run application
go run cmd/server/main.go

# Test API
curl http://localhost:8081/api/health
```

## Next Steps

Setelah push ke GitHub:
1. Setup CI/CD pipeline (optional)
2. Setup GitHub Actions untuk auto-build
3. Tambahkan contributors
4. Setup project board
5. Tambahkan issue templates
