package handlers

import (
	"github.com/Danuson17-8/corn-backend/services"

	"github.com/gofiber/fiber/v2"
)

type MenuHandler struct {
	Service *services.MenuService
}

func (h *MenuHandler) GetMenu(c *fiber.Ctx) error {
	menus, err := h.Service.GetMenu()
	if err != nil {
		return Error(c, 500, "DB Error")
	}
	return c.JSON(menus)
}
