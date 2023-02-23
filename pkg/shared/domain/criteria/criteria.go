package criteria

import (
	"strings"

	"<MODULE_URL_REPLACE>/pkg/shared/domain/errors"
)

const (
	AMOUNT_OF_PARTS          = 3
	CONDITION_PARTS_SPLITTER = "::"
	CONDITION_SPLITTER       = ","
	PARAMETERS_SPLITTER      = "|"
	ZERO_CONDITIONS          = 0
)

type Filter struct {
	Field      string
	Operator   string
	Parameters []string
}

type Filters []Filter

type CriteriaFilter struct {
	Filters Filters
}

func (cf *CriteriaFilter) GetFilter(filter string) *Filter {
	for _, f := range cf.Filters {
		if f.Field == filter {
			return &f
		}
	}

	return nil
}

func NewCriteriaFilter(f string) (*CriteriaFilter, *errors.AppError) {

	filters, err := stringToFilters(f)

	if err != nil {
		return nil, err
	}

	return &CriteriaFilter{Filters: filters}, nil
}

func stringToFilters(input string) (Filters, *errors.AppError) {

	var filters Filters

	if len(input) == ZERO_CONDITIONS {
		return filters, nil
	}

	conditions := strings.Split(input, CONDITION_SPLITTER)

	for _, c := range conditions {
		parts := strings.Split(c, CONDITION_PARTS_SPLITTER)

		if len(parts) != AMOUNT_OF_PARTS {
			return Filters{}, errors.NewAppError(errors.MALFORMED_FILTER_ERROR)
		}

		if invalidOperator(parts[1]) {
			return Filters{}, errors.NewAppError(errors.INVALID_OPERATOR_FILTER_ERROR)
		}

		newCondtion := Filter{
			Field:      parts[0],
			Operator:   parts[1],
			Parameters: parametersToArray(parts[2]),
		}

		filters = append(filters, newCondtion)
	}

	return filters, nil
}

// The valid operators ar:
// eq = equal
// gt = gratter than
// gte = gratter and equal than
// lt = less than
// lte = less and equal than
// between = between two values
// in = a value in a list
// not-in = as in but not in a list
func invalidOperator(o string) bool {
	set := make(map[string]bool)
	list := []string{"eq", "gt", "gte", "lt", "lte", "between", "in", "not", "like"}

	for _, v := range list {
		set[v] = true
	}

	return !set[o]
}

func parametersToArray(parameters string) []string {
	array := strings.Split(parameters, PARAMETERS_SPLITTER)

	return array
}

type Criteria interface {
	Filter() interface{}
}

type Builder interface {
	GetFilters(filers string) (*CriteriaFilter, *errors.AppError)
	Auth(field string, value string) Criteria
	And(c1 Criteria, c2 Criteria) Criteria
	Between(field string, values any) Criteria
	Eq(field string, value string) Criteria
	Gt(field string, value any) Criteria
	Gte(field string, value any) Criteria
	In(field string, values any) Criteria
	Like(field string, value string) Criteria
	Lt(field string, value any) Criteria
	Lte(field string, value any) Criteria
	Not(field string, c Criteria) Criteria
	Or(c1 Criteria, c2 Criteria) Criteria
	Order(field string, sort string) Criteria
	Paginator(page int, perPage int) Criteria
}
