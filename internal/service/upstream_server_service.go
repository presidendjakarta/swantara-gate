package service

import (
	"fmt"

	"github.com/presidendjakarta/swantara-gate/internal/model"
	"github.com/presidendjakarta/swantara-gate/internal/repository"
)

// UpstreamServerService menangani business logic untuk Upstream Server
type UpstreamServerService struct {
	UpstreamRepo *repository.UpstreamServerRepository
}

// NewUpstreamServerService membuat instance baru UpstreamServerService
func NewUpstreamServerService(repo *repository.UpstreamServerRepository) *UpstreamServerService {
	return &UpstreamServerService{UpstreamRepo: repo}
}

// CreateUpstreamServer membuat upstream server baru dengan validasi
func (s *UpstreamServerService) CreateUpstreamServer(req *model.CreateUpstreamServerRequest) (*model.UpstreamServer, error) {
	if req.TargetHost == "" {
		return nil, fmt.Errorf("target_host wajib diisi")
	}
	if req.TargetPort <= 0 || req.TargetPort > 65535 {
		return nil, fmt.Errorf("target_port harus antara 1-65535")
	}

	// Default values
	if req.Protocol == "" {
		req.Protocol = "http"
	}
	if req.Weight <= 0 {
		req.Weight = 1
	}
	if req.HealthCheckIntervalSeconds <= 0 {
		req.HealthCheckIntervalSeconds = 30
	}
	if req.HealthCheckTimeoutSeconds <= 0 {
		req.HealthCheckTimeoutSeconds = 5
	}
	if req.MaxFails <= 0 {
		req.MaxFails = 3
	}
	if req.FailTimeoutSeconds <= 0 {
		req.FailTimeoutSeconds = 30
	}

	server, err := s.UpstreamRepo.Create(req)
	if err != nil {
		return nil, fmt.Errorf("gagal membuat upstream server: %w", err)
	}

	return server, nil
}

// GetUpstreamServerByID mengambil upstream server berdasarkan ID
func (s *UpstreamServerService) GetUpstreamServerByID(id int64) (*model.UpstreamServer, error) {
	server, err := s.UpstreamRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return server, nil
}

// GetAllUpstreamServers mengambil semua upstream server dengan pagination
func (s *UpstreamServerService) GetAllUpstreamServers(page, limit int, search string) ([]model.UpstreamServer, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	total, err := s.UpstreamRepo.Count(search)
	if err != nil {
		return nil, 0, err
	}

	servers, err := s.UpstreamRepo.GetAll(page, limit, search)
	if err != nil {
		return nil, 0, err
	}

	return servers, total, nil
}

// GetUpstreamServersByVHostID mengambil upstream servers berdasarkan virtual host
func (s *UpstreamServerService) GetUpstreamServersByVHostID(vhostID int64) ([]model.UpstreamServer, error) {
	servers, err := s.UpstreamRepo.GetByVirtualHostID(vhostID)
	if err != nil {
		return nil, err
	}
	return servers, nil
}

// UpdateUpstreamServer memperbarui upstream server
func (s *UpstreamServerService) UpdateUpstreamServer(id int64, req *model.UpdateUpstreamServerRequest) error {
	_, err := s.UpstreamRepo.GetByID(id)
	if err != nil {
		return fmt.Errorf("upstream server tidak ditemukan")
	}

	return s.UpstreamRepo.Update(id, req)
}

// DeleteUpstreamServer menghapus upstream server
func (s *UpstreamServerService) DeleteUpstreamServer(id int64) error {
	_, err := s.UpstreamRepo.GetByID(id)
	if err != nil {
		return fmt.Errorf("upstream server tidak ditemukan")
	}

	return s.UpstreamRepo.Delete(id)
}
