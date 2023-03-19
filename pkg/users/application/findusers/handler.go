package findusers

import (
	"context"

	"<MODULE_URL_REPLACE>/pkg/shared/domain/collection"
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
	var finalCriteria criteria.Criteria
	var qtyCriteria int

	command, ok := bus.(Command)

	if !ok {
		return nil, errors.NewAppError(errors.UNEXPECTED_COMMAND_ERROR)
	}

	filters, err := cmd.CriteriaBuilder.GetFilters(command.filter)

	if err != nil {
		return nil, err
	}

	for _, filter := range filters.Filters {

		criteria, err := cmd.GetCriteriaByFilter(filter)

		if err != nil {
			return collection.Collection{}, err
		}

		if qtyCriteria == 0 {
			finalCriteria = criteria
			qtyCriteria++
			continue
		}

		finalCriteria = cmd.CriteriaBuilder.And(finalCriteria, criteria)
		qtyCriteria++
	}

	sortCriteria := cmd.CriteriaBuilder.Sort(command.orderBy, command.orderSort)
	paginatorCriteria := cmd.CriteriaBuilder.Paginator(command.page, command.perPage)

	return cmd.Service.exec(ctx, finalCriteria, sortCriteria, paginatorCriteria)
}

func (cmd *CommandHandler) GetCriteriaByFilter(filter criteria.Filter) (criteria.Criteria, errors.App) {
	var criteria criteria.Criteria

	switch filter.Operator {
	case "eq":
		criteria = cmd.CriteriaBuilder.Eq(filter.Field, filter.Parameters[0])

	case "lt":
		criteria = cmd.CriteriaBuilder.Lt(filter.Field, filter.Parameters[0])

	case "lte":
		criteria = cmd.CriteriaBuilder.Lte(filter.Field, filter.Parameters[0])

	case "gt":
		criteria = cmd.CriteriaBuilder.Gt(filter.Field, filter.Parameters[0])

	case "gte":
		criteria = cmd.CriteriaBuilder.Gte(filter.Field, filter.Parameters[0])

	case "in":
		criteria = cmd.CriteriaBuilder.In(filter.Field, filter.Parameters)

	case "like":
		criteria = cmd.CriteriaBuilder.Like(filter.Field, filter.Parameters[0])

	case "between":
		if len(filter.Parameters) != 2 {
			return criteria, errors.NewAppError(errors.MALFORMED_FILTER_ERROR)
		}

		criteria = cmd.CriteriaBuilder.Between(filter.Field, filter.Parameters)

	default:
		return criteria, errors.NewAppError(errors.INVALID_OPERATOR_FILTER_ERROR)
	}

	return criteria, nil
}

var _ commandbus.Handler = (*CommandHandler)(nil)
