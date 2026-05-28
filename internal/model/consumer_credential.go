package model

import "time"

// ConsumerCredential merepresentasikan kredensial autentikasi consumer
type ConsumerCredential struct {
	ID           int64      `json:"id"`
	ConsumerID   int64      `json:"consumer_id"`
	AuthType     string     `json:"auth_type"`
	Username     string     `json:"username,omitempty"`
	PasswordHash string     `json:"-"`
	APIKey       string     `json:"api_key,omitempty"`
	JWTSecret    string     `json:"jwt_secret,omitempty"`
	ExpiredAt    *time.Time `json:"expired_at,omitempty"`
	IsActive     bool       `json:"is_active"`
	CreatedAt    time.Time  `json:"created_at"`

	// Data join
	ConsumerName string `json:"consumer_name,omitempty"`
}

// CreateConsumerCredentialRequest request untuk membuat credential baru
type CreateConsumerCredentialRequest struct {
	ConsumerID int64  `json:"consumer_id" validate:"required"`
	AuthType   string `json:"auth_type" validate:"required"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	APIKey     string `json:"api_key"`
	JWTSecret  string `json:"jwt_secret"`
	ExpiredAt  string `json:"expired_at"`
	IsActive   bool   `json:"is_active"`
}

// UpdateConsumerCredentialRequest request untuk update credential
type UpdateConsumerCredentialRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	APIKey    string `json:"api_key"`
	JWTSecret string `json:"jwt_secret"`
	ExpiredAt string `json:"expired_at"`
	IsActive  bool   `json:"is_active"`
}

// APIKey merepresentasikan API key khusus untuk consumer
type APIKey struct {
	ID                int64      `json:"id"`
	ConsumerID        int64      `json:"consumer_id"`
	Key               string     `json:"api_key"`
	Description       string     `json:"description"`
	ExpiredAt         *time.Time `json:"expired_at,omitempty"`
	RateLimitOverride *int       `json:"rate_limit_override,omitempty"`
	IsActive          bool       `json:"is_active"`
	CreatedAt         time.Time  `json:"created_at"`

	// Data join
	ConsumerName string `json:"consumer_name,omitempty"`
}

// CreateAPIKeyRequest request untuk membuat API key baru
type CreateAPIKeyRequest struct {
	ConsumerID        int64  `json:"consumer_id" validate:"required"`
	Description       string `json:"description"`
	ExpiredAt         string `json:"expired_at"`
	RateLimitOverride *int   `json:"rate_limit_override"`
	IsActive          bool   `json:"is_active"`
}

// UpdateAPIKeyRequest request untuk update API key
type UpdateAPIKeyRequest struct {
	Description       string `json:"description"`
	ExpiredAt         string `json:"expired_at"`
	RateLimitOverride *int   `json:"rate_limit_override"`
	IsActive          bool   `json:"is_active"`
}

// RouteConsumerAccess merepresentasikan akses consumer ke route tertentu
type RouteConsumerAccess struct {
	ID                 int64     `json:"id"`
	VirtualDirectoryID int64     `json:"virtual_directory_id"`
	ConsumerID         int64     `json:"consumer_id"`
	IsActive           bool      `json:"is_active"`
	CreatedAt          time.Time `json:"created_at"`

	// Data join
	ConsumerName string `json:"consumer_name,omitempty"`
	SourcePath   string `json:"source_path,omitempty"`
}

// CreateRouteConsumerAccessRequest request untuk membuat akses baru
type CreateRouteConsumerAccessRequest struct {
	VirtualDirectoryID int64 `json:"virtual_directory_id" validate:"required"`
	ConsumerID         int64 `json:"consumer_id" validate:"required"`
	IsActive           bool  `json:"is_active"`
}

// UpdateRouteConsumerAccessRequest request untuk update akses
type UpdateRouteConsumerAccessRequest struct {
	IsActive bool `json:"is_active"`
}
