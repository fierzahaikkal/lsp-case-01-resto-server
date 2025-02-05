package entity

import (
	"time"

	"github.com/google/uuid"
)

type Payment struct {
    ID            uuid.UUID `gorm:"type:uuid;primary_key"`
    OrderID       uuid.UUID `gorm:"type:uuid;not null;unique"`
    Amount        float64   `gorm:"type:decimal(10,2);not null"`
    PaymentMethod string    `gorm:"size:20;not null"`
    Status        string    `gorm:"size:20;not null;default:'pending'"`
    PaymentDate   *time.Time
    CreatedAt     time.Time
    UpdatedAt     time.Time
}