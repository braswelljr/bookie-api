package store

import "time"

type Label struct {
	ID        string    `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Color     string    `json:"color" db:"color"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}

type LabelRequest struct {
	Name  string `json:"name" db:"name" validate:"required"`
	Color string `json:"color" db:"color" validate:"omitempty,hexcolor"`
}
