package service

import (
	"fmt"

	"github.com/presidendjakarta/swantara-gate/internal/model"
	"github.com/presidendjakarta/swantara-gate/internal/repository"
)

// VirtualDirectoryService menangani business logic untuk Virtual Directory
type VirtualDirectoryService struct {
	DirRepo    *repository.VirtualDirectoryRepository
	MethodRepo *repository.VirtualDirectoryMethodRepository
}

// NewVirtualDirectoryService membuat instance baru VirtualDirectoryService
func NewVirtualDirectoryService(dirRepo *repository.VirtualDirectoryRepository, methodRepo *repository.VirtualDirectoryMethodRepository) *VirtualDirectoryService {
	return &VirtualDirectoryService{DirRepo: dirRepo, MethodRepo: methodRepo}
}

// CreateVirtualDirectory membuat virtual directory baru dengan validasi
func (s *VirtualDirectoryService) CreateVirtualDirectory(req *model.CreateVirtualDirectoryRequest) (*model.VirtualDirectory, error) {
	if req.SourcePath == "" {
		return nil, fmt.Errorf("source_path wajib diisi")
	}
	if req.TargetPath == "" {
		return nil, fmt.Errorf("target_path wajib diisi")
	}

	// Default values
	if req.MatchType == "" {
		req.MatchType = "prefix"
	}
	if req.AuthType == "" {
		req.AuthType = "none"
	}
	if req.ProxyTimeoutSeconds <= 0 {
		req.ProxyTimeoutSeconds = 30
	}
	if req.MaxRequestSizeMB <= 0 {
		req.MaxRequestSizeMB = 10
	}

	dir, err := s.DirRepo.Create(req)
	if err != nil {
		return nil, fmt.Errorf("gagal membuat virtual directory: %w", err)
	}

	return dir, nil
}

// GetVirtualDirectoryByID mengambil virtual directory berdasarkan ID
func (s *VirtualDirectoryService) GetVirtualDirectoryByID(id int64) (*model.VirtualDirectory, error) {
	return s.DirRepo.GetByID(id)
}

// GetAllVirtualDirectories mengambil semua virtual directory dengan pagination
func (s *VirtualDirectoryService) GetAllVirtualDirectories(page, limit int) ([]model.VirtualDirectory, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	total, err := s.DirRepo.Count()
	if err != nil {
		return nil, 0, err
	}

	dirs, err := s.DirRepo.GetAll(page, limit)
	if err != nil {
		return nil, 0, err
	}

	return dirs, total, nil
}

// GetVirtualDirectoriesByVHostID mengambil directories berdasarkan virtual host
func (s *VirtualDirectoryService) GetVirtualDirectoriesByVHostID(vhostID int64) ([]model.VirtualDirectory, error) {
	return s.DirRepo.GetByVirtualHostID(vhostID)
}

// UpdateVirtualDirectory memperbarui virtual directory
func (s *VirtualDirectoryService) UpdateVirtualDirectory(id int64, req *model.UpdateVirtualDirectoryRequest) error {
	_, err := s.DirRepo.GetByID(id)
	if err != nil {
		return fmt.Errorf("virtual directory tidak ditemukan")
	}

	return s.DirRepo.Update(id, req)
}

// DeleteVirtualDirectory menghapus virtual directory beserta methods-nya
func (s *VirtualDirectoryService) DeleteVirtualDirectory(id int64) error {
	_, err := s.DirRepo.GetByID(id)
	if err != nil {
		return fmt.Errorf("virtual directory tidak ditemukan")
	}

	// Hapus methods terkait terlebih dahulu
	_ = s.MethodRepo.DeleteByDirectoryID(id)

	return s.DirRepo.Delete(id)
}

// === Virtual Directory Methods Service ===

// GetMethods mengambil methods untuk directory tertentu
func (s *VirtualDirectoryService) GetMethods(dirID int64) ([]model.VirtualDirectoryMethod, error) {
	return s.MethodRepo.GetByDirectoryID(dirID)
}

// SetMethods mengatur methods untuk directory tertentu
func (s *VirtualDirectoryService) SetMethods(dirID int64, methods []string) ([]model.VirtualDirectoryMethod, error) {
	// Validasi directory ada
	_, err := s.DirRepo.GetByID(dirID)
	if err != nil {
		return nil, fmt.Errorf("virtual directory tidak ditemukan")
	}

	// Validasi HTTP methods
	validMethods := map[string]bool{
		"GET": true, "POST": true, "PUT": true, "PATCH": true,
		"DELETE": true, "HEAD": true, "OPTIONS": true,
	}
	for _, m := range methods {
		if !validMethods[m] {
			return nil, fmt.Errorf("method tidak valid: %s (gunakan GET/POST/PUT/PATCH/DELETE/HEAD/OPTIONS)", m)
		}
	}

	return s.MethodRepo.SetMethods(dirID, methods)
}
