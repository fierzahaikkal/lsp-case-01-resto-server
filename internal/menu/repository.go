package menu

import "gorm.io/gorm"

type MenuRepository interface {
	CreateMenu(menu Menu) error
	GetMenuByID(id string) (*Menu, error)
	UpdateMenu(menu Menu) error
	DeleteMenu(id string) error
}

type menuRepo struct {
	db *gorm.DB
}

func NewMenuRepository(db *gorm.DB) MenuRepository {
	return &menuRepo{db}
}

func (r *menuRepo) CreateMenu(menu Menu) error {
	return r.db.Create(&menu).Error
}

func (r *menuRepo) GetMenuByID(id string) (*Menu, error) {
	var menu Menu
	err := r.db.First(&menu, "id = ?", id).Error
	return &menu, err
}

func (r *menuRepo) UpdateMenu(menu Menu) error {
	return r.db.Save(&menu).Error
}

func (r *menuRepo) DeleteMenu(id string) error {
	return r.db.Delete(&Menu{}, "id = ?", id).Error
}
