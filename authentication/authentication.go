package authentication

import (
	"context"
	"errors"

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

	// return the response
	return &store.Response{
		Message: "Signup successful",
		Token:   token,
		Payload: user,
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

	return &store.Response{
		Message: "Login successful",
		Token:   token,
		Payload: user,
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
func Logout(ctx context.Context) (*store.Response, error) {
	return &store.Response{}, nil
}
