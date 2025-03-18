package handler

import (
	"github.com/fierzahaikkal/lsp-case-01-resto-server/internal/model"
	"github.com/fierzahaikkal/lsp-case-01-resto-server/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type RolesHandler struct{
	rolesUsecase usecase.RolesUsecase
}

func NewRolesHandler(rolesUsecase usecase.RolesUsecase) *RolesHandler{
	return &RolesHandler{rolesUsecase}
}

func (h *RolesHandler) CreateRoles(c *fiber.Ctx) error{
	var req model.RequestAddRoles
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request"})
	}

	role, err := h.rolesUsecase.CreateRoles(&req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to create role"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Role created successfully", "data": role})
}

func (h *RolesHandler) GetAllRoles(c *fiber.Ctx) error{
	roles, err := h.rolesUsecase.GetAllRoles()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to get roles"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Roles fetched successfully", "data": roles})
}

func (h *RolesHandler) GetRolesByID(c *fiber.Ctx) error{
	id := c.Params("id")

	roleID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid role ID"})
	}
	
	role, err := h.rolesUsecase.GetRolesByID(roleID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to get role"})
	}
	
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Role fetched successfully", "data": role})
}

func (h *RolesHandler) UpdateRoles(c *fiber.Ctx) error{
	id := c.Params("id")

	roleID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid role ID"})
	}

	var req model.RequestUpdateRoles
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request"})
	}
	
	err = h.rolesUsecase.UpdateRoles(roleID, &req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to update role"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Role updated successfully"})
}

func (h *RolesHandler) DeleteRoles(c *fiber.Ctx) error{
	id := c.Params("id")

	roleID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid role ID"})
	}

	err = h.rolesUsecase.DeleteRoles(roleID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to delete role"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Role deleted successfully"})
}

