package valueobjects

import (
	"github.com/google/uuid"
)

// GENERIC ID VALUE OBJECT
type Id struct {
	value uuid.UUID
}

func NewId(value uuid.UUID) (Id, error) {
	return Id{value: value}, nil
}

func GenerateNewId() (Id, error) {
	return Id{value: uuid.New()}, nil
}
