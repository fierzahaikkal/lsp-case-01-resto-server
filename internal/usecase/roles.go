package usecase

import (
	"github.com/fierzahaikkal/lsp-case-01-resto-server/internal/entity"
	"github.com/fierzahaikkal/lsp-case-01-resto-server/internal/model"
	"github.com/fierzahaikkal/lsp-case-01-resto-server/internal/repository"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type RolesUsecase interface{
	CreateRoles(req *model.RequestAddRoles) (*entity.Roles, error)
	GetAllRoles() ([]entity.Roles, error)
	GetRolesByID(id uuid.UUID) (*entity.Roles, error)
	UpdateRoles(id uuid.UUID, req *model.RequestUpdateRoles) error
	DeleteRoles(id uuid.UUID) error
}

type rolesUsecase struct{
	repo repository.RolesRepository
	validator *validator.Validate
}

func NewRolesUsecase(repo repository.RolesRepository) RolesUsecase{
	return &rolesUsecase{
		repo: repo,
		validator: validator.New(),
	}
}

func (u *rolesUsecase) CreateRoles(req *model.RequestAddRoles) (*entity.Roles, error){
	err := u.validator.Struct(req)
	if err != nil {
		return nil, err
	}

	role := &entity.Roles{
		ID: uuid.New(),
		Name: req.Name,
		Level: req.Level,
	}

	return u.repo.CreateRoles(role)
}

func (u *rolesUsecase) GetAllRoles() ([]entity.Roles, error){
	return u.repo.GetAllRoles()
}

func (u *rolesUsecase) GetRolesByID(id uuid.UUID) (*entity.Roles, error){
	return u.repo.GetRolesByID(id)
}

func (u *rolesUsecase) UpdateRoles(id uuid.UUID, req *model.RequestUpdateRoles) error{
	err := u.validator.Struct(req)
	if err != nil {
		return err
	}

	role := &entity.Roles{
		Name: req.Name,
		Level: req.Level,
	}

	return u.repo.UpdateRoles(id, role)
}

func (u *rolesUsecase) DeleteRoles(id uuid.UUID) error{
	return u.repo.DeleteRoles(id)
}
