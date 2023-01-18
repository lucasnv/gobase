package createtool

import (
	"context"
	"errors"

	"<MODULE_URL_REPLACE>/pkg/shared/infrastructure/commandbus"
)

type CommandHandler struct {
	service Service
}

func NewCommandHandler(service Service) CommandHandler {
	return CommandHandler{
		service: service,
	}
}

func (h CommandHandler) Handle(ctx context.Context, cmd commandbus.Command) error {
	command, ok := cmd.(Command)

	if !ok {
		return errors.New("unexpected command")
	}

	return h.service.exec(
		ctx,
		command.id,
		command.name,
		command.link,
		command.description,
	)
}
