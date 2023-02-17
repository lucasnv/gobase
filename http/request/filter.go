package request

// This request works as a generic struct when we have to filter a list
// Examples of valid URIs:
// - /users?page=1&per_page=20
// - /users?sort_by=name&sort_order=asc&page=1&per_page=20
// - /users?filter=name:eq:lucas vazquez&page=1&per_page=20
// - /users?filter=name:eq:lucas vazquez&sort_by=name&sort_order=asc
// - /users?filter=name:eq:lucas vazquez,age:lte:50&sort_by=name&sort_order=asc&page=1&per_page=20
// --------------------
// The parameter filter has the folling struct
// filter=[field]:[operator]:[value],[field]:[operator]:[value]
// --------------------
// The valid operators:
// e = equal
// gt = gratter than
// gte = gratter and equal than
// lt = less than
// lte = less and equal than
// between = between two values
// in = a value in a list
// not-in = as in but not in a list

type FilterRequest struct {
	Filter    string `form:"filter"`
	SortBy    string `form:"sort_by"`
	SortOrder string `form:"sort_order" binding:"eq=asc|eq=desc"`
	Page      int    `form:"page" binding:"numeric|gte=1"`
	PerPage   int    `form:"per_page" binding:"numeric|gte=1|lte=50"`
}
