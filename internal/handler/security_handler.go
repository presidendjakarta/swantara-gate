package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/presidendjakarta/swantara-gate/internal/model"
	"github.com/presidendjakarta/swantara-gate/internal/response"
	"github.com/presidendjakarta/swantara-gate/internal/service"
)

// JWTConfigHandler menangani HTTP request untuk JWT Config
type JWTConfigHandler struct {
	Service *service.JWTConfigService
}

// NewJWTConfigHandler membuat instance baru
func NewJWTConfigHandler(svc *service.JWTConfigService) *JWTConfigHandler {
	return &JWTConfigHandler{Service: svc}
}

// Create handler untuk membuat JWT config baru
func (h *JWTConfigHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req model.CreateJWTConfigRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Request body tidak valid")
		return
	}
	if req.VirtualDirectoryID <= 0 {
		response.BadRequest(w, "virtual_directory_id wajib diisi")
		return
	}
	item, err := h.Service.Create(&req)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}
	response.Created(w, "JWT config berhasil dibuat", item)
}

// GetByID handler untuk mengambil JWT config berdasarkan ID
func (h *JWTConfigHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "ID tidak valid")
		return
	}
	item, err := h.Service.GetByID(id)
	if err != nil {
		response.NotFound(w, "JWT config tidak ditemukan")
		return
	}
	response.Success(w, "JWT config berhasil diambil", item)
}

// GetAll handler untuk mengambil semua JWT configs
func (h *JWTConfigHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	page, limit := parsePagination(r)
	items, total, err := h.Service.GetAll(page, limit)
	if err != nil {
		response.InternalServerError(w, "Gagal mengambil daftar JWT config")
		return
	}
	result := map[string]interface{}{
		"jwt_configs": items,
		"pagination":  paginationData(page, limit, total),
	}
	response.Success(w, "Daftar JWT config berhasil diambil", result)
}

// Update handler untuk memperbarui JWT config
func (h *JWTConfigHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "ID tidak valid")
		return
	}
	var req model.UpdateJWTConfigRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Request body tidak valid")
		return
	}
	err = h.Service.Update(id, &req)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}
	response.Success(w, "JWT config berhasil diupdate", nil)
}

// Delete handler untuk menghapus JWT config
func (h *JWTConfigHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "ID tidak valid")
		return
	}
	err = h.Service.Delete(id)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}
	response.Success(w, "JWT config berhasil dihapus", nil)
}

// === External Auth Handler ===

// ExternalAuthHandler menangani HTTP request untuk External Auth
type ExternalAuthHandler struct {
	Service *service.ExternalAuthService
}

// NewExternalAuthHandler membuat instance baru
func NewExternalAuthHandler(svc *service.ExternalAuthService) *ExternalAuthHandler {
	return &ExternalAuthHandler{Service: svc}
}

// Create handler untuk membuat external auth baru
func (h *ExternalAuthHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req model.CreateExternalAuthRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Request body tidak valid")
		return
	}
	if req.VirtualDirectoryID <= 0 {
		response.BadRequest(w, "virtual_directory_id wajib diisi")
		return
	}
	item, err := h.Service.Create(&req)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}
	response.Created(w, "External auth berhasil dibuat", item)
}

// GetByID handler untuk mengambil external auth berdasarkan ID
func (h *ExternalAuthHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "ID tidak valid")
		return
	}
	item, err := h.Service.GetByID(id)
	if err != nil {
		response.NotFound(w, "External auth tidak ditemukan")
		return
	}
	response.Success(w, "External auth berhasil diambil", item)
}

// GetAll handler untuk mengambil semua external auth
func (h *ExternalAuthHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	page, limit := parsePagination(r)
	items, total, err := h.Service.GetAll(page, limit)
	if err != nil {
		response.InternalServerError(w, "Gagal mengambil daftar external auth")
		return
	}
	result := map[string]interface{}{
		"external_auth": items,
		"pagination":    paginationData(page, limit, total),
	}
	response.Success(w, "Daftar external auth berhasil diambil", result)
}

// Update handler untuk memperbarui external auth
func (h *ExternalAuthHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "ID tidak valid")
		return
	}
	var req model.UpdateExternalAuthRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Request body tidak valid")
		return
	}
	err = h.Service.Update(id, &req)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}
	response.Success(w, "External auth berhasil diupdate", nil)
}

