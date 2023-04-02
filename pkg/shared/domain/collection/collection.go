package collection

type Collection struct {
	items    []any
	metadata Metadata
}

type Metadata struct {
	Page       uint32
	PageSize   uint32
	TotalPages uint32
	Total      uint32
}

func (c Collection) Data() []any {
	return c.items
}

func (c Collection) Metadata() Metadata {
	return c.metadata
}

func (c *Collection) SetMetadata(page uint32, pageSize uint32, total uint32) Collection {
	totalPages := total / pageSize

	if total%pageSize != 0 {
		totalPages = totalPages + 1
	}

	c.metadata = Metadata{
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
		Total:      total,
	}

	return *c
}

func (c *Collection) Add(item any) {
	c.items = append(c.items, item)
}

func (c *Collection) Remove(item interface{}) {
	for i, v := range c.items {
		if v == item {
			c.items = append(c.items[:i], c.items[i+1:]...)
			break
		}
	}
}

func (c Collection) Transform(fn func(any) any) Collection {
	result := Collection{
		metadata: c.Metadata(),
	}

	for _, item := range c.items {
		result.Add(fn(item))
	}

	return result
}

func (c Collection) First() any {
	return c.items[0]
}

func New() Collection {
	return Collection{}
}

func calculateTotalPages(totalItems int, pageSize int) int {
	totalPages := totalItems / pageSize

	if totalItems%pageSize != 0 {
		totalPages++
	}

	return totalPages
}
