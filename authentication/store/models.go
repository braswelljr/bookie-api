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
	DateOfBirth string    `json:"dateOfBirth" db:"date_of_birth"`
	Password    string    `json:"password" db:"password"`
	Phone       string    `json:"phone" db:"phone"`
	Address     string    `json:"address" db:"address"`
	City        string    `json:"city" db:"city"`
	Country     string    `json:"country" db:"country"`
	Role        string    `json:"role" db:"role"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time `json:"updatedAt" db:"updated_at"`
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

type UserResponse struct {
	ID          string    `json:"id" db:"id"`
	Firstname   string    `json:"firstname" db:"firstname"`
	Lastname    string    `json:"lastname" db:"lastname"`
	Othernames  string    `json:"othernames" db:"othernames"`
	Username    string    `json:"username" db:"username"`
	DateOfBirth string    `json:"dateOfBirth" db:"date_of_birth"`
	Email       string    `json:"email" db:"email"`
	Phone       string    `json:"phone" db:"phone"`
	Address     string    `json:"address" db:"address"`
	City        string    `json:"city" db:"city"`
	Country     string    `json:"country" db:"country"`
	Role        []string  `json:"role" db:"role"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time `json:"updatedAt" db:"updated_at"`
}

type LoginPayload struct {
	Email    string `json:"email" validate:"required,email"` // required
	Password string `json:"password" validate:"required"`    // required
}

type Response struct {
	Message string `json:"message"`
	Token   string `json:"token"`
	Payload *User  `json:"payload"`
}
