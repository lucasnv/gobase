package collection

type Collection struct {
	dataValue     any
	metadataValue Metadata
}

type Metadata struct {
	Page       int
	PageSize   int
	TotalPages int
	Total      int
}

func (c Collection) Data() any {
	return c.dataValue
}

func (c Collection) Metadata() Metadata {
	return c.metadataValue
}

func (c Collection) SetData(data any) Collection {
	return Collection{
		dataValue:     data,
		metadataValue: c.metadataValue,
	}
}

func NewCollection(data any, page int, pageSize int, total int) Collection {

	totalPages := total / pageSize

	if total%pageSize != 0 {
		totalPages = totalPages + 1
	}

	return Collection{
		dataValue: data,
		metadataValue: Metadata{
			Page:       page,
			PageSize:   pageSize,
			TotalPages: totalPages,
			Total:      total,
		},
	}
}

func calculateTotalPages(totalItems int, pageSize int) int {
	totalPages := totalItems / pageSize

	if totalItems%pageSize != 0 {
		totalPages++
	}

	return totalPages
}
