package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/presidendjakarta/swantara-gate/internal/model"
	"github.com/presidendjakarta/swantara-gate/internal/response"
	"github.com/presidendjakarta/swantara-gate/internal/service"
)

// =========================================================
// REQUEST HEADER RULE HANDLER
// =========================================================

// RequestHeaderRuleHandler menangani HTTP request untuk Request Header Rules
type RequestHeaderRuleHandler struct {
	Service *service.RequestHeaderRuleService
}

// NewRequestHeaderRuleHandler membuat instance baru
func NewRequestHeaderRuleHandler(svc *service.RequestHeaderRuleService) *RequestHeaderRuleHandler {
	return &RequestHeaderRuleHandler{Service: svc}
}

// Create handler untuk membuat request header rule baru
func (h *RequestHeaderRuleHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req model.CreateRequestHeaderRuleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Request body tidak valid")
		return
	}
	item, err := h.Service.Create(&req)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}
	response.Created(w, "Request header rule berhasil dibuat", item)
}

// GetByID handler untuk mengambil request header rule berdasarkan ID
func (h *RequestHeaderRuleHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "ID tidak valid")
		return
	}
	item, err := h.Service.GetByID(id)
	if err != nil {
		response.NotFound(w, "Request header rule tidak ditemukan")
		return
	}
	response.Success(w, "Request header rule berhasil diambil", item)
}

// GetAll handler untuk mengambil semua request header rules
func (h *RequestHeaderRuleHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	page, limit := parsePagination(r)
	items, total, err := h.Service.GetAll(page, limit)
	if err != nil {
		response.InternalServerError(w, "Gagal mengambil daftar request header rules")
		return
	}
	result := map[string]interface{}{
		"request_header_rules": items,
		"pagination":           paginationData(page, limit, total),
	}
	response.Success(w, "Daftar request header rules berhasil diambil", result)
}

// GetByDirectory handler untuk mengambil rules berdasarkan directory
func (h *RequestHeaderRuleHandler) GetByDirectory(w http.ResponseWriter, r *http.Request) {
	dirID, err := strconv.ParseInt(r.PathValue("dir_id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "Directory ID tidak valid")
		return
	}
	items, err := h.Service.GetByDirectoryID(dirID)
	if err != nil {
		response.InternalServerError(w, "Gagal mengambil request header rules")
		return
	}
	response.Success(w, "Request header rules berhasil diambil", items)
}

// Update handler untuk memperbarui request header rule
func (h *RequestHeaderRuleHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "ID tidak valid")
		return
	}
	var req model.UpdateRequestHeaderRuleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Request body tidak valid")
		return
	}
	if err := h.Service.Update(id, &req); err != nil {
		response.BadRequest(w, err.Error())
		return
	}
	response.Success(w, "Request header rule berhasil diperbarui", nil)
}

// Delete handler untuk menghapus request header rule
func (h *RequestHeaderRuleHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "ID tidak valid")
		return
	}
	if err := h.Service.Delete(id); err != nil {
		response.BadRequest(w, err.Error())
		return
	}
	response.Success(w, "Request header rule berhasil dihapus", nil)
}

// =========================================================
// RESPONSE HEADER RULE HANDLER
// =========================================================

// ResponseHeaderRuleHandler menangani HTTP request untuk Response Header Rules
type ResponseHeaderRuleHandler struct {
	Service *service.ResponseHeaderRuleService
}

// NewResponseHeaderRuleHandler membuat instance baru
func NewResponseHeaderRuleHandler(svc *service.ResponseHeaderRuleService) *ResponseHeaderRuleHandler {
	return &ResponseHeaderRuleHandler{Service: svc}
}

// Create handler untuk membuat response header rule baru
func (h *ResponseHeaderRuleHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req model.CreateResponseHeaderRuleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Request body tidak valid")
		return
	}
	item, err := h.Service.Create(&req)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}
	response.Created(w, "Response header rule berhasil dibuat", item)
}

// GetByID handler untuk mengambil response header rule berdasarkan ID
func (h *ResponseHeaderRuleHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "ID tidak valid")
		return
	}
	item, err := h.Service.GetByID(id)
	if err != nil {
		response.NotFound(w, "Response header rule tidak ditemukan")
		return
	}
	response.Success(w, "Response header rule berhasil diambil", item)
}

