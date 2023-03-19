package registeruser

import (
	"context"

	"<MODULE_URL_REPLACE>/pkg/shared/domain/commandbus"
	"<MODULE_URL_REPLACE>/pkg/shared/domain/errors"
	"<MODULE_URL_REPLACE>/pkg/shared/domain/valueobjects"
	"<MODULE_URL_REPLACE>/pkg/users/domain"
)

type CommandHandler struct {
	Service Service
}

func NewCommandHandler(s Service) *CommandHandler {
	return &CommandHandler{
		Service: s,
	}
}

func (h *CommandHandler) Handle(ctx *context.Context, cmd commandbus.Command) (commandbus.Response, errors.App) {

	command, ok := cmd.(Command)

	if !ok {
		return nil, errors.NewAppError(errors.UNEXPECTED_COMMAND_ERROR)
	}

	firstName, err := domain.NewFirstName(command.firstName)

	if err != nil {
		return nil, err
	}

	lastName, err := domain.NewLastName(command.lastName)

	if err != nil {
		return nil, err
	}

	email, err := domain.NewEmail(command.email)

	if err != nil {
		return nil, err
	}

	return nil, h.Service.exec(
		ctx,
		valueobjects.NewId(),
		firstName,
		lastName,
		email,
	)
}

var _ commandbus.Handler = (*CommandHandler)(nil)
