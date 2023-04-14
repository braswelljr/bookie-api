package middleware

import (
	"golang.org/x/crypto/bcrypt"
)

// ComparePasswords - ComparePasswords is a function that handles the comparison of passwords.
//
//	@param hash - string
//	@param password - string
//	@return bool
//	@return error
func ComparePasswords(hash, password string) (bool, error) {
	// compare the passwords
	// check if there is an error and return it
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return false, err
	}

	// return true if there is no error
	return true, nil
}

// HashPassword - HashPassword is a function that handles the hashing of passwords.
//
//	@param password - string
//	@return string
//	@return error
func HashPassword(password string) (string, error) {
	// hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	// check if there is an error and return it
	if err != nil {
		return "", err
	}

	// return the hash if there is no error
	return string(hash), nil
}
