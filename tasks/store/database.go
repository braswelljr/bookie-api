package store

import (
	"context"
	"fmt"
	"time"

	"encore.dev/storage/sqldb"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"encore.app/pkg/database"
)

// get the service name
var tasksDatabase = sqlx.NewDb(sqldb.Named("tasks").Stdlib(), "postgres")

// FindByOneField - get task by field
//
//	@param ctx - context.Context
//	@param field - string
//	@param ops - string
//	@param value - interface{}
//	@return user
//	@return error
func FindOneByField(ctx context.Context, field, ops string, value interface{}) (Task, error) {
	// set the data fields for the query
	data := map[string]interface{}{
		field: value,
	}

	// query statement to be executed
	q := "SELECT * FROM tasks WHERE %v %v :%v LIMIT 1"
	// format query parameters
	q = fmt.Sprintf(q, field, ops, field)

	// declare task
	var task Task
	// execute query
	if err := database.NamedStructQuery(ctx, tasksDatabase, q, data, &task); err != nil {
		if err == database.ErrNotFound {
			return Task{}, ErrNotFound
		}
		return Task{}, fmt.Errorf("selecting tasks by ID[%v]: %w", value, err)
	}

	return task, nil
}

// FindManyByField - get tasks by field
//
//	@param ctx - context.Context
//	@param field - string
//	@param ops - string
//	@param value - interface{}
//	@return tasks
//	@return error
func FindManyByField(ctx context.Context, field, ops string, value interface{}) ([]Task, error) {
	// set the data fields for the query
	data := map[string]interface{}{
		field: value,
	}

	// query statement to be executed
	q := "SELECT * FROM tasks WHERE %v %v :%v"
	// format query parameters
	q = fmt.Sprintf(q, field, ops, field)

	// declare tasks
	var tasks []Task
	// execute query
	if err := database.NamedStructQuery(ctx, tasksDatabase, q, data, &tasks); err != nil {
		if err == database.ErrNotFound {
			return []Task{}, ErrNotFound
		}
		return []Task{}, fmt.Errorf("selecting tasks by ID[%v]: %w", value, err)
	}

	return tasks, nil
}

// Create - Create is a function that creates a new task.
//
// @param ctx - context.Context
// @param payload
// @return task
// @return error
func Create(ctx context.Context, id string, payload *CreateTaskPayload) (*Task, error) {
	// declare task
	task := Task{}

	// set the fields for the task
	task.ID = uuid.New().String()
	task.UserID = id
	task.Title = payload.Title
	task.Description = payload.Description
	task.Status = "pending"
	task.Pinned = false
	task.Archived = false
	task.Color = "default"
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	// query statement to be executed
	q := `INSERT INTO tasks (id, user_id, title, description, status, pinned, archived, color, created_at, updated_at) VALUES (:id, :user_id, :title, :description, :status, :pinned, :archived, :color, :created_at, :updated_at) RETURNING *`

	// execute query
	if err := database.NamedExecQuery(ctx, tasksDatabase, q, task); err != nil {
		return nil, fmt.Errorf("inserting task: %w", err)
	}

	// check if task was created
	tsk, err := FindOneByField(ctx, "id", "=", task.ID)
	if err != nil {
		return nil, fmt.Errorf("selecting task: %w", err)
	}

	return &tsk, nil
}

// Get - Get is a function that gets a task.
//
// @param ctx - context.Context
// @param id - string
// @return task
// @return error
func Get(ctx context.Context, id string) (*Task, error) {
	// check if task exists
	task, err := FindOneByField(ctx, "id", "=", id)
	if err != nil {
		return nil, fmt.Errorf("selecting task: %w", err)
	}

	// return task
	return &task, nil
}

// Delete - Delete is a function that deletes a task.
//
// @param ctx - context.Context
// @param id - string
// @return error
func Delete(ctx context.Context, id string) error {
	// query statement to be executed
	q := "DELETE FROM tasks WHERE id = :id"

	// execute query
	if err := database.NamedExecQuery(ctx, tasksDatabase, q, map[string]interface{}{
		"id": id,
	}); err != nil {
		return fmt.Errorf("deleting task: %w", err)
	}

	// Delete was successful
	return nil
}
