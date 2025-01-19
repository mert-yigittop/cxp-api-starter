package usecase

import (
	"context"
	"github.com/mert-yigittop/cxp-api-starter/internal/todo/dto"
	"github.com/mert-yigittop/cxp-api-starter/internal/todo/entity"
	"github.com/mert-yigittop/cxp-api-starter/internal/todo/repository"
	"net/http"
)

type Usecase interface {
	GetList(ctx context.Context, payload dto.GetTodoListRequest) (dto.GetTodoListResponse, int, error)
	Create(ctx context.Context, payload entity.Todo) (dto.CreateTodoResponse, int, error)
	Update(ctx context.Context, payload dto.UpdateTodoRequest, userId uint) (dto.UpdateTodoResponse, int, error)
	Delete(ctx context.Context, payload dto.DeleteTodoRequest, userId uint) (dto.DeleteTodoResponse, int, error)
}

type usecase struct {
	repo repository.Repository
}

func New(repo repository.Repository) Usecase {
	return &usecase{
		repo: repo,
	}
}

func (uc *usecase) GetList(ctx context.Context, payload dto.GetTodoListRequest) (dto.GetTodoListResponse, int, error) {
	response, err := uc.repo.GetList(ctx, payload)
	if err != nil {
		return dto.GetTodoListResponse{}, http.StatusBadRequest, err
	}

	return response, http.StatusOK, nil
}

func (uc *usecase) Create(ctx context.Context, payload entity.Todo) (dto.CreateTodoResponse, int, error) {
	response, err := uc.repo.Create(ctx, payload)
	if err != nil {
		return dto.CreateTodoResponse{}, http.StatusBadRequest, err
	}

	return response, http.StatusOK, nil
}

func (uc *usecase) Update(ctx context.Context, payload dto.UpdateTodoRequest, userId uint) (dto.UpdateTodoResponse, int, error) {
	response, err := uc.repo.Update(ctx, payload, userId)
	if err != nil {
		return dto.UpdateTodoResponse{}, http.StatusBadRequest, err
	}

	return response, http.StatusOK, nil
}

func (uc *usecase) Delete(ctx context.Context, payload dto.DeleteTodoRequest, userId uint) (dto.DeleteTodoResponse, int, error) {
	response, err := uc.repo.Delete(ctx, payload, userId)
	if err != nil {
		return dto.DeleteTodoResponse{}, http.StatusBadRequest, err
	}

	return response, http.StatusOK, nil
}
