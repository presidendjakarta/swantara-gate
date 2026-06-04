package model

import "time"

// JWTProvider merepresentasikan master data provider JWT (reusable)
type JWTProvider struct {
	ID               int64     `json:"id"`
	Name             string    `json:"name"`
	Description      string    `json:"description,omitempty"`
	Algorithm        string    `json:"algorithm"`
	JWTSecret        string    `json:"jwt_secret"`
	Issuer           string    `json:"issuer,omitempty"`
	Audience         string    `json:"audience,omitempty"`
	ExpiredInSeconds int       `json:"expired_in_seconds"`
	RequireExp       bool      `json:"require_exp"`
	RequireIat       bool      `json:"require_iat"`
	IsActive         bool      `json:"is_active"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`

	// Statistics
	UsedByCount int64 `json:"used_by_count,omitempty"` // Jumlah vdir yang menggunakan provider ini
}

// CreateJWTProviderRequest request untuk membuat JWT provider baru
type CreateJWTProviderRequest struct {
	Name             string `json:"name" validate:"required"`
	Description      string `json:"description"`
	Algorithm        string `json:"algorithm"`
	JWTSecret        string `json:"jwt_secret" validate:"required"`
	Issuer           string `json:"issuer"`
	Audience         string `json:"audience"`
	ExpiredInSeconds int    `json:"expired_in_seconds"`
	RequireExp       bool   `json:"require_exp"`
	RequireIat       bool   `json:"require_iat"`
	IsActive         bool   `json:"is_active"`
}

// UpdateJWTProviderRequest request untuk update JWT provider
type UpdateJWTProviderRequest struct {
	Name             string `json:"name"`
	Description      string `json:"description"`
	Algorithm        string `json:"algorithm"`
	JWTSecret        string `json:"jwt_secret"`
	Issuer           string `json:"issuer"`
	Audience         string `json:"audience"`
	ExpiredInSeconds int    `json:"expired_in_seconds"`
	RequireExp       bool   `json:"require_exp"`
	RequireIat       bool   `json:"require_iat"`
	IsActive         bool   `json:"is_active"`
}

// VirtualDirectoryJWTProviderMapping mapping antara virtual directory dan JWT provider
type VirtualDirectoryJWTProviderMapping struct {
	VirtualDirectoryID int64     `json:"virtual_directory_id"`
	JWTProviderID      int64     `json:"jwt_provider_id"`
	CreatedAt          time.Time `json:"created_at"`

	// Data join
	SourcePath   string `json:"source_path,omitempty"`
	TargetPath   string `json:"target_path,omitempty"`
	ProviderName string `json:"provider_name,omitempty"`
	Algorithm    string `json:"algorithm,omitempty"`
}

// JWTConfig merepresentasikan konfigurasi JWT per route (legacy)
type JWTConfig struct {
	ID                 int64     `json:"id"`
	VirtualDirectoryID int64     `json:"virtual_directory_id"`
	Algorithm          string    `json:"algorithm"`
	JWTSecret          string    `json:"jwt_secret"`
	Issuer             string    `json:"issuer"`
	Audience           string    `json:"audience"`
	ExpiredInSeconds   int       `json:"expired_in_seconds"`
	ClockSkewSeconds   int       `json:"clock_skew_seconds"`
	RequireExp         bool      `json:"require_exp"`
	RequireIat         bool      `json:"require_iat"`
	IsActive           bool      `json:"is_active"`
	CreatedAt          time.Time `json:"created_at"`

	// Data join
	SourcePath string `json:"source_path,omitempty"`
}

// CreateJWTConfigRequest request untuk membuat JWT config baru
type CreateJWTConfigRequest struct {
	VirtualDirectoryID int64  `json:"virtual_directory_id" validate:"required"`
	Algorithm          string `json:"algorithm"`
	JWTSecret          string `json:"jwt_secret" validate:"required"`
	Issuer             string `json:"issuer"`
	Audience           string `json:"audience"`
	ExpiredInSeconds   int    `json:"expired_in_seconds"`
	ClockSkewSeconds   int    `json:"clock_skew_seconds"`
	RequireExp         bool   `json:"require_exp"`
	RequireIat         bool   `json:"require_iat"`
	IsActive           bool   `json:"is_active"`
}

