package v1

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/mert-yigittop/cxp-api-starter/internal/user/dto"
	"github.com/mert-yigittop/cxp-api-starter/internal/user/usecase"
	"github.com/mert-yigittop/cxp-api-starter/pkg/utils/response"
	"net/http"
	"time"
)

type Handler interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	Logout(c *fiber.Ctx) error
}

type handler struct {
	uc usecase.Usecase
}

func New(uc usecase.Usecase) Handler {
	return &handler{uc: uc}
}

func (h *handler) Register(c *fiber.Ctx) error {
	var (
		ctx, cancel = context.WithTimeout(c.Context(), time.Duration(10*time.Second))
		payload     dto.RegisterRequest
	)
	defer cancel()

	if err := c.BodyParser(&payload); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Parse error: %v", err.Error()))
	}

	if err := payload.Validate(); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Validate error: %v", err.Error()))
	}

	userID, status, err := h.uc.Register(ctx, payload)
	if err != nil {
		return response.ErrorResponse(c, status, fmt.Sprintf("Register error: %v", err.Error()))
	}

	return response.SuccessResponse(c, status, userID)
}

func (h *handler) Login(c *fiber.Ctx) error {
	var (
		ctx, cancel = context.WithTimeout(c.Context(), time.Duration(10*time.Second))
		payload     dto.LoginRequest
	)
	defer cancel()

	if err := c.BodyParser(&payload); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Parse error: %v", err.Error()))
	}

	if err := payload.Validate(); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Validation error: %v", err.Error()))
	}

	user, status, err := h.uc.Login(ctx, payload)
	if err != nil {
		return response.ErrorResponse(c, status, fmt.Sprintf("Login error: %v", err.Error()))
	}

	c.Cookie(&fiber.Cookie{
		Name:    "jwt",
		Value:   user.AccessToken,
		Expires: time.Now().Add(time.Hour * 24),
	})

	return response.SuccessResponse(c, status, user)
}

func (h *handler) Logout(c *fiber.Ctx) error {
	c.Locals("userId", nil)
	c.Cookie(&fiber.Cookie{
		Name:    "jwt",
		Value:   "",
		Expires: time.Now().Add(-time.Hour),
	})

	return response.SuccessResponse(c, http.StatusOK, "Successfully logged out")
}
