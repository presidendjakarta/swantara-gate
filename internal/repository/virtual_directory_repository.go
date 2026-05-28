package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/presidendjakarta/swantara-gate/internal/model"
)

// VirtualDirectoryRepository menangani operasi database untuk Virtual Directory
type VirtualDirectoryRepository struct {
	DB *sql.DB
}

// NewVirtualDirectoryRepository membuat instance baru VirtualDirectoryRepository
func NewVirtualDirectoryRepository(db *sql.DB) *VirtualDirectoryRepository {
	return &VirtualDirectoryRepository{DB: db}
}

// Create membuat virtual directory baru
func (r *VirtualDirectoryRepository) Create(req *model.CreateVirtualDirectoryRequest) (*model.VirtualDirectory, error) {
	query := `
		INSERT INTO virtual_directories (
			virtual_host_id, source_path, target_path, match_type, strip_prefix,
			preserve_host_header, auth_type, is_active, proxy_timeout_seconds,
			retry_count, retry_delay_ms, max_request_size_mb, websocket_enabled,
			cache_enabled, cache_ttl_seconds
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		RETURNING id, created_at, updated_at
	`

	var vd model.VirtualDirectory
	err := r.DB.QueryRow(query,
		req.VirtualHostID, req.SourcePath, req.TargetPath, req.MatchType,
		req.StripPrefix, req.PreserveHostHeader, req.AuthType, req.IsActive,
		req.ProxyTimeoutSeconds, req.RetryCount, req.RetryDelayMs,
		req.MaxRequestSizeMB, req.WebsocketEnabled, req.CacheEnabled, req.CacheTTLSeconds,
	).Scan(&vd.ID, &vd.CreatedAt, &vd.UpdatedAt)

	if err != nil {
		return nil, fmt.Errorf("gagal membuat virtual directory: %w", err)
	}

	vd.VirtualHostID = req.VirtualHostID
	vd.SourcePath = req.SourcePath
	vd.TargetPath = req.TargetPath
	vd.MatchType = req.MatchType
	vd.StripPrefix = req.StripPrefix
	vd.PreserveHostHeader = req.PreserveHostHeader
	vd.AuthType = req.AuthType
	vd.IsActive = req.IsActive
	vd.ProxyTimeoutSeconds = req.ProxyTimeoutSeconds
	vd.RetryCount = req.RetryCount
	vd.RetryDelayMs = req.RetryDelayMs
	vd.MaxRequestSizeMB = req.MaxRequestSizeMB
	vd.WebsocketEnabled = req.WebsocketEnabled
	vd.CacheEnabled = req.CacheEnabled
	vd.CacheTTLSeconds = req.CacheTTLSeconds

	return &vd, nil
}

// GetByID mengambil virtual directory berdasarkan ID
func (r *VirtualDirectoryRepository) GetByID(id int64) (*model.VirtualDirectory, error) {
	query := `
		SELECT vd.id, vd.virtual_host_id, vd.source_path, vd.target_path, vd.match_type,
		       vd.strip_prefix, vd.preserve_host_header, vd.auth_type, vd.is_active,
		       vd.proxy_timeout_seconds, vd.retry_count, vd.retry_delay_ms,
		       vd.max_request_size_mb, vd.websocket_enabled, vd.cache_enabled,
		       vd.cache_ttl_seconds, vd.created_at, vd.updated_at,
		       vh.vhost_name
		FROM virtual_directories vd
		LEFT JOIN virtual_hosts vh ON vd.virtual_host_id = vh.id
		WHERE vd.id = ?
	`

	var vd model.VirtualDirectory
	err := r.DB.QueryRow(query, id).Scan(
		&vd.ID, &vd.VirtualHostID, &vd.SourcePath, &vd.TargetPath, &vd.MatchType,
		&vd.StripPrefix, &vd.PreserveHostHeader, &vd.AuthType, &vd.IsActive,
		&vd.ProxyTimeoutSeconds, &vd.RetryCount, &vd.RetryDelayMs,
		&vd.MaxRequestSizeMB, &vd.WebsocketEnabled, &vd.CacheEnabled,
		&vd.CacheTTLSeconds, &vd.CreatedAt, &vd.UpdatedAt,
		&vd.VHostName,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("virtual directory tidak ditemukan")
		}
		return nil, fmt.Errorf("gagal mengambil virtual directory: %w", err)
	}

	return &vd, nil
}

