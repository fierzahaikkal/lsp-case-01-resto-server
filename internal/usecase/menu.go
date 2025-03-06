package usecase

import (
	"github.com/fierzahaikkal/lsp-case-01-resto-server/internal/entity"
	"github.com/fierzahaikkal/lsp-case-01-resto-server/internal/repository"
	"github.com/go-playground/validator/v10"
)

type MenuService interface {
	CreateMenu(menu entity.Menu) error
	GetMenu() ([]entity.Menu, error)
	GetMenuByID(id string) (*entity.Menu, error)
	UpdateMenu(menu entity.Menu) error
	DeleteMenu(id string) error
}

type menuService struct {
	repo      repository.MenuRepository
	validator *validator.Validate
}

func NewMenuService(repo repository.MenuRepository) MenuService {
	return &menuService{repo: repo, validator: validator.New()}
}

func (s *menuService) CreateMenu(menu entity.Menu) error {
	err := s.validator.Struct(menu)
	if err != nil {
		return err
	}
	return s.repo.CreateMenu(menu)
}

func (s *menuService) GetMenu() ([]entity.Menu, error) {
	return s.repo.GetMenu()
}

func (s *menuService) GetMenuByID(id string) (*entity.Menu, error) {
	return s.repo.GetMenuByID(id)
}

func (s *menuService) UpdateMenu(menu entity.Menu) error {
	return s.repo.UpdateMenu(menu)
}

func (s *menuService) DeleteMenu(id string) error {
	return s.repo.DeleteMenu(id)
}
