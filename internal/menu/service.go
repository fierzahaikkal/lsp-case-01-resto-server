package menu

import "github.com/go-playground/validator/v10"

type MenuService interface {
	CreateMenu(menu Menu) error
	GetMenuByID(id string) (*Menu, error)
	UpdateMenu(menu Menu) error
	DeleteMenu(id string) error
}

type menuService struct {
	repo      MenuRepository
	validator *validator.Validate
}

func NewMenuService(repo MenuRepository) MenuService {
	return &menuService{repo: repo, validator: validator.New()}
}

func (s *menuService) CreateMenu(menu Menu) error {
	err := s.validator.Struct(menu)
	if err != nil {
		return err
	}
	return s.repo.CreateMenu(menu)
}

func (s *menuService) GetMenuByID(id string) (*Menu, error) {
	return s.repo.GetMenuByID(id)
}

func (s *menuService) UpdateMenu(menu Menu) error {
	return s.repo.UpdateMenu(menu)
}

func (s *menuService) DeleteMenu(id string) error {
	return s.repo.DeleteMenu(id)
}
