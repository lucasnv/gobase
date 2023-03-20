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

type Criteria interface {
	Type() string
}

type SorterCriteria interface {
	By() string
	Sort() string
}

type PaginatorCriteria interface {
	Limit() uint32
	Offset() uint32
	Page() uint32
	PageSize() uint32
}

type Builder interface {
	FiltersToCriteria(filters string) (*Criteria, *errors.AppError)
	And(c1 Criteria, c2 Criteria) Criteria
	Between(field string, values []string) Criteria
	Eq(field string, value string) Criteria
	Gt(field string, value any) Criteria
	Gte(field string, value any) Criteria
	In(field string, values any) Criteria
	Like(field string, value string) Criteria
	Lt(field string, value any) Criteria
	Lte(field string, value any) Criteria
	Not(field string, c Criteria) Criteria
	Or(c1 Criteria, c2 Criteria) Criteria
	Sort(field string, order string) SorterCriteria
	Paginator(page uint32, perPage uint32) PaginatorCriteria
}

type CriteriaBuilder struct {
}

func NewCriterBuilder() CriteriaBuilder {
	return CriteriaBuilder{}
}

var _ Builder = (*CriteriaBuilder)(nil)

// Get filters from uri
func (builder CriteriaBuilder) FiltersToCriteria(f string) (*Criteria, *errors.AppError) {
	var finalCriteria Criteria
	var qtyCriteria int

	filters, err := getFilters(f)

	if err != nil {
		return nil, err
	}

	for _, filter := range filters.Filters {
		criteria, err := builder.getCriteriaByFilter(filter)

		if err != nil {
			return nil, err
		}

		if qtyCriteria == 0 {
			finalCriteria = criteria
			qtyCriteria++
			continue
		}

		finalCriteria = builder.And(finalCriteria, criteria)
		qtyCriteria++
	}

	return &finalCriteria, nil
}

func (builder CriteriaBuilder) getCriteriaByFilter(filter Filter) (Criteria, *errors.AppError) {
	var criteria Criteria

	switch filter.Operator {
	case "eq":
		criteria = builder.Eq(filter.Field, filter.Parameters[0])

	case "lt":
		criteria = builder.Lt(filter.Field, filter.Parameters[0])

	case "lte":
		criteria = builder.Lte(filter.Field, filter.Parameters[0])

	case "gt":
		criteria = builder.Gt(filter.Field, filter.Parameters[0])

	case "gte":
		criteria = builder.Gte(filter.Field, filter.Parameters[0])

	case "in":
		criteria = builder.In(filter.Field, filter.Parameters)

	case "like":
		criteria = builder.Like(filter.Field, filter.Parameters[0])

	case "between":
		criteria = builder.Between(filter.Field, filter.Parameters)

	default:
		return criteria, errors.NewAppError(errors.INVALID_OPERATOR_FILTER_ERROR)
	}

	return criteria, nil
}

