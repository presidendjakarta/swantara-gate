package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/presidendjakarta/swantara-gate/internal/model"
	"github.com/presidendjakarta/swantara-gate/internal/response"
	"github.com/presidendjakarta/swantara-gate/internal/service"
)

// UpstreamServerHandler menangani HTTP request untuk Upstream Server
type UpstreamServerHandler struct {
	Service *service.UpstreamServerService
}

// NewUpstreamServerHandler membuat instance baru UpstreamServerHandler
func NewUpstreamServerHandler(svc *service.UpstreamServerService) *UpstreamServerHandler {
	return &UpstreamServerHandler{Service: svc}
}

// CreateUpstreamServer handler untuk membuat upstream server baru
func (h *UpstreamServerHandler) CreateUpstreamServer(w http.ResponseWriter, r *http.Request) {
	var req model.CreateUpstreamServerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Request body tidak valid")
		return
	}

	if req.VirtualHostID <= 0 {
		response.BadRequest(w, "virtual_host_id wajib diisi")
		return
	}

	server, err := h.Service.CreateUpstreamServer(&req)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	response.Created(w, "Upstream server berhasil dibuat", server)
}

// GetUpstreamServerByID handler untuk mengambil upstream server berdasarkan ID
func (h *UpstreamServerHandler) GetUpstreamServerByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "ID tidak valid")
		return
	}

	server, err := h.Service.GetUpstreamServerByID(id)
	if err != nil {
		response.NotFound(w, "Upstream server tidak ditemukan")
		return
	}

	response.Success(w, "Upstream server berhasil diambil", server)
}

// GetAllUpstreamServers handler untuk mengambil semua upstream servers
func (h *UpstreamServerHandler) GetAllUpstreamServers(w http.ResponseWriter, r *http.Request) {
	page, limit := parsePagination(r)

	servers, total, err := h.Service.GetAllUpstreamServers(page, limit)
	if err != nil {
		response.InternalServerError(w, "Gagal mengambil daftar upstream server")
		return
	}

	result := map[string]interface{}{
		"upstream_servers": servers,
		"pagination":       paginationData(page, limit, total),
	}

	response.Success(w, "Daftar upstream server berhasil diambil", result)
}

// GetUpstreamServersByVHost handler untuk mengambil upstream servers per virtual host
func (h *UpstreamServerHandler) GetUpstreamServersByVHost(w http.ResponseWriter, r *http.Request) {
	vhostID, err := strconv.ParseInt(r.PathValue("vhost_id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "vhost_id tidak valid")
		return
	}

	servers, err := h.Service.GetUpstreamServersByVHostID(vhostID)
	if err != nil {
		response.InternalServerError(w, "Gagal mengambil upstream servers")
		return
	}

	response.Success(w, "Upstream servers berhasil diambil", servers)
}

// UpdateUpstreamServer handler untuk memperbarui upstream server
func (h *UpstreamServerHandler) UpdateUpstreamServer(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "ID tidak valid")
		return
	}

	var req model.UpdateUpstreamServerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Request body tidak valid")
		return
	}

	err = h.Service.UpdateUpstreamServer(id, &req)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	response.Success(w, "Upstream server berhasil diupdate", nil)
}

// DeleteUpstreamServer handler untuk menghapus upstream server
func (h *UpstreamServerHandler) DeleteUpstreamServer(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "ID tidak valid")
		return
	}

	err = h.Service.DeleteUpstreamServer(id)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	response.Success(w, "Upstream server berhasil dihapus", nil)
}

// === Helper Functions ===

// parsePagination mengambil parameter pagination dari query string
func parsePagination(r *http.Request) (int, int) {
	page := 1
	limit := 10

	if p, err := strconv.Atoi(r.URL.Query().Get("page")); err == nil && p > 0 {
		page = p
	}
	if l, err := strconv.Atoi(r.URL.Query().Get("limit")); err == nil && l > 0 {
		limit = l
	}

	return page, limit
}

// paginationData membuat data pagination untuk response
func paginationData(page, limit int, total int64) map[string]interface{} {
	return map[string]interface{}{
		"page":  page,
		"limit": limit,
		"total": total,
	}
}
