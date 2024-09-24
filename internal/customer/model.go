package customer

import "github.com/google/uuid"

type Customer struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key"`
	Nama     string    `gorm:"size:50;not null"`
	Alamat   string    `gorm:"size:100;not null"`
	Telepon  string    `gorm:"size:13;not null"`
	Username string    `gorm:"size:10;not null"`
	Sandi    string    `gorm:"size:8;not null"`
	RoleID   uuid.UUID `gorm:"type:uuid;not null"`
	OrderID  uuid.UUID `gorm:"type:uuid"`
}
