package store

import (
	"time"
)

type Category struct {
	ID          string    `json:"id" db:"id"`
	UID         string    `json:"uid" db:"uid"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"createdAt" db:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt" db:"updatedAt"`
}

type CreateCategoryPayload struct {
	UID         string `json:"uid" db:"uid" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"omitempty"`
}

type UpdateCategoryPayload struct {
	Name        string `json:"name" db:"name" validate:"omitempty"`
	Description string `json:"description" db:"description" validate:"omitempty"`
}
