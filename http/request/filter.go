package request

// This request works as a generic struct when we have to filter a list
type FilterRequest struct {
	Filter    string `form:"filter" binding:"omitempty"`
	SortBy    string `form:"sort_by" binding:"omitempty,required_with=sort_order"`
	SortOrder string `form:"sort_order" binding:"omitempty,required_with=sort_by,oneof=asc desc"`
	Page      int    `form:"page" binding:"omitempty,numeric,gte=1"`
	PerPage   int    `form:"per_page" binding:"omitempty,numeric,gte=1,lte=50"`
}
