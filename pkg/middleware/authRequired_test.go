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
	app := fiber.New()

	mockJWTService := new(MockJWTService)

	app.Use(func(c *fiber.Ctx) error {
		return AuthRequired()(c)
	})

	t.Run("Valid Token", func(t *testing.T) {
		mockJWTService.On("Verify", "valid_token").Return(uint(123), nil)

		req := httptest.NewRequest("GET", "/protected", nil)
		req.AddCookie(&http.Cookie{Name: "jwt", Value: "valid_token"})
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		mockJWTService.AssertExpectations(t)
	})

	t.Run("Invalid Token", func(t *testing.T) {
		mockJWTService.On("Verify", "invalid_token").Return(uint(0), fmt.Errorf("invalid token"))

		req := httptest.NewRequest("GET", "/protected", nil)
		req.AddCookie(&http.Cookie{Name: "jwt", Value: "invalid_token"})
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
		mockJWTService.AssertExpectations(t)
	})

	t.Run("Missing Token", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/protected", nil)
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	})
}
