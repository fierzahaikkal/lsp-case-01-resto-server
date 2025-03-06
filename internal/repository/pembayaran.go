package repository

import (
	"github.com/fierzahaikkal/lsp-case-01-resto-server/internal/entity"
	"gorm.io/gorm"
)

type PembayaranRepository interface {
	CreatePembayaran(pembayaran entity.Pembayaran) error
	GetPembayaran() ([]entity.Pembayaran, error)
	GetPembayaranByID(id string) (*entity.Pembayaran, error)
	UpdatePembayaran(pembayaran entity.Pembayaran) error
	DeletePembayaran(id string) error
}

type pembayaranRepo struct {
	db *gorm.DB
}

func NewPembayaranRepository(db *gorm.DB) PembayaranRepository {
	return &pembayaranRepo{db}
}

func (r *pembayaranRepo) CreatePembayaran(pembayaran entity.Pembayaran) error {
	return r.db.Create(&pembayaran).Error
}

func (r *pembayaranRepo) GetPembayaran() ([]entity.Pembayaran, error) {
	var pembayaran []entity.Pembayaran
	err := r.db.Find(&pembayaran).Error
	return pembayaran, err
}

func (r *pembayaranRepo) GetPembayaranByID(id string) (*entity.Pembayaran, error) {
	var pembayaran entity.Pembayaran
	err := r.db.First(&pembayaran, "id = ?", id).Error
	return &pembayaran, err
}

func (r *pembayaranRepo) UpdatePembayaran(pembayaran entity.Pembayaran) error {
	return r.db.Save(&pembayaran).Error
}

func (r *pembayaranRepo) DeletePembayaran(id string) error {
	return r.db.Delete(&entity.Pembayaran{}, "id = ?", id).Error
}
