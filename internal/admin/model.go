package admin

import "github.com/google/uuid"

type Admin struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key"`
	Username string    `gorm:"not null"`
	Password string    `gorm:"not null"`
	RoleID   uuid.UUID `gorm:"type:uuid"`
}
