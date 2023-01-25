package middleware

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var (
	Data  *DataI
	UData *User
)

var secrets struct {
	PrivateKey string
}

// ValidateToken - ValidateToken is a function that handles the verification of tokens.
//
//	@param ctx - context.Context
//	@param token - string
//	@param next - func(context.Context) (interface{}, error)
//	@return response
//	@return error
func ValidateToken(token string) (*SignedParams, error) {
	// get the private key from the secrets
	privateKey := secrets.PrivateKey

	// parse the token
	parsedToken, err := jwt.ParseWithClaims(token, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(privateKey), nil
	})
	if err != nil {
		return nil, err
	}

	// get the claims
	claims, ok := parsedToken.Claims.(*jwt.RegisteredClaims)

	// check if the claims are valid
	if !ok || !parsedToken.Valid {
		return nil, err
	}

	// return the claims
	return &SignedParams{User: UData, RegisteredClaims: *claims}, nil
}

// GetTokens - is a function that handles the retrieval of tokens.
//
//	@param user - *store.User
//	@return string
//	@return string
//	@return error
func GetToken(user *User) (string, error) {
	// get the private key from the secrets
	privateKey := secrets.PrivateKey

	// create a new token
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, &SignedParams{
		User: user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}).SignedString([]byte(privateKey))
	if err != nil {
		return "", err
	}

	// return the token and refresh token
	return token, nil
}
