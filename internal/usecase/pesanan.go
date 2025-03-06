package usecase

import (
	"github.com/fierzahaikkal/lsp-case-01-resto-server/internal/entity"
	"github.com/fierzahaikkal/lsp-case-01-resto-server/internal/repository"
	"github.com/go-playground/validator/v10"
)

type PesananService interface {
	CreatePesanan(pesanan entity.Pesanan) error
	GetPesanan() ([]entity.Pesanan, error)
	GetPesananByID(id string) (*entity.Pesanan, error)
	UpdatePesanan(pesanan entity.Pesanan) error
	CetakPesanan() ([]entity.Pesanan, error)
	CetakPesananByID(id string) (*entity.Pesanan, error)
	DeletePesanan(id string) error
}

type pesananService struct {
	repo      repository.PesananRepository
	validator *validator.Validate
}

func NewPesananService(repo repository.PesananRepository) PesananService {
	return &pesananService{repo: repo, validator: validator.New()}
}

func (s *pesananService) CreatePesanan(pesanan entity.Pesanan) error {
	err := s.validator.Struct(pesanan)
	if err != nil {
		return err
	}
	return s.repo.CreatePesanan(pesanan)
}

//TODO: GetPesanan
func (s *pesananService) GetPesanan() ([]entity.Pesanan, error){
	return s.repo.GetPesanan()
}


func (s *pesananService) GetPesananByID(id string) (*entity.Pesanan, error) {
	return s.repo.GetPesananByID(id)
}

func (s *pesananService) CetakPesanan() ([]entity.Pesanan, error) {
	return s.repo.CetakPesanan()
}

func (s *pesananService) CetakPesananByID(id string) (*entity.Pesanan, error){
	return s.repo.CetakPesananByID(id)
}

func (s *pesananService) UpdatePesanan(pesanan entity.Pesanan) error {
	return s.repo.UpdatePesanan(pesanan)
}

func (s *pesananService) DeletePesanan(id string) error {
	return s.repo.DeletePesanan(id)
}
