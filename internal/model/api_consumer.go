package model

import "time"

// APIConsumer merepresentasikan konsumen/aplikasi yang terdaftar
type APIConsumer struct {
	ID           int64     `json:"id"`
	ConsumerName string    `json:"consumer_name"`
	Description  string    `json:"description"`
	ContactEmail string    `json:"contact_email"`
	IsActive     bool      `json:"is_active"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// CreateAPIConsumerRequest request untuk membuat consumer baru
type CreateAPIConsumerRequest struct {
	ConsumerName string `json:"consumer_name" validate:"required"`
	Description  string `json:"description"`
	ContactEmail string `json:"contact_email"`
	IsActive     bool   `json:"is_active"`
}

// UpdateAPIConsumerRequest request untuk update consumer
type UpdateAPIConsumerRequest struct {
	Description  string `json:"description"`
	ContactEmail string `json:"contact_email"`
	IsActive     bool   `json:"is_active"`
}
