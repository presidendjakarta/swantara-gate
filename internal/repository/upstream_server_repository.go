package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/presidendjakarta/swantara-gate/internal/model"
)

// UpstreamServerRepository menangani operasi database untuk Upstream Server
type UpstreamServerRepository struct {
	DB *sql.DB
}

// NewUpstreamServerRepository membuat instance baru UpstreamServerRepository
func NewUpstreamServerRepository(db *sql.DB) *UpstreamServerRepository {
	return &UpstreamServerRepository{DB: db}
}

// Create membuat upstream server baru di database
func (r *UpstreamServerRepository) Create(req *model.CreateUpstreamServerRequest) (*model.UpstreamServer, error) {
	query := `
		INSERT INTO upstream_servers (
			virtual_host_id, target_host, target_port, protocol, priority, weight,
			is_backup, is_active, health_check_enabled, health_check_path,
			health_check_interval_seconds, health_check_timeout_seconds, max_fails, fail_timeout_seconds
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		RETURNING id, created_at, updated_at
	`

	var us model.UpstreamServer
	err := r.DB.QueryRow(query,
		req.VirtualHostID, req.TargetHost, req.TargetPort, req.Protocol,
		req.Priority, req.Weight, req.IsBackup, req.IsActive,
		req.HealthCheckEnabled, req.HealthCheckPath,
		req.HealthCheckIntervalSeconds, req.HealthCheckTimeoutSeconds,
		req.MaxFails, req.FailTimeoutSeconds,
	).Scan(&us.ID, &us.CreatedAt, &us.UpdatedAt)

	if err != nil {
		return nil, fmt.Errorf("gagal membuat upstream server: %w", err)
	}

	us.VirtualHostID = req.VirtualHostID
	us.TargetHost = req.TargetHost
	us.TargetPort = req.TargetPort
	us.Protocol = req.Protocol
	us.Priority = req.Priority
	us.Weight = req.Weight
	us.IsBackup = req.IsBackup
	us.IsActive = req.IsActive
	us.HealthCheckEnabled = req.HealthCheckEnabled
	us.HealthCheckPath = req.HealthCheckPath
	us.HealthCheckIntervalSeconds = req.HealthCheckIntervalSeconds
	us.HealthCheckTimeoutSeconds = req.HealthCheckTimeoutSeconds
	us.MaxFails = req.MaxFails
	us.FailTimeoutSeconds = req.FailTimeoutSeconds

	return &us, nil
}

// GetByID mengambil upstream server berdasarkan ID
func (r *UpstreamServerRepository) GetByID(id int64) (*model.UpstreamServer, error) {
	query := `
		SELECT us.id, us.virtual_host_id, us.target_host, us.target_port, us.protocol,
		       us.priority, us.weight, us.is_backup, us.is_active,
		       us.health_check_enabled, us.health_check_path,
		       us.health_check_interval_seconds, us.health_check_timeout_seconds,
		       us.max_fails, us.fail_timeout_seconds, us.created_at, us.updated_at,
		       vh.vhost_name
		FROM upstream_servers us
		LEFT JOIN virtual_hosts vh ON us.virtual_host_id = vh.id
		WHERE us.id = ?
	`

	var us model.UpstreamServer
	err := r.DB.QueryRow(query, id).Scan(
		&us.ID, &us.VirtualHostID, &us.TargetHost, &us.TargetPort, &us.Protocol,
		&us.Priority, &us.Weight, &us.IsBackup, &us.IsActive,
		&us.HealthCheckEnabled, &us.HealthCheckPath,
		&us.HealthCheckIntervalSeconds, &us.HealthCheckTimeoutSeconds,
		&us.MaxFails, &us.FailTimeoutSeconds, &us.CreatedAt, &us.UpdatedAt,
		&us.VHostName,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("upstream server tidak ditemukan")
		}
		return nil, fmt.Errorf("gagal mengambil upstream server: %w", err)
	}

	return &us, nil
}

// GetAll mengambil semua upstream server dengan pagination
func (r *UpstreamServerRepository) GetAll(page, limit int, search string) ([]model.UpstreamServer, error) {
	offset := (page - 1) * limit
	
	var args []interface{}
	
	query := `
		SELECT us.id, us.virtual_host_id, us.target_host, us.target_port, us.protocol,
		       us.priority, us.weight, us.is_backup, us.is_active,
		       us.health_check_enabled, us.health_check_path,
		       us.health_check_interval_seconds, us.health_check_timeout_seconds,
		       us.max_fails, us.fail_timeout_seconds, us.created_at, us.updated_at,
		       vh.vhost_name
		FROM upstream_servers us
		LEFT JOIN virtual_hosts vh ON us.virtual_host_id = vh.id
	`
	
	// Add search filter if provided
	if search != "" {
		query += ` WHERE us.target_host LIKE ? OR vh.vhost_name LIKE ? OR us.protocol LIKE ?`
		searchPattern := "%" + search + "%"
		args = append(args, searchPattern, searchPattern, searchPattern)
	}
	
	query += ` ORDER BY us.created_at DESC LIMIT ? OFFSET ?`
	args = append(args, limit, offset)

	rows, err := r.DB.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil daftar upstream server: %w", err)
	}
	defer rows.Close()

	var servers []model.UpstreamServer
	for rows.Next() {
		var us model.UpstreamServer
		err := rows.Scan(
			&us.ID, &us.VirtualHostID, &us.TargetHost, &us.TargetPort, &us.Protocol,
			&us.Priority, &us.Weight, &us.IsBackup, &us.IsActive,
			&us.HealthCheckEnabled, &us.HealthCheckPath,
			&us.HealthCheckIntervalSeconds, &us.HealthCheckTimeoutSeconds,
			&us.MaxFails, &us.FailTimeoutSeconds, &us.CreatedAt, &us.UpdatedAt,
			&us.VHostName,
		)
		if err != nil {
			return nil, fmt.Errorf("gagal memindai upstream server: %w", err)
		}
		servers = append(servers, us)
	}

	return servers, nil
}

