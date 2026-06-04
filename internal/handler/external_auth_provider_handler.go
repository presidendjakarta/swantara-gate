package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/presidendjakarta/swantara-gate/internal/model"
	"github.com/presidendjakarta/swantara-gate/internal/response"
	"github.com/presidendjakarta/swantara-gate/internal/service"
)

// ExternalAuthProviderHandler menangani HTTP request untuk external auth providers
type ExternalAuthProviderHandler struct {
	Service *service.ExternalAuthProviderService
}

// NewExternalAuthProviderHandler membuat instance baru
func NewExternalAuthProviderHandler(svc *service.ExternalAuthProviderService) *ExternalAuthProviderHandler {
	return &ExternalAuthProviderHandler{Service: svc}
}

// Create handler untuk membuat provider baru
func (h *ExternalAuthProviderHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req model.CreateExternalAuthProviderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Request body tidak valid")
		return
	}

	provider, err := h.Service.Create(&req)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	response.Created(w, "External auth provider berhasil dibuat", provider)
}

// GetByID handler untuk mengambil provider berdasarkan ID
func (h *ExternalAuthProviderHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "ID tidak valid")
		return
	}

	provider, err := h.Service.GetByID(id)
	if err != nil {
		response.NotFound(w, "External auth provider tidak ditemukan")
		return
	}

	response.Success(w, "External auth provider berhasil diambil", provider)
}

// GetAll handler untuk mengambil semua provider
func (h *ExternalAuthProviderHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	page, limit := parsePagination(r)
	search := r.URL.Query().Get("search")

	providers, total, err := h.Service.GetAll(page, limit, search)
	if err != nil {
		response.InternalServerError(w, "Gagal mengambil daftar external auth providers")
		return
	}

	result := map[string]interface{}{
		"providers":  providers,
		"pagination": paginationData(page, limit, total),
	}

	response.Success(w, "Daftar external auth provider berhasil diambil", result)
}

// Update handler untuk mengupdate provider
func (h *ExternalAuthProviderHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "ID tidak valid")
		return
	}

	var req model.UpdateExternalAuthProviderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Request body tidak valid")
		return
	}

	err = h.Service.Update(id, &req)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	response.Success(w, "External auth provider berhasil diupdate", nil)
}

// Delete handler untuk menghapus provider
func (h *ExternalAuthProviderHandler) Delete(w http.ResponseWriter, r *http.Request) {
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

	response.Success(w, "External auth provider berhasil dihapus", nil)
}

// AssignToVirtualDirectory handler untuk assign provider ke virtual directory
func (h *ExternalAuthProviderHandler) AssignToVirtualDirectory(w http.ResponseWriter, r *http.Request) {
	providerID, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "Provider ID tidak valid")
		return
	}

	var req struct {
		VirtualDirectoryID int64 `json:"virtual_directory_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Request body tidak valid")
		return
	}

	if req.VirtualDirectoryID <= 0 {
		response.BadRequest(w, "virtual_directory_id wajib diisi")
		return
	}

	err = h.Service.AssignProviderToVirtualDirectory(req.VirtualDirectoryID, providerID)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	response.Success(w, "Provider berhasil di-assign ke virtual directory", nil)
}

// RemoveFromVirtualDirectory handler untuk remove provider dari virtual directory
func (h *ExternalAuthProviderHandler) RemoveFromVirtualDirectory(w http.ResponseWriter, r *http.Request) {
	providerID, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "Provider ID tidak valid")
		return
	}

	vdirID, err := strconv.ParseInt(r.PathValue("vdir_id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "Virtual directory ID tidak valid")
		return
	}

	err = h.Service.RemoveProviderFromVirtualDirectory(vdirID, providerID)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	response.Success(w, "Provider berhasil di-remove dari virtual directory", nil)
}

// GetByVirtualDirectory handler untuk mengambil semua provider untuk virtual directory
func (h *ExternalAuthProviderHandler) GetByVirtualDirectory(w http.ResponseWriter, r *http.Request) {
	vdirID, err := strconv.ParseInt(r.PathValue("vdir_id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "Virtual directory ID tidak valid")
		return
	}

	providers, err := h.Service.GetProvidersByVirtualDirectory(vdirID)
	if err != nil {
		response.InternalServerError(w, "Gagal mengambil providers untuk virtual directory")
		return
	}

	result := map[string]interface{}{
		"providers": providers,
	}

	response.Success(w, "Daftar provider untuk virtual directory berhasil diambil", result)
}

// GetMappingsByVirtualDirectory handler untuk mengambil mapping untuk virtual directory
func (h *ExternalAuthProviderHandler) GetMappingsByVirtualDirectory(w http.ResponseWriter, r *http.Request) {
	vdirID, err := strconv.ParseInt(r.PathValue("vdir_id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "Virtual directory ID tidak valid")
		return
	}

	mappings, err := h.Service.GetMappingsByVirtualDirectory(vdirID)
	if err != nil {
		response.InternalServerError(w, "Gagal mengambil mappings")
		return
	}

	result := map[string]interface{}{
		"mappings": mappings,
	}

	response.Success(w, "Mappings berhasil diambil", result)
}
