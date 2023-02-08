package pagination

// PaginatedResults - paginated results
type PaginatedResults struct {
	Data  []interface{} `json:"data" db:"data"`
	Page  int           `json:"page" db:"page"`
	Limit int           `json:"limit" db:"limit"`
	Total int           `json:"total" db:"total"`
}

// Query - query
type Query struct {
	Where  string `json:"where" db:"where"`
	Tasks  string `json:"tasks" db:"tasks"`
	Fields string `json:"fields" db:"fields"` // the fields to be selected
}

// pagination options
type Options struct {
	Limit int `json:"limit" db:"limit" url:"limit"` // the number of items
	Page  int `json:"page" db:"page" url:"limit"`   // the page
}

// PaginationResponse - pagination
type PaginationResponse struct {
	Offset      int `json:"offset" db:"offset"`  // where to start from
	Limit       int `json:"limit" db:"limit"`    // the number of items
	Total       int `json:"total" total:"total"` // total number of items
	CurrentPage int `json:"page" db:"page"`      // the current page
}

// Pagination represents a pagination.
type Pagination struct {
	page     int
	perPage  int
	total    int
	pages    int
	previous int
	next     int
	offset   int
}
