package middleware

import (
	"github.com/golang-jwt/jwt/v4"
)

type DataI struct {
	User  *User
	Roles []string
}

type User struct {
	ID          string
	Firstname   string
	Lastname    string
	Othernames  string
	Username    string
	Email       string
	DateOfBirth string
	Phone       string
	Role        string
}

type SignedParams struct {
	User *User
	jwt.RegisteredClaims
}

// Store Structs
type CtxKey interface{}
type CtxValues struct {
	m map[string]interface{}
}
