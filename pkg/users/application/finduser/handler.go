package finduser

import (
	"context"

	"<MODULE_URL_REPLACE>/pkg/shared/domain/commandbus"
	"<MODULE_URL_REPLACE>/pkg/shared/domain/errors"
	"<MODULE_URL_REPLACE>/pkg/shared/domain/valueobjects"
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

func (h CommandHandler) Handle(ctx context.Context, cmd commandbus.Command) (commandbus.Reponse, errors.App) {

	command, ok := cmd.(Command)

	if !ok {
		return nil, errors.NewAppError(errors.UNEXPECTED_COMMAND_ERROR)
	}

	uuid, err := valueobjects.NewId(command.id)

	if err != nil {
		return nil, err
	}

	return h.Service.exec(ctx, uuid)
}

var _ commandbus.Handler = (*CommandHandler)(nil)