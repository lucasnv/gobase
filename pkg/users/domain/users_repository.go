package domain

import (
	"context"

	"<MODULE_URL_REPLACE>/pkg/shared/domain/criteria"
	"<MODULE_URL_REPLACE>/pkg/shared/domain/errors"
	vo "<MODULE_URL_REPLACE>/pkg/shared/domain/valueobjects"
)

type UserRepository interface {
	Save(ctx *context.Context, u User) *errors.AppError
	Find(ctx *context.Context, id vo.Id) (User, *errors.AppError)
	FindByCriteria(ctx *context.Context, c criteria.Criteria) (Users, *errors.AppError)
	Delete(ctx *context.Context, id vo.Id) *errors.AppError
}
