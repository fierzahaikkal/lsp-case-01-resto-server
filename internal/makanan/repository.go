package makanan

import "gorm.io/gorm"

type MakananRepository interface {
	CreateMakanan(makanan Makanan) error
	GetMakananByID(id string) (*Makanan, error)
	UpdateMakanan(makanan Makanan) error
	DeleteMakanan(id string) error
}

type makananRepo struct {
	db *gorm.DB
}

func NewMakananRepository(db *gorm.DB) MakananRepository {
	return &makananRepo{db}
}

func (r *makananRepo) CreateMakanan(makanan Makanan) error {
	return r.db.Create(&makanan).Error
}

func (r *makananRepo) GetMakananByID(id string) (*Makanan, error) {
	var makanan Makanan
	err := r.db.First(&makanan, "id = ?", id).Error
	return &makanan, err
}

func (r *makananRepo) UpdateMakanan(makanan Makanan) error {
	return r.db.Save(&makanan).Error
}

func (r *makananRepo) DeleteMakanan(id string) error {
	return r.db.Delete(&Makanan{}, "id = ?", id).Error
}
