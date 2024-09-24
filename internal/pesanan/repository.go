package pesanan

import "gorm.io/gorm"

type PesananRepository interface {
	CreatePesanan(pesanan Pesanan) error
	GetPesananByID(id string) (*Pesanan, error)
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

func (r *pesananRepo) GetPesananByID(id string) (*Pesanan, error) {
	var pesanan Pesanan
	err := r.db.First(&pesanan, "id = ?", id).Error
	return &pesanan, err
}

func (r *pesananRepo) UpdatePesanan(pesanan Pesanan) error {
	return r.db.Save(&pesanan).Error
}

func (r *pesananRepo) DeletePesanan(id string) error {
	return r.db.Delete(&Pesanan{}, "id = ?", id).Error
}
