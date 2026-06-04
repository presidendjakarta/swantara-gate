package repository

import (
	"database/sql"
	"fmt"
	"github.com/presidendjakarta/swantara-gate/internal/model"
	"time"
)

// ExternalAuthProviderRepository menangani operasi database untuk external auth providers
type ExternalAuthProviderRepository struct {
	DB *sql.DB
}

// NewExternalAuthProviderRepository membuat instance baru
func NewExternalAuthProviderRepository(db *sql.DB) *ExternalAuthProviderRepository {
	return &ExternalAuthProviderRepository{DB: db}
}

// Create membuat provider baru
func (r *ExternalAuthProviderRepository) Create(req *model.CreateExternalAuthProviderRequest) (*model.ExternalAuthProvider, error) {
	now := time.Now()

	query := `
		INSERT INTO external_auth_providers (
			name, description, auth_url, http_method, request_timeout_seconds,
			send_headers, send_body, success_key, success_value, message_key,
			token_key, is_active, created_at, updated_at
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	result, err := r.DB.Exec(query,
		req.Name, req.Description, req.AuthURL, req.HTTPMethod, req.RequestTimeoutSeconds,
		req.SendHeaders, req.SendBody, req.SuccessKey, req.SuccessValue, req.MessageKey,
		req.TokenKey, req.IsActive, now, now,
	)
	if err != nil {
		return nil, fmt.Errorf("gagal membuat external auth provider: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("gagal mendapatkan ID provider: %w", err)
	}

	return r.GetByID(id)
}

// GetByID mengambil provider berdasarkan ID
func (r *ExternalAuthProviderRepository) GetByID(id int64) (*model.ExternalAuthProvider, error) {
	query := `
		SELECT 
			eap.id, eap.name, eap.description, eap.auth_url, eap.http_method,
			eap.request_timeout_seconds, eap.send_headers, eap.send_body,
			eap.success_key, eap.success_value, eap.message_key, eap.token_key,
			eap.is_active, eap.created_at, eap.updated_at,
			COUNT(vdea.virtual_directory_id) as used_by_count
		FROM external_auth_providers eap
		LEFT JOIN virtual_directory_external_auth vdea ON eap.id = vdea.external_auth_provider_id
		WHERE eap.id = ?
		GROUP BY eap.id
	`

	var provider model.ExternalAuthProvider
	err := r.DB.QueryRow(query, id).Scan(
		&provider.ID, &provider.Name, &provider.Description, &provider.AuthURL, &provider.HTTPMethod,
		&provider.RequestTimeoutSeconds, &provider.SendHeaders, &provider.SendBody,
		&provider.SuccessKey, &provider.SuccessValue, &provider.MessageKey, &provider.TokenKey,
		&provider.IsActive, &provider.CreatedAt, &provider.UpdatedAt,
		&provider.UsedByCount,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("external auth provider tidak ditemukan")
		}
		return nil, fmt.Errorf("gagal mengambil external auth provider: %w", err)
	}

	return &provider, nil
}

// GetAll mengambil semua provider dengan pagination dan search
func (r *ExternalAuthProviderRepository) GetAll(page, limit int, search string) ([]model.ExternalAuthProvider, error) {
	offset := (page - 1) * limit

	var args []interface{}

	query := `
		SELECT 
			eap.id, eap.name, eap.description, eap.auth_url, eap.http_method,
			eap.request_timeout_seconds, eap.send_headers, eap.send_body,
			eap.success_key, eap.success_value, eap.message_key, eap.token_key,
			eap.is_active, eap.created_at, eap.updated_at,
			COUNT(vdea.virtual_directory_id) as used_by_count
		FROM external_auth_providers eap
		LEFT JOIN virtual_directory_external_auth vdea ON eap.id = vdea.external_auth_provider_id
	`

	if search != "" {
		query += ` WHERE eap.name LIKE ? OR eap.description LIKE ? OR eap.auth_url LIKE ?`
		searchPattern := "%" + search + "%"
		args = append(args, searchPattern, searchPattern, searchPattern)
	}

	query += `
		GROUP BY eap.id
		ORDER BY eap.created_at DESC
		LIMIT ? OFFSET ?
	`
	args = append(args, limit, offset)

	rows, err := r.DB.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil daftar external auth providers: %w", err)
	}
	defer rows.Close()

	var providers []model.ExternalAuthProvider
	for rows.Next() {
		var p model.ExternalAuthProvider
		err := rows.Scan(
			&p.ID, &p.Name, &p.Description, &p.AuthURL, &p.HTTPMethod,
			&p.RequestTimeoutSeconds, &p.SendHeaders, &p.SendBody,
			&p.SuccessKey, &p.SuccessValue, &p.MessageKey, &p.TokenKey,
			&p.IsActive, &p.CreatedAt, &p.UpdatedAt,
			&p.UsedByCount,
		)
		if err != nil {
			return nil, fmt.Errorf("gagal memindai provider: %w", err)
		}
		providers = append(providers, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error saat iterasi rows: %w", err)
	}

	return providers, nil
}

// Count menghitung jumlah provider
func (r *ExternalAuthProviderRepository) Count(search string) (int64, error) {
	var count int64
	var err error

	query := `SELECT COUNT(*) FROM external_auth_providers`

	if search != "" {
		query += ` WHERE name LIKE ? OR description LIKE ? OR auth_url LIKE ?`
		searchPattern := "%" + search + "%"
		err = r.DB.QueryRow(query, searchPattern, searchPattern, searchPattern).Scan(&count)
	} else {
		err = r.DB.QueryRow(query).Scan(&count)
	}

	if err != nil {
		return 0, fmt.Errorf("gagal menghitung external auth providers: %w", err)
	}

	return count, nil
}

// Update mengupdate provider
func (r *ExternalAuthProviderRepository) Update(id int64, req *model.UpdateExternalAuthProviderRequest) error {
	now := time.Now()

	query := `
		UPDATE external_auth_providers SET
			name = COALESCE(NULLIF(?, ''), name),
			description = ?,
			auth_url = COALESCE(NULLIF(?, ''), auth_url),
			http_method = COALESCE(NULLIF(?, ''), http_method),
			request_timeout_seconds = ?,
			send_headers = ?,
			send_body = ?,
			success_key = COALESCE(NULLIF(?, ''), success_key),
			success_value = COALESCE(NULLIF(?, ''), success_value),
			message_key = COALESCE(NULLIF(?, ''), message_key),
			token_key = ?,
			is_active = ?,
			updated_at = ?
		WHERE id = ?
	`

	result, err := r.DB.Exec(query,
		req.Name, req.Description, req.AuthURL, req.HTTPMethod, req.RequestTimeoutSeconds,
		req.SendHeaders, req.SendBody, req.SuccessKey, req.SuccessValue, req.MessageKey,
		req.TokenKey, req.IsActive, now, id,
	)
	if err != nil {
		return fmt.Errorf("gagal mengupdate external auth provider: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("gagal mendapatkan rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("external auth provider tidak ditemukan")
	}

	return nil
}

// Delete menghapus provider
func (r *ExternalAuthProviderRepository) Delete(id int64) error {
	result, err := r.DB.Exec("DELETE FROM external_auth_providers WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("gagal menghapus external auth provider: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("external auth provider tidak ditemukan")
	}

	return nil
}

// GetProvidersByVirtualDirectoryID mengambil semua provider yang di-assign ke virtual directory
func (r *ExternalAuthProviderRepository) GetProvidersByVirtualDirectoryID(vdirID int64) ([]model.ExternalAuthProvider, error) {
	query := `
		SELECT 
			eap.id, eap.name, eap.description, eap.auth_url, eap.http_method,
			eap.request_timeout_seconds, eap.send_headers, eap.send_body,
			eap.success_key, eap.success_value, eap.message_key, eap.token_key,
			eap.is_active, eap.created_at, eap.updated_at,
			0 as used_by_count
		FROM external_auth_providers eap
		INNER JOIN virtual_directory_external_auth vdea ON eap.id = vdea.external_auth_provider_id
		WHERE vdea.virtual_directory_id = ?
		ORDER BY eap.name
	`

	rows, err := r.DB.Query(query, vdirID)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil providers untuk virtual directory: %w", err)
	}
	defer rows.Close()

	var providers []model.ExternalAuthProvider
	for rows.Next() {
		var p model.ExternalAuthProvider
		err := rows.Scan(
			&p.ID, &p.Name, &p.Description, &p.AuthURL, &p.HTTPMethod,
			&p.RequestTimeoutSeconds, &p.SendHeaders, &p.SendBody,
			&p.SuccessKey, &p.SuccessValue, &p.MessageKey, &p.TokenKey,
			&p.IsActive, &p.CreatedAt, &p.UpdatedAt,
			&p.UsedByCount,
		)
		if err != nil {
			return nil, fmt.Errorf("gagal memindai provider: %w", err)
		}
		providers = append(providers, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error saat iterasi rows: %w", err)
	}

	return providers, nil
}

// AssignProviderToVirtualDirectory meng-assign provider ke virtual directory
func (r *ExternalAuthProviderRepository) AssignProviderToVirtualDirectory(vdirID, providerID int64) error {
	query := `
		INSERT OR IGNORE INTO virtual_directory_external_auth (
			virtual_directory_id, external_auth_provider_id
		) VALUES (?, ?)
	`

	_, err := r.DB.Exec(query, vdirID, providerID)
	if err != nil {
		return fmt.Errorf("gagal meng-assign provider ke virtual directory: %w", err)
	}

	return nil
}

// RemoveProviderFromVirtualDirectory menghapus assignment provider dari virtual directory
func (r *ExternalAuthProviderRepository) RemoveProviderFromVirtualDirectory(vdirID, providerID int64) error {
	query := `DELETE FROM virtual_directory_external_auth WHERE virtual_directory_id = ? AND external_auth_provider_id = ?`

	result, err := r.DB.Exec(query, vdirID, providerID)
	if err != nil {
		return fmt.Errorf("gagal menghapus assignment: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("assignment tidak ditemukan")
	}

	return nil
}

// GetMappingsByVirtualDirectoryID mengambil semua mapping untuk virtual directory
func (r *ExternalAuthProviderRepository) GetMappingsByVirtualDirectoryID(vdirID int64) ([]model.VirtualDirectoryExternalAuthMapping, error) {
	query := `
		SELECT 
			vdea.virtual_directory_id,
			vdea.external_auth_provider_id,
			vdea.created_at,
			vd.source_path,
			vd.target_path,
			eap.name as provider_name,
			eap.auth_url
		FROM virtual_directory_external_auth vdea
		INNER JOIN virtual_directories vd ON vdea.virtual_directory_id = vd.id
		INNER JOIN external_auth_providers eap ON vdea.external_auth_provider_id = eap.id
		WHERE vdea.virtual_directory_id = ?
		ORDER BY eap.name
	`

	rows, err := r.DB.Query(query, vdirID)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil mappings: %w", err)
	}
	defer rows.Close()

	var mappings []model.VirtualDirectoryExternalAuthMapping
	for rows.Next() {
		var m model.VirtualDirectoryExternalAuthMapping
		err := rows.Scan(
			&m.VirtualDirectoryID, &m.ExternalAuthProviderID, &m.CreatedAt,
			&m.SourcePath, &m.TargetPath, &m.ProviderName, &m.AuthURL,
		)
		if err != nil {
			return nil, fmt.Errorf("gagal memindai mapping: %w", err)
		}
		mappings = append(mappings, m)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error saat iterasi rows: %w", err)
	}

	return mappings, nil
}
