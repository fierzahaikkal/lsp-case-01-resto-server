package repository

import (
	"github.com/fierzahaikkal/lsp-case-01-resto-server/internal/entity"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	CreateCustomer(customer entity.Customer) error
	GetCustomerByID(id string) (*entity.Customer, error)
	UpdateCustomer(customer entity.Customer) error
	DeleteCustomer(id string) error
}

type customerRepo struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepo{db}
}

func (r *customerRepo) CreateCustomer(customer entity.Customer) error {
	return r.db.Create(&customer).Error
}

func (r *customerRepo) GetCustomerByID(id string) (*entity.Customer, error) {
	var customer entity.Customer
	err := r.db.First(&customer, "id = ?", id).Error
	return &customer, err
}

func (r *customerRepo) UpdateCustomer(customer entity.Customer) error {
	return r.db.Save(&customer).Error
}

func (r *customerRepo) DeleteCustomer(id string) error {
	return r.db.Delete(&entity.Customer{}, "id = ?", id).Error
}
