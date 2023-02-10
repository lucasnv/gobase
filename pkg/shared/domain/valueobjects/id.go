package valueobjects

import (
	"github.com/google/uuid"

	"<MODULE_URL_REPLACE>/pkg/shared/domain/errors"
)

// GENERIC ID VALUE OBJECT
type Id struct {
	value uuid.UUID
}

func NewId(value uuid.UUID) (Id, errors.App) {
	return Id{value: value}, nil
}

func GenerateNewId() (Id, errors.App) {
	return Id{value: uuid.New()}, nil
}
