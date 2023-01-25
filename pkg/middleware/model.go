package middleware

import (
	"github.com/golang-jwt/jwt/v4"
)

type DataI struct {
	User  *User
	Roles []string
}

type User struct {
	Firstname   string
	Lastname    string
	Othernames  string
	Username    string
	Email       string
	DateOfBirth string
	Phone       string
	Address     string
	City        string
	Country     string
	Role        string
}

type SignedParams struct {
	User *User
	jwt.RegisteredClaims
}
