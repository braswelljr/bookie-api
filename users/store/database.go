package store

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
	"encore.app/pkg/middleware"
	"encore.app/pkg/pagination"
)

// get the service name
var usersDatabase = sqlx.NewDb(sqldb.Named("users").Stdlib(), "postgres")

// FindByOneField - get user by field
//
//	@param ctx - context.Context
//	@param field - string
//	@param ops - string
//	@param value - interface{}
//	@return user
//	@return error
func FindOneByField(ctx context.Context, field, ops string, value interface{}) (User, error) {
	// set the data fields for the query
	data := map[string]interface{}{
		field: value,
	}

	// query statement to be executed
	q := "SELECT * FROM users WHERE %v %v :%v LIMIT 1"
	// format query parameters
	q = fmt.Sprintf(q, field, ops, field)

	// declare user
	var user User
	// execute query
	if err := database.NamedStructQuery(ctx, usersDatabase, q, data, &user); err != nil {
		if err == database.ErrNotFound {
			return User{}, ErrNotFound
		}
		return User{}, fmt.Errorf("selecting users by ID[%v]: %w", value, err)
	}

	return user, nil
}

// Create - Create is a function that creates a new user.

// @param ctx - context.Context
// @param payload
// @return user
// @return error
func Create(ctx context.Context, payload *SignupPayload) (*User, error) {
	// create user from payload
	user := User{}

	user.ID = uuid.New().String()
	user.Firstname = strings.TrimSpace(payload.Firstname)
	user.Lastname = strings.TrimSpace(payload.Lastname)
	user.Othernames = strings.TrimSpace(payload.Othernames)
	user.Username = strings.TrimSpace(payload.Username)
	user.Email = strings.TrimSpace(payload.Email)
	user.DateOfBirth = payload.DateOfBirth
	user.Phone = strings.TrimSpace(payload.Phone)
	user.Address = strings.TrimSpace(payload.Address)
	user.City = strings.TrimSpace(payload.City)
	user.Country = strings.TrimSpace(payload.Country)
	user.Role = "admin"

	// hash password
	password, err := middleware.HashPassword(payload.Password)
	if err != nil {
		return &User{}, err
	}
	user.Password = password
	// set created at and updated at
	user.CreatedAt = time.Now().UTC()
	user.UpdatedAt = time.Now().UTC()

	// create query
	query := `
    INSERT INTO users (
      id, firstname, lastname, othernames, username, email, date_of_birth, password, phone, address, city, country, role, created_at, updated_at
    ) 
    VALUES (
      :id, :firstname, :lastname, :othernames, :username, :email, :date_of_birth, :password, :phone, :address, :city, :country, :role, :created_at, :updated_at
    )
  `
	// ON CONFLICT (email) DO NOTHING
	//   ON CONFLICT (username) DO NOTHING

	// insert user into database
	if err := database.NamedExecQuery(ctx, usersDatabase, query, user); err != nil {
		return &User{}, err
	}

	// query user from database
	usr, err := FindOneByField(ctx, "ID", "=", user.ID)
	if err != nil {
		return &User{}, err
	}

	return &usr, nil
}

// Get - Get is a function that gets a user by field.
//
//	@param ctx - context.Context
//	@param email
//	@return user
//	@return error
func Get(ctx context.Context, email string) (*User, error) {
	// query user from database
	user, err := FindOneByField(ctx, "email", "=", email)
	if err != nil {
		return &User{}, err
	}

	return &user, nil
}

// GetWithID - GetWithID is a function that gets a user by field.
//
//	@param ctx - context.Context
//	@param id
//	@return user
//	@return error
func GetWithID(ctx context.Context, id string) (*User, error) {
	// query user from database
	user, err := FindOneByField(ctx, "id", "=", id)
	if err != nil {
		return &User{}, err
	}

	return &user, nil
}

// Update - Update is a function that updates a user.
//
//	@param ctx - context.Context
//	@param id
//	@param payload
//	@return user
//	@return error
func Update(ctx context.Context, id string, payload UpdatePayload) (*User, error) {
	// query user from database
	user, err := GetWithID(ctx, id)
	if err != nil {
		return &User{}, err
	}

	// map for query fields
	fields := map[string]interface{}{}

	// if not empty, update user field
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
	fields["id"] = user.ID

	// loop through fields and create query fields
	for k := range fields {
		ks = append(ks, fmt.Sprintf("%v = :%v", k, k))
	}

	// create query with query fields and join them with commas
	query := fmt.Sprintf("UPDATE users SET %v WHERE id = :id", strings.Join(ks, ", "))

	// update user in database
	if err := database.NamedExecQuery(ctx, usersDatabase, query, fields); err != nil {
		return &User{}, err
	}

	// query user from database
	usr, err := GetWithID(ctx, id)
	if err != nil {
		return &User{}, err
	}

	return usr, nil
}

// UpdateRole - UpdateRole is a function that updates a user's role.
//
//	@param ctx - context.Context
//	@param id
//	@param role
//	@return user
//	@return error
func UpdateRole(ctx context.Context, payload *UpdateRolePayload, role string) error {
	// query user from database
	user, err := GetWithID(ctx, payload.ID)
	if err != nil {
		return err
	}

	// update user in database
	if err := database.NamedExecQuery(ctx, usersDatabase, "UPDATE users SET role = :role, updated_at = :updated_at WHERE id = :id", map[string]interface{}{
		"role":       role,
		"updated_at": time.Now().UTC(),
		"id":         user.ID,
	}); err != nil {
		return err
	}

	return nil
}

// GetAll - GetAll is a function that gets all users.
//
//	@param ctx - context.Context
//	@return users
//	@return error
func GetAll(ctx context.Context, pag *pagination.Options) (*PaginatedUsersResponse, error) {
	var users []User = []User{}

	// create query
	countQuery := `SELECT COUNT(*) FROM users`

	// get total count of users
	// get count of categories
	count, err := database.NamedCountQuery(ctx, usersDatabase, countQuery, map[string]interface{}{})
	if err != nil {
		return nil, fmt.Errorf("getting count of users: %w", err)
	}

	// set limit to 20 if it is less than 0 or greater than count
	if pag.Limit < 1 || pag.Limit > count {
		pag.Limit = 20
	}

	// calculate for pagination
	// var paginate pagination.Paginate
	// initialize pagination
	paging := pagination.New(pag.Page, pag.Limit, count)

	// if page is greater than total pages, set page to total pages
	if pag.Page > paging.Pages() {
		paging.SetPage(paging.Pages())
	}

	// query to set offset and limit
	const query = `SELECT * FROM users LIMIT :limit OFFSET :offset`
	// data to be passed to the query
	p := struct {
		Limit  int `db:"limit" json:"limit" validate:"omitempty"`
		Offset int `db:"offset" json:"offset" validate:"omitempty"`
	}{
		Limit:  paging.PerPage(),
		Offset: paging.Offset(),
	}

	// execute query
	if err := database.NamedSliceQuery(ctx, usersDatabase, query, p, &users); err != nil {
		return nil, fmt.Errorf("getting categories: %w", err)
	}

	usersResponse := []UserResponse{}

	for _, user := range users {
		usersResponse = append(usersResponse, UserResponse{
			ID:          user.ID,
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
			CreatedAt:   user.CreatedAt,
			UpdatedAt:   user.UpdatedAt,
		})
	}

	return &PaginatedUsersResponse{
		TotalPages:  paging.Pages(),
		Total:       paging.Total(),
		CurrentPage: paging.Page(),
		Users:       usersResponse,
	}, nil
}
