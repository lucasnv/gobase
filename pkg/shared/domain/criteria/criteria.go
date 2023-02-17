package criteria

type Offset int

type Limit int

type SortBy struct {
	field string
	order string
}

type Filters struct {
	data interface{}
}

type Criteria struct {
	filters Filters
	order   SortBy
	offset  Offset
	limit   Limit
}

func NewCriteria(f Filters, o SortBy, of Offset, l Limit) *Criteria {
	return &Criteria{
		filters: f,
		order:   o,
		offset:  of,
		limit:   l,
	}
}

func NewSortBy(f string, o string) SortBy {
	return SortBy{
		field: f,
		order: o,
	}
}

func NewFilters(data interface{}) Filters {
	return Filters{
		data: data,
	}
}
