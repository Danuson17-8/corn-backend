package handlers

import "github.com/gofiber/fiber/v2"

func Error(c *fiber.Ctx, code int, msg string) error {
	return c.Status(code).JSON(fiber.Map{
		"success": false,
		"message": msg,
	})
}

func Success(c *fiber.Ctx, msg string, data ...fiber.Map) error {
	resp := fiber.Map{
		"success": true,
		"message": msg,
	}

	if len(data) > 0 {
		resp["data"] = data[0]
	}

	return c.JSON(resp)
}
