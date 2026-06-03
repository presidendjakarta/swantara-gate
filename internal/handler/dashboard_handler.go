package handler

import (
	"net/http"

	"github.com/presidendjakarta/swantara-gate/internal/response"
	"github.com/presidendjakarta/swantara-gate/internal/service"
)

// DashboardHandler menangani HTTP request untuk dashboard
type DashboardHandler struct {
	Service *service.DashboardService
}

// NewDashboardHandler membuat instance baru DashboardHandler
func NewDashboardHandler(svc *service.DashboardService) *DashboardHandler {
	return &DashboardHandler{Service: svc}
}

// GetStats handler untuk mengambil statistik dashboard
func (h *DashboardHandler) GetStats(w http.ResponseWriter, r *http.Request) {
	stats, err := h.Service.GetDashboardStats()
	if err != nil {
		response.InternalServerError(w, "Gagal mengambil statistik dashboard")
		return
	}

	response.Success(w, "Statistik dashboard berhasil diambil", stats)
}
