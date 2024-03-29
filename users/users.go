package users

import (
	"context"
	"fmt"

	"github.com/go-playground/validator/v10"

	"encore.app/pkg/events"
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

// QueryAll - Get all users
//
//	@param ctx - context.Context
//	@return users
//	@return error
//
// encore:api public method=GET path=/users
func QueryAll(ctx context.Context, options *pagination.Options) (*store.PaginatedUsersResponse, error) {
	// check if user is admin or superadmin
	claims, err := middleware.GetVerifiedClaims(ctx, "")
	if err != nil {
		return &store.PaginatedUsersResponse{}, err
	}

	// check for the roles
	if !claims.HasRole(middleware.RoleSuperAdmin, middleware.RoleAdmin) {
		return &store.PaginatedUsersResponse{}, fmt.Errorf("unauthorized: you are not authorized to perform this action")
	}

	// query users
	users, err := store.GetAll(ctx, options)
	if err != nil {
		return &store.PaginatedUsersResponse{}, fmt.Errorf("querying users: %w", err)
	}

	// return users, nil
	return users, nil
}

// Get - Get a user
//
//	@param ctx - context.Context
//	@param id
//	@return user
//	@return error
//
// encore:api auth method=GET path=/users/:id
func Get(ctx context.Context, id string) (*store.User, error) {
	// check for claims
	claims, err := middleware.GetVerifiedClaims(ctx, "")
	if err != nil {
		return &store.User{}, err
	}

	// check for the roles
	if !claims.HasRole(middleware.RoleSuperAdmin, middleware.RoleAdmin) || claims.Subject.ID != id {
		return &store.User{}, fmt.Errorf("unauthorized: you are not authorized to perform this action")
	}

	// get user
	user, err := store.GetWithID(ctx, id)
	if err != nil {
		return nil, err
	}

	// return user
	return user, nil
}

// Delete - Delete a user
//
//	@param ctx - context.Context
//	@param id
//	@return error
//
// encore:api auth method=DELETE path=/users/:id
func Delete(ctx context.Context, id string) error {
	// check for claims
	claims, err := middleware.GetVerifiedClaims(ctx, "")
	if err != nil {
		return err
	}

	// check for the roles
	if !claims.HasRole(middleware.RoleSuperAdmin) || claims.Subject.ID != id {
		return fmt.Errorf("unauthorized: you are not authorized to perform this action")
	}

	// delete user
	if err := store.Delete(ctx, id); err != nil {
		return err
	}

	// publish delete user's tasks event
	if _, err := events.DeleteAllUserTasks.Publish(ctx, &events.DeleteAllUserTasksEvent{
		UserID: id,
	}); err != nil {
		return err
	}

	// return nil on a delete event
	return nil
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
	// check if the user matches the authenticated user
	claims, err := middleware.GetVerifiedClaims(ctx, "")
	if err != nil {
		return &store.UserUpdateResponse{}, err
	}

	// check if the user is authorized to perform this action
	if claims.Subject.ID != id {
		return &store.UserUpdateResponse{}, fmt.Errorf("unauthorized: you are not authorized to perform this action")
	}

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

// UpdateRole - Updates a user's role
//
//	@param ctx - context.Context
//	@param id
//	@param role
//	@return user
//	@return error
//
// encore:api auth method=PATCH path=/users/update/:id/role-to-admin
func UpdateRole(ctx context.Context, id string) error {
	// check if user is admin or superadmin
	claims, err := middleware.GetVerifiedClaims(ctx, "")
	if err != nil {
		return err
	}

	// check for the roles
	if !claims.HasRole(middleware.RoleSuperAdmin, middleware.RoleAdmin) {
		return fmt.Errorf("unauthorized: you are not authorized to perform this action")
	}

	// update user
	if err := store.UpdateRole(ctx, id, middleware.RoleAdmin); err != nil {
		return err
	}

	// return user
	return nil
}
