package db

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Admin model
type Admin struct {
    ID       uuid.UUID `gorm:"type:uuid;primary_key"`
    Username string    `gorm:"not null"`
    Password string    `gorm:"not null"`
    RoleID   uuid.UUID `gorm:"type:uuid"`
}

// Peranan (Roles) model
type Peranan struct {
    ID    uuid.UUID `gorm:"type:uuid;primary_key"`
    Peran string    `gorm:"size:20;not null"`
}

// Kustomer (Customers) model
type Kustomer struct {
    ID       uuid.UUID `gorm:"type:uuid;primary_key"`
    Nama     string    `gorm:"size:50;not null"`
    Alamat   string    `gorm:"size:255;not null"`
    Telepon  string    `gorm:"size:13;not null"`
    Username string    `gorm:"not null"`
    Sandi    string    `gorm:"not null"`
    OrderID  uuid.UUID `gorm:"type:uuid"`
    RoleID   uuid.UUID `gorm:"type:uuid"`
}

// Menu (Food) model
type Menu struct {
    ID       uuid.UUID `gorm:"type:uuid;primary_key"`
    Nama     string    `gorm:"size:30;not null"`
    Stok     int       `gorm:"not null"`
    Kategori string    `gorm:"size:20;not null"`
    URI_image string `gorm:"not null"`  // Enum: 'menu utama', 'appetizer', 'minuman'
}

// Pesanan (Order) model
type Pesanan struct {
    ID             uuid.UUID `gorm:"type:uuid;primary_key"`
    StatusOrder    string    `gorm:"size:20;not null"`  // Enum: 'menunggu', 'dibuat', 'selesai'
    Jumlah         int       `gorm:"not null"`
    Harga          int       `gorm:"not null"`
    JenisTransaksi string    `gorm:"size:20;not null"`  // Enum: 'tunai', 'non tunai'
    UserID         uuid.UUID `gorm:"type:uuid"`
    MenuID      uuid.UUID `gorm:"type:uuid"`
}

func Migrate(db *gorm.DB) error {
    // Run the migrations
    return db.AutoMigrate(&Admin{}, &Peranan{}, &Kustomer{}, &Menu{}, &Pesanan{})
}
