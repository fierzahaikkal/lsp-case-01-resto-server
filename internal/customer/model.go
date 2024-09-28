package customer

import "github.com/google/uuid"

type Customer struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key"`
	Nama     string    `gorm:"size:50;not null"`
	Alamat   string    `gorm:"size:100;not null"`
	Telepon  string    `gorm:"size:13;not null"`
	Username string    `gorm:"not null"`
	Sandi    string    `gorm:"not null"`
	RoleID   uuid.UUID `gorm:"type:uuid;not null"`
	OrderID  uuid.UUID `gorm:"type:uuid"`
}
