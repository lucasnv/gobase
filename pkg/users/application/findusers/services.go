package findusers

import (
	"context"

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

func (s *Service) exec(ctx *context.Context, c criteria.Criteria) (commandbus.Response, errors.App) {

	users, err := s.UserRepository.FindByCriteria(ctx, c)

	if err != nil {
		return UsersResponse{}, err
	}

	return NewUsersResponse(users), nil
}
