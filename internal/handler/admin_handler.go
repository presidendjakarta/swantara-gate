package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/presidendjakarta/swantara-gate/internal/model"
	"github.com/presidendjakarta/swantara-gate/internal/response"
	"github.com/presidendjakarta/swantara-gate/internal/service"
)

// APIConsumerHandler menangani HTTP request untuk API Consumer
type APIConsumerHandler struct {
	ConsumerService *service.APIConsumerService
}

// NewAPIConsumerHandler membuat instance baru APIConsumerHandler
func NewAPIConsumerHandler(consumerService *service.APIConsumerService) *APIConsumerHandler {
	return &APIConsumerHandler{ConsumerService: consumerService}
}

// CreateConsumer handler untuk membuat consumer baru
func (h *APIConsumerHandler) CreateConsumer(w http.ResponseWriter, r *http.Request) {
	var req model.CreateAPIConsumerRequest
	
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Request body tidak valid")
		return
	}

	if req.ConsumerName == "" {
		response.BadRequest(w, "Consumer name wajib diisi")
		return
	}

	consumer, err := h.ConsumerService.CreateConsumer(&req)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	response.Created(w, "Consumer berhasil dibuat", consumer)
}

// GetConsumerByID handler untuk mengambil consumer berdasarkan ID
func (h *APIConsumerHandler) GetConsumerByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.BadRequest(w, "ID tidak valid")
		return
	}

	consumer, err := h.ConsumerService.GetConsumerByID(id)
	if err != nil {
		response.NotFound(w, "Consumer tidak ditemukan")
		return
	}

	response.Success(w, "Consumer berhasil diambil", consumer)
}

// GetAllConsumers handler untuk mengambil semua consumer
func (h *APIConsumerHandler) GetAllConsumers(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	page := 1
	limit := 10

	if pageStr != "" {
		p, err := strconv.Atoi(pageStr)
		if err == nil && p > 0 {
			page = p
		}
	}

	if limitStr != "" {
		l, err := strconv.Atoi(limitStr)
		if err == nil && l > 0 {
			limit = l
		}
	}

	consumers, total, err := h.ConsumerService.GetAllConsumers(page, limit)
	if err != nil {
		response.InternalServerError(w, "Gagal mengambil daftar consumer")
		return
	}

	result := map[string]interface{}{
		"consumers": consumers,
		"pagination": map[string]interface{}{
			"page":  page,
			"limit": limit,
			"total": total,
		},
	}

	response.Success(w, "Daftar consumer berhasil diambil", result)
}

// UpdateConsumer handler untuk memperbarui consumer
func (h *APIConsumerHandler) UpdateConsumer(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.BadRequest(w, "ID tidak valid")
		return
	}

	var req model.UpdateAPIConsumerRequest
	
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Request body tidak valid")
		return
	}

	err = h.ConsumerService.UpdateConsumer(id, &req)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	response.Success(w, "Consumer berhasil diupdate", nil)
}

// DeleteConsumer handler untuk menghapus consumer
func (h *APIConsumerHandler) DeleteConsumer(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.BadRequest(w, "ID tidak valid")
		return
	}

	err = h.ConsumerService.DeleteConsumer(id)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	response.Success(w, "Consumer berhasil dihapus", nil)
}

// HostHandler menangani HTTP request untuk Host
type HostHandler struct {
	HostService *service.HostService
}

// NewHostHandler membuat instance baru HostHandler
func NewHostHandler(hostService *service.HostService) *HostHandler {
	return &HostHandler{HostService: hostService}
}

// CreateHost handler untuk membuat host baru
func (h *HostHandler) CreateHost(w http.ResponseWriter, r *http.Request) {
	var req model.CreateHostRequest
	
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Request body tidak valid")
		return
	}

	if req.HostName == "" {
		response.BadRequest(w, "Host name wajib diisi")
		return
	}

	host, err := h.HostService.CreateHost(&req)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	response.Created(w, "Host berhasil dibuat", host)
}

// GetHostByID handler untuk mengambil host berdasarkan ID
func (h *HostHandler) GetHostByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.BadRequest(w, "ID tidak valid")
		return
	}

	host, err := h.HostService.GetHostByID(id)
	if err != nil {
		response.NotFound(w, "Host tidak ditemukan")
		return
	}

	response.Success(w, "Host berhasil diambil", host)
}

