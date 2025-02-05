package usecase

import "github.com/go-playground/validator/v10"

type CustomerService interface {
	CreateCustomer(customer Customer) error
	GetCustomerByID(id string) (*Customer, error)
	UpdateCustomer(customer Customer) error
	DeleteCustomer(id string) error
}

type customerService struct {
	repo      CustomerRepository
	validator *validator.Validate
}

func NewCustomerService(repo CustomerRepository) CustomerService {
	return &customerService{repo: repo, validator: validator.New()}
}

func (s *customerService) CreateCustomer(customer Customer) error {
	err := s.validator.Struct(customer)
	if err != nil {
		return err
	}
	return s.repo.CreateCustomer(customer)
}

func (s *customerService) GetCustomerByID(id string) (*Customer, error) {
	return s.repo.GetCustomerByID(id)
}

func (s *customerService) UpdateCustomer(customer Customer) error {
	return s.repo.UpdateCustomer(customer)
}

func (s *customerService) DeleteCustomer(id string) error {
	return s.repo.DeleteCustomer(id)
}
