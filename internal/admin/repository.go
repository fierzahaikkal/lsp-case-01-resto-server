package admin

import (
	"gorm.io/gorm"
)

type AdminRepository interface {
	CreateAdmin(admin Admin) error
	GetAdminByID(id string) (*Admin, error)
	UpdateAdmin(admin Admin) error
	DeleteAdmin(id string) error
}

type adminRepo struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return &adminRepo{db}
}

func (r *adminRepo) CreateAdmin(admin Admin) error {
	return r.db.Create(&admin).Error
}

func (r *adminRepo) GetAdminByID(id string) (*Admin, error) {
	var admin Admin
	err := r.db.First(&admin, "id = ?", id).Error
	return &admin, err
}

func (r *adminRepo) UpdateAdmin(admin Admin) error {
	return r.db.Save(&admin).Error
}

func (r *adminRepo) DeleteAdmin(id string) error {
	return r.db.Delete(&Admin{}, "id = ?", id).Error
}
