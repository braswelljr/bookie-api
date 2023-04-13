package store

// import (
// 	"context"
// 	"fmt"
// 	"strings"
// 	"time"

// 	"encore.app/pkg/database"
// 	"encore.dev/storage/sqldb"
// 	"github.com/google/uuid"
// 	"github.com/jmoiron/sqlx"
// )

// // // get the service name
// var categoriesDatabase = sqlx.NewDb(sqldb.Named("categories").Stdlib(), "postgres")

// // FindByOneField - get user by field
// //
// //	@param ctx - context.Context
// //	@param field - string
// //	@param ops - string
// //	@param value - interface{}
// //	@return user
// //	@return error
// func FindOneByField(ctx context.Context, field, ops string, value interface{}) (Category, error) {
// 	// set the data fields for the query
// 	data := map[string]interface{}{
// 		field: value,
// 	}

// 	// query statement to be executed
// 	q := "SELECT * FROM categories WHERE %v %v :%v LIMIT 1"
// 	// format query parameters
// 	q = fmt.Sprintf(q, field, ops, field)

// 	// declare category
// 	var category Category
// 	// execute query
// 	if err := database.NamedStructQuery(ctx, categoriesDatabase, q, data, &category); err != nil {
// 		if err == database.ErrNotFound {
// 			return Category{}, ErrNotFound
// 		}
// 		return Category{}, fmt.Errorf("selecting categories by ID[%v]: %w", value, err)
// 	}

// 	return category, nil
// }

// // Create - Create a new category
// //
// //	@param ctx - context.Context
// //	@param payload - *CreateCategoryPayload
// //	@return error
// func Create(ctx context.Context, uid string, payload *CreateCategoryPayload) error {
// 	// check if category already exists
// 	cat, err := FindOneByField(ctx, "name", "=", strings.ToLower(payload.Name))
// 	if err == nil {
// 		return err
// 	}
// 	// check if category ID is empty (if not, category already exists)
// 	if len(cat.ID) > 0 {
// 		return ErrAlreadyExists
// 	}

// 	// create category
// 	category := Category{
// 		ID:          uuid.New().String(),
// 		UID:         uid,
// 		Name:        strings.ToLower(payload.Name),
// 		Description: payload.Description,
// 		CreatedAt:   time.Now(),
// 		UpdatedAt:   time.Now(),
// 	}

// 	// query statement to be executed
// 	query := `
//     INSERT INTO categories (id, uid, name, description, createdAt, updatedAt)
//     VALUES (:id, :uid, :name, :description, :createdAt, :updatedAt)
//   `

// 	// create category
// 	if err := database.NamedExecQuery(ctx, categoriesDatabase, query, category); err != nil {
// 		return fmt.Errorf("creating category: %w", err)
// 	}

// 	return nil
// }
