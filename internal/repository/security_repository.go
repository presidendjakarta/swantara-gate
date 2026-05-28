package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/presidendjakarta/swantara-gate/internal/model"
)

// JWTConfigRepository menangani operasi database untuk JWT Config
type JWTConfigRepository struct {
	DB *sql.DB
}

// NewJWTConfigRepository membuat instance baru JWTConfigRepository
func NewJWTConfigRepository(db *sql.DB) *JWTConfigRepository {
	return &JWTConfigRepository{DB: db}
}

// Create membuat JWT config baru
func (r *JWTConfigRepository) Create(req *model.CreateJWTConfigRequest) (*model.JWTConfig, error) {
	query := `
		INSERT INTO jwt_configs (virtual_directory_id, algorithm, jwt_secret, issuer, audience,
			expired_in_seconds, clock_skew_seconds, require_exp, require_iat, is_active)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		RETURNING id, created_at
	`
	var jc model.JWTConfig
	err := r.DB.QueryRow(query,
		req.VirtualDirectoryID, req.Algorithm, req.JWTSecret, req.Issuer, req.Audience,
		req.ExpiredInSeconds, req.ClockSkewSeconds, req.RequireExp, req.RequireIat, req.IsActive,
	).Scan(&jc.ID, &jc.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("gagal membuat JWT config: %w", err)
	}

	jc.VirtualDirectoryID = req.VirtualDirectoryID
	jc.Algorithm = req.Algorithm
	jc.JWTSecret = req.JWTSecret
	jc.Issuer = req.Issuer
	jc.Audience = req.Audience
	jc.ExpiredInSeconds = req.ExpiredInSeconds
	jc.ClockSkewSeconds = req.ClockSkewSeconds
	jc.RequireExp = req.RequireExp
	jc.RequireIat = req.RequireIat
	jc.IsActive = req.IsActive
	return &jc, nil
}

// GetByID mengambil JWT config berdasarkan ID
func (r *JWTConfigRepository) GetByID(id int64) (*model.JWTConfig, error) {
	query := `
		SELECT jc.id, jc.virtual_directory_id, jc.algorithm, jc.jwt_secret, jc.issuer, jc.audience,
		       jc.expired_in_seconds, jc.clock_skew_seconds, jc.require_exp, jc.require_iat,
		       jc.is_active, jc.created_at, vd.source_path
		FROM jwt_configs jc
		LEFT JOIN virtual_directories vd ON jc.virtual_directory_id = vd.id
		WHERE jc.id = ?
	`
	var jc model.JWTConfig
	err := r.DB.QueryRow(query, id).Scan(
		&jc.ID, &jc.VirtualDirectoryID, &jc.Algorithm, &jc.JWTSecret, &jc.Issuer, &jc.Audience,
		&jc.ExpiredInSeconds, &jc.ClockSkewSeconds, &jc.RequireExp, &jc.RequireIat,
		&jc.IsActive, &jc.CreatedAt, &jc.SourcePath,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("JWT config tidak ditemukan")
		}
		return nil, fmt.Errorf("gagal mengambil JWT config: %w", err)
	}
	return &jc, nil
}

// GetAll mengambil semua JWT configs dengan pagination
func (r *JWTConfigRepository) GetAll(page, limit int) ([]model.JWTConfig, error) {
	offset := (page - 1) * limit
	query := `
		SELECT jc.id, jc.virtual_directory_id, jc.algorithm, jc.jwt_secret, jc.issuer, jc.audience,
		       jc.expired_in_seconds, jc.clock_skew_seconds, jc.require_exp, jc.require_iat,
		       jc.is_active, jc.created_at, vd.source_path
		FROM jwt_configs jc
		LEFT JOIN virtual_directories vd ON jc.virtual_directory_id = vd.id
		ORDER BY jc.created_at DESC
		LIMIT ? OFFSET ?
	`
	rows, err := r.DB.Query(query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil daftar JWT config: %w", err)
	}
	defer rows.Close()

	var configs []model.JWTConfig
	for rows.Next() {
		var jc model.JWTConfig
		err := rows.Scan(
			&jc.ID, &jc.VirtualDirectoryID, &jc.Algorithm, &jc.JWTSecret, &jc.Issuer, &jc.Audience,
			&jc.ExpiredInSeconds, &jc.ClockSkewSeconds, &jc.RequireExp, &jc.RequireIat,
			&jc.IsActive, &jc.CreatedAt, &jc.SourcePath,
		)
		if err != nil {
			return nil, fmt.Errorf("gagal memindai JWT config: %w", err)
		}
		configs = append(configs, jc)
	}
	return configs, nil
}

