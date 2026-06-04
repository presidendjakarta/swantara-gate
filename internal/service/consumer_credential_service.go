package service

import (
	"fmt"

	"github.com/presidendjakarta/swantara-gate/internal/model"
	"github.com/presidendjakarta/swantara-gate/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

// ConsumerCredentialService menangani business logic untuk Consumer Credentials
type ConsumerCredentialService struct {
	CredRepo *repository.ConsumerCredentialRepository
}

// NewConsumerCredentialService membuat instance baru
func NewConsumerCredentialService(credRepo *repository.ConsumerCredentialRepository) *ConsumerCredentialService {
	return &ConsumerCredentialService{CredRepo: credRepo}
}

// CreateCredential membuat credential baru dengan validasi
func (s *ConsumerCredentialService) CreateCredential(req *model.CreateConsumerCredentialRequest) (*model.ConsumerCredential, error) {
	if req.AuthType == "" {
		return nil, fmt.Errorf("auth_type wajib diisi (basic/api_key/jwt)")
	}

	// Validasi berdasarkan auth type
	validTypes := map[string]bool{"basic": true, "api_key": true, "jwt": true}
	if !validTypes[req.AuthType] {
		return nil, fmt.Errorf("auth_type tidak valid: gunakan basic, api_key, atau jwt")
	}

	var passwordHash string
	if req.AuthType == "basic" {
		if req.Username == "" || req.Password == "" {
			return nil, fmt.Errorf("username dan password wajib untuk auth_type basic")
		}
		hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, fmt.Errorf("gagal hash password: %w", err)
		}
		passwordHash = string(hash)
	}

	cred, err := s.CredRepo.Create(req, passwordHash)
	if err != nil {
		return nil, err
	}

	return cred, nil
}

// GetCredentialByID mengambil credential berdasarkan ID
func (s *ConsumerCredentialService) GetCredentialByID(id int64) (*model.ConsumerCredential, error) {
	return s.CredRepo.GetByID(id)
}

// GetAllCredentials mengambil semua credentials dengan pagination
func (s *ConsumerCredentialService) GetAllCredentials(page, limit int, search string) ([]model.ConsumerCredential, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	total, err := s.CredRepo.Count(search)
	if err != nil {
		return nil, 0, err
	}

	creds, err := s.CredRepo.GetAll(page, limit, search)
	if err != nil {
		return nil, 0, err
	}

	return creds, total, nil
}

// GetCredentialsByConsumerID mengambil credentials berdasarkan consumer
func (s *ConsumerCredentialService) GetCredentialsByConsumerID(consumerID int64) ([]model.ConsumerCredential, error) {
	return s.CredRepo.GetByConsumerID(consumerID)
}

// UpdateCredential memperbarui credential
func (s *ConsumerCredentialService) UpdateCredential(id int64, req *model.UpdateConsumerCredentialRequest) error {
	_, err := s.CredRepo.GetByID(id)
	if err != nil {
		return fmt.Errorf("credential tidak ditemukan")
	}

	var passwordHash string
	if req.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("gagal hash password: %w", err)
		}
		passwordHash = string(hash)
	}

	return s.CredRepo.Update(id, req, passwordHash)
}

// DeleteCredential menghapus credential
func (s *ConsumerCredentialService) DeleteCredential(id int64) error {
	_, err := s.CredRepo.GetByID(id)
	if err != nil {
		return fmt.Errorf("credential tidak ditemukan")
	}

	return s.CredRepo.Delete(id)
}

// === API Key Service ===

// APIKeyService menangani business logic untuk API Keys
type APIKeyService struct {
	KeyRepo *repository.APIKeyRepository
}

// NewAPIKeyService membuat instance baru APIKeyService
func NewAPIKeyService(keyRepo *repository.APIKeyRepository) *APIKeyService {
	return &APIKeyService{KeyRepo: keyRepo}
}

// CreateAPIKey membuat API key baru (key di-generate otomatis)
func (s *APIKeyService) CreateAPIKey(req *model.CreateAPIKeyRequest) (*model.APIKey, error) {
	if req.ConsumerID <= 0 {
		return nil, fmt.Errorf("consumer_id wajib diisi")
	}

	key, err := s.KeyRepo.Create(req)
	if err != nil {
		return nil, err
	}

	return key, nil
}

