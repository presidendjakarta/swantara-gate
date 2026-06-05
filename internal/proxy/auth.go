package proxy

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// AuthConfig holds authentication configuration for a route
type AuthConfig struct {
	VDirID   int64
	AuthType string // none, api_key, jwt, basic, external
}

// JWTRouteConfig holds JWT validation config for a route
type JWTRouteConfig struct {
	Algorithm        string
	Secret           string
	Issuer           string
	Audience         string
	ClockSkewSeconds int
	RequireExp       bool
	RequireIat       bool
}

// ExternalAuthConfig holds external (forward) auth config
type ExternalAuthConfig struct {
	AuthURL               string
	HTTPMethod            string
	RequestTimeoutSeconds int
	SendHeaders           bool
	SendBody              bool
	SuccessKey            string
	SuccessValue          string
	MessageKey            string
}

// Authenticator handles all proxy authentication types
type Authenticator struct {
	mu             sync.RWMutex
	db             *sql.DB
	jwtConfigs     map[int64]*JWTRouteConfig    // vdir_id -> jwt config
	externalAuths  map[int64]*ExternalAuthConfig // vdir_id -> external auth config
	httpClient     *http.Client
}

// NewAuthenticator creates a new Authenticator
func NewAuthenticator(db *sql.DB) *Authenticator {
	a := &Authenticator{
		db:            db,
		jwtConfigs:    make(map[int64]*JWTRouteConfig),
		externalAuths: make(map[int64]*ExternalAuthConfig),
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
	a.Reload()
	return a
}

// Reload loads authentication configs from database
func (a *Authenticator) Reload() {
	jwtConfigs := make(map[int64]*JWTRouteConfig)
	externalAuths := make(map[int64]*ExternalAuthConfig)

	// Load JWT configs
	rows, err := a.db.Query(`
		SELECT virtual_directory_id, algorithm, jwt_secret, 
		       COALESCE(issuer, ''), COALESCE(audience, ''),
		       clock_skew_seconds, require_exp, require_iat
		FROM jwt_configs WHERE is_active = 1
	`)
	if err != nil {
		log.Printf("[Proxy:Auth] Error loading JWT configs: %v", err)
	} else {
		defer rows.Close()
		for rows.Next() {
			var vdirID int64
			var cfg JWTRouteConfig
			var requireExp, requireIat int
			if err := rows.Scan(&vdirID, &cfg.Algorithm, &cfg.Secret,
				&cfg.Issuer, &cfg.Audience, &cfg.ClockSkewSeconds,
				&requireExp, &requireIat); err != nil {
				log.Printf("[Proxy:Auth] Error scanning JWT config: %v", err)
				continue
			}
			cfg.RequireExp = requireExp == 1
			cfg.RequireIat = requireIat == 1
			jwtConfigs[vdirID] = &cfg
		}
	}

	// Load external auth configs
	rows2, err := a.db.Query(`
		SELECT virtual_directory_id, auth_url, http_method, request_timeout_seconds,
		       send_headers, send_body, COALESCE(success_key, 'status'),
		       COALESCE(success_value, 'true'), COALESCE(message_key, 'message')
		FROM external_auth WHERE is_active = 1
	`)
	if err != nil {
		log.Printf("[Proxy:Auth] Error loading external auth configs: %v", err)
	} else {
		defer rows2.Close()
		for rows2.Next() {
			var vdirID int64
			var cfg ExternalAuthConfig
			var sendHeaders, sendBody int
			if err := rows2.Scan(&vdirID, &cfg.AuthURL, &cfg.HTTPMethod,
				&cfg.RequestTimeoutSeconds, &sendHeaders, &sendBody,
				&cfg.SuccessKey, &cfg.SuccessValue, &cfg.MessageKey); err != nil {
				log.Printf("[Proxy:Auth] Error scanning external auth: %v", err)
				continue
			}
			cfg.SendHeaders = sendHeaders == 1
			cfg.SendBody = sendBody == 1
			externalAuths[vdirID] = &cfg
		}
	}

	a.mu.Lock()
	a.jwtConfigs = jwtConfigs
	a.externalAuths = externalAuths
	a.mu.Unlock()

	log.Printf("[Proxy:Auth] Loaded %d JWT configs, %d external auth configs",
		len(jwtConfigs), len(externalAuths))
}

// Authenticate checks the request against the route's auth_type
func (a *Authenticator) Authenticate(r *http.Request, route *VDirConfig) (bool, string) {
	switch route.AuthType {
	case "none", "":
		return true, ""
	case "api_key":
		return a.authenticateAPIKey(r, route)
	case "jwt":
		return a.authenticateJWT(r, route)
	case "basic":
		return a.authenticateBasic(r, route)
	case "external":
		return a.authenticateExternal(r, route)
	default:
		return true, ""
	}
}

// authenticateAPIKey validates API key from header or query param
func (a *Authenticator) authenticateAPIKey(r *http.Request, route *VDirConfig) (bool, string) {
	// Check X-API-Key header first, then query param
	apiKey := r.Header.Get("X-API-Key")
	if apiKey == "" {
		apiKey = r.URL.Query().Get("api_key")
	}
	if apiKey == "" {
		return false, "API key required"
	}

	// Validate against database
	var consumerID int64
	var isActive int
	err := a.db.QueryRow(`
		SELECT ak.consumer_id, ak.is_active FROM api_keys ak
		JOIN api_consumers ac ON ac.id = ak.consumer_id
		WHERE ak.api_key = ? AND ak.is_active = 1 AND ac.is_active = 1
		AND (ak.expired_at IS NULL OR ak.expired_at > datetime('now'))
	`, apiKey).Scan(&consumerID, &isActive)
	if err != nil {
		return false, "invalid API key"
	}

	// Check route consumer access (ACL)
	if !a.checkRouteAccess(route.ID, consumerID) {
		return false, "access denied for this consumer"
	}

	return true, ""
}

// authenticateJWT validates JWT token from Authorization header
func (a *Authenticator) authenticateJWT(r *http.Request, route *VDirConfig) (bool, string) {
	// Extract token from Authorization: Bearer <token>
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return false, "Bearer token required"
	}
	token := strings.TrimPrefix(authHeader, "Bearer ")

	// Get JWT config for this route
	a.mu.RLock()
	cfg, ok := a.jwtConfigs[route.ID]
	a.mu.RUnlock()
	if !ok {
		return false, "JWT not configured for this route"
	}

	// Validate the JWT token
	valid, err := validateJWTToken(token, cfg)
	if err != nil {
		return false, fmt.Sprintf("JWT validation failed: %v", err)
	}
	if !valid {
		return false, "invalid JWT token"
	}

	return true, ""
}

