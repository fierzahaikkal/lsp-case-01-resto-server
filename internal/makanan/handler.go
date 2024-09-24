package makanan

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type MakananHandler struct {
	service MakananService
}

func NewMakananHandler(service MakananService) *MakananHandler {
	return &MakananHandler{service}
}

func (h *MakananHandler) CreateMakanan(c *fiber.Ctx) error {
	var makanan Makanan
	if err := c.BodyParser(&makanan); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	makanan.ID = uuid.New()
	if err := h.service.CreateMakanan(makanan); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(makanan)
}

func (h *MakananHandler) GetMakanan(c *fiber.Ctx) error {
	id := c.Params("id")
	makanan, err := h.service.GetMakananByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Makanan not found"})
	}
	return c.JSON(makanan)
}

func (h *MakananHandler) UpdateMakanan(c *fiber.Ctx) error {
	var makanan Makanan
	if err := c.BodyParser(&makanan); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := h.service.UpdateMakanan(makanan); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(makanan)
}

func (h *MakananHandler) DeleteMakanan(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.service.DeleteMakanan(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusNoContent).Send(nil)
}
