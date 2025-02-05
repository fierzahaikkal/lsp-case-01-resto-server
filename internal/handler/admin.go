package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type AdminHandler struct {
	service AdminService
}

func NewAdminHandler(service AdminService) *AdminHandler {
	return &AdminHandler{service}
}

func (h *AdminHandler) CreateAdmin(c *fiber.Ctx) error {
	var admin Admin
	if err := c.BodyParser(&admin); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	admin.ID = uuid.New()
	if err := h.service.CreateAdmin(admin); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(admin)
}

func (h *AdminHandler) GetAdmin(c *fiber.Ctx) error {
	id := c.Params("id")
	admin, err := h.service.GetAdminByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Admin not found"})
	}
	return c.JSON(admin)
}

func (h *AdminHandler) UpdateAdmin(c *fiber.Ctx) error {
	var admin Admin
	if err := c.BodyParser(&admin); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := h.service.UpdateAdmin(admin); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(admin)
}

func (h *AdminHandler) DeleteAdmin(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.service.DeleteAdmin(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusNoContent).Send(nil)
}
