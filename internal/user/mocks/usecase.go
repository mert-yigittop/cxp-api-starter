package mocks

import (
	"context"
	"github.com/mert-yigittop/cxp-api-starter/internal/user/dto"
	"github.com/stretchr/testify/mock"
)

type Usecase struct {
	mock.Mock
}

func (m *Usecase) Register(ctx context.Context, payload dto.RegisterRequest) (dto.RegisterResponse, int, error) {
	args := m.Called(ctx, payload)
	return args.Get(0).(dto.RegisterResponse), args.Int(1), args.Error(2)
}

func (m *Usecase) Login(ctx context.Context, payload dto.LoginRequest) (dto.LoginResponse, int, error) {
	args := m.Called(ctx, payload)
	return args.Get(0).(dto.LoginResponse), args.Int(1), args.Error(2)
}
