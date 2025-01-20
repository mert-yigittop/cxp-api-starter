package v1

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/mert-yigittop/cxp-api-starter/internal/todo/dto"
	"github.com/mert-yigittop/cxp-api-starter/internal/todo/entity"
	"github.com/mert-yigittop/cxp-api-starter/internal/todo/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func setupApp(mockUsecase *mocks.Usecase) *fiber.App {
	app := fiber.New()

	// Register routes
	handler := New(mockUsecase)
	app.Get("/todos", handler.GetList)
	app.Post("/todos", handler.Create)
	app.Put("/todos/:id", handler.Update)
	app.Delete("/todos/:id", handler.Delete)

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("userId", uint(1))
		c.Cookie(&fiber.Cookie{
			Name:    "jwt",
			Value:   "mockToken",
			Expires: time.Now().Add(time.Hour * 24),
		})
		return c.Next()
	})

	return app
}

func TestGetListSuccess(t *testing.T) {
	mockUsecase := new(mocks.Usecase)
	app := setupApp(mockUsecase)

	mockResponse := dto.GetTodoListResponse{
		Todos: []entity.Todo{
			{Content: "Test Todo", UserID: 1, Completed: false},
		},
	}

	mockUsecase.On("GetList", mock.Anything, mock.Anything).Return(mockResponse, http.StatusOK, nil)

	req := httptest.NewRequest(http.MethodGet, "/todos", nil)
	resp, _ := app.Test(req, -1)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	mockUsecase.AssertCalled(t, "GetList", mock.Anything, mock.Anything)
}

func TestCreateSuccess(t *testing.T) {
	mockUsecase := new(mocks.Usecase)
	app := setupApp(mockUsecase)

	payload := dto.CreateTodoRequest{Content: "New Todo"}
	mockResponse := dto.CreateTodoResponse{
		ID:        1,
		Content:   "New Todo",
		Completed: false,
	}

	mockUsecase.On("Create", mock.Anything, mock.Anything).Return(mockResponse, http.StatusCreated, nil)

	body, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/todos", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req, -1)

	assert.Equal(t, http.StatusCreated, resp.StatusCode)
	mockUsecase.AssertCalled(t, "Create", mock.Anything, mock.Anything)
}

func TestUpdateSuccess(t *testing.T) {
	mockUsecase := new(mocks.Usecase)
	app := setupApp(mockUsecase)

	payload := dto.UpdateTodoRequest{Content: "Updated Todo", Completed: true}
	mockResponse := dto.UpdateTodoResponse{
		ID:        1,
		Content:   "Updated Todo",
		Completed: true,
	}

	mockUsecase.On("Update", mock.Anything, mock.Anything, uint(1)).Return(mockResponse, http.StatusOK, nil)

	body, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPut, "/todos/1", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req, -1)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	mockUsecase.AssertCalled(t, "Update", mock.Anything, mock.Anything, uint(1))
}

func TestDeleteSuccess(t *testing.T) {
	mockUsecase := new(mocks.Usecase)
	app := setupApp(mockUsecase)

	mockResponse := dto.DeleteTodoResponse{ID: 1}

	mockUsecase.On("Delete", mock.Anything, mock.Anything, uint(1)).Return(mockResponse, http.StatusOK, nil)

	req := httptest.NewRequest(http.MethodDelete, "/todos/1", nil)

	resp, _ := app.Test(req, -1)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	mockUsecase.AssertCalled(t, "Delete", mock.Anything, mock.Anything, uint(1))
}
