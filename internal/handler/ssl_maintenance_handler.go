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
// ACME ACCOUNT HANDLER
// =========================================================

type ACMEAccountHandler struct{ Service *service.ACMEAccountService }

func NewACMEAccountHandler(svc *service.ACMEAccountService) *ACMEAccountHandler {
	return &ACMEAccountHandler{Service: svc}
}

func (h *ACMEAccountHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req model.CreateACMEAccountRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil { response.BadRequest(w, "Request body tidak valid"); return }
	item, err := h.Service.Create(&req)
	if err != nil { response.BadRequest(w, err.Error()); return }
	response.Created(w, "ACME account berhasil dibuat", item)
}

func (h *ACMEAccountHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil { response.BadRequest(w, "ID tidak valid"); return }
	item, err := h.Service.GetByID(id)
	if err != nil { response.NotFound(w, err.Error()); return }
	response.Success(w, "ACME account berhasil diambil", item)
}

func (h *ACMEAccountHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	page, limit := parsePagination(r)
	items, total, err := h.Service.GetAll(page, limit)
	if err != nil { response.InternalServerError(w, "Gagal mengambil ACME accounts"); return }
	response.Success(w, "Daftar ACME accounts", map[string]interface{}{"acme_accounts": items, "pagination": paginationData(page, limit, total)})
}

func (h *ACMEAccountHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil { response.BadRequest(w, "ID tidak valid"); return }
	var req model.UpdateACMEAccountRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil { response.BadRequest(w, "Request body tidak valid"); return }
	if err := h.Service.Update(id, &req); err != nil { response.BadRequest(w, err.Error()); return }
	response.Success(w, "ACME account berhasil diperbarui", nil)
}

func (h *ACMEAccountHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil { response.BadRequest(w, "ID tidak valid"); return }
	if err := h.Service.Delete(id); err != nil { response.BadRequest(w, err.Error()); return }
	response.Success(w, "ACME account berhasil dihapus", nil)
}

// =========================================================
// SSL CERTIFICATE HANDLER
// =========================================================

type SSLCertificateHandler struct{ Service *service.SSLCertificateService }

func NewSSLCertificateHandler(svc *service.SSLCertificateService) *SSLCertificateHandler {
	return &SSLCertificateHandler{Service: svc}
}

func (h *SSLCertificateHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req model.CreateSSLCertificateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil { response.BadRequest(w, "Request body tidak valid"); return }
	item, err := h.Service.Create(&req)
	if err != nil { response.BadRequest(w, err.Error()); return }
	response.Created(w, "SSL certificate berhasil dibuat", item)
}

func (h *SSLCertificateHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil { response.BadRequest(w, "ID tidak valid"); return }
	item, err := h.Service.GetByID(id)
	if err != nil { response.NotFound(w, err.Error()); return }
	response.Success(w, "SSL certificate berhasil diambil", item)
}

func (h *SSLCertificateHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	page, limit := parsePagination(r)
	items, total, err := h.Service.GetAll(page, limit)
	if err != nil { response.InternalServerError(w, "Gagal mengambil SSL certificates"); return }
	response.Success(w, "Daftar SSL certificates", map[string]interface{}{"ssl_certificates": items, "pagination": paginationData(page, limit, total)})
}

func (h *SSLCertificateHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil { response.BadRequest(w, "ID tidak valid"); return }
	var req model.UpdateSSLCertificateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil { response.BadRequest(w, "Request body tidak valid"); return }
	if err := h.Service.Update(id, &req); err != nil { response.BadRequest(w, err.Error()); return }
	response.Success(w, "SSL certificate berhasil diperbarui", nil)
}

func (h *SSLCertificateHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil { response.BadRequest(w, "ID tidak valid"); return }
	if err := h.Service.Delete(id); err != nil { response.BadRequest(w, err.Error()); return }
	response.Success(w, "SSL certificate berhasil dihapus", nil)
}

// =========================================================
// CERTIFICATE DOMAIN HANDLER
// =========================================================

type CertificateDomainHandler struct{ Service *service.CertificateDomainService }

