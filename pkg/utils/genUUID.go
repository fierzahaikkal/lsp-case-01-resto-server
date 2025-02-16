package utils

import "github.com/google/uuid"

func GenUUID() uuid.UUID {
	return uuid.New()
}