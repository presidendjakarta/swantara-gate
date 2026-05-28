package repository

import (
	"database/sql"
	"fmt"

	"github.com/presidendjakarta/swantara-gate/internal/model"
)

// =========================================================
// ACME ACCOUNT REPOSITORY
// =========================================================

type ACMEAccountRepository struct {
	DB *sql.DB
}

func NewACMEAccountRepository(db *sql.DB) *ACMEAccountRepository {
	return &ACMEAccountRepository{DB: db}
}

func (r *ACMEAccountRepository) Create(req *model.CreateACMEAccountRequest) (*model.ACMEAccount, error) {
	query := `INSERT INTO acme_accounts (email, provider_url, account_key_path, is_default) VALUES (?, ?, ?, ?) RETURNING id, created_at`
	var item model.ACMEAccount
	err := r.DB.QueryRow(query, req.Email, req.ProviderURL, req.AccountKeyPath, req.IsDefault).Scan(&item.ID, &item.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("gagal membuat ACME account: %w", err)
	}
	item.Email = req.Email
	item.ProviderURL = req.ProviderURL
	item.AccountKeyPath = req.AccountKeyPath
	item.IsDefault = req.IsDefault
	return &item, nil
}

func (r *ACMEAccountRepository) GetByID(id int64) (*model.ACMEAccount, error) {
	query := `SELECT id, email, provider_url, account_key_path, is_default, created_at FROM acme_accounts WHERE id = ?`
	var item model.ACMEAccount
	err := r.DB.QueryRow(query, id).Scan(&item.ID, &item.Email, &item.ProviderURL, &item.AccountKeyPath, &item.IsDefault, &item.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("ACME account tidak ditemukan")
		}
		return nil, fmt.Errorf("gagal mengambil ACME account: %w", err)
	}
	return &item, nil
}

func (r *ACMEAccountRepository) GetAll(page, limit int) ([]model.ACMEAccount, error) {
	offset := (page - 1) * limit
	query := `SELECT id, email, provider_url, account_key_path, is_default, created_at FROM acme_accounts ORDER BY id DESC LIMIT ? OFFSET ?`
	rows, err := r.DB.Query(query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil ACME accounts: %w", err)
	}
	defer rows.Close()
	var items []model.ACMEAccount
	for rows.Next() {
		var item model.ACMEAccount
		if err := rows.Scan(&item.ID, &item.Email, &item.ProviderURL, &item.AccountKeyPath, &item.IsDefault, &item.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (r *ACMEAccountRepository) Update(id int64, req *model.UpdateACMEAccountRequest) error {
	query := `UPDATE acme_accounts SET email = ?, provider_url = ?, account_key_path = ?, is_default = ? WHERE id = ?`
	_, err := r.DB.Exec(query, req.Email, req.ProviderURL, req.AccountKeyPath, req.IsDefault, id)
	return err
}

func (r *ACMEAccountRepository) Delete(id int64) error {
	_, err := r.DB.Exec("DELETE FROM acme_accounts WHERE id = ?", id)
	return err
}

func (r *ACMEAccountRepository) Count() (int64, error) {
	var count int64
	err := r.DB.QueryRow("SELECT COUNT(*) FROM acme_accounts").Scan(&count)
	return count, err
}

// =========================================================
// SSL CERTIFICATE REPOSITORY
// =========================================================

type SSLCertificateRepository struct {
	DB *sql.DB
}

func NewSSLCertificateRepository(db *sql.DB) *SSLCertificateRepository {
	return &SSLCertificateRepository{DB: db}
}

func (r *SSLCertificateRepository) Create(req *model.CreateSSLCertificateRequest) (*model.SSLCertificate, error) {
	query := `INSERT INTO ssl_certificates (acme_account_id, provider, challenge_type, certificate_path, private_key_path, chain_path, auto_renew, renew_before_days, expired_at, is_active) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?) RETURNING id, created_at`
	var item model.SSLCertificate
	var expAt *string
	if req.ExpiredAt != "" {
		expAt = &req.ExpiredAt
	}
	err := r.DB.QueryRow(query, req.ACMEAccountID, req.Provider, req.ChallengeType, req.CertificatePath, req.PrivateKeyPath, req.ChainPath, req.AutoRenew, req.RenewBeforeDays, expAt, req.IsActive).Scan(&item.ID, &item.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("gagal membuat SSL certificate: %w", err)
	}
	item.ACMEAccountID = req.ACMEAccountID
	item.Provider = req.Provider
	item.ChallengeType = req.ChallengeType
	item.CertificatePath = req.CertificatePath
	item.PrivateKeyPath = req.PrivateKeyPath
	item.ChainPath = req.ChainPath
	item.AutoRenew = req.AutoRenew
	item.RenewBeforeDays = req.RenewBeforeDays
	item.IsActive = req.IsActive
	return &item, nil
}

func (r *SSLCertificateRepository) GetByID(id int64) (*model.SSLCertificate, error) {
	query := `SELECT sc.id, sc.acme_account_id, sc.provider, sc.challenge_type, sc.certificate_path, sc.private_key_path, sc.chain_path, sc.auto_renew, sc.renew_before_days, sc.last_renew_at, sc.expired_at, COALESCE(sc.renew_status, ''), COALESCE(sc.last_error, ''), sc.is_active, sc.created_at, COALESCE(sc.updated_at, sc.created_at), COALESCE(aa.email, '') FROM ssl_certificates sc LEFT JOIN acme_accounts aa ON sc.acme_account_id = aa.id WHERE sc.id = ?`
	var item model.SSLCertificate
	err := r.DB.QueryRow(query, id).Scan(&item.ID, &item.ACMEAccountID, &item.Provider, &item.ChallengeType, &item.CertificatePath, &item.PrivateKeyPath, &item.ChainPath, &item.AutoRenew, &item.RenewBeforeDays, &item.LastRenewAt, &item.ExpiredAt, &item.RenewStatus, &item.LastError, &item.IsActive, &item.CreatedAt, &item.UpdatedAt, &item.ACMEEmail)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("SSL certificate tidak ditemukan")
		}
		return nil, fmt.Errorf("gagal mengambil SSL certificate: %w", err)
	}
	return &item, nil
}

func (r *SSLCertificateRepository) GetAll(page, limit int) ([]model.SSLCertificate, error) {
	offset := (page - 1) * limit
	query := `SELECT sc.id, sc.acme_account_id, sc.provider, sc.challenge_type, sc.certificate_path, sc.private_key_path, sc.chain_path, sc.auto_renew, sc.renew_before_days, sc.last_renew_at, sc.expired_at, COALESCE(sc.renew_status, ''), COALESCE(sc.last_error, ''), sc.is_active, sc.created_at, COALESCE(sc.updated_at, sc.created_at), COALESCE(aa.email, '') FROM ssl_certificates sc LEFT JOIN acme_accounts aa ON sc.acme_account_id = aa.id ORDER BY sc.id DESC LIMIT ? OFFSET ?`
	rows, err := r.DB.Query(query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil SSL certificates: %w", err)
	}
	defer rows.Close()
	var items []model.SSLCertificate
	for rows.Next() {
		var item model.SSLCertificate
		if err := rows.Scan(&item.ID, &item.ACMEAccountID, &item.Provider, &item.ChallengeType, &item.CertificatePath, &item.PrivateKeyPath, &item.ChainPath, &item.AutoRenew, &item.RenewBeforeDays, &item.LastRenewAt, &item.ExpiredAt, &item.RenewStatus, &item.LastError, &item.IsActive, &item.CreatedAt, &item.UpdatedAt, &item.ACMEEmail); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (r *SSLCertificateRepository) Update(id int64, req *model.UpdateSSLCertificateRequest) error {
	query := `UPDATE ssl_certificates SET acme_account_id = ?, provider = ?, challenge_type = ?, certificate_path = ?, private_key_path = ?, chain_path = ?, auto_renew = ?, renew_before_days = ?, expired_at = ?, renew_status = ?, is_active = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?`
	var expAt *string
	if req.ExpiredAt != "" {
		expAt = &req.ExpiredAt
	}
	_, err := r.DB.Exec(query, req.ACMEAccountID, req.Provider, req.ChallengeType, req.CertificatePath, req.PrivateKeyPath, req.ChainPath, req.AutoRenew, req.RenewBeforeDays, expAt, req.RenewStatus, req.IsActive, id)
	return err
}

func (r *SSLCertificateRepository) Delete(id int64) error {
	_, err := r.DB.Exec("DELETE FROM ssl_certificates WHERE id = ?", id)
	return err
}

func (r *SSLCertificateRepository) Count() (int64, error) {
	var count int64
	err := r.DB.QueryRow("SELECT COUNT(*) FROM ssl_certificates").Scan(&count)
	return count, err
}

// =========================================================
// CERTIFICATE DOMAIN REPOSITORY
// =========================================================

type CertificateDomainRepository struct {
	DB *sql.DB
}

func NewCertificateDomainRepository(db *sql.DB) *CertificateDomainRepository {
	return &CertificateDomainRepository{DB: db}
}

func (r *CertificateDomainRepository) Create(req *model.CreateCertificateDomainRequest) (*model.CertificateDomain, error) {
	query := `INSERT INTO certificate_domains (ssl_certificate_id, domain_name, is_wildcard) VALUES (?, ?, ?) RETURNING id, created_at`
	var item model.CertificateDomain
	err := r.DB.QueryRow(query, req.SSLCertificateID, req.DomainName, req.IsWildcard).Scan(&item.ID, &item.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("gagal membuat certificate domain: %w", err)
	}
	item.SSLCertificateID = req.SSLCertificateID
	item.DomainName = req.DomainName
	item.IsWildcard = req.IsWildcard
	return &item, nil
}

func (r *CertificateDomainRepository) GetByID(id int64) (*model.CertificateDomain, error) {
	query := `SELECT id, ssl_certificate_id, domain_name, is_wildcard, created_at FROM certificate_domains WHERE id = ?`
	var item model.CertificateDomain
	err := r.DB.QueryRow(query, id).Scan(&item.ID, &item.SSLCertificateID, &item.DomainName, &item.IsWildcard, &item.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("certificate domain tidak ditemukan")
		}
		return nil, err
	}
	return &item, nil
}

func (r *CertificateDomainRepository) GetAll(page, limit int) ([]model.CertificateDomain, error) {
	offset := (page - 1) * limit
	query := `SELECT id, ssl_certificate_id, domain_name, is_wildcard, created_at FROM certificate_domains ORDER BY id DESC LIMIT ? OFFSET ?`
	rows, err := r.DB.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []model.CertificateDomain
	for rows.Next() {
		var item model.CertificateDomain
		if err := rows.Scan(&item.ID, &item.SSLCertificateID, &item.DomainName, &item.IsWildcard, &item.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (r *CertificateDomainRepository) GetByCertificateID(certID int64) ([]model.CertificateDomain, error) {
	query := `SELECT id, ssl_certificate_id, domain_name, is_wildcard, created_at FROM certificate_domains WHERE ssl_certificate_id = ? ORDER BY id`
	rows, err := r.DB.Query(query, certID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []model.CertificateDomain
	for rows.Next() {
		var item model.CertificateDomain
		if err := rows.Scan(&item.ID, &item.SSLCertificateID, &item.DomainName, &item.IsWildcard, &item.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (r *CertificateDomainRepository) Update(id int64, req *model.UpdateCertificateDomainRequest) error {
	_, err := r.DB.Exec("UPDATE certificate_domains SET domain_name = ?, is_wildcard = ? WHERE id = ?", req.DomainName, req.IsWildcard, id)
	return err
}

func (r *CertificateDomainRepository) Delete(id int64) error {
	_, err := r.DB.Exec("DELETE FROM certificate_domains WHERE id = ?", id)
	return err
}

func (r *CertificateDomainRepository) Count() (int64, error) {
	var count int64
	err := r.DB.QueryRow("SELECT COUNT(*) FROM certificate_domains").Scan(&count)
	return count, err
}

// =========================================================
// SSL CERTIFICATE BINDING REPOSITORY
// =========================================================

type SSLCertificateBindingRepository struct {
	DB *sql.DB
}

func NewSSLCertificateBindingRepository(db *sql.DB) *SSLCertificateBindingRepository {
	return &SSLCertificateBindingRepository{DB: db}
}

func (r *SSLCertificateBindingRepository) Create(req *model.CreateSSLCertificateBindingRequest) (*model.SSLCertificateBinding, error) {
	query := `INSERT INTO ssl_certificate_bindings (ssl_certificate_id, binding_type, host_id, virtual_host_id, is_default, priority) VALUES (?, ?, ?, ?, ?, ?) RETURNING id, created_at`
	var item model.SSLCertificateBinding
	err := r.DB.QueryRow(query, req.SSLCertificateID, req.BindingType, req.HostID, req.VirtualHostID, req.IsDefault, req.Priority).Scan(&item.ID, &item.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("gagal membuat SSL binding: %w", err)
	}
	item.SSLCertificateID = req.SSLCertificateID
	item.BindingType = req.BindingType
	item.HostID = req.HostID
	item.VirtualHostID = req.VirtualHostID
	item.IsDefault = req.IsDefault
	item.Priority = req.Priority
	return &item, nil
}

func (r *SSLCertificateBindingRepository) GetByID(id int64) (*model.SSLCertificateBinding, error) {
	query := `SELECT b.id, b.ssl_certificate_id, b.binding_type, b.host_id, b.virtual_host_id, b.is_default, b.priority, b.created_at, COALESCE(h.host_name, ''), COALESCE(vh.vhost_name, '') FROM ssl_certificate_bindings b LEFT JOIN hosts h ON b.host_id = h.id LEFT JOIN virtual_hosts vh ON b.virtual_host_id = vh.id WHERE b.id = ?`
	var item model.SSLCertificateBinding
	err := r.DB.QueryRow(query, id).Scan(&item.ID, &item.SSLCertificateID, &item.BindingType, &item.HostID, &item.VirtualHostID, &item.IsDefault, &item.Priority, &item.CreatedAt, &item.HostName, &item.VHostName)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("SSL binding tidak ditemukan")
		}
		return nil, err
	}
	return &item, nil
}

func (r *SSLCertificateBindingRepository) GetAll(page, limit int) ([]model.SSLCertificateBinding, error) {
	offset := (page - 1) * limit
	query := `SELECT b.id, b.ssl_certificate_id, b.binding_type, b.host_id, b.virtual_host_id, b.is_default, b.priority, b.created_at, COALESCE(h.host_name, ''), COALESCE(vh.vhost_name, '') FROM ssl_certificate_bindings b LEFT JOIN hosts h ON b.host_id = h.id LEFT JOIN virtual_hosts vh ON b.virtual_host_id = vh.id ORDER BY b.priority ASC, b.id DESC LIMIT ? OFFSET ?`
	rows, err := r.DB.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []model.SSLCertificateBinding
	for rows.Next() {
		var item model.SSLCertificateBinding
		if err := rows.Scan(&item.ID, &item.SSLCertificateID, &item.BindingType, &item.HostID, &item.VirtualHostID, &item.IsDefault, &item.Priority, &item.CreatedAt, &item.HostName, &item.VHostName); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (r *SSLCertificateBindingRepository) Update(id int64, req *model.UpdateSSLCertificateBindingRequest) error {
	query := `UPDATE ssl_certificate_bindings SET binding_type = ?, host_id = ?, virtual_host_id = ?, is_default = ?, priority = ? WHERE id = ?`
	_, err := r.DB.Exec(query, req.BindingType, req.HostID, req.VirtualHostID, req.IsDefault, req.Priority, id)
	return err
}

func (r *SSLCertificateBindingRepository) Delete(id int64) error {
	_, err := r.DB.Exec("DELETE FROM ssl_certificate_bindings WHERE id = ?", id)
	return err
}

func (r *SSLCertificateBindingRepository) Count() (int64, error) {
	var count int64
	err := r.DB.QueryRow("SELECT COUNT(*) FROM ssl_certificate_bindings").Scan(&count)
	return count, err
}

// =========================================================
// TLS OPTION REPOSITORY
// =========================================================

type TLSOptionRepository struct {
	DB *sql.DB
}

func NewTLSOptionRepository(db *sql.DB) *TLSOptionRepository {
	return &TLSOptionRepository{DB: db}
}

func (r *TLSOptionRepository) Create(req *model.CreateTLSOptionRequest) (*model.TLSOption, error) {
	query := `INSERT INTO tls_options (binding_type, host_id, virtual_host_id, min_tls_version, http2_enabled, hsts_enabled, hsts_max_age) VALUES (?, ?, ?, ?, ?, ?, ?) RETURNING id, created_at`
	var item model.TLSOption
	err := r.DB.QueryRow(query, req.BindingType, req.HostID, req.VirtualHostID, req.MinTLSVersion, req.HTTP2Enabled, req.HSTSEnabled, req.HSTSMaxAge).Scan(&item.ID, &item.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("gagal membuat TLS option: %w", err)
	}
	item.BindingType = req.BindingType
	item.HostID = req.HostID
	item.VirtualHostID = req.VirtualHostID
	item.MinTLSVersion = req.MinTLSVersion
	item.HTTP2Enabled = req.HTTP2Enabled
	item.HSTSEnabled = req.HSTSEnabled
	item.HSTSMaxAge = req.HSTSMaxAge
	return &item, nil
}

func (r *TLSOptionRepository) GetByID(id int64) (*model.TLSOption, error) {
	query := `SELECT t.id, t.binding_type, t.host_id, t.virtual_host_id, t.min_tls_version, t.http2_enabled, t.hsts_enabled, t.hsts_max_age, t.created_at, COALESCE(h.host_name, ''), COALESCE(vh.vhost_name, '') FROM tls_options t LEFT JOIN hosts h ON t.host_id = h.id LEFT JOIN virtual_hosts vh ON t.virtual_host_id = vh.id WHERE t.id = ?`
	var item model.TLSOption
	err := r.DB.QueryRow(query, id).Scan(&item.ID, &item.BindingType, &item.HostID, &item.VirtualHostID, &item.MinTLSVersion, &item.HTTP2Enabled, &item.HSTSEnabled, &item.HSTSMaxAge, &item.CreatedAt, &item.HostName, &item.VHostName)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("TLS option tidak ditemukan")
		}
		return nil, err
	}
	return &item, nil
}

func (r *TLSOptionRepository) GetAll(page, limit int) ([]model.TLSOption, error) {
	offset := (page - 1) * limit
	query := `SELECT t.id, t.binding_type, t.host_id, t.virtual_host_id, t.min_tls_version, t.http2_enabled, t.hsts_enabled, t.hsts_max_age, t.created_at, COALESCE(h.host_name, ''), COALESCE(vh.vhost_name, '') FROM tls_options t LEFT JOIN hosts h ON t.host_id = h.id LEFT JOIN virtual_hosts vh ON t.virtual_host_id = vh.id ORDER BY t.id DESC LIMIT ? OFFSET ?`
	rows, err := r.DB.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []model.TLSOption
	for rows.Next() {
		var item model.TLSOption
		if err := rows.Scan(&item.ID, &item.BindingType, &item.HostID, &item.VirtualHostID, &item.MinTLSVersion, &item.HTTP2Enabled, &item.HSTSEnabled, &item.HSTSMaxAge, &item.CreatedAt, &item.HostName, &item.VHostName); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (r *TLSOptionRepository) Update(id int64, req *model.UpdateTLSOptionRequest) error {
	query := `UPDATE tls_options SET binding_type = ?, host_id = ?, virtual_host_id = ?, min_tls_version = ?, http2_enabled = ?, hsts_enabled = ?, hsts_max_age = ? WHERE id = ?`
	_, err := r.DB.Exec(query, req.BindingType, req.HostID, req.VirtualHostID, req.MinTLSVersion, req.HTTP2Enabled, req.HSTSEnabled, req.HSTSMaxAge, id)
	return err
}

func (r *TLSOptionRepository) Delete(id int64) error {
	_, err := r.DB.Exec("DELETE FROM tls_options WHERE id = ?", id)
	return err
}

func (r *TLSOptionRepository) Count() (int64, error) {
	var count int64
	err := r.DB.QueryRow("SELECT COUNT(*) FROM tls_options").Scan(&count)
	return count, err
}

// =========================================================
// SERVICE DISCOVERY REPOSITORY
// =========================================================

type ServiceDiscoveryRepository struct {
	DB *sql.DB
}

func NewServiceDiscoveryRepository(db *sql.DB) *ServiceDiscoveryRepository {
	return &ServiceDiscoveryRepository{DB: db}
}

func (r *ServiceDiscoveryRepository) Create(req *model.CreateServiceDiscoveryRequest) (*model.ServiceDiscovery, error) {
	query := `INSERT INTO service_discovery (virtual_host_id, provider, endpoint_url, refresh_interval_seconds, is_active) VALUES (?, ?, ?, ?, ?) RETURNING id, created_at`
	var item model.ServiceDiscovery
	err := r.DB.QueryRow(query, req.VirtualHostID, req.Provider, req.EndpointURL, req.RefreshIntervalSeconds, req.IsActive).Scan(&item.ID, &item.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("gagal membuat service discovery: %w", err)
	}
	item.VirtualHostID = req.VirtualHostID
	item.Provider = req.Provider
	item.EndpointURL = req.EndpointURL
	item.RefreshIntervalSeconds = req.RefreshIntervalSeconds
	item.IsActive = req.IsActive
	return &item, nil
}

func (r *ServiceDiscoveryRepository) GetByID(id int64) (*model.ServiceDiscovery, error) {
	query := `SELECT sd.id, sd.virtual_host_id, sd.provider, sd.endpoint_url, sd.refresh_interval_seconds, sd.is_active, sd.created_at, COALESCE(vh.vhost_name, '') FROM service_discovery sd LEFT JOIN virtual_hosts vh ON sd.virtual_host_id = vh.id WHERE sd.id = ?`
	var item model.ServiceDiscovery
	err := r.DB.QueryRow(query, id).Scan(&item.ID, &item.VirtualHostID, &item.Provider, &item.EndpointURL, &item.RefreshIntervalSeconds, &item.IsActive, &item.CreatedAt, &item.VHostName)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("service discovery tidak ditemukan")
		}
		return nil, err
	}
	return &item, nil
}

func (r *ServiceDiscoveryRepository) GetAll(page, limit int) ([]model.ServiceDiscovery, error) {
	offset := (page - 1) * limit
	query := `SELECT sd.id, sd.virtual_host_id, sd.provider, sd.endpoint_url, sd.refresh_interval_seconds, sd.is_active, sd.created_at, COALESCE(vh.vhost_name, '') FROM service_discovery sd LEFT JOIN virtual_hosts vh ON sd.virtual_host_id = vh.id ORDER BY sd.id DESC LIMIT ? OFFSET ?`
	rows, err := r.DB.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []model.ServiceDiscovery
	for rows.Next() {
		var item model.ServiceDiscovery
		if err := rows.Scan(&item.ID, &item.VirtualHostID, &item.Provider, &item.EndpointURL, &item.RefreshIntervalSeconds, &item.IsActive, &item.CreatedAt, &item.VHostName); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (r *ServiceDiscoveryRepository) Update(id int64, req *model.UpdateServiceDiscoveryRequest) error {
	query := `UPDATE service_discovery SET provider = ?, endpoint_url = ?, refresh_interval_seconds = ?, is_active = ? WHERE id = ?`
	_, err := r.DB.Exec(query, req.Provider, req.EndpointURL, req.RefreshIntervalSeconds, req.IsActive, id)
	return err
}

func (r *ServiceDiscoveryRepository) Delete(id int64) error {
	_, err := r.DB.Exec("DELETE FROM service_discovery WHERE id = ?", id)
	return err
}

func (r *ServiceDiscoveryRepository) Count() (int64, error) {
	var count int64
	err := r.DB.QueryRow("SELECT COUNT(*) FROM service_discovery").Scan(&count)
	return count, err
}

// =========================================================
// CONFIG VERSION REPOSITORY
// =========================================================

type ConfigVersionRepository struct {
	DB *sql.DB
}

func NewConfigVersionRepository(db *sql.DB) *ConfigVersionRepository {
	return &ConfigVersionRepository{DB: db}
}

func (r *ConfigVersionRepository) Create(req *model.CreateConfigVersionRequest) (*model.ConfigVersion, error) {
	query := `INSERT INTO config_versions (config_name, version_number, changed_by, notes) VALUES (?, ?, ?, ?) RETURNING id, created_at`
	var item model.ConfigVersion
	err := r.DB.QueryRow(query, req.ConfigName, req.VersionNumber, req.ChangedBy, req.Notes).Scan(&item.ID, &item.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("gagal membuat config version: %w", err)
	}
	item.ConfigName = req.ConfigName
	item.VersionNumber = req.VersionNumber
	item.ChangedBy = req.ChangedBy
	item.Notes = req.Notes
	return &item, nil
}

func (r *ConfigVersionRepository) GetByID(id int64) (*model.ConfigVersion, error) {
	query := `SELECT id, config_name, version_number, changed_by, notes, created_at FROM config_versions WHERE id = ?`
	var item model.ConfigVersion
	err := r.DB.QueryRow(query, id).Scan(&item.ID, &item.ConfigName, &item.VersionNumber, &item.ChangedBy, &item.Notes, &item.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("config version tidak ditemukan")
		}
		return nil, err
	}
	return &item, nil
}

func (r *ConfigVersionRepository) GetAll(page, limit int) ([]model.ConfigVersion, error) {
	offset := (page - 1) * limit
	query := `SELECT id, config_name, version_number, changed_by, notes, created_at FROM config_versions ORDER BY id DESC LIMIT ? OFFSET ?`
	rows, err := r.DB.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []model.ConfigVersion
	for rows.Next() {
		var item model.ConfigVersion
		if err := rows.Scan(&item.ID, &item.ConfigName, &item.VersionNumber, &item.ChangedBy, &item.Notes, &item.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (r *ConfigVersionRepository) Update(id int64, req *model.UpdateConfigVersionRequest) error {
	_, err := r.DB.Exec("UPDATE config_versions SET config_name = ?, version_number = ?, changed_by = ?, notes = ? WHERE id = ?", req.ConfigName, req.VersionNumber, req.ChangedBy, req.Notes, id)
	return err
}

func (r *ConfigVersionRepository) Delete(id int64) error {
	_, err := r.DB.Exec("DELETE FROM config_versions WHERE id = ?", id)
	return err
}

func (r *ConfigVersionRepository) Count() (int64, error) {
	var count int64
	err := r.DB.QueryRow("SELECT COUNT(*) FROM config_versions").Scan(&count)
	return count, err
}

// =========================================================
// MAINTENANCE WINDOW REPOSITORY
// =========================================================

type MaintenanceWindowRepository struct {
	DB *sql.DB
}

func NewMaintenanceWindowRepository(db *sql.DB) *MaintenanceWindowRepository {
	return &MaintenanceWindowRepository{DB: db}
}

func (r *MaintenanceWindowRepository) Create(req *model.CreateMaintenanceWindowRequest) (*model.MaintenanceWindow, error) {
	query := `INSERT INTO maintenance_windows (virtual_host_id, title, start_at, end_at, maintenance_response_code, maintenance_message, is_active) VALUES (?, ?, ?, ?, ?, ?, ?) RETURNING id, created_at`
	var item model.MaintenanceWindow
	err := r.DB.QueryRow(query, req.VirtualHostID, req.Title, req.StartAt, req.EndAt, req.MaintenanceResponseCode, req.MaintenanceMessage, req.IsActive).Scan(&item.ID, &item.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("gagal membuat maintenance window: %w", err)
	}
	item.VirtualHostID = req.VirtualHostID
	item.Title = req.Title
	item.StartAt = req.StartAt
	item.EndAt = req.EndAt
	item.MaintenanceResponseCode = req.MaintenanceResponseCode
	item.MaintenanceMessage = req.MaintenanceMessage
	item.IsActive = req.IsActive
	return &item, nil
}

func (r *MaintenanceWindowRepository) GetByID(id int64) (*model.MaintenanceWindow, error) {
	query := `SELECT mw.id, mw.virtual_host_id, mw.title, mw.start_at, mw.end_at, mw.maintenance_response_code, mw.maintenance_message, mw.is_active, mw.created_at, COALESCE(vh.vhost_name, '') FROM maintenance_windows mw LEFT JOIN virtual_hosts vh ON mw.virtual_host_id = vh.id WHERE mw.id = ?`
	var item model.MaintenanceWindow
	err := r.DB.QueryRow(query, id).Scan(&item.ID, &item.VirtualHostID, &item.Title, &item.StartAt, &item.EndAt, &item.MaintenanceResponseCode, &item.MaintenanceMessage, &item.IsActive, &item.CreatedAt, &item.VHostName)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("maintenance window tidak ditemukan")
		}
		return nil, err
	}
	return &item, nil
}

func (r *MaintenanceWindowRepository) GetAll(page, limit int) ([]model.MaintenanceWindow, error) {
	offset := (page - 1) * limit
	query := `SELECT mw.id, mw.virtual_host_id, mw.title, mw.start_at, mw.end_at, mw.maintenance_response_code, mw.maintenance_message, mw.is_active, mw.created_at, COALESCE(vh.vhost_name, '') FROM maintenance_windows mw LEFT JOIN virtual_hosts vh ON mw.virtual_host_id = vh.id ORDER BY mw.start_at DESC LIMIT ? OFFSET ?`
	rows, err := r.DB.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []model.MaintenanceWindow
	for rows.Next() {
		var item model.MaintenanceWindow
		if err := rows.Scan(&item.ID, &item.VirtualHostID, &item.Title, &item.StartAt, &item.EndAt, &item.MaintenanceResponseCode, &item.MaintenanceMessage, &item.IsActive, &item.CreatedAt, &item.VHostName); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (r *MaintenanceWindowRepository) Update(id int64, req *model.UpdateMaintenanceWindowRequest) error {
	query := `UPDATE maintenance_windows SET virtual_host_id = ?, title = ?, start_at = ?, end_at = ?, maintenance_response_code = ?, maintenance_message = ?, is_active = ? WHERE id = ?`
	_, err := r.DB.Exec(query, req.VirtualHostID, req.Title, req.StartAt, req.EndAt, req.MaintenanceResponseCode, req.MaintenanceMessage, req.IsActive, id)
	return err
}

func (r *MaintenanceWindowRepository) Delete(id int64) error {
	_, err := r.DB.Exec("DELETE FROM maintenance_windows WHERE id = ?", id)
	return err
}

func (r *MaintenanceWindowRepository) Count() (int64, error) {
	var count int64
	err := r.DB.QueryRow("SELECT COUNT(*) FROM maintenance_windows").Scan(&count)
	return count, err
}
