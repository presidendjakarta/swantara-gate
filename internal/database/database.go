package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"regexp"

	_ "modernc.org/sqlite" // Driver SQLite pure Go (tidak perlu CGO)
)

// DB menyimpan instance koneksi database
var DB *sql.DB

// Daftar tabel yang wajib ada di database
var requiredTables = []string{
	"users",
	"api_consumers",
	"consumer_credentials",
	"hosts",
	"virtual_hosts",
	"upstream_servers",
	"virtual_directories",
	"virtual_directory_methods",
	"route_consumer_access",
	"external_auth_providers",
	"virtual_directory_external_auth",
	"external_auth",
	"jwt_providers",
	"virtual_directory_jwt_providers",
	"jwt_configs",
	"jwt_tokens",
	"rate_limits",
	"cors_configs",
	"circuit_breakers",
	"request_header_rules",
	"response_header_rules",
	"query_rewrites",
	"acme_accounts",
	"ssl_certificates",
	"certificate_domains",
	"ssl_certificate_bindings",
	"tls_options",
	"ip_whitelists",
	"ip_blacklists",
	"service_discovery",
	"config_versions",
	"api_keys",
	"maintenance_windows",
}

// InitDatabase menginisialisasi koneksi database dan menjalankan migrasi
func InitDatabase(dbPath string, sqlPath string) error {
	// Membuat direktori database jika belum ada
	dir := "./data"
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("gagal membuat direktori database: %w", err)
	}

	// Membuka koneksi ke database SQLite
	var err error
	DB, err = sql.Open("sqlite", dbPath)
	if err != nil {
		return fmt.Errorf("gagal membuka database: %w", err)
	}

	// Mengatur koneksi pool untuk performa
	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(10)

	// Memastikan koneksi dapat digunakan
	if err := DB.Ping(); err != nil {
		return fmt.Errorf("gagal melakukan ping ke database: %w", err)
	}

	log.Println("✓ Koneksi database berhasil dibuka")

	// Cek dan jalankan migrasi jika diperlukan
	if err := checkAndMigrate(sqlPath); err != nil {
		return fmt.Errorf("gagal menjalankan migrasi: %w", err)
	}

	return nil
}

// checkAndMigrate mengecek apakah database sudah lengkap, jika belum jalankan migrasi
func checkAndMigrate(sqlPath string) error {
	// Cek apakah semua tabel wajib sudah ada
	missingTables, err := checkMissingTables()
	if err != nil {
		return fmt.Errorf("gagal mengecek tabel: %w", err)
	}

	// Jika tidak ada tabel yang missing, database sudah lengkap
	if len(missingTables) == 0 {
		log.Println("✓ Database sudah lengkap - tidak perlu migrasi")
		return nil
	}

	// Jika ada tabel yang missing, jalankan migrasi
	log.Printf("⚠ Ditemukan %d tabel yang belum ada, menjalankan migrasi...", len(missingTables))
	log.Printf("  Tabel missing: %v", missingTables)

	if err := runMigration(sqlPath); err != nil {
		return fmt.Errorf("gagal menjalankan migrasi: %w", err)
	}

	log.Println("✓ Migrasi database berhasil dijalankan")
	return nil
}

// checkMissingTables mengecek tabel-tabel apa saja yang belum ada di database
func checkMissingTables() ([]string, error) {
	var missingTables []string

	for _, tableName := range requiredTables {
		exists, err := tableExists(tableName)
		if err != nil {
			return nil, fmt.Errorf("gagal mengecek tabel %s: %w", tableName, err)
		}

		if !exists {
			missingTables = append(missingTables, tableName)
		}
	}

	return missingTables, nil
}

// tableExists mengecek apakah sebuah tabel sudah ada di database
func tableExists(tableName string) (bool, error) {
	query := `
		SELECT COUNT(*) 
		FROM sqlite_master 
		WHERE type='table' AND name=?
	`

	var count int
	err := DB.QueryRow(query, tableName).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("gagal query tabel %s: %w", tableName, err)
	}

	return count > 0, nil
}