// Examples of valid filter string:
// - page=1&per_page=20
// - sort_by=name&sort_order=asc&page=1&per_page=20
// - filter=name::eq::lucas vazquez&page=1&per_page=20
// - filter=name::eq::lucas vazquez&sort_by=name&sort_order=asc
// - filter=date::between::2023-03-19T13:24:21Z|2023-03-19T13:24:21Z
// - filter=status::in::active|suspended
// - filter=name::eq::lucas vazquez,age::lte::50&sort_by=name&sort_order=asc&page=1&per_page=20
// --------------------
// The parameter filter has the folling struct
// [criteria]::[operator]::[parameters],[criteria]::[operator]::[parameters]
func getFilters(f string) (*CriteriaFilter, *errors.AppError) {
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

// And Criteria
func (CriteriaBuilder) And(c1 Criteria, c2 Criteria) Criteria {
	return AndCriteria{
		C1: c1,
		C2: c2,
	}
}

type AndCriteria struct {
	C1 Criteria
	C2 Criteria
}

func (AndCriteria) Type() string {
	return "and"
}

// Between criteria
func (CriteriaBuilder) Between(f string, values []string) Criteria {

	if len(values) != 2 {
		return BetweenCriteria{
			Field: f,
			V1:    "",
			V2:    "",
		}
	}

	return BetweenCriteria{
		Field: f,
		V1:    values[0],
		V2:    values[1],
	}
}

type BetweenCriteria struct {
	Field string
	V1    any
	V2    any
}

func (BetweenCriteria) Type() string {
	return "between"
}

// Equal Criteria
func (CriteriaBuilder) Eq(field string, value string) Criteria {
	return EqCriteria{
		Field: field,
		Value: value,
	}
}

type EqCriteria struct {
	Field string
	Value any
}

func (EqCriteria) Type() string {
	return "eq"
}

// Gt Criteria
func (CriteriaBuilder) Gt(field string, value any) Criteria {
	return GtCriteria{
		Field: field,
		Value: value,
	}
}

type GtCriteria struct {
	Field string
	Value any
}

func (GtCriteria) Type() string {
	return "gt"
}

// Gte Criteria
func (CriteriaBuilder) Gte(field string, value any) Criteria {
	return GteCriteria{
		Field: field,
		Value: value,
	}
}

type GteCriteria struct {
	Field string
	Value any
}

func (GteCriteria) Type() string {
	return "gte"
}

// In Criteria
func (CriteriaBuilder) In(field string, value any) Criteria {
	return InCriteria{
		Field: field,
		Value: value,
	}
}

type InCriteria struct {
	Field string
	Value any
}

func (InCriteria) Type() string {
	return "in"
}

// Like Criteria
func (CriteriaBuilder) Like(field string, value string) Criteria {
	return LikeCriteria{
		Field: field,
		Value: value,
	}
}

type LikeCriteria struct {
	Field string
	Value string
}

func (LikeCriteria) Type() string {
	return "like"
}

// Lt Criteria
func (CriteriaBuilder) Lt(field string, value any) Criteria {
	return LtCriteria{
		Field: field,
		Value: value,
	}
}

type LtCriteria struct {
	Field string
	Value any
}

func (LtCriteria) Type() string {
	return "lt"
}

// Lte Criteria
func (CriteriaBuilder) Lte(field string, value any) Criteria {
	return LteCriteria{
		Field: field,
		Value: value,
	}
}

type LteCriteria struct {
	Field string
	Value any
}

func (LteCriteria) Type() string {
	return "lte"
}

// Not Criteria
func (CriteriaBuilder) Not(field string, c Criteria) Criteria {
	return NotCriteria{
		Field:    field,
		Criteria: c,
	}
}

type NotCriteria struct {
	Field    string
	Criteria Criteria
}

func (NotCriteria) Type() string {
	return "not"
}

// Or criteria
func (CriteriaBuilder) Or(c1 Criteria, c2 Criteria) Criteria {
	return OrCriteria{
		C1: c1,
		C2: c2,
	}
}

type OrCriteria struct {
	C1 Criteria
	C2 Criteria
}

func (OrCriteria) Type() string {
	return "or"
}

// Sort Criteria
func (CriteriaBuilder) Sort(field string, order string) SorterCriteria {
	return SortCriteria{
		Field: field,
		Order: order,
	}
}

type SortCriteria struct {
	Field string
	Order string
}

func (c SortCriteria) By() string {
	return c.Field
}

func (c SortCriteria) Sort() string {
	return c.Order
}

var _ SorterCriteria = (*SortCriteria)(nil)

// Paginator Criteria
func (CriteriaBuilder) Paginator(page uint32, pageSize uint32) PaginatorCriteria {

	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	return PaginateCriteria{
		offset:   (page - 1) * pageSize,
		limit:    pageSize,
		page:     page,
		pageSize: pageSize,
	}
}

type PaginateCriteria struct {
	limit    uint32
	offset   uint32
	page     uint32
	pageSize uint32
}

func (c PaginateCriteria) Limit() uint32 {
	return c.limit
}

func (c PaginateCriteria) Offset() uint32 {
	return c.offset
}

func (c PaginateCriteria) Page() uint32 {
	return c.page
}

func (c PaginateCriteria) PageSize() uint32 {
	return c.pageSize
}

var _ PaginatorCriteria = (*PaginateCriteria)(nil)