// GetByDirectoryID mengambil JWT configs berdasarkan virtual directory
func (r *JWTConfigRepository) GetByDirectoryID(dirID int64) ([]model.JWTConfig, error) {
	query := `
		SELECT jc.id, jc.virtual_directory_id, jc.algorithm, jc.jwt_secret, jc.issuer, jc.audience,
		       jc.expired_in_seconds, jc.clock_skew_seconds, jc.require_exp, jc.require_iat,
		       jc.is_active, jc.created_at, vd.source_path
		FROM jwt_configs jc
		LEFT JOIN virtual_directories vd ON jc.virtual_directory_id = vd.id
		WHERE jc.virtual_directory_id = ?
	`
	rows, err := r.DB.Query(query, dirID)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil JWT configs: %w", err)
	}
	defer rows.Close()

	var configs []model.JWTConfig
	for rows.Next() {
		var jc model.JWTConfig
		err := rows.Scan(
			&jc.ID, &jc.VirtualDirectoryID, &jc.Algorithm, &jc.JWTSecret, &jc.Issuer, &jc.Audience,
			&jc.ExpiredInSeconds, &jc.ClockSkewSeconds, &jc.RequireExp, &jc.RequireIat,
			&jc.IsActive, &jc.CreatedAt, &jc.SourcePath,
		)
		if err != nil {
			return nil, fmt.Errorf("gagal memindai JWT config: %w", err)
		}
		configs = append(configs, jc)
	}
	return configs, nil
}

// Update memperbarui JWT config
func (r *JWTConfigRepository) Update(id int64, req *model.UpdateJWTConfigRequest) error {
	query := `
		UPDATE jwt_configs
		SET algorithm = ?, jwt_secret = ?, issuer = ?, audience = ?, expired_in_seconds = ?,
		    clock_skew_seconds = ?, require_exp = ?, require_iat = ?, is_active = ?
		WHERE id = ?
	`
	result, err := r.DB.Exec(query,
		req.Algorithm, req.JWTSecret, req.Issuer, req.Audience, req.ExpiredInSeconds,
		req.ClockSkewSeconds, req.RequireExp, req.RequireIat, req.IsActive, id,
	)
	if err != nil {
		return fmt.Errorf("gagal mengupdate JWT config: %w", err)
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("JWT config tidak ditemukan")
	}
	return nil
}

// Delete menghapus JWT config
func (r *JWTConfigRepository) Delete(id int64) error {
	result, err := r.DB.Exec("DELETE FROM jwt_configs WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("gagal menghapus JWT config: %w", err)
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("JWT config tidak ditemukan")
	}
	return nil
}

// Count menghitung total JWT configs
func (r *JWTConfigRepository) Count() (int64, error) {
	var count int64
	err := r.DB.QueryRow("SELECT COUNT(*) FROM jwt_configs").Scan(&count)
	return count, err
}

// === External Auth Repository ===

// ExternalAuthRepository menangani operasi database untuk External Auth
type ExternalAuthRepository struct {
	DB *sql.DB
}

// NewExternalAuthRepository membuat instance baru
func NewExternalAuthRepository(db *sql.DB) *ExternalAuthRepository {
	return &ExternalAuthRepository{DB: db}
}

// Create membuat external auth baru
func (r *ExternalAuthRepository) Create(req *model.CreateExternalAuthRequest) (*model.ExternalAuth, error) {
	query := `
		INSERT INTO external_auth (virtual_directory_id, auth_url, http_method, request_timeout_seconds,
			send_headers, send_body, success_key, success_value, message_key, token_key, is_active)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		RETURNING id, created_at
	`
	var ea model.ExternalAuth
	err := r.DB.QueryRow(query,
		req.VirtualDirectoryID, req.AuthURL, req.HTTPMethod, req.RequestTimeoutSeconds,
		req.SendHeaders, req.SendBody, req.SuccessKey, req.SuccessValue,
		req.MessageKey, req.TokenKey, req.IsActive,
	).Scan(&ea.ID, &ea.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("gagal membuat external auth: %w", err)
	}

	ea.VirtualDirectoryID = req.VirtualDirectoryID
	ea.AuthURL = req.AuthURL
	ea.HTTPMethod = req.HTTPMethod
	ea.RequestTimeoutSeconds = req.RequestTimeoutSeconds
	ea.SendHeaders = req.SendHeaders
	ea.SendBody = req.SendBody
	ea.SuccessKey = req.SuccessKey
	ea.SuccessValue = req.SuccessValue
	ea.MessageKey = req.MessageKey
	ea.TokenKey = req.TokenKey
	ea.IsActive = req.IsActive
	return &ea, nil
}

// GetByID mengambil external auth berdasarkan ID
func (r *ExternalAuthRepository) GetByID(id int64) (*model.ExternalAuth, error) {
	query := `
		SELECT ea.id, ea.virtual_directory_id, ea.auth_url, ea.http_method, ea.request_timeout_seconds,
		       ea.send_headers, ea.send_body, ea.success_key, ea.success_value, ea.message_key,
		       ea.token_key, ea.is_active, ea.created_at, vd.source_path
		FROM external_auth ea
		LEFT JOIN virtual_directories vd ON ea.virtual_directory_id = vd.id
		WHERE ea.id = ?
	`
	var ea model.ExternalAuth
	err := r.DB.QueryRow(query, id).Scan(
		&ea.ID, &ea.VirtualDirectoryID, &ea.AuthURL, &ea.HTTPMethod, &ea.RequestTimeoutSeconds,
		&ea.SendHeaders, &ea.SendBody, &ea.SuccessKey, &ea.SuccessValue, &ea.MessageKey,
		&ea.TokenKey, &ea.IsActive, &ea.CreatedAt, &ea.SourcePath,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("external auth tidak ditemukan")
		}
		return nil, fmt.Errorf("gagal mengambil external auth: %w", err)
	}
	return &ea, nil
}

