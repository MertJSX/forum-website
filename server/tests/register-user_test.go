package tests

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/MertJSX/forum-website/server/routes"
	"github.com/MertJSX/forum-website/server/types"
	"github.com/gofiber/fiber/v2"
)

func TestHandleRegisterUser(t *testing.T) {
	app := fiber.New()
	var db *sql.DB = GetTestDB() // Mock or use a test database connection
	app.Post("/register-user", func(c *fiber.Ctx) error {
		return routes.HandleRegisterUser(c, db)
	})

	t.Run("User Registration", func(t *testing.T) {
		requestBody := types.User{
			Name:     "testuser",
			Email:    "testuser@example.com",
			Password: "password123",
		}
		body, _ := json.Marshal(requestBody)
		req := httptest.NewRequest(http.MethodPost, "/register-user", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("Register existing user", func(t *testing.T) {
		requestBody := types.User{
			Name:     "testuser",
			Email:    "testuser@example.com",
			Password: "password123",
		}
		body, _ := json.Marshal(requestBody)
		req := httptest.NewRequest(http.MethodPost, "/register-user", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

		var errorResponse types.ErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&errorResponse); err != nil {
			t.Fatalf("Failed to parse response body: %v", err)
		}

		assert.True(t, errorResponse.ErrorMsg == "Email already exists" || errorResponse.ErrorMsg == "Username already exists")

	})

	t.Run("Register user with missing email", func(t *testing.T) {
		requestBody := types.User{
			Name:     "registeruserwithmissingemail",
			Email:    "",
			Password: "password123",
		}
		body, _ := json.Marshal(requestBody)
		req := httptest.NewRequest(http.MethodPost, "/register-user", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

		var errorResponse types.ErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&errorResponse); err != nil {
			t.Fatalf("Failed to parse response body: %v", err)
		}

		assert.Equal(t, "Missing required fields", errorResponse.ErrorMsg)
	})

	t.Run("Register user with missing username", func(t *testing.T) {
		requestBody := types.User{
			Name:     "",
			Email:    "testuser@example.com",
			Password: "password123",
		}
		body, _ := json.Marshal(requestBody)
		req := httptest.NewRequest(http.MethodPost, "/register-user", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

		var errorResponse types.ErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&errorResponse); err != nil {
			t.Fatalf("Failed to parse response body: %v", err)
		}

		assert.Equal(t, "Missing required fields", errorResponse.ErrorMsg)
	})

	t.Run("Bad Request", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/register-user", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})
}
