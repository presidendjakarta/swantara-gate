package test

import (
	"net/http"
	"testing"
)

// =========================================================
// TEST AUTH - LOGIN
// =========================================================

func TestAuthLogin(t *testing.T) {
	ts := SetupTestServer(t)
	defer ts.Close()

	// Create test user
	ts.CreateTestUser(t, "admin_test", "password123", "admin")

	t.Run("Login success", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/auth/login", map[string]interface{}{
			"username": "admin_test",
			"password": "password123",
		})
		AssertStatus(t, resp, http.StatusOK)
		res := AssertSuccess(t, body)
		if res.Data["access_token"] == nil {
			t.Fatal("Expected access_token in response")
		}
		if res.Data["refresh_token"] == nil {
			t.Fatal("Expected refresh_token in response")
		}
		if res.Data["token_type"] != "Bearer" {
			t.Errorf("Expected token_type=Bearer, got %v", res.Data["token_type"])
		}
	})

	t.Run("Login - wrong password", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/auth/login", map[string]interface{}{
			"username": "admin_test",
			"password": "wrongpassword",
		})
		AssertStatus(t, resp, http.StatusUnauthorized)
		AssertError(t, body)
	})

	t.Run("Login - user not found", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/auth/login", map[string]interface{}{
			"username": "nonexistent",
			"password": "password123",
		})
		AssertStatus(t, resp, http.StatusUnauthorized)
		AssertError(t, body)
	})

	t.Run("Login - missing username", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/auth/login", map[string]interface{}{
			"password": "password123",
		})
		AssertStatus(t, resp, http.StatusUnauthorized)
		AssertError(t, body)
	})

	t.Run("Login - missing password", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/auth/login", map[string]interface{}{
			"username": "admin_test",
		})
		AssertStatus(t, resp, http.StatusUnauthorized)
		AssertError(t, body)
	})

	t.Run("Login - inactive user", func(t *testing.T) {
		// Create inactive user
		ts.DoRequest("POST", "/api/admin/users", map[string]interface{}{
			"username":  "inactive_user",
			"password":  "password123",
			"role":      "admin",
			"is_active": false,
		})
		resp, body := ts.DoRequest("POST", "/api/admin/auth/login", map[string]interface{}{
			"username": "inactive_user",
			"password": "password123",
		})
		AssertStatus(t, resp, http.StatusUnauthorized)
		AssertError(t, body)
	})
}

// =========================================================
// TEST AUTH - TOKEN REFRESH
// =========================================================

func TestAuthRefreshToken(t *testing.T) {
	ts := SetupTestServer(t)
	defer ts.Close()

	// Create user and login
	ts.CreateTestUser(t, "refresh_user", "password123", "admin")
	_, refreshToken := ts.LoginTestUser(t, "refresh_user", "password123")

	t.Run("Refresh token success", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/auth/refresh", map[string]interface{}{
			"refresh_token": refreshToken,
		})
		AssertStatus(t, resp, http.StatusOK)
		res := AssertSuccess(t, body)
		if res.Data["access_token"] == nil {
			t.Fatal("Expected new access_token")
		}
		if res.Data["refresh_token"] == nil {
			t.Fatal("Expected new refresh_token")
		}
	})

	t.Run("Refresh - old token revoked after use", func(t *testing.T) {
		// The old refresh token should be revoked after first use
		resp, body := ts.DoRequest("POST", "/api/admin/auth/refresh", map[string]interface{}{
			"refresh_token": refreshToken,
		})
		AssertStatus(t, resp, http.StatusUnauthorized)
		AssertError(t, body)
	})

	t.Run("Refresh - invalid token", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/auth/refresh", map[string]interface{}{
			"refresh_token": "invalid-token-string",
		})
		AssertStatus(t, resp, http.StatusUnauthorized)
		AssertError(t, body)
	})

	t.Run("Refresh - empty token", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/auth/refresh", map[string]interface{}{
			"refresh_token": "",
		})
		AssertStatus(t, resp, http.StatusUnauthorized)
		AssertError(t, body)
	})
}

// =========================================================
// TEST AUTH - LOGOUT
// =========================================================

func TestAuthLogout(t *testing.T) {
	ts := SetupTestServer(t)
	defer ts.Close()

	// Create user and login
	ts.CreateTestUser(t, "logout_user", "password123", "admin")
	_, refreshToken := ts.LoginTestUser(t, "logout_user", "password123")

	t.Run("Logout success", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/auth/logout", map[string]interface{}{
			"refresh_token": refreshToken,
		})
		AssertStatus(t, resp, http.StatusOK)
		AssertSuccess(t, body)
	})

	t.Run("Refresh after logout should fail", func(t *testing.T) {
		resp, body := ts.DoRequest("POST", "/api/admin/auth/refresh", map[string]interface{}{
			"refresh_token": refreshToken,
		})
		AssertStatus(t, resp, http.StatusUnauthorized)
		AssertError(t, body)
	})
}

// =========================================================
// TEST AUTH - PROTECTED ROUTE (GET /api/admin/auth/me)
// =========================================================

func TestAuthMe(t *testing.T) {
	ts := SetupTestServer(t)
	defer ts.Close()

	// Create user and login
	ts.CreateTestUser(t, "me_user", "password123", "super_admin")
	accessToken, _ := ts.LoginTestUser(t, "me_user", "password123")

	t.Run("Get profile with valid token", func(t *testing.T) {
		resp, body := ts.DoAuthRequest("GET", "/api/admin/auth/me", nil, accessToken)
		AssertStatus(t, resp, http.StatusOK)
		res := AssertSuccess(t, body)
		if res.Data["username"] != "me_user" {
			t.Errorf("Expected username=me_user, got %v", res.Data["username"])
		}
		if res.Data["role"] != "super_admin" {
			t.Errorf("Expected role=super_admin, got %v", res.Data["role"])
		}
	})

	t.Run("Get profile without token", func(t *testing.T) {
		resp, body := ts.DoAuthRequest("GET", "/api/admin/auth/me", nil, "")
		AssertStatus(t, resp, http.StatusUnauthorized)
		AssertError(t, body)
	})

	t.Run("Get profile with invalid token", func(t *testing.T) {
		resp, body := ts.DoAuthRequest("GET", "/api/admin/auth/me", nil, "invalid-token")
		AssertStatus(t, resp, http.StatusUnauthorized)
		AssertError(t, body)
	})

	t.Run("Get profile with malformed auth header", func(t *testing.T) {
		// Send request with wrong format
		req, _ := http.NewRequest("GET", ts.Server.URL+"/api/admin/auth/me", nil)
		req.Header.Set("Authorization", "Token "+accessToken) // wrong prefix
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusUnauthorized {
			t.Errorf("Expected 401, got %d", resp.StatusCode)
		}
	})
}

// =========================================================
// TEST AUTH - RATE LIMITING
// =========================================================

func TestAuthRateLimiting(t *testing.T) {
	ts := SetupTestServer(t)
	defer ts.Close()

	t.Run("Rate limit after multiple failed attempts", func(t *testing.T) {
		// Test server has rate limit of 50 per 60 seconds
		// Make 50 attempts to exhaust the limit
		for i := 0; i < 50; i++ {
			ts.DoRequest("POST", "/api/admin/auth/login", map[string]interface{}{
				"username": "nobody",
				"password": "wrong",
			})
		}

		// 51st attempt should be rate limited
		resp, body := ts.DoRequest("POST", "/api/admin/auth/login", map[string]interface{}{
			"username": "nobody",
			"password": "wrong",
		})
		AssertStatus(t, resp, http.StatusTooManyRequests)
		AssertError(t, body)
	})
}