// authenticateBasic validates HTTP Basic Auth credentials
func (a *Authenticator) authenticateBasic(r *http.Request, route *VDirConfig) (bool, string) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Basic ") {
		return false, "Basic authentication required"
	}

	decoded, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(authHeader, "Basic "))
	if err != nil {
		return false, "invalid basic auth encoding"
	}

	parts := strings.SplitN(string(decoded), ":", 2)
	if len(parts) != 2 {
		return false, "invalid basic auth format"
	}
	username, password := parts[0], parts[1]

	// Validate against consumer_credentials
	var passwordHash string
	var consumerID int64
	err = a.db.QueryRow(`
		SELECT cc.consumer_id, cc.password_hash FROM consumer_credentials cc
		JOIN api_consumers ac ON ac.id = cc.consumer_id
		WHERE cc.auth_type = 'basic' AND cc.username = ? AND cc.is_active = 1 AND ac.is_active = 1
		AND (cc.expired_at IS NULL OR cc.expired_at > datetime('now'))
	`, username).Scan(&consumerID, &passwordHash)
	if err != nil {
		return false, "invalid credentials"
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password)); err != nil {
		return false, "invalid credentials"
	}

	// Check route consumer access (ACL)
	if !a.checkRouteAccess(route.ID, consumerID) {
		return false, "access denied for this consumer"
	}

	return true, ""
}

// authenticateExternal calls an external auth service
func (a *Authenticator) authenticateExternal(r *http.Request, route *VDirConfig) (bool, string) {
	a.mu.RLock()
	cfg, ok := a.externalAuths[route.ID]
	a.mu.RUnlock()
	if !ok {
		return false, "external auth not configured for this route"
	}

	timeout := time.Duration(cfg.RequestTimeoutSeconds) * time.Second
	if timeout <= 0 {
		timeout = 5 * time.Second
	}
	ctx, cancel := context.WithTimeout(r.Context(), timeout)
	defer cancel()

	method := cfg.HTTPMethod
	if method == "" {
		method = "POST"
	}

	// Build external auth request
	var body io.Reader
	if cfg.SendBody && r.Body != nil {
		body = r.Body
	}

	authReq, err := http.NewRequestWithContext(ctx, method, cfg.AuthURL, body)
	if err != nil {
		return false, "failed to create auth request"
	}

	// Forward headers if configured
	if cfg.SendHeaders {
		for k, vv := range r.Header {
			for _, v := range vv {
				authReq.Header.Add(k, v)
			}
		}
	}

	resp, err := a.httpClient.Do(authReq)
	if err != nil {
		log.Printf("[Proxy:Auth] External auth request failed: %v", err)
		return false, "external auth service unavailable"
	}
	defer resp.Body.Close()

	// Read response
	respBody, _ := io.ReadAll(resp.Body)

	// Check HTTP status
	if resp.StatusCode != http.StatusOK {
		return false, "external auth denied"
	}

	// Parse JSON response and check success key
	var result map[string]interface{}
	if err := json.Unmarshal(respBody, &result); err == nil {
		if val, ok := result[cfg.SuccessKey]; ok {
			valStr := fmt.Sprintf("%v", val)
			if valStr != cfg.SuccessValue {
				msg := "access denied"
				if msgVal, ok := result[cfg.MessageKey]; ok {
					msg = fmt.Sprintf("%v", msgVal)
				}
				return false, msg
			}
		}
	}

	return true, ""
}