// GetAll mengambil semua external auth dengan pagination
func (r *ExternalAuthRepository) GetAll(page, limit int) ([]model.ExternalAuth, error) {
	offset := (page - 1) * limit
	query := `
		SELECT ea.id, ea.virtual_directory_id, ea.auth_url, ea.http_method, ea.request_timeout_seconds,
		       ea.send_headers, ea.send_body, ea.success_key, ea.success_value, ea.message_key,
		       ea.token_key, ea.is_active, ea.created_at, vd.source_path
		FROM external_auth ea
		LEFT JOIN virtual_directories vd ON ea.virtual_directory_id = vd.id
		ORDER BY ea.created_at DESC LIMIT ? OFFSET ?
	`
	rows, err := r.DB.Query(query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil daftar external auth: %w", err)
	}
	defer rows.Close()

	var items []model.ExternalAuth
	for rows.Next() {
		var ea model.ExternalAuth
		err := rows.Scan(
			&ea.ID, &ea.VirtualDirectoryID, &ea.AuthURL, &ea.HTTPMethod, &ea.RequestTimeoutSeconds,
			&ea.SendHeaders, &ea.SendBody, &ea.SuccessKey, &ea.SuccessValue, &ea.MessageKey,
			&ea.TokenKey, &ea.IsActive, &ea.CreatedAt, &ea.SourcePath,
		)
		if err != nil {
			return nil, fmt.Errorf("gagal memindai external auth: %w", err)
		}
		items = append(items, ea)
	}
	return items, nil
}

// Update memperbarui external auth
func (r *ExternalAuthRepository) Update(id int64, req *model.UpdateExternalAuthRequest) error {
	query := `
		UPDATE external_auth
		SET auth_url = ?, http_method = ?, request_timeout_seconds = ?, send_headers = ?,
		    send_body = ?, success_key = ?, success_value = ?, message_key = ?, token_key = ?, is_active = ?
		WHERE id = ?
	`
	result, err := r.DB.Exec(query,
		req.AuthURL, req.HTTPMethod, req.RequestTimeoutSeconds, req.SendHeaders,
		req.SendBody, req.SuccessKey, req.SuccessValue, req.MessageKey, req.TokenKey, req.IsActive, id,
	)
	if err != nil {
		return fmt.Errorf("gagal mengupdate external auth: %w", err)
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("external auth tidak ditemukan")
	}
	return nil
}

// Delete menghapus external auth
func (r *ExternalAuthRepository) Delete(id int64) error {
	result, err := r.DB.Exec("DELETE FROM external_auth WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("gagal menghapus external auth: %w", err)
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("external auth tidak ditemukan")
	}
	return nil
}

// Count menghitung total external auth
func (r *ExternalAuthRepository) Count() (int64, error) {
	var count int64
	err := r.DB.QueryRow("SELECT COUNT(*) FROM external_auth").Scan(&count)
	return count, err
}

// === Rate Limit Repository ===

// RateLimitRepository menangani operasi database untuk Rate Limit
type RateLimitRepository struct {
	DB *sql.DB
}

// NewRateLimitRepository membuat instance baru
func NewRateLimitRepository(db *sql.DB) *RateLimitRepository {
	return &RateLimitRepository{DB: db}
}

// Create membuat rate limit baru
func (r *RateLimitRepository) Create(req *model.CreateRateLimitRequest) (*model.RateLimit, error) {
	query := `
		INSERT INTO rate_limits (virtual_directory_id, limit_by, requests_per_minute, burst, block_duration_seconds, is_active)
		VALUES (?, ?, ?, ?, ?, ?)
		RETURNING id, created_at
	`
	var rl model.RateLimit
	err := r.DB.QueryRow(query,
		req.VirtualDirectoryID, req.LimitBy, req.RequestsPerMinute, req.Burst, req.BlockDurationSeconds, req.IsActive,
	).Scan(&rl.ID, &rl.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("gagal membuat rate limit: %w", err)
	}
	rl.VirtualDirectoryID = req.VirtualDirectoryID
	rl.LimitBy = req.LimitBy
	rl.RequestsPerMinute = req.RequestsPerMinute
	rl.Burst = req.Burst
	rl.BlockDurationSeconds = req.BlockDurationSeconds
	rl.IsActive = req.IsActive
	return &rl, nil
}

// GetByID mengambil rate limit berdasarkan ID
func (r *RateLimitRepository) GetByID(id int64) (*model.RateLimit, error) {
	query := `
		SELECT rl.id, rl.virtual_directory_id, rl.limit_by, rl.requests_per_minute, rl.burst,
		       rl.block_duration_seconds, rl.is_active, rl.created_at, vd.source_path
		FROM rate_limits rl
		LEFT JOIN virtual_directories vd ON rl.virtual_directory_id = vd.id
		WHERE rl.id = ?
	`
	var rl model.RateLimit
	err := r.DB.QueryRow(query, id).Scan(
		&rl.ID, &rl.VirtualDirectoryID, &rl.LimitBy, &rl.RequestsPerMinute, &rl.Burst,
		&rl.BlockDurationSeconds, &rl.IsActive, &rl.CreatedAt, &rl.SourcePath,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("rate limit tidak ditemukan")
		}
		return nil, fmt.Errorf("gagal mengambil rate limit: %w", err)
	}
	return &rl, nil
}

