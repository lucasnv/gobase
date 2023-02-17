package domain

import (
	"<MODULE_URL_REPLACE>/pkg/shared/domain/criteria"
	"<MODULE_URL_REPLACE>/pkg/shared/domain/errors"
	vo "<MODULE_URL_REPLACE>/pkg/shared/domain/valueobjects"
)

type UserRepository interface {
	Save(u User) *errors.AppError
	Find(id vo.Id) (*User, *errors.AppError)
	FindBy(c criteria.Criteria) (*Users, *errors.AppError)
	Delete(id vo.Id) *errors.AppError
}