// GetAllHosts handler untuk mengambil semua host
func (h *HostHandler) GetAllHosts(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	page := 1
	limit := 10

	if pageStr != "" {
		p, err := strconv.Atoi(pageStr)
		if err == nil && p > 0 {
			page = p
		}
	}

	if limitStr != "" {
		l, err := strconv.Atoi(limitStr)
		if err == nil && l > 0 {
			limit = l
		}
	}

	hosts, total, err := h.HostService.GetAllHosts(page, limit)
	if err != nil {
		response.InternalServerError(w, "Gagal mengambil daftar host")
		return
	}

	result := map[string]interface{}{
		"hosts": hosts,
		"pagination": map[string]interface{}{
			"page":  page,
			"limit": limit,
			"total": total,
		},
	}

	response.Success(w, "Daftar host berhasil diambil", result)
}

// UpdateHost handler untuk memperbarui host
func (h *HostHandler) UpdateHost(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.BadRequest(w, "ID tidak valid")
		return
	}

	var req model.UpdateHostRequest
	
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Request body tidak valid")
		return
	}

	err = h.HostService.UpdateHost(id, &req)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	response.Success(w, "Host berhasil diupdate", nil)
}

// DeleteHost handler untuk menghapus host
func (h *HostHandler) DeleteHost(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.BadRequest(w, "ID tidak valid")
		return
	}

	err = h.HostService.DeleteHost(id)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	response.Success(w, "Host berhasil dihapus", nil)
}

// VirtualHostHandler menangani HTTP request untuk Virtual Host
type VirtualHostHandler struct {
	VHostService *service.VirtualHostService
}

// NewVirtualHostHandler membuat instance baru VirtualHostHandler
func NewVirtualHostHandler(vhostService *service.VirtualHostService) *VirtualHostHandler {
	return &VirtualHostHandler{VHostService: vhostService}
}

// CreateVirtualHost handler untuk membuat virtual host baru
func (h *VirtualHostHandler) CreateVirtualHost(w http.ResponseWriter, r *http.Request) {
	var req model.CreateVirtualHostRequest
	
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Request body tidak valid")
		return
	}

	if req.VHostName == "" {
		response.BadRequest(w, "Virtual host name wajib diisi")
		return
	}
	if req.HostID == 0 {
		response.BadRequest(w, "Host ID wajib diisi")
		return
	}

	vhost, err := h.VHostService.CreateVirtualHost(&req)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	response.Created(w, "Virtual host berhasil dibuat", vhost)
}

// GetVirtualHostByID handler untuk mengambil virtual host berdasarkan ID
func (h *VirtualHostHandler) GetVirtualHostByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.BadRequest(w, "ID tidak valid")
		return
	}

	vhost, err := h.VHostService.GetVirtualHostByID(id)
	if err != nil {
		response.NotFound(w, "Virtual host tidak ditemukan")
		return
	}

	response.Success(w, "Virtual host berhasil diambil", vhost)
}

// GetAllVirtualHosts handler untuk mengambil semua virtual host
func (h *VirtualHostHandler) GetAllVirtualHosts(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	page := 1
	limit := 10

	if pageStr != "" {
		p, err := strconv.Atoi(pageStr)
		if err == nil && p > 0 {
			page = p
		}
	}

	if limitStr != "" {
		l, err := strconv.Atoi(limitStr)
		if err == nil && l > 0 {
			limit = l
		}
	}

	vhosts, total, err := h.VHostService.GetAllVirtualHosts(page, limit)
	if err != nil {
		response.InternalServerError(w, "Gagal mengambil daftar virtual host")
		return
	}

	result := map[string]interface{}{
		"virtual_hosts": vhosts,
		"pagination": map[string]interface{}{
			"page":  page,
			"limit": limit,
			"total": total,
		},
	}

	response.Success(w, "Daftar virtual host berhasil diambil", result)
}

// UpdateVirtualHost handler untuk memperbarui virtual host
func (h *VirtualHostHandler) UpdateVirtualHost(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.BadRequest(w, "ID tidak valid")
		return
	}

	var req model.UpdateVirtualHostRequest
	
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Request body tidak valid")
		return
	}

	err = h.VHostService.UpdateVirtualHost(id, &req)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	response.Success(w, "Virtual host berhasil diupdate", nil)
}

// DeleteVirtualHost handler untuk menghapus virtual host
func (h *VirtualHostHandler) DeleteVirtualHost(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.BadRequest(w, "ID tidak valid")
		return
	}

	err = h.VHostService.DeleteVirtualHost(id)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	response.Success(w, "Virtual host berhasil dihapus", nil)
}
