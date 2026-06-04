package service

import (
	"fmt"

	"github.com/presidendjakarta/swantara-gate/internal/model"
	"github.com/presidendjakarta/swantara-gate/internal/repository"
)

// JWTConfigService menangani business logic untuk JWT Config
type JWTConfigService struct {
	Repo *repository.JWTConfigRepository
}

// NewJWTConfigService membuat instance baru
func NewJWTConfigService(repo *repository.JWTConfigRepository) *JWTConfigService {
	return &JWTConfigService{Repo: repo}
}

// Create membuat JWT config baru dengan validasi
func (s *JWTConfigService) Create(req *model.CreateJWTConfigRequest) (*model.JWTConfig, error) {
	if req.JWTSecret == "" {
		return nil, fmt.Errorf("jwt_secret wajib diisi")
	}
	if req.Algorithm == "" {
		req.Algorithm = "HS256"
	}
	if req.ExpiredInSeconds <= 0 {
		req.ExpiredInSeconds = 3600
	}
	if req.ClockSkewSeconds <= 0 {
		req.ClockSkewSeconds = 30
	}
	return s.Repo.Create(req)
}

// GetByID mengambil JWT config berdasarkan ID
func (s *JWTConfigService) GetByID(id int64) (*model.JWTConfig, error) {
	return s.Repo.GetByID(id)
}

// GetAll mengambil semua JWT configs dengan pagination
func (s *JWTConfigService) GetAll(page, limit int, search string) ([]model.JWTConfig, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}
	total, err := s.Repo.Count(search)
	if err != nil {
		return nil, 0, err
	}
	items, err := s.Repo.GetAll(page, limit, search)
	if err != nil {
		return nil, 0, err
	}
	return items, total, nil
}

// GetByDirectoryID mengambil JWT configs berdasarkan directory
func (s *JWTConfigService) GetByDirectoryID(dirID int64) ([]model.JWTConfig, error) {
	return s.Repo.GetByDirectoryID(dirID)
}

// Update memperbarui JWT config
func (s *JWTConfigService) Update(id int64, req *model.UpdateJWTConfigRequest) error {
	_, err := s.Repo.GetByID(id)
	if err != nil {
		return fmt.Errorf("JWT config tidak ditemukan")
	}
	return s.Repo.Update(id, req)
}

// Delete menghapus JWT config
func (s *JWTConfigService) Delete(id int64) error {
	_, err := s.Repo.GetByID(id)
	if err != nil {
		return fmt.Errorf("JWT config tidak ditemukan")
	}
	return s.Repo.Delete(id)
}

// === External Auth Service ===

// ExternalAuthService menangani business logic untuk External Auth
type ExternalAuthService struct {
	Repo *repository.ExternalAuthRepository
}

// NewExternalAuthService membuat instance baru
func NewExternalAuthService(repo *repository.ExternalAuthRepository) *ExternalAuthService {
	return &ExternalAuthService{Repo: repo}
}

// Create membuat external auth baru dengan validasi
func (s *ExternalAuthService) Create(req *model.CreateExternalAuthRequest) (*model.ExternalAuth, error) {
	if req.AuthURL == "" {
		return nil, fmt.Errorf("auth_url wajib diisi")
	}
	if req.HTTPMethod == "" {
		req.HTTPMethod = "POST"
	}
	if req.RequestTimeoutSeconds <= 0 {
		req.RequestTimeoutSeconds = 5
	}
	if req.SuccessKey == "" {
		req.SuccessKey = "status"
	}
	if req.SuccessValue == "" {
		req.SuccessValue = "true"
	}
	if req.MessageKey == "" {
		req.MessageKey = "message"
	}
	return s.Repo.Create(req)
}

// GetByID mengambil external auth berdasarkan ID
func (s *ExternalAuthService) GetByID(id int64) (*model.ExternalAuth, error) {
	return s.Repo.GetByID(id)
}

