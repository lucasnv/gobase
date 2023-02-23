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
	command, ok := cmd.(Command)

	if !ok {
		return nil, errors.NewAppError(errors.UNEXPECTED_COMMAND_ERROR)
	}

	filters, err := h.CriteriaBuilder.GetFilters(command.filter)

	if err != nil {
		return nil, err
	}

	var finalCriteria criteria.Criteria
	var qtyCriteria int

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

/*
	if email := filters.GetFilter("email"); email != nil {

		if email.Operator != "eq" {
			return UsersResponse{}, errors.NewAppError(errors.INVALID_OPERATOR_FILTER_ERROR)
		}

		emailCriteria = h.CriteriaBuilder.Eq("email", email.Parameters[0])
	}

	if firstName := filters.GetFilter("first_name"); firstName != nil {

		if email.Operator != "eq" {
			return UsersResponse{}, errors.NewAppError(errors.INVALID_OPERATOR_FILTER_ERROR)
		}

		emailCriteria = h.CriteriaBuilder.Eq("email", email.Parameters[0])
	}

	if createdAt := filters.GetFilter("created_at"); createdAt != nil {
		var err errors.App
		createdAtCriteria, err = h.GetCriteriaByFilter(*createdAt)

		if err != nil {
			return UsersResponse{}, err
		}

	} else {
		now := time.Now().UTC()
		daysAgo := now.AddDate(0, 0, -15)
		createdAtCriteria = h.CriteriaBuilder.Gte("created_at", daysAgo.Format(time.RFC3339))
	}

	if emailCriteria != nil {
		criteria = h.CriteriaBuilder.And(emailCriteria, createdAtCriteria)
	} else {
		criteria = createdAtCriteria
	}
*/