func NewCertificateDomainHandler(svc *service.CertificateDomainService) *CertificateDomainHandler {
	return &CertificateDomainHandler{Service: svc}
}

func (h *CertificateDomainHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req model.CreateCertificateDomainRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil { response.BadRequest(w, "Request body tidak valid"); return }
	item, err := h.Service.Create(&req)
	if err != nil { response.BadRequest(w, err.Error()); return }
	response.Created(w, "Certificate domain berhasil dibuat", item)
}

func (h *CertificateDomainHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil { response.BadRequest(w, "ID tidak valid"); return }
	item, err := h.Service.GetByID(id)
	if err != nil { response.NotFound(w, err.Error()); return }
	response.Success(w, "Certificate domain berhasil diambil", item)
}

func (h *CertificateDomainHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	page, limit := parsePagination(r)
	items, total, err := h.Service.GetAll(page, limit)
	if err != nil { response.InternalServerError(w, "Gagal mengambil certificate domains"); return }
	response.Success(w, "Daftar certificate domains", map[string]interface{}{"certificate_domains": items, "pagination": paginationData(page, limit, total)})
}

func (h *CertificateDomainHandler) GetByCertificate(w http.ResponseWriter, r *http.Request) {
	certID, err := strconv.ParseInt(r.PathValue("cert_id"), 10, 64)
	if err != nil { response.BadRequest(w, "Certificate ID tidak valid"); return }
	items, err := h.Service.GetByCertificateID(certID)
	if err != nil { response.InternalServerError(w, "Gagal mengambil domains"); return }
	response.Success(w, "Certificate domains berhasil diambil", items)
}

func (h *CertificateDomainHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil { response.BadRequest(w, "ID tidak valid"); return }
	var req model.UpdateCertificateDomainRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil { response.BadRequest(w, "Request body tidak valid"); return }
	if err := h.Service.Update(id, &req); err != nil { response.BadRequest(w, err.Error()); return }
	response.Success(w, "Certificate domain berhasil diperbarui", nil)
}

func (h *CertificateDomainHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil { response.BadRequest(w, "ID tidak valid"); return }
	if err := h.Service.Delete(id); err != nil { response.BadRequest(w, err.Error()); return }
	response.Success(w, "Certificate domain berhasil dihapus", nil)
}

// =========================================================
// SSL CERTIFICATE BINDING HANDLER
// =========================================================

type SSLCertificateBindingHandler struct{ Service *service.SSLCertificateBindingService }

func NewSSLCertificateBindingHandler(svc *service.SSLCertificateBindingService) *SSLCertificateBindingHandler {
	return &SSLCertificateBindingHandler{Service: svc}
}

func (h *SSLCertificateBindingHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req model.CreateSSLCertificateBindingRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil { response.BadRequest(w, "Request body tidak valid"); return }
	item, err := h.Service.Create(&req)
	if err != nil { response.BadRequest(w, err.Error()); return }
	response.Created(w, "SSL binding berhasil dibuat", item)
}

func (h *SSLCertificateBindingHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil { response.BadRequest(w, "ID tidak valid"); return }
	item, err := h.Service.GetByID(id)
	if err != nil { response.NotFound(w, err.Error()); return }
	response.Success(w, "SSL binding berhasil diambil", item)
}

func (h *SSLCertificateBindingHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	page, limit := parsePagination(r)
	items, total, err := h.Service.GetAll(page, limit)
	if err != nil { response.InternalServerError(w, "Gagal mengambil SSL bindings"); return }
	response.Success(w, "Daftar SSL bindings", map[string]interface{}{"ssl_bindings": items, "pagination": paginationData(page, limit, total)})
}

func (h *SSLCertificateBindingHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil { response.BadRequest(w, "ID tidak valid"); return }
	var req model.UpdateSSLCertificateBindingRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil { response.BadRequest(w, "Request body tidak valid"); return }
	if err := h.Service.Update(id, &req); err != nil { response.BadRequest(w, err.Error()); return }
	response.Success(w, "SSL binding berhasil diperbarui", nil)
}

