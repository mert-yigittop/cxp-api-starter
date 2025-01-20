package mocks

import (
	"context"
	"github.com/mert-yigittop/cxp-api-starter/internal/user/dto"
	"github.com/mert-yigittop/cxp-api-starter/internal/user/entity"
	"github.com/stretchr/testify/mock"
)

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
