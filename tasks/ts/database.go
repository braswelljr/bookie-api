package ts

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"time"

	"encore.dev/storage/sqldb"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"encore.app/pkg/database"
	"encore.app/pkg/pagination"
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
func Create(ctx context.Context, id string, payload *CreateTaskPayload) error {
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
	q := `
    INSERT INTO tasks (id, uid, title, description, status, pinned, archived, color, created_at, updated_at) 
    VALUES (:id, :uid, :title, :description, :status, :pinned, :archived, :color, :created_at, :updated_at) 
    RETURNING *
  `

	// execute query
	if err := database.NamedExecQuery(ctx, tasksDatabase, q, task); err != nil {
		return fmt.Errorf("inserting task: %w", err)
	}

	// check if task was created
	if tsk, err := FindOneByField(ctx, "id", "=", task.ID); err != nil || reflect.DeepEqual(tsk, Task{}) {
		return fmt.Errorf("selecting task: %w", err)
	}

	return nil
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

// GetMany - GetMany is a function that gets many tasks.
//
// @param ctx - context.Context
// @param ids - []string
// @return tasks
// @return error
func GetMany(ctx context.Context, ids []string) ([]Task, error) {
	// check if task exists
	tasks, err := FindManyByField(ctx, "id", "=", ids)
	if err != nil {
		return nil, fmt.Errorf("selecting task: %w", err)
	}

	// return task
	return tasks, nil
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

// DeleteMany - DeleteMany is a function that deletes many tasks.
//
// @param ctx - context.Context
// @param ids - []string
// @return error
func DeleteMany(ctx context.Context, ids []string) error {
	// query statement to be executed
	q := `
    DELETE FROM tasks
    WHERE id IN (:ids)
  `

	// execute query
	if err := database.NamedExecQuery(ctx, tasksDatabase, q, map[string]interface{}{
		"ids": ids,
	}); err != nil {
		return fmt.Errorf("deleting tasks: %w", err)
	}

	// Delete was successful
	return nil
}

// DeleteAllWithUserID - DeleteAllWithUserID is a function that deletes all tasks with a user ID.
//
// @param ctx - context.Context
// @param id - string
// @return error
func DeleteAllWithUserID(ctx context.Context, id string) error {
	// query statement to be executed
	q := `
    DELETE FROM tasks
    WHERE uid = :uid
  `

	// execute query
	if err := database.NamedExecQuery(ctx, tasksDatabase, q, map[string]interface{}{
		"uid": id,
	}); err != nil {
		return fmt.Errorf("deleting tasks: %w", err)
	}

	// Delete was successful
	return nil
}

// Update - Update is a function that updates a task.
//
// @param ctx - context.Context
// @param id - string
// @param payload
// @return task
// @return error
func Update(ctx context.Context, id string, payload *UpdateTaskPayload) error {
	// check if task exists
	task, err := FindOneByField(ctx, "id", "=", id)
	if err != nil {
		return fmt.Errorf("selecting task: %w", err)
	}

	// map for query fields
	fields := map[string]interface{}{}

	// if not empty, update task field
	vp := reflect.ValueOf(payload)

	// loop through payload fields and check for empty values
	for i := 0; i < vp.NumField(); i++ {
		// get the db tag name of the field
		field := vp.Type().Field(i).Tag.Get("db")
		// get the value of the field
		value := vp.Field(i).Interface()

		// if the value is not empty, add it to the fields map
		if len(strings.TrimSpace(value.(string))) > 0 {
			fields[field] = value
		}
	}

	// create query fields
	var ks []string

	fields["updated_at"] = time.Now().UTC()
	fields["id"] = task.ID

	// loop through fields and create query fields
	for k := range fields {
		ks = append(ks, fmt.Sprintf("%v = :%v", k, k))
	}

	// query statement to be executed
	q := fmt.Sprintf("UPDATE tasks SET %v WHERE id = :id RETURNING *", strings.Join(ks, ", "))

	// execute query
	if err := database.NamedExecQuery(ctx, tasksDatabase, q, fields); err != nil {
		return fmt.Errorf("updating task: %w", err)
	}

	// return task
	return nil
}

// GetUserTasks - GetUserTasks is a function that gets a user's tasks.
//
// @param ctx - context.Context
// @param uid - string
// @return tasks
// @return error
func GetUserTasks(ctx context.Context, uid string, options *pagination.Options) (*PaginatedTasksResponse, error) {
	// declare tasks
	var tasks []Task = []Task{}

	// query statement to be executed
	countQuery := "SELECT COUNT(*) FROM tasks WHERE uid = :uid"

	// execute query
	count, err := database.NamedCountQuery(ctx, tasksDatabase, countQuery, map[string]interface{}{"uid": uid})

	// check for errors
	if err != nil {
		return nil, fmt.Errorf("counting tasks: %w", err)
	}

	// set limit to 20 if it is less than 0 or greater than count
	if options.Limit < 1 || options.Limit > count {
		options.Limit = 20
	}

	// calculate for pagination
	// var paginate pagination.Paginate
	// initialize pagination
	paging := pagination.New(options.Page, options.Limit, count)

	// if page is greater than total pages, set page to total pages
	if options.Page > paging.Pages() {
		paging.SetPage(paging.Pages())
	}

	// query statement to be executed
	query := `
    SELECT * FROM tasks
    WHERE uid = :uid
    ORDER BY created_at
    DESC LIMIT :limit OFFSET :offset
  `

	p := struct {
		UID    string `db:"uid" json:"uid" validate:"required" url:"uid"`
		Limit  int    `db:"limit" json:"limit" validate:"omitempty" url:"limit"`
		Offset int    `db:"offset" json:"offset" validate:"omitempty" url:"offset"`
	}{
		UID:    uid,
		Limit:  paging.PerPage(),
		Offset: paging.Offset(),
	}

	// execute query
	if err := database.NamedSliceQuery(ctx, tasksDatabase, query, p, &tasks); err != nil {
		return nil, fmt.Errorf("selecting tasks: %w", err)
	}

	return &PaginatedTasksResponse{
		TotalPages:  paging.Pages(),
		Total:       paging.Total(),
		CurrentPage: paging.Page(),
		Tasks:       tasks,
	}, nil
}

// ToggleComplete - ToggleComplete is a function that toggles a task's complete status.
//
// @param ctx - context.Context
// @param id - string
// @return error
func ToggleComplete(ctx context.Context, id string) error {
	// check if task exists
	task, err := FindOneByField(ctx, "id", "=", id)
	if err != nil {
		return fmt.Errorf("selecting task: %w", err)
	}

	// query statement to be executed
	query := `
    UPDATE tasks 
    SET completed = :completed, completed_at = :completed_at, updated_at = :updated_at  
    WHERE id = :id
  `

	// execute query
	if err := database.NamedExecQuery(ctx, tasksDatabase, query, map[string]interface{}{
		"id":          task.ID,
		"completed":   !task.Completed,
		"completedAt": time.Now().UTC(),
		"updatedAt":   time.Now().UTC(),
	}); err != nil {
		return fmt.Errorf("updating task: %w", err)
	}

	// return task
	return nil
}

// ToggleMultipleComplete - ToggleMultipleComplete is a function that toggles multiple tasks' complete status.
//
// @param ctx - context.Context
// @param ids - []string
// @return error
func ToggleMultipleComplete(ctx context.Context, ids []string) error {
	// query statement to be executed
	query := `
    UPDATE tasks 
    SET completed = :completed, completed_at = :completed_at, updated_at = :updated_at 
    WHERE id = ANY(:ids)
  `

	// execute query
	if err := database.NamedExecQuery(ctx, tasksDatabase, query, map[string]interface{}{
		"ids":          ids,
		"completed":    true,
		"completed_at": time.Now().UTC(),
		"updated_at":   time.Now().UTC(),
	}); err != nil {
		return fmt.Errorf("updating tasks: %w", err)
	}

	// return task
	return nil
}
