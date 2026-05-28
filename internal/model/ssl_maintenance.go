package model

import "time"

// =========================================================
// ACME ACCOUNTS
// =========================================================

// ACMEAccount merepresentasikan akun ACME (Let's Encrypt)
type ACMEAccount struct {
	ID             int64     `json:"id"`
	Email          string    `json:"email"`
	ProviderURL    string    `json:"provider_url"`
	AccountKeyPath string    `json:"account_key_path"`
	IsDefault      bool      `json:"is_default"`
	CreatedAt      time.Time `json:"created_at"`
}

// CreateACMEAccountRequest request untuk membuat ACME account baru
type CreateACMEAccountRequest struct {
	Email          string `json:"email" validate:"required"`
	ProviderURL    string `json:"provider_url" validate:"required"`
	AccountKeyPath string `json:"account_key_path" validate:"required"`
	IsDefault      bool   `json:"is_default"`
}

// UpdateACMEAccountRequest request untuk update ACME account
type UpdateACMEAccountRequest struct {
	Email          string `json:"email"`
	ProviderURL    string `json:"provider_url"`
	AccountKeyPath string `json:"account_key_path"`
	IsDefault      bool   `json:"is_default"`
}

// =========================================================
// SSL CERTIFICATES
// =========================================================

// SSLCertificate merepresentasikan sertifikat SSL
type SSLCertificate struct {
	ID              int64     `json:"id"`
	ACMEAccountID   *int64    `json:"acme_account_id"`
	Provider        string    `json:"provider"`
	ChallengeType   string    `json:"challenge_type"`
	CertificatePath string    `json:"certificate_path"`
	PrivateKeyPath  string    `json:"private_key_path"`
	ChainPath       string    `json:"chain_path"`
	AutoRenew       bool      `json:"auto_renew"`
	RenewBeforeDays int       `json:"renew_before_days"`
	LastRenewAt     *string   `json:"last_renew_at"`
	ExpiredAt       *string   `json:"expired_at"`
	RenewStatus     string    `json:"renew_status"`
	LastError       string    `json:"last_error"`
	IsActive        bool      `json:"is_active"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       string    `json:"updated_at"`

	// Join data
	ACMEEmail string `json:"acme_email,omitempty"`
}

// CreateSSLCertificateRequest request untuk membuat SSL certificate
type CreateSSLCertificateRequest struct {
	ACMEAccountID   *int64 `json:"acme_account_id"`
	Provider        string `json:"provider"`
	ChallengeType   string `json:"challenge_type"`
	CertificatePath string `json:"certificate_path" validate:"required"`
	PrivateKeyPath  string `json:"private_key_path" validate:"required"`
	ChainPath       string `json:"chain_path"`
	AutoRenew       bool   `json:"auto_renew"`
	RenewBeforeDays int    `json:"renew_before_days"`
	ExpiredAt       string `json:"expired_at"`
	IsActive        bool   `json:"is_active"`
}

// UpdateSSLCertificateRequest request untuk update SSL certificate
type UpdateSSLCertificateRequest struct {
	ACMEAccountID   *int64 `json:"acme_account_id"`
	Provider        string `json:"provider"`
	ChallengeType   string `json:"challenge_type"`
	CertificatePath string `json:"certificate_path"`
	PrivateKeyPath  string `json:"private_key_path"`
	ChainPath       string `json:"chain_path"`
	AutoRenew       bool   `json:"auto_renew"`
	RenewBeforeDays int    `json:"renew_before_days"`
	ExpiredAt       string `json:"expired_at"`
	RenewStatus     string `json:"renew_status"`
	IsActive        bool   `json:"is_active"`
}

// =========================================================
// CERTIFICATE DOMAINS
// =========================================================

// CertificateDomain merepresentasikan domain pada sertifikat SSL
type CertificateDomain struct {
	ID               int64     `json:"id"`
	SSLCertificateID int64     `json:"ssl_certificate_id"`
	DomainName       string    `json:"domain_name"`
	IsWildcard       bool      `json:"is_wildcard"`
	CreatedAt        time.Time `json:"created_at"`
}

// CreateCertificateDomainRequest request untuk membuat certificate domain
type CreateCertificateDomainRequest struct {
	SSLCertificateID int64  `json:"ssl_certificate_id" validate:"required"`
	DomainName       string `json:"domain_name" validate:"required"`
	IsWildcard       bool   `json:"is_wildcard"`
}

// UpdateCertificateDomainRequest request untuk update certificate domain
type UpdateCertificateDomainRequest struct {
	DomainName string `json:"domain_name"`
	IsWildcard bool   `json:"is_wildcard"`
}

// =========================================================
// SSL CERTIFICATE BINDINGS
// =========================================================

// SSLCertificateBinding merepresentasikan binding sertifikat ke host/vhost
type SSLCertificateBinding struct {
	ID               int64     `json:"id"`
	SSLCertificateID int64     `json:"ssl_certificate_id"`
	BindingType      string    `json:"binding_type"`
	HostID           *int64    `json:"host_id"`
	VirtualHostID    *int64    `json:"virtual_host_id"`
	IsDefault        bool      `json:"is_default"`
	Priority         int       `json:"priority"`
	CreatedAt        time.Time `json:"created_at"`

	// Join data
	HostName  string `json:"host_name,omitempty"`
	VHostName string `json:"vhost_name,omitempty"`
}

// CreateSSLCertificateBindingRequest request untuk membuat binding
type CreateSSLCertificateBindingRequest struct {
	SSLCertificateID int64  `json:"ssl_certificate_id" validate:"required"`
	BindingType      string `json:"binding_type" validate:"required"`
	HostID           *int64 `json:"host_id"`
	VirtualHostID    *int64 `json:"virtual_host_id"`
	IsDefault        bool   `json:"is_default"`
	Priority         int    `json:"priority"`
}

// UpdateSSLCertificateBindingRequest request untuk update binding
type UpdateSSLCertificateBindingRequest struct {
	BindingType   string `json:"binding_type"`
	HostID        *int64 `json:"host_id"`
	VirtualHostID *int64 `json:"virtual_host_id"`
	IsDefault     bool   `json:"is_default"`
	Priority      int    `json:"priority"`
}

// =========================================================
// TLS OPTIONS
// =========================================================

// TLSOption merepresentasikan opsi TLS per host/vhost
type TLSOption struct {
	ID            int64     `json:"id"`
	BindingType   string    `json:"binding_type"`
	HostID        *int64    `json:"host_id"`
	VirtualHostID *int64    `json:"virtual_host_id"`
	MinTLSVersion string    `json:"min_tls_version"`
	HTTP2Enabled  bool      `json:"http2_enabled"`
	HSTSEnabled   bool      `json:"hsts_enabled"`
	HSTSMaxAge    int       `json:"hsts_max_age"`
	CreatedAt     time.Time `json:"created_at"`

	// Join data
	HostName  string `json:"host_name,omitempty"`
	VHostName string `json:"vhost_name,omitempty"`
}

// CreateTLSOptionRequest request untuk membuat TLS option
type CreateTLSOptionRequest struct {
	BindingType   string `json:"binding_type" validate:"required"`
	HostID        *int64 `json:"host_id"`
	VirtualHostID *int64 `json:"virtual_host_id"`
	MinTLSVersion string `json:"min_tls_version"`
	HTTP2Enabled  bool   `json:"http2_enabled"`
	HSTSEnabled   bool   `json:"hsts_enabled"`
	HSTSMaxAge    int    `json:"hsts_max_age"`
}

// UpdateTLSOptionRequest request untuk update TLS option
type UpdateTLSOptionRequest struct {
	BindingType   string `json:"binding_type"`
	HostID        *int64 `json:"host_id"`
	VirtualHostID *int64 `json:"virtual_host_id"`
	MinTLSVersion string `json:"min_tls_version"`
	HTTP2Enabled  bool   `json:"http2_enabled"`
	HSTSEnabled   bool   `json:"hsts_enabled"`
	HSTSMaxAge    int    `json:"hsts_max_age"`
}

// =========================================================
// SERVICE DISCOVERY
// =========================================================

// ServiceDiscovery merepresentasikan konfigurasi service discovery
type ServiceDiscovery struct {
	ID                     int64     `json:"id"`
	VirtualHostID          int64     `json:"virtual_host_id"`
	Provider               string    `json:"provider"`
	EndpointURL            string    `json:"endpoint_url"`
	RefreshIntervalSeconds int       `json:"refresh_interval_seconds"`
	IsActive               bool      `json:"is_active"`
	CreatedAt              time.Time `json:"created_at"`

	// Join data
	VHostName string `json:"vhost_name,omitempty"`
}

// CreateServiceDiscoveryRequest request untuk membuat service discovery
type CreateServiceDiscoveryRequest struct {
	VirtualHostID          int64  `json:"virtual_host_id" validate:"required"`
	Provider               string `json:"provider" validate:"required"`
	EndpointURL            string `json:"endpoint_url" validate:"required"`
	RefreshIntervalSeconds int    `json:"refresh_interval_seconds"`
	IsActive               bool   `json:"is_active"`
}

// UpdateServiceDiscoveryRequest request untuk update service discovery
type UpdateServiceDiscoveryRequest struct {
	Provider               string `json:"provider"`
	EndpointURL            string `json:"endpoint_url"`
	RefreshIntervalSeconds int    `json:"refresh_interval_seconds"`
	IsActive               bool   `json:"is_active"`
}

// =========================================================
// CONFIG VERSIONS
// =========================================================

// ConfigVersion merepresentasikan versi konfigurasi
type ConfigVersion struct {
	ID            int64     `json:"id"`
	ConfigName    string    `json:"config_name"`
	VersionNumber int       `json:"version_number"`
	ChangedBy     string    `json:"changed_by"`
	Notes         string    `json:"notes"`
	CreatedAt     time.Time `json:"created_at"`
}

// CreateConfigVersionRequest request untuk membuat config version
type CreateConfigVersionRequest struct {
	ConfigName    string `json:"config_name" validate:"required"`
	VersionNumber int    `json:"version_number" validate:"required"`
	ChangedBy     string `json:"changed_by"`
	Notes         string `json:"notes"`
}

// UpdateConfigVersionRequest request untuk update config version
type UpdateConfigVersionRequest struct {
	ConfigName    string `json:"config_name"`
	VersionNumber int    `json:"version_number"`
	ChangedBy     string `json:"changed_by"`
	Notes         string `json:"notes"`
}

// =========================================================
// MAINTENANCE WINDOWS
// =========================================================

// MaintenanceWindow merepresentasikan jadwal maintenance
type MaintenanceWindow struct {
	ID                      int64     `json:"id"`
	VirtualHostID           *int64    `json:"virtual_host_id"`
	Title                   string    `json:"title"`
	StartAt                 string    `json:"start_at"`
	EndAt                   string    `json:"end_at"`
	MaintenanceResponseCode int       `json:"maintenance_response_code"`
	MaintenanceMessage      string    `json:"maintenance_message"`
	IsActive                bool      `json:"is_active"`
	CreatedAt               time.Time `json:"created_at"`

	// Join data
	VHostName string `json:"vhost_name,omitempty"`
}

// CreateMaintenanceWindowRequest request untuk membuat maintenance window
type CreateMaintenanceWindowRequest struct {
	VirtualHostID           *int64 `json:"virtual_host_id"`
	Title                   string `json:"title" validate:"required"`
	StartAt                 string `json:"start_at" validate:"required"`
	EndAt                   string `json:"end_at" validate:"required"`
	MaintenanceResponseCode int    `json:"maintenance_response_code"`
	MaintenanceMessage      string `json:"maintenance_message"`
	IsActive                bool   `json:"is_active"`
}

// UpdateMaintenanceWindowRequest request untuk update maintenance window
type UpdateMaintenanceWindowRequest struct {
	VirtualHostID           *int64 `json:"virtual_host_id"`
	Title                   string `json:"title"`
	StartAt                 string `json:"start_at"`
	EndAt                   string `json:"end_at"`
	MaintenanceResponseCode int    `json:"maintenance_response_code"`
	MaintenanceMessage      string `json:"maintenance_message"`
	IsActive                bool   `json:"is_active"`
}
