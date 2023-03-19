package valueobjects

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"<MODULE_URL_REPLACE>/pkg/shared/domain/errors"
)

type Id struct {
	Value primitive.ObjectID
}

func NewId() Id {
	return Id{Value: primitive.NewObjectID()}
}

func NewIdFromString(value string) (Id, errors.App) {

	newId, err := primitive.ObjectIDFromHex(value)

	if err != nil {
		return Id{Value: newId}, errors.NewAppError(errors.INVALID_UUID_ERROR)
	}

	return Id{Value: newId}, nil
}

func (i *Id) ToString() string {
	return i.Value.Hex()
}