func (h *SSLCertificateBindingHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil { response.BadRequest(w, "ID tidak valid"); return }
	if err := h.Service.Delete(id); err != nil { response.BadRequest(w, err.Error()); return }
	response.Success(w, "SSL binding berhasil dihapus", nil)
}

// =========================================================
// TLS OPTION HANDLER
// =========================================================

type TLSOptionHandler struct{ Service *service.TLSOptionService }

func NewTLSOptionHandler(svc *service.TLSOptionService) *TLSOptionHandler {
	return &TLSOptionHandler{Service: svc}
}

func (h *TLSOptionHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req model.CreateTLSOptionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil { response.BadRequest(w, "Request body tidak valid"); return }
	item, err := h.Service.Create(&req)
	if err != nil { response.BadRequest(w, err.Error()); return }
	response.Created(w, "TLS option berhasil dibuat", item)
}

func (h *TLSOptionHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil { response.BadRequest(w, "ID tidak valid"); return }
	item, err := h.Service.GetByID(id)
	if err != nil { response.NotFound(w, err.Error()); return }
	response.Success(w, "TLS option berhasil diambil", item)
}

func (h *TLSOptionHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	page, limit := parsePagination(r)
	items, total, err := h.Service.GetAll(page, limit)
	if err != nil { response.InternalServerError(w, "Gagal mengambil TLS options"); return }
	response.Success(w, "Daftar TLS options", map[string]interface{}{"tls_options": items, "pagination": paginationData(page, limit, total)})
}

func (h *TLSOptionHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil { response.BadRequest(w, "ID tidak valid"); return }
	var req model.UpdateTLSOptionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil { response.BadRequest(w, "Request body tidak valid"); return }
	if err := h.Service.Update(id, &req); err != nil { response.BadRequest(w, err.Error()); return }
	response.Success(w, "TLS option berhasil diperbarui", nil)
}

func (h *TLSOptionHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil { response.BadRequest(w, "ID tidak valid"); return }
	if err := h.Service.Delete(id); err != nil { response.BadRequest(w, err.Error()); return }
	response.Success(w, "TLS option berhasil dihapus", nil)
}

// =========================================================
// SERVICE DISCOVERY HANDLER
// =========================================================

type ServiceDiscoveryHandler struct{ Service *service.ServiceDiscoveryService }

func NewServiceDiscoveryHandler(svc *service.ServiceDiscoveryService) *ServiceDiscoveryHandler {
	return &ServiceDiscoveryHandler{Service: svc}
}

func (h *ServiceDiscoveryHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req model.CreateServiceDiscoveryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil { response.BadRequest(w, "Request body tidak valid"); return }
	item, err := h.Service.Create(&req)
	if err != nil { response.BadRequest(w, err.Error()); return }
	response.Created(w, "Service discovery berhasil dibuat", item)
}

func (h *ServiceDiscoveryHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil { response.BadRequest(w, "ID tidak valid"); return }
	item, err := h.Service.GetByID(id)
	if err != nil { response.NotFound(w, err.Error()); return }
	response.Success(w, "Service discovery berhasil diambil", item)
}

func (h *ServiceDiscoveryHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	page, limit := parsePagination(r)
	items, total, err := h.Service.GetAll(page, limit)
	if err != nil { response.InternalServerError(w, "Gagal mengambil service discoveries"); return }
	response.Success(w, "Daftar service discoveries", map[string]interface{}{"service_discoveries": items, "pagination": paginationData(page, limit, total)})
}

func (h *ServiceDiscoveryHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil { response.BadRequest(w, "ID tidak valid"); return }
	var req model.UpdateServiceDiscoveryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil { response.BadRequest(w, "Request body tidak valid"); return }
	if err := h.Service.Update(id, &req); err != nil { response.BadRequest(w, err.Error()); return }
	response.Success(w, "Service discovery berhasil diperbarui", nil)
}