// GetAll mengambil semua external auth dengan pagination
func (s *ExternalAuthService) GetAll(page, limit int, search string) ([]model.ExternalAuth, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}
	total, err := s.Repo.Count(search)
	if err != nil {
		return nil, 0, err
	}
	items, err := s.Repo.GetAll(page, limit, search)
	if err != nil {
		return nil, 0, err
	}
	return items, total, nil
}

// Update memperbarui external auth
func (s *ExternalAuthService) Update(id int64, req *model.UpdateExternalAuthRequest) error {
	_, err := s.Repo.GetByID(id)
	if err != nil {
		return fmt.Errorf("external auth tidak ditemukan")
	}
	return s.Repo.Update(id, req)
}

// Delete menghapus external auth
func (s *ExternalAuthService) Delete(id int64) error {
	_, err := s.Repo.GetByID(id)
	if err != nil {
		return fmt.Errorf("external auth tidak ditemukan")
	}
	return s.Repo.Delete(id)
}

// === Rate Limit Service ===

// RateLimitService menangani business logic untuk Rate Limit
type RateLimitService struct {
	Repo *repository.RateLimitRepository
}

// NewRateLimitService membuat instance baru
func NewRateLimitService(repo *repository.RateLimitRepository) *RateLimitService {
	return &RateLimitService{Repo: repo}
}

// Create membuat rate limit baru dengan validasi
func (s *RateLimitService) Create(req *model.CreateRateLimitRequest) (*model.RateLimit, error) {
	if req.LimitBy == "" {
		req.LimitBy = "ip"
	}
	if req.RequestsPerMinute <= 0 {
		req.RequestsPerMinute = 60
	}
	if req.Burst <= 0 {
		req.Burst = 10
	}
	if req.BlockDurationSeconds <= 0 {
		req.BlockDurationSeconds = 60
	}
	return s.Repo.Create(req)
}

// GetByID mengambil rate limit berdasarkan ID
func (s *RateLimitService) GetByID(id int64) (*model.RateLimit, error) {
	return s.Repo.GetByID(id)
}

// GetAll mengambil semua rate limits dengan pagination
func (s *RateLimitService) GetAll(page, limit int, search string) ([]model.RateLimit, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}
	total, err := s.Repo.Count(search)
	if err != nil {
		return nil, 0, err
	}
	items, err := s.Repo.GetAll(page, limit, search)
	if err != nil {
		return nil, 0, err
	}
	return items, total, nil
}

// Update memperbarui rate limit
func (s *RateLimitService) Update(id int64, req *model.UpdateRateLimitRequest) error {
	_, err := s.Repo.GetByID(id)
	if err != nil {
		return fmt.Errorf("rate limit tidak ditemukan")
	}
	return s.Repo.Update(id, req)
}

// Delete menghapus rate limit
func (s *RateLimitService) Delete(id int64) error {
	_, err := s.Repo.GetByID(id)
	if err != nil {
		return fmt.Errorf("rate limit tidak ditemukan")
	}
	return s.Repo.Delete(id)
}

// === CORS Config Service ===

// CORSConfigService menangani business logic untuk CORS Config
type CORSConfigService struct {
	Repo *repository.CORSConfigRepository
}

// NewCORSConfigService membuat instance baru
func NewCORSConfigService(repo *repository.CORSConfigRepository) *CORSConfigService {
	return &CORSConfigService{Repo: repo}
}

// Create membuat CORS config baru dengan validasi
func (s *CORSConfigService) Create(req *model.CreateCORSConfigRequest) (*model.CORSConfig, error) {
	if req.AllowedOrigins == "" {
		req.AllowedOrigins = "*"
	}
	if req.AllowedMethods == "" {
		req.AllowedMethods = "GET,POST,PUT,DELETE,OPTIONS"
	}
	if req.AllowedHeaders == "" {
		req.AllowedHeaders = "*"
	}
	if req.MaxAgeSeconds <= 0 {
		req.MaxAgeSeconds = 3600
	}
	return s.Repo.Create(req)
}

