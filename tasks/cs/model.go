package cs

import (
	"time"
)

type Category struct {
	ID          string    `json:"id" db:"id"`
	UID         string    `json:"uid" db:"uid"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time `json:"updatedAt" db:"updated_at"`
}

type CreateCategoryPayload struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"omitempty"`
}

type UpdateCategoryPayload struct {
	Name        string `json:"name" db:"name" validate:"omitempty"`
	Description string `json:"description" db:"description" validate:"omitempty"`
}

type PaginatedCategoriesResponse struct {
	Categories  []Category `json:"data"`
	Total       int        `json:"total" db:"total"`
	TotalPages  int        `json:"totalPages" db:"total_pages"`
	CurrentPage int        `json:"currentPage" db:"current_page"`
}

type MultiIdsPayload struct {
	Ids []string `json:"ids" db:"ids"`
}