// GetAll mengambil semua rate limits dengan pagination
func (r *RateLimitRepository) GetAll(page, limit int) ([]model.RateLimit, error) {
	offset := (page - 1) * limit
	query := `
		SELECT rl.id, rl.virtual_directory_id, rl.limit_by, rl.requests_per_minute, rl.burst,
		       rl.block_duration_seconds, rl.is_active, rl.created_at, vd.source_path
		FROM rate_limits rl
		LEFT JOIN virtual_directories vd ON rl.virtual_directory_id = vd.id
		ORDER BY rl.created_at DESC LIMIT ? OFFSET ?
	`
	rows, err := r.DB.Query(query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil daftar rate limit: %w", err)
	}
	defer rows.Close()

	var items []model.RateLimit
	for rows.Next() {
		var rl model.RateLimit
		err := rows.Scan(
			&rl.ID, &rl.VirtualDirectoryID, &rl.LimitBy, &rl.RequestsPerMinute, &rl.Burst,
			&rl.BlockDurationSeconds, &rl.IsActive, &rl.CreatedAt, &rl.SourcePath,
		)
		if err != nil {
			return nil, fmt.Errorf("gagal memindai rate limit: %w", err)
		}
		items = append(items, rl)
	}
	return items, nil
}

// Update memperbarui rate limit
func (r *RateLimitRepository) Update(id int64, req *model.UpdateRateLimitRequest) error {
	query := `UPDATE rate_limits SET limit_by = ?, requests_per_minute = ?, burst = ?, block_duration_seconds = ?, is_active = ? WHERE id = ?`
	result, err := r.DB.Exec(query, req.LimitBy, req.RequestsPerMinute, req.Burst, req.BlockDurationSeconds, req.IsActive, id)
	if err != nil {
		return fmt.Errorf("gagal mengupdate rate limit: %w", err)
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("rate limit tidak ditemukan")
	}
	return nil
}

// Delete menghapus rate limit
func (r *RateLimitRepository) Delete(id int64) error {
	result, err := r.DB.Exec("DELETE FROM rate_limits WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("gagal menghapus rate limit: %w", err)
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("rate limit tidak ditemukan")
	}
	return nil
}

// Count menghitung total rate limits
func (r *RateLimitRepository) Count() (int64, error) {
	var count int64
	err := r.DB.QueryRow("SELECT COUNT(*) FROM rate_limits").Scan(&count)
	return count, err
}

// === CORS Config Repository ===

// CORSConfigRepository menangani operasi database untuk CORS Config
type CORSConfigRepository struct {
	DB *sql.DB
}

// NewCORSConfigRepository membuat instance baru
func NewCORSConfigRepository(db *sql.DB) *CORSConfigRepository {
	return &CORSConfigRepository{DB: db}
}

// Create membuat CORS config baru
func (r *CORSConfigRepository) Create(req *model.CreateCORSConfigRequest) (*model.CORSConfig, error) {
	query := `
		INSERT INTO cors_configs (virtual_directory_id, allowed_origins, allowed_methods, allowed_headers,
			exposed_headers, allow_credentials, max_age_seconds, is_active)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
		RETURNING id, created_at
	`
	var cc model.CORSConfig
	err := r.DB.QueryRow(query,
		req.VirtualDirectoryID, req.AllowedOrigins, req.AllowedMethods, req.AllowedHeaders,
		req.ExposedHeaders, req.AllowCredentials, req.MaxAgeSeconds, req.IsActive,
	).Scan(&cc.ID, &cc.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("gagal membuat CORS config: %w", err)
	}
	cc.VirtualDirectoryID = req.VirtualDirectoryID
	cc.AllowedOrigins = req.AllowedOrigins
	cc.AllowedMethods = req.AllowedMethods
	cc.AllowedHeaders = req.AllowedHeaders
	cc.ExposedHeaders = req.ExposedHeaders
	cc.AllowCredentials = req.AllowCredentials
	cc.MaxAgeSeconds = req.MaxAgeSeconds
	cc.IsActive = req.IsActive
	return &cc, nil
}

