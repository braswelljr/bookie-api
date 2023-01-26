package users

import (
	"context"

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
func Update(ctx context.Context, id string) (*store.User, error) {
	return &store.User{
		ID: id,
	}, nil
}
