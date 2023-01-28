package registeruser

import (
	"context"

	internalError "<MODULE_URL_REPLACE>/pkg/shared/domain/errors"
	"<MODULE_URL_REPLACE>/pkg/shared/domain/valueobjects"
	"<MODULE_URL_REPLACE>/pkg/shared/infrastructure/commandbus"
	"<MODULE_URL_REPLACE>/pkg/users/domain"
)

type CommandHandler struct {
	Service Service
}

/*
func NewCommandHandler(service Service) CommandHandler {
	return CommandHandler{
		service: service,
	}
}
*/
// TODO : Data sanitization
func (h CommandHandler) Handle(ctx context.Context, cmd commandbus.Command) error {

	command, ok := cmd.(Command)

	if !ok {
		return internalError.NewUnexpectedCommand()
	}

	id, err := valueobjects.GenerateNewId()

	if err != nil {
		return err
	}

	firstName, err := domain.NewFirstName(command.firstName)

	if err != nil {
		return err
	}

	lastName, err := domain.NewLastName(command.lastName)

	if err != nil {
		return err
	}

	email, err := domain.NewEmail(command.email)

	if err != nil {
		return err
	}

	password, err := domain.NewPassword(command.password)

	if err != nil {
		return err
	}

	return h.Service.exec(
		ctx,
		id,
		firstName,
		lastName,
		email,
		password,
	)
}

var _ commandbus.Handler = (*CommandHandler)(nil)
