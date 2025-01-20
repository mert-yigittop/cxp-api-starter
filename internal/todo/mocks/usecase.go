package mocks

import (
	"context"
	"github.com/mert-yigittop/cxp-api-starter/internal/todo/dto"
	"github.com/mert-yigittop/cxp-api-starter/internal/todo/entity"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) GetList(ctx context.Context, payload dto.GetTodoListRequest) (dto.GetTodoListResponse, error) {
	args := m.Called(ctx, payload)
	return args.Get(0).(dto.GetTodoListResponse), args.Error(1)
}

func (m *MockRepository) Create(ctx context.Context, payload entity.Todo) (dto.CreateTodoResponse, error) {
	args := m.Called(ctx, payload)
	return args.Get(0).(dto.CreateTodoResponse), args.Error(1)
}

func (m *MockRepository) Update(ctx context.Context, payload dto.UpdateTodoRequest, userId uint) (dto.UpdateTodoResponse, error) {
	args := m.Called(ctx, payload, userId)
	return args.Get(0).(dto.UpdateTodoResponse), args.Error(1)
}

func (m *MockRepository) Delete(ctx context.Context, payload dto.DeleteTodoRequest, userId uint) (dto.DeleteTodoResponse, error) {
	args := m.Called(ctx, payload, userId)
	return args.Get(0).(dto.DeleteTodoResponse), args.Error(1)
}
