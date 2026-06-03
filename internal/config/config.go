package config

import (
	"os"
	"strconv"
)

// Config menyimpan semua konfigurasi aplikasi
type Config struct {
	// Konfigurasi Admin Panel
	AdminHTTPPort  int
	AdminHTTPSPort int

	// Konfigurasi Proxy Gateway
	ProxyHTTPPort  int
	ProxyHTTPSPort int

	// Konfigurasi Database
	DatabasePath    string
	DatabaseSQLPath string

	// Konfigurasi SSL
	AdminSSLCertPath string
	AdminSSLKeyPath  string
	ProxySSLCertPath string
	ProxySSLKeyPath  string

	// Konfigurasi Umum
	AppEnv  string
	LogLevel string
	DevMode bool

	// Konfigurasi JWT
	JWTSecret             string
	JWTAccessExpireMinutes int
	JWTRefreshExpireDays  int
}

// LoadConfig membaca konfigurasi dari environment variables
func LoadConfig() *Config {
	return &Config{
		// Konfigurasi Admin Panel - membaca port dari environment
		AdminHTTPPort:  getEnvInt("ADMIN_HTTP_PORT", 8080),
		AdminHTTPSPort: getEnvInt("ADMIN_HTTPS_PORT", 8443),

		// Konfigurasi Proxy Gateway - membaca port dari environment
		ProxyHTTPPort:  getEnvInt("PROXY_HTTP_PORT", 8000),
		ProxyHTTPSPort: getEnvInt("PROXY_HTTPS_PORT", 8440),

		// Konfigurasi Database - membaca path dari environment
		DatabasePath:    getEnvString("DATABASE_PATH", "./data/database.db"),
		DatabaseSQLPath: getEnvString("DATABASE_SQL_PATH", "./data/database.sql"),

		// Konfigurasi SSL - membaca path sertifikat dari environment
		AdminSSLCertPath: getEnvString("ADMIN_SSL_CERT_PATH", "./ssl/admin/cert.pem"),
		AdminSSLKeyPath:  getEnvString("ADMIN_SSL_KEY_PATH", "./ssl/admin/key.pem"),
		ProxySSLCertPath: getEnvString("PROXY_SSL_CERT_PATH", "./ssl/proxy/cert.pem"),
		ProxySSLKeyPath:  getEnvString("PROXY_SSL_KEY_PATH", "./ssl/proxy/key.pem"),

		// Konfigurasi Umum - membaca environment dan log level
		AppEnv:   getEnvString("APP_ENV", "development"),
		LogLevel: getEnvString("LOG_LEVEL", "info"),
		DevMode:  getEnvBool("DEV_MODE", true),

		// Konfigurasi JWT
		JWTSecret:             getEnvString("JWT_SECRET", "swantara-gate-secret-key-change-in-production"),
		JWTAccessExpireMinutes: getEnvInt("JWT_ACCESS_EXPIRE_MINUTES", 30),
		JWTRefreshExpireDays:  getEnvInt("JWT_REFRESH_EXPIRE_DAYS", 7),
	}
}

// getEnvString membaca string dari environment, mengembalikan default jika tidak ada
func getEnvString(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// getEnvInt membaca integer dari environment, mengembalikan default jika tidak ada atau error
func getEnvInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	
	return intValue
}

// getEnvBool membaca boolean dari environment, mengembalikan default jika tidak ada atau error
func getEnvBool(key string, defaultValue bool) bool {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	
	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		return defaultValue
	}
	
	return boolValue
}
