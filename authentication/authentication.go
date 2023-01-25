package authentication

import (
	"context"
	"errors"
	"strings"

	"encore.dev/beta/auth"
	"github.com/go-playground/validator/v10"

	"encore.app/authentication/store"
	"encore.app/pkg/middleware"
)

// Signup is a function that handles the signup process.
//
//	@route POST /signup
//	@param ctx - context.Context
//	@param payload
//	@return response
//	@return error
//
// encore:api public method=POST path=/signup
func Signup(ctx context.Context, payload *store.SignupPayload) (*store.Response, error) {
	// validate user details
	if err := validator.New().Struct(payload); err != nil {
		return &store.Response{}, err
	}

	// create a new user
	user, err := store.Create(ctx, payload)
	if err != nil {
		return nil, err
	}

	// generate tokens
	token, err := middleware.GetToken(&middleware.User{
		ID:          user.ID,
		Firstname:   user.Firstname,
		Lastname:    user.Lastname,
		Othernames:  user.Othernames,
		Username:    user.Username,
		Email:       user.Email,
		DateOfBirth: user.DateOfBirth,
		Phone:       user.Phone,
		Address:     user.Address,
		City:        user.City,
		Country:     user.Country,
		Role:        user.Role,
	})
	if err != nil {
		return &store.Response{}, errors.New("authentication failed: unable to generate token")
	}

	// set the user id and token in the context
	uid := middleware.Store.SetCtxValue("uid", user.ID)
	tid := middleware.Store.SetCtxValue("token", token)
	role := middleware.Store.SetCtxValue("role", user.Role)

	// check if the values were set
	if uid == "" || tid == "" || role == "" {
		return &store.Response{}, errors.New("authentication failed: unable to set context values")
	}

	// return the response
	return &store.Response{
		Message: "Signup successful",
		Token:   token,
		Payload: &store.UserResponse{
			ID:          user.ID,
			Firstname:   user.Firstname,
			Lastname:    user.Lastname,
			Othernames:  user.Othernames,
			Username:    user.Username,
			Email:       user.Email,
			DateOfBirth: user.DateOfBirth,
			Phone:       user.Phone,
			Address:     user.Address,
			City:        user.City,
			Country:     user.Country,
			Role:        user.Role,
			CreatedAt:   user.CreatedAt,
			UpdatedAt:   user.UpdatedAt,
		},
	}, nil
}

// Login - Login is a function that handles the login process for a user.
//
//	@route POST /login
//	@param ctx - context.Context
//	@param payload
//	@return response
//	@return error
//
// encore:api public method=POST path=/login
func Login(ctx context.Context, payload *store.LoginPayload) (*store.Response, error) {
	// validate user details
	if err := validator.New().Struct(payload); err != nil {
		return &store.Response{}, err
	}

	// get the user
	user, err := store.Get(ctx, payload.Email)
	if err != nil {
		return nil, err
	}

	// check if the password is correct
	isCorrect, err := middleware.ComparePasswords(user.Password, payload.Password)
	if err != nil || !isCorrect {
		return nil, errors.New("authentication failed: incorrect password")
	}

	// generate tokens
	token, err := middleware.GetToken(&middleware.User{
		ID:          user.ID,
		Firstname:   user.Firstname,
		Lastname:    user.Lastname,
		Othernames:  user.Othernames,
		Username:    user.Username,
		Email:       user.Email,
		DateOfBirth: user.DateOfBirth,
		Phone:       user.Phone,
		Address:     user.Address,
		City:        user.City,
		Country:     user.Country,
		Role:        user.Role,
	})
	if err != nil {
		return &store.Response{}, errors.New("authentication failed: unable to generate token")
	}

	// set the user id and token in the context
	uid := middleware.Store.SetCtxValue("uid", user.ID)
	tid := middleware.Store.SetCtxValue("token", token)
	role := middleware.Store.SetCtxValue("role", user.Role)

	// check if the values were set
	if uid == "" || tid == "" || role == "" {
		return &store.Response{}, errors.New("authentication failed: unable to set context values")
	}

	return &store.Response{
		Message: "Login successful",
		Token:   token,
		Payload: &store.UserResponse{
			ID:          user.ID,
			Firstname:   user.Firstname,
			Lastname:    user.Lastname,
			Othernames:  user.Othernames,
			Username:    user.Username,
			Email:       user.Email,
			DateOfBirth: user.DateOfBirth,
			Phone:       user.Phone,
			Address:     user.Address,
			City:        user.City,
			Country:     user.Country,
			Role:        user.Role,
			CreatedAt:   user.CreatedAt,
			UpdatedAt:   user.UpdatedAt,
		},
	}, nil
}

// Logout - Logout is a function that handles the logout process for a user.
//
//	@route POST /logout
//	@param ctx - context.Context
//	@return response
//	@return error
//
// encore:api public method=POST path=/logout
func Logout(_ context.Context) (*store.Response, error) {
	// get the user id, token, role from the context
	uid := middleware.Store.SetCtxValue("uid", "")
	token := middleware.Store.SetCtxValue("token", "")
	role := middleware.Store.SetCtxValue("role", "")

	// check if the values are empty
	if uid != "" || token != "" || role != "" {
		return &store.Response{}, errors.New("authentication failed: unable to remove context values")
	}

	return &store.Response{
		Message: "Logout successful",
		Payload: nil,
		Token:   "",
	}, nil
}

// Auth - Auth is a function that handles the authentication process for a user.
//
//		@route POST /auth
//		@param ctx - context.Context
//	 @param token - string
//		@return response
//		@return error
//
// encore:authhandler
func Auth(_ context.Context, token string) (auth.UID, *store.UserResponse, error) {
	// check for empty token
	if token == "" || len(strings.TrimSpace(token)) < 1 {
		return "", &store.UserResponse{}, errors.New("authentication failed: token is empty")
	}

	// validate token
	claims, err := middleware.ValidateToken(token)
	if err != nil {
		return "", &store.UserResponse{}, errors.New("authentication failed: invalid token")
	}

	// get the user
	user, err := store.GetWithID(context.Background(), claims.User.ID)
	if err != nil {
		return "", &store.UserResponse{}, errors.New("authentication failed: unable to get user")
	}

	return auth.UID(claims.User.ID), &store.UserResponse{
		ID:          user.ID,
		Firstname:   user.Firstname,
		Lastname:    user.Lastname,
		Othernames:  user.Othernames,
		Username:    user.Username,
		Email:       user.Email,
		DateOfBirth: user.DateOfBirth,
		Phone:       user.Phone,
		Address:     user.Address,
		City:        user.City,
		Country:     user.Country,
		Role:        user.Role,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}, nil
}
