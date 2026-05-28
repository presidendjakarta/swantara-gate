package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"

	_ "modernc.org/sqlite"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "create-user":
		handleCreateUser(os.Args[2:])
	case "reset-password":
		handleResetPassword(os.Args[2:])
	default:
		fmt.Printf("❌ Perintah tidak dikenal: %s\n\n", command)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Swantara Gate CLI Tool")
	fmt.Println("======================")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  swantara-cli <command> [options]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  create-user      Membuat user baru")
	fmt.Println("  reset-password   Reset password user")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  swantara-cli create-user -username admin -password secret123 -role super_admin")
	fmt.Println("  swantara-cli reset-password -username admin -password newsecret123")
}

func handleCreateUser(args []string) {
	fs := flag.NewFlagSet("create-user", flag.ExitOnError)

	username := fs.String("username", "", "Username (wajib)")
	password := fs.String("password", "", "Password (wajib)")
	fullName := fs.String("fullname", "", "Nama lengkap")
	email := fs.String("email", "", "Email")
	role := fs.String("role", "admin", "Role: super_admin, admin, operator, viewer (default: admin)")
	isActive := fs.Bool("active", true, "Status aktif (default: true)")
	dbPath := fs.String("db", "./data/database.db", "Path ke file database SQLite")

	fs.Usage = func() {
		fmt.Println("Usage: swantara-cli create-user [options]")
		fmt.Println()
		fmt.Println("Options:")
		fs.PrintDefaults()
		fmt.Println()
		fmt.Println("Example:")
		fmt.Println("  swantara-cli create-user -username admin -password secret123 -fullname \"Super Admin\" -email admin@example.com -role super_admin")
	}

	if err := fs.Parse(args); err != nil {
		os.Exit(1)
	}

	// Validasi input
	if *username == "" {
		fmt.Println("❌ Username wajib diisi (-username)")
		fs.Usage()
		os.Exit(1)
	}
	if *password == "" {
		fmt.Println("❌ Password wajib diisi (-password)")
		fs.Usage()
		os.Exit(1)
	}

	// Validasi role
	validRoles := map[string]bool{
		"super_admin": true,
		"admin":       true,
		"operator":    true,
		"viewer":      true,
	}
	if !validRoles[*role] {
		fmt.Printf("❌ Role tidak valid: %s. Pilihan: super_admin, admin, operator, viewer\n", *role)
		os.Exit(1)
	}

	// Buka database
	db, err := openDB(*dbPath)
	if err != nil {
		log.Fatalf("❌ Gagal membuka database: %v", err)
	}
	defer db.Close()

	// Cek apakah username sudah ada
	var existingID int64
	err = db.QueryRow("SELECT id FROM users WHERE username = ?", *username).Scan(&existingID)
	if err == nil {
		fmt.Printf("❌ Username '%s' sudah digunakan (ID: %d)\n", *username, existingID)
		os.Exit(1)
	}

	// Hash password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("❌ Gagal menghash password: %v", err)
	}

	// Insert user
	now := time.Now()
	result, err := db.Exec(`
		INSERT INTO users (username, password_hash, full_name, email, role, is_active, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`, *username, string(passwordHash), *fullName, *email, *role, *isActive, now, now)
	if err != nil {
		log.Fatalf("❌ Gagal membuat user: %v", err)
	}

	id, _ := result.LastInsertId()

	fmt.Println("✅ User berhasil dibuat!")
	fmt.Println(strings.Repeat("-", 40))
	fmt.Printf("  ID       : %d\n", id)
	fmt.Printf("  Username : %s\n", *username)
	fmt.Printf("  Full Name: %s\n", *fullName)
	fmt.Printf("  Email    : %s\n", *email)
	fmt.Printf("  Role     : %s\n", *role)
	fmt.Printf("  Active   : %v\n", *isActive)
	fmt.Println(strings.Repeat("-", 40))
}

func handleResetPassword(args []string) {
	fs := flag.NewFlagSet("reset-password", flag.ExitOnError)

	username := fs.String("username", "", "Username yang akan di-reset passwordnya (wajib)")
	password := fs.String("password", "", "Password baru (wajib)")
	dbPath := fs.String("db", "./data/database.db", "Path ke file database SQLite")

	fs.Usage = func() {
		fmt.Println("Usage: swantara-cli reset-password [options]")
		fmt.Println()
		fmt.Println("Options:")
		fs.PrintDefaults()
		fmt.Println()
		fmt.Println("Example:")
		fmt.Println("  swantara-cli reset-password -username admin -password newsecret123")
	}

	if err := fs.Parse(args); err != nil {
		os.Exit(1)
	}

	// Validasi input
	if *username == "" {
		fmt.Println("❌ Username wajib diisi (-username)")
		fs.Usage()
		os.Exit(1)
	}
	if *password == "" {
		fmt.Println("❌ Password baru wajib diisi (-password)")
		fs.Usage()
		os.Exit(1)
	}

	// Buka database
	db, err := openDB(*dbPath)
	if err != nil {
		log.Fatalf("❌ Gagal membuka database: %v", err)
	}
	defer db.Close()

	// Cek apakah username ada
	var userID int64
	var userRole string
	err = db.QueryRow("SELECT id, role FROM users WHERE username = ?", *username).Scan(&userID, &userRole)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("❌ User dengan username '%s' tidak ditemukan\n", *username)
		} else {
			fmt.Printf("❌ Gagal mencari user: %v\n", err)
		}
		os.Exit(1)
	}

	// Hash password baru
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("❌ Gagal menghash password: %v", err)
	}

	// Update password
	now := time.Now()
	_, err = db.Exec("UPDATE users SET password_hash = ?, updated_at = ? WHERE username = ?",
		string(passwordHash), now, *username)
	if err != nil {
		log.Fatalf("❌ Gagal mengupdate password: %v", err)
	}

	fmt.Println("✅ Password berhasil di-reset!")
	fmt.Println(strings.Repeat("-", 40))
	fmt.Printf("  ID       : %d\n", userID)
	fmt.Printf("  Username : %s\n", *username)
	fmt.Printf("  Role     : %s\n", userRole)
	fmt.Println(strings.Repeat("-", 40))
}

func openDB(dbPath string) (*sql.DB, error) {
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("file database tidak ditemukan: %s", dbPath)
	}

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, fmt.Errorf("gagal membuka database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("gagal ping database: %w", err)
	}

	return db, nil
}
