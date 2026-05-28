package model

import "time"

// UpstreamServer merepresentasikan server backend yang menerima traffic dari proxy
type UpstreamServer struct {
	ID                        int64     `json:"id"`
	VirtualHostID             int64     `json:"virtual_host_id"`
	TargetHost                string    `json:"target_host"`
	TargetPort                int       `json:"target_port"`
	Protocol                  string    `json:"protocol"`
	Priority                  int       `json:"priority"`
	Weight                    int       `json:"weight"`
	IsBackup                  bool      `json:"is_backup"`
	IsActive                  bool      `json:"is_active"`
	HealthCheckEnabled        bool      `json:"health_check_enabled"`
	HealthCheckPath           string    `json:"health_check_path"`
	HealthCheckIntervalSeconds int      `json:"health_check_interval_seconds"`
	HealthCheckTimeoutSeconds int       `json:"health_check_timeout_seconds"`
	MaxFails                  int       `json:"max_fails"`
	FailTimeoutSeconds        int       `json:"fail_timeout_seconds"`
	CreatedAt                 time.Time `json:"created_at"`
	UpdatedAt                 time.Time `json:"updated_at"`

	// Data join
	VHostName string `json:"vhost_name,omitempty"`
}

// CreateUpstreamServerRequest request untuk membuat upstream server baru
type CreateUpstreamServerRequest struct {
	VirtualHostID             int64  `json:"virtual_host_id" validate:"required"`
	TargetHost                string `json:"target_host" validate:"required"`
	TargetPort                int    `json:"target_port" validate:"required"`
	Protocol                  string `json:"protocol"`
	Priority                  int    `json:"priority"`
	Weight                    int    `json:"weight"`
	IsBackup                  bool   `json:"is_backup"`
	IsActive                  bool   `json:"is_active"`
	HealthCheckEnabled        bool   `json:"health_check_enabled"`
	HealthCheckPath           string `json:"health_check_path"`
	HealthCheckIntervalSeconds int   `json:"health_check_interval_seconds"`
	HealthCheckTimeoutSeconds int    `json:"health_check_timeout_seconds"`
	MaxFails                  int    `json:"max_fails"`
	FailTimeoutSeconds        int    `json:"fail_timeout_seconds"`
}

// UpdateUpstreamServerRequest request untuk update upstream server
type UpdateUpstreamServerRequest struct {
	TargetHost                string `json:"target_host"`
	TargetPort                int    `json:"target_port"`
	Protocol                  string `json:"protocol"`
	Priority                  int    `json:"priority"`
	Weight                    int    `json:"weight"`
	IsBackup                  bool   `json:"is_backup"`
	IsActive                  bool   `json:"is_active"`
	HealthCheckEnabled        bool   `json:"health_check_enabled"`
	HealthCheckPath           string `json:"health_check_path"`
	HealthCheckIntervalSeconds int   `json:"health_check_interval_seconds"`
	HealthCheckTimeoutSeconds int    `json:"health_check_timeout_seconds"`
	MaxFails                  int    `json:"max_fails"`
	FailTimeoutSeconds        int    `json:"fail_timeout_seconds"`
}
