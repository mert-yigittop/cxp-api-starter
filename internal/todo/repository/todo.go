package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/mert-yigittop/cxp-api-starter/internal/todo/dto"
	"github.com/mert-yigittop/cxp-api-starter/internal/todo/entity"
	"gorm.io/gorm"
)

type Repository interface {
	GetList(ctx context.Context, payload dto.GetTodoListRequest) (dto.GetTodoListResponse, error)
	Create(ctx context.Context, todo entity.Todo) (dto.CreateTodoResponse, error)
	Update(ctx context.Context, todo dto.UpdateTodoRequest, userId uint) (dto.UpdateTodoResponse, error)
	Delete(ctx context.Context, payload dto.DeleteTodoRequest, userId uint) (dto.DeleteTodoResponse, error)
}

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository { return &repository{db: db} }

func (repo *repository) GetList(ctx context.Context, payload dto.GetTodoListRequest) (dto.GetTodoListResponse, error) {
	var todos []entity.Todo
	err := repo.db.Where("user_id = ?", payload.UserID).Find(&todos).Error
	if err != nil {
		return dto.GetTodoListResponse{}, err
	}

	return dto.GetTodoListResponse{
		Todos: todos,
	}, nil

}

func (repo *repository) Create(ctx context.Context, todo entity.Todo) (dto.CreateTodoResponse, error) {
	err := repo.db.Create(&todo).Error
	if err != nil {
		return dto.CreateTodoResponse{}, err
	}

	return dto.CreateTodoResponse{
		ID:        todo.ID,
		Content:   todo.Content,
		Completed: todo.Completed,
	}, nil
}

func (repo *repository) Update(ctx context.Context, payload dto.UpdateTodoRequest, userId uint) (dto.UpdateTodoResponse, error) {
	var todo entity.Todo
	if err := repo.db.First(&todo, payload.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.UpdateTodoResponse{}, err
		}
		return dto.UpdateTodoResponse{}, err
	}

	if todo.UserID != userId {
		return dto.UpdateTodoResponse{}, fmt.Errorf("You are not authorized to update this todo")
	}

	err := repo.db.Model(&todo).Updates(payload).Error
	if err != nil {
		return dto.UpdateTodoResponse{}, err
	}

	return dto.UpdateTodoResponse{
		ID:        todo.ID,
		Content:   todo.Content,
		Completed: todo.Completed,
	}, nil
}

func (repo *repository) Delete(ctx context.Context, payload dto.DeleteTodoRequest, userId uint) (dto.DeleteTodoResponse, error) {
	var todo entity.Todo
	if err := repo.db.First(&todo, payload.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.DeleteTodoResponse{}, err
		}
		return dto.DeleteTodoResponse{}, err
	}

	if todo.UserID != userId {
		return dto.DeleteTodoResponse{}, fmt.Errorf("You are not authorized to update this todo")
	}

	err := repo.db.Delete(&todo).Error
	if err != nil {
		return dto.DeleteTodoResponse{}, err
	}

	return dto.DeleteTodoResponse{
		ID: payload.ID,
	}, nil
}
