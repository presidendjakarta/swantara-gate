package middleware

import (
	"context"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/presidendjakarta/swantara-gate/internal/model"
	"github.com/presidendjakarta/swantara-gate/internal/response"
	"github.com/presidendjakarta/swantara-gate/internal/service"
)

// AuthMiddleware middleware untuk autentikasi JWT
type AuthMiddleware struct {
	AuthService *service.AuthService
}

// NewAuthMiddleware membuat instance baru AuthMiddleware
func NewAuthMiddleware(authService *service.AuthService) *AuthMiddleware {
	return &AuthMiddleware{AuthService: authService}
}

// RequireAuth memvalidasi JWT token di header Authorization
func (m *AuthMiddleware) RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			response.Error(w, http.StatusUnauthorized, "Token tidak ditemukan")
			return
		}

		// Ambil token dari "Bearer <token>"
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			response.Error(w, http.StatusUnauthorized, "Format token tidak valid")
			return
		}

		tokenStr := parts[1]

		// Validasi token
		claims, err := m.AuthService.ValidateAccessToken(tokenStr)
		if err != nil {
			response.Error(w, http.StatusUnauthorized, "Token tidak valid atau sudah expired")
			return
		}

		// Set claims ke context
		ctx := context.WithValue(r.Context(), model.ContextKeyClaims, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// RequireRole memvalidasi role user
func (m *AuthMiddleware) RequireRole(roles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			claims, ok := r.Context().Value(model.ContextKeyClaims).(*model.JWTClaims)
			if !ok {
				response.Error(w, http.StatusUnauthorized, "Token tidak valid")
				return
			}

			// Cek apakah role user ada di daftar yang diizinkan
			allowed := false
			for _, role := range roles {
				if claims.Role == role {
					allowed = true
					break
				}
			}

			if !allowed {
				response.Error(w, http.StatusForbidden, "Akses ditolak: role tidak memiliki izin")
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

// RequireAuthExcept membungkus handler dengan auth check kecuali untuk path tertentu
func (m *AuthMiddleware) RequireAuthExcept(next http.Handler, publicPaths ...string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Cek apakah path termasuk public (tidak perlu auth)
		for _, publicPath := range publicPaths {
			if r.URL.Path == publicPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		// Untuk path yang bukan public, wajib auth
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			response.Error(w, http.StatusUnauthorized, "Token tidak ditemukan")
			return
		}

		// Ambil token dari "Bearer <token>"
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			response.Error(w, http.StatusUnauthorized, "Format token tidak valid")
			return
		}

		tokenStr := parts[1]

		// Validasi token
		claims, err := m.AuthService.ValidateAccessToken(tokenStr)
		if err != nil {
			response.Error(w, http.StatusUnauthorized, "Token tidak valid atau sudah expired")
			return
		}

		// Set claims ke context
		ctx := context.WithValue(r.Context(), model.ContextKeyClaims, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// =========================================================
// LOGIN RATE LIMITER
// =========================================================

// LoginRateLimiter rate limiter sederhana untuk endpoint login
type LoginRateLimiter struct {
	mu       sync.Mutex
	attempts map[string]*rateLimitEntry
	maxAttempts int
	windowSec   int
}

type rateLimitEntry struct {
	count    int
	resetAt  time.Time
}

// NewLoginRateLimiter membuat instance login rate limiter
func NewLoginRateLimiter(maxAttempts, windowSec int) *LoginRateLimiter {
	return &LoginRateLimiter{
		attempts:    make(map[string]*rateLimitEntry),
		maxAttempts: maxAttempts,
		windowSec:   windowSec,
	}
}

// RateLimit middleware rate limiter untuk login
func (rl *LoginRateLimiter) RateLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := getClientIP(r)

		rl.mu.Lock()
		entry, exists := rl.attempts[ip]
		now := time.Now()

		if !exists || now.After(entry.resetAt) {
			// Reset atau buat entry baru
			rl.attempts[ip] = &rateLimitEntry{
				count:   1,
				resetAt: now.Add(time.Duration(rl.windowSec) * time.Second),
			}
			rl.mu.Unlock()
			next.ServeHTTP(w, r)
			return
		}

		entry.count++
		if entry.count > rl.maxAttempts {
			rl.mu.Unlock()
			response.Error(w, http.StatusTooManyRequests, "Terlalu banyak percobaan login. Coba lagi nanti.")
			return
		}
		rl.mu.Unlock()

		next.ServeHTTP(w, r)
	})
}

// getClientIP mengambil IP address client
func getClientIP(r *http.Request) string {
	// Cek X-Forwarded-For header
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		parts := strings.Split(xff, ",")
		return strings.TrimSpace(parts[0])
	}
	// Cek X-Real-IP header
	if xri := r.Header.Get("X-Real-IP"); xri != "" {
		return xri
	}
	// Gunakan RemoteAddr
	parts := strings.Split(r.RemoteAddr, ":")
	if len(parts) > 0 {
		return parts[0]
	}
	return r.RemoteAddr
}
