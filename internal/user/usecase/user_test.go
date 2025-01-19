package usecase

import (
	"context"
	"errors"
	"github.com/mert-yigittop/cxp-api-starter/internal/user/dto"
	"github.com/mert-yigittop/cxp-api-starter/internal/user/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"testing"
)

// MockRepository is a mock implementation of the Repository interface
type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) Register(ctx context.Context, user entity.User) (dto.RegisterResponse, error) {
	args := m.Called(ctx, user)
	return args.Get(0).(dto.RegisterResponse), args.Error(1)
}

func (m *MockRepository) Login(ctx context.Context, payload dto.LoginRequest) (uint, error) {
	args := m.Called(ctx, payload)
	return args.Get(0).(uint), args.Error(1)
}

func TestUsecase(t *testing.T) {
	mockRepo := new(MockRepository)
	uc := New(mockRepo)

	t.Run("Register - Success", func(t *testing.T) {
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
	})

	t.Run("Register - Error", func(t *testing.T) {
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
	})

	t.Run("Login - Success", func(t *testing.T) {
		request := dto.LoginRequest{
			Username: "testuser",
			Password: "password123",
		}

		mockRepo.On("Login", mock.Anything, request).Return(uint(1), nil)

		// Mock JWT sign function
		//jwt.Sign = func(userId uint, duration time.Duration) (string, error) {
		//	return "mockToken", nil
		//}

		response, status, err := uc.Login(context.Background(), request)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, status)
		assert.Equal(t, dto.LoginResponse{
			ID:          1,
			AccessToken: "mockToken",
		}, response)
	})

	t.Run("Login - Invalid Credentials", func(t *testing.T) {
		request := dto.LoginRequest{
			Username: "testuser",
			Password: "wrongpassword",
		}

		mockRepo.On("Login", mock.Anything, request).Return(uint(0), errors.New("invalid credentials"))

		response, status, err := uc.Login(context.Background(), request)
		assert.Error(t, err)
		assert.Equal(t, http.StatusBadRequest, status)
		assert.Equal(t, dto.LoginResponse{}, response)
	})

	t.Run("Login - JWT Error", func(t *testing.T) {
		request := dto.LoginRequest{
			Username: "testuser",
			Password: "password123",
		}

		mockRepo.On("Login", mock.Anything, request).Return(uint(1), nil)

		// Mock JWT sign function to return an error
		//jwt.Sign = func(userId uint, duration time.Duration) (string, error) {
		//	return "", errors.New("jwt error")
		//}

		response, status, err := uc.Login(context.Background(), request)
		assert.Error(t, err)
		assert.Equal(t, http.StatusUnauthorized, status)
		assert.Equal(t, dto.LoginResponse{}, response)
	})
}
