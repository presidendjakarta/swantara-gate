package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "modernc.org/sqlite" // Driver SQLite pure Go (tidak perlu CGO)
)

// DB menyimpan instance koneksi database
var DB *sql.DB

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

	// Menjalankan migrasi dari file SQL
	if err := runMigration(sqlPath); err != nil {
		log.Printf("⚠ Migrasi database gagal (mungkin sudah ada): %v", err)
	} else {
		log.Println("✓ Migrasi database berhasil dijalankan")
	}

	return nil
}

// runMigration menjalankan file SQL untuk membuat tabel-tabel
func runMigration(sqlPath string) error {
	// Membaca file SQL
	sqlBytes, err := os.ReadFile(sqlPath)
	if err != nil {
		// Jika file tidak ada, lewati migrasi (tabel mungkin sudah dibuat manual)
		log.Printf("⚠ File SQL tidak ditemukan: %s, melewatkan migrasi", sqlPath)
		return nil
	}

	// Menjalankan semua statement SQL
	_, err = DB.Exec(string(sqlBytes))
	if err != nil {
		return fmt.Errorf("gagal mengeksekusi SQL migration: %w", err)
	}

	return nil
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
