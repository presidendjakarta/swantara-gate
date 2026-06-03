package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/presidendjakarta/swantara-gate/internal/model"
)

// HostRepository menangani operasi database untuk Host
type HostRepository struct {
	DB *sql.DB
}

// NewHostRepository membuat instance baru HostRepository
func NewHostRepository(db *sql.DB) *HostRepository {
	return &HostRepository{DB: db}
}

// Create membuat host baru
func (r *HostRepository) Create(host *model.CreateHostRequest) (*model.Host, error) {
	query := `
		INSERT INTO hosts (host_name, description, is_active)
		VALUES (?, ?, ?)
		RETURNING id, created_at, updated_at
	`

	var createdHost model.Host
	err := r.DB.QueryRow(
		query,
		host.HostName,
		host.Description,
		host.IsActive,
	).Scan(&createdHost.ID, &createdHost.CreatedAt, &createdHost.UpdatedAt)

	if err != nil {
		return nil, fmt.Errorf("gagal membuat host: %w", err)
	}

	createdHost.HostName = host.HostName
	createdHost.Description = host.Description
	createdHost.IsActive = host.IsActive

	return &createdHost, nil
}

// GetByID mengambil host berdasarkan ID
func (r *HostRepository) GetByID(id int64) (*model.Host, error) {
	query := `
		SELECT id, host_name, description, is_active, created_at, updated_at
		FROM hosts
		WHERE id = ?
	`

	var host model.Host
	err := r.DB.QueryRow(query, id).Scan(
		&host.ID,
		&host.HostName,
		&host.Description,
		&host.IsActive,
		&host.CreatedAt,
		&host.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("host tidak ditemukan")
		}
		return nil, fmt.Errorf("gagal mengambil host: %w", err)
	}

	return &host, nil
}

// GetAll mengambil semua host dengan pagination dan search
func (r *HostRepository) GetAll(page, limit int, search string) ([]model.Host, error) {
	offset := (page - 1) * limit

	query := `
		SELECT id, host_name, description, is_active, created_at, updated_at
		FROM hosts
	`
	
	var args []interface{}
	
	// Add search filter if provided
	if search != "" {
		query += ` WHERE host_name LIKE ? OR description LIKE ?`
		searchPattern := "%" + search + "%"
		args = append(args, searchPattern, searchPattern)
	}
	
	query += ` ORDER BY created_at DESC LIMIT ? OFFSET ?`
	args = append(args, limit, offset)

	rows, err := r.DB.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil daftar host: %w", err)
	}
	defer rows.Close()

	var hosts []model.Host
	for rows.Next() {
		var host model.Host
		err := rows.Scan(
			&host.ID,
			&host.HostName,
			&host.Description,
			&host.IsActive,
			&host.CreatedAt,
			&host.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("gagal memindai host: %w", err)
		}
		hosts = append(hosts, host)
	}

	return hosts, nil
}

// Update memperbarui data host
func (r *HostRepository) Update(id int64, host *model.UpdateHostRequest) error {
	query := `
		UPDATE hosts
		SET host_name = ?, description = ?, is_active = ?, updated_at = ?
		WHERE id = ?
	`

	result, err := r.DB.Exec(query, host.HostName, host.Description, host.IsActive, time.Now(), id)
	if err != nil {
		return fmt.Errorf("gagal mengupdate host: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("gagal mendapatkan rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("host tidak ditemukan")
	}

	return nil
}

// Delete menghapus host
func (r *HostRepository) Delete(id int64) error {
	query := `DELETE FROM hosts WHERE id = ?`

	result, err := r.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("gagal menghapus host: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("gagal mendapatkan rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("host tidak ditemukan")
	}

	return nil
}

// Count menghitung total host dengan search filter
func (r *HostRepository) Count(search string) (int64, error) {
	query := `SELECT COUNT(*) FROM hosts`
	
	var args []interface{}
	
	// Add search filter if provided
	if search != "" {
		query += ` WHERE host_name LIKE ? OR description LIKE ?`
		searchPattern := "%" + search + "%"
		args = append(args, searchPattern, searchPattern)
	}

	var count int64
	err := r.DB.QueryRow(query, args...).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("gagal menghitung host: %w", err)
	}

	return count, nil
}
