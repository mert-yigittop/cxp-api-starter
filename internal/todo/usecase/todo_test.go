package usecase

import (
	"context"
	"github.com/mert-yigittop/cxp-api-starter/internal/todo/dto"
	"github.com/mert-yigittop/cxp-api-starter/internal/todo/entity"
	"github.com/mert-yigittop/cxp-api-starter/internal/todo/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"testing"
)

func TestUsecase_GetList(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	usecase := New(mockRepo)

	todoList := dto.GetTodoListResponse{
		Todos: []entity.Todo{
			{Content: "Test Todo 1", Completed: false, UserID: 1},
			{Content: "Test Todo 2", Completed: false, UserID: 1},
		},
	}

	mockRepo.On("GetList", mock.Anything, mock.Anything).Return(todoList, nil)

	payload := dto.GetTodoListRequest{UserID: 1}

	response, status, err := usecase.GetList(context.Background(), payload)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, status)
	assert.Equal(t, 2, len(response.Todos))
	mockRepo.AssertExpectations(t)
}

func TestUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	usecase := New(mockRepo)

	todo := entity.Todo{
		UserID:  1,
		Content: "Test Todo",
	}

	createResponse := dto.CreateTodoResponse{
		ID:        1,
		Content:   "Test Todo",
		Completed: false,
	}

	mockRepo.On("Create", mock.Anything, mock.Anything).Return(createResponse, nil)

	response, status, err := usecase.Create(context.Background(), todo)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, status)
	assert.Equal(t, createResponse.ID, response.ID)
	mockRepo.AssertExpectations(t)
}

func TestUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	usecase := New(mockRepo)

	updateRequest := dto.UpdateTodoRequest{
		ID:        1,
		Content:   "Updated Todo",
		Completed: true,
	}

	updateResponse := dto.UpdateTodoResponse{
		ID:        1,
		Content:   "Updated Todo",
		Completed: true,
	}

	mockRepo.On("Update", mock.Anything, mock.Anything, uint(1)).Return(updateResponse, nil)

	response, status, err := usecase.Update(context.Background(), updateRequest, 1)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, status)
	assert.Equal(t, updateResponse.Content, response.Content)
	mockRepo.AssertExpectations(t)
}

func TestUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	usecase := New(mockRepo)

	deleteRequest := dto.DeleteTodoRequest{ID: 1}
	deleteResponse := dto.DeleteTodoResponse{ID: 1}

	mockRepo.On("Delete", mock.Anything, mock.Anything, uint(1)).Return(deleteResponse, nil)

	response, status, err := usecase.Delete(context.Background(), deleteRequest, 1)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, status)
	assert.Equal(t, deleteResponse.ID, response.ID)
	mockRepo.AssertExpectations(t)
}
