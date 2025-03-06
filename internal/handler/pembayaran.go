package handler

import (
	"github.com/fierzahaikkal/lsp-case-01-resto-server/internal/entity"
	"github.com/fierzahaikkal/lsp-case-01-resto-server/internal/model"
	"github.com/fierzahaikkal/lsp-case-01-resto-server/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type PembayaranHandler struct {
	service usecase.PembayaranService
}

func NewPembayaranHandler(service usecase.PembayaranService) *PembayaranHandler {
	return &PembayaranHandler{service}
}

func (h *PembayaranHandler) CreatePembayaran(c *fiber.Ctx) error {
	var req model.RequestCreatePayment
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := h.service.CreatePembayaran(req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Payment created successfully",
	})
}

func (h *PembayaranHandler) GetPembayaran(c *fiber.Ctx) error {
	payments, err := h.service.GetPembayaran()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(payments)
}

func (h *PembayaranHandler) GetPembayaranByID(c *fiber.Ctx) error {
	id := c.Params("id")
	payment, err := h.service.GetPembayaranByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Payment not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(payment)
}

func (h *PembayaranHandler) UpdatePembayaran(c *fiber.Ctx) error {
	id := c.Params("id")
	paymentID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid payment ID",
		})
	}

	var payment entity.Pembayaran
	if err := c.BodyParser(&payment); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	payment.ID = paymentID
	if err := h.service.UpdatePembayaran(payment); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Payment updated successfully",
	})
}

func (h *PembayaranHandler) DeletePembayaran(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.service.DeletePembayaran(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Payment deleted successfully",
	})
}

