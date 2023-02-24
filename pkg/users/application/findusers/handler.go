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

func NewCommandHandler(s Service, cb criteria.Builder) *CommandHandler {
	return &CommandHandler{
		Service:         s,
		CriteriaBuilder: cb,
	}
}

func (h *CommandHandler) Handle(ctx *context.Context, cmd commandbus.Command) (commandbus.Response, errors.App) {
	var finalCriteria criteria.Criteria
	var qtyCriteria int

	command, ok := cmd.(Command)

	if !ok {
		return nil, errors.NewAppError(errors.UNEXPECTED_COMMAND_ERROR)
	}

	filters, err := h.CriteriaBuilder.GetFilters(command.filter)

	if err != nil {
		return nil, err
	}

	for _, filter := range filters.Filters {

		criteria, err := h.GetCriteriaByFilter(filter)

		if err != nil {
			return UsersResponse{}, errors.NewAppError(errors.INVALID_OPERATOR_FILTER_ERROR)
		}

		if qtyCriteria == 0 {
			finalCriteria = criteria
			continue
		}

		finalCriteria = h.CriteriaBuilder.And(finalCriteria, criteria)
	}

	return h.Service.exec(ctx, finalCriteria)
}

var _ commandbus.Handler = (*CommandHandler)(nil)

func (h *CommandHandler) GetCriteriaByFilter(filter criteria.Filter) (criteria.Criteria, errors.App) {
	var criteria criteria.Criteria

	switch filter.Operator {
	case "eq":
		criteria = h.CriteriaBuilder.Eq(filter.Field, filter.Parameters[0])

	case "lt":
		criteria = h.CriteriaBuilder.Lt(filter.Field, filter.Parameters[0])

	case "lte":
		criteria = h.CriteriaBuilder.Lte(filter.Field, filter.Parameters[0])

	case "gt":
		criteria = h.CriteriaBuilder.Gt(filter.Field, filter.Parameters[0])

	case "gte":
		criteria = h.CriteriaBuilder.Gte(filter.Field, filter.Parameters[0])

	case "in":
		criteria = h.CriteriaBuilder.In(filter.Field, filter.Parameters)

	case "like":
		criteria = h.CriteriaBuilder.Like(filter.Field, filter.Parameters[0])

	case "between":
		criteria = h.CriteriaBuilder.Between(filter.Field, filter.Parameters)

	default:
		return criteria, errors.NewAppError(errors.INVALID_OPERATOR_FILTER_ERROR)
	}

	return criteria, nil

}
