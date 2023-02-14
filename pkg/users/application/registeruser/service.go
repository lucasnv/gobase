package registeruser

import (
	"context"

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

func (s Service) exec(ctx context.Context, id valueobjects.Id, fn user.FirstName, ln user.LastName, e user.Email, p user.Password) errors.App {

	var newUser user.User = user.NewUser(id, fn, ln, e, p)

	err := s.UserRepository.Save(newUser)

	if err != nil {
		return user.NewUserError(user.REPOSITORY_USER_ERROR)
	}

	return nil
}