// GetAll mengambil semua virtual directory dengan pagination
func (r *VirtualDirectoryRepository) GetAll(page, limit int) ([]model.VirtualDirectory, error) {
	offset := (page - 1) * limit
	query := `
		SELECT vd.id, vd.virtual_host_id, vd.source_path, vd.target_path, vd.match_type,
		       vd.strip_prefix, vd.preserve_host_header, vd.auth_type, vd.is_active,
		       vd.proxy_timeout_seconds, vd.retry_count, vd.retry_delay_ms,
		       vd.max_request_size_mb, vd.websocket_enabled, vd.cache_enabled,
		       vd.cache_ttl_seconds, vd.created_at, vd.updated_at,
		       vh.vhost_name
		FROM virtual_directories vd
		LEFT JOIN virtual_hosts vh ON vd.virtual_host_id = vh.id
		ORDER BY vd.created_at DESC
		LIMIT ? OFFSET ?
	`

	rows, err := r.DB.Query(query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil daftar virtual directory: %w", err)
	}
	defer rows.Close()

	var dirs []model.VirtualDirectory
	for rows.Next() {
		var vd model.VirtualDirectory
		err := rows.Scan(
			&vd.ID, &vd.VirtualHostID, &vd.SourcePath, &vd.TargetPath, &vd.MatchType,
			&vd.StripPrefix, &vd.PreserveHostHeader, &vd.AuthType, &vd.IsActive,
			&vd.ProxyTimeoutSeconds, &vd.RetryCount, &vd.RetryDelayMs,
			&vd.MaxRequestSizeMB, &vd.WebsocketEnabled, &vd.CacheEnabled,
			&vd.CacheTTLSeconds, &vd.CreatedAt, &vd.UpdatedAt,
			&vd.VHostName,
		)
		if err != nil {
			return nil, fmt.Errorf("gagal memindai virtual directory: %w", err)
		}
		dirs = append(dirs, vd)
	}

	return dirs, nil
}

// GetByVirtualHostID mengambil virtual directories berdasarkan virtual host ID
func (r *VirtualDirectoryRepository) GetByVirtualHostID(vhostID int64) ([]model.VirtualDirectory, error) {
	query := `
		SELECT vd.id, vd.virtual_host_id, vd.source_path, vd.target_path, vd.match_type,
		       vd.strip_prefix, vd.preserve_host_header, vd.auth_type, vd.is_active,
		       vd.proxy_timeout_seconds, vd.retry_count, vd.retry_delay_ms,
		       vd.max_request_size_mb, vd.websocket_enabled, vd.cache_enabled,
		       vd.cache_ttl_seconds, vd.created_at, vd.updated_at,
		       vh.vhost_name
		FROM virtual_directories vd
		LEFT JOIN virtual_hosts vh ON vd.virtual_host_id = vh.id
		WHERE vd.virtual_host_id = ?
		ORDER BY vd.source_path ASC
	`

	rows, err := r.DB.Query(query, vhostID)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil virtual directories: %w", err)
	}
	defer rows.Close()

	var dirs []model.VirtualDirectory
	for rows.Next() {
		var vd model.VirtualDirectory
		err := rows.Scan(
			&vd.ID, &vd.VirtualHostID, &vd.SourcePath, &vd.TargetPath, &vd.MatchType,
			&vd.StripPrefix, &vd.PreserveHostHeader, &vd.AuthType, &vd.IsActive,
			&vd.ProxyTimeoutSeconds, &vd.RetryCount, &vd.RetryDelayMs,
			&vd.MaxRequestSizeMB, &vd.WebsocketEnabled, &vd.CacheEnabled,
			&vd.CacheTTLSeconds, &vd.CreatedAt, &vd.UpdatedAt,
			&vd.VHostName,
		)
		if err != nil {
			return nil, fmt.Errorf("gagal memindai virtual directory: %w", err)
		}
		dirs = append(dirs, vd)
	}

	return dirs, nil
}

