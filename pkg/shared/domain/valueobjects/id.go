package valueobjects

import (
	"github.com/google/uuid"

	"<MODULE_URL_REPLACE>/pkg/shared/domain/errors"
)

// Generic Id Value Object
type Id struct {
	value uuid.UUID
}

func (i *Id) Value() string {
	return i.value.String()
}

func NewId(value string) (Id, errors.App) {
	u, err := uuid.Parse(value)

	if err != nil {
		return Id{value: uuid.New()}, errors.NewAppError(errors.INVALID_UUID_ERROR)
	}

	return Id{value: u}, nil
}

func GenerateNewId() (Id, errors.App) {
	return Id{value: uuid.New()}, nil
}
