package tasks

import (
	"context"
	"fmt"

	"encore.dev/pubsub"
	"github.com/go-playground/validator/v10"

	"encore.app/pkg/events"
	"encore.app/pkg/pagination"
	"encore.app/tasks/cs"
	"encore.app/tasks/ts"
	"encore.app/users"
)

// Create - Create is a function that creates a new task.
//
// @param ctx - context.Context
// @param payload
// @return task
// @return error
//
// encore:api auth method=POST path=/tasks/:uid/create
func CreateTask(ctx context.Context, uid string, payload *ts.CreateTaskPayload) error {
	// validate payload
	if err := validator.New().Struct(payload); err != nil {
		return err
	}

	// check if user exists
	if user, err := users.Get(ctx, uid); err != nil || user == nil || user.ID != uid {
		return err
	}

	// create task
	if err := ts.Create(ctx, uid, payload); err != nil {
		return err
	}

	return nil
}

// Get - Get a task
//
// @param ctx - context.Context
// @param id - string
// @return task
// @return error
//
// encore:api public method=GET path=/tasks/:id
func GetTask(ctx context.Context, id string) (*ts.Task, error) {
	// get task
	task, err := ts.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return task, nil
}

// Update - Update a task
//
// @param ctx - context.Context
// @param id - string
// @param payload
// @return task
// @return error
//
// encore:api auth method=PATCH path=/tasks/:id
func UpdateTask(ctx context.Context, id string, payload *ts.UpdateTaskPayload) error {
	// validate payload
	if err := validator.New().Struct(payload); err != nil {
		return err
	}

	// update task
	if err := ts.Update(ctx, id, payload); err != nil {
		return err
	}

	// return nil if no error
	return nil
}

// Delete - Delete a task
//
// @param ctx - context.Context
// @param id - string
// @return task
// @return error
//
// encore:api auth method=DELETE path=/tasks/:id
func DeleteTask(ctx context.Context, id string) error {
	// delete task
	if err := ts.Delete(ctx, id); err != nil {
		return err
	}

	// return nil if no error
	return nil
}

// ToggleComplete - Toggle a task's complete status
//
// @param ctx - context.Context
// @param id - string
// @return task
// @return error
//
// encore:api auth method=PATCH path=/tasks/:id/toggle-complete
func ToggleTaskComplete(ctx context.Context, id string) error {
	// toggle complete
	if err := ts.ToggleComplete(ctx, id); err != nil {
		return err
	}

	// return nil if no error
	return nil
}

// ToggleMultipleComplete - Toggle multiple tasks' complete status
//
// @param ctx - context.Context
// @param {*ts.ToggleMultipleTasksCompletePayload} ids - ids of tasks to toggle
// @return error
//
// encore:api auth method=PATCH path=/tasks/toggle-complete/multiple
func ToggleMultipleTaskComplete(ctx context.Context, ids *ts.MultiIdsPayload) error {
	// toggle complete
	if err := ts.ToggleMultipleComplete(ctx, ids.Ids); err != nil {
		return err
	}

	// return nil if no error
	return nil
}

// GetUserTasks - Get all tasks for a user
//
// @param ctx - context.Context
// @param uid - string
// @return tasks
// @return error
//
// encore:api auth method=GET path=/users/:uid/tasks
func GetUserTasks(ctx context.Context, uid string, options *pagination.Options) (*ts.PaginatedTasksResponse, error) {
	// get user tasks
	tasks, err := ts.GetUserTasks(ctx, uid, options)
	if err != nil {
		return nil, fmt.Errorf("querying tasks: %w", err)
	}

	// return tasks and nil if no error
	return tasks, nil
}

// DeleteAllWithUserID - Delete all tasks for a user
//
// @param ctx - context.Context
// @param uid - string
// @return task
// @return error
//
// encore:api auth method=DELETE path=/users/:uid/tasks/delete
func DeleteAllTasksWithUserID(ctx context.Context, uid string) error {
	// delete tasks
	if err := ts.DeleteAllWithUserID(ctx, uid); err != nil {
		return err
	}

	// return nil if no error
	return nil
}

// SUBSCRIPTIONS - Subscriptions to delete all tasks for a user
//
// @param ctx - context.Context
// @param uid - string
var _ = pubsub.NewSubscription(
	events.DeleteAllUserTasks,
	"delete-all-tasks-with-user-id",
	pubsub.SubscriptionConfig[*events.DeleteAllUserTasksEvent]{
		Handler: func(ctx context.Context, event *events.DeleteAllUserTasksEvent) error {
			return DeleteAllTasksWithUserID(ctx, event.UserID)
		},
	},
)

// =====================================================================================================================
// CATEGORY
// =====================================================================================================================

// Create - Create a new category
//
//	@param ctx - context.Context
//	@param payload - *CreateCategoryPayload
//
// encore:api auth method=POST path=/categories/:uid/create
func CreateCategory(ctx context.Context, uid string, payload *cs.CreateCategoryPayload) error {
	// validate payload
	if err := validator.New().Struct(payload); err != nil {
		return err
	}

	// create category
	if err := cs.Create(ctx, uid, payload); err != nil {
		return err
	}

	return nil
}

// Get - Get a category
//
//	@param ctx - context.Context
//	@param id - string
//	@return category
//	@return error
//
// encore:api public method=GET path=/categories/:id
func GetCategory(ctx context.Context, id string) (*cs.Category, error) {
	// get category
	category, err := cs.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return category, nil
}

// Update - Update a category
//
//	@param ctx - context.Context
//	@param id - string
//	@param payload
//	@return category
//	@return error
//
// encore:api auth method=PATCH path=/categories/:id
func UpdateCategory(ctx context.Context, id string, payload *cs.UpdateCategoryPayload) error {
	// validate payload
	if err := validator.New().Struct(payload); err != nil {
		return err
	}

	// update category
	if err := cs.Update(ctx, id, payload); err != nil {
		return err
	}

	// return nil if no error
	return nil
}

// Delete - Delete a category
//
//	@param ctx - context.Context
//	@param id - string
//	@return error
//
// encore:api auth method=DELETE path=/categories/:id
func DeleteCategory(ctx context.Context, id string) error {
	// delete category
	if err := cs.Delete(ctx, id); err != nil {
		return err
	}

	// return nil if no error
	return nil
}

// GetUserCategories - Get all categories for a user
//
//	@param ctx - context.Context
//	@param uid - string
//	@return categories
//	@return error
//
// encore:api auth method=GET path=/users/:uid/categories
func GetUserCategories(ctx context.Context, uid string, options *pagination.Options) (*cs.PaginatedCategoriesResponse, error) {
	// get user categories
	categories, err := cs.GetUserCategories(ctx, uid, options)
	if err != nil {
		return nil, fmt.Errorf("querying categories: %w", err)
	}

	// return categories and nil if no error
	return categories, nil
}

// DeleteAllUserCategories - Delete all categories for a user
//
//	@param ctx - context.Context
//	@param uid - string
//	@return error
//
// encore:api auth method=DELETE path=/users/:uid/categories/delete
func DeleteAllUserCategories(ctx context.Context, uid string) error {
	// delete categories
	if err := cs.DeleteAllUserCategories(ctx, uid); err != nil {
		return err
	}

	// return nil if no error
	return nil
}

// =====================================================================================================================
// LABEL
// =====================================================================================================================
