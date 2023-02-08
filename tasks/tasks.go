package tasks

import (
	"context"

	"github.com/go-playground/validator/v10"

	"encore.app/tasks/store"
)

// Create - Create is a function that creates a new task.
//
// @param ctx - context.Context
// @param payload
// @return task
// @return error
//
// encore:api auth method=POST path=/tasks/:uid/create
func Create(ctx context.Context, uid string, payload *store.CreateTaskPayload) (*store.Task, error) {
	// validate payload
	if err := validator.New().Struct(payload); err != nil {
		return nil, err
	}

	// create task
	task, err := store.Create(ctx, uid, payload)
	if err != nil {
		return nil, err
	}

	return task, nil
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

// Delete - Delete a task
//
// @param ctx - context.Context
// @param id - string
// @return task
// @return error
//
// encore:api auth method=DELETE path=/tasks/:id
func Delete(ctx context.Context, id string) error {
	// delete task
	if err := store.Delete(ctx, id); err != nil {
		return err
	}

	// return nil if no error
	return nil
}
