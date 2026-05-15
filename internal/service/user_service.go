package service

import (
	"fmt"

	"github.com/presidendjakarta/swantara-gate/internal/model"
	"github.com/presidendjakarta/swantara-gate/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

// UserService menangani business logic untuk User
type UserService struct {
	UserRepo *repository.UserRepository
}

// NewUserService membuat instance baru UserService
func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{UserRepo: userRepo}
}

// CreateUser membuat user baru dengan validasi dan hash password
func (s *UserService) CreateUser(req *model.CreateUserRequest) (*model.User, error) {
	// Validasi role
	validRoles := map[string]bool{
		"super_admin": true,
		"admin":       true,
		"operator":    true,
		"viewer":      true,
	}

	if req.Role == "" {
		req.Role = "admin" // Default role
	}

	if !validRoles[req.Role] {
		return nil, fmt.Errorf("role tidak valid. Harus salah satu dari: super_admin, admin, operator, viewer")
	}

	// Cek apakah username sudah ada
	existingUser, err := s.UserRepo.GetByUsername(req.Username)
	if err == nil && existingUser != nil {
		return nil, fmt.Errorf("username sudah digunakan")
	}

	// Hash password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("gagal menghash password: %w", err)
	}

	// Membuat user baru
	user, err := s.UserRepo.Create(req, string(passwordHash))
	if err != nil {
		return nil, fmt.Errorf("gagal membuat user: %w", err)
	}

	return user, nil
}

// GetUserByID mengambil user berdasarkan ID
func (s *UserService) GetUserByID(id int64) (*model.User, error) {
	user, err := s.UserRepo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil user: %w", err)
	}

	return user, nil
}

// GetAllUsers mengambil semua user dengan pagination
func (s *UserService) GetAllUsers(page, limit int) ([]model.User, int64, error) {
	// Validasi pagination
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	// Mengambil total count
	total, err := s.UserRepo.Count()
	if err != nil {
		return nil, 0, fmt.Errorf("gagal menghitung user: %w", err)
	}

	// Mengambil data user
	users, err := s.UserRepo.GetAll(page, limit)
	if err != nil {
		return nil, 0, fmt.Errorf("gagal mengambil daftar user: %w", err)
	}

	return users, total, nil
}

// UpdateUser memperbarui data user
func (s *UserService) UpdateUser(id int64, req *model.UpdateUserRequest) error {
	// Cek apakah user ada
	_, err := s.UserRepo.GetByID(id)
	if err != nil {
		return fmt.Errorf("gagal update user: %w", err)
	}

	// Validasi role jika diubah
	if req.Role != "" {
		validRoles := map[string]bool{
			"super_admin": true,
			"admin":       true,
			"operator":    true,
			"viewer":      true,
		}

		if !validRoles[req.Role] {
			return fmt.Errorf("role tidak valid. Harus salah satu dari: super_admin, admin, operator, viewer")
		}
	}

	// Update user
	err = s.UserRepo.Update(id, req)
	if err != nil {
		return fmt.Errorf("gagal mengupdate user: %w", err)
	}

	return nil
}

// DeleteUser menghapus user
func (s *UserService) DeleteUser(id int64) error {
	// Cek apakah user ada
	_, err := s.UserRepo.GetByID(id)
	if err != nil {
		return fmt.Errorf("gagal hapus user: %w", err)
	}

	// Delete user
	err = s.UserRepo.Delete(id)
	if err != nil {
		return fmt.Errorf("gagal menghapus user: %w", err)
	}

	return nil
}

// ChangePassword mengubah password user
func (s *UserService) ChangePassword(id int64, req *model.ChangePasswordRequest) error {
	// Ambil user dengan password hash
	user, err := s.UserRepo.GetByID(id)
	if err != nil {
		return fmt.Errorf("user tidak ditemukan")
	}

	// Verifikasi password lama
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.OldPassword))
	if err != nil {
		return fmt.Errorf("password lama salah")
	}

	// Hash password baru
	_, err = bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("gagal menghash password baru: %w", err)
	}

	// Update password di database (perlu query khusus)
	// Untuk sekarang, kita gunakan repository update dengan cara khusus
	// Ini perlu ditambahkan di repository
	return fmt.Errorf("fitur change password belum diimplementasikan di repository")
}
