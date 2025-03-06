package pagination

// Pagination represents pagination information
type Pagination struct {
	CurrentPage  int `json:"current_page"`
	TotalPages   int `json:"total_pages"`
	TotalItems   int `json:"total_items"`
	ItemsPerPage int `json:"items_per_page"`
}

// PaginatedResponse represents a generic paginated response
type PaginatedResponse[T any] struct {
	Items      []T         `json:"items"`
	Pagination Pagination `json:"pagination"`
}

// NewPaginatedResponse creates a new paginated response
func NewPaginatedResponse[T any](items []T, currentPage, totalItems, itemsPerPage int) PaginatedResponse[T] {
	totalPages := (totalItems + itemsPerPage - 1) / itemsPerPage
	if totalPages < 1 {
		totalPages = 1
	}

	return PaginatedResponse[T]{
		Items: items,
		Pagination: Pagination{
			CurrentPage:  currentPage,
			TotalPages:   totalPages,
			TotalItems:   totalItems,
			ItemsPerPage: itemsPerPage,
		},
	}
}

// GetOffset calculates the offset for SQL queries
func GetOffset(page, itemsPerPage int) int {
	if page < 1 {
		page = 1
	}
	return (page - 1) * itemsPerPage
} 