package service

import (
	"fmt"
	"github.com/presidendjakarta/swantara-gate/internal/model"
	"github.com/presidendjakarta/swantara-gate/internal/repository"
)

// JWTProviderService menangani business logic untuk JWT providers
type JWTProviderService struct {
	ProviderRepo *repository.JWTProviderRepository
	VDirRepo     *repository.VirtualDirectoryRepository
}

// NewJWTProviderService membuat instance baru
func NewJWTProviderService(providerRepo *repository.JWTProviderRepository, vdirRepo *repository.VirtualDirectoryRepository) *JWTProviderService {
	return &JWTProviderService{
		ProviderRepo: providerRepo,
		VDirRepo:     vdirRepo,
	}
}

// Create membuat provider baru
func (s *JWTProviderService) Create(req *model.CreateJWTProviderRequest) (*model.JWTProvider, error) {
	// Validasi required fields
	if req.Name == "" {
		return nil, fmt.Errorf("nama provider wajib diisi")
	}
	if req.JWTSecret == "" {
		return nil, fmt.Errorf("JWT secret wajib diisi")
	}
	if len(req.JWTSecret) < 32 {
		return nil, fmt.Errorf("JWT secret minimal 32 karakter untuk keamanan")
	}

	// Set defaults
	if req.Algorithm == "" {
		req.Algorithm = "HS256"
	}
	if req.ExpiredInSeconds == 0 {
		req.ExpiredInSeconds = 3600
	}

	return s.ProviderRepo.Create(req)
}

// GetByID mengambil provider berdasarkan ID
func (s *JWTProviderService) GetByID(id int64) (*model.JWTProvider, error) {
	return s.ProviderRepo.GetByID(id)
}

// GetAll mengambil semua provider dengan pagination
func (s *JWTProviderService) GetAll(page, limit int, search string) ([]model.JWTProvider, int64, error) {
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
func (s *JWTProviderService) Update(id int64, req *model.UpdateJWTProviderRequest) error {
	// Cek apakah provider exists
	_, err := s.ProviderRepo.GetByID(id)
	if err != nil {
		return err
	}

	// Validasi secret jika ada
	if req.JWTSecret != "" && len(req.JWTSecret) < 32 {
		return fmt.Errorf("JWT secret minimal 32 karakter untuk keamanan")
	}

	return s.ProviderRepo.Update(id, req)
}

// Delete menghapus provider
func (s *JWTProviderService) Delete(id int64) error {
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
func (s *JWTProviderService) AssignProviderToVirtualDirectory(vdirID, providerID int64) error {
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
func (s *JWTProviderService) RemoveProviderFromVirtualDirectory(vdirID, providerID int64) error {
	return s.ProviderRepo.RemoveProviderFromVirtualDirectory(vdirID, providerID)
}

// GetProvidersByVirtualDirectory mengambil semua provider untuk virtual directory
func (s *JWTProviderService) GetProvidersByVirtualDirectory(vdirID int64) ([]model.JWTProvider, error) {
	return s.ProviderRepo.GetProvidersByVirtualDirectoryID(vdirID)
}

// GetMappingsByVirtualDirectory mengambil semua mapping untuk virtual directory
func (s *JWTProviderService) GetMappingsByVirtualDirectory(vdirID int64) ([]model.VirtualDirectoryJWTProviderMapping, error) {
	return s.ProviderRepo.GetMappingsByVirtualDirectoryID(vdirID)
}
