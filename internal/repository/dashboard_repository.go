package repository

import (
	"database/sql"
)

// DashboardRepository menangani operasi database untuk statistik dashboard
type DashboardRepository struct {
	DB *sql.DB
}

// NewDashboardRepository membuat instance baru DashboardRepository
func NewDashboardRepository(db *sql.DB) *DashboardRepository {
	return &DashboardRepository{DB: db}
}

// DashboardStats menyimpan semua statistik dashboard
type DashboardStats struct {
	HostCount           int64 `json:"host_count"`
	VirtualHostCount    int64 `json:"virtual_host_count"`
	UpstreamCount       int64 `json:"upstream_count"`
	RouteCount          int64 `json:"route_count"`
	ConsumerCount       int64 `json:"consumer_count"`
	RateLimitCount      int64 `json:"rate_limit_count"`
	CORSCount           int64 `json:"cors_count"`
	CircuitBreakerCount int64 `json:"circuit_breaker_count"`
	SSLBindingCount     int64 `json:"ssl_binding_count"`
	ActiveUserCount     int64 `json:"active_user_count"`
}

// GetDashboardStats mengambil semua statistik dashboard dalam satu query
func (r *DashboardRepository) GetDashboardStats() (*DashboardStats, error) {
	stats := &DashboardStats{}

	queries := map[string]*int64{
		"SELECT COUNT(*) FROM hosts":                    &stats.HostCount,
		"SELECT COUNT(*) FROM virtual_hosts":            &stats.VirtualHostCount,
		"SELECT COUNT(*) FROM upstream_servers":         &stats.UpstreamCount,
		"SELECT COUNT(*) FROM virtual_directories":      &stats.RouteCount,
		"SELECT COUNT(*) FROM api_consumers":            &stats.ConsumerCount,
		"SELECT COUNT(*) FROM rate_limits":              &stats.RateLimitCount,
		"SELECT COUNT(*) FROM cors_configs":             &stats.CORSCount,
		"SELECT COUNT(*) FROM circuit_breakers":         &stats.CircuitBreakerCount,
		"SELECT COUNT(*) FROM ssl_certificate_bindings": &stats.SSLBindingCount,
		"SELECT COUNT(*) FROM users WHERE is_active = 1": &stats.ActiveUserCount,
	}

	for query, dest := range queries {
		err := r.DB.QueryRow(query).Scan(dest)
		if err != nil {
			// Jika tabel belum ada, set 0
			*dest = 0
		}
	}

	return stats, nil
}
