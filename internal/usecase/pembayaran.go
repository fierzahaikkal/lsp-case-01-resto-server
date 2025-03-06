package usecase

import (
	"time"

	"github.com/fierzahaikkal/lsp-case-01-resto-server/internal/entity"
	"github.com/fierzahaikkal/lsp-case-01-resto-server/internal/model"
	"github.com/fierzahaikkal/lsp-case-01-resto-server/internal/repository"
	"github.com/google/uuid"
)

type PembayaranService interface {
	CreatePembayaran(req model.RequestCreatePayment) error
	GetPembayaran() ([]entity.Pembayaran, error)
	GetPembayaranByID(id string) (*entity.Pembayaran, error)
	UpdatePembayaran(pembayaran entity.Pembayaran) error
	DeletePembayaran(id string) error
}

type pembayaranService struct {
	repo repository.PembayaranRepository
}

func NewPembayaranService(repo repository.PembayaranRepository) PembayaranService {
	return &pembayaranService{repo}
}

func (s *pembayaranService) CreatePembayaran(req model.RequestCreatePayment) error {
	orderID, err := uuid.Parse(req.OrderID)
	if err != nil {
		return err
	}

	now := time.Now()
	pembayaran := entity.Pembayaran{
		ID:            uuid.New(),
		OrderID:       orderID,
		Amount:        req.Amount,
		PaymentMethod: req.PaymentMethod,
		Status:        "pending",
		PaymentDate:   &now,
		CreatedAt:     now,
		UpdatedAt:     now,
	}

	return s.repo.CreatePembayaran(pembayaran)
}

func (s *pembayaranService) GetPembayaran() ([]entity.Pembayaran, error) {
	return s.repo.GetPembayaran()
}

func (s *pembayaranService) GetPembayaranByID(id string) (*entity.Pembayaran, error) {
	return s.repo.GetPembayaranByID(id)
}

func (s *pembayaranService) UpdatePembayaran(pembayaran entity.Pembayaran) error {
	return s.repo.UpdatePembayaran(pembayaran)
}

func (s *pembayaranService) DeletePembayaran(id string) error {
	return s.repo.DeletePembayaran(id)
}
