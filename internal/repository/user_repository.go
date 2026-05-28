package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/presidendjakarta/swantara-gate/internal/model"
)

// UserRepository menangani operasi database untuk User
type UserRepository struct {
	DB *sql.DB
}

// NewUserRepository membuat instance baru UserRepository
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// Create membuat user baru di database
func (r *UserRepository) Create(user *model.CreateUserRequest, passwordHash string) (*model.User, error) {
	query := `
		INSERT INTO users (username, password_hash, full_name, email, role, is_active)
		VALUES (?, ?, ?, ?, ?, ?)
		RETURNING id, created_at, updated_at
	`

	var createdUser model.User
	err := r.DB.QueryRow(
		query,
		user.Username,
		passwordHash,
		user.FullName,
		user.Email,
		user.Role,
		user.IsActive,
	).Scan(&createdUser.ID, &createdUser.CreatedAt, &createdUser.UpdatedAt)

	if err != nil {
		return nil, fmt.Errorf("gagal membuat user: %w", err)
	}

	// Mengisi data user yang baru dibuat
	createdUser.Username = user.Username
	createdUser.FullName = user.FullName
	createdUser.Email = user.Email
	createdUser.Role = user.Role
	createdUser.IsActive = user.IsActive

	return &createdUser, nil
}

// GetByID mengambil user berdasarkan ID
func (r *UserRepository) GetByID(id int64) (*model.User, error) {
	query := `
		SELECT id, username, full_name, email, role, is_active, last_login_at, created_at, updated_at
		FROM users
		WHERE id = ?
	`

	var user model.User
	err := r.DB.QueryRow(query, id).Scan(
		&user.ID,
		&user.Username,
		&user.FullName,
		&user.Email,
		&user.Role,
		&user.IsActive,
		&user.LastLoginAt,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user tidak ditemukan")
		}
		return nil, fmt.Errorf("gagal mengambil user: %w", err)
	}

	return &user, nil
}

// GetAll mengambil semua user dengan pagination
func (r *UserRepository) GetAll(page, limit int) ([]model.User, error) {
	offset := (page - 1) * limit
	
	query := `
		SELECT id, username, full_name, email, role, is_active, last_login_at, created_at, updated_at
		FROM users
		ORDER BY created_at DESC
		LIMIT ? OFFSET ?
	`

	rows, err := r.DB.Query(query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil daftar user: %w", err)
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.FullName,
			&user.Email,
			&user.Role,
			&user.IsActive,
			&user.LastLoginAt,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("gagal memindai user: %w", err)
		}
		users = append(users, user)
	}

	return users, nil
}

// GetByUsername mengambil user berdasarkan username
func (r *UserRepository) GetByUsername(username string) (*model.User, error) {
	query := `
		SELECT id, username, password_hash, full_name, email, role, is_active, last_login_at, created_at, updated_at
		FROM users
		WHERE username = ?
	`

	var user model.User
	err := r.DB.QueryRow(query, username).Scan(
		&user.ID,
		&user.Username,
		&user.PasswordHash,
		&user.FullName,
		&user.Email,
		&user.Role,
		&user.IsActive,
		&user.LastLoginAt,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user tidak ditemukan")
		}
		return nil, fmt.Errorf("gagal mengambil user: %w", err)
	}

	return &user, nil
}

// Update memperbarui data user
func (r *UserRepository) Update(id int64, user *model.UpdateUserRequest) error {
	query := `
		UPDATE users
		SET full_name = ?, email = ?, role = ?, is_active = ?, updated_at = ?
		WHERE id = ?
	`

	result, err := r.DB.Exec(query, user.FullName, user.Email, user.Role, user.IsActive, time.Now(), id)
	if err != nil {
		return fmt.Errorf("gagal mengupdate user: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("gagal mendapatkan rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user tidak ditemukan")
	}

	return nil
}

// Delete menghapus user berdasarkan ID
func (r *UserRepository) Delete(id int64) error {
	query := `DELETE FROM users WHERE id = ?`

	result, err := r.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("gagal menghapus user: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("gagal mendapatkan rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user tidak ditemukan")
	}

	return nil
}

// UpdateLastLogin memperbarui waktu terakhir login
func (r *UserRepository) UpdateLastLogin(id int64) error {
	query := `UPDATE users SET last_login_at = ?, updated_at = ? WHERE id = ?`

	_, err := r.DB.Exec(query, time.Now(), time.Now(), id)
	if err != nil {
		return fmt.Errorf("gagal mengupdate last_login_at: %w", err)
	}

	return nil
}

// Count menghitung total jumlah user
func (r *UserRepository) Count() (int64, error) {
	query := `SELECT COUNT(*) FROM users`

	var count int64
	err := r.DB.QueryRow(query).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("gagal menghitung user: %w", err)
	}

	return count, nil
}

// UpdatePassword memperbarui password hash user berdasarkan ID
func (r *UserRepository) UpdatePassword(id int64, passwordHash string) error {
	query := `UPDATE users SET password_hash = ?, updated_at = ? WHERE id = ?`

	result, err := r.DB.Exec(query, passwordHash, time.Now(), id)
	if err != nil {
		return fmt.Errorf("gagal mengupdate password: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("gagal mendapatkan rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user tidak ditemukan")
	}

	return nil
}

// UpdatePasswordByUsername memperbarui password hash user berdasarkan username
func (r *UserRepository) UpdatePasswordByUsername(username string, passwordHash string) error {
	query := `UPDATE users SET password_hash = ?, updated_at = ? WHERE username = ?`

	result, err := r.DB.Exec(query, passwordHash, time.Now(), username)
	if err != nil {
		return fmt.Errorf("gagal mengupdate password: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("gagal mendapatkan rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user dengan username '%s' tidak ditemukan", username)
	}

	return nil
}
