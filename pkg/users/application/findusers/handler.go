package findusers

import (
	"context"

	"<MODULE_URL_REPLACE>/pkg/shared/domain/commandbus"
	"<MODULE_URL_REPLACE>/pkg/shared/domain/criteria"
	"<MODULE_URL_REPLACE>/pkg/shared/domain/errors"
)

type CommandHandler struct {
	Service Service
}

func NewCommandHandler(s Service) *CommandHandler {
	return &CommandHandler{
		Service: s,
	}
}

func (h *CommandHandler) Handle(ctx context.Context, cmd commandbus.Command) (commandbus.Response, errors.App) {

	command, ok := cmd.(Command)

	if !ok {
		return nil, errors.NewAppError(errors.UNEXPECTED_COMMAND_ERROR)
	}

	sortBy := criteria.NewSortBy(command.orderBy, command.orderSort)
	filters := criteria.NewFilters(command.filter)
	criteria := criteria.NewCriteria(filters, sortBy, criteria.Offset(command.page), criteria.Limit(command.perPage))

	return h.Service.exec(ctx, *criteria)
}

var _ commandbus.Handler = (*CommandHandler)(nil)
