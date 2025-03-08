package repository

import (
	"github.com/fierzahaikkal/lsp-case-01-resto-server/internal/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MenuRepository interface {
	CreateMenu(menu entity.Menu) error
	GetMenu() ([]entity.Menu, error)
	GetMenuByID(id string) (*entity.Menu, error)
	UpdateMenu(menu entity.Menu) error
	DeleteMenu(id uuid.UUID) error
}

type menuRepo struct {
	db *gorm.DB
}

func NewMenuRepository(db *gorm.DB) MenuRepository {
	return &menuRepo{db}
}

func (r *menuRepo) CreateMenu(menu entity.Menu) error {
	return r.db.Create(&menu).Error
}

func (r *menuRepo) GetMenu() ([]entity.Menu, error) {
	var menu []entity.Menu
	err := r.db.Find(&menu).Error
	return menu, err
} 

func (r *menuRepo) GetMenuByID(id string) (*entity.Menu, error) {
	var menu entity.Menu
	err := r.db.First(&menu, "id = ?", id).Error
	return &menu, err
}

func (r *menuRepo) UpdateMenu(menu entity.Menu) error {
	return r.db.Save(&menu).Error
}

func (r *menuRepo) DeleteMenu(id uuid.UUID) error {
	return r.db.Delete(&entity.Menu{}, "id = ?", id).Error
}
