package mocks

import (
	"context"

	"github.com/mert-yigittop/cxp-api-starter/internal/todo/dto"
	"github.com/mert-yigittop/cxp-api-starter/internal/todo/entity"
	"github.com/stretchr/testify/mock"
)

type Usecase struct {
	mock.Mock
}

func (m *Usecase) GetList(ctx context.Context, req dto.GetTodoListRequest) (dto.GetTodoListResponse, int, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(dto.GetTodoListResponse), args.Int(1), args.Error(2)
}

func (m *Usecase) Create(ctx context.Context, todo entity.Todo) (dto.CreateTodoResponse, int, error) {
	args := m.Called(ctx, todo)
	return args.Get(0).(dto.CreateTodoResponse), args.Int(1), args.Error(2)
}

func (m *Usecase) Update(ctx context.Context, req dto.UpdateTodoRequest, userID uint) (dto.UpdateTodoResponse, int, error) {
	args := m.Called(ctx, req, userID)
	return args.Get(0).(dto.UpdateTodoResponse), args.Int(1), args.Error(2)
}

func (m *Usecase) Delete(ctx context.Context, req dto.DeleteTodoRequest, userID uint) (dto.DeleteTodoResponse, int, error) {
	args := m.Called(ctx, req, userID)
	return args.Get(0).(dto.DeleteTodoResponse), args.Int(1), args.Error(2)
}
