package menu

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type MenuHandler struct {
	service MenuService
}

func NewMenuHandler(service MenuService) *MenuHandler {
	return &MenuHandler{service}
}

func (h *MenuHandler) CreateMenu(c *fiber.Ctx) error {
	var menu Menu
	if err := c.BodyParser(&menu); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	menu.ID = uuid.New()
	if err := h.service.CreateMenu(menu); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(menu)
}

func (h *MenuHandler) GetMenu(c *fiber.Ctx) error {
	id := c.Params("id")
	menu, err := h.service.GetMenuByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Menu not found"})
	}
	return c.JSON(menu)
}

func (h *MenuHandler) UpdateMenu(c *fiber.Ctx) error {
	var menu Menu
	if err := c.BodyParser(&menu); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := h.service.UpdateMenu(menu); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(menu)
}

func (h *MenuHandler) DeleteMenu(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.service.DeleteMenu(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusNoContent).Send(nil)
}
