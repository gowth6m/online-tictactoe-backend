package utils

import "github.com/google/uuid"

// GenerateID generates a new UUID
func GenerateID() string {
	return uuid.New().String()
}
