package middleware

import (
	"context"

	"encore.dev/beta/auth"
	"encore.dev/beta/errs"
)

// GetVerifiedClaims - is a function that handles the retrieval of verified claims.
//
//	@param ctx - context.Context
//	@param role - string
//	@return *SignedParams
//	@return error
func GetVerifiedClaims(ctx context.Context, role string) (*DataI, error) {
	// get the claims from the context
	claims, ok := ctx.Value(ContextKey).(*DataI)
	if !ok {
		// get the claims from the encore auth
		claims, ok = auth.Data().(*DataI)
		if !ok {
			return &DataI{}, &errs.Error{Code: errs.Unauthenticated, Message: "unable to get claims"}
		}
	}

	// check if the user has the required role
	if role != "" && !claims.HasRole(role) {
		return &DataI{}, &errs.Error{Code: errs.Unauthenticated, Message: "user does not have the required role"}
	}

	return claims, nil
}
