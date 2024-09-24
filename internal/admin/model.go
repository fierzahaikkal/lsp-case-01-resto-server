package admin

import "github.com/google/uuid"

type Admin struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key"`
	Username string    `gorm:"size:10;not null"`
	Password string    `gorm:"size:8;not null"`
	RoleID   uuid.UUID `gorm:"type:uuid"`
}
