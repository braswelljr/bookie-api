package middleware

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var (
	Data     *DataI
	UserData *User
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
	parsedToken, err := jwt.ParseWithClaims(token, &SignedParams{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(privateKey), nil
	})
	if err != nil {
		return nil, errors.New("authentication failed: invalid token")
	}

	// get the claims
	claims, ok := parsedToken.Claims.(*SignedParams)

	// check if the claims are valid
	if !ok || !parsedToken.Valid {
		return nil, errors.New("authentication failed: invalid token")
	}

	// Ensure token is valid not expired
	if !claims.VerifyExpiresAt(time.Now().Local(), true) {
		return nil, errors.New("authentication failed: token has expired")
	}

	// set the role in the context
	role := Store.SetCtxValue("role", claims.User.Role)
	if role != claims.User.Role {
		return nil, errors.New("authentication failed: unable to set role in context")
	}

	// return the claims
	return claims, nil
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
			Issuer:    ContextKey.(string),
			Subject:   user.ID,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}).SignedString([]byte(privateKey))
	if err != nil {
		return "", err
	}

	// return the token and refresh token
	return token, nil
}

// GetUserID - GetUserID is a function that handles the retrieval of user id from the context.
//
//	@param ctx - context.Context
//	@return string
func GetUserID() string {
	// get the user id from the context
	userID := Store.GetCtxValue("userID")

	// return the user id
	return (userID).(string)
}

// IsAdmin - IsAdmin is a function that handles the verification of admin privileges.
//
//	@param ctx - context.Context
//	@return bool
func IsAdmin() bool {
	// get the role from the context
	role := Store.GetCtxValue("role")

	// check if the role is admin
	return role == "admin"
}

// IsSuperAdmin - IsSuperAdmin is a function that handles the verification of superadmin privileges.
//
//	@param ctx - context.Context
//	@return bool
func IsSuperAdmin() bool {
	// get the role from the context
	role := Store.GetCtxValue("role")

	// check if the role is superadmin
	return role == "superadmin"
}
