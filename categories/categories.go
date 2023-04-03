package category

import (
	"context"

	"encore.app/categories/store"
	"github.com/go-playground/validator/v10"
	// "encore.app/categories/store"
)

// Create - Create a new category
//
// POST /categories/create/:uid
//
//	{
//	  "name": "string",
//	  "description": "string"
//	}
//
// encore:api auth method=POST path=/categories/create
func Create(ctx context.Context, payload *store.CreateCategoryPayload) error {
	// validate payload
	if err := validator.New().Struct(payload); err != nil {
		return err
	}

	// create category
	// if err := store.Create(ctx, payload); err != nil {
	// 	return err
	// }

	return nil
}