// GetByID mengambil CORS config berdasarkan ID
func (r *CORSConfigRepository) GetByID(id int64) (*model.CORSConfig, error) {
	query := `
		SELECT cc.id, cc.virtual_directory_id, cc.allowed_origins, cc.allowed_methods, cc.allowed_headers,
		       cc.exposed_headers, cc.allow_credentials, cc.max_age_seconds, cc.is_active, cc.created_at, vd.source_path
		FROM cors_configs cc
		LEFT JOIN virtual_directories vd ON cc.virtual_directory_id = vd.id
		WHERE cc.id = ?
	`
	var cc model.CORSConfig
	err := r.DB.QueryRow(query, id).Scan(
		&cc.ID, &cc.VirtualDirectoryID, &cc.AllowedOrigins, &cc.AllowedMethods, &cc.AllowedHeaders,
		&cc.ExposedHeaders, &cc.AllowCredentials, &cc.MaxAgeSeconds, &cc.IsActive, &cc.CreatedAt, &cc.SourcePath,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("CORS config tidak ditemukan")
		}
		return nil, fmt.Errorf("gagal mengambil CORS config: %w", err)
	}
	return &cc, nil
}

// GetAll mengambil semua CORS configs dengan pagination
func (r *CORSConfigRepository) GetAll(page, limit int) ([]model.CORSConfig, error) {
	offset := (page - 1) * limit
	query := `
		SELECT cc.id, cc.virtual_directory_id, cc.allowed_origins, cc.allowed_methods, cc.allowed_headers,
		       cc.exposed_headers, cc.allow_credentials, cc.max_age_seconds, cc.is_active, cc.created_at, vd.source_path
		FROM cors_configs cc
		LEFT JOIN virtual_directories vd ON cc.virtual_directory_id = vd.id
		ORDER BY cc.created_at DESC LIMIT ? OFFSET ?
	`
	rows, err := r.DB.Query(query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil daftar CORS config: %w", err)
	}
	defer rows.Close()

	var items []model.CORSConfig
	for rows.Next() {
		var cc model.CORSConfig
		err := rows.Scan(
			&cc.ID, &cc.VirtualDirectoryID, &cc.AllowedOrigins, &cc.AllowedMethods, &cc.AllowedHeaders,
			&cc.ExposedHeaders, &cc.AllowCredentials, &cc.MaxAgeSeconds, &cc.IsActive, &cc.CreatedAt, &cc.SourcePath,
		)
		if err != nil {
			return nil, fmt.Errorf("gagal memindai CORS config: %w", err)
		}
		items = append(items, cc)
	}
	return items, nil
}

// Update memperbarui CORS config
func (r *CORSConfigRepository) Update(id int64, req *model.UpdateCORSConfigRequest) error {
	query := `UPDATE cors_configs SET allowed_origins = ?, allowed_methods = ?, allowed_headers = ?,
		exposed_headers = ?, allow_credentials = ?, max_age_seconds = ?, is_active = ? WHERE id = ?`
	result, err := r.DB.Exec(query,
		req.AllowedOrigins, req.AllowedMethods, req.AllowedHeaders,
		req.ExposedHeaders, req.AllowCredentials, req.MaxAgeSeconds, req.IsActive, id,
	)
	if err != nil {
		return fmt.Errorf("gagal mengupdate CORS config: %w", err)
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("CORS config tidak ditemukan")
	}
	return nil
}

// Delete menghapus CORS config
func (r *CORSConfigRepository) Delete(id int64) error {
	result, err := r.DB.Exec("DELETE FROM cors_configs WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("gagal menghapus CORS config: %w", err)
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("CORS config tidak ditemukan")
	}
	return nil
}

// Count menghitung total CORS configs
func (r *CORSConfigRepository) Count() (int64, error) {
	var count int64
	err := r.DB.QueryRow("SELECT COUNT(*) FROM cors_configs").Scan(&count)
	return count, err
}

// === Circuit Breaker Repository ===

// CircuitBreakerRepository menangani operasi database untuk Circuit Breaker
type CircuitBreakerRepository struct {
	DB *sql.DB
}

// NewCircuitBreakerRepository membuat instance baru
func NewCircuitBreakerRepository(db *sql.DB) *CircuitBreakerRepository {
	return &CircuitBreakerRepository{DB: db}
}

// Create membuat circuit breaker baru
func (r *CircuitBreakerRepository) Create(req *model.CreateCircuitBreakerRequest) (*model.CircuitBreaker, error) {
	query := `
		INSERT INTO circuit_breakers (virtual_directory_id, enabled, failure_threshold, recovery_timeout_seconds, half_open_max_requests)
		VALUES (?, ?, ?, ?, ?)
		RETURNING id, created_at
	`
	var cb model.CircuitBreaker
	err := r.DB.QueryRow(query,
		req.VirtualDirectoryID, req.Enabled, req.FailureThreshold, req.RecoveryTimeoutSeconds, req.HalfOpenMaxRequests,
	).Scan(&cb.ID, &cb.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("gagal membuat circuit breaker: %w", err)
	}
	cb.VirtualDirectoryID = req.VirtualDirectoryID
	cb.Enabled = req.Enabled
	cb.FailureThreshold = req.FailureThreshold
	cb.RecoveryTimeoutSeconds = req.RecoveryTimeoutSeconds
	cb.HalfOpenMaxRequests = req.HalfOpenMaxRequests
	return &cb, nil
}

// GetByID mengambil circuit breaker berdasarkan ID
func (r *CircuitBreakerRepository) GetByID(id int64) (*model.CircuitBreaker, error) {
	query := `
		SELECT cb.id, cb.virtual_directory_id, cb.enabled, cb.failure_threshold,
		       cb.recovery_timeout_seconds, cb.half_open_max_requests, cb.created_at, vd.source_path
		FROM circuit_breakers cb
		LEFT JOIN virtual_directories vd ON cb.virtual_directory_id = vd.id
		WHERE cb.id = ?
	`
	var cb model.CircuitBreaker
	err := r.DB.QueryRow(query, id).Scan(
		&cb.ID, &cb.VirtualDirectoryID, &cb.Enabled, &cb.FailureThreshold,
		&cb.RecoveryTimeoutSeconds, &cb.HalfOpenMaxRequests, &cb.CreatedAt, &cb.SourcePath,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("circuit breaker tidak ditemukan")
		}
		return nil, fmt.Errorf("gagal mengambil circuit breaker: %w", err)
	}
	return &cb, nil
}

// GetAll mengambil semua circuit breakers dengan pagination
func (r *CircuitBreakerRepository) GetAll(page, limit int) ([]model.CircuitBreaker, error) {
	offset := (page - 1) * limit
	query := `
		SELECT cb.id, cb.virtual_directory_id, cb.enabled, cb.failure_threshold,
		       cb.recovery_timeout_seconds, cb.half_open_max_requests, cb.created_at, vd.source_path
		FROM circuit_breakers cb
		LEFT JOIN virtual_directories vd ON cb.virtual_directory_id = vd.id
		ORDER BY cb.created_at DESC LIMIT ? OFFSET ?
	`
	rows, err := r.DB.Query(query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil daftar circuit breaker: %w", err)
	}
	defer rows.Close()

	var items []model.CircuitBreaker
	for rows.Next() {
		var cb model.CircuitBreaker
		err := rows.Scan(
			&cb.ID, &cb.VirtualDirectoryID, &cb.Enabled, &cb.FailureThreshold,
			&cb.RecoveryTimeoutSeconds, &cb.HalfOpenMaxRequests, &cb.CreatedAt, &cb.SourcePath,
		)
		if err != nil {
			return nil, fmt.Errorf("gagal memindai circuit breaker: %w", err)
		}
		items = append(items, cb)
	}
	return items, nil
}

// Update memperbarui circuit breaker
func (r *CircuitBreakerRepository) Update(id int64, req *model.UpdateCircuitBreakerRequest) error {
	query := `UPDATE circuit_breakers SET enabled = ?, failure_threshold = ?, recovery_timeout_seconds = ?, half_open_max_requests = ? WHERE id = ?`
	result, err := r.DB.Exec(query, req.Enabled, req.FailureThreshold, req.RecoveryTimeoutSeconds, req.HalfOpenMaxRequests, id)
	if err != nil {
		return fmt.Errorf("gagal mengupdate circuit breaker: %w", err)
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("circuit breaker tidak ditemukan")
	}
	return nil
}

// Delete menghapus circuit breaker
func (r *CircuitBreakerRepository) Delete(id int64) error {
	result, err := r.DB.Exec("DELETE FROM circuit_breakers WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("gagal menghapus circuit breaker: %w", err)
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("circuit breaker tidak ditemukan")
	}
	return nil
}

// Count menghitung total circuit breakers
func (r *CircuitBreakerRepository) Count() (int64, error) {
	var count int64
	err := r.DB.QueryRow("SELECT COUNT(*) FROM circuit_breakers").Scan(&count)
	return count, err
}

// === IP Whitelist Repository ===

// IPWhitelistRepository menangani operasi database untuk IP Whitelist
type IPWhitelistRepository struct {
	DB *sql.DB
}

// NewIPWhitelistRepository membuat instance baru
func NewIPWhitelistRepository(db *sql.DB) *IPWhitelistRepository {
	return &IPWhitelistRepository{DB: db}
}

// Create menambahkan IP ke whitelist
func (r *IPWhitelistRepository) Create(req *model.CreateIPWhitelistRequest) (*model.IPWhitelist, error) {
	query := `
		INSERT INTO ip_whitelists (virtual_directory_id, ip_address, description, is_active)
		VALUES (?, ?, ?, ?)
		RETURNING id, created_at
	`
	var ipw model.IPWhitelist
	err := r.DB.QueryRow(query, req.VirtualDirectoryID, req.IPAddress, req.Description, req.IsActive).Scan(&ipw.ID, &ipw.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("gagal menambahkan IP whitelist: %w", err)
	}
	ipw.VirtualDirectoryID = req.VirtualDirectoryID
	ipw.IPAddress = req.IPAddress
	ipw.Description = req.Description
	ipw.IsActive = req.IsActive
	return &ipw, nil
}

// GetByID mengambil IP whitelist berdasarkan ID
func (r *IPWhitelistRepository) GetByID(id int64) (*model.IPWhitelist, error) {
	query := `
		SELECT iw.id, iw.virtual_directory_id, iw.ip_address, iw.description, iw.is_active, iw.created_at, vd.source_path
		FROM ip_whitelists iw
		LEFT JOIN virtual_directories vd ON iw.virtual_directory_id = vd.id
		WHERE iw.id = ?
	`
	var ipw model.IPWhitelist
	err := r.DB.QueryRow(query, id).Scan(&ipw.ID, &ipw.VirtualDirectoryID, &ipw.IPAddress, &ipw.Description, &ipw.IsActive, &ipw.CreatedAt, &ipw.SourcePath)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("IP whitelist tidak ditemukan")
		}
		return nil, fmt.Errorf("gagal mengambil IP whitelist: %w", err)
	}
	return &ipw, nil
}