// runMigration menjalankan file SQL untuk membuat tabel-tabel yang belum ada
func runMigration(sqlPath string) error {
	// Membaca file SQL
	sqlBytes, err := os.ReadFile(sqlPath)
	if err != nil {
		return fmt.Errorf("gagal membaca file SQL %s: %w", sqlPath, err)
	}

	sqlContent := string(sqlBytes)

	// Parse dan execute hanya untuk tabel yang missing
	for _, tableName := range requiredTables {
		exists, err := tableExists(tableName)
		if err != nil {
			return fmt.Errorf("gagal mengecek tabel %s: %w", tableName, err)
		}

		if !exists {
			log.Printf("  Membuat tabel: %s", tableName)
			
			// Extract CREATE TABLE statement untuk tabel ini
			createStmt, err := extractCreateTableStatement(sqlContent, tableName)
			if err != nil {
				return fmt.Errorf("gagal extract CREATE TABLE untuk %s: %w", tableName, err)
			}

			if createStmt == "" {
				log.Printf("  ⚠ CREATE TABLE untuk %s tidak ditemukan di SQL file, skip", tableName)
				continue
			}

			// Execute CREATE TABLE
			_, err = DB.Exec(createStmt)
			if err != nil {
				return fmt.Errorf("gagal membuat tabel %s: %w", tableName, err)
			}

			log.Printf("  ✓ Tabel %s berhasil dibuat", tableName)
		}
	}

	return nil
}

// extractCreateTableStatement ekstrak CREATE TABLE statement untuk tabel tertentu dari SQL content
func extractCreateTableStatement(sqlContent, tableName string) (string, error) {
	// Pattern: CREATE TABLE table_name ( ... );
	// Case insensitive search
	searchPattern := "(?i)CREATE TABLE\\s+" + tableName + "\\s*\\("
	
	re := regexp.MustCompile(searchPattern)
	match := re.FindStringIndex(sqlContent)
	
	if match == nil {
		return "", nil // Not found
	}

	// Find the start of CREATE TABLE
	startIdx := match[0]
	
	// Find the closing ); for this CREATE TABLE
	// We need to count parentheses to find the correct closing
	parenCount := 0
	endIdx := -1
	inString := false
	escaped := false
	
	for i := startIdx; i < len(sqlContent); i++ {
		char := sqlContent[i]
		
		if escaped {
			escaped = false
			continue
		}
		
		if char == '\\' {
			escaped = true
			continue
		}
		
		if char == '\'' {
			inString = !inString
			continue
		}
		
		if inString {
			continue
		}
		
		if char == '(' {
			parenCount++
		} else if char == ')' {
			parenCount--
			if parenCount == 0 {
				// Found closing ), now find the semicolon
				for j := i + 1; j < len(sqlContent); j++ {
					if sqlContent[j] == ';' {
						endIdx = j + 1
						break
					}
					if sqlContent[j] != ' ' && sqlContent[j] != '\t' && sqlContent[j] != '\n' && sqlContent[j] != '\r' {
						// Non-whitespace character before semicolon
						endIdx = i + 1
						break
					}
				}
				break
			}
		}
	}
	
	if endIdx == -1 {
		return "", fmt.Errorf("tidak menemukan penutup CREATE TABLE untuk %s", tableName)
	}
	
	return sqlContent[startIdx:endIdx], nil
}

// CloseDatabase menutup koneksi database dengan aman
func CloseDatabase() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}

// GetDB mengembalikan instance database untuk digunakan di repository
func GetDB() *sql.DB {
	return DB
}

// ListTables mengembalikan daftar semua tabel yang ada di database
func ListTables() ([]string, error) {
	query := `
		SELECT name 
		FROM sqlite_master 
		WHERE type='table' 
		ORDER BY name
	`

	rows, err := DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("gagal query daftar tabel: %w", err)
	}
	defer rows.Close()

	var tables []string
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			return nil, fmt.Errorf("gagal memindai tabel: %w", err)
		}
		tables = append(tables, tableName)
	}

	return tables, nil
}
