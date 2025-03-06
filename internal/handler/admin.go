package handler

import (
	"github.com/fierzahaikkal/lsp-case-01-resto-server/internal/model"
	"github.com/fierzahaikkal/lsp-case-01-resto-server/internal/usecase"
	"github.com/fierzahaikkal/lsp-case-01-resto-server/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type AdminHandler struct {
	uc usecase.AdminUsecase
}

func NewAdminHandler(uc usecase.AdminUsecase) *AdminHandler {
	return &AdminHandler{uc}
}

func (h *AdminHandler) CreateAdmin(c *fiber.Ctx) error {
	var req model.RequestSignUpAdmin
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	// req.ID = uuid.New() -> uuid will automatically generate from postgresql
	admin, err := h.uc.Create(&req)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(admin)
}

func (h *AdminHandler) GetAdmin(c *fiber.Ctx) error {
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

	admin, err := h.uc.GetByID(c.Context(), parsedUUID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Admin not found"})
	}
	return c.JSON(admin)
}

func (h *AdminHandler) UpdateAdmin(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid UUID format",
		})
	}
	var req model.RequestUpdateAdmin
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := h.uc.UpdatePartial(c.Context(), id, &req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(req)
}

func (h *AdminHandler) DeleteAdmin(c *fiber.Ctx) error {
	id := c.Params("id")
	parsedUUID, err := utils.ParseUUID(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	admin, err := h.uc.Delete(c.Context(), parsedUUID)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(admin)
}
