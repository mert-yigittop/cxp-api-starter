package v1

import (
	"bytes"
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/mert-yigittop/cxp-api-starter/internal/todo/dto"
	"github.com/mert-yigittop/cxp-api-starter/internal/todo/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockUsecase struct {
	mock.Mock
}

func (m *MockUsecase) GetList(ctx context.Context, payload dto.GetTodoListRequest) (dto.GetTodoListResponse, int, error) {
	args := m.Called(ctx, payload)
	return args.Get(0).(dto.GetTodoListResponse), args.Int(1), args.Error(2)
}

func (m *MockUsecase) Create(ctx context.Context, payload entity.Todo) (dto.CreateTodoResponse, int, error) {
	args := m.Called(ctx, payload)
	return args.Get(0).(dto.CreateTodoResponse), args.Int(1), args.Error(2)
}

func (m *MockUsecase) Update(ctx context.Context, payload dto.UpdateTodoRequest, userId uint) (dto.UpdateTodoResponse, int, error) {
	args := m.Called(ctx, payload, userId)
	return args.Get(0).(dto.UpdateTodoResponse), args.Int(1), args.Error(2)
}

func (m *MockUsecase) Delete(ctx context.Context, payload dto.DeleteTodoRequest, userId uint) (dto.DeleteTodoResponse, int, error) {
	args := m.Called(ctx, payload, userId)
	return args.Get(0).(dto.DeleteTodoResponse), args.Int(1), args.Error(2)
}

func TestHandler_GetList(t *testing.T) {
	mockUsecase := new(MockUsecase)
	handler := New(mockUsecase)

	todoList := dto.GetTodoListResponse{
		Todos: []entity.Todo{
			{Content: "Test Todo 1", Completed: false, UserID: 1},
			{Content: "Test Todo 2", Completed: true, UserID: 1},
		},
	}

	mockUsecase.On("GetList", mock.Anything, mock.Anything).Return(todoList, http.StatusOK, nil)

	app := fiber.New()

	app.Get("/todos", handler.GetList)

	req := httptest.NewRequest("GET", "/todos", nil)
	req.Header.Set("Authorization", "Bearer token")
	req.Cookies("jwt", "dummy_token")

	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	mockUsecase.AssertExpectations(t)
}

func TestHandler_Create(t *testing.T) {
	mockUsecase := new(MockUsecase)
	handler := New(mockUsecase)

	payload := dto.CreateTodoRequest{
		Content: "New Todo",
	}

	createResponse := dto.CreateTodoResponse{
		ID:        1,
		Content:   "New Todo",
		Completed: false,
	}

	mockUsecase.On("Create", mock.Anything, mock.Anything).Return(createResponse, http.StatusOK, nil)

	app := fiber.New()

	app.Post("/todos", handler.Create)

	reqBody := []byte(`{"content":"New Todo"}`)
	req := httptest.NewRequest("POST", "/todos", bytes.NewReader(reqBody))
	req.Cookies("jwt", "dummy_token")

	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	mockUsecase.AssertExpectations(t)
}

func TestHandler_Update(t *testing.T) {
	mockUsecase := new(MockUsecase)
	handler := New(mockUsecase)

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

	mockUsecase.On("Update", mock.Anything, mock.Anything, uint(1)).Return(updateResponse, http.StatusOK, nil)

	app := fiber.New()

	app.Put("/todos/:id", handler.Update)

	reqBody := []byte(`{"content":"Updated Todo","completed":true}`)
	req := httptest.NewRequest("PUT", "/todos/1", bytes.NewReader(reqBody))
	req.Cookies("jwt", "dummy_token")

	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	mockUsecase.AssertExpectations(t)
}

func TestHandler_Delete(t *testing.T) {
	mockUsecase := new(MockUsecase)
	handler := New(mockUsecase)

	deleteResponse := dto.DeleteTodoResponse{
		ID: 1,
	}

	mockUsecase.On("Delete", mock.Anything, mock.Anything, uint(1)).Return(deleteResponse, http.StatusOK, nil)

	app := fiber.New()

	app.Delete("/todos/:id", handler.Delete)

	req := httptest.NewRequest("DELETE", "/todos/1", nil)
	req.Cookies("jwt", "dummy_token")

	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	mockUsecase.AssertExpectations(t)
}
