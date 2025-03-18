package repository

import (
	"github.com/fierzahaikkal/lsp-case-01-resto-server/internal/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RolesRepository interface{
	CreateRoles(role *entity.Roles) (*entity.Roles, error)
	GetAllRoles() ([]entity.Roles, error)
	GetRolesByID(id uuid.UUID) (*entity.Roles, error)
	UpdateRoles(id uuid.UUID, role *entity.Roles) error
	DeleteRoles(id uuid.UUID) error
}

type rolesRepository struct{
	db *gorm.DB
}

func NewRolesRepository(db *gorm.DB) RolesRepository{
	return &rolesRepository{db}
}

func (r *rolesRepository) CreateRoles(role *entity.Roles) (*entity.Roles, error){
	err := r.db.Create(role).Error
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (r *rolesRepository) GetAllRoles() ([]entity.Roles, error){
	var roles []entity.Roles
	err := r.db.Find(&roles).Error
	return roles, err
}

func (r *rolesRepository) GetRolesByID(id uuid.UUID) (*entity.Roles, error){
	var role entity.Roles
	err := r.db.First(&role, "id = ?", id).Error
	return &role, err
}

func (r *rolesRepository) UpdateRoles(id uuid.UUID, role *entity.Roles) error {
	return r.db.Model(&entity.Roles{}).Where("id = ?", id).Updates(role).Error
}

func (r *rolesRepository) DeleteRoles(id uuid.UUID) error {
	return r.db.Delete(&entity.Roles{}, "id = ?", id).Error
}

