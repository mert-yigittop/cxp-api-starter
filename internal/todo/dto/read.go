package dto

import "github.com/mert-yigittop/cxp-api-starter/internal/todo/entity"

type TodoResponse struct {
	ID        uint   `json:"id"`
	Content   string `json:"content"`
	Completed bool   `json:"completed"`
	UpdatedAt string `json:"updated_at"`
}

type GetTodoListRequest struct {
	UserID uint `json:"id"`
}

type GetTodoListResponse struct {
	Todos []entity.Todo `json:"todos"`
}
