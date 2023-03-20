package findusers

import (
	"context"

	"<MODULE_URL_REPLACE>/pkg/shared/domain/commandbus"
	"<MODULE_URL_REPLACE>/pkg/shared/domain/criteria"
	"<MODULE_URL_REPLACE>/pkg/shared/domain/errors"
)

type CommandHandler struct {
	Service         Service
	CriteriaBuilder criteria.Builder
}

func NewCommandHandler(s Service) *CommandHandler {
	return &CommandHandler{
		Service:         s,
		CriteriaBuilder: criteria.NewCriterBuilder(),
	}
}

func (cmd *CommandHandler) Handle(ctx *context.Context, bus commandbus.Command) (commandbus.Response, errors.App) {
	command, ok := bus.(Command)

	if !ok {
		return nil, errors.NewAppError(errors.UNEXPECTED_COMMAND_ERROR)
	}

	filterCriteria, err := cmd.CriteriaBuilder.FiltersToCriteria(command.filter)

	if err != nil {
		return nil, err
	}

	sortCriteria := cmd.CriteriaBuilder.Sort(command.orderBy, command.orderSort)
	paginatorCriteria := cmd.CriteriaBuilder.Paginator(command.page, command.perPage)

	return cmd.Service.exec(ctx, *filterCriteria, sortCriteria, paginatorCriteria)
}

var _ commandbus.Handler = (*CommandHandler)(nil)
