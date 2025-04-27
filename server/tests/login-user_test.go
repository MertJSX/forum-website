package tests

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MertJSX/forum-website/server/routes"
	"github.com/MertJSX/forum-website/server/types"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestHandleLoginUser(t *testing.T) {
	app := fiber.New()
	var db *sql.DB = GetTestDB() // Mock or use a test database connection
	app.Post("/login-user", func(c *fiber.Ctx) error {
		return routes.HandleLoginUser(c, db)
	})

	t.Run("User Login", func(t *testing.T) {
		requestBody := types.User{
			Name:     "exampleexistinguser",
			Email:    "exampleexistingemail@gmail.com",
			Password: "password123",
		}
		body, _ := json.Marshal(requestBody)
		req := httptest.NewRequest(http.MethodPost, "/login-user", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var response types.LoginResponse
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			t.Fatalf("Failed to parse response body: %v", err)
		}
		assert.IsType(t, response, types.LoginResponse{}, "Expected response type to be LoginResponse")
		assert.NotEmpty(t, response.Token, "Expected token to be present in the response")
	})

	t.Run("User login wrong password", func(t *testing.T) {
		requestBody := types.User{
			Name:     "exampleexistinguser",
			Email:    "exampleexistingemail@gmail.com",
			Password: "password1234",
		}
		body, _ := json.Marshal(requestBody)
		req := httptest.NewRequest(http.MethodPost, "/login-user", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var errResponse types.ErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&errResponse); err != nil {
			t.Fatalf("Failed to parse response body: %v", err)
		}

		assert.True(t, errResponse.ErrorMsg == "Email or password is wrong" || errResponse.ErrorMsg == "Username or password is wrong")
	})
}