// GetAll handler untuk mengambil semua response header rules
func (h *ResponseHeaderRuleHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	page, limit := parsePagination(r)
	items, total, err := h.Service.GetAll(page, limit)
	if err != nil {
		response.InternalServerError(w, "Gagal mengambil daftar response header rules")
		return
	}
	result := map[string]interface{}{
		"response_header_rules": items,
		"pagination":            paginationData(page, limit, total),
	}
	response.Success(w, "Daftar response header rules berhasil diambil", result)
}

// GetByDirectory handler untuk mengambil rules berdasarkan directory
func (h *ResponseHeaderRuleHandler) GetByDirectory(w http.ResponseWriter, r *http.Request) {
	dirID, err := strconv.ParseInt(r.PathValue("dir_id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "Directory ID tidak valid")
		return
	}
	items, err := h.Service.GetByDirectoryID(dirID)
	if err != nil {
		response.InternalServerError(w, "Gagal mengambil response header rules")
		return
	}
	response.Success(w, "Response header rules berhasil diambil", items)
}

// Update handler untuk memperbarui response header rule
func (h *ResponseHeaderRuleHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "ID tidak valid")
		return
	}
	var req model.UpdateResponseHeaderRuleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Request body tidak valid")
		return
	}
	if err := h.Service.Update(id, &req); err != nil {
		response.BadRequest(w, err.Error())
		return
	}
	response.Success(w, "Response header rule berhasil diperbarui", nil)
}

// Delete handler untuk menghapus response header rule
func (h *ResponseHeaderRuleHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "ID tidak valid")
		return
	}
	if err := h.Service.Delete(id); err != nil {
		response.BadRequest(w, err.Error())
		return
	}
	response.Success(w, "Response header rule berhasil dihapus", nil)
}

// =========================================================
// QUERY REWRITE HANDLER
// =========================================================

// QueryRewriteHandler menangani HTTP request untuk Query Rewrites
type QueryRewriteHandler struct {
	Service *service.QueryRewriteService
}

// NewQueryRewriteHandler membuat instance baru
func NewQueryRewriteHandler(svc *service.QueryRewriteService) *QueryRewriteHandler {
	return &QueryRewriteHandler{Service: svc}
}

// Create handler untuk membuat query rewrite baru
func (h *QueryRewriteHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req model.CreateQueryRewriteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Request body tidak valid")
		return
	}
	item, err := h.Service.Create(&req)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}
	response.Created(w, "Query rewrite berhasil dibuat", item)
}

// GetByID handler untuk mengambil query rewrite berdasarkan ID
func (h *QueryRewriteHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "ID tidak valid")
		return
	}
	item, err := h.Service.GetByID(id)
	if err != nil {
		response.NotFound(w, "Query rewrite tidak ditemukan")
		return
	}
	response.Success(w, "Query rewrite berhasil diambil", item)
}

// GetAll handler untuk mengambil semua query rewrites
func (h *QueryRewriteHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	page, limit := parsePagination(r)
	items, total, err := h.Service.GetAll(page, limit)
	if err != nil {
		response.InternalServerError(w, "Gagal mengambil daftar query rewrites")
		return
	}
	result := map[string]interface{}{
		"query_rewrites": items,
		"pagination":     paginationData(page, limit, total),
	}
	response.Success(w, "Daftar query rewrites berhasil diambil", result)
}

// GetByDirectory handler untuk mengambil rewrites berdasarkan directory
func (h *QueryRewriteHandler) GetByDirectory(w http.ResponseWriter, r *http.Request) {
	dirID, err := strconv.ParseInt(r.PathValue("dir_id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "Directory ID tidak valid")
		return
	}
	items, err := h.Service.GetByDirectoryID(dirID)
	if err != nil {
		response.InternalServerError(w, "Gagal mengambil query rewrites")
		return
	}
	response.Success(w, "Query rewrites berhasil diambil", items)
}

// Update handler untuk memperbarui query rewrite
func (h *QueryRewriteHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "ID tidak valid")
		return
	}
	var req model.UpdateQueryRewriteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Request body tidak valid")
		return
	}
	if err := h.Service.Update(id, &req); err != nil {
		response.BadRequest(w, err.Error())
		return
	}
	response.Success(w, "Query rewrite berhasil diperbarui", nil)
}

// Delete handler untuk menghapus query rewrite
func (h *QueryRewriteHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "ID tidak valid")
		return
	}
	if err := h.Service.Delete(id); err != nil {
		response.BadRequest(w, err.Error())
		return
	}
	response.Success(w, "Query rewrite berhasil dihapus", nil)
}
