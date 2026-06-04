package service

import (
	"fmt"
	"github.com/presidendjakarta/swantara-gate/internal/model"
	"github.com/presidendjakarta/swantara-gate/internal/repository"
)

// ExternalAuthProviderService menangani business logic untuk external auth providers
type ExternalAuthProviderService struct {
	ProviderRepo *repository.ExternalAuthProviderRepository
	VDirRepo     *repository.VirtualDirectoryRepository
}

// NewExternalAuthProviderService membuat instance baru
func NewExternalAuthProviderService(providerRepo *repository.ExternalAuthProviderRepository, vdirRepo *repository.VirtualDirectoryRepository) *ExternalAuthProviderService {
	return &ExternalAuthProviderService{
		ProviderRepo: providerRepo,
		VDirRepo:     vdirRepo,
	}
}

// Create membuat provider baru
func (s *ExternalAuthProviderService) Create(req *model.CreateExternalAuthProviderRequest) (*model.ExternalAuthProvider, error) {
	// Validasi required fields
	if req.Name == "" {
		return nil, fmt.Errorf("nama provider wajib diisi")
	}
	if req.AuthURL == "" {
		return nil, fmt.Errorf("auth URL wajib diisi")
	}

	// Set defaults
	if req.HTTPMethod == "" {
		req.HTTPMethod = "POST"
	}
	if req.RequestTimeoutSeconds == 0 {
		req.RequestTimeoutSeconds = 5
	}
	if req.SuccessKey == "" {
		req.SuccessKey = "status"
	}
	if req.SuccessValue == "" {
		req.SuccessValue = "true"
	}
	if req.MessageKey == "" {
		req.MessageKey = "message"
	}

	return s.ProviderRepo.Create(req)
}

// GetByID mengambil provider berdasarkan ID
func (s *ExternalAuthProviderService) GetByID(id int64) (*model.ExternalAuthProvider, error) {
	return s.ProviderRepo.GetByID(id)
}

// GetAll mengambil semua provider dengan pagination
func (s *ExternalAuthProviderService) GetAll(page, limit int, search string) ([]model.ExternalAuthProvider, int64, error) {
	total, err := s.ProviderRepo.Count(search)
	if err != nil {
		return nil, 0, fmt.Errorf("gagal menghitung total providers: %w", err)
	}

	providers, err := s.ProviderRepo.GetAll(page, limit, search)
	if err != nil {
		return nil, 0, fmt.Errorf("gagal mengambil daftar providers: %w", err)
	}

	return providers, total, nil
}

// Update mengupdate provider
func (s *ExternalAuthProviderService) Update(id int64, req *model.UpdateExternalAuthProviderRequest) error {
	// Cek apakah provider exists
	_, err := s.ProviderRepo.GetByID(id)
	if err != nil {
		return err
	}

	return s.ProviderRepo.Update(id, req)
}

// Delete menghapus provider
func (s *ExternalAuthProviderService) Delete(id int64) error {
	// Cek apakah provider masih digunakan
	provider, err := s.ProviderRepo.GetByID(id)
	if err != nil {
		return err
	}

	if provider.UsedByCount > 0 {
		return fmt.Errorf("provider masih digunakan oleh %d virtual directory, hapus mapping terlebih dahulu", provider.UsedByCount)
	}

	return s.ProviderRepo.Delete(id)
}

// AssignProviderToVirtualDirectory meng-assign provider ke virtual directory
func (s *ExternalAuthProviderService) AssignProviderToVirtualDirectory(vdirID, providerID int64) error {
	// Validasi virtual directory exists
	_, err := s.VDirRepo.GetByID(vdirID)
	if err != nil {
		return fmt.Errorf("virtual directory tidak ditemukan: %w", err)
	}

	// Validasi provider exists
	_, err = s.ProviderRepo.GetByID(providerID)
	if err != nil {
		return fmt.Errorf("provider tidak ditemukan: %w", err)
	}

	return s.ProviderRepo.AssignProviderToVirtualDirectory(vdirID, providerID)
}

// RemoveProviderFromVirtualDirectory menghapus assignment
func (s *ExternalAuthProviderService) RemoveProviderFromVirtualDirectory(vdirID, providerID int64) error {
	return s.ProviderRepo.RemoveProviderFromVirtualDirectory(vdirID, providerID)
}

// GetProvidersByVirtualDirectory mengambil semua provider untuk virtual directory
func (s *ExternalAuthProviderService) GetProvidersByVirtualDirectory(vdirID int64) ([]model.ExternalAuthProvider, error) {
	return s.ProviderRepo.GetProvidersByVirtualDirectoryID(vdirID)
}

// GetMappingsByVirtualDirectory mengambil semua mapping untuk virtual directory
func (s *ExternalAuthProviderService) GetMappingsByVirtualDirectory(vdirID int64) ([]model.VirtualDirectoryExternalAuthMapping, error) {
	return s.ProviderRepo.GetMappingsByVirtualDirectoryID(vdirID)
}
