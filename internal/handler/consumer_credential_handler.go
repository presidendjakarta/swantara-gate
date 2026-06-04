package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/presidendjakarta/swantara-gate/internal/model"
	"github.com/presidendjakarta/swantara-gate/internal/response"
	"github.com/presidendjakarta/swantara-gate/internal/service"
)

// ConsumerCredentialHandler menangani HTTP request untuk Consumer Credentials
type ConsumerCredentialHandler struct {
	Service *service.ConsumerCredentialService
}

// NewConsumerCredentialHandler membuat instance baru
func NewConsumerCredentialHandler(svc *service.ConsumerCredentialService) *ConsumerCredentialHandler {
	return &ConsumerCredentialHandler{Service: svc}
}

// CreateCredential handler untuk membuat credential baru
func (h *ConsumerCredentialHandler) CreateCredential(w http.ResponseWriter, r *http.Request) {
	var req model.CreateConsumerCredentialRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Request body tidak valid")
		return
	}

	if req.ConsumerID <= 0 {
		response.BadRequest(w, "consumer_id wajib diisi")
		return
	}

	cred, err := h.Service.CreateCredential(&req)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	response.Created(w, "Consumer credential berhasil dibuat", cred)
}

// GetCredentialByID handler untuk mengambil credential berdasarkan ID
func (h *ConsumerCredentialHandler) GetCredentialByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "ID tidak valid")
		return
	}

	cred, err := h.Service.GetCredentialByID(id)
	if err != nil {
		response.NotFound(w, "Credential tidak ditemukan")
		return
	}

	response.Success(w, "Credential berhasil diambil", cred)
}

// GetAllCredentials handler untuk mengambil semua credentials
func (h *ConsumerCredentialHandler) GetAllCredentials(w http.ResponseWriter, r *http.Request) {
	page, limit := parsePagination(r)
	search := r.URL.Query().Get("search")

	creds, total, err := h.Service.GetAllCredentials(page, limit, search)
	if err != nil {
		response.InternalServerError(w, "Gagal mengambil daftar credential")
		return
	}

	result := map[string]interface{}{
		"credentials": creds,
		"pagination":  paginationData(page, limit, total),
	}

	response.Success(w, "Daftar credential berhasil diambil", result)
}

// GetCredentialsByConsumer handler untuk mengambil credentials per consumer
func (h *ConsumerCredentialHandler) GetCredentialsByConsumer(w http.ResponseWriter, r *http.Request) {
	consumerID, err := strconv.ParseInt(r.PathValue("consumer_id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "consumer_id tidak valid")
		return
	}

	creds, err := h.Service.GetCredentialsByConsumerID(consumerID)
	if err != nil {
		response.InternalServerError(w, "Gagal mengambil credentials")
		return
	}

	response.Success(w, "Credentials berhasil diambil", creds)
}

// UpdateCredential handler untuk memperbarui credential
func (h *ConsumerCredentialHandler) UpdateCredential(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "ID tidak valid")
		return
	}

	var req model.UpdateConsumerCredentialRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Request body tidak valid")
		return
	}

	err = h.Service.UpdateCredential(id, &req)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	response.Success(w, "Credential berhasil diupdate", nil)
}

// DeleteCredential handler untuk menghapus credential
func (h *ConsumerCredentialHandler) DeleteCredential(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "ID tidak valid")
		return
	}

	err = h.Service.DeleteCredential(id)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	response.Success(w, "Credential berhasil dihapus", nil)
}

// === API Key Handler ===

// APIKeyHandler menangani HTTP request untuk API Keys
type APIKeyHandler struct {
	Service *service.APIKeyService
}

// NewAPIKeyHandler membuat instance baru APIKeyHandler
func NewAPIKeyHandler(svc *service.APIKeyService) *APIKeyHandler {
	return &APIKeyHandler{Service: svc}
}

// CreateAPIKey handler untuk membuat API key baru
func (h *APIKeyHandler) CreateAPIKey(w http.ResponseWriter, r *http.Request) {
	var req model.CreateAPIKeyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Request body tidak valid")
		return
	}

	if req.ConsumerID <= 0 {
		response.BadRequest(w, "consumer_id wajib diisi")
		return
	}

	key, err := h.Service.CreateAPIKey(&req)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	response.Created(w, "API key berhasil dibuat", key)
}

// GetAPIKeyByID handler untuk mengambil API key berdasarkan ID
func (h *APIKeyHandler) GetAPIKeyByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "ID tidak valid")
		return
	}

	key, err := h.Service.GetAPIKeyByID(id)
	if err != nil {
		response.NotFound(w, "API key tidak ditemukan")
		return
	}

	response.Success(w, "API key berhasil diambil", key)
}

