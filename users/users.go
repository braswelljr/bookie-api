package users

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"

	"encore.app/pkg/middleware"
	"encore.app/pkg/pagination"
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
func Update(ctx context.Context, id string, payload store.UpdatePayload) (*store.UserUpdateResponse, error) {
	// validate user details
	if err := validator.New().Struct(payload); err != nil {
		return &store.UserUpdateResponse{}, err
	}

	// update user
	if err := store.Update(ctx, id, payload); err != nil {
		return &store.UserUpdateResponse{}, err
	}

	return &store.UserUpdateResponse{
		Message: fmt.Sprintf("user with id %s updated successfully", id),
	}, nil
}

// Get - Get a user
//
//	@param ctx - context.Context
//	@param id
//	@return user
//	@return error
//
// encore:api public method=GET path=/users/:id tag:cache
func Get(ctx context.Context, id string) (*store.User, error) {
	// get user
	user, err := store.GetWithID(ctx, id)
	if err != nil {
		return nil, err
	}

	// return user
	return user, nil
}

// UpdateRole - Updates a user's role
//
//	@param ctx - context.Context
//	@param id
//	@param role
//	@return user
//	@return error
//
// encore:api auth method=PATCH path=/users/update-role-to-admin
func UpdateRoleAsAdmin(ctx context.Context, payload *store.UpdateRolePayload) error {
	// check if user is admin or superadmin
	if !middleware.IsAdmin() && !middleware.IsSuperAdmin() {
		return errors.New("unauthorized: you are not authorized to perform this action")
	}

	// validate user details
	if err := validator.New().Struct(payload); err != nil {
		return err
	}

	// update user
	if err := store.UpdateRole(ctx, payload, "admin"); err != nil {
		return err
	}

	// return user
	return nil
}

// GetAll - Get all users
//
//	@param ctx - context.Context
//	@return users
//	@return error
//
// encore:api public method=GET path=/users
func QueryAll(ctx context.Context, options *pagination.Options) (*store.PaginatedUsersResponse, error) {
	// query users
	users, err := store.GetAll(ctx, options)
	if err != nil {
		return &store.PaginatedUsersResponse{}, fmt.Errorf("querying users: %w", err)
	}

	// return users, nil
	return users, nil
}

// Delete - Delete a user
//
//	@param ctx - context.Context
//	@param id
//	@return error
//
// encore:api auth method=DELETE path=/users/:id
func Delete(ctx context.Context, id string) error {
	// delete user
	if err := store.Delete(ctx, id); err != nil {
		return err
	}

	// return nil on a delete event
	return nil
}
