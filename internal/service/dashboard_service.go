package service

import (
	"github.com/presidendjakarta/swantara-gate/internal/repository"
)

// DashboardService menangani bisnis logika untuk statistik dashboard
type DashboardService struct {
	Repo *repository.DashboardRepository
}

// NewDashboardService membuat instance baru DashboardService
func NewDashboardService(repo *repository.DashboardRepository) *DashboardService {
	return &DashboardService{Repo: repo}
}

// GetDashboardStats mengambil semua statistik dashboard
func (s *DashboardService) GetDashboardStats() (*repository.DashboardStats, error) {
	return s.Repo.GetDashboardStats()
}