// UpdateJWTConfigRequest request untuk update JWT config
type UpdateJWTConfigRequest struct {
	Algorithm        string `json:"algorithm"`
	JWTSecret        string `json:"jwt_secret"`
	Issuer           string `json:"issuer"`
	Audience         string `json:"audience"`
	ExpiredInSeconds int    `json:"expired_in_seconds"`
	ClockSkewSeconds int    `json:"clock_skew_seconds"`
	RequireExp       bool   `json:"require_exp"`
	RequireIat       bool   `json:"require_iat"`
	IsActive         bool   `json:"is_active"`
}

// ExternalAuthProvider merepresentasikan master data provider autentikasi eksternal
type ExternalAuthProvider struct {
	ID                    int64     `json:"id"`
	Name                  string    `json:"name"`
	Description           string    `json:"description,omitempty"`
	AuthURL               string    `json:"auth_url"`
	HTTPMethod            string    `json:"http_method"`
	RequestTimeoutSeconds int       `json:"request_timeout_seconds"`
	SendHeaders           bool      `json:"send_headers"`
	SendBody              bool      `json:"send_body"`
	SuccessKey            string    `json:"success_key"`
	SuccessValue          string    `json:"success_value"`
	MessageKey            string    `json:"message_key"`
	TokenKey              string    `json:"token_key"`
	IsActive              bool      `json:"is_active"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`

	// Statistics
	UsedByCount int64 `json:"used_by_count,omitempty"` // Jumlah vdir yang menggunakan provider ini
}

// CreateExternalAuthProviderRequest request untuk membuat provider baru
type CreateExternalAuthProviderRequest struct {
	Name                  string `json:"name" validate:"required"`
	Description           string `json:"description"`
	AuthURL               string `json:"auth_url" validate:"required"`
	HTTPMethod            string `json:"http_method"`
	RequestTimeoutSeconds int    `json:"request_timeout_seconds"`
	SendHeaders           bool   `json:"send_headers"`
	SendBody              bool   `json:"send_body"`
	SuccessKey            string `json:"success_key"`
	SuccessValue          string `json:"success_value"`
	MessageKey            string `json:"message_key"`
	TokenKey              string `json:"token_key"`
	IsActive              bool   `json:"is_active"`
}

// UpdateExternalAuthProviderRequest request untuk update provider
type UpdateExternalAuthProviderRequest struct {
	Name                  string `json:"name"`
	Description           string `json:"description"`
	AuthURL               string `json:"auth_url"`
	HTTPMethod            string `json:"http_method"`
	RequestTimeoutSeconds int    `json:"request_timeout_seconds"`
	SendHeaders           bool   `json:"send_headers"`
	SendBody              bool   `json:"send_body"`
	SuccessKey            string `json:"success_key"`
	SuccessValue          string `json:"success_value"`
	MessageKey            string `json:"message_key"`
	TokenKey              string `json:"token_key"`
	IsActive              bool   `json:"is_active"`
}

// VirtualDirectoryExternalAuthMapping mapping antara virtual directory dan auth provider
type VirtualDirectoryExternalAuthMapping struct {
	VirtualDirectoryID       int64     `json:"virtual_directory_id"`
	ExternalAuthProviderID   int64     `json:"external_auth_provider_id"`
	CreatedAt                time.Time `json:"created_at"`

	// Data join
	SourcePath  string `json:"source_path,omitempty"`
	TargetPath  string `json:"target_path,omitempty"`
	ProviderName string `json:"provider_name,omitempty"`
	AuthURL     string `json:"auth_url,omitempty"`
}

// ExternalAuth merepresentasikan konfigurasi autentikasi eksternal (backward compatibility)
type ExternalAuth struct {
	ID                    int64     `json:"id"`
	VirtualDirectoryID    int64     `json:"virtual_directory_id"`
	AuthURL               string    `json:"auth_url"`
	HTTPMethod            string    `json:"http_method"`
	RequestTimeoutSeconds int       `json:"request_timeout_seconds"`
	SendHeaders           bool      `json:"send_headers"`
	SendBody              bool      `json:"send_body"`
	SuccessKey            string    `json:"success_key"`
	SuccessValue          string    `json:"success_value"`
	MessageKey            string    `json:"message_key"`
	TokenKey              string    `json:"token_key"`
	IsActive              bool      `json:"is_active"`
	CreatedAt             time.Time `json:"created_at"`

	// Data join
	SourcePath string `json:"source_path,omitempty"`
}