// GetAll mengambil semua IP whitelist dengan pagination
func (r *IPWhitelistRepository) GetAll(page, limit int) ([]model.IPWhitelist, error) {
	offset := (page - 1) * limit
	query := `
		SELECT iw.id, iw.virtual_directory_id, iw.ip_address, iw.description, iw.is_active, iw.created_at, vd.source_path
		FROM ip_whitelists iw
		LEFT JOIN virtual_directories vd ON iw.virtual_directory_id = vd.id
		ORDER BY iw.created_at DESC LIMIT ? OFFSET ?
	`
	rows, err := r.DB.Query(query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil daftar IP whitelist: %w", err)
	}
	defer rows.Close()

	var items []model.IPWhitelist
	for rows.Next() {
		var ipw model.IPWhitelist
		err := rows.Scan(&ipw.ID, &ipw.VirtualDirectoryID, &ipw.IPAddress, &ipw.Description, &ipw.IsActive, &ipw.CreatedAt, &ipw.SourcePath)
		if err != nil {
			return nil, fmt.Errorf("gagal memindai IP whitelist: %w", err)
		}
		items = append(items, ipw)
	}
	return items, nil
}

// GetByDirectoryID mengambil IP whitelist berdasarkan directory
func (r *IPWhitelistRepository) GetByDirectoryID(dirID int64) ([]model.IPWhitelist, error) {
	query := `
		SELECT iw.id, iw.virtual_directory_id, iw.ip_address, iw.description, iw.is_active, iw.created_at, vd.source_path
		FROM ip_whitelists iw
		LEFT JOIN virtual_directories vd ON iw.virtual_directory_id = vd.id
		WHERE iw.virtual_directory_id = ?
	`
	rows, err := r.DB.Query(query, dirID)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil IP whitelist: %w", err)
	}
	defer rows.Close()

	var items []model.IPWhitelist
	for rows.Next() {
		var ipw model.IPWhitelist
		err := rows.Scan(&ipw.ID, &ipw.VirtualDirectoryID, &ipw.IPAddress, &ipw.Description, &ipw.IsActive, &ipw.CreatedAt, &ipw.SourcePath)
		if err != nil {
			return nil, fmt.Errorf("gagal memindai IP whitelist: %w", err)
		}
		items = append(items, ipw)
	}
	return items, nil
}

// Update memperbarui IP whitelist
func (r *IPWhitelistRepository) Update(id int64, req *model.UpdateIPWhitelistRequest) error {
	query := `UPDATE ip_whitelists SET ip_address = ?, description = ?, is_active = ? WHERE id = ?`
	result, err := r.DB.Exec(query, req.IPAddress, req.Description, req.IsActive, id)
	if err != nil {
		return fmt.Errorf("gagal mengupdate IP whitelist: %w", err)
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("IP whitelist tidak ditemukan")
	}
	return nil
}

// Delete menghapus IP whitelist
func (r *IPWhitelistRepository) Delete(id int64) error {
	result, err := r.DB.Exec("DELETE FROM ip_whitelists WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("gagal menghapus IP whitelist: %w", err)
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("IP whitelist tidak ditemukan")
	}
	return nil
}

// Count menghitung total IP whitelist
func (r *IPWhitelistRepository) Count() (int64, error) {
	var count int64
	err := r.DB.QueryRow("SELECT COUNT(*) FROM ip_whitelists").Scan(&count)
	return count, err
}

// === IP Blacklist Repository ===

// IPBlacklistRepository menangani operasi database untuk IP Blacklist
type IPBlacklistRepository struct {
	DB *sql.DB
}

// NewIPBlacklistRepository membuat instance baru
func NewIPBlacklistRepository(db *sql.DB) *IPBlacklistRepository {
	return &IPBlacklistRepository{DB: db}
}

// Create menambahkan IP ke blacklist
func (r *IPBlacklistRepository) Create(req *model.CreateIPBlacklistRequest) (*model.IPBlacklist, error) {
	var expiredAt *time.Time
	if req.ExpiredAt != "" {
		t, err := time.Parse("2006-01-02 15:04:05", req.ExpiredAt)
		if err != nil {
			return nil, fmt.Errorf("format expired_at tidak valid: %w", err)
		}
		expiredAt = &t
	}

	query := `
		INSERT INTO ip_blacklists (virtual_directory_id, ip_address, reason, expired_at, is_active)
		VALUES (?, ?, ?, ?, ?)
		RETURNING id, created_at
	`
	var ipb model.IPBlacklist
	err := r.DB.QueryRow(query, req.VirtualDirectoryID, req.IPAddress, req.Reason, expiredAt, req.IsActive).Scan(&ipb.ID, &ipb.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("gagal menambahkan IP blacklist: %w", err)
	}
	ipb.VirtualDirectoryID = req.VirtualDirectoryID
	ipb.IPAddress = req.IPAddress
	ipb.Reason = req.Reason
	ipb.ExpiredAt = expiredAt
	ipb.IsActive = req.IsActive
	return &ipb, nil
}

