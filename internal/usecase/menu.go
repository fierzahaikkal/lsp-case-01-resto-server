package usecase

import (
	"context"

	"github.com/fierzahaikkal/lsp-case-01-resto-server/internal/entity"
	"github.com/fierzahaikkal/lsp-case-01-resto-server/internal/model"
	"github.com/fierzahaikkal/lsp-case-01-resto-server/internal/repository"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type MenuUsecase interface {
	CreateMenu(req *model.RequestCreateMenu) (*entity.Menu, error)
	GetAll() ([]entity.Menu, error) 
	GetByID(ctx context.Context, id uuid.UUID) (*entity.Menu, error)
	Update(ctx context.Context, id uuid.UUID, req *model.RequestUpdateMenu) error
	Delete(id uuid.UUID) error
}

type menuUsecase struct {
	repo      repository.MenuRepository
	validator *validator.Validate
}

func NewMenuUsecase(repo repository.MenuRepository) MenuUsecase {
	return &menuUsecase{
		repo:      repo,
		validator: validator.New(),
	}
}

func (u *menuUsecase) CreateMenu(req *model.RequestCreateMenu) (*entity.Menu, error) {
	menu := &entity.Menu{
		ID:        uuid.New(),
		Nama:      req.Nama,
		Deskripsi: req.Deskripsi,
		Stok:      int(req.Stok),
		Harga:     int(req.Harga),
		Kategori:  req.Kategori,
		URI_image: req.URI_image,
	}

	if err := u.validator.Struct(menu); err != nil {
		return nil, err
	}

	if err := u.repo.CreateMenu(*menu); err != nil {
		return nil, err
	}

	return menu, nil
}

func (u *menuUsecase) GetAll() ([]entity.Menu, error) {
	return u.repo.GetMenu()
}

func (u *menuUsecase) GetByID(ctx context.Context, id uuid.UUID) (*entity.Menu, error) {
	return u.repo.GetMenuByID(id.String())
}

func (u *menuUsecase) Update(ctx context.Context, id uuid.UUID, req *model.RequestUpdateMenu) error {
	updates := make(map[string]interface{})

	if req.Nama != nil {
		updates["nama"] = *req.Nama
	}
	if req.Deskripsi != nil {
		updates["deskripsi"] = *req.Deskripsi
	}
	if req.Stok != nil {
		updates["stok"] = *req.Stok
	}
	if req.Harga != nil {
		updates["harga"] = *req.Harga
	}
	if req.Kategori != nil {
		updates["kategori"] = *req.Kategori
	}
	if req.URI_image != nil {
		updates["uri_image"] = *req.URI_image
	}

	menu := &entity.Menu{ID: id}
	for field, value := range updates {
		switch field {
		case "nama":
			menu.Nama = value.(string)
		case "deskripsi":
			menu.Deskripsi = value.(string)
		case "stok":
			menu.Stok = value.(int)
		case "harga":
			menu.Harga = value.(int)
		case "kategori":
			menu.Kategori = value.(string)
		case "uri_image":
			menu.URI_image = value.(string)
		}
	}

	if err := u.validator.Struct(menu); err != nil {
		return err
	}

	return u.repo.UpdateMenu(*menu)
}

func (u *menuUsecase) Delete(id uuid.UUID) error {
	return u.repo.DeleteMenu(id)
}
