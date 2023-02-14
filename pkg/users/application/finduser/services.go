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

// TODO: aca falta determinar si tengo que usar el contexto para los erores o implemento el error de forma tradcional
// El 404 es un error, tambien pudo tener el error de que no se puede acceder al repositorio
func (s *Service) exec(ctx context.Context, id valueobjects.Id) (commandbus.Reponse, errors.App) {

	user := s.UserRepository.Find(id)

	/*
		if err != nil {
			return user.NewUserError(user.REPOSITORY_USER_ERROR)
		}
	*/
	return NewUserResponse(user), nil
}
