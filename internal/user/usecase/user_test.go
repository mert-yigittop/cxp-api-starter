package usecase

import (
	"context"
	"errors"
	"github.com/mert-yigittop/cxp-api-starter/internal/user/dto"
	"github.com/mert-yigittop/cxp-api-starter/internal/user/entity"
	"github.com/mert-yigittop/cxp-api-starter/internal/user/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"testing"
)

func TestUsecase_Register_Success(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	uc := New(mockRepo)

	request := dto.RegisterRequest{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123",
	}

	user := entity.New(request)
	expectedResponse := dto.RegisterResponse{
		ID: 1,
	}

	mockRepo.On("Register", mock.Anything, user).Return(expectedResponse, nil)

	response, status, err := uc.Register(context.Background(), request)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, status)
	assert.Equal(t, expectedResponse, response)
	mockRepo.AssertCalled(t, "Register", mock.Anything, user)
}

func TestUsecase_Register_Error(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	uc := New(mockRepo)

	request := dto.RegisterRequest{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123",
	}

	user := entity.New(request)
	mockRepo.On("Register", mock.Anything, user).Return(dto.RegisterResponse{}, errors.New("registration failed"))

	response, status, err := uc.Register(context.Background(), request)
	assert.Error(t, err)
	assert.Equal(t, http.StatusBadRequest, status)
	assert.Equal(t, dto.RegisterResponse{}, response)
}

func TestUsecase_Login_Success(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	uc := New(mockRepo)

	request := dto.LoginRequest{
		Username: "testuser",
		Password: "password123",
	}

	mockRepo.On("Login", mock.Anything, request).Return(uint(1), nil)

	response, status, err := uc.Login(context.Background(), request)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, status)
	assert.Equal(t, dto.LoginResponse{
		ID:          1,
		AccessToken: "mockToken",
	}, response)
}

func TestUsecase_Login_Invalid_Token(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	uc := New(mockRepo)

	request := dto.LoginRequest{
		Username: "testuser",
		Password: "wrongpassword",
	}

	mockRepo.On("Login", mock.Anything, request).Return(uint(0), errors.New("invalid credentials"))

	response, status, err := uc.Login(context.Background(), request)
	assert.Error(t, err)
	assert.Equal(t, http.StatusBadRequest, status)
	assert.Equal(t, dto.LoginResponse{}, response)
}

func TestUsecase_Login_JWT_Error(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	uc := New(mockRepo)

	request := dto.LoginRequest{
		Username: "testuser",
		Password: "password123",
	}

	mockRepo.On("Login", mock.Anything, request).Return(uint(1), nil)

	response, status, err := uc.Login(context.Background(), request)
	assert.Error(t, err)
	assert.Equal(t, http.StatusUnauthorized, status)
	assert.Equal(t, dto.LoginResponse{}, response)
}
