package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/presidendjakarta/swantara-gate/internal/model"
	"github.com/presidendjakarta/swantara-gate/internal/response"
	"github.com/presidendjakarta/swantara-gate/internal/service"
)

// UserHandler menangani HTTP request untuk User
type UserHandler struct {
	UserService *service.UserService
}

// NewUserHandler membuat instance baru UserHandler
func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

// CreateUser handler untuk membuat user baru
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req model.CreateUserRequest
	
	// Decode JSON request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Request body tidak valid")
		return
	}

	// Validasi input
	if req.Username == "" {
		response.BadRequest(w, "Username wajib diisi")
		return
	}
	if req.Password == "" {
		response.BadRequest(w, "Password wajib diisi")
		return
	}

	// Membuat user baru melalui service
	user, err := h.UserService.CreateUser(&req)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	response.Created(w, "User berhasil dibuat", user)
}

// GetUserByID handler untuk mengambil user berdasarkan ID
func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	// Mengambil ID dari URL parameter
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.BadRequest(w, "ID tidak valid")
		return
	}

	// Mengambil user dari service
	user, err := h.UserService.GetUserByID(id)
	if err != nil {
		response.NotFound(w, "User tidak ditemukan")
		return
	}

	response.Success(w, "User berhasil diambil", user)
}

// GetAllUsers handler untuk mengambil semua user
func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	// Mengambil pagination parameters dari query string
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

	// Mengambil daftar user dari service
	users, total, err := h.UserService.GetAllUsers(page, limit)
	if err != nil {
		response.InternalServerError(w, "Gagal mengambil daftar user")
		return
	}

	// Response dengan pagination info
	result := map[string]interface{}{
		"users": users,
		"pagination": map[string]interface{}{
			"page":  page,
			"limit": limit,
			"total": total,
		},
	}

	response.Success(w, "Daftar user berhasil diambil", result)
}

// UpdateUser handler untuk memperbarui user
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Mengambil ID dari URL parameter
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.BadRequest(w, "ID tidak valid")
		return
	}

	var req model.UpdateUserRequest
	
	// Decode JSON request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Request body tidak valid")
		return
	}

	// Update user melalui service
	err = h.UserService.UpdateUser(id, &req)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	response.Success(w, "User berhasil diupdate", nil)
}

// DeleteUser handler untuk menghapus user
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Mengambil ID dari URL parameter
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.BadRequest(w, "ID tidak valid")
		return
	}

	// Delete user melalui service
	err = h.UserService.DeleteUser(id)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	response.Success(w, "User berhasil dihapus", nil)
}
