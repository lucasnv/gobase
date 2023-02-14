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

func (s Service) exec(ctx context.Context, id valueobjects.Id) (commandbus.Reponse, errors.App) {

	user := s.UserRepository.Find(id)

	/*
		if err != nil {
			return user.NewUserError(user.REPOSITORY_USER_ERROR)
		}
	*/
	return user, nil
}
