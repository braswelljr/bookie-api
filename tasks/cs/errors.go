package cs

import "errors"

var (
	ErrNotFound      = errors.New("category not found")
	ErrAlreadyExists = errors.New("category already exists")
)
