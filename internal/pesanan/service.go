package pesanan

import "github.com/go-playground/validator/v10"

type PesananService interface {
	CreatePesanan(pesanan Pesanan) error
	GetPesananByID(id string) (*Pesanan, error)
	UpdatePesanan(pesanan Pesanan) error
	DeletePesanan(id string) error
}

type pesananService struct {
	repo      PesananRepository
	validator *validator.Validate
}

func NewPesananService(repo PesananRepository) PesananService {
	return &pesananService{repo: repo, validator: validator.New()}
}

func (s *pesananService) CreatePesanan(pesanan Pesanan) error {
	err := s.validator.Struct(pesanan)
	if err != nil {
		return err
	}
	return s.repo.CreatePesanan(pesanan)
}

//TODO: GetPesanan

func (s *pesananService) GetPesananByID(id string) (*Pesanan, error) {
	return s.repo.GetPesananByID(id)
}

func (s *pesananService) UpdatePesanan(pesanan Pesanan) error {
	return s.repo.UpdatePesanan(pesanan)
}

func (s *pesananService) DeletePesanan(id string) error {
	return s.repo.DeletePesanan(id)
}
