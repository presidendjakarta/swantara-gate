package repository

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/presidendjakarta/swantara-gate/internal/model"
)

// ConsumerCredentialRepository menangani operasi database untuk Consumer Credential
type ConsumerCredentialRepository struct {
	DB *sql.DB
}

// NewConsumerCredentialRepository membuat instance baru
func NewConsumerCredentialRepository(db *sql.DB) *ConsumerCredentialRepository {
	return &ConsumerCredentialRepository{DB: db}
}

// Create membuat consumer credential baru
func (r *ConsumerCredentialRepository) Create(req *model.CreateConsumerCredentialRequest, passwordHash string) (*model.ConsumerCredential, error) {
	query := `
		INSERT INTO consumer_credentials (consumer_id, auth_type, username, password_hash, api_key, jwt_secret, expired_at, is_active)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
		RETURNING id, created_at
	`

	var cc model.ConsumerCredential
	var expiredAt *time.Time
	if req.ExpiredAt != "" {
		t, err := time.Parse("2006-01-02 15:04:05", req.ExpiredAt)
		if err != nil {
			return nil, fmt.Errorf("format expired_at tidak valid (gunakan YYYY-MM-DD HH:MM:SS): %w", err)
		}
		expiredAt = &t
	}

	err := r.DB.QueryRow(query,
		req.ConsumerID, req.AuthType, req.Username, passwordHash,
		req.APIKey, req.JWTSecret, expiredAt, req.IsActive,
	).Scan(&cc.ID, &cc.CreatedAt)

	if err != nil {
		return nil, fmt.Errorf("gagal membuat consumer credential: %w", err)
	}

	cc.ConsumerID = req.ConsumerID
	cc.AuthType = req.AuthType
	cc.Username = req.Username
	cc.APIKey = req.APIKey
	cc.JWTSecret = req.JWTSecret
	cc.ExpiredAt = expiredAt
	cc.IsActive = req.IsActive

	return &cc, nil
}

// GetByID mengambil credential berdasarkan ID
func (r *ConsumerCredentialRepository) GetByID(id int64) (*model.ConsumerCredential, error) {
	query := `
		SELECT cc.id, cc.consumer_id, cc.auth_type, cc.username, cc.api_key, cc.jwt_secret,
		       cc.expired_at, cc.is_active, cc.created_at, ac.consumer_name
		FROM consumer_credentials cc
		LEFT JOIN api_consumers ac ON cc.consumer_id = ac.id
		WHERE cc.id = ?
	`

	var cc model.ConsumerCredential
	err := r.DB.QueryRow(query, id).Scan(
		&cc.ID, &cc.ConsumerID, &cc.AuthType, &cc.Username, &cc.APIKey, &cc.JWTSecret,
		&cc.ExpiredAt, &cc.IsActive, &cc.CreatedAt, &cc.ConsumerName,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("consumer credential tidak ditemukan")
		}
		return nil, fmt.Errorf("gagal mengambil consumer credential: %w", err)
	}

	return &cc, nil
}

// GetAll mengambil semua credentials dengan pagination
func (r *ConsumerCredentialRepository) GetAll(page, limit int) ([]model.ConsumerCredential, error) {
	offset := (page - 1) * limit
	query := `
		SELECT cc.id, cc.consumer_id, cc.auth_type, cc.username, cc.api_key, cc.jwt_secret,
		       cc.expired_at, cc.is_active, cc.created_at, ac.consumer_name
		FROM consumer_credentials cc
		LEFT JOIN api_consumers ac ON cc.consumer_id = ac.id
		ORDER BY cc.created_at DESC
		LIMIT ? OFFSET ?
	`

	rows, err := r.DB.Query(query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil daftar credential: %w", err)
	}
	defer rows.Close()

	var creds []model.ConsumerCredential
	for rows.Next() {
		var cc model.ConsumerCredential
		err := rows.Scan(
			&cc.ID, &cc.ConsumerID, &cc.AuthType, &cc.Username, &cc.APIKey, &cc.JWTSecret,
			&cc.ExpiredAt, &cc.IsActive, &cc.CreatedAt, &cc.ConsumerName,
		)
		if err != nil {
			return nil, fmt.Errorf("gagal memindai credential: %w", err)
		}
		creds = append(creds, cc)
	}

	return creds, nil
}

