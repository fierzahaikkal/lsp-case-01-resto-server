package entity

import (
	"time"

	"github.com/google/uuid"
)

type Admin struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key"`
	Nama	string `gorm:"not null"`
	Email string `gorm:"not null"`
	Username string    `gorm:"not null"`
	Password string    `gorm:"not null"`
	RoleID   uuid.UUID `gorm:"type:uuid"`
	createdAt time.Time
	updatedAt time.Time
}