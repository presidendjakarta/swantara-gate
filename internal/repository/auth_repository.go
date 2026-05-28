package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/presidendjakarta/swantara-gate/internal/model"
)

// AuthRepository menangani operasi database untuk autentikasi
type AuthRepository struct {
	DB *sql.DB
}

// NewAuthRepository membuat instance baru AuthRepository
func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{DB: db}
}

// GetUserByUsername mengambil user dengan password hash untuk autentikasi
func (r *AuthRepository) GetUserByUsername(username string) (*model.User, error) {
	query := `SELECT id, username, password_hash, full_name, email, role, is_active, last_login_at, created_at, updated_at FROM users WHERE username = ?`
	var user model.User
	err := r.DB.QueryRow(query, username).Scan(
		&user.ID, &user.Username, &user.PasswordHash, &user.FullName,
		&user.Email, &user.Role, &user.IsActive, &user.LastLoginAt,
		&user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user tidak ditemukan")
		}
		return nil, fmt.Errorf("gagal mengambil user: %w", err)
	}
	return &user, nil
}

// UpdateLastLogin memperbarui waktu login terakhir
func (r *AuthRepository) UpdateLastLogin(userID int64) error {
	_, err := r.DB.Exec("UPDATE users SET last_login_at = ?, updated_at = ? WHERE id = ?", time.Now(), time.Now(), userID)
	return err
}

// SaveRefreshToken menyimpan refresh token ke database
func (r *AuthRepository) SaveRefreshToken(userID int64, token string, expiresAt time.Time) error {
	query := `INSERT INTO refresh_tokens (user_id, token, expires_at) VALUES (?, ?, ?)`
	_, err := r.DB.Exec(query, userID, token, expiresAt)
	if err != nil {
		return fmt.Errorf("gagal menyimpan refresh token: %w", err)
	}
	return nil
}

// GetRefreshToken mengambil refresh token dari database
func (r *AuthRepository) GetRefreshToken(token string) (*model.RefreshToken, error) {
	query := `SELECT id, user_id, token, expires_at, is_revoked, created_at FROM refresh_tokens WHERE token = ?`
	var rt model.RefreshToken
	err := r.DB.QueryRow(query, token).Scan(&rt.ID, &rt.UserID, &rt.Token, &rt.ExpiresAt, &rt.IsRevoked, &rt.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("refresh token tidak ditemukan")
		}
		return nil, fmt.Errorf("gagal mengambil refresh token: %w", err)
	}
	return &rt, nil
}

// RevokeRefreshToken mencabut refresh token
func (r *AuthRepository) RevokeRefreshToken(token string) error {
	_, err := r.DB.Exec("UPDATE refresh_tokens SET is_revoked = 1 WHERE token = ?", token)
	return err
}

// RevokeAllUserTokens mencabut semua refresh token milik user
func (r *AuthRepository) RevokeAllUserTokens(userID int64) error {
	_, err := r.DB.Exec("UPDATE refresh_tokens SET is_revoked = 1 WHERE user_id = ?", userID)
	return err
}

// CleanExpiredTokens menghapus token yang sudah expired
func (r *AuthRepository) CleanExpiredTokens() error {
	_, err := r.DB.Exec("DELETE FROM refresh_tokens WHERE expires_at < ? OR is_revoked = 1", time.Now())
	return err
}
