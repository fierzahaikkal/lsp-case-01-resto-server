package utils

import (
	"errors"

	"github.com/google/uuid"
)

func ParseUUID(id string) (uuid.UUID, error) {
    if err := uuid.Validate(id); err != nil {
        return uuid.Nil, errors.New("invalid UUID format")
    }
    
    return uuid.MustParse(id), nil
}

// Helper function to validate multiple UUIDs at once
func ValidateUUIDs(ids ...string) error {
    for _, id := range ids {
        if err := uuid.Validate(id); err != nil {
            return errors.New("invalid UUID format: " + id)
        }
    }
    return nil
}

// Helper function to convert string slice to UUID slice
func ParseUUIDs(ids []string) ([]uuid.UUID, error) {
    uuids := make([]uuid.UUID, len(ids))
    
    for i, id := range ids {
        if err := uuid.Validate(id); err != nil {
            return nil, errors.New("invalid UUID format at index " + string(i))
        }
        uuids[i] = uuid.MustParse(id)
    }
    
    return uuids, nil
}

// Helper function to check if a UUID is zero/empty
func IsEmptyUUID(id uuid.UUID) bool {
    return id == uuid.UUID{}
}
