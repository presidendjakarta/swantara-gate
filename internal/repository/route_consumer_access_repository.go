package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/presidendjakarta/swantara-gate/internal/model"
)

// RouteConsumerAccessRepository menangani operasi database untuk Route Consumer Access (ACL)
type RouteConsumerAccessRepository struct {
	DB *sql.DB
}

// NewRouteConsumerAccessRepository membuat instance baru
func NewRouteConsumerAccessRepository(db *sql.DB) *RouteConsumerAccessRepository {
	return &RouteConsumerAccessRepository{DB: db}
}

// Create membuat akses consumer ke route baru
func (r *RouteConsumerAccessRepository) Create(req *model.CreateRouteConsumerAccessRequest) (*model.RouteConsumerAccess, error) {
	query := `
		INSERT INTO route_consumer_access (virtual_directory_id, consumer_id, is_active)
		VALUES (?, ?, ?)
		RETURNING id, created_at
	`

	var rca model.RouteConsumerAccess
	err := r.DB.QueryRow(query, req.VirtualDirectoryID, req.ConsumerID, req.IsActive).Scan(&rca.ID, &rca.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("gagal membuat route consumer access: %w", err)
	}

	rca.VirtualDirectoryID = req.VirtualDirectoryID
	rca.ConsumerID = req.ConsumerID
	rca.IsActive = req.IsActive

	return &rca, nil
}

// GetByID mengambil access berdasarkan ID
func (r *RouteConsumerAccessRepository) GetByID(id int64) (*model.RouteConsumerAccess, error) {
	query := `
		SELECT rca.id, rca.virtual_directory_id, rca.consumer_id, rca.is_active, rca.created_at,
		       ac.consumer_name, vd.source_path
		FROM route_consumer_access rca
		LEFT JOIN api_consumers ac ON rca.consumer_id = ac.id
		LEFT JOIN virtual_directories vd ON rca.virtual_directory_id = vd.id
		WHERE rca.id = ?
	`

	var rca model.RouteConsumerAccess
	err := r.DB.QueryRow(query, id).Scan(
		&rca.ID, &rca.VirtualDirectoryID, &rca.ConsumerID, &rca.IsActive, &rca.CreatedAt,
		&rca.ConsumerName, &rca.SourcePath,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("route consumer access tidak ditemukan")
		}
		return nil, fmt.Errorf("gagal mengambil route consumer access: %w", err)
	}

	return &rca, nil
}

// GetAll mengambil semua access dengan pagination
func (r *RouteConsumerAccessRepository) GetAll(page, limit int) ([]model.RouteConsumerAccess, error) {
	offset := (page - 1) * limit
	query := `
		SELECT rca.id, rca.virtual_directory_id, rca.consumer_id, rca.is_active, rca.created_at,
		       ac.consumer_name, vd.source_path
		FROM route_consumer_access rca
		LEFT JOIN api_consumers ac ON rca.consumer_id = ac.id
		LEFT JOIN virtual_directories vd ON rca.virtual_directory_id = vd.id
		ORDER BY rca.created_at DESC
		LIMIT ? OFFSET ?
	`

	rows, err := r.DB.Query(query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil daftar access: %w", err)
	}
	defer rows.Close()

	var accesses []model.RouteConsumerAccess
	for rows.Next() {
		var rca model.RouteConsumerAccess
		err := rows.Scan(
			&rca.ID, &rca.VirtualDirectoryID, &rca.ConsumerID, &rca.IsActive, &rca.CreatedAt,
			&rca.ConsumerName, &rca.SourcePath,
		)
		if err != nil {
			return nil, fmt.Errorf("gagal memindai access: %w", err)
		}
		accesses = append(accesses, rca)
	}

	return accesses, nil
}

// GetByDirectoryID mengambil access berdasarkan virtual directory ID
func (r *RouteConsumerAccessRepository) GetByDirectoryID(dirID int64) ([]model.RouteConsumerAccess, error) {
	query := `
		SELECT rca.id, rca.virtual_directory_id, rca.consumer_id, rca.is_active, rca.created_at,
		       ac.consumer_name, vd.source_path
		FROM route_consumer_access rca
		LEFT JOIN api_consumers ac ON rca.consumer_id = ac.id
		LEFT JOIN virtual_directories vd ON rca.virtual_directory_id = vd.id
		WHERE rca.virtual_directory_id = ?
		ORDER BY rca.created_at DESC
	`

	rows, err := r.DB.Query(query, dirID)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil access: %w", err)
	}
	defer rows.Close()

	var accesses []model.RouteConsumerAccess
	for rows.Next() {
		var rca model.RouteConsumerAccess
		err := rows.Scan(
			&rca.ID, &rca.VirtualDirectoryID, &rca.ConsumerID, &rca.IsActive, &rca.CreatedAt,
			&rca.ConsumerName, &rca.SourcePath,
		)
		if err != nil {
			return nil, fmt.Errorf("gagal memindai access: %w", err)
		}
		accesses = append(accesses, rca)
	}

	return accesses, nil
}

// GetByConsumerID mengambil access berdasarkan consumer ID
func (r *RouteConsumerAccessRepository) GetByConsumerID(consumerID int64) ([]model.RouteConsumerAccess, error) {
	query := `
		SELECT rca.id, rca.virtual_directory_id, rca.consumer_id, rca.is_active, rca.created_at,
		       ac.consumer_name, vd.source_path
		FROM route_consumer_access rca
		LEFT JOIN api_consumers ac ON rca.consumer_id = ac.id
		LEFT JOIN virtual_directories vd ON rca.virtual_directory_id = vd.id
		WHERE rca.consumer_id = ?
		ORDER BY rca.created_at DESC
	`

	rows, err := r.DB.Query(query, consumerID)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil access: %w", err)
	}
	defer rows.Close()

	var accesses []model.RouteConsumerAccess
	for rows.Next() {
		var rca model.RouteConsumerAccess
		err := rows.Scan(
			&rca.ID, &rca.VirtualDirectoryID, &rca.ConsumerID, &rca.IsActive, &rca.CreatedAt,
			&rca.ConsumerName, &rca.SourcePath,
		)
		if err != nil {
			return nil, fmt.Errorf("gagal memindai access: %w", err)
		}
		accesses = append(accesses, rca)
	}

	return accesses, nil
}

// Update memperbarui status akses
func (r *RouteConsumerAccessRepository) Update(id int64, req *model.UpdateRouteConsumerAccessRequest) error {
	query := `UPDATE route_consumer_access SET is_active = ?, created_at = ? WHERE id = ?`

	result, err := r.DB.Exec(query, req.IsActive, time.Now(), id)
	if err != nil {
		return fmt.Errorf("gagal mengupdate access: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("route consumer access tidak ditemukan")
	}

	return nil
}

// Delete menghapus access
func (r *RouteConsumerAccessRepository) Delete(id int64) error {
	result, err := r.DB.Exec("DELETE FROM route_consumer_access WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("gagal menghapus access: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("route consumer access tidak ditemukan")
	}

	return nil
}

// Count menghitung total access
func (r *RouteConsumerAccessRepository) Count() (int64, error) {
	var count int64
	err := r.DB.QueryRow("SELECT COUNT(*) FROM route_consumer_access").Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("gagal menghitung access: %w", err)
	}
	return count, nil
}
