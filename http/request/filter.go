package request

// This request works as a generic struct when we have to filter a list
// Examples of valid URIs:
// - /users?page=1&per_page=20
// - /users?sort_by=name&sort_order=asc&page=1&per_page=20
// - /users?filter=name::eq::lucas vazquez&page=1&per_page=20
// - /users?filter=name::eq::lucas vazquez&sort_by=name&sort_order=asc
// - /users?filter=date::between::2023-10-10|2023-10-12
// - /users?filter=status::in::active|suspended
// - /users?filter=name::eq::lucas vazquez,age::lte::50&sort_by=name&sort_order=asc&page=1&per_page=20
// --------------------
// The parameter filter has the folling struct
// [criteria]::[operator]::[parameters],[criteria]::[operator]::[parameters]

type FilterRequest struct {
	Filter    string `form:"filter" binding:"omitempty"`
	SortBy    string `form:"sort_by" binding:"omitempty,alpha,required_with=sort_order"`
	SortOrder string `form:"sort_order" binding:"omitempty,required_with=sort_by,oneof=asc desc"`
	Page      int    `form:"page" binding:"omitempty,numeric,gte=1"`
	PerPage   int    `form:"per_page" binding:"omitempty,numeric,gte=1,lte=50"`
}