// GetByID mengambil CORS config berdasarkan ID
func (s *CORSConfigService) GetByID(id int64) (*model.CORSConfig, error) {
	return s.Repo.GetByID(id)
}

// GetAll mengambil semua CORS configs dengan pagination
func (s *CORSConfigService) GetAll(page, limit int, search string) ([]model.CORSConfig, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}
	total, err := s.Repo.Count(search)
	if err != nil {
		return nil, 0, err
	}
	items, err := s.Repo.GetAll(page, limit, search)
	if err != nil {
		return nil, 0, err
	}
	return items, total, nil
}

// Update memperbarui CORS config
func (s *CORSConfigService) Update(id int64, req *model.UpdateCORSConfigRequest) error {
	_, err := s.Repo.GetByID(id)
	if err != nil {
		return fmt.Errorf("CORS config tidak ditemukan")
	}
	return s.Repo.Update(id, req)
}

// Delete menghapus CORS config
func (s *CORSConfigService) Delete(id int64) error {
	_, err := s.Repo.GetByID(id)
	if err != nil {
		return fmt.Errorf("CORS config tidak ditemukan")
	}
	return s.Repo.Delete(id)
}

// === Circuit Breaker Service ===

// CircuitBreakerService menangani business logic untuk Circuit Breaker
type CircuitBreakerService struct {
	Repo *repository.CircuitBreakerRepository
}

// NewCircuitBreakerService membuat instance baru
func NewCircuitBreakerService(repo *repository.CircuitBreakerRepository) *CircuitBreakerService {
	return &CircuitBreakerService{Repo: repo}
}

// Create membuat circuit breaker baru dengan validasi
func (s *CircuitBreakerService) Create(req *model.CreateCircuitBreakerRequest) (*model.CircuitBreaker, error) {
	if req.FailureThreshold <= 0 {
		req.FailureThreshold = 5
	}
	if req.RecoveryTimeoutSeconds <= 0 {
		req.RecoveryTimeoutSeconds = 30
	}
	if req.HalfOpenMaxRequests <= 0 {
		req.HalfOpenMaxRequests = 3
	}
	return s.Repo.Create(req)
}

// GetByID mengambil circuit breaker berdasarkan ID
func (s *CircuitBreakerService) GetByID(id int64) (*model.CircuitBreaker, error) {
	return s.Repo.GetByID(id)
}

// GetAll mengambil semua circuit breakers dengan pagination
func (s *CircuitBreakerService) GetAll(page, limit int, search string) ([]model.CircuitBreaker, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}
	total, err := s.Repo.Count(search)
	if err != nil {
		return nil, 0, err
	}
	items, err := s.Repo.GetAll(page, limit, search)
	if err != nil {
		return nil, 0, err
	}
	return items, total, nil
}

// Update memperbarui circuit breaker
func (s *CircuitBreakerService) Update(id int64, req *model.UpdateCircuitBreakerRequest) error {
	_, err := s.Repo.GetByID(id)
	if err != nil {
		return fmt.Errorf("circuit breaker tidak ditemukan")
	}
	return s.Repo.Update(id, req)
}

// Delete menghapus circuit breaker
func (s *CircuitBreakerService) Delete(id int64) error {
	_, err := s.Repo.GetByID(id)
	if err != nil {
		return fmt.Errorf("circuit breaker tidak ditemukan")
	}
	return s.Repo.Delete(id)
}

// === IP Whitelist Service ===

// IPWhitelistService menangani business logic untuk IP Whitelist
type IPWhitelistService struct {
	Repo *repository.IPWhitelistRepository
}

// NewIPWhitelistService membuat instance baru
func NewIPWhitelistService(repo *repository.IPWhitelistRepository) *IPWhitelistService {
	return &IPWhitelistService{Repo: repo}
}

