package users

import (
	"context"

	"github.com/go-playground/validator/v10"

	"encore.app/users/store"
)

// Create - Create a new user
//
//	@param ctx - context.Context
//	@param payload
//	@return user
//	@return error
//
// encore:api private method=POST path=/users/create
func Create(ctx context.Context, payload *store.SignupPayload) (*store.User, error) {

	// create user
	user, err := store.Create(ctx, payload)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Update - Updates a user
//
//	@param ctx - context.Context
//	@param payload
//	@return user
//	@return error
//
// encore:api auth method=PATCH path=/users/update/:id
func Update(ctx context.Context, id string, payload store.UpdatePayload) (*store.User, error) {
	// validate user details
	if err := validator.New().Struct(payload); err != nil {
		return &store.User{}, err
	}

	// update user
	user, err := store.Update(ctx, id, payload)
	if err != nil {
		return nil, err
	}

	return user, nil
}
