package repository

import (
	"github.com/fierzahaikkal/lsp-case-01-resto-server/internal/entity"
	"gorm.io/gorm"
)

type PesananRepository interface {
	CreatePesanan(pesanan entity.Pesanan) error
	GetPesanan() ([]entity.Pesanan, error)
	GetPesananByID(id string) (*entity.Pesanan, error)
	CetakPesanan() ([]entity.Pesanan, error)
	CetakPesananByID(id string) (*entity.Pesanan, error)
	UpdatePesanan(pesanan entity.Pesanan) error
	DeletePesanan(id string) error
}

type pesananRepo struct {
	db *gorm.DB
}

func NewPesananRepository(db *gorm.DB) PesananRepository {
	return &pesananRepo{db}
}

func (r *pesananRepo) CreatePesanan(pesanan entity.Pesanan) error {
	return r.db.Create(&pesanan).Error
}

//TODO: GetPesanan
func (r *pesananRepo) GetPesanan() ([]entity.Pesanan, error){
	var pesanan []entity.Pesanan
	err := r.db.Find(&pesanan).Error
	return pesanan, err
}

func (r *pesananRepo) GetPesananByID(id string) (*entity.Pesanan, error) {
	var pesanan entity.Pesanan
	err := r.db.First(&pesanan, "id = ?", id).Error
	return &pesanan, err
}

//TODO: CetakPesanan
func (r *pesananRepo) CetakPesanan() ([]entity.Pesanan, error) {
	var pesanan []entity.Pesanan
	err := r.db.Where("StatausOrder = ?", "selesai").Find(&pesanan).Error
	return pesanan, err
}

func (r *pesananRepo) CetakPesananByID(id string) (*entity.Pesanan, error){
	var pesanan entity.Pesanan
	err := r.db.Where("StatausOrder = ?", "selesai").First("id = ?", id).Find(&pesanan).Error
	return &pesanan, err
}

func (r *pesananRepo) UpdatePesanan(pesanan entity.Pesanan) error {
	return r.db.Save(&pesanan).Error
}

func (r *pesananRepo) DeletePesanan(id string) error {
	return r.db.Delete(&entity.Pesanan{}, "id = ?", id).Error
}

