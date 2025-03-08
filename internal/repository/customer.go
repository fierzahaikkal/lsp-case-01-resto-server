package repository

import (
	"context"

	"github.com/fierzahaikkal/lsp-case-01-resto-server/internal/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	CreateCustomer(customer entity.Customer) error
	GetCustomerByID(ctx context.Context, id uuid.UUID) (*entity.Customer, error)
	UpdateCustomer(ctx context.Context, id uuid.UUID, updates map[string]interface{}) error
	DeleteCustomer(id uuid.UUID) error
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

func (r *customerRepo) GetCustomerByID(ctx context.Context, id uuid.UUID) (*entity.Customer, error) {
	var customer entity.Customer
	err := r.db.First(&customer, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &customer, err
}

func (r *customerRepo) UpdateCustomer(ctx context.Context, id uuid.UUID, updates map[string]interface{}) error {
	return r.db.Model(&entity.Customer{}).Where("id = ?", id).Updates(updates).Error
}

func (r *customerRepo) DeleteCustomer(id uuid.UUID) error {
	return r.db.Delete(&entity.Customer{}, "id = ?", id).Error
}