// GetByVirtualHostID mengambil upstream servers berdasarkan virtual host ID
func (r *UpstreamServerRepository) GetByVirtualHostID(vhostID int64) ([]model.UpstreamServer, error) {
	query := `
		SELECT us.id, us.virtual_host_id, us.target_host, us.target_port, us.protocol,
		       us.priority, us.weight, us.is_backup, us.is_active,
		       us.health_check_enabled, us.health_check_path,
		       us.health_check_interval_seconds, us.health_check_timeout_seconds,
		       us.max_fails, us.fail_timeout_seconds, us.created_at, us.updated_at,
		       vh.vhost_name
		FROM upstream_servers us
		LEFT JOIN virtual_hosts vh ON us.virtual_host_id = vh.id
		WHERE us.virtual_host_id = ?
		ORDER BY us.priority ASC, us.weight DESC
	`

	rows, err := r.DB.Query(query, vhostID)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil upstream servers: %w", err)
	}
	defer rows.Close()

	var servers []model.UpstreamServer
	for rows.Next() {
		var us model.UpstreamServer
		err := rows.Scan(
			&us.ID, &us.VirtualHostID, &us.TargetHost, &us.TargetPort, &us.Protocol,
			&us.Priority, &us.Weight, &us.IsBackup, &us.IsActive,
			&us.HealthCheckEnabled, &us.HealthCheckPath,
			&us.HealthCheckIntervalSeconds, &us.HealthCheckTimeoutSeconds,
			&us.MaxFails, &us.FailTimeoutSeconds, &us.CreatedAt, &us.UpdatedAt,
			&us.VHostName,
		)
		if err != nil {
			return nil, fmt.Errorf("gagal memindai upstream server: %w", err)
		}
		servers = append(servers, us)
	}

	return servers, nil
}

// Update memperbarui data upstream server
func (r *UpstreamServerRepository) Update(id int64, req *model.UpdateUpstreamServerRequest) error {
	query := `
		UPDATE upstream_servers
		SET target_host = ?, target_port = ?, protocol = ?, priority = ?, weight = ?,
		    is_backup = ?, is_active = ?, health_check_enabled = ?, health_check_path = ?,
		    health_check_interval_seconds = ?, health_check_timeout_seconds = ?,
		    max_fails = ?, fail_timeout_seconds = ?, updated_at = ?
		WHERE id = ?
	`

	result, err := r.DB.Exec(query,
		req.TargetHost, req.TargetPort, req.Protocol, req.Priority, req.Weight,
		req.IsBackup, req.IsActive, req.HealthCheckEnabled, req.HealthCheckPath,
		req.HealthCheckIntervalSeconds, req.HealthCheckTimeoutSeconds,
		req.MaxFails, req.FailTimeoutSeconds, time.Now(), id,
	)
	if err != nil {
		return fmt.Errorf("gagal mengupdate upstream server: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("upstream server tidak ditemukan")
	}

	return nil
}

// Delete menghapus upstream server
func (r *UpstreamServerRepository) Delete(id int64) error {
	result, err := r.DB.Exec("DELETE FROM upstream_servers WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("gagal menghapus upstream server: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("upstream server tidak ditemukan")
	}

	return nil
}

// Count menghitung total upstream server
func (r *UpstreamServerRepository) Count(search string) (int64, error) {
	var count int64
	var err error
	
	query := `SELECT COUNT(*) FROM upstream_servers us LEFT JOIN virtual_hosts vh ON us.virtual_host_id = vh.id`
	
	if search != "" {
		query += ` WHERE us.target_host LIKE ? OR vh.vhost_name LIKE ? OR us.protocol LIKE ?`
		searchPattern := "%" + search + "%"
		err = r.DB.QueryRow(query, searchPattern, searchPattern, searchPattern).Scan(&count)
	} else {
		err = r.DB.QueryRow(query).Scan(&count)
	}
	
	if err != nil {
		return 0, fmt.Errorf("gagal menghitung upstream server: %w", err)
	}
	return count, nil
}
