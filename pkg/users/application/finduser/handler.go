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

	id, err := valueobjects.NewIdFromString(command.id)

	if err != nil {
		return nil, err
	}

	return h.Service.exec(ctx, id)
}

var _ commandbus.Handler = (*CommandHandler)(nil)
