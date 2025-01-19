package middleware

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockJWTService struct {
	mock.Mock
}

func (m *MockJWTService) Verify(token string) (uint, error) {
	args := m.Called(token)
	return args.Get(0).(uint), args.Error(1)
}

func TestAuthRequired(t *testing.T) {
	// Create a new fiber app
	app := fiber.New()

	// Create a mock JWT service
	mockJWTService := new(MockJWTService)

	// Middleware'ı kullanacak bir route oluştur
	app.Use(func(c *fiber.Ctx) error {
		// Mock servis kullanarak AuthRequired fonksiyonunu çalıştırıyoruz
		return AuthRequired()(c)
	})

	// Test valid token
	t.Run("Valid Token", func(t *testing.T) {
		// Mock the Verify function to return a valid user ID
		mockJWTService.On("Verify", "valid_token").Return(uint(123), nil)

		// Create a mock request with a valid JWT cookie
		req := httptest.NewRequest("GET", "/protected", nil)
		req.AddCookie(&http.Cookie{Name: "jwt", Value: "valid_token"})
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		mockJWTService.AssertExpectations(t)
	})

	// Test invalid token
	t.Run("Invalid Token", func(t *testing.T) {
		// Mock the Verify function to return an error for an invalid token
		mockJWTService.On("Verify", "invalid_token").Return(uint(0), fmt.Errorf("invalid token"))

		// Create a mock request with an invalid JWT cookie
		req := httptest.NewRequest("GET", "/protected", nil)
		req.AddCookie(&http.Cookie{Name: "jwt", Value: "invalid_token"})
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
		mockJWTService.AssertExpectations(t)
	})

	// Test missing token
	t.Run("Missing Token", func(t *testing.T) {
		// Create a mock request without JWT cookie
		req := httptest.NewRequest("GET", "/protected", nil)
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	})
}
