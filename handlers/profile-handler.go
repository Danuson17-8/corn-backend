package handlers

import (
	"github.com/Danuson17-8/corn-backend/services"

	"github.com/gofiber/fiber/v2"
)

type ProfileHandler struct {
	Profile *services.ProfileService
}

func (h *ProfileHandler) GetProfile(c *fiber.Ctx) error {
	email := c.Locals("email").(string)

	user, err := h.Profile.GetProfileByEmail(email)
	if err != nil {
		return Error(c, 404, "")
	}

	return c.JSON(user)
}
