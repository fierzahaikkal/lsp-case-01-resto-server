package repository

import (
	"github.com/fierzahaikkal/lsp-case-01-resto-server/internal/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AdminRepository interface {
    CreateAdmin(admin entity.Admin) error
    GetAdminByID(id uuid.UUID) (*entity.Admin, error)
    UpdateAdmin(admin entity.Admin) error
    DeleteAdmin(id uuid.UUID) error
}

type adminRepo struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return &adminRepo{db}
}

func (r *adminRepo) CreateAdmin(admin entity.Admin) error {
	return r.db.Create(&admin).Error
}

func (r *adminRepo) GetAdminByID(id uuid.UUID) (*entity.Admin, error) {
	var admin entity.Admin
	err := r.db.First(&admin, "id = ?", id).Error
	return &admin, err
}

func (r *adminRepo) UpdateAdmin(admin entity.Admin) error {
	return r.db.Save(&admin).Error
}

func (r *adminRepo) DeleteAdmin(id uuid.UUID) error {
	return r.db.Delete(&entity.Admin{}, "id = ?", id).Error
}