// GetByConsumerID mengambil credentials berdasarkan consumer ID
func (r *ConsumerCredentialRepository) GetByConsumerID(consumerID int64) ([]model.ConsumerCredential, error) {
	query := `
		SELECT cc.id, cc.consumer_id, cc.auth_type, cc.username, cc.api_key, cc.jwt_secret,
		       cc.expired_at, cc.is_active, cc.created_at, ac.consumer_name
		FROM consumer_credentials cc
		LEFT JOIN api_consumers ac ON cc.consumer_id = ac.id
		WHERE cc.consumer_id = ?
		ORDER BY cc.created_at DESC
	`

	rows, err := r.DB.Query(query, consumerID)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil credentials: %w", err)
	}
	defer rows.Close()

	var creds []model.ConsumerCredential
	for rows.Next() {
		var cc model.ConsumerCredential
		err := rows.Scan(
			&cc.ID, &cc.ConsumerID, &cc.AuthType, &cc.Username, &cc.APIKey, &cc.JWTSecret,
			&cc.ExpiredAt, &cc.IsActive, &cc.CreatedAt, &cc.ConsumerName,
		)
		if err != nil {
			return nil, fmt.Errorf("gagal memindai credential: %w", err)
		}
		creds = append(creds, cc)
	}

	return creds, nil
}

// Update memperbarui credential
func (r *ConsumerCredentialRepository) Update(id int64, req *model.UpdateConsumerCredentialRequest, passwordHash string) error {
	var expiredAt *time.Time
	if req.ExpiredAt != "" {
		t, err := time.Parse("2006-01-02 15:04:05", req.ExpiredAt)
		if err != nil {
			return fmt.Errorf("format expired_at tidak valid: %w", err)
		}
		expiredAt = &t
	}

	query := `
		UPDATE consumer_credentials
		SET username = ?, password_hash = ?, api_key = ?, jwt_secret = ?, expired_at = ?, is_active = ?
		WHERE id = ?
	`

	result, err := r.DB.Exec(query,
		req.Username, passwordHash, req.APIKey, req.JWTSecret, expiredAt, req.IsActive, id,
	)
	if err != nil {
		return fmt.Errorf("gagal mengupdate credential: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("consumer credential tidak ditemukan")
	}

	return nil
}

// Delete menghapus credential
func (r *ConsumerCredentialRepository) Delete(id int64) error {
	result, err := r.DB.Exec("DELETE FROM consumer_credentials WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("gagal menghapus credential: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("consumer credential tidak ditemukan")
	}

	return nil
}

// Count menghitung total credentials
func (r *ConsumerCredentialRepository) Count() (int64, error) {
	var count int64
	err := r.DB.QueryRow("SELECT COUNT(*) FROM consumer_credentials").Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("gagal menghitung credential: %w", err)
	}
	return count, nil
}

// === API Key Repository ===

// APIKeyRepository menangani operasi database untuk API Keys
type APIKeyRepository struct {
	DB *sql.DB
}

// NewAPIKeyRepository membuat instance baru APIKeyRepository
func NewAPIKeyRepository(db *sql.DB) *APIKeyRepository {
	return &APIKeyRepository{DB: db}
}

// generateAPIKey menghasilkan API key acak dengan panjang 32 karakter hex
func generateAPIKey() (string, error) {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", fmt.Errorf("gagal generate API key: %w", err)
	}
	return "sgk_" + hex.EncodeToString(bytes), nil
}

// Create membuat API key baru
func (r *APIKeyRepository) Create(req *model.CreateAPIKeyRequest) (*model.APIKey, error) {
	key, err := generateAPIKey()
	if err != nil {
		return nil, err
	}

	var expiredAt *time.Time
	if req.ExpiredAt != "" {
		t, err := time.Parse("2006-01-02 15:04:05", req.ExpiredAt)
		if err != nil {
			return nil, fmt.Errorf("format expired_at tidak valid (gunakan YYYY-MM-DD HH:MM:SS): %w", err)
		}
		expiredAt = &t
	}

	query := `
		INSERT INTO api_keys (consumer_id, api_key, description, expired_at, rate_limit_override, is_active)
		VALUES (?, ?, ?, ?, ?, ?)
		RETURNING id, created_at
	`

	var ak model.APIKey
	err = r.DB.QueryRow(query,
		req.ConsumerID, key, req.Description, expiredAt, req.RateLimitOverride, req.IsActive,
	).Scan(&ak.ID, &ak.CreatedAt)

	if err != nil {
		return nil, fmt.Errorf("gagal membuat API key: %w", err)
	}

	ak.ConsumerID = req.ConsumerID
	ak.Key = key
	ak.Description = req.Description
	ak.ExpiredAt = expiredAt
	ak.RateLimitOverride = req.RateLimitOverride
	ak.IsActive = req.IsActive

	return &ak, nil
}

