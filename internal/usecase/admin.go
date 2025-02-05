package usecase

import (
	"github.com/go-playground/validator/v10"
)

type AdminService interface {
	CreateAdmin(admin Admin) error
	GetAdminByID(id string) (*Admin, error)
	UpdateAdmin(admin Admin) error
	DeleteAdmin(id string) error
}

type adminService struct {
	repo      AdminRepository
	validator *validator.Validate
}

func NewAdminService(repo AdminRepository) AdminService {
	return &adminService{repo: repo, validator: validator.New()}
}

func (s *adminService) CreateAdmin(admin Admin) error {
	err := s.validator.Struct(admin)
	if err != nil {
		return err
	}
	return s.repo.CreateAdmin(admin)
}

func (s *adminService) GetAdminByID(id string) (*Admin, error) {
	return s.repo.GetAdminByID(id)
}

func (s *adminService) UpdateAdmin(admin Admin) error {
	return s.repo.UpdateAdmin(admin)
}

func (s *adminService) DeleteAdmin(id string) error {
	return s.repo.DeleteAdmin(id)
}