// GetAPIKeyByID mengambil API key berdasarkan ID
func (s *APIKeyService) GetAPIKeyByID(id int64) (*model.APIKey, error) {
	return s.KeyRepo.GetByID(id)
}

// GetAllAPIKeys mengambil semua API keys dengan pagination
func (s *APIKeyService) GetAllAPIKeys(page, limit int, search string) ([]model.APIKey, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	total, err := s.KeyRepo.Count(search)
	if err != nil {
		return nil, 0, err
	}

	keys, err := s.KeyRepo.GetAll(page, limit, search)
	if err != nil {
		return nil, 0, err
	}

	return keys, total, nil
}

// GetAPIKeysByConsumerID mengambil API keys berdasarkan consumer
func (s *APIKeyService) GetAPIKeysByConsumerID(consumerID int64) ([]model.APIKey, error) {
	return s.KeyRepo.GetByConsumerID(consumerID)
}

// UpdateAPIKey memperbarui API key
func (s *APIKeyService) UpdateAPIKey(id int64, req *model.UpdateAPIKeyRequest) error {
	_, err := s.KeyRepo.GetByID(id)
	if err != nil {
		return fmt.Errorf("API key tidak ditemukan")
	}

	return s.KeyRepo.Update(id, req)
}

// DeleteAPIKey menghapus API key
func (s *APIKeyService) DeleteAPIKey(id int64) error {
	_, err := s.KeyRepo.GetByID(id)
	if err != nil {
		return fmt.Errorf("API key tidak ditemukan")
	}

	return s.KeyRepo.Delete(id)
}

// === Route Consumer Access Service ===

// RouteConsumerAccessService menangani business logic untuk ACL
type RouteConsumerAccessService struct {
	AccessRepo *repository.RouteConsumerAccessRepository
}

// NewRouteConsumerAccessService membuat instance baru
func NewRouteConsumerAccessService(accessRepo *repository.RouteConsumerAccessRepository) *RouteConsumerAccessService {
	return &RouteConsumerAccessService{AccessRepo: accessRepo}
}

// CreateAccess membuat akses baru
func (s *RouteConsumerAccessService) CreateAccess(req *model.CreateRouteConsumerAccessRequest) (*model.RouteConsumerAccess, error) {
	if req.VirtualDirectoryID <= 0 {
		return nil, fmt.Errorf("virtual_directory_id wajib diisi")
	}
	if req.ConsumerID <= 0 {
		return nil, fmt.Errorf("consumer_id wajib diisi")
	}

	access, err := s.AccessRepo.Create(req)
	if err != nil {
		return nil, err
	}

	return access, nil
}

// GetAccessByID mengambil access berdasarkan ID
func (s *RouteConsumerAccessService) GetAccessByID(id int64) (*model.RouteConsumerAccess, error) {
	return s.AccessRepo.GetByID(id)
}

// GetAllAccess mengambil semua access dengan pagination
func (s *RouteConsumerAccessService) GetAllAccess(page, limit int) ([]model.RouteConsumerAccess, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	total, err := s.AccessRepo.Count()
	if err != nil {
		return nil, 0, err
	}

	accesses, err := s.AccessRepo.GetAll(page, limit)
	if err != nil {
		return nil, 0, err
	}

	return accesses, total, nil
}

// GetAccessByDirectoryID mengambil access berdasarkan directory
func (s *RouteConsumerAccessService) GetAccessByDirectoryID(dirID int64) ([]model.RouteConsumerAccess, error) {
	return s.AccessRepo.GetByDirectoryID(dirID)
}

// GetAccessByConsumerID mengambil access berdasarkan consumer
func (s *RouteConsumerAccessService) GetAccessByConsumerID(consumerID int64) ([]model.RouteConsumerAccess, error) {
	return s.AccessRepo.GetByConsumerID(consumerID)
}

// UpdateAccess memperbarui status akses
func (s *RouteConsumerAccessService) UpdateAccess(id int64, req *model.UpdateRouteConsumerAccessRequest) error {
	_, err := s.AccessRepo.GetByID(id)
	if err != nil {
		return fmt.Errorf("access tidak ditemukan")
	}

	return s.AccessRepo.Update(id, req)
}

// DeleteAccess menghapus akses
func (s *RouteConsumerAccessService) DeleteAccess(id int64) error {
	_, err := s.AccessRepo.GetByID(id)
	if err != nil {
		return fmt.Errorf("access tidak ditemukan")
	}

	return s.AccessRepo.Delete(id)
}
