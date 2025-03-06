package repository

import (
	"context"

	"github.com/fierzahaikkal/lsp-case-01-resto-server/internal/entity"
	"github.com/fierzahaikkal/lsp-case-01-resto-server/internal/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AdminRepository interface {
	CreateAdmin(admin entity.Admin) error
	GetAdminByID(ctx context.Context, id uuid.UUID) (*entity.Admin, error)
	UpdateAdmin(ctx context.Context, id uuid.UUID, updates map[string]interface{}) error
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

func (r *adminRepo) GetAdminByID(ctx context.Context, id uuid.UUID) (*entity.Admin, error) {
	var admin entity.Admin
	err := r.db.First(&admin, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &admin, err
}

func (r *adminRepo) UpdateAdmin(ctx context.Context, id uuid.UUID, updates map[string]interface{}) error {
	return r.db.Model(&model.RequestUpdateAdmin{}).Where("id = ?", id).Updates(updates).Error
}

func (r *adminRepo) DeleteAdmin(id uuid.UUID) error {
	return r.db.Delete(&entity.Admin{}, "id = ?", id).Error
}
