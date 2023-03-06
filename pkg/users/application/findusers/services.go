package findusers

import (
	"context"

	"<MODULE_URL_REPLACE>/pkg/shared/domain/collection"
	"<MODULE_URL_REPLACE>/pkg/shared/domain/commandbus"
	"<MODULE_URL_REPLACE>/pkg/shared/domain/criteria"
	"<MODULE_URL_REPLACE>/pkg/shared/domain/errors"
	user "<MODULE_URL_REPLACE>/pkg/users/domain"
)

type Service struct {
	UserRepository user.UserRepository
}

func NewService(r user.UserRepository) *Service {
	return &Service{
		UserRepository: r,
	}
}

func (s *Service) exec(ctx *context.Context, f criteria.Criteria, o criteria.SortCriteria, p criteria.PaginatorCriteria) (commandbus.Response, errors.App) {

	users, err := s.UserRepository.FindByCriteria(ctx, f, o, p)

	if err != nil {
		return collection.Collection{}, err
	}

	return NewUsersResponse(users), nil
}
