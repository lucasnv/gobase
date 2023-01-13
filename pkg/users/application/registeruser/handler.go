package registeruser

import (
	"context"

	"[MODULE_URL]/internal/commandbus"
	internalError "[MODULE_URL]/pkg/shared/domain/errors"
	"[MODULE_URL]/pkg/shared/domain/valueobjects"
	users "[MODULE_URL]/pkg/users/domain"
)

type CommandHandler struct {
	service Service
}

func NewCommandHandler(service Service) CommandHandler {
	return CommandHandler{
		service: service,
	}
}

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

	firstName, err := users.NewFirstName(command.firstName)

	if err != nil {
		return err
	}

	lastName, err := users.NewLastName(command.lastName)

	if err != nil {
		return err
	}

	email, err := users.NewEmail(command.email)

	if err != nil {
		return err
	}

	password, err := users.NewPassword(command.password)

	if err != nil {
		return err
	}

	return h.service.exec(
		ctx,
		id,
		firstName,
		lastName,
		email,
		password,
	)
}
