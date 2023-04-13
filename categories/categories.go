package category

import (
	"context"

	"github.com/go-playground/validator/v10"

	"encore.app/categories/store"
)

// Create - Create a new category
//
//	@param ctx - context.Context
//	@param payload - *CreateCategoryPayload
//
// encore:api auth method=POST path=/categories/:uid/create
func Create(ctx context.Context, uid string, payload *store.CreateCategoryPayload) error {
	// validate payload
	if err := validator.New().Struct(payload); err != nil {
		return err
	}

	// create category
	// if err := store.Create(ctx, uid, payload); err != nil {
	// 	return err
	// }

	return nil
}

// Get - Get a category
//
//	@param ctx - context.Context
