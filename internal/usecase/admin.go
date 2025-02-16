package usecase

import (
	"errors"

	"github.com/fierzahaikkal/lsp-case-01-resto-server/internal/entity"
	"github.com/fierzahaikkal/lsp-case-01-resto-server/internal/model"
	"github.com/fierzahaikkal/lsp-case-01-resto-server/internal/repository"
	"github.com/fierzahaikkal/lsp-case-01-resto-server/pkg/utils"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// AdminUsecase implements AdminUsecase interface
type AdminUsecase struct {
    adminRepo repository.AdminRepository
    validator *validator.Validate
}

// NewAdminUsecase creates a new instance of AdminUsecase
func NewAdminUsecase(repo repository.AdminRepository) *AdminUsecase {
    return &AdminUsecase{
        adminRepo: repo,
        validator: validator.New(),
    }
}

// Create handles the business logic for creating a new admin
func (u *AdminUsecase) Create(req *model.RequestSignUpAdmin) (*entity.Admin, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("Password cannot be hash: " + err.Error())
	}

    if req.Password != req.ValidatePassword {
        return nil, errors.New("Password not match")
    }

    admin := entity.Admin {
        ID: utils.GenUUID(),
        Nama: req.Nama,
        Email: req.Email,
        Username: req.Password,
        Password: string(hashedPassword),
    }

    if req == nil {
        return nil, errors.New("please fill the request")
    }

    // Validate admin data
    if err := u.validator.Struct(req); err != nil {
        return nil, errors.New("invalid admin data: " + err.Error())
    }

    // Create admin in repository
    if err := u.adminRepo.CreateAdmin(admin); err != nil {
        return nil, errors.New("failed to create admin: " + err.Error())
    }

    return &admin, err
}

// GetByID retrieves an admin by their ID
func (u *AdminUsecase) GetByID(id uuid.UUID) (*entity.Admin, error) {
    if id == uuid.Nil {
        return nil, errors.New("ID not found")
    }

    admin, err := u.adminRepo.GetAdminByID(id)
    if err != nil {
        return nil, errors.New("Failed to get admin: " + err.Error())
    }

    if admin == nil {
        return nil, errors.New("Admin not found")
    }

    return admin, nil
}

// Update handles the business logic for updating an existing admin
func (u *AdminUsecase) Update(admin *entity.Admin) error {
    if admin == nil {
        return errors.New("admin cannot be nil")
    }

    if admin.ID == uuid.Nil {
        return errors.New("invalid UUID")
    }

    // Validate admin data
    if err := u.validator.Struct(admin); err != nil {
        return errors.New("invalid admin data: " + err.Error())
    }

    // Check if admin exists
    existingAdmin, err := u.adminRepo.GetAdminByID(admin.ID)
    if err != nil {
        return errors.New("failed to check existing admin: " + err.Error())
    }
    if existingAdmin == nil {
        return errors.New("admin not found")
    }

    // Update admin in repository
    if err := u.adminRepo.UpdateAdmin(*admin); err != nil {
        return errors.New("failed to update admin: " + err.Error())
    }

    return nil
}

// Delete handles the business logic for deleting an admin
func (u *AdminUsecase) Delete(id uuid.UUID) (*entity.Admin, error) {
    if id == uuid.Nil {
        return nil, errors.New("invalid UUID")
    }

    // Check if admin exists
    existingAdmin, err := u.adminRepo.GetAdminByID(id)
    if err != nil {
        return nil, errors.New("failed to check existing admin: " + err.Error())
    }
    if existingAdmin == nil {
        return nil, errors.New("admin not found")
    }

    // Delete admin from repository
    if err := u.adminRepo.DeleteAdmin(id); err != nil {
        return nil, errors.New("failed to delete admin: " + err.Error())
    }

    return existingAdmin, nil
}