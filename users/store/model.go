package store

import (
	"time"
)

type User struct {
	ID          string    `json:"id" db:"id"`
	Firstname   string    `json:"firstname" db:"firstname"`
	Lastname    string    `json:"lastname" db:"lastname"`
	Othernames  string    `json:"othernames" db:"othernames"`
	Username    string    `json:"username" db:"username"`
	Email       string    `json:"email" db:"email"`
	DateOfBirth string    `json:"dateOfBirth" db:"dateOfBirth"`
	Password    string    `json:"password" db:"password"`
	Phone       string    `json:"phone" db:"phone"`
	Address     string    `json:"address" db:"address"`
	City        string    `json:"city" db:"city"`
	Country     string    `json:"country" db:"country"`
	Role        string    `json:"role" db:"role"`
	CreatedAt   time.Time `json:"createdAt" db:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt" db:"updatedAt"`
}

type SignupPayload struct {
	Firstname   string `json:"firstname" validate:"required"`        // required
	Lastname    string `json:"lastname" validate:"required"`         // required
	Othernames  string `json:"othernames" validate:"omitempty"`      // optional
	Username    string `json:"username" validate:"required"`         // required
	DateOfBirth string `json:"dateOfBirth" validate:"required"`      // required
	Email       string `json:"email" validate:"required,email"`      // required
	Password    string `json:"password" validate:"required" min:"8"` // required
	Phone       string `json:"phone" validate:"required"`            // required
	Address     string `json:"address" validate:"omitempty"`         // optional
	City        string `json:"city" validate:"omitempty"`            // optional
	Country     string `json:"country" validate:"omitempty"`         // optional
}

type UpdatePayload struct {
	Firstname   string `json:"firstname" db:"firstname" validate:"omitempty"`     // not required
	Lastname    string `json:"lastname" db:"lastname" validate:"omitempty"`       // not required
	Othernames  string `json:"othernames" db:"othernames" validate:"omitempty"`   // not required
	Username    string `json:"username" db:"username" validate:"omitempty"`       // not required
	DateOfBirth string `json:"dateOfBirth" db:"dateOfBirth" validate:"omitempty"` // not required
	Email       string `json:"email" db:"email" validate:"omitempty,email"`       // not required
	Phone       string `json:"phone" db:"phone" validate:"omitempty"`             // not required
	Address     string `json:"address" db:"address" validate:"omitempty"`         // not required
	City        string `json:"city" db:"city" validate:"omitempty"`               // not required
	Country     string `json:"country" db:"country" validate:"omitempty"`         // not required
}

type UpdateRolePayload struct {
	ID string `json:"id" db:"id" validate:"required"` // required
}

type UserResponse struct {
	ID          string    `json:"id" db:"id"`
	Firstname   string    `json:"firstname" db:"firstname"`
	Lastname    string    `json:"lastname" db:"lastname"`
	Othernames  string    `json:"othernames" db:"othernames"`
	Username    string    `json:"username" db:"username"`
	DateOfBirth string    `json:"dateOfBirth" db:"dateOfBirth"`
	Email       string    `json:"email" db:"email"`
	Phone       string    `json:"phone" db:"phone"`
	Address     string    `json:"address" db:"address"`
	City        string    `json:"city" db:"city"`
	Country     string    `json:"country" db:"country"`
	Role        string    `json:"role" db:"role"`
	CreatedAt   time.Time `json:"createdAt" db:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt" db:"updatedAt"`
}

type PaginatedUsersResponse struct {
	Users       []UserResponse `json:"data"`
	Total       int            `json:"total" db:"total"`
	TotalPages  int            `json:"totalPages" db:"totalPages"`
	CurrentPage int            `json:"currentPage" db:"currentPage"`
}

type UserUpdateResponse struct {
	Message string `json:"message"`
}

type DeleteUserEvent struct {
	ID string `json:"id" db:"id"`
}
