package handler

import (
	"github.com/fierzahaikkal/lsp-case-01-resto-server/internal/model"
	"github.com/fierzahaikkal/lsp-case-01-resto-server/internal/usecase"
	"github.com/fierzahaikkal/lsp-case-01-resto-server/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type CustomerHandler struct {
	uc usecase.CustomerUsecase
}

func NewCustomerHandler(uc usecase.CustomerUsecase) *CustomerHandler {
	return &CustomerHandler{uc}
}

func (h *CustomerHandler) CreateCustomer(c *fiber.Ctx) error {
	var req model.RequestSignUpCustomer
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	customer, err := h.uc.Create(&req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(customer)
}

func (h *CustomerHandler) GetCustomer(c *fiber.Ctx) error {
	id := c.Params("id")

	parsedUUID, err := utils.ParseUUID(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if utils.IsEmptyUUID(parsedUUID) {
		return c.Status(400).JSON(fiber.Map{
			"message": "UUID cannot be empty",
		})
	}

	customer, err := h.uc.GetByID(c.Context(), parsedUUID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Customer not found"})
	}
	return c.JSON(customer)
}

func (h *CustomerHandler) UpdateCustomer(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid UUID format",
		})
	}

	var req model.RequestSignUpCustomer
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.uc.UpdatePartial(c.Context(), id, &req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Customer updated successfully"})
}

func (h *CustomerHandler) DeleteCustomer(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid UUID format",
		})
	}

	deletedCustomer, err := h.uc.Delete(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(deletedCustomer)
}
