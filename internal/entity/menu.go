package entity

import (
	"time"

	"github.com/google/uuid"
)

type Menu struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key"`
	Nama      string    `gorm:"size:30;not null"`
	Deskripsi string	`gorm:"null"`
	Stok      int       `gorm:"not null"`
	Harga	  int		`gorm:"not null"`
	Kategori  string    `gorm:"type:enum('menu utama', 'appetizer', 'minuman');not null"`
	URI_image string    `gorm:"not null"`
	createdAt time.Time
	updatedAt time.Time
}