func (h *ServiceDiscoveryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil { response.BadRequest(w, "ID tidak valid"); return }
	if err := h.Service.Delete(id); err != nil { response.BadRequest(w, err.Error()); return }
	response.Success(w, "Service discovery berhasil dihapus", nil)
}

// =========================================================
// CONFIG VERSION HANDLER
// =========================================================

type ConfigVersionHandler struct{ Service *service.ConfigVersionService }

func NewConfigVersionHandler(svc *service.ConfigVersionService) *ConfigVersionHandler {
	return &ConfigVersionHandler{Service: svc}
}

func (h *ConfigVersionHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req model.CreateConfigVersionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil { response.BadRequest(w, "Request body tidak valid"); return }
	item, err := h.Service.Create(&req)
	if err != nil { response.BadRequest(w, err.Error()); return }
	response.Created(w, "Config version berhasil dibuat", item)
}

func (h *ConfigVersionHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil { response.BadRequest(w, "ID tidak valid"); return }
	item, err := h.Service.GetByID(id)
	if err != nil { response.NotFound(w, err.Error()); return }
	response.Success(w, "Config version berhasil diambil", item)
}

func (h *ConfigVersionHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	page, limit := parsePagination(r)
	items, total, err := h.Service.GetAll(page, limit)
	if err != nil { response.InternalServerError(w, "Gagal mengambil config versions"); return }
	response.Success(w, "Daftar config versions", map[string]interface{}{"config_versions": items, "pagination": paginationData(page, limit, total)})
}

func (h *ConfigVersionHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil { response.BadRequest(w, "ID tidak valid"); return }
	var req model.UpdateConfigVersionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil { response.BadRequest(w, "Request body tidak valid"); return }
	if err := h.Service.Update(id, &req); err != nil { response.BadRequest(w, err.Error()); return }
	response.Success(w, "Config version berhasil diperbarui", nil)
}

func (h *ConfigVersionHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil { response.BadRequest(w, "ID tidak valid"); return }
	if err := h.Service.Delete(id); err != nil { response.BadRequest(w, err.Error()); return }
	response.Success(w, "Config version berhasil dihapus", nil)
}

// =========================================================
// MAINTENANCE WINDOW HANDLER
// =========================================================

type MaintenanceWindowHandler struct{ Service *service.MaintenanceWindowService }

func NewMaintenanceWindowHandler(svc *service.MaintenanceWindowService) *MaintenanceWindowHandler {
	return &MaintenanceWindowHandler{Service: svc}
}

func (h *MaintenanceWindowHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req model.CreateMaintenanceWindowRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil { response.BadRequest(w, "Request body tidak valid"); return }
	item, err := h.Service.Create(&req)
	if err != nil { response.BadRequest(w, err.Error()); return }
	response.Created(w, "Maintenance window berhasil dibuat", item)
}

func (h *MaintenanceWindowHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil { response.BadRequest(w, "ID tidak valid"); return }
	item, err := h.Service.GetByID(id)
	if err != nil { response.NotFound(w, err.Error()); return }
	response.Success(w, "Maintenance window berhasil diambil", item)
}

func (h *MaintenanceWindowHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	page, limit := parsePagination(r)
	items, total, err := h.Service.GetAll(page, limit)
	if err != nil { response.InternalServerError(w, "Gagal mengambil maintenance windows"); return }
	response.Success(w, "Daftar maintenance windows", map[string]interface{}{"maintenance_windows": items, "pagination": paginationData(page, limit, total)})
}

func (h *MaintenanceWindowHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil { response.BadRequest(w, "ID tidak valid"); return }
	var req model.UpdateMaintenanceWindowRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil { response.BadRequest(w, "Request body tidak valid"); return }
	if err := h.Service.Update(id, &req); err != nil { response.BadRequest(w, err.Error()); return }
	response.Success(w, "Maintenance window berhasil diperbarui", nil)
}

func (h *MaintenanceWindowHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil { response.BadRequest(w, "ID tidak valid"); return }
	if err := h.Service.Delete(id); err != nil { response.BadRequest(w, err.Error()); return }
	response.Success(w, "Maintenance window berhasil dihapus", nil)
}
