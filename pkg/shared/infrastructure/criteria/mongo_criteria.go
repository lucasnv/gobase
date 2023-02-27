package criteria

import (
	"<MODULE_URL_REPLACE>/pkg/shared/domain/criteria"
	"<MODULE_URL_REPLACE>/pkg/shared/domain/errors"
)

type MongoBuilder struct {
}

func (MongoBuilder) GetFilters(f string) (*criteria.CriteriaFilter, *errors.AppError) {
	return criteria.NewCriteriaFilter(f)
}

func (MongoBuilder) And(c1 criteria.Criteria, c2 criteria.Criteria) criteria.Criteria {
	return AndMongoCriteria{}
}

func (MongoBuilder) Owner(field string, value string) criteria.Criteria {
	return OwnerMongoCriteria{
		field: field,
		value: value,
	}
}

func (MongoBuilder) Between(field string, values []string) criteria.Criteria {
	return BetweenMongoCriteria{}
}

func (MongoBuilder) Eq(field string, value string) criteria.Criteria {
	return EqMongoCriteria{
		field: field,
		value: value,
	}
}

func (MongoBuilder) Gt(field string, value any) criteria.Criteria {
	return GtMongoCriteria{}
}

func (MongoBuilder) Gte(field string, value any) criteria.Criteria {
	return GteMongoCriteria{}
}

func (MongoBuilder) In(field string, values []string) criteria.Criteria {
	return InMongoCriteria{}
}

func (MongoBuilder) Like(field string, value string) criteria.Criteria {
	return LikeMongoCriteria{}
}

func (MongoBuilder) Lt(field string, value any) criteria.Criteria {
	return LtMongoCriteria{}
}

func (MongoBuilder) Lte(field string, value any) criteria.Criteria {
	return LteMongoCriteria{}
}

func (MongoBuilder) Not(field string, c criteria.Criteria) criteria.Criteria {
	return NotMongoCriteria{}
}

func (MongoBuilder) Or(c1 criteria.Criteria, c2 criteria.Criteria) criteria.Criteria {
	return OrMongoCriteria{}
}

func (MongoBuilder) Sort(field string, order string) criteria.SortCriteria {
	return SortMongoCriteria{}
}

func (MongoBuilder) Paginator(page int, perPage int) criteria.PaginatorCriteria {
	return PaginatorMongoCriteria{}
}

func NewMongoBuilder() (criteria.Builder, *errors.AppError) {
	return &MongoBuilder{}, nil
}

var _ criteria.Builder = (*MongoBuilder)(nil)

type AndMongoCriteria struct {
}

func (c AndMongoCriteria) Filter() interface{} {
	return true
}

var _ criteria.Criteria = (*AndMongoCriteria)(nil)

type OwnerMongoCriteria struct {
	field string
	value any
}

func (c OwnerMongoCriteria) Filter() interface{} {
	return true
}

var _ criteria.Criteria = (*OwnerMongoCriteria)(nil)

type EqMongoCriteria struct {
	field string
	value any
}

func (c EqMongoCriteria) Filter() interface{} {
	return true
}

var _ criteria.Criteria = (*EqMongoCriteria)(nil)

type GtMongoCriteria struct {
	field string
	value any
}

func (c GtMongoCriteria) Filter() interface{} {
	return true
}

var _ criteria.Criteria = (*GtMongoCriteria)(nil)

type GteMongoCriteria struct {
	field string
	value any
}

func (c GteMongoCriteria) Filter() interface{} {
	return true
}

var _ criteria.Criteria = (*GteMongoCriteria)(nil)

type InMongoCriteria struct {
	field string
	value any
}

func (c InMongoCriteria) Filter() interface{} {
	return true
}

var _ criteria.Criteria = (*InMongoCriteria)(nil)

type LikeMongoCriteria struct {
	field string
	value string
}

func (c LikeMongoCriteria) Filter() interface{} {
	return true
}

var _ criteria.Criteria = (*LikeMongoCriteria)(nil)

type LtMongoCriteria struct {
	field string
	value any
}

func (c LtMongoCriteria) Filter() interface{} {
	return true
}

var _ criteria.Criteria = (*LtMongoCriteria)(nil)

type LteMongoCriteria struct {
	field string
	value any
}

func (c LteMongoCriteria) Filter() interface{} {
	return true
}

var _ criteria.Criteria = (*LteMongoCriteria)(nil)

type NotMongoCriteria struct {
	field string
	value criteria.Criteria
}

func (c NotMongoCriteria) Filter() interface{} {
	return true
}

var _ criteria.Criteria = (*NotMongoCriteria)(nil)

type OrMongoCriteria struct {
}

func (c OrMongoCriteria) Filter() interface{} {
	return true
}

var _ criteria.Criteria = (*OrMongoCriteria)(nil)

type SortMongoCriteria struct {
	Field string
	Order string
}

func (c SortMongoCriteria) By() string {
	return c.Field
}

func (c SortMongoCriteria) Sort() string {
	return c.Order
}

var _ criteria.SortCriteria = (*SortMongoCriteria)(nil)

type PaginatorMongoCriteria struct {
}

func (c PaginatorMongoCriteria) Filter() interface{} {
	return true
}

var _ criteria.Criteria = (*PaginatorMongoCriteria)(nil)

type BetweenMongoCriteria struct {
}

func (c BetweenMongoCriteria) Filter() interface{} {
	return true
}

var _ criteria.Criteria = (*BetweenMongoCriteria)(nil)