// Delete handler untuk menghapus external auth
func (h *ExternalAuthHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "ID tidak valid")
		return
	}
	err = h.Service.Delete(id)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}
	response.Success(w, "External auth berhasil dihapus", nil)
}

// === Rate Limit Handler ===

// RateLimitHandler menangani HTTP request untuk Rate Limit
type RateLimitHandler struct {
	Service *service.RateLimitService
}

// NewRateLimitHandler membuat instance baru
func NewRateLimitHandler(svc *service.RateLimitService) *RateLimitHandler {
	return &RateLimitHandler{Service: svc}
}

// Create handler untuk membuat rate limit baru
func (h *RateLimitHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req model.CreateRateLimitRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Request body tidak valid")
		return
	}
	if req.VirtualDirectoryID <= 0 {
		response.BadRequest(w, "virtual_directory_id wajib diisi")
		return
	}
	item, err := h.Service.Create(&req)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}
	response.Created(w, "Rate limit berhasil dibuat", item)
}

// GetByID handler untuk mengambil rate limit berdasarkan ID
func (h *RateLimitHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "ID tidak valid")
		return
	}
	item, err := h.Service.GetByID(id)
	if err != nil {
		response.NotFound(w, "Rate limit tidak ditemukan")
		return
	}
	response.Success(w, "Rate limit berhasil diambil", item)
}

// GetAll handler untuk mengambil semua rate limits
func (h *RateLimitHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	page, limit := parsePagination(r)
	items, total, err := h.Service.GetAll(page, limit)
	if err != nil {
		response.InternalServerError(w, "Gagal mengambil daftar rate limit")
		return
	}
	result := map[string]interface{}{
		"rate_limits": items,
		"pagination":  paginationData(page, limit, total),
	}
	response.Success(w, "Daftar rate limit berhasil diambil", result)
}

// Update handler untuk memperbarui rate limit
func (h *RateLimitHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "ID tidak valid")
		return
	}
	var req model.UpdateRateLimitRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Request body tidak valid")
		return
	}
	err = h.Service.Update(id, &req)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}
	response.Success(w, "Rate limit berhasil diupdate", nil)
}

// Delete handler untuk menghapus rate limit
func (h *RateLimitHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "ID tidak valid")
		return
	}
	err = h.Service.Delete(id)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}
	response.Success(w, "Rate limit berhasil dihapus", nil)
}

// === CORS Config Handler ===

// CORSConfigHandler menangani HTTP request untuk CORS Config
type CORSConfigHandler struct {
	Service *service.CORSConfigService
}

// NewCORSConfigHandler membuat instance baru
func NewCORSConfigHandler(svc *service.CORSConfigService) *CORSConfigHandler {
	return &CORSConfigHandler{Service: svc}
}

// Create handler untuk membuat CORS config baru
func (h *CORSConfigHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req model.CreateCORSConfigRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Request body tidak valid")
		return
	}
	if req.VirtualDirectoryID <= 0 {
		response.BadRequest(w, "virtual_directory_id wajib diisi")
		return
	}
	item, err := h.Service.Create(&req)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}
	response.Created(w, "CORS config berhasil dibuat", item)
}

// GetByID handler untuk mengambil CORS config berdasarkan ID
func (h *CORSConfigHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "ID tidak valid")
		return
	}
	item, err := h.Service.GetByID(id)
	if err != nil {
		response.NotFound(w, "CORS config tidak ditemukan")
		return
	}
	response.Success(w, "CORS config berhasil diambil", item)
}

// GetAll handler untuk mengambil semua CORS configs
func (h *CORSConfigHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	page, limit := parsePagination(r)
	items, total, err := h.Service.GetAll(page, limit)
	if err != nil {
		response.InternalServerError(w, "Gagal mengambil daftar CORS config")
		return
	}
	result := map[string]interface{}{
		"cors_configs": items,
		"pagination":   paginationData(page, limit, total),
	}
	response.Success(w, "Daftar CORS config berhasil diambil", result)
}

