package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/mert-yigittop/cxp-api-starter/pkg/jwt"
	"github.com/mert-yigittop/cxp-api-starter/pkg/utils/response"
	"net/http"
)

func AuthRequired() fiber.Handler {
	return func(c *fiber.Ctx) error {
		jwtCookie := c.Cookies("jwt")
		if jwtCookie == "" {
			return response.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
		}

		tokenString := jwtCookie

		userId, err := jwt.Verify(tokenString)
		if err != nil {
			return response.ErrorResponse(c, http.StatusUnauthorized, fmt.Sprintf("Unauthorized: %v", err.Error()))
		}

		c.Locals("userId", userId)

		return c.Next()
	}
}
