package store

import "time"

type Task struct {
	ID          string      `json:"id" db:"id"`
	UserID      string      `json:"uid" db:"uid"`
	Title       string      `json:"title" db:"title"`
	Description string      `json:"description" db:"description"`
	Completed   bool        `json:"completed" db:"completed"`
	Tags        []string    `json:"tags" db:"tags"`
	Attributes  *Attributes `json:"attributes" db:"attributes"`
	CreatedAt   time.Time   `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at" db:"updated_at"`
}

type Pin struct {
	Pinned   bool `json:"pinned" db:"pinned"`
	Position int  `json:"position" db:"position"`
}

type Attributes struct {
	Pin      *Pin      `json:"pinned" db:"pinned"`
	Color    string    `json:"color" db:"color"`
	Icon     string    `json:"icon" db:"icon"`
	Category *Category `json:"category" db:"category"`
}

type Category struct {
	ID          string `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
}