// Update handler untuk memperbarui CORS config
func (h *CORSConfigHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "ID tidak valid")
		return
	}
	var req model.UpdateCORSConfigRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Request body tidak valid")
		return
	}
	err = h.Service.Update(id, &req)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}
	response.Success(w, "CORS config berhasil diupdate", nil)
}

// Delete handler untuk menghapus CORS config
func (h *CORSConfigHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "ID tidak valid")
		return
	}
	err = h.Service.Delete(id)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}
	response.Success(w, "CORS config berhasil dihapus", nil)
}

// === Circuit Breaker Handler ===

// CircuitBreakerHandler menangani HTTP request untuk Circuit Breaker
type CircuitBreakerHandler struct {
	Service *service.CircuitBreakerService
}

// NewCircuitBreakerHandler membuat instance baru
func NewCircuitBreakerHandler(svc *service.CircuitBreakerService) *CircuitBreakerHandler {
	return &CircuitBreakerHandler{Service: svc}
}

// Create handler untuk membuat circuit breaker baru
func (h *CircuitBreakerHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req model.CreateCircuitBreakerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Request body tidak valid")
		return
	}
	if req.VirtualDirectoryID <= 0 {
		response.BadRequest(w, "virtual_directory_id wajib diisi")
		return
	}
	item, err := h.Service.Create(&req)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}
	response.Created(w, "Circuit breaker berhasil dibuat", item)
}

// GetByID handler untuk mengambil circuit breaker berdasarkan ID
func (h *CircuitBreakerHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "ID tidak valid")
		return
	}
	item, err := h.Service.GetByID(id)
	if err != nil {
		response.NotFound(w, "Circuit breaker tidak ditemukan")
		return
	}
	response.Success(w, "Circuit breaker berhasil diambil", item)
}

// GetAll handler untuk mengambil semua circuit breakers
func (h *CircuitBreakerHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	page, limit := parsePagination(r)
	items, total, err := h.Service.GetAll(page, limit)
	if err != nil {
		response.InternalServerError(w, "Gagal mengambil daftar circuit breaker")
		return
	}
	result := map[string]interface{}{
		"circuit_breakers": items,
		"pagination":       paginationData(page, limit, total),
	}
	response.Success(w, "Daftar circuit breaker berhasil diambil", result)
}

// Update handler untuk memperbarui circuit breaker
func (h *CircuitBreakerHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "ID tidak valid")
		return
	}
	var req model.UpdateCircuitBreakerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Request body tidak valid")
		return
	}
	err = h.Service.Update(id, &req)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}
	response.Success(w, "Circuit breaker berhasil diupdate", nil)
}

// Delete handler untuk menghapus circuit breaker
func (h *CircuitBreakerHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "ID tidak valid")
		return
	}
	err = h.Service.Delete(id)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}
	response.Success(w, "Circuit breaker berhasil dihapus", nil)
}

// === IP Whitelist Handler ===

// IPWhitelistHandler menangani HTTP request untuk IP Whitelist
type IPWhitelistHandler struct {
	Service *service.IPWhitelistService
}

// NewIPWhitelistHandler membuat instance baru
func NewIPWhitelistHandler(svc *service.IPWhitelistService) *IPWhitelistHandler {
	return &IPWhitelistHandler{Service: svc}
}

// Create handler untuk menambahkan IP ke whitelist
func (h *IPWhitelistHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req model.CreateIPWhitelistRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Request body tidak valid")
		return
	}
	if req.VirtualDirectoryID <= 0 {
		response.BadRequest(w, "virtual_directory_id wajib diisi")
		return
	}
	item, err := h.Service.Create(&req)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}
	response.Created(w, "IP whitelist berhasil ditambahkan", item)
}

// GetByID handler untuk mengambil IP whitelist berdasarkan ID
func (h *IPWhitelistHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "ID tidak valid")
		return
	}
	item, err := h.Service.GetByID(id)
	if err != nil {
		response.NotFound(w, "IP whitelist tidak ditemukan")
		return
	}
	response.Success(w, "IP whitelist berhasil diambil", item)
}

