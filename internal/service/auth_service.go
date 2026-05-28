package service

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/presidendjakarta/swantara-gate/internal/model"
	"github.com/presidendjakarta/swantara-gate/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

// AuthService menangani business logic autentikasi
type AuthService struct {
	AuthRepo             *repository.AuthRepository
	JWTSecret            string
	AccessExpireMinutes  int
	RefreshExpireDays    int
}

// NewAuthService membuat instance baru AuthService
func NewAuthService(authRepo *repository.AuthRepository, jwtSecret string, accessExpMin, refreshExpDays int) *AuthService {
	return &AuthService{
		AuthRepo:            authRepo,
		JWTSecret:           jwtSecret,
		AccessExpireMinutes: accessExpMin,
		RefreshExpireDays:   refreshExpDays,
	}
}

// Login melakukan autentikasi dan menghasilkan token pair
func (s *AuthService) Login(req *model.LoginRequest) (*model.TokenResponse, error) {
	if req.Username == "" {
		return nil, fmt.Errorf("username wajib diisi")
	}
	if req.Password == "" {
		return nil, fmt.Errorf("password wajib diisi")
	}

	// Ambil user berdasarkan username
	user, err := s.AuthRepo.GetUserByUsername(req.Username)
	if err != nil {
		return nil, fmt.Errorf("username atau password salah")
	}

	// Cek apakah user aktif
	if !user.IsActive {
		return nil, fmt.Errorf("akun tidak aktif")
	}

	// Verifikasi password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return nil, fmt.Errorf("username atau password salah")
	}

	// Generate access token
	accessToken, err := s.generateAccessToken(user)
	if err != nil {
		return nil, fmt.Errorf("gagal membuat access token: %w", err)
	}

	// Generate refresh token
	refreshToken, err := s.generateRefreshToken()
	if err != nil {
		return nil, fmt.Errorf("gagal membuat refresh token: %w", err)
	}

	// Simpan refresh token ke database
	refreshExpiry := time.Now().Add(time.Duration(s.RefreshExpireDays) * 24 * time.Hour)
	if err := s.AuthRepo.SaveRefreshToken(user.ID, refreshToken, refreshExpiry); err != nil {
		return nil, fmt.Errorf("gagal menyimpan refresh token: %w", err)
	}

	// Update last login
	_ = s.AuthRepo.UpdateLastLogin(user.ID)

	return &model.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "Bearer",
		ExpiresIn:    s.AccessExpireMinutes * 60,
	}, nil
}

// RefreshToken menghasilkan access token baru dari refresh token
func (s *AuthService) RefreshToken(refreshTokenStr string) (*model.TokenResponse, error) {
	if refreshTokenStr == "" {
		return nil, fmt.Errorf("refresh_token wajib diisi")
	}

	// Cari refresh token di database
	rt, err := s.AuthRepo.GetRefreshToken(refreshTokenStr)
	if err != nil {
		return nil, fmt.Errorf("refresh token tidak valid")
	}

	// Cek apakah sudah direvoke
	if rt.IsRevoked {
		return nil, fmt.Errorf("refresh token sudah dicabut")
	}

	// Cek apakah sudah expired
	if time.Now().After(rt.ExpiresAt) {
		return nil, fmt.Errorf("refresh token sudah expired")
	}

	// Ambil user
	user, err := s.AuthRepo.GetUserByUsername("")
	if err != nil {
		// Cari user by ID via raw query
		user, err = s.getUserByID(rt.UserID)
		if err != nil {
			return nil, fmt.Errorf("user tidak ditemukan")
		}
	}

	if !user.IsActive {
		return nil, fmt.Errorf("akun tidak aktif")
	}

	// Revoke refresh token lama
	_ = s.AuthRepo.RevokeRefreshToken(refreshTokenStr)

	// Generate access token baru
	accessToken, err := s.generateAccessToken(user)
	if err != nil {
		return nil, fmt.Errorf("gagal membuat access token: %w", err)
	}

	// Generate refresh token baru
	newRefreshToken, err := s.generateRefreshToken()
	if err != nil {
		return nil, fmt.Errorf("gagal membuat refresh token: %w", err)
	}

	// Simpan refresh token baru
	refreshExpiry := time.Now().Add(time.Duration(s.RefreshExpireDays) * 24 * time.Hour)
	if err := s.AuthRepo.SaveRefreshToken(rt.UserID, newRefreshToken, refreshExpiry); err != nil {
		return nil, fmt.Errorf("gagal menyimpan refresh token: %w", err)
	}

	return &model.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
		TokenType:    "Bearer",
		ExpiresIn:    s.AccessExpireMinutes * 60,
	}, nil
}

// Logout mencabut refresh token
func (s *AuthService) Logout(refreshTokenStr string) error {
	if refreshTokenStr == "" {
		return fmt.Errorf("refresh_token wajib diisi")
	}
	return s.AuthRepo.RevokeRefreshToken(refreshTokenStr)
}

// ValidateAccessToken memvalidasi access token dan mengembalikan claims
func (s *AuthService) ValidateAccessToken(tokenString string) (*model.JWTClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("metode signing tidak valid")
		}
		return []byte(s.JWTSecret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("token tidak valid: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("token claims tidak valid")
	}

	userID, _ := claims["user_id"].(float64)
	username, _ := claims["username"].(string)
	role, _ := claims["role"].(string)

	return &model.JWTClaims{
		UserID:   int64(userID),
		Username: username,
		Role:     role,
	}, nil
}

// generateAccessToken membuat JWT access token
func (s *AuthService) generateAccessToken(user *model.User) (string, error) {
	now := time.Now()
	claims := jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"role":     user.Role,
		"iat":      now.Unix(),
		"exp":      now.Add(time.Duration(s.AccessExpireMinutes) * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.JWTSecret))
}

// generateRefreshToken membuat random refresh token
func (s *AuthService) generateRefreshToken() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// getUserByID mengambil user berdasarkan ID (internal helper)
func (s *AuthService) getUserByID(userID int64) (*model.User, error) {
	// Kita gunakan query langsung melalui AuthRepo
	query := `SELECT id, username, password_hash, full_name, email, role, is_active, last_login_at, created_at, updated_at FROM users WHERE id = ?`
	var user model.User
	err := s.AuthRepo.DB.QueryRow(query, userID).Scan(
		&user.ID, &user.Username, &user.PasswordHash, &user.FullName,
		&user.Email, &user.Role, &user.IsActive, &user.LastLoginAt,
		&user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("user tidak ditemukan")
	}
	return &user, nil
}
