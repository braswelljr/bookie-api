package ts

import "time"

type Task struct {
	ID             string    `json:"id" db:"id"`
	UserID         string    `json:"uid" db:"uid"`
	Title          string    `json:"title" db:"title"`
	Description    string    `json:"description" db:"description"`
	Status         string    `json:"status" db:"status"`     // pending, completed, archived
	Category       string    `json:"category" db:"category"` // default: "general", "work", "personal", "shopping", "others"
	Pinned         bool      `json:"pinned" db:"pinned"`
	PinnedAt       time.Time `json:"pinnedAt" db:"pinned_at"`
	PinnedPosition int       `json:"pinnedPosition" db:"pinned_position"` // default -1 -> not pinned
	Archived       bool      `json:"archived" db:"archived"`
	ArchivedAt     time.Time `json:"archivedAt" db:"archived_at"`
	Completed      bool      `json:"completed" db:"completed"` // default: false
	CompletedAt    time.Time `json:"completedAt" db:"completed_at"`
	Color          string    `json:"color" db:"color"` // default: "default", "red", "orange", "yellow", "green", "blue", "purple", "pink", "brown", "grey"
	CreatedAt      time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt      time.Time `json:"updatedAt" db:"updated_at"`
}

type Pin struct {
	Pinned   bool `json:"pinned" db:"pinned"`
	Position int  `json:"position" db:"position"`
}

type Attributes struct {
	Pin      *Pin      `json:"pin" db:"pin"`
	Color    string    `json:"color" db:"color"`
	Category *Category `json:"category" db:"category"`
	Archived bool      `json:"archived" db:"archived"`
}

type Category struct {
	ID          string `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
}

type CreateTaskPayload struct {
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description" validate:"omitempty"`             // optional
	Status      string `json:"status" db:"status" validate:"omitempty" default:"pending"`     // pending, completed, archived
	Category    string `json:"category" db:"category" validate:"omitempty" default:"general"` // default: "general", "work", "personal", "shopping", "others"
}

type UpdateTaskPayload struct {
	Title       string `json:"title" db:"title" validate:"omitempty"`                         // optional
	Description string `json:"description" db:"description" validate:"omitempty"`             // optional
	Status      string `json:"status" db:"status" validate:"omitempty" default:"pending"`     // pending, completed, archived
	Category    string `json:"category" db:"category" validate:"omitempty" default:"general"` // default: "general", "work", "personal", "shopping", "others"
}

type PaginatedTasksResponse struct {
	Tasks       []Task `json:"data"`
	Total       int    `json:"total" db:"total"`
	TotalPages  int    `json:"totalPages" db:"total_pages"`
	CurrentPage int    `json:"currentPage" db:"current_page"`
}

type MultiIdsPayload struct {
	Ids []string `json:"ids" db:"ids"`
}
