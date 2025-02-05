package repository

import "gorm.io/gorm"

type PesananRepository interface {
	CreatePesanan(pesanan Pesanan) error
	GetPesanan() ([]Pesanan, error)
	GetPesananByID(id string) (*Pesanan, error)
	CetakPesanan() ([]Pesanan, error)
	CetakPesananByID(id string) (*Pesanan, error)
	UpdatePesanan(pesanan Pesanan) error
	DeletePesanan(id string) error
}

type pesananRepo struct {
	db *gorm.DB
}

func NewPesananRepository(db *gorm.DB) PesananRepository {
	return &pesananRepo{db}
}

func (r *pesananRepo) CreatePesanan(pesanan Pesanan) error {
	return r.db.Create(&pesanan).Error
}

//TODO: GetPesanan
func (r *pesananRepo) GetPesanan() ([]Pesanan, error){
	var pesanan []Pesanan
	err := r.db.Find(&pesanan).Error
	return pesanan, err
}

func (r *pesananRepo) GetPesananByID(id string) (*Pesanan, error) {
	var pesanan Pesanan
	err := r.db.First(&pesanan, "id = ?", id).Error
	return &pesanan, err
}

//TODO: CetakPesanan
func (r *pesananRepo) CetakPesanan() ([]Pesanan, error) {
	var pesanan []Pesanan
	err := r.db.Where("StatausOrder = ?", "selesai").Find(&pesanan).Error
	return pesanan, err
}

func (r *pesananRepo) CetakPesananByID(id string) (*Pesanan, error){
	var pesanan Pesanan
	err := r.db.Where("StatausOrder = ?", "selesai").First("id = ?", id).Find(&pesanan).Error
	return &pesanan, err
}

func (r *pesananRepo) UpdatePesanan(pesanan Pesanan) error {
	return r.db.Save(&pesanan).Error
}

func (r *pesananRepo) DeletePesanan(id string) error {
	return r.db.Delete(&Pesanan{}, "id = ?", id).Error
}

