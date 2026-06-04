package repository

import (
	"database/sql"
	"fmt"
	"github.com/presidendjakarta/swantara-gate/internal/model"
	"time"
)

// JWTProviderRepository menangani operasi database untuk JWT providers
type JWTProviderRepository struct {
	DB *sql.DB
}

// NewJWTProviderRepository membuat instance baru
func NewJWTProviderRepository(db *sql.DB) *JWTProviderRepository {
	return &JWTProviderRepository{DB: db}
}

// Create membuat provider baru
func (r *JWTProviderRepository) Create(req *model.CreateJWTProviderRequest) (*model.JWTProvider, error) {
	now := time.Now()

	query := `
		INSERT INTO jwt_providers (
			name, description, algorithm, jwt_secret, issuer, audience,
			expired_in_seconds, require_exp, require_iat, is_active,
			created_at, updated_at
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	result, err := r.DB.Exec(query,
		req.Name, req.Description, req.Algorithm, req.JWTSecret,
		req.Issuer, req.Audience, req.ExpiredInSeconds,
		req.RequireExp, req.RequireIat, req.IsActive,
		now, now,
	)
	if err != nil {
		return nil, fmt.Errorf("gagal membuat JWT provider: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("gagal mendapatkan ID provider: %w", err)
	}

	return r.GetByID(id)
}

// GetByID mengambil provider berdasarkan ID dengan used_by_count
func (r *JWTProviderRepository) GetByID(id int64) (*model.JWTProvider, error) {
	query := `
		SELECT 
			jp.id, jp.name, jp.description, jp.algorithm, jp.jwt_secret,
			jp.issuer, jp.audience, jp.expired_in_seconds, jp.require_exp,
			jp.require_iat, jp.is_active, jp.created_at, jp.updated_at,
			COUNT(vdjp.virtual_directory_id) as used_by_count
		FROM jwt_providers jp
		LEFT JOIN virtual_directory_jwt_providers vdjp ON jp.id = vdjp.jwt_provider_id
		WHERE jp.id = ?
		GROUP BY jp.id
	`

	var provider model.JWTProvider
	var createdAt, updatedAt string

	err := r.DB.QueryRow(query, id).Scan(
		&provider.ID, &provider.Name, &provider.Description, &provider.Algorithm,
		&provider.JWTSecret, &provider.Issuer, &provider.Audience,
		&provider.ExpiredInSeconds, &provider.RequireExp, &provider.RequireIat,
		&provider.IsActive, &createdAt, &updatedAt, &provider.UsedByCount,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("JWT provider tidak ditemukan")
		}
		return nil, fmt.Errorf("gagal mengambil JWT provider: %w", err)
	}

	provider.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", createdAt)
	provider.UpdatedAt, _ = time.Parse("2006-01-02 15:04:05", updatedAt)

	return &provider, nil
}

// GetAll mengambil semua provider dengan pagination dan search
func (r *JWTProviderRepository) GetAll(page, limit int, search string) ([]model.JWTProvider, error) {
	query := `
		SELECT 
			jp.id, jp.name, jp.description, jp.algorithm, jp.jwt_secret,
			jp.issuer, jp.audience, jp.expired_in_seconds, jp.require_exp,
			jp.require_iat, jp.is_active, jp.created_at, jp.updated_at,
			COUNT(vdjp.virtual_directory_id) as used_by_count
		FROM jwt_providers jp
		LEFT JOIN virtual_directory_jwt_providers vdjp ON jp.id = vdjp.jwt_provider_id
	`

	args := []interface{}{}

	if search != "" {
		query += " WHERE jp.name LIKE ? OR jp.description LIKE ?"
		searchPattern := "%" + search + "%"
		args = append(args, searchPattern, searchPattern)
	}

	query += " GROUP BY jp.id ORDER BY jp.created_at DESC LIMIT ? OFFSET ?"
	args = append(args, limit, (page-1)*limit)

	rows, err := r.DB.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil daftar JWT providers: %w", err)
	}
	defer rows.Close()

	var providers []model.JWTProvider
	for rows.Next() {
		var provider model.JWTProvider
		var createdAt, updatedAt string

		err := rows.Scan(
			&provider.ID, &provider.Name, &provider.Description, &provider.Algorithm,
			&provider.JWTSecret, &provider.Issuer, &provider.Audience,
			&provider.ExpiredInSeconds, &provider.RequireExp, &provider.RequireIat,
			&provider.IsActive, &createdAt, &updatedAt, &provider.UsedByCount,
		)
		if err != nil {
			return nil, fmt.Errorf("gagal scan JWT provider: %w", err)
		}

		provider.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", createdAt)
		provider.UpdatedAt, _ = time.Parse("2006-01-02 15:04:05", updatedAt)

		providers = append(providers, provider)
	}

	return providers, nil
}

// Count menghitung total provider dengan filter search
func (r *JWTProviderRepository) Count(search string) (int64, error) {
	query := "SELECT COUNT(*) FROM jwt_providers"
	args := []interface{}{}

	if search != "" {
		query += " WHERE name LIKE ? OR description LIKE ?"
		searchPattern := "%" + search + "%"
		args = append(args, searchPattern, searchPattern)
	}

	var count int64
	err := r.DB.QueryRow(query, args...).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("gagal menghitung JWT providers: %w", err)
	}

	return count, nil
}

// Update mengupdate provider
func (r *JWTProviderRepository) Update(id int64, req *model.UpdateJWTProviderRequest) error {
	query := `
		UPDATE jwt_providers SET
			name = COALESCE(NULLIF(?, ''), name),
			description = ?,
			algorithm = COALESCE(NULLIF(?, ''), algorithm),
			issuer = ?,
			audience = ?,
			expired_in_seconds = CASE WHEN ? > 0 THEN ? ELSE expired_in_seconds END,
			require_exp = ?,
			require_iat = ?,
			is_active = ?,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = ?
	`

	args := []interface{}{
		req.Name, req.Description, req.Algorithm,
		req.Issuer, req.Audience, req.ExpiredInSeconds, req.ExpiredInSeconds,
		req.RequireExp, req.RequireIat, req.IsActive, id,
	}

	// Add jwt_secret only if provided
	if req.JWTSecret != "" {
		query = `
			UPDATE jwt_providers SET
				name = COALESCE(NULLIF(?, ''), name),
				description = ?,
				algorithm = COALESCE(NULLIF(?, ''), algorithm),
				jwt_secret = ?,
				issuer = ?,
				audience = ?,
				expired_in_seconds = CASE WHEN ? > 0 THEN ? ELSE expired_in_seconds END,
				require_exp = ?,
				require_iat = ?,
				is_active = ?,
				updated_at = CURRENT_TIMESTAMP
			WHERE id = ?
		`
		args = []interface{}{
			req.Name, req.Description, req.Algorithm, req.JWTSecret,
			req.Issuer, req.Audience, req.ExpiredInSeconds, req.ExpiredInSeconds,
			req.RequireExp, req.RequireIat, req.IsActive, id,
		}
	}

	result, err := r.DB.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("gagal update JWT provider: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("gagal mendapatkan rows affected: %w", err)
	}

	if rows == 0 {
		return fmt.Errorf("JWT provider tidak ditemukan")
	}

	return nil
}

// Delete menghapus provider
func (r *JWTProviderRepository) Delete(id int64) error {
	result, err := r.DB.Exec("DELETE FROM jwt_providers WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("gagal menghapus JWT provider: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("gagal mendapatkan rows affected: %w", err)
	}

	if rows == 0 {
		return fmt.Errorf("JWT provider tidak ditemukan")
	}

	return nil
}

// GetProvidersByVirtualDirectoryID mengambil semua provider untuk virtual directory
func (r *JWTProviderRepository) GetProvidersByVirtualDirectoryID(vdirID int64) ([]model.JWTProvider, error) {
	query := `
		SELECT jp.id, jp.name, jp.description, jp.algorithm, jp.jwt_secret,
			jp.issuer, jp.audience, jp.expired_in_seconds, jp.require_exp,
			jp.require_iat, jp.is_active, jp.created_at, jp.updated_at,
			0 as used_by_count
		FROM jwt_providers jp
		INNER JOIN virtual_directory_jwt_providers vdjp ON jp.id = vdjp.jwt_provider_id
		WHERE vdjp.virtual_directory_id = ?
		ORDER BY jp.name
	`

	rows, err := r.DB.Query(query, vdirID)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil providers untuk virtual directory: %w", err)
	}
	defer rows.Close()

	var providers []model.JWTProvider
	for rows.Next() {
		var provider model.JWTProvider
		var createdAt, updatedAt string

		err := rows.Scan(
			&provider.ID, &provider.Name, &provider.Description, &provider.Algorithm,
			&provider.JWTSecret, &provider.Issuer, &provider.Audience,
			&provider.ExpiredInSeconds, &provider.RequireExp, &provider.RequireIat,
			&provider.IsActive, &createdAt, &updatedAt, &provider.UsedByCount,
		)
		if err != nil {
			return nil, fmt.Errorf("gagal scan provider: %w", err)
		}

		provider.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", createdAt)
		provider.UpdatedAt, _ = time.Parse("2006-01-02 15:04:05", updatedAt)

		providers = append(providers, provider)
	}

	return providers, nil
}

// AssignProviderToVirtualDirectory meng-assign provider ke virtual directory
func (r *JWTProviderRepository) AssignProviderToVirtualDirectory(vdirID, providerID int64) error {
	query := `
		INSERT OR IGNORE INTO virtual_directory_jwt_providers (
			virtual_directory_id, jwt_provider_id, created_at
		) VALUES (?, ?, CURRENT_TIMESTAMP)
	`

	_, err := r.DB.Exec(query, vdirID, providerID)
	if err != nil {
		return fmt.Errorf("gagal assign provider ke virtual directory: %w", err)
	}

	return nil
}

// RemoveProviderFromVirtualDirectory menghapus assignment
func (r *JWTProviderRepository) RemoveProviderFromVirtualDirectory(vdirID, providerID int64) error {
	result, err := r.DB.Exec(
		"DELETE FROM virtual_directory_jwt_providers WHERE virtual_directory_id = ? AND jwt_provider_id = ?",
		vdirID, providerID,
	)
	if err != nil {
		return fmt.Errorf("gagal menghapus assignment: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("gagal mendapatkan rows affected: %w", err)
	}

	if rows == 0 {
		return fmt.Errorf("assignment tidak ditemukan")
	}

	return nil
}

// GetMappingsByVirtualDirectoryID mengambil semua mapping untuk virtual directory
func (r *JWTProviderRepository) GetMappingsByVirtualDirectoryID(vdirID int64) ([]model.VirtualDirectoryJWTProviderMapping, error) {
	query := `
		SELECT 
			vdjp.virtual_directory_id,
			vdjp.jwt_provider_id,
			vdjp.created_at,
			vd.source_path,
			vd.target_path,
			jp.name as provider_name,
			jp.algorithm
		FROM virtual_directory_jwt_providers vdjp
		INNER JOIN virtual_directories vd ON vdjp.virtual_directory_id = vd.id
		INNER JOIN jwt_providers jp ON vdjp.jwt_provider_id = jp.id
		WHERE vdjp.virtual_directory_id = ?
		ORDER BY vdjp.created_at DESC
	`

	rows, err := r.DB.Query(query, vdirID)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil mappings: %w", err)
	}
	defer rows.Close()

	var mappings []model.VirtualDirectoryJWTProviderMapping
	for rows.Next() {
		var mapping model.VirtualDirectoryJWTProviderMapping
		var createdAt string

		err := rows.Scan(
			&mapping.VirtualDirectoryID, &mapping.JWTProviderID, &createdAt,
			&mapping.SourcePath, &mapping.TargetPath,
			&mapping.ProviderName, &mapping.Algorithm,
		)
		if err != nil {
			return nil, fmt.Errorf("gagal scan mapping: %w", err)
		}

		mapping.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", createdAt)

		mappings = append(mappings, mapping)
	}

	return mappings, nil
}
