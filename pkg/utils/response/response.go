package response

import (
	"github.com/gofiber/fiber/v2"
)

func newError(message string) *fiber.Map {
	return &fiber.Map{
		"status":  false,
		"message": message,
	}
}

func newSuccess(data any) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
	}
}

func ErrorResponse(c *fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(newError(message))
}

func SuccessResponse(c *fiber.Ctx, status int, data any) error {
	return c.Status(status).JSON(newSuccess(data))
}
