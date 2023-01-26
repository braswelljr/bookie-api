package store

import (
	"context"
	"fmt"
	"strings"
	"time"

	"encore.dev/storage/sqldb"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"encore.app/pkg/database"
	"encore.app/pkg/middleware"
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