// Create menambahkan IP ke whitelist dengan validasi
func (s *IPWhitelistService) Create(req *model.CreateIPWhitelistRequest) (*model.IPWhitelist, error) {
	if req.IPAddress == "" {
		return nil, fmt.Errorf("ip_address wajib diisi")
	}
	return s.Repo.Create(req)
}

// GetByID mengambil IP whitelist berdasarkan ID
func (s *IPWhitelistService) GetByID(id int64) (*model.IPWhitelist, error) {
	return s.Repo.GetByID(id)
}

// GetAll mengambil semua IP whitelist dengan pagination
func (s *IPWhitelistService) GetAll(page, limit int, search string) ([]model.IPWhitelist, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}
	total, err := s.Repo.Count(search)
	if err != nil {
		return nil, 0, err
	}
	items, err := s.Repo.GetAll(page, limit, search)
	if err != nil {
		return nil, 0, err
	}
	return items, total, nil
}

// GetByDirectoryID mengambil IP whitelist berdasarkan directory
func (s *IPWhitelistService) GetByDirectoryID(dirID int64) ([]model.IPWhitelist, error) {
	return s.Repo.GetByDirectoryID(dirID)
}

// Update memperbarui IP whitelist
func (s *IPWhitelistService) Update(id int64, req *model.UpdateIPWhitelistRequest) error {
	_, err := s.Repo.GetByID(id)
	if err != nil {
		return fmt.Errorf("IP whitelist tidak ditemukan")
	}
	return s.Repo.Update(id, req)
}

// Delete menghapus IP whitelist
func (s *IPWhitelistService) Delete(id int64) error {
	_, err := s.Repo.GetByID(id)
	if err != nil {
		return fmt.Errorf("IP whitelist tidak ditemukan")
	}
	return s.Repo.Delete(id)
}

// === IP Blacklist Service ===

// IPBlacklistService menangani business logic untuk IP Blacklist
type IPBlacklistService struct {
	Repo *repository.IPBlacklistRepository
}

// NewIPBlacklistService membuat instance baru
func NewIPBlacklistService(repo *repository.IPBlacklistRepository) *IPBlacklistService {
	return &IPBlacklistService{Repo: repo}
}

// Create menambahkan IP ke blacklist dengan validasi
func (s *IPBlacklistService) Create(req *model.CreateIPBlacklistRequest) (*model.IPBlacklist, error) {
	if req.IPAddress == "" {
		return nil, fmt.Errorf("ip_address wajib diisi")
	}
	return s.Repo.Create(req)
}

// GetByID mengambil IP blacklist berdasarkan ID
func (s *IPBlacklistService) GetByID(id int64) (*model.IPBlacklist, error) {
	return s.Repo.GetByID(id)
}

// GetAll mengambil semua IP blacklist dengan pagination
func (s *IPBlacklistService) GetAll(page, limit int, search string) ([]model.IPBlacklist, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}
	total, err := s.Repo.Count(search)
	if err != nil {
		return nil, 0, err
	}
	items, err := s.Repo.GetAll(page, limit, search)
	if err != nil {
		return nil, 0, err
	}
	return items, total, nil
}

// GetByDirectoryID mengambil IP blacklist berdasarkan directory
func (s *IPBlacklistService) GetByDirectoryID(dirID int64) ([]model.IPBlacklist, error) {
	return s.Repo.GetByDirectoryID(dirID)
}

// Update memperbarui IP blacklist
func (s *IPBlacklistService) Update(id int64, req *model.UpdateIPBlacklistRequest) error {
	_, err := s.Repo.GetByID(id)
	if err != nil {
		return fmt.Errorf("IP blacklist tidak ditemukan")
	}
	return s.Repo.Update(id, req)
}

// Delete menghapus IP blacklist
func (s *IPBlacklistService) Delete(id int64) error {
	_, err := s.Repo.GetByID(id)
	if err != nil {
		return fmt.Errorf("IP blacklist tidak ditemukan")
	}
	return s.Repo.Delete(id)
}
