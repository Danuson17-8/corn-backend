package middlewares

import (
	"github.com/Danuson17-8/corn-backend/services"

	"github.com/gofiber/fiber/v2"
)

func RequireAuth(jwt *services.JWTService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Cookies("access_token")
		if token == "" {
			return fiber.ErrUnauthorized
		}

		claims, err := jwt.Verify(token)
		if err != nil {
			return fiber.ErrUnauthorized
		}

		// inject user info
		c.Locals("email", claims.Email)
		c.Locals("role", claims.Role)

		return c.Next()
	}
}
