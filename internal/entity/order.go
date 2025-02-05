package entity

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
    ID              uuid.UUID     `gorm:"type:uuid;primary_key"`
    CustomerID      uuid.UUID     `gorm:"type:uuid;not null"`
    Customer        Customer      `gorm:"foreignKey:CustomerID"`
    Items           []OrderItem   `gorm:"foreignKey:OrderID"`
    Status          string        `gorm:"type:varchar(20);not null;default:'pending'"`
    TotalAmount     float64       `gorm:"type:decimal(10,2);not null"`
    DeliveryAddress string        `gorm:"type:text;not null"`
    Notes           string        `gorm:"type:text"`
    CreatedAt       time.Time
    UpdatedAt       time.Time
    Payment         Payment       `gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
    ID        uuid.UUID `gorm:"type:uuid;primary_key"`
    OrderID   uuid.UUID `gorm:"type:uuid;not null"`
    MenuID    uuid.UUID `gorm:"type:uuid;not null"`
    Menu      Menu      `gorm:"foreignKey:MenuID"`
    Quantity  int       `gorm:"not null"`
    Price     float64   `gorm:"type:decimal(10,2);not null"`
    CreatedAt time.Time
    UpdatedAt time.Time
}