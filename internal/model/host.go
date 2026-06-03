package model

import "time"

// Host merepresentasikan host utama
type Host struct {
	ID          int64     `json:"id"`
	HostName    string    `json:"host_name"`
	Description string    `json:"description"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// CreateHostRequest request untuk membuat host baru
type CreateHostRequest struct {
	HostName    string `json:"host_name" validate:"required"`
	Description string `json:"description"`
	IsActive    bool   `json:"is_active"`
}

// UpdateHostRequest request untuk update host
type UpdateHostRequest struct {
	HostName    string `json:"host_name"`
	Description string `json:"description"`
	IsActive    bool   `json:"is_active"`
}
