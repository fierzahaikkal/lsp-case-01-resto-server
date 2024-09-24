package makanan

import "github.com/google/uuid"

type Makanan struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key"`
	Nama     string    `gorm:"size:30;not null"`
	Stok     int       `gorm:"not null"`
	Kategori string    `gorm:"type:enum('makanan utama', 'appetizer', 'minuman');not null"`
}
