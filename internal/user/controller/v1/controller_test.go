package v1

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/mert-yigittop/cxp-api-starter/internal/user/dto"
	"github.com/mert-yigittop/cxp-api-starter/internal/user/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_Register(t *testing.T) {
	mockUsecase := new(mocks.Usecase)
	handler := New(mockUsecase)

	app := fiber.New()
	app.Post("/register", handler.Register)

	payload := dto.RegisterRequest{
		Username: "testuser",
		Password: "password123",
		Email:    "test@example.com",
	}
	mockResponse := dto.RegisterResponse{ID: 1}

	mockUsecase.On("Register", mock.Anything, payload).Return(mockResponse, http.StatusOK, nil)

	body, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req, -1)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	mockUsecase.AssertCalled(t, "Register", mock.Anything, payload)
	mockUsecase.AssertExpectations(t) // Mock beklentilerini kontrol et

}

func TestHandler_Login(t *testing.T) {
	mockUsecase := new(mocks.Usecase)
	handler := New(mockUsecase)

	app := fiber.New()
	app.Post("/login", handler.Login)

	payload := dto.LoginRequest{
		Username: "testuser",
		Password: "password123",
	}
	mockResponse := dto.LoginResponse{
		ID:          1,
		AccessToken: "mockAccessToken",
	}

	mockUsecase.On("Login", mock.Anything, payload).Return(mockResponse, http.StatusOK, nil)

	body, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req, -1)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	mockUsecase.AssertCalled(t, "Login", mock.Anything, payload)
	mockUsecase.AssertExpectations(t) // Mock beklentilerini kontrol et

}

func TestHandler_Logout(t *testing.T) {
	mockUsecase := new(mocks.Usecase)
	handler := New(mockUsecase)

	app := fiber.New()

	app.Post("/logout", handler.Logout)

	req := httptest.NewRequest(http.MethodPost, "/logout", nil)

	resp, _ := app.Test(req, -1)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