// GetByID mengambil API key berdasarkan ID
func (r *APIKeyRepository) GetByID(id int64) (*model.APIKey, error) {
	query := `
		SELECT ak.id, ak.consumer_id, ak.api_key, ak.description, ak.expired_at,
		       ak.rate_limit_override, ak.is_active, ak.created_at, ac.consumer_name
		FROM api_keys ak
		LEFT JOIN api_consumers ac ON ak.consumer_id = ac.id
		WHERE ak.id = ?
	`

	var ak model.APIKey
	err := r.DB.QueryRow(query, id).Scan(
		&ak.ID, &ak.ConsumerID, &ak.Key, &ak.Description, &ak.ExpiredAt,
		&ak.RateLimitOverride, &ak.IsActive, &ak.CreatedAt, &ak.ConsumerName,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("API key tidak ditemukan")
		}
		return nil, fmt.Errorf("gagal mengambil API key: %w", err)
	}

	return &ak, nil
}

// GetAll mengambil semua API keys dengan pagination
func (r *APIKeyRepository) GetAll(page, limit int) ([]model.APIKey, error) {
	offset := (page - 1) * limit
	query := `
		SELECT ak.id, ak.consumer_id, ak.api_key, ak.description, ak.expired_at,
		       ak.rate_limit_override, ak.is_active, ak.created_at, ac.consumer_name
		FROM api_keys ak
		LEFT JOIN api_consumers ac ON ak.consumer_id = ac.id
		ORDER BY ak.created_at DESC
		LIMIT ? OFFSET ?
	`

	rows, err := r.DB.Query(query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil daftar API keys: %w", err)
	}
	defer rows.Close()

	var keys []model.APIKey
	for rows.Next() {
		var ak model.APIKey
		err := rows.Scan(
			&ak.ID, &ak.ConsumerID, &ak.Key, &ak.Description, &ak.ExpiredAt,
			&ak.RateLimitOverride, &ak.IsActive, &ak.CreatedAt, &ak.ConsumerName,
		)
		if err != nil {
			return nil, fmt.Errorf("gagal memindai API key: %w", err)
		}
		keys = append(keys, ak)
	}

	return keys, nil
}

// GetByConsumerID mengambil API keys berdasarkan consumer ID
func (r *APIKeyRepository) GetByConsumerID(consumerID int64) ([]model.APIKey, error) {
	query := `
		SELECT ak.id, ak.consumer_id, ak.api_key, ak.description, ak.expired_at,
		       ak.rate_limit_override, ak.is_active, ak.created_at, ac.consumer_name
		FROM api_keys ak
		LEFT JOIN api_consumers ac ON ak.consumer_id = ac.id
		WHERE ak.consumer_id = ?
		ORDER BY ak.created_at DESC
	`

	rows, err := r.DB.Query(query, consumerID)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil API keys: %w", err)
	}
	defer rows.Close()

	var keys []model.APIKey
	for rows.Next() {
		var ak model.APIKey
		err := rows.Scan(
			&ak.ID, &ak.ConsumerID, &ak.Key, &ak.Description, &ak.ExpiredAt,
			&ak.RateLimitOverride, &ak.IsActive, &ak.CreatedAt, &ak.ConsumerName,
		)
		if err != nil {
			return nil, fmt.Errorf("gagal memindai API key: %w", err)
		}
		keys = append(keys, ak)
	}

	return keys, nil
}

// Update memperbarui API key
func (r *APIKeyRepository) Update(id int64, req *model.UpdateAPIKeyRequest) error {
	var expiredAt *time.Time
	if req.ExpiredAt != "" {
		t, err := time.Parse("2006-01-02 15:04:05", req.ExpiredAt)
		if err != nil {
			return fmt.Errorf("format expired_at tidak valid: %w", err)
		}
		expiredAt = &t
	}

	query := `
		UPDATE api_keys
		SET description = ?, expired_at = ?, rate_limit_override = ?, is_active = ?
		WHERE id = ?
	`

	result, err := r.DB.Exec(query, req.Description, expiredAt, req.RateLimitOverride, req.IsActive, id)
	if err != nil {
		return fmt.Errorf("gagal mengupdate API key: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("API key tidak ditemukan")
	}

	return nil
}

// Delete menghapus API key
func (r *APIKeyRepository) Delete(id int64) error {
	result, err := r.DB.Exec("DELETE FROM api_keys WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("gagal menghapus API key: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("API key tidak ditemukan")
	}

	return nil
}

// Count menghitung total API keys
func (r *APIKeyRepository) Count() (int64, error) {
	var count int64
	err := r.DB.QueryRow("SELECT COUNT(*) FROM api_keys").Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("gagal menghitung API keys: %w", err)
	}
	return count, nil
}
