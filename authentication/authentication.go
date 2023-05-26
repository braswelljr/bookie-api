package authentication

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"encore.dev/beta/auth"
	"github.com/go-playground/validator/v10"

	"encore.app/authentication/store"
	"encore.app/pkg/middleware"
	"encore.app/users"
	us "encore.app/users/store"
)

// Signup is a function that handles the signup process.
//
//	@route POST /signup
//	@param ctx - context.Context
//	@param payload
//	@return response
//	@return error
//
// encore:api public method=POST path=/signup
func Signup(ctx context.Context, payload *us.SignupPayload) (*store.Response, error) {
	// validate user details
	if err := validator.New().Struct(payload); err != nil {
		return &store.Response{}, err
	}

	// create a new user
	user, err := users.Create(ctx, payload)
	if err != nil {
		return nil, err
	}

	// generate tokens
	token, err := middleware.GetToken(&middleware.User{
		ID:          user.ID,
		Firstname:   user.Firstname,
		Lastname:    user.Lastname,
		Othernames:  user.Othernames,
		Username:    user.Username,
		Email:       user.Email,
		DateOfBirth: user.DateOfBirth,
		Phone:       user.Phone,
		Role:        user.Role,
	})
	if err != nil {
		return &store.Response{}, errors.New("authentication failed: unable to generate token")
	}

	// return the response
	return &store.Response{
		Message: "Signup successful",
		Token:   token,
		Payload: &store.UserResponse{
			ID:          user.ID,
			Firstname:   user.Firstname,
			Lastname:    user.Lastname,
			Othernames:  user.Othernames,
			Username:    user.Username,
			Email:       user.Email,
			DateOfBirth: user.DateOfBirth,
			Phone:       user.Phone,
			Role:        user.Role,
			CreatedAt:   user.CreatedAt,
			UpdatedAt:   user.UpdatedAt,
		},
	}, nil
}

// Login - Login is a function that handles the login process for a user.
//
//	@route POST /login
//	@param w http.ResponseWriter
//	@param req *http.Request
//
// encore:api public raw method=POST path=/login
func Login(w http.ResponseWriter, req *http.Request) {
	// Set headers for general response
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	// Get the user details from the request
	email, password, ok := req.BasicAuth()
	if !ok {
		writeJSONErrorResponse(w, "authentication failed: invalid credentials", http.StatusUnauthorized)
		return
	}

	// Get the user
	user, err := us.Get(req.Context(), email)
	if err != nil {
		writeJSONErrorResponse(w, "authentication failed: invalid credentials", http.StatusUnauthorized)
		return
	}

	// Check if the password is correct
	isCorrect, err := middleware.ComparePasswords(user.Password, password)
	if err != nil || !isCorrect {
		writeJSONErrorResponse(w, "authentication failed: invalid credentials", http.StatusUnauthorized)
		return
	}

	// Generate tokens
	token, err := middleware.GetToken(&middleware.User{
		ID:          user.ID,
		Firstname:   user.Firstname,
		Lastname:    user.Lastname,
		Othernames:  user.Othernames,
		Username:    user.Username,
		Email:       user.Email,
		DateOfBirth: user.DateOfBirth,
		Phone:       user.Phone,
		Role:        user.Role,
	})
	if err != nil {
		writeJSONErrorResponse(w, "authentication failed: unable to generate token", http.StatusInternalServerError)
		return
	}

	// Create the response object
	response := &store.Response{
		Message: "Login successful",
		Token:   token,
		Payload: &store.UserResponse{
			ID:          user.ID,
			Firstname:   user.Firstname,
			Lastname:    user.Lastname,
			Othernames:  user.Othernames,
			Username:    user.Username,
			Email:       user.Email,
			DateOfBirth: user.DateOfBirth,
			Phone:       user.Phone,
			Role:        user.Role,
			CreatedAt:   user.CreatedAt,
			UpdatedAt:   user.UpdatedAt,
		},
	}

	// Convert the response to JSON
	responseJSON, err := json.Marshal(response)
	if err != nil {
		writeJSONErrorResponse(w, "authentication failed: unable to generate token", http.StatusInternalServerError)
		return
	}

	// Write the response
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(responseJSON); err != nil {
		writeJSONErrorResponse(w, "authentication failed: unable to write response", http.StatusInternalServerError)
		return
	}
}

// writeJSONErrorResponse writes the specified error message as a JSON response with the provided status code.
func writeJSONErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	response := map[string]interface{}{
		"message": message,
		"code":    statusCode,
		"token":   "",
		"payload": "",
	}

	responseJSON, err := json.Marshal(response)
	if err != nil {
		// If unable to marshal the error response, fallback to a generic error message
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"error": "An internal server error occurred"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, _ = w.Write(responseJSON)
}

// Logout - Logout is a function that handles the logout process for a user.
//
//	@route POST /logout
//	@param ctx - context.Context
//	@return response
//	@return error
//
// encore:api public method=POST path=/logout
func Logout(_ context.Context) (*store.Response, error) {

	return &store.Response{
		Message: "Logout successful",
		Payload: nil,
		Token:   "",
	}, nil
}

// Auth - Auth is a function that handles the authentication process for a user.
//
//	@route POST /auth
//	@param ctx - context.Context
//	@param token - string
//	@return response
//	@return error
//
// encore:authhandler
func Auth(_ context.Context, token string) (auth.UID, *middleware.DataI, error) {
	// check for empty token
	if len(strings.TrimSpace(token)) < 1 {
		return "", &middleware.DataI{}, errors.New("authentication failed: token is empty")
	}

	// validate token
	claims, err := middleware.ValidateToken(token)
	if err != nil {
		return "", &middleware.DataI{}, errors.New("authentication failed: invalid token")
	}

	return auth.UID(claims.User.ID), &middleware.DataI{
		Subject: claims.User,
		Roles:   []string{claims.User.Role},
	}, nil
}
