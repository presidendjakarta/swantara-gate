package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/presidendjakarta/swantara-gate/internal/model"
)

// VirtualHostRepository menangani operasi database untuk Virtual Host
type VirtualHostRepository struct {
	DB *sql.DB
}

// NewVirtualHostRepository membuat instance baru VirtualHostRepository
func NewVirtualHostRepository(db *sql.DB) *VirtualHostRepository {
	return &VirtualHostRepository{DB: db}
}

// Create membuat virtual host baru
func (r *VirtualHostRepository) Create(vhost *model.CreateVirtualHostRequest) (*model.VirtualHost, error) {
	query := `
		INSERT INTO virtual_hosts (host_id, vhost_name, lb_algorithm, sticky_session, failover_mode, is_active)
		VALUES (?, ?, ?, ?, ?, ?)
		RETURNING id, created_at, updated_at
	`

	var createdVHost model.VirtualHost
	err := r.DB.QueryRow(
		query,
		vhost.HostID,
		vhost.VHostName,
		vhost.LBAlgorithm,
		vhost.StickySession,
		vhost.FailoverMode,
		vhost.IsActive,
	).Scan(&createdVHost.ID, &createdVHost.CreatedAt, &createdVHost.UpdatedAt)

	if err != nil {
		return nil, fmt.Errorf("gagal membuat virtual host: %w", err)
	}

	createdVHost.HostID = vhost.HostID
	createdVHost.VHostName = vhost.VHostName
	createdVHost.LBAlgorithm = vhost.LBAlgorithm
	createdVHost.StickySession = vhost.StickySession
	createdVHost.FailoverMode = vhost.FailoverMode
	createdVHost.IsActive = vhost.IsActive

	return &createdVHost, nil
}

// GetByID mengambil virtual host berdasarkan ID
func (r *VirtualHostRepository) GetByID(id int64) (*model.VirtualHost, error) {
	query := `
		SELECT vh.id, vh.host_id, vh.vhost_name, vh.lb_algorithm, vh.sticky_session, 
		       vh.failover_mode, vh.is_active, vh.created_at, vh.updated_at,
		       h.host_name
		FROM virtual_hosts vh
		LEFT JOIN hosts h ON vh.host_id = h.id
		WHERE vh.id = ?
	`

	var vhost model.VirtualHost
	err := r.DB.QueryRow(query, id).Scan(
		&vhost.ID,
		&vhost.HostID,
		&vhost.VHostName,
		&vhost.LBAlgorithm,
		&vhost.StickySession,
		&vhost.FailoverMode,
		&vhost.IsActive,
		&vhost.CreatedAt,
		&vhost.UpdatedAt,
		&vhost.HostName,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("virtual host tidak ditemukan")
		}
		return nil, fmt.Errorf("gagal mengambil virtual host: %w", err)
	}

	return &vhost, nil
}

// GetAll mengambil semua virtual host dengan pagination
func (r *VirtualHostRepository) GetAll(page, limit int) ([]model.VirtualHost, error) {
	offset := (page - 1) * limit

	query := `
		SELECT vh.id, vh.host_id, vh.vhost_name, vh.lb_algorithm, vh.sticky_session, 
		       vh.failover_mode, vh.is_active, vh.created_at, vh.updated_at,
		       h.host_name
		FROM virtual_hosts vh
		LEFT JOIN hosts h ON vh.host_id = h.id
		ORDER BY vh.created_at DESC
		LIMIT ? OFFSET ?
	`

	rows, err := r.DB.Query(query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil daftar virtual host: %w", err)
	}
	defer rows.Close()

	var vhosts []model.VirtualHost
	for rows.Next() {
		var vhost model.VirtualHost
		err := rows.Scan(
			&vhost.ID,
			&vhost.HostID,
			&vhost.VHostName,
			&vhost.LBAlgorithm,
			&vhost.StickySession,
			&vhost.FailoverMode,
			&vhost.IsActive,
			&vhost.CreatedAt,
			&vhost.UpdatedAt,
			&vhost.HostName,
		)
		if err != nil {
			return nil, fmt.Errorf("gagal memindai virtual host: %w", err)
		}
		vhosts = append(vhosts, vhost)
	}

	return vhosts, nil
}

// GetByHostID mengambil semua virtual host berdasarkan host_id
func (r *VirtualHostRepository) GetByHostID(hostID int64) ([]model.VirtualHost, error) {
	query := `
		SELECT vh.id, vh.host_id, vh.vhost_name, vh.lb_algorithm, vh.sticky_session, 
		       vh.failover_mode, vh.is_active, vh.created_at, vh.updated_at,
		       h.host_name
		FROM virtual_hosts vh
		LEFT JOIN hosts h ON vh.host_id = h.id
		WHERE vh.host_id = ?
		ORDER BY vh.created_at DESC
	`

	rows, err := r.DB.Query(query, hostID)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil virtual host by host_id: %w", err)
	}
	defer rows.Close()

	var vhosts []model.VirtualHost
	for rows.Next() {
		var vhost model.VirtualHost
		err := rows.Scan(
			&vhost.ID,
			&vhost.HostID,
			&vhost.VHostName,
			&vhost.LBAlgorithm,
			&vhost.StickySession,
			&vhost.FailoverMode,
			&vhost.IsActive,
			&vhost.CreatedAt,
			&vhost.UpdatedAt,
			&vhost.HostName,
		)
		if err != nil {
			return nil, fmt.Errorf("gagal memindai virtual host: %w", err)
		}
		vhosts = append(vhosts, vhost)
	}

	return vhosts, nil
}

// Update memperbarui data virtual host
func (r *VirtualHostRepository) Update(id int64, vhost *model.UpdateVirtualHostRequest) error {
	query := `
		UPDATE virtual_hosts
		SET lb_algorithm = ?, sticky_session = ?, failover_mode = ?, is_active = ?, updated_at = ?
		WHERE id = ?
	`

	result, err := r.DB.Exec(query, vhost.LBAlgorithm, vhost.StickySession, vhost.FailoverMode, vhost.IsActive, time.Now(), id)
	if err != nil {
		return fmt.Errorf("gagal mengupdate virtual host: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("gagal mendapatkan rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("virtual host tidak ditemukan")
	}

	return nil
}

// Delete menghapus virtual host
func (r *VirtualHostRepository) Delete(id int64) error {
	query := `DELETE FROM virtual_hosts WHERE id = ?`

	result, err := r.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("gagal menghapus virtual host: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("gagal mendapatkan rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("virtual host tidak ditemukan")
	}

	return nil
}

// Count menghitung total virtual host
func (r *VirtualHostRepository) Count() (int64, error) {
	query := `SELECT COUNT(*) FROM virtual_hosts`

	var count int64
	err := r.DB.QueryRow(query).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("gagal menghitung virtual host: %w", err)
	}

	return count, nil
}
