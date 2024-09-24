package pesanan

import "github.com/google/uuid"

type Pesanan struct {
	ID             uuid.UUID `gorm:"type:uuid;primary_key"`
	StatusOrder    string    `gorm:"type:enum('menunggu', 'dibuat', 'selesai');not null"`
	Jumlah         int       `gorm:"not null"`
	Harga          int       `gorm:"not null"`
	JenisTransaksi string    `gorm:"type:enum('tunai', 'non tunai');not null"`
	CustomerID     uuid.UUID `gorm:"type:uuid;not null"`
	MakananID      uuid.UUID `gorm:"type:uuid;not null"`
}
