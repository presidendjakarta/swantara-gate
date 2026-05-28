package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/presidendjakarta/swantara-gate/internal/model"
	"github.com/presidendjakarta/swantara-gate/internal/response"
	"github.com/presidendjakarta/swantara-gate/internal/service"
)

// AuthHandler menangani request autentikasi
type AuthHandler struct {
	Service *service.AuthService
}

// NewAuthHandler membuat instance baru AuthHandler
func NewAuthHandler(svc *service.AuthService) *AuthHandler {
	return &AuthHandler{Service: svc}
}

// Login melakukan autentikasi user dan mengembalikan token
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req model.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Request body tidak valid")
		return
	}

	tokenResp, err := h.Service.Login(&req)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, err.Error())
		return
	}

	response.Success(w, "Login berhasil", tokenResp)
}

// Refresh menghasilkan access token baru dari refresh token
func (h *AuthHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	var req model.RefreshTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Request body tidak valid")
		return
	}

	tokenResp, err := h.Service.RefreshToken(req.RefreshToken)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, err.Error())
		return
	}

	response.Success(w, "Token berhasil diperbarui", tokenResp)
}

// Logout mencabut refresh token dan blacklist access token
func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	// Parse body (opsional, bisa kosong)
	var req model.RefreshTokenRequest
	_ = json.NewDecoder(r.Body).Decode(&req)

	// Ambil access token dari Authorization header untuk di-blacklist
	var accessToken string
	authHeader := r.Header.Get("Authorization")
	if authHeader != "" {
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) == 2 && strings.ToLower(parts[0]) == "bearer" {
			accessToken = parts[1]
		}
	}

	// Minimal harus ada salah satu: access token atau refresh token
	if accessToken == "" && req.RefreshToken == "" {
		response.BadRequest(w, "Kirim Authorization header atau refresh_token di body")
		return
	}

	if err := h.Service.Logout(req.RefreshToken, accessToken); err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	response.Success(w, "Logout berhasil", nil)
}

// Me mengembalikan profil user yang sedang login
func (h *AuthHandler) Me(w http.ResponseWriter, r *http.Request) {
	// Claims diambil dari context yang di-set oleh auth middleware
	claims := GetClaimsFromContext(r)
	if claims == nil {
		response.Error(w, http.StatusUnauthorized, "Token tidak valid")
		return
	}

	response.Success(w, "Profil user", map[string]interface{}{
		"user_id":  claims.UserID,
		"username": claims.Username,
		"role":     claims.Role,
	})
}

// GetClaimsFromContext mengambil JWT claims dari request context
func GetClaimsFromContext(r *http.Request) *model.JWTClaims {
	claims, _ := r.Context().Value(model.ContextKeyClaims).(*model.JWTClaims)
	return claims
}