// GetByID mengambil IP blacklist berdasarkan ID
func (r *IPBlacklistRepository) GetByID(id int64) (*model.IPBlacklist, error) {
	query := `
		SELECT ib.id, ib.virtual_directory_id, ib.ip_address, ib.reason, ib.expired_at, ib.is_active, ib.created_at, vd.source_path
		FROM ip_blacklists ib
		LEFT JOIN virtual_directories vd ON ib.virtual_directory_id = vd.id
		WHERE ib.id = ?
	`
	var ipb model.IPBlacklist
	err := r.DB.QueryRow(query, id).Scan(&ipb.ID, &ipb.VirtualDirectoryID, &ipb.IPAddress, &ipb.Reason, &ipb.ExpiredAt, &ipb.IsActive, &ipb.CreatedAt, &ipb.SourcePath)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("IP blacklist tidak ditemukan")
		}
		return nil, fmt.Errorf("gagal mengambil IP blacklist: %w", err)
	}
	return &ipb, nil
}

// GetAll mengambil semua IP blacklist dengan pagination
func (r *IPBlacklistRepository) GetAll(page, limit int) ([]model.IPBlacklist, error) {
	offset := (page - 1) * limit
	query := `
		SELECT ib.id, ib.virtual_directory_id, ib.ip_address, ib.reason, ib.expired_at, ib.is_active, ib.created_at, vd.source_path
		FROM ip_blacklists ib
		LEFT JOIN virtual_directories vd ON ib.virtual_directory_id = vd.id
		ORDER BY ib.created_at DESC LIMIT ? OFFSET ?
	`
	rows, err := r.DB.Query(query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil daftar IP blacklist: %w", err)
	}
	defer rows.Close()

	var items []model.IPBlacklist
	for rows.Next() {
		var ipb model.IPBlacklist
		err := rows.Scan(&ipb.ID, &ipb.VirtualDirectoryID, &ipb.IPAddress, &ipb.Reason, &ipb.ExpiredAt, &ipb.IsActive, &ipb.CreatedAt, &ipb.SourcePath)
		if err != nil {
			return nil, fmt.Errorf("gagal memindai IP blacklist: %w", err)
		}
		items = append(items, ipb)
	}
	return items, nil
}

// GetByDirectoryID mengambil IP blacklist berdasarkan directory
func (r *IPBlacklistRepository) GetByDirectoryID(dirID int64) ([]model.IPBlacklist, error) {
	query := `
		SELECT ib.id, ib.virtual_directory_id, ib.ip_address, ib.reason, ib.expired_at, ib.is_active, ib.created_at, vd.source_path
		FROM ip_blacklists ib
		LEFT JOIN virtual_directories vd ON ib.virtual_directory_id = vd.id
		WHERE ib.virtual_directory_id = ?
	`
	rows, err := r.DB.Query(query, dirID)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil IP blacklist: %w", err)
	}
	defer rows.Close()

	var items []model.IPBlacklist
	for rows.Next() {
		var ipb model.IPBlacklist
		err := rows.Scan(&ipb.ID, &ipb.VirtualDirectoryID, &ipb.IPAddress, &ipb.Reason, &ipb.ExpiredAt, &ipb.IsActive, &ipb.CreatedAt, &ipb.SourcePath)
		if err != nil {
			return nil, fmt.Errorf("gagal memindai IP blacklist: %w", err)
		}
		items = append(items, ipb)
	}
	return items, nil
}

// Update memperbarui IP blacklist
func (r *IPBlacklistRepository) Update(id int64, req *model.UpdateIPBlacklistRequest) error {
	var expiredAt *time.Time
	if req.ExpiredAt != "" {
		t, err := time.Parse("2006-01-02 15:04:05", req.ExpiredAt)
		if err != nil {
			return fmt.Errorf("format expired_at tidak valid: %w", err)
		}
		expiredAt = &t
	}
	query := `UPDATE ip_blacklists SET ip_address = ?, reason = ?, expired_at = ?, is_active = ? WHERE id = ?`
	result, err := r.DB.Exec(query, req.IPAddress, req.Reason, expiredAt, req.IsActive, id)
	if err != nil {
		return fmt.Errorf("gagal mengupdate IP blacklist: %w", err)
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("IP blacklist tidak ditemukan")
	}
	return nil
}

// Delete menghapus IP blacklist
func (r *IPBlacklistRepository) Delete(id int64) error {
	result, err := r.DB.Exec("DELETE FROM ip_blacklists WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("gagal menghapus IP blacklist: %w", err)
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("IP blacklist tidak ditemukan")
	}
	return nil
}

// Count menghitung total IP blacklist
func (r *IPBlacklistRepository) Count() (int64, error) {
	var count int64
	err := r.DB.QueryRow("SELECT COUNT(*) FROM ip_blacklists").Scan(&count)
	return count, err
}