// CreateExternalAuthRequest request untuk membuat external auth baru
type CreateExternalAuthRequest struct {
	VirtualDirectoryID    int64  `json:"virtual_directory_id" validate:"required"`
	AuthURL               string `json:"auth_url" validate:"required"`
	HTTPMethod            string `json:"http_method"`
	RequestTimeoutSeconds int    `json:"request_timeout_seconds"`
	SendHeaders           bool   `json:"send_headers"`
	SendBody              bool   `json:"send_body"`
	SuccessKey            string `json:"success_key"`
	SuccessValue          string `json:"success_value"`
	MessageKey            string `json:"message_key"`
	TokenKey              string `json:"token_key"`
	IsActive              bool   `json:"is_active"`
}

// UpdateExternalAuthRequest request untuk update external auth
type UpdateExternalAuthRequest struct {
	AuthURL               string `json:"auth_url"`
	HTTPMethod            string `json:"http_method"`
	RequestTimeoutSeconds int    `json:"request_timeout_seconds"`
	SendHeaders           bool   `json:"send_headers"`
	SendBody              bool   `json:"send_body"`
	SuccessKey            string `json:"success_key"`
	SuccessValue          string `json:"success_value"`
	MessageKey            string `json:"message_key"`
	TokenKey              string `json:"token_key"`
	IsActive              bool   `json:"is_active"`
}

// RateLimit merepresentasikan konfigurasi rate limiting per route
type RateLimit struct {
	ID                   int64     `json:"id"`
	VirtualDirectoryID   int64     `json:"virtual_directory_id"`
	LimitBy              string    `json:"limit_by"`
	RequestsPerMinute    int       `json:"requests_per_minute"`
	Burst                int       `json:"burst"`
	BlockDurationSeconds int       `json:"block_duration_seconds"`
	IsActive             bool      `json:"is_active"`
	CreatedAt            time.Time `json:"created_at"`

	// Data join
	SourcePath string `json:"source_path,omitempty"`
}

// CreateRateLimitRequest request untuk membuat rate limit baru
type CreateRateLimitRequest struct {
	VirtualDirectoryID   int64  `json:"virtual_directory_id" validate:"required"`
	LimitBy              string `json:"limit_by"`
	RequestsPerMinute    int    `json:"requests_per_minute"`
	Burst                int    `json:"burst"`
	BlockDurationSeconds int    `json:"block_duration_seconds"`
	IsActive             bool   `json:"is_active"`
}

// UpdateRateLimitRequest request untuk update rate limit
type UpdateRateLimitRequest struct {
	LimitBy              string `json:"limit_by"`
	RequestsPerMinute    int    `json:"requests_per_minute"`
	Burst                int    `json:"burst"`
	BlockDurationSeconds int    `json:"block_duration_seconds"`
	IsActive             bool   `json:"is_active"`
}

// CORSConfig merepresentasikan konfigurasi CORS per route
type CORSConfig struct {
	ID                 int64     `json:"id"`
	VirtualDirectoryID int64     `json:"virtual_directory_id"`
	AllowedOrigins     string    `json:"allowed_origins"`
	AllowedMethods     string    `json:"allowed_methods"`
	AllowedHeaders     string    `json:"allowed_headers"`
	ExposedHeaders     string    `json:"exposed_headers"`
	AllowCredentials   bool      `json:"allow_credentials"`
	MaxAgeSeconds      int       `json:"max_age_seconds"`
	IsActive           bool      `json:"is_active"`
	CreatedAt          time.Time `json:"created_at"`

	// Data join
	SourcePath string `json:"source_path,omitempty"`
}

// CreateCORSConfigRequest request untuk membuat CORS config baru
type CreateCORSConfigRequest struct {
	VirtualDirectoryID int64  `json:"virtual_directory_id" validate:"required"`
	AllowedOrigins     string `json:"allowed_origins"`
	AllowedMethods     string `json:"allowed_methods"`
	AllowedHeaders     string `json:"allowed_headers"`
	ExposedHeaders     string `json:"exposed_headers"`
	AllowCredentials   bool   `json:"allow_credentials"`
	MaxAgeSeconds      int    `json:"max_age_seconds"`
	IsActive           bool   `json:"is_active"`
}

