package usecase

import (
	"github.com/fierzahaikkal/lsp-case-01-resto-server/internal/entity"
	"github.com/fierzahaikkal/lsp-case-01-resto-server/internal/repository"
	"github.com/go-playground/validator/v10"
)

type CustomerService interface {
	CreateCustomer(customer entity.Customer) error
	GetCustomerByID(id string) (*entity.Customer, error)
	UpdateCustomer(customer entity.Customer) error
	DeleteCustomer(id string) error
}

type customerService struct {
	repo      repository.CustomerRepository
	validator *validator.Validate
}

func NewCustomerService(repo repository.CustomerRepository) CustomerService {
	return &customerService{repo: repo, validator: validator.New()}
}

func (s *customerService) CreateCustomer(customer entity.Customer) error {
	err := s.validator.Struct(customer)
	if err != nil {
		return err
	}
	return s.repo.CreateCustomer(customer)
}

func (s *customerService) GetCustomerByID(id string) (*entity.Customer, error) {
	return s.repo.GetCustomerByID(id)
}

func (s *customerService) UpdateCustomer(customer entity.Customer) error {
	return s.repo.UpdateCustomer(customer)
}

func (s *customerService) DeleteCustomer(id string) error {
	return s.repo.DeleteCustomer(id)
}
