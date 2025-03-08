package usecase

import (
	"context"
	"errors"

	"github.com/fierzahaikkal/lsp-case-01-resto-server/internal/entity"
	"github.com/fierzahaikkal/lsp-case-01-resto-server/internal/model"
	"github.com/fierzahaikkal/lsp-case-01-resto-server/internal/repository"
	"github.com/fierzahaikkal/lsp-case-01-resto-server/pkg/utils"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type CustomerUsecase interface {
	Create(req *model.RequestSignUpCustomer) (*entity.Customer, error)
	GetByID(ctx context.Context, id uuid.UUID) (*entity.Customer, error)
	UpdatePartial(ctx context.Context, id uuid.UUID, req *model.RequestSignUpCustomer) error
	Delete(ctx context.Context, id uuid.UUID) (*entity.Customer, error)
}

type customerUsecase struct {
	customerRepo repository.CustomerRepository
	validator    *validator.Validate
}

func NewCustomerUsecase(repo repository.CustomerRepository) CustomerUsecase {
	return &customerUsecase{
		customerRepo: repo,
		validator:    validator.New(),
	}
}

func (u *customerUsecase) Create(req *model.RequestSignUpCustomer) (*entity.Customer, error) {
	if req == nil {
		return nil, errors.New("please fill the request")
	}

	// Validate customer data
	if err := u.validator.Struct(req); err != nil {
		return nil, errors.New("invalid customer data: " + err.Error())
	}

	customer := entity.Customer{
		ID:      utils.GenUUID(),
		Nama:    req.Nama,
		Email:   req.Email,
		Telepon: req.Telepon,
		Alamat:  req.Alamat,
	}

	// Create customer in repository
	if err := u.customerRepo.CreateCustomer(customer); err != nil {
		return nil, errors.New("failed to create customer: " + err.Error())
	}

	return &customer, nil
}

func (u *customerUsecase) GetByID(ctx context.Context, id uuid.UUID) (*entity.Customer, error) {
	if id == uuid.Nil {
		return nil, errors.New("ID not found")
	}

	customer, err := u.customerRepo.GetCustomerByID(ctx, id)
	if err != nil {
		return nil, errors.New("Failed to get customer: " + err.Error())
	}

	if customer == nil {
		return nil, errors.New("Customer not found")
	}

	return customer, nil
}

func (u *customerUsecase) UpdatePartial(ctx context.Context, id uuid.UUID, req *model.RequestSignUpCustomer) error {
	updates := make(map[string]interface{})

	if req.Nama != "" {
		updates["nama"] = req.Nama
	}
	if req.Email != "" {
		updates["email"] = req.Email
	}
	if req.Telepon != "" {
		updates["telepon"] = req.Telepon
	}
	if req.Alamat != "" {
		updates["alamat"] = req.Alamat
	}

	return u.customerRepo.UpdateCustomer(ctx, id, updates)
}

func (u *customerUsecase) Delete(ctx context.Context, id uuid.UUID) (*entity.Customer, error) {
	if id == uuid.Nil {
		return nil, errors.New("invalid UUID")
	}

	// Check if customer exists
	existingCustomer, err := u.customerRepo.GetCustomerByID(ctx, id)
	if err != nil {
		return nil, errors.New("failed to check existing customer: " + err.Error())
	}
	if existingCustomer == nil {
		return nil, errors.New("customer not found")
	}

	// Delete customer from repository
	if err := u.customerRepo.DeleteCustomer(id); err != nil {
		return nil, errors.New("failed to delete customer: " + err.Error())
	}

	return existingCustomer, nil
}
