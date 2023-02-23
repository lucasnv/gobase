package finduser

import (
	"context"

	"<MODULE_URL_REPLACE>/pkg/shared/domain/commandbus"
	"<MODULE_URL_REPLACE>/pkg/shared/domain/errors"
	"<MODULE_URL_REPLACE>/pkg/shared/domain/valueobjects"
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

func (s *Service) exec(ctx *context.Context, id valueobjects.Id) (commandbus.Response, errors.App) {

	user, err := s.UserRepository.Find(ctx, id)

	if err != nil {
		return UserResponse{}, err
	}

	return NewUserResponse(user), nil
}
