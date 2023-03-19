package updateuser

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

func (s *Service) exec(ctx *context.Context, id valueobjects.Id, fn user.FirstName, ln user.LastName, e user.Email) errors.App {
	user, err := s.UserRepository.Find(ctx, id)

	if err != nil {
		return err
	}

	userUpdated := user.UpdateProfile(fn, ln, e)

	err = s.UserRepository.Save(ctx, userUpdated)

	if err != nil {
		return err
	}

	return nil
}
