package handlers

import (
	"github.com/Danuson17-8/corn-backend/models"
	"github.com/Danuson17-8/corn-backend/services"

	"github.com/gofiber/fiber/v2"
)

type ContactHandler struct {
	Service *services.ContactService
}

func (h *ContactHandler) Create(c *fiber.Ctx) error {
	var body models.ContactUser

	if err := c.BodyParser(&body); err != nil {
		return Error(c, 400, "Invalid JSON")
	}

	err := h.Service.Create(&body)
	if err != nil {
		switch err {
		case services.ErrContactInvalid:
			return Error(c, 400, "All fields are required")
		default:
			return Error(c, 500, "Internal server error")
		}
	}

	return Success(c, "Contact created successfully")
}
