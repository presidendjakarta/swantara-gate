package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/presidendjakarta/swantara-gate/internal/model"
)

// APIConsumerRepository menangani operasi database untuk API Consumer
type APIConsumerRepository struct {
	DB *sql.DB
}

// NewAPIConsumerRepository membuat instance baru APIConsumerRepository
func NewAPIConsumerRepository(db *sql.DB) *APIConsumerRepository {
	return &APIConsumerRepository{DB: db}
}

// Create membuat consumer baru
func (r *APIConsumerRepository) Create(consumer *model.CreateAPIConsumerRequest) (*model.APIConsumer, error) {
	query := `
		INSERT INTO api_consumers (consumer_name, description, contact_email, is_active)
		VALUES (?, ?, ?, ?)
		RETURNING id, created_at, updated_at
	`

	var createdConsumer model.APIConsumer
	err := r.DB.QueryRow(
		query,
		consumer.ConsumerName,
		consumer.Description,
		consumer.ContactEmail,
		consumer.IsActive,
	).Scan(&createdConsumer.ID, &createdConsumer.CreatedAt, &createdConsumer.UpdatedAt)

	if err != nil {
		return nil, fmt.Errorf("gagal membuat consumer: %w", err)
	}

	createdConsumer.ConsumerName = consumer.ConsumerName
	createdConsumer.Description = consumer.Description
	createdConsumer.ContactEmail = consumer.ContactEmail
	createdConsumer.IsActive = consumer.IsActive

	return &createdConsumer, nil
}

// GetByID mengambil consumer berdasarkan ID
func (r *APIConsumerRepository) GetByID(id int64) (*model.APIConsumer, error) {
	query := `
		SELECT id, consumer_name, description, contact_email, is_active, created_at, updated_at
		FROM api_consumers
		WHERE id = ?
	`

	var consumer model.APIConsumer
	err := r.DB.QueryRow(query, id).Scan(
		&consumer.ID,
		&consumer.ConsumerName,
		&consumer.Description,
		&consumer.ContactEmail,
		&consumer.IsActive,
		&consumer.CreatedAt,
		&consumer.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("consumer tidak ditemukan")
		}
		return nil, fmt.Errorf("gagal mengambil consumer: %w", err)
	}

	return &consumer, nil
}

// GetAll mengambil semua consumer dengan pagination
func (r *APIConsumerRepository) GetAll(page, limit int) ([]model.APIConsumer, error) {
	offset := (page - 1) * limit

	query := `
		SELECT id, consumer_name, description, contact_email, is_active, created_at, updated_at
		FROM api_consumers
		ORDER BY created_at DESC
		LIMIT ? OFFSET ?
	`

	rows, err := r.DB.Query(query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil daftar consumer: %w", err)
	}
	defer rows.Close()

	var consumers []model.APIConsumer
	for rows.Next() {
		var consumer model.APIConsumer
		err := rows.Scan(
			&consumer.ID,
			&consumer.ConsumerName,
			&consumer.Description,
			&consumer.ContactEmail,
			&consumer.IsActive,
			&consumer.CreatedAt,
			&consumer.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("gagal memindai consumer: %w", err)
		}
		consumers = append(consumers, consumer)
	}

	return consumers, nil
}

// Update memperbarui data consumer
func (r *APIConsumerRepository) Update(id int64, consumer *model.UpdateAPIConsumerRequest) error {
	query := `
		UPDATE api_consumers
		SET description = ?, contact_email = ?, is_active = ?, updated_at = ?
		WHERE id = ?
	`

	result, err := r.DB.Exec(query, consumer.Description, consumer.ContactEmail, consumer.IsActive, time.Now(), id)
	if err != nil {
		return fmt.Errorf("gagal mengupdate consumer: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("gagal mendapatkan rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("consumer tidak ditemukan")
	}

	return nil
}

// Delete menghapus consumer
func (r *APIConsumerRepository) Delete(id int64) error {
	query := `DELETE FROM api_consumers WHERE id = ?`

	result, err := r.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("gagal menghapus consumer: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("gagal mendapatkan rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("consumer tidak ditemukan")
	}

	return nil
}

// Count menghitung total consumer
func (r *APIConsumerRepository) Count() (int64, error) {
	query := `SELECT COUNT(*) FROM api_consumers`

	var count int64
	err := r.DB.QueryRow(query).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("gagal menghitung consumer: %w", err)
	}

	return count, nil
}
