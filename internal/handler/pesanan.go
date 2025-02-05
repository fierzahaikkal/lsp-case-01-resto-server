package pesanan

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type PesananHandler struct {
	service PesananService
}

func NewPesananHandler(service PesananService) *PesananHandler {
	return &PesananHandler{service}
}

func (h *PesananHandler) CreatePesanan(c *fiber.Ctx) error {
	var pesanan Pesanan
	if err := c.BodyParser(&pesanan); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	pesanan.ID = uuid.New()
	if err := h.service.CreatePesanan(pesanan); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(pesanan)
}

func (h *PesananHandler) GetPesanan(c *fiber.Ctx) error {
    pesanans, err := h.service.GetPesanan()
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(pesanans)
}

func (h *PesananHandler) GetPesananByID(c *fiber.Ctx) error {
    id := c.Params("id")
    pesanan, err := h.service.GetPesananByID(id)
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Pesanan not found"})
    }
    return c.JSON(pesanan)
}

func (h *PesananHandler) CetakPesanan(c *fiber.Ctx) error {
    pesanans, err := h.service.CetakPesanan()
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(pesanans)
}

func (h *PesananHandler) CetakPesananByID(c *fiber.Ctx) error {
    id := c.Params("id")
    pesanan, err := h.service.CetakPesananByID(id)
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Pesanan not found or not completed"})
    }
    return c.JSON(pesanan)
}

func (h *PesananHandler) UpdatePesanan(c *fiber.Ctx) error {
	var pesanan Pesanan
	if err := c.BodyParser(&pesanan); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := h.service.UpdatePesanan(pesanan); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(pesanan)
}

func (h *PesananHandler) DeletePesanan(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.service.DeletePesanan(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusNoContent).Send(nil)
}