// GetAll handler untuk mengambil semua IP whitelist
func (h *IPWhitelistHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	page, limit := parsePagination(r)
	items, total, err := h.Service.GetAll(page, limit)
	if err != nil {
		response.InternalServerError(w, "Gagal mengambil daftar IP whitelist")
		return
	}
	result := map[string]interface{}{
		"ip_whitelists": items,
		"pagination":    paginationData(page, limit, total),
	}
	response.Success(w, "Daftar IP whitelist berhasil diambil", result)
}

// GetByDirectory handler untuk mengambil IP whitelist per directory
func (h *IPWhitelistHandler) GetByDirectory(w http.ResponseWriter, r *http.Request) {
	dirID, err := strconv.ParseInt(r.PathValue("dir_id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "dir_id tidak valid")
		return
	}
	items, err := h.Service.GetByDirectoryID(dirID)
	if err != nil {
		response.InternalServerError(w, "Gagal mengambil IP whitelist")
		return
	}
	response.Success(w, "IP whitelist berhasil diambil", items)
}

// Update handler untuk memperbarui IP whitelist
func (h *IPWhitelistHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "ID tidak valid")
		return
	}
	var req model.UpdateIPWhitelistRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Request body tidak valid")
		return
	}
	err = h.Service.Update(id, &req)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}
	response.Success(w, "IP whitelist berhasil diupdate", nil)
}

// Delete handler untuk menghapus IP whitelist
func (h *IPWhitelistHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "ID tidak valid")
		return
	}
	err = h.Service.Delete(id)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}
	response.Success(w, "IP whitelist berhasil dihapus", nil)
}

// === IP Blacklist Handler ===

// IPBlacklistHandler menangani HTTP request untuk IP Blacklist
type IPBlacklistHandler struct {
	Service *service.IPBlacklistService
}

// NewIPBlacklistHandler membuat instance baru
func NewIPBlacklistHandler(svc *service.IPBlacklistService) *IPBlacklistHandler {
	return &IPBlacklistHandler{Service: svc}
}

// Create handler untuk menambahkan IP ke blacklist
func (h *IPBlacklistHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req model.CreateIPBlacklistRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Request body tidak valid")
		return
	}
	if req.VirtualDirectoryID <= 0 {
		response.BadRequest(w, "virtual_directory_id wajib diisi")
		return
	}
	item, err := h.Service.Create(&req)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}
	response.Created(w, "IP blacklist berhasil ditambahkan", item)
}

// GetByID handler untuk mengambil IP blacklist berdasarkan ID
func (h *IPBlacklistHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "ID tidak valid")
		return
	}
	item, err := h.Service.GetByID(id)
	if err != nil {
		response.NotFound(w, "IP blacklist tidak ditemukan")
		return
	}
	response.Success(w, "IP blacklist berhasil diambil", item)
}

// GetAll handler untuk mengambil semua IP blacklist
func (h *IPBlacklistHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	page, limit := parsePagination(r)
	items, total, err := h.Service.GetAll(page, limit)
	if err != nil {
		response.InternalServerError(w, "Gagal mengambil daftar IP blacklist")
		return
	}
	result := map[string]interface{}{
		"ip_blacklists": items,
		"pagination":    paginationData(page, limit, total),
	}
	response.Success(w, "Daftar IP blacklist berhasil diambil", result)
}

// GetByDirectory handler untuk mengambil IP blacklist per directory
func (h *IPBlacklistHandler) GetByDirectory(w http.ResponseWriter, r *http.Request) {
	dirID, err := strconv.ParseInt(r.PathValue("dir_id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "dir_id tidak valid")
		return
	}
	items, err := h.Service.GetByDirectoryID(dirID)
	if err != nil {
		response.InternalServerError(w, "Gagal mengambil IP blacklist")
		return
	}
	response.Success(w, "IP blacklist berhasil diambil", items)
}

// Update handler untuk memperbarui IP blacklist
func (h *IPBlacklistHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "ID tidak valid")
		return
	}
	var req model.UpdateIPBlacklistRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Request body tidak valid")
		return
	}
	err = h.Service.Update(id, &req)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}
	response.Success(w, "IP blacklist berhasil diupdate", nil)
}

// Delete handler untuk menghapus IP blacklist
func (h *IPBlacklistHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "ID tidak valid")
		return
	}
	err = h.Service.Delete(id)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}
	response.Success(w, "IP blacklist berhasil dihapus", nil)
}
