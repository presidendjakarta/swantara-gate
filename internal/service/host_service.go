package service

import (
	"fmt"

	"github.com/presidendjakarta/swantara-gate/internal/model"
	"github.com/presidendjakarta/swantara-gate/internal/repository"
)

// HostService menangani business logic untuk Host
type HostService struct {
	HostRepo *repository.HostRepository
}

// NewHostService membuat instance baru HostService
func NewHostService(hostRepo *repository.HostRepository) *HostService {
	return &HostService{HostRepo: hostRepo}
}

// CreateHost membuat host baru
func (s *HostService) CreateHost(req *model.CreateHostRequest) (*model.Host, error) {
	host, err := s.HostRepo.Create(req)
	if err != nil {
		return nil, fmt.Errorf("gagal membuat host: %w", err)
	}

	return host, nil
}

// GetHostByID mengambil host berdasarkan ID
func (s *HostService) GetHostByID(id int64) (*model.Host, error) {
	host, err := s.HostRepo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil host: %w", err)
	}

	return host, nil
}

// GetAllHosts mengambil semua host dengan pagination dan search
func (s *HostService) GetAllHosts(page, limit int, search string) ([]model.Host, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	total, err := s.HostRepo.Count(search)
	if err != nil {
		return nil, 0, fmt.Errorf("gagal menghitung host: %w", err)
	}

	hosts, err := s.HostRepo.GetAll(page, limit, search)
	if err != nil {
		return nil, 0, fmt.Errorf("gagal mengambil daftar host: %w", err)
	}

	return hosts, total, nil
}

// UpdateHost memperbarui data host
func (s *HostService) UpdateHost(id int64, req *model.UpdateHostRequest) error {
	_, err := s.HostRepo.GetByID(id)
	if err != nil {
		return fmt.Errorf("host tidak ditemukan: %w", err)
	}

	err = s.HostRepo.Update(id, req)
	if err != nil {
		return fmt.Errorf("gagal mengupdate host: %w", err)
	}

	return nil
}

// DeleteHost menghapus host
func (s *HostService) DeleteHost(id int64) error {
	_, err := s.HostRepo.GetByID(id)
	if err != nil {
		return fmt.Errorf("host tidak ditemukan: %w", err)
	}

	err = s.HostRepo.Delete(id)
	if err != nil {
		return fmt.Errorf("gagal menghapus host: %w", err)
	}

	return nil
}

// VirtualHostService menangani business logic untuk Virtual Host
type VirtualHostService struct {
	VHostRepo *repository.VirtualHostRepository
}

// NewVirtualHostService membuat instance baru VirtualHostService
func NewVirtualHostService(vhostRepo *repository.VirtualHostRepository) *VirtualHostService {
	return &VirtualHostService{VHostRepo: vhostRepo}
}

// CreateVirtualHost membuat virtual host baru
func (s *VirtualHostService) CreateVirtualHost(req *model.CreateVirtualHostRequest) (*model.VirtualHost, error) {
	// Set default values
	if req.LBAlgorithm == "" {
		req.LBAlgorithm = "round_robin"
	}
	if req.FailoverMode == "" {
		req.FailoverMode = "active-active"
	}

	vhost, err := s.VHostRepo.Create(req)
	if err != nil {
		return nil, fmt.Errorf("gagal membuat virtual host: %w", err)
	}

	return vhost, nil
}

// GetVirtualHostByID mengambil virtual host berdasarkan ID
func (s *VirtualHostService) GetVirtualHostByID(id int64) (*model.VirtualHost, error) {
	vhost, err := s.VHostRepo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil virtual host: %w", err)
	}

	return vhost, nil
}

// GetAllVirtualHosts mengambil semua virtual host dengan pagination dan search
func (s *VirtualHostService) GetAllVirtualHosts(page, limit int, search string) ([]model.VirtualHost, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	total, err := s.VHostRepo.Count(search)
	if err != nil {
		return nil, 0, fmt.Errorf("gagal menghitung virtual host: %w", err)
	}

	vhosts, err := s.VHostRepo.GetAll(page, limit, search)
	if err != nil {
		return nil, 0, fmt.Errorf("gagal mengambil daftar virtual host: %w", err)
	}

	return vhosts, total, nil
}

// GetVirtualHostsByHostID mengambil virtual host berdasarkan host_id
func (s *VirtualHostService) GetVirtualHostsByHostID(hostID int64) ([]model.VirtualHost, error) {
	vhosts, err := s.VHostRepo.GetByHostID(hostID)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil virtual host: %w", err)
	}

	return vhosts, nil
}

// UpdateVirtualHost memperbarui data virtual host
func (s *VirtualHostService) UpdateVirtualHost(id int64, req *model.UpdateVirtualHostRequest) error {
	_, err := s.VHostRepo.GetByID(id)
	if err != nil {
		return fmt.Errorf("virtual host tidak ditemukan: %w", err)
	}

	err = s.VHostRepo.Update(id, req)
	if err != nil {
		return fmt.Errorf("gagal mengupdate virtual host: %w", err)
	}

	return nil
}

// DeleteVirtualHost menghapus virtual host
func (s *VirtualHostService) DeleteVirtualHost(id int64) error {
	_, err := s.VHostRepo.GetByID(id)
	if err != nil {
		return fmt.Errorf("virtual host tidak ditemukan: %w", err)
	}

	err = s.VHostRepo.Delete(id)
	if err != nil {
		return fmt.Errorf("gagal menghapus virtual host: %w", err)
	}

	return nil
}
