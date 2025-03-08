package handler

import (
	"github.com/fierzahaikkal/lsp-case-01-resto-server/internal/model"
	"github.com/fierzahaikkal/lsp-case-01-resto-server/internal/usecase"
	"github.com/fierzahaikkal/lsp-case-01-resto-server/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type MenuHandler struct {
	uc usecase.MenuUsecase
}

func NewMenuHandler(uc usecase.MenuUsecase) *MenuHandler {
	return &MenuHandler{uc}
}

func (h *MenuHandler) CreateMenu(c *fiber.Ctx) error {
	var req model.RequestCreateMenu
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	menu, err := h.uc.CreateMenu(&req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(menu)
}

func (h *MenuHandler) GetMenu(c *fiber.Ctx) error {
	menus, err := h.uc.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(menus)
}

func (h *MenuHandler) GetMenuByID(c *fiber.Ctx) error {
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

	menu, err := h.uc.GetByID(c.Context(), parsedUUID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Menu not found"})
	}
	return c.JSON(menu)
}

func (h *MenuHandler) UpdateMenu(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid UUID format",
		})
	}

	var req model.RequestUpdateMenu
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.uc.Update(c.Context(), id, &req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Menu updated successfully"})
}

func (h *MenuHandler) DeleteMenu(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid UUID format",
		})
	}

	if err := h.uc.Delete(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusNoContent).Send(nil)
}