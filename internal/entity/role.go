package entity

import "github.com/google/uuid"

type Roles struct{
	ID uuid.UUID
	Name string `gorm:"size=10,not null"`
}