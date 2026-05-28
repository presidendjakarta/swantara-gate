package service

import (
	"fmt"

	"github.com/presidendjakarta/swantara-gate/internal/model"
	"github.com/presidendjakarta/swantara-gate/internal/repository"
)

// =========================================================
// ACME ACCOUNT SERVICE
// =========================================================

type ACMEAccountService struct{ Repo *repository.ACMEAccountRepository }

func NewACMEAccountService(repo *repository.ACMEAccountRepository) *ACMEAccountService {
	return &ACMEAccountService{Repo: repo}
}

func (s *ACMEAccountService) Create(req *model.CreateACMEAccountRequest) (*model.ACMEAccount, error) {
	if req.Email == "" {
		return nil, fmt.Errorf("email wajib diisi")
	}
	if req.ProviderURL == "" {
		return nil, fmt.Errorf("provider_url wajib diisi")
	}
	if req.AccountKeyPath == "" {
		return nil, fmt.Errorf("account_key_path wajib diisi")
	}
	return s.Repo.Create(req)
}

func (s *ACMEAccountService) GetByID(id int64) (*model.ACMEAccount, error) { return s.Repo.GetByID(id) }

func (s *ACMEAccountService) GetAll(page, limit int) ([]model.ACMEAccount, int64, error) {
	if page < 1 { page = 1 }
	if limit < 1 || limit > 100 { limit = 10 }
	total, _ := s.Repo.Count()
	items, err := s.Repo.GetAll(page, limit)
	return items, total, err
}

func (s *ACMEAccountService) Update(id int64, req *model.UpdateACMEAccountRequest) error {
	if _, err := s.Repo.GetByID(id); err != nil { return fmt.Errorf("ACME account tidak ditemukan") }
	return s.Repo.Update(id, req)
}

func (s *ACMEAccountService) Delete(id int64) error {
	if _, err := s.Repo.GetByID(id); err != nil { return fmt.Errorf("ACME account tidak ditemukan") }
	return s.Repo.Delete(id)
}

// =========================================================
// SSL CERTIFICATE SERVICE
// =========================================================

type SSLCertificateService struct{ Repo *repository.SSLCertificateRepository }

func NewSSLCertificateService(repo *repository.SSLCertificateRepository) *SSLCertificateService {
	return &SSLCertificateService{Repo: repo}
}

func (s *SSLCertificateService) Create(req *model.CreateSSLCertificateRequest) (*model.SSLCertificate, error) {
	if req.CertificatePath == "" {
		return nil, fmt.Errorf("certificate_path wajib diisi")
	}
	if req.PrivateKeyPath == "" {
		return nil, fmt.Errorf("private_key_path wajib diisi")
	}
	if req.Provider == "" { req.Provider = "lets_encrypt" }
	if req.ChallengeType == "" { req.ChallengeType = "http01" }
	if req.RenewBeforeDays <= 0 { req.RenewBeforeDays = 30 }
	return s.Repo.Create(req)
}

func (s *SSLCertificateService) GetByID(id int64) (*model.SSLCertificate, error) { return s.Repo.GetByID(id) }

func (s *SSLCertificateService) GetAll(page, limit int) ([]model.SSLCertificate, int64, error) {
	if page < 1 { page = 1 }
	if limit < 1 || limit > 100 { limit = 10 }
	total, _ := s.Repo.Count()
	items, err := s.Repo.GetAll(page, limit)
	return items, total, err
}

func (s *SSLCertificateService) Update(id int64, req *model.UpdateSSLCertificateRequest) error {
	if _, err := s.Repo.GetByID(id); err != nil { return fmt.Errorf("SSL certificate tidak ditemukan") }
	return s.Repo.Update(id, req)
}

func (s *SSLCertificateService) Delete(id int64) error {
	if _, err := s.Repo.GetByID(id); err != nil { return fmt.Errorf("SSL certificate tidak ditemukan") }
	return s.Repo.Delete(id)
}

// =========================================================
// CERTIFICATE DOMAIN SERVICE
// =========================================================

type CertificateDomainService struct{ Repo *repository.CertificateDomainRepository }

func NewCertificateDomainService(repo *repository.CertificateDomainRepository) *CertificateDomainService {
	return &CertificateDomainService{Repo: repo}
}

func (s *CertificateDomainService) Create(req *model.CreateCertificateDomainRequest) (*model.CertificateDomain, error) {
	if req.SSLCertificateID <= 0 {
		return nil, fmt.Errorf("ssl_certificate_id wajib diisi")
	}
	if req.DomainName == "" {
		return nil, fmt.Errorf("domain_name wajib diisi")
	}
	return s.Repo.Create(req)
}

func (s *CertificateDomainService) GetByID(id int64) (*model.CertificateDomain, error) { return s.Repo.GetByID(id) }

func (s *CertificateDomainService) GetAll(page, limit int) ([]model.CertificateDomain, int64, error) {
	if page < 1 { page = 1 }
	if limit < 1 || limit > 100 { limit = 10 }
	total, _ := s.Repo.Count()
	items, err := s.Repo.GetAll(page, limit)
	return items, total, err
}

func (s *CertificateDomainService) GetByCertificateID(certID int64) ([]model.CertificateDomain, error) {
	return s.Repo.GetByCertificateID(certID)
}

func (s *CertificateDomainService) Update(id int64, req *model.UpdateCertificateDomainRequest) error {
	if _, err := s.Repo.GetByID(id); err != nil { return fmt.Errorf("certificate domain tidak ditemukan") }
	return s.Repo.Update(id, req)
}

func (s *CertificateDomainService) Delete(id int64) error {
	if _, err := s.Repo.GetByID(id); err != nil { return fmt.Errorf("certificate domain tidak ditemukan") }
	return s.Repo.Delete(id)
}

// =========================================================
// SSL CERTIFICATE BINDING SERVICE
// =========================================================

var validBindingTypes = map[string]bool{"host": true, "virtual_host": true, "global": true}

type SSLCertificateBindingService struct{ Repo *repository.SSLCertificateBindingRepository }

func NewSSLCertificateBindingService(repo *repository.SSLCertificateBindingRepository) *SSLCertificateBindingService {
	return &SSLCertificateBindingService{Repo: repo}
}

func (s *SSLCertificateBindingService) Create(req *model.CreateSSLCertificateBindingRequest) (*model.SSLCertificateBinding, error) {
	if req.SSLCertificateID <= 0 {
		return nil, fmt.Errorf("ssl_certificate_id wajib diisi")
	}
	if req.BindingType == "" || !validBindingTypes[req.BindingType] {
		return nil, fmt.Errorf("binding_type wajib diisi (host, virtual_host, global)")
	}
	if req.Priority <= 0 { req.Priority = 1 }
	return s.Repo.Create(req)
}

func (s *SSLCertificateBindingService) GetByID(id int64) (*model.SSLCertificateBinding, error) { return s.Repo.GetByID(id) }

func (s *SSLCertificateBindingService) GetAll(page, limit int) ([]model.SSLCertificateBinding, int64, error) {
	if page < 1 { page = 1 }
	if limit < 1 || limit > 100 { limit = 10 }
	total, _ := s.Repo.Count()
	items, err := s.Repo.GetAll(page, limit)
	return items, total, err
}

func (s *SSLCertificateBindingService) Update(id int64, req *model.UpdateSSLCertificateBindingRequest) error {
	if _, err := s.Repo.GetByID(id); err != nil { return fmt.Errorf("SSL binding tidak ditemukan") }
	return s.Repo.Update(id, req)
}

func (s *SSLCertificateBindingService) Delete(id int64) error {
	if _, err := s.Repo.GetByID(id); err != nil { return fmt.Errorf("SSL binding tidak ditemukan") }
	return s.Repo.Delete(id)
}

// =========================================================
// TLS OPTION SERVICE
// =========================================================

type TLSOptionService struct{ Repo *repository.TLSOptionRepository }

func NewTLSOptionService(repo *repository.TLSOptionRepository) *TLSOptionService {
	return &TLSOptionService{Repo: repo}
}

func (s *TLSOptionService) Create(req *model.CreateTLSOptionRequest) (*model.TLSOption, error) {
	if req.BindingType == "" || !validBindingTypes[req.BindingType] {
		return nil, fmt.Errorf("binding_type wajib diisi (host, virtual_host, global)")
	}
	if req.MinTLSVersion == "" { req.MinTLSVersion = "1.2" }
	if req.HSTSMaxAge <= 0 { req.HSTSMaxAge = 31536000 }
	return s.Repo.Create(req)
}

func (s *TLSOptionService) GetByID(id int64) (*model.TLSOption, error) { return s.Repo.GetByID(id) }

func (s *TLSOptionService) GetAll(page, limit int) ([]model.TLSOption, int64, error) {
	if page < 1 { page = 1 }
	if limit < 1 || limit > 100 { limit = 10 }
	total, _ := s.Repo.Count()
	items, err := s.Repo.GetAll(page, limit)
	return items, total, err
}

func (s *TLSOptionService) Update(id int64, req *model.UpdateTLSOptionRequest) error {
	if _, err := s.Repo.GetByID(id); err != nil { return fmt.Errorf("TLS option tidak ditemukan") }
	return s.Repo.Update(id, req)
}

func (s *TLSOptionService) Delete(id int64) error {
	if _, err := s.Repo.GetByID(id); err != nil { return fmt.Errorf("TLS option tidak ditemukan") }
	return s.Repo.Delete(id)
}

// =========================================================
// SERVICE DISCOVERY SERVICE
// =========================================================

type ServiceDiscoveryService struct{ Repo *repository.ServiceDiscoveryRepository }

func NewServiceDiscoveryService(repo *repository.ServiceDiscoveryRepository) *ServiceDiscoveryService {
	return &ServiceDiscoveryService{Repo: repo}
}

func (s *ServiceDiscoveryService) Create(req *model.CreateServiceDiscoveryRequest) (*model.ServiceDiscovery, error) {
	if req.VirtualHostID <= 0 {
		return nil, fmt.Errorf("virtual_host_id wajib diisi")
	}
	if req.Provider == "" {
		return nil, fmt.Errorf("provider wajib diisi")
	}
	if req.EndpointURL == "" {
		return nil, fmt.Errorf("endpoint_url wajib diisi")
	}
	if req.RefreshIntervalSeconds <= 0 { req.RefreshIntervalSeconds = 30 }
	return s.Repo.Create(req)
}

func (s *ServiceDiscoveryService) GetByID(id int64) (*model.ServiceDiscovery, error) { return s.Repo.GetByID(id) }

func (s *ServiceDiscoveryService) GetAll(page, limit int) ([]model.ServiceDiscovery, int64, error) {
	if page < 1 { page = 1 }
	if limit < 1 || limit > 100 { limit = 10 }
	total, _ := s.Repo.Count()
	items, err := s.Repo.GetAll(page, limit)
	return items, total, err
}

func (s *ServiceDiscoveryService) Update(id int64, req *model.UpdateServiceDiscoveryRequest) error {
	if _, err := s.Repo.GetByID(id); err != nil { return fmt.Errorf("service discovery tidak ditemukan") }
	return s.Repo.Update(id, req)
}

func (s *ServiceDiscoveryService) Delete(id int64) error {
	if _, err := s.Repo.GetByID(id); err != nil { return fmt.Errorf("service discovery tidak ditemukan") }
	return s.Repo.Delete(id)
}

// =========================================================
// CONFIG VERSION SERVICE
// =========================================================

type ConfigVersionService struct{ Repo *repository.ConfigVersionRepository }

func NewConfigVersionService(repo *repository.ConfigVersionRepository) *ConfigVersionService {
	return &ConfigVersionService{Repo: repo}
}

func (s *ConfigVersionService) Create(req *model.CreateConfigVersionRequest) (*model.ConfigVersion, error) {
	if req.ConfigName == "" {
		return nil, fmt.Errorf("config_name wajib diisi")
	}
	if req.VersionNumber <= 0 {
		return nil, fmt.Errorf("version_number wajib diisi")
	}
	return s.Repo.Create(req)
}

func (s *ConfigVersionService) GetByID(id int64) (*model.ConfigVersion, error) { return s.Repo.GetByID(id) }

func (s *ConfigVersionService) GetAll(page, limit int) ([]model.ConfigVersion, int64, error) {
	if page < 1 { page = 1 }
	if limit < 1 || limit > 100 { limit = 10 }
	total, _ := s.Repo.Count()
	items, err := s.Repo.GetAll(page, limit)
	return items, total, err
}

func (s *ConfigVersionService) Update(id int64, req *model.UpdateConfigVersionRequest) error {
	if _, err := s.Repo.GetByID(id); err != nil { return fmt.Errorf("config version tidak ditemukan") }
	return s.Repo.Update(id, req)
}

func (s *ConfigVersionService) Delete(id int64) error {
	if _, err := s.Repo.GetByID(id); err != nil { return fmt.Errorf("config version tidak ditemukan") }
	return s.Repo.Delete(id)
}

// =========================================================
// MAINTENANCE WINDOW SERVICE
// =========================================================

type MaintenanceWindowService struct{ Repo *repository.MaintenanceWindowRepository }

func NewMaintenanceWindowService(repo *repository.MaintenanceWindowRepository) *MaintenanceWindowService {
	return &MaintenanceWindowService{Repo: repo}
}

func (s *MaintenanceWindowService) Create(req *model.CreateMaintenanceWindowRequest) (*model.MaintenanceWindow, error) {
	if req.Title == "" {
		return nil, fmt.Errorf("title wajib diisi")
	}
	if req.StartAt == "" {
		return nil, fmt.Errorf("start_at wajib diisi")
	}
	if req.EndAt == "" {
		return nil, fmt.Errorf("end_at wajib diisi")
	}
	if req.MaintenanceResponseCode <= 0 { req.MaintenanceResponseCode = 503 }
	return s.Repo.Create(req)
}

func (s *MaintenanceWindowService) GetByID(id int64) (*model.MaintenanceWindow, error) { return s.Repo.GetByID(id) }

func (s *MaintenanceWindowService) GetAll(page, limit int) ([]model.MaintenanceWindow, int64, error) {
	if page < 1 { page = 1 }
	if limit < 1 || limit > 100 { limit = 10 }
	total, _ := s.Repo.Count()
	items, err := s.Repo.GetAll(page, limit)
	return items, total, err
}

func (s *MaintenanceWindowService) Update(id int64, req *model.UpdateMaintenanceWindowRequest) error {
	if _, err := s.Repo.GetByID(id); err != nil { return fmt.Errorf("maintenance window tidak ditemukan") }
	return s.Repo.Update(id, req)
}

func (s *MaintenanceWindowService) Delete(id int64) error {
	if _, err := s.Repo.GetByID(id); err != nil { return fmt.Errorf("maintenance window tidak ditemukan") }
	return s.Repo.Delete(id)
}
