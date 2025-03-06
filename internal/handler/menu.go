package handler

import (
	"github.com/fierzahaikkal/lsp-case-01-resto-server/internal/entity"
	"github.com/fierzahaikkal/lsp-case-01-resto-server/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type MenuHandler struct {
	service usecase.MenuService
}

func NewMenuHandler(service usecase.MenuService) *MenuHandler {
	return &MenuHandler{service}
}

func (h *MenuHandler) CreateMenu(c *fiber.Ctx) error {
	var menu entity.Menu
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
    pesanans, err := h.service.GetMenu()
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(pesanans)
}

func (h *MenuHandler) GetMenuByID(c *fiber.Ctx) error {
    id := c.Params("id")
    pesanan, err := h.service.GetMenuByID(id)
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Pesanan not found"})
    }
    return c.JSON(pesanan)
}

func (h *MenuHandler) UpdateMenu(c *fiber.Ctx) error {
	var menu entity.Menu
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
	