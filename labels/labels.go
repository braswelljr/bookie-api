package labels

import (
	"context"

	"github.com/go-playground/validator/v10"

	"encore.app/labels/store"
	"encore.app/users"
)

// Create - creates a new label.
//
//	@route POST /label/create
//	@param ctx - context.Context
//	@param payload
//	@return error
//
// encore:api public method=POST path=/label/:uid/create
func Create(ctx context.Context, uid string, payload *store.LabelRequest) error {
	// validate label details
	if err := validator.New().Struct(payload); err != nil {
		return err
	}

	// check if user exists
	user, err := users.Get(ctx, uid)
	if err != nil || user == nil {
		return err
	}

	// create a new label
	// if err := store.Create(ctx, payload); err != nil {
	//   return err
	// }

	// return the response
	return nil
}
