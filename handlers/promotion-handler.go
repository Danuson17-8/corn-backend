package handlers

import (
	"github.com/Danuson17-8/corn-backend/services"

	"github.com/gofiber/fiber/v2"
)

type PromotionHandler struct {
	Service *services.PromotionService
}

func (h *PromotionHandler) GetActive(c *fiber.Ctx) error {
	promotions, err := h.Service.GetActivePromotions()
	if err != nil {
		return Error(c, 500, "DB Error")
	}

	return c.JSON(promotions)
}
