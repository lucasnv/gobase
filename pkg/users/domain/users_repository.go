package domain

import (
	"context"

	"<MODULE_URL_REPLACE>/pkg/shared/domain/collection"
	"<MODULE_URL_REPLACE>/pkg/shared/domain/criteria"
	"<MODULE_URL_REPLACE>/pkg/shared/domain/errors"
	vo "<MODULE_URL_REPLACE>/pkg/shared/domain/valueobjects"
)

type UserRepository interface {
	Save(ctx *context.Context, u User) *errors.AppError
	Find(ctx *context.Context, id vo.Id) (User, *errors.AppError)
	FindByCriteria(ctx *context.Context, c criteria.Criteria, o criteria.SorterCriteria, p criteria.PaginatorCriteria) (collection.Collection, *errors.AppError)
	Delete(ctx *context.Context, id vo.Id) *errors.AppError
}