// checkRouteAccess checks if a consumer has access to a route via ACL
func (a *Authenticator) checkRouteAccess(vdirID, consumerID int64) bool {
	// Check if this route has any consumer ACL entries
	var count int
	err := a.db.QueryRow(`SELECT COUNT(*) FROM route_consumer_access WHERE virtual_directory_id = ? AND is_active = 1`, vdirID).Scan(&count)
	if err != nil {
		log.Printf("[Proxy:Auth] Error checking route access for vdir=%d: %v", vdirID, err)
		return false // Error = deny by default (secure)
	}
	
	// SECURITY: If no ACL entries exist for this route, DENY all access
	// This enforces mandatory consumer selection (deny-by-default policy)
	if count == 0 {
		return false // No ACL configured = deny all (must select at least 1 consumer)
	}

	// Check if this specific consumer has access
	var accessCount int
	err = a.db.QueryRow(`SELECT COUNT(*) FROM route_consumer_access WHERE virtual_directory_id = ? AND consumer_id = ? AND is_active = 1`, vdirID, consumerID).Scan(&accessCount)
	return err == nil && accessCount > 0
}

// validateJWTToken validates a JWT token with HMAC-SHA256
func validateJWTToken(token string, cfg *JWTRouteConfig) (bool, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return false, fmt.Errorf("invalid token format")
	}

	// Verify signature (HMAC-SHA256)
	signingInput := parts[0] + "." + parts[1]
	expectedSig, err := computeHMAC(signingInput, cfg.Secret)
	if err != nil {
		return false, err
	}

	actualSig := parts[2]
	// Compare base64url encoded signatures
	if !hmacEqual(expectedSig, actualSig) {
		return false, fmt.Errorf("signature mismatch")
	}

	// Decode payload
	payload, err := base64URLDecode(parts[1])
	if err != nil {
		return false, fmt.Errorf("invalid payload encoding")
	}

	var claims map[string]interface{}
	if err := json.Unmarshal(payload, &claims); err != nil {
		return false, fmt.Errorf("invalid payload JSON")
	}

	now := time.Now()
	clockSkew := time.Duration(cfg.ClockSkewSeconds) * time.Second

	// Check expiration
	if cfg.RequireExp {
		exp, ok := claims["exp"].(float64)
		if !ok {
			return false, fmt.Errorf("exp claim required")
		}
		expTime := time.Unix(int64(exp), 0)
		if now.After(expTime.Add(clockSkew)) {
			return false, fmt.Errorf("token expired")
		}
	}

	// Check issued at
	if cfg.RequireIat {
		if _, ok := claims["iat"]; !ok {
			return false, fmt.Errorf("iat claim required")
		}
	}

	// Check issuer
	if cfg.Issuer != "" {
		if iss, ok := claims["iss"].(string); !ok || iss != cfg.Issuer {
			return false, fmt.Errorf("invalid issuer")
		}
	}

	// Check audience
	if cfg.Audience != "" {
		if aud, ok := claims["aud"].(string); !ok || aud != cfg.Audience {
			return false, fmt.Errorf("invalid audience")
		}
	}

	return true, nil
}

func computeHMAC(input, secret string) (string, error) {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(input))
	sig := mac.Sum(nil)
	return base64URLEncode(sig), nil
}

func hmacEqual(expected, actual string) bool {
	// Normalize base64url (handle padding differences)
	a := strings.TrimRight(expected, "=")
	b := strings.TrimRight(actual, "=")
	return a == b
}

func base64URLDecode(s string) ([]byte, error) {
	// Add padding if needed
	switch len(s) % 4 {
	case 2:
		s += "=="
	case 3:
		s += "="
	}
	return base64.URLEncoding.DecodeString(s)
}

func base64URLEncode(data []byte) string {
	return strings.TrimRight(base64.URLEncoding.EncodeToString(data), "=")
}
