package store

import (
	"time"
)

type UserResponse struct {
	ID          string    `json:"id" db:"id"`
	Firstname   string    `json:"firstname" db:"firstname"`
	Lastname    string    `json:"lastname" db:"lastname"`
	Othernames  string    `json:"othernames" db:"othernames"`
	Username    string    `json:"username" db:"username"`
	DateOfBirth string    `json:"dateOfBirth" db:"date_of_birth"`
	Email       string    `json:"email" db:"email"`
	Phone       string    `json:"phone" db:"phone"`
	Role        string    `json:"role" db:"role"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time `json:"updatedAt" db:"updated_at"`
}

type LoginPayload struct {
	Email    string `json:"email" validate:"required,email"` // required
	Password string `json:"password" validate:"required"`    // required
}

type Response struct {
	Message string        `json:"message"`
	Code    int           `json:"code"`
	Token   string        `json:"token"`
	Payload *UserResponse `json:"payload"`
}
