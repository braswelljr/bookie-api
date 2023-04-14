package tasks

import (
	"context"
	"fmt"

	"encore.dev/pubsub"
	"github.com/go-playground/validator/v10"

	"encore.app/pkg/events"
	"encore.app/pkg/pagination"
	"encore.app/tasks/store"
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
func Create(ctx context.Context, uid string, payload *store.CreateTaskPayload) error {
	// validate payload
	if err := validator.New().Struct(payload); err != nil {
		return err
	}

	// check if user exists
	if user, err := users.Get(ctx, uid); err != nil || user == nil || user.ID != uid {
		return err
	}

	// create task
	if err := store.Create(ctx, uid, payload); err != nil {
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
func Get(ctx context.Context, id string) (*store.Task, error) {
	// get task
	task, err := store.Get(ctx, id)
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
// encore:api auth method=PATCH path=/tasks/:id/update
func Update(ctx context.Context, id string, payload *store.UpdateTaskPayload) error {
	// validate payload
	if err := validator.New().Struct(payload); err != nil {
		return err
	}

	// update task
	if err := store.Update(ctx, id, payload); err != nil {
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
// encore:api auth method=DELETE path=/tasks/:id/delete
func Delete(ctx context.Context, id string) error {
	// delete task
	if err := store.Delete(ctx, id); err != nil {
		return err
	}

	// return nil if no error
	return nil
}

// DeleteAllWithUserID - Delete all tasks for a user
//
// @param ctx - context.Context
// @param uid - string
// @return task
// @return error
//
// encore:api auth method=DELETE path=/users/:uid/tasks/delete
func DeleteAllWithUserID(ctx context.Context, uid string) error {
	// delete tasks
	if err := store.DeleteAllWithUserID(ctx, uid); err != nil {
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
			return DeleteAllWithUserID(ctx, event.UserID)
		},
	},
)

// GetUserTasks - Get all tasks for a user
//
// @param ctx - context.Context
// @param uid - string
// @return tasks
// @return error
//
// encore:api auth method=GET path=/users/:uid/tasks
func GetUserTasks(ctx context.Context, uid string, options *pagination.Options) (*store.PaginatedTasksResponse, error) {
	// get user tasks
	tasks, err := store.GetUserTasks(ctx, uid, options)
	if err != nil {
		return nil, fmt.Errorf("querying tasks: %w", err)
	}

	// return tasks and nil if no error
	return tasks, nil
}

// ToggleComplete - Toggle a task's complete status
//
// @param ctx - context.Context
// @param id - string
// @return task
// @return error
//
// encore:api auth method=PATCH path=/tasks/toggle-complete/:id
func ToggleComplete(ctx context.Context, id string) error {
	// toggle complete
	if err := store.ToggleComplete(ctx, id); err != nil {
		return err
	}

	// return nil if no error
	return nil
}

// ToggleMultipleComplete - Toggle multiple tasks' complete status
//
// @param ctx - context.Context
// @param {*store.ToggleMultipleTasksCompletePayload} ids - ids of tasks to toggle
// @return error
//
// encore:api auth method=PATCH path=/tasks/toggle-multiple-complete
func ToggleMultipleComplete(ctx context.Context, ids *store.MultiIdsPayload) error {
	// toggle complete
	if err := store.ToggleMultipleComplete(ctx, ids.Ids); err != nil {
		return err
	}

	// return nil if no error
	return nil
}