// UpdateCORSConfigRequest request untuk update CORS config
type UpdateCORSConfigRequest struct {
	AllowedOrigins   string `json:"allowed_origins"`
	AllowedMethods   string `json:"allowed_methods"`
	AllowedHeaders   string `json:"allowed_headers"`
	ExposedHeaders   string `json:"exposed_headers"`
	AllowCredentials bool   `json:"allow_credentials"`
	MaxAgeSeconds    int    `json:"max_age_seconds"`
	IsActive         bool   `json:"is_active"`
}

// CircuitBreaker merepresentasikan konfigurasi circuit breaker per route
type CircuitBreaker struct {
	ID                     int64     `json:"id"`
	VirtualDirectoryID     int64     `json:"virtual_directory_id"`
	Enabled                bool      `json:"enabled"`
	FailureThreshold       int       `json:"failure_threshold"`
	RecoveryTimeoutSeconds int       `json:"recovery_timeout_seconds"`
	HalfOpenMaxRequests    int       `json:"half_open_max_requests"`
	CreatedAt              time.Time `json:"created_at"`

	// Data join
	SourcePath string `json:"source_path,omitempty"`
}

// CreateCircuitBreakerRequest request untuk membuat circuit breaker baru
type CreateCircuitBreakerRequest struct {
	VirtualDirectoryID     int64 `json:"virtual_directory_id" validate:"required"`
	Enabled                bool  `json:"enabled"`
	FailureThreshold       int   `json:"failure_threshold"`
	RecoveryTimeoutSeconds int   `json:"recovery_timeout_seconds"`
	HalfOpenMaxRequests    int   `json:"half_open_max_requests"`
}

// UpdateCircuitBreakerRequest request untuk update circuit breaker
type UpdateCircuitBreakerRequest struct {
	Enabled                bool `json:"enabled"`
	FailureThreshold       int  `json:"failure_threshold"`
	RecoveryTimeoutSeconds int  `json:"recovery_timeout_seconds"`
	HalfOpenMaxRequests    int  `json:"half_open_max_requests"`
}

// IPWhitelist merepresentasikan IP yang diperbolehkan mengakses route
type IPWhitelist struct {
	ID                 int64     `json:"id"`
	VirtualDirectoryID int64     `json:"virtual_directory_id"`
	IPAddress          string    `json:"ip_address"`
	Description        string    `json:"description"`
	IsActive           bool      `json:"is_active"`
	CreatedAt          time.Time `json:"created_at"`

	// Data join
	SourcePath string `json:"source_path,omitempty"`
}

// CreateIPWhitelistRequest request untuk menambahkan IP ke whitelist
type CreateIPWhitelistRequest struct {
	VirtualDirectoryID int64  `json:"virtual_directory_id" validate:"required"`
	IPAddress          string `json:"ip_address" validate:"required"`
	Description        string `json:"description"`
	IsActive           bool   `json:"is_active"`
}

// UpdateIPWhitelistRequest request untuk update whitelist entry
type UpdateIPWhitelistRequest struct {
	IPAddress   string `json:"ip_address"`
	Description string `json:"description"`
	IsActive    bool   `json:"is_active"`
}

// IPBlacklist merepresentasikan IP yang diblokir dari route
type IPBlacklist struct {
	ID                 int64      `json:"id"`
	VirtualDirectoryID int64      `json:"virtual_directory_id"`
	IPAddress          string     `json:"ip_address"`
	Reason             string     `json:"reason"`
	ExpiredAt          *time.Time `json:"expired_at,omitempty"`
	IsActive           bool       `json:"is_active"`
	CreatedAt          time.Time  `json:"created_at"`

	// Data join
	SourcePath string `json:"source_path,omitempty"`
}

// CreateIPBlacklistRequest request untuk menambahkan IP ke blacklist
type CreateIPBlacklistRequest struct {
	VirtualDirectoryID int64  `json:"virtual_directory_id" validate:"required"`
	IPAddress          string `json:"ip_address" validate:"required"`
	Reason             string `json:"reason"`
	ExpiredAt          string `json:"expired_at"`
	IsActive           bool   `json:"is_active"`
}

// UpdateIPBlacklistRequest request untuk update blacklist entry
type UpdateIPBlacklistRequest struct {
	IPAddress string `json:"ip_address"`
	Reason    string `json:"reason"`
	ExpiredAt string `json:"expired_at"`
	IsActive  bool   `json:"is_active"`
}
