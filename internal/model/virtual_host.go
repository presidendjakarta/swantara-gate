package model

import "time"

// VirtualHost merepresentasikan virtual domain dengan load balancing
type VirtualHost struct {
	ID            int64     `json:"id"`
	HostID        int64     `json:"host_id"`
	VHostName     string    `json:"vhost_name"`
	LBAlgorithm   string    `json:"lb_algorithm"`
	StickySession bool      `json:"sticky_session"`
	FailoverMode  string    `json:"failover_mode"`
	IsActive      bool      `json:"is_active"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	
	// Host data (untuk join)
	HostName string `json:"host_name,omitempty"`
}

// CreateVirtualHostRequest request untuk membuat virtual host baru
type CreateVirtualHostRequest struct {
	HostID        int64  `json:"host_id" validate:"required"`
	VHostName     string `json:"vhost_name" validate:"required"`
	LBAlgorithm   string `json:"lb_algorithm"`
	StickySession bool   `json:"sticky_session"`
	FailoverMode  string `json:"failover_mode"`
	IsActive      bool   `json:"is_active"`
}

// UpdateVirtualHostRequest request untuk update virtual host
type UpdateVirtualHostRequest struct {
	LBAlgorithm   string `json:"lb_algorithm"`
	StickySession bool   `json:"sticky_session"`
	FailoverMode  string `json:"failover_mode"`
	IsActive      bool   `json:"is_active"`
}
