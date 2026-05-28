package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/presidendjakarta/swantara-gate/internal/model"
	"github.com/presidendjakarta/swantara-gate/internal/response"
	"github.com/presidendjakarta/swantara-gate/internal/service"
)

// VirtualDirectoryHandler menangani HTTP request untuk Virtual Directory
type VirtualDirectoryHandler struct {
	Service *service.VirtualDirectoryService
}

// NewVirtualDirectoryHandler membuat instance baru VirtualDirectoryHandler
func NewVirtualDirectoryHandler(svc *service.VirtualDirectoryService) *VirtualDirectoryHandler {
	return &VirtualDirectoryHandler{Service: svc}
}

// CreateVirtualDirectory handler untuk membuat virtual directory baru
func (h *VirtualDirectoryHandler) CreateVirtualDirectory(w http.ResponseWriter, r *http.Request) {
	var req model.CreateVirtualDirectoryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Request body tidak valid")
		return
	}

	if req.VirtualHostID <= 0 {
		response.BadRequest(w, "virtual_host_id wajib diisi")
		return
	}

	dir, err := h.Service.CreateVirtualDirectory(&req)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	response.Created(w, "Virtual directory berhasil dibuat", dir)
}

// GetVirtualDirectoryByID handler untuk mengambil virtual directory berdasarkan ID
func (h *VirtualDirectoryHandler) GetVirtualDirectoryByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "ID tidak valid")
		return
	}

	dir, err := h.Service.GetVirtualDirectoryByID(id)
	if err != nil {
		response.NotFound(w, "Virtual directory tidak ditemukan")
		return
	}

	response.Success(w, "Virtual directory berhasil diambil", dir)
}

// GetAllVirtualDirectories handler untuk mengambil semua virtual directories
func (h *VirtualDirectoryHandler) GetAllVirtualDirectories(w http.ResponseWriter, r *http.Request) {
	page, limit := parsePagination(r)

	dirs, total, err := h.Service.GetAllVirtualDirectories(page, limit)
	if err != nil {
		response.InternalServerError(w, "Gagal mengambil daftar virtual directory")
		return
	}

	result := map[string]interface{}{
		"virtual_directories": dirs,
		"pagination":          paginationData(page, limit, total),
	}

	response.Success(w, "Daftar virtual directory berhasil diambil", result)
}

// GetVirtualDirectoriesByVHost handler untuk mengambil directories per virtual host
func (h *VirtualDirectoryHandler) GetVirtualDirectoriesByVHost(w http.ResponseWriter, r *http.Request) {
	vhostID, err := strconv.ParseInt(r.PathValue("vhost_id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "vhost_id tidak valid")
		return
	}

	dirs, err := h.Service.GetVirtualDirectoriesByVHostID(vhostID)
	if err != nil {
		response.InternalServerError(w, "Gagal mengambil virtual directories")
		return
	}

	response.Success(w, "Virtual directories berhasil diambil", dirs)
}

// UpdateVirtualDirectory handler untuk memperbarui virtual directory
func (h *VirtualDirectoryHandler) UpdateVirtualDirectory(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "ID tidak valid")
		return
	}

	var req model.UpdateVirtualDirectoryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Request body tidak valid")
		return
	}

	err = h.Service.UpdateVirtualDirectory(id, &req)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	response.Success(w, "Virtual directory berhasil diupdate", nil)
}

// DeleteVirtualDirectory handler untuk menghapus virtual directory
func (h *VirtualDirectoryHandler) DeleteVirtualDirectory(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "ID tidak valid")
		return
	}

	err = h.Service.DeleteVirtualDirectory(id)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	response.Success(w, "Virtual directory berhasil dihapus", nil)
}

// === Virtual Directory Methods Handlers ===

// GetMethods handler untuk mengambil methods dari directory
func (h *VirtualDirectoryHandler) GetMethods(w http.ResponseWriter, r *http.Request) {
	dirID, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "ID tidak valid")
		return
	}

	methods, err := h.Service.GetMethods(dirID)
	if err != nil {
		response.InternalServerError(w, "Gagal mengambil methods")
		return
	}

	response.Success(w, "Methods berhasil diambil", methods)
}

// SetMethods handler untuk mengatur methods pada directory
func (h *VirtualDirectoryHandler) SetMethods(w http.ResponseWriter, r *http.Request) {
	dirID, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "ID tidak valid")
		return
	}

	var req model.SetMethodsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Request body tidak valid")
		return
	}

	if len(req.Methods) == 0 {
		response.BadRequest(w, "Minimal satu HTTP method wajib diisi")
		return
	}

	methods, err := h.Service.SetMethods(dirID, req.Methods)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	response.Success(w, "Methods berhasil diatur", methods)
}
