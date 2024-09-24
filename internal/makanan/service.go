package makanan

import "github.com/go-playground/validator/v10"

type MakananService interface {
	CreateMakanan(makanan Makanan) error
	GetMakananByID(id string) (*Makanan, error)
	UpdateMakanan(makanan Makanan) error
	DeleteMakanan(id string) error
}

type makananService struct {
	repo      MakananRepository
	validator *validator.Validate
}

func NewMakananService(repo MakananRepository) MakananService {
	return &makananService{repo: repo, validator: validator.New()}
}

func (s *makananService) CreateMakanan(makanan Makanan) error {
	err := s.validator.Struct(makanan)
	if err != nil {
		return err
	}
	return s.repo.CreateMakanan(makanan)
}

func (s *makananService) GetMakananByID(id string) (*Makanan, error) {
	return s.repo.GetMakananByID(id)
}

func (s *makananService) UpdateMakanan(makanan Makanan) error {
	return s.repo.UpdateMakanan(makanan)
}

func (s *makananService) DeleteMakanan(id string) error {
	return s.repo.DeleteMakanan(id)
}