// Update memperbarui data virtual directory
func (r *VirtualDirectoryRepository) Update(id int64, req *model.UpdateVirtualDirectoryRequest) error {
	query := `
		UPDATE virtual_directories
		SET source_path = ?, target_path = ?, match_type = ?, strip_prefix = ?,
		    preserve_host_header = ?, auth_type = ?, is_active = ?,
		    proxy_timeout_seconds = ?, retry_count = ?, retry_delay_ms = ?,
		    max_request_size_mb = ?, websocket_enabled = ?, cache_enabled = ?,
		    cache_ttl_seconds = ?, updated_at = ?
		WHERE id = ?
	`

	result, err := r.DB.Exec(query,
		req.SourcePath, req.TargetPath, req.MatchType, req.StripPrefix,
		req.PreserveHostHeader, req.AuthType, req.IsActive,
		req.ProxyTimeoutSeconds, req.RetryCount, req.RetryDelayMs,
		req.MaxRequestSizeMB, req.WebsocketEnabled, req.CacheEnabled,
		req.CacheTTLSeconds, time.Now(), id,
	)
	if err != nil {
		return fmt.Errorf("gagal mengupdate virtual directory: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("virtual directory tidak ditemukan")
	}

	return nil
}

// Delete menghapus virtual directory
func (r *VirtualDirectoryRepository) Delete(id int64) error {
	result, err := r.DB.Exec("DELETE FROM virtual_directories WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("gagal menghapus virtual directory: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("virtual directory tidak ditemukan")
	}

	return nil
}

// Count menghitung total virtual directory
func (r *VirtualDirectoryRepository) Count() (int64, error) {
	var count int64
	err := r.DB.QueryRow("SELECT COUNT(*) FROM virtual_directories").Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("gagal menghitung virtual directory: %w", err)
	}
	return count, nil
}

// === Virtual Directory Methods ===

// VirtualDirectoryMethodRepository menangani operasi database untuk methods
type VirtualDirectoryMethodRepository struct {
	DB *sql.DB
}

// NewVirtualDirectoryMethodRepository membuat instance baru
func NewVirtualDirectoryMethodRepository(db *sql.DB) *VirtualDirectoryMethodRepository {
	return &VirtualDirectoryMethodRepository{DB: db}
}

// GetByDirectoryID mengambil methods berdasarkan virtual directory ID
func (r *VirtualDirectoryMethodRepository) GetByDirectoryID(dirID int64) ([]model.VirtualDirectoryMethod, error) {
	query := `SELECT id, virtual_directory_id, http_method FROM virtual_directory_methods WHERE virtual_directory_id = ?`

	rows, err := r.DB.Query(query, dirID)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil methods: %w", err)
	}
	defer rows.Close()

	var methods []model.VirtualDirectoryMethod
	for rows.Next() {
		var m model.VirtualDirectoryMethod
		err := rows.Scan(&m.ID, &m.VirtualDirectoryID, &m.HTTPMethod)
		if err != nil {
			return nil, fmt.Errorf("gagal memindai method: %w", err)
		}
		methods = append(methods, m)
	}

	return methods, nil
}

// SetMethods mengatur ulang methods untuk virtual directory (hapus lalu insert ulang)
func (r *VirtualDirectoryMethodRepository) SetMethods(dirID int64, methods []string) ([]model.VirtualDirectoryMethod, error) {
	tx, err := r.DB.Begin()
	if err != nil {
		return nil, fmt.Errorf("gagal memulai transaksi: %w", err)
	}
	defer tx.Rollback()

	// Hapus methods lama
	_, err = tx.Exec("DELETE FROM virtual_directory_methods WHERE virtual_directory_id = ?", dirID)
	if err != nil {
		return nil, fmt.Errorf("gagal menghapus methods lama: %w", err)
	}

	// Insert methods baru
	var result []model.VirtualDirectoryMethod
	for _, method := range methods {
		var m model.VirtualDirectoryMethod
		err := tx.QueryRow(
			"INSERT INTO virtual_directory_methods (virtual_directory_id, http_method) VALUES (?, ?) RETURNING id",
			dirID, method,
		).Scan(&m.ID)
		if err != nil {
			return nil, fmt.Errorf("gagal menambahkan method %s: %w", method, err)
		}
		m.VirtualDirectoryID = dirID
		m.HTTPMethod = method
		result = append(result, m)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("gagal commit transaksi: %w", err)
	}

	return result, nil
}

// DeleteByDirectoryID menghapus semua methods untuk virtual directory
func (r *VirtualDirectoryMethodRepository) DeleteByDirectoryID(dirID int64) error {
	_, err := r.DB.Exec("DELETE FROM virtual_directory_methods WHERE virtual_directory_id = ?", dirID)
	if err != nil {
		return fmt.Errorf("gagal menghapus methods: %w", err)
	}
	return nil
}
