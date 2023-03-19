package collection

type Collection struct {
	dataValue     any
	metadataValue Metadata
}

type Metadata struct {
	Page       uint32
	PageSize   uint32
	TotalPages uint32
	Total      uint32
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

func NewCollection(data any, page uint32, pageSize uint32, total uint32) Collection {

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
