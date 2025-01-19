package v1

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/mert-yigittop/cxp-api-starter/internal/todo/dto"
	"github.com/mert-yigittop/cxp-api-starter/internal/todo/entity"
	"github.com/mert-yigittop/cxp-api-starter/internal/todo/usecase"
	"github.com/mert-yigittop/cxp-api-starter/pkg/utils/response"
	"net/http"
	"strconv"
	"time"
)

type Handler interface {
	GetList(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type handler struct {
	uc usecase.Usecase
}

func New(uc usecase.Usecase) Handler {
	return &handler{uc: uc}
}

func (h *handler) GetList(c *fiber.Ctx) error {
	var (
		ctx, cancel = context.WithTimeout(c.Context(), time.Duration(10*time.Second))
	)
	defer cancel()

	userId := c.Locals("userId").(uint)

	req := dto.GetTodoListRequest{
		UserID: userId,
	}

	todos, status, err := h.uc.GetList(ctx, req)
	if err != nil {
		return response.ErrorResponse(c, status, fmt.Sprintf("Get Todo List error: %v", err.Error()))
	}

	return response.SuccessResponse(c, status, todos)
}

func (h *handler) Create(c *fiber.Ctx) error {
	var (
		ctx, cancel = context.WithTimeout(c.Context(), time.Duration(10*time.Second))
		payload     dto.CreateTodoRequest
	)
	defer cancel()

	if err := c.BodyParser(&payload); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Parse error: %v", err.Error()))
	}

	if err := payload.Validate(); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Validate error: %v", err.Error()))
	}

	userId := c.Locals("userId").(uint)

	todo := entity.Todo{
		UserID:  userId,
		Content: payload.Content,
	}

	user, status, err := h.uc.Create(ctx, todo)
	if err != nil {
		return response.ErrorResponse(c, status, fmt.Sprintf("Create Todo error: %v", err.Error()))
	}

	return response.SuccessResponse(c, status, user)
}

func (h *handler) Update(c *fiber.Ctx) error {
	var (
		ctx, cancel = context.WithTimeout(c.Context(), time.Duration(10*time.Second))
		payload     dto.UpdateTodoRequest
	)
	defer cancel()

	if err := c.BodyParser(&payload); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Parse error: %v", err.Error()))
	}

	if err := payload.Validate(); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Validate error: %v", err.Error()))
	}

	userId := c.Locals("userId").(uint)

	todoId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Params error: %v", err.Error()))
	}

	todo := dto.UpdateTodoRequest{
		ID:        uint(todoId),
		Content:   payload.Content,
		Completed: payload.Completed,
		UserID:    payload.UserID,
	}

	user, status, err := h.uc.Update(ctx, todo, userId)
	if err != nil {
		return response.ErrorResponse(c, status, fmt.Sprintf("Update Todo error: %v", err.Error()))
	}

	return response.SuccessResponse(c, status, user)
}

func (h *handler) Delete(c *fiber.Ctx) error {
	var (
		ctx, cancel = context.WithTimeout(c.Context(), time.Duration(10*time.Second))
	)
	defer cancel()

	userId := c.Locals("userId").(uint)

	todoId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Params error: %v", err.Error()))
	}

	todo := dto.DeleteTodoRequest{
		ID: uint(todoId),
	}

	user, status, err := h.uc.Delete(ctx, todo, userId)
	if err != nil {
		return response.ErrorResponse(c, status, fmt.Sprintf("Update Todo error: %v", err.Error()))
	}

	return response.SuccessResponse(c, status, user)
}
