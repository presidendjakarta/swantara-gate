package model

import "time"

// VirtualDirectory merepresentasikan konfigurasi route API di gateway
type VirtualDirectory struct {
	ID                  int64     `json:"id"`
	VirtualHostID       int64     `json:"virtual_host_id"`
	SourcePath          string    `json:"source_path"`
	TargetPath          string    `json:"target_path"`
	MatchType           string    `json:"match_type"`
	StripPrefix         bool      `json:"strip_prefix"`
	PreserveHostHeader  bool      `json:"preserve_host_header"`
	AuthType            string    `json:"auth_type"`
	IsActive            bool      `json:"is_active"`
	ProxyTimeoutSeconds int       `json:"proxy_timeout_seconds"`
	RetryCount          int       `json:"retry_count"`
	RetryDelayMs        int       `json:"retry_delay_ms"`
	MaxRequestSizeMB    int       `json:"max_request_size_mb"`
	WebsocketEnabled    bool      `json:"websocket_enabled"`
	CacheEnabled        bool      `json:"cache_enabled"`
	CacheTTLSeconds     int       `json:"cache_ttl_seconds"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`

	// Data join
	VHostName string   `json:"vhost_name,omitempty"`
	Methods   []string `json:"methods,omitempty"`
}

// CreateVirtualDirectoryRequest request untuk membuat virtual directory baru
type CreateVirtualDirectoryRequest struct {
	VirtualHostID       int64  `json:"virtual_host_id" validate:"required"`
	SourcePath          string `json:"source_path" validate:"required"`
	TargetPath          string `json:"target_path" validate:"required"`
	MatchType           string `json:"match_type"`
	StripPrefix         bool   `json:"strip_prefix"`
	PreserveHostHeader  bool   `json:"preserve_host_header"`
	AuthType            string `json:"auth_type"`
	IsActive            bool   `json:"is_active"`
	ProxyTimeoutSeconds int    `json:"proxy_timeout_seconds"`
	RetryCount          int    `json:"retry_count"`
	RetryDelayMs        int    `json:"retry_delay_ms"`
	MaxRequestSizeMB    int    `json:"max_request_size_mb"`
	WebsocketEnabled    bool   `json:"websocket_enabled"`
	CacheEnabled        bool   `json:"cache_enabled"`
	CacheTTLSeconds     int    `json:"cache_ttl_seconds"`
}

// UpdateVirtualDirectoryRequest request untuk update virtual directory
type UpdateVirtualDirectoryRequest struct {
	SourcePath          string `json:"source_path"`
	TargetPath          string `json:"target_path"`
	MatchType           string `json:"match_type"`
	StripPrefix         bool   `json:"strip_prefix"`
	PreserveHostHeader  bool   `json:"preserve_host_header"`
	AuthType            string `json:"auth_type"`
	IsActive            bool   `json:"is_active"`
	ProxyTimeoutSeconds int    `json:"proxy_timeout_seconds"`
	RetryCount          int    `json:"retry_count"`
	RetryDelayMs        int    `json:"retry_delay_ms"`
	MaxRequestSizeMB    int    `json:"max_request_size_mb"`
	WebsocketEnabled    bool   `json:"websocket_enabled"`
	CacheEnabled        bool   `json:"cache_enabled"`
	CacheTTLSeconds     int    `json:"cache_ttl_seconds"`
}

// VirtualDirectoryMethod merepresentasikan HTTP method yang diizinkan
type VirtualDirectoryMethod struct {
	ID                 int64  `json:"id"`
	VirtualDirectoryID int64  `json:"virtual_directory_id"`
	HTTPMethod         string `json:"http_method"`
}

// SetMethodsRequest request untuk mengatur methods pada virtual directory
type SetMethodsRequest struct {
	Methods []string `json:"methods" validate:"required"`
}
