package valueobjects

import (
	"github.com/google/uuid"

	"<MODULE_URL_REPLACE>/pkg/shared/domain/errors"
)

type Id struct {
	Value uuid.UUID
}

func NewId(value string) (Id, errors.App) {
	u, err := uuid.Parse(value)

	if err != nil {
		return Id{Value: uuid.New()}, errors.NewAppError(errors.INVALID_UUID_ERROR)
	}

	return Id{Value: u}, nil
}

func GenerateNewId() (Id, errors.App) {
	return Id{Value: uuid.New()}, nil
}

func (i *Id) ToString() string {
	return i.Value.String()
}
