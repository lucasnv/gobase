package criteria

import (
	"<MODULE_URL_REPLACE>/pkg/shared/domain/criteria"
	"<MODULE_URL_REPLACE>/pkg/shared/domain/errors"
)

type InmemoryBuilder struct {
	order  string
	limit  int
	offset int
}

func (InmemoryBuilder) GetFilters(f string) (*criteria.CriteriaFilter, *errors.AppError) {
	return criteria.NewCriteriaFilter(f)
}

func (InmemoryBuilder) And(c1 criteria.Criteria, c2 criteria.Criteria) criteria.Criteria {
	return AndInmemoryCriteria{
		C1: c1,
		C2: c2,
	}
}

func (InmemoryBuilder) Owner(field string, value string) criteria.Criteria {
	return OwnerInmemoryCriteria{
		Field: field,
		Value: value,
	}
}

func (InmemoryBuilder) Between(f string, values []string) criteria.Criteria {

	if len(values) != 2 {
		return BetweenInmemoryCriteria{
			Field: f,
			V1:    "",
			V2:    "",
		}
	}

	return BetweenInmemoryCriteria{
		Field: f,
		V1:    values[0],
		V2:    values[1],
	}
}

func (InmemoryBuilder) Eq(field string, value string) criteria.Criteria {
	return EqInmemoryCriteria{
		Field: field,
		Value: value,
	}
}

func (InmemoryBuilder) Gt(field string, value any) criteria.Criteria {
	return GtInmemoryCriteria{
		Field: field,
		Value: value,
	}
}

func (InmemoryBuilder) Gte(field string, value any) criteria.Criteria {
	return GteInmemoryCriteria{
		Field: field,
		Value: value,
	}
}

func (InmemoryBuilder) In(field string, values []string) criteria.Criteria {
	return InInmemoryCriteria{}
}

func (InmemoryBuilder) Like(field string, value string) criteria.Criteria {
	return LikeInmemoryCriteria{
		Field: field,
		Value: value,
	}
}

func (InmemoryBuilder) Lt(field string, value any) criteria.Criteria {
	return LtInmemoryCriteria{
		Field: field,
		Value: value,
	}
}

func (InmemoryBuilder) Lte(field string, value any) criteria.Criteria {
	return LteInmemoryCriteria{
		Field: field,
		Value: value,
	}
}

func (InmemoryBuilder) Not(field string, c criteria.Criteria) criteria.Criteria {
	return NotInmemoryCriteria{
		Field:    field,
		Criteria: c,
	}
}

func (InmemoryBuilder) Or(c1 criteria.Criteria, c2 criteria.Criteria) criteria.Criteria {
	return OrInmemoryCriteria{
		C1: c1,
		C2: c2,
	}
}

func (InmemoryBuilder) Sort(field string, order string) criteria.SortCriteria {
	return SortInmemoryCriteria{
		Field: field,
		Order: order,
	}
}

func (InmemoryBuilder) Paginator(page int, perPage int) criteria.PaginatorCriteria {
	return PaginatorInmemoryCriteria{}
}

func NewInmemoryBuilder() (criteria.Builder, *errors.AppError) {
	return &InmemoryBuilder{}, nil
}

var _ criteria.Builder = (*InmemoryBuilder)(nil)

type AndInmemoryCriteria struct {
	C1 criteria.Criteria
	C2 criteria.Criteria
}

func (c AndInmemoryCriteria) Filter() interface{} {
	return "and"
}

var _ criteria.Criteria = (*AndInmemoryCriteria)(nil)

type OwnerInmemoryCriteria struct {
	Field string
	Value any
}

func (c OwnerInmemoryCriteria) Filter() interface{} {
	return "owner"
}

var _ criteria.Criteria = (*OwnerInmemoryCriteria)(nil)

type EqInmemoryCriteria struct {
	Field string
	Value any
}

func (c EqInmemoryCriteria) Filter() interface{} {
	return "eq"
}

var _ criteria.Criteria = (*EqInmemoryCriteria)(nil)

type GtInmemoryCriteria struct {
	Field string
	Value any
}

func (c GtInmemoryCriteria) Filter() interface{} {
	return "gt"
}

var _ criteria.Criteria = (*GtInmemoryCriteria)(nil)

// In memory criterials definition
type GteInmemoryCriteria struct {
	Field string
	Value any
}

func (c GteInmemoryCriteria) Filter() interface{} {
	return "gte"
}

var _ criteria.Criteria = (*GteInmemoryCriteria)(nil)

type InInmemoryCriteria struct {
	Field string
	Value any
}

func (c InInmemoryCriteria) Filter() interface{} {
	return "in"
}

var _ criteria.Criteria = (*InInmemoryCriteria)(nil)

type LikeInmemoryCriteria struct {
	Field string
	Value string
}

func (c LikeInmemoryCriteria) Filter() interface{} {
	return "like"
}

var _ criteria.Criteria = (*LikeInmemoryCriteria)(nil)

type LtInmemoryCriteria struct {
	Field string
	Value any
}

func (c LtInmemoryCriteria) Filter() interface{} {
	return "lt"
}

var _ criteria.Criteria = (*LtInmemoryCriteria)(nil)

type LteInmemoryCriteria struct {
	Field string
	Value any
}

func (c LteInmemoryCriteria) Filter() interface{} {
	return "lte"
}

var _ criteria.Criteria = (*LteInmemoryCriteria)(nil)

type NotInmemoryCriteria struct {
	Field    string
	Criteria criteria.Criteria
}

func (c NotInmemoryCriteria) Filter() interface{} {
	return "not"
}

var _ criteria.Criteria = (*NotInmemoryCriteria)(nil)

type OrInmemoryCriteria struct {
	C1 criteria.Criteria
	C2 criteria.Criteria
}

func (c OrInmemoryCriteria) Filter() interface{} {
	return "or"
}

var _ criteria.Criteria = (*OrInmemoryCriteria)(nil)

type PaginatorInmemoryCriteria struct {
}

func (c PaginatorInmemoryCriteria) Filter() interface{} {
	return "paginator"
}

var _ criteria.Criteria = (*PaginatorInmemoryCriteria)(nil)

type BetweenInmemoryCriteria struct {
	Field string
	V1    any
	V2    any
}

func (c BetweenInmemoryCriteria) Filter() interface{} {
	return "between"
}

var _ criteria.Criteria = (*BetweenInmemoryCriteria)(nil)

type SortInmemoryCriteria struct {
	Field string
	Order string
}

func (c SortInmemoryCriteria) By() string {
	return c.Field
}

func (c SortInmemoryCriteria) Sort() string {
	return c.Order
}

var _ criteria.SortCriteria = (*SortInmemoryCriteria)(nil)
