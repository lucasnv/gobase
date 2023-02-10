package registeruser

import (
	"context"
	"fmt"

	"<MODULE_URL_REPLACE>/pkg/shared/domain/errors"
	"<MODULE_URL_REPLACE>/pkg/shared/domain/valueobjects"
	"<MODULE_URL_REPLACE>/pkg/shared/infrastructure/commandbus"
	"<MODULE_URL_REPLACE>/pkg/users/domain"
)

type Service struct {
	UserRepository domain.UserRepository
}

func (s Service) exec(ctx context.Context, id valueobjects.Id, firstName domain.FirstName, lastName domain.LastName, email domain.Email, password domain.Password) errors.App {
	fmt.Println("exec service registeruser")
	s.UserRepository.Save()
	return nil
}

var _ commandbus.Handler = (*CommandHandler)(nil)
