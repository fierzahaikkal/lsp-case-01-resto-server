package repository

import "gorm.io/gorm"

type CustomerRepository interface {
	CreateCustomer(customer Customer) error
	GetCustomerByID(id string) (*Customer, error)
	UpdateCustomer(customer Customer) error
	DeleteCustomer(id string) error
}

type customerRepo struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepo{db}
}

func (r *customerRepo) CreateCustomer(customer Customer) error {
	return r.db.Create(&customer).Error
}

func (r *customerRepo) GetCustomerByID(id string) (*Customer, error) {
	var customer Customer
	err := r.db.First(&customer, "id = ?", id).Error
	return &customer, err
}

func (r *customerRepo) UpdateCustomer(customer Customer) error {
	return r.db.Save(&customer).Error
}

func (r *customerRepo) DeleteCustomer(id string) error {
	return r.db.Delete(&Customer{}, "id = ?", id).Error
}