// GetAllAPIKeys handler untuk mengambil semua API keys
func (h *APIKeyHandler) GetAllAPIKeys(w http.ResponseWriter, r *http.Request) {
	page, limit := parsePagination(r)
	search := r.URL.Query().Get("search")

	keys, total, err := h.Service.GetAllAPIKeys(page, limit, search)
	if err != nil {
		response.InternalServerError(w, "Gagal mengambil daftar API keys")
		return
	}

	result := map[string]interface{}{
		"api_keys":   keys,
		"pagination": paginationData(page, limit, total),
	}

	response.Success(w, "Daftar API keys berhasil diambil", result)
}

// GetAPIKeysByConsumer handler untuk mengambil API keys per consumer
func (h *APIKeyHandler) GetAPIKeysByConsumer(w http.ResponseWriter, r *http.Request) {
	consumerID, err := strconv.ParseInt(r.PathValue("consumer_id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "consumer_id tidak valid")
		return
	}

	keys, err := h.Service.GetAPIKeysByConsumerID(consumerID)
	if err != nil {
		response.InternalServerError(w, "Gagal mengambil API keys")
		return
	}

	response.Success(w, "API keys berhasil diambil", keys)
}

// UpdateAPIKey handler untuk memperbarui API key
func (h *APIKeyHandler) UpdateAPIKey(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "ID tidak valid")
		return
	}

	var req model.UpdateAPIKeyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Request body tidak valid")
		return
	}

	err = h.Service.UpdateAPIKey(id, &req)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	response.Success(w, "API key berhasil diupdate", nil)
}

// DeleteAPIKey handler untuk menghapus API key
func (h *APIKeyHandler) DeleteAPIKey(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "ID tidak valid")
		return
	}

	err = h.Service.DeleteAPIKey(id)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	response.Success(w, "API key berhasil dihapus", nil)
}

// === Route Consumer Access Handler ===

// RouteConsumerAccessHandler menangani HTTP request untuk ACL
type RouteConsumerAccessHandler struct {
	Service *service.RouteConsumerAccessService
}

// NewRouteConsumerAccessHandler membuat instance baru
func NewRouteConsumerAccessHandler(svc *service.RouteConsumerAccessService) *RouteConsumerAccessHandler {
	return &RouteConsumerAccessHandler{Service: svc}
}

// CreateAccess handler untuk membuat akses baru
func (h *RouteConsumerAccessHandler) CreateAccess(w http.ResponseWriter, r *http.Request) {
	var req model.CreateRouteConsumerAccessRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Request body tidak valid")
		return
	}

	access, err := h.Service.CreateAccess(&req)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	response.Created(w, "Route consumer access berhasil dibuat", access)
}

// GetAccessByID handler untuk mengambil access berdasarkan ID
func (h *RouteConsumerAccessHandler) GetAccessByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "ID tidak valid")
		return
	}

	access, err := h.Service.GetAccessByID(id)
	if err != nil {
		response.NotFound(w, "Access tidak ditemukan")
		return
	}

	response.Success(w, "Access berhasil diambil", access)
}

// GetAllAccess handler untuk mengambil semua access
func (h *RouteConsumerAccessHandler) GetAllAccess(w http.ResponseWriter, r *http.Request) {
	page, limit := parsePagination(r)

	accesses, total, err := h.Service.GetAllAccess(page, limit)
	if err != nil {
		response.InternalServerError(w, "Gagal mengambil daftar access")
		return
	}

	result := map[string]interface{}{
		"route_consumer_access": accesses,
		"pagination":            paginationData(page, limit, total),
	}

	response.Success(w, "Daftar access berhasil diambil", result)
}

// GetAccessByDirectory handler untuk mengambil access per directory
func (h *RouteConsumerAccessHandler) GetAccessByDirectory(w http.ResponseWriter, r *http.Request) {
	dirID, err := strconv.ParseInt(r.PathValue("dir_id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "dir_id tidak valid")
		return
	}

	accesses, err := h.Service.GetAccessByDirectoryID(dirID)
	if err != nil {
		response.InternalServerError(w, "Gagal mengambil access")
		return
	}

	response.Success(w, "Access berhasil diambil", accesses)
}

// UpdateAccess handler untuk memperbarui access
func (h *RouteConsumerAccessHandler) UpdateAccess(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "ID tidak valid")
		return
	}

	var req model.UpdateRouteConsumerAccessRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Request body tidak valid")
		return
	}

	err = h.Service.UpdateAccess(id, &req)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	response.Success(w, "Access berhasil diupdate", nil)
}

// DeleteAccess handler untuk menghapus access
func (h *RouteConsumerAccessHandler) DeleteAccess(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "ID tidak valid")
		return
	}

	err = h.Service.DeleteAccess(id)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	response.Success(w, "Access berhasil dihapus", nil)
}
