package shared

import (
	"encoding/json"
	"net/http"
)

// APIError represents an API error response
type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

// Error implements the error interface
func (e *APIError) Error() string {
	return e.Message
}

// Predefined errors
var (
	// 400 Bad Request
	ErrBadRequest = &APIError{
		Code:    http.StatusBadRequest,
		Message: "Bad Request",
	}

	ErrInvalidInput = &APIError{
		Code:    http.StatusBadRequest,
		Message: "Invalid input data",
	}

	ErrMissingFields = &APIError{
		Code:    http.StatusBadRequest,
		Message: "Missing required fields",
	}

	// 401 Unauthorized
	ErrUnauthorized = &APIError{
		Code:    http.StatusUnauthorized,
		Message: "Unauthorized",
	}

	ErrInvalidCredentials = &APIError{
		Code:    http.StatusUnauthorized,
		Message: "Invalid credentials",
	}

	// 409 Conflict
	ErrUserAlreadyExists = &APIError{
		Code:    http.StatusConflict,
		Message: "User already exists",
	}

	// 500 Internal Server Error
	ErrInternalServer = &APIError{
		Code:    http.StatusInternalServerError,
		Message: "Internal server error",
	}

	ErrDatabaseError = &APIError{
		Code:    http.StatusInternalServerError,
		Message: "Database error",
	}

	ErrTokenGeneration = &APIError{
		Code:    http.StatusInternalServerError,
		Message: "Token generation failed",
	}
)

// NewAPIError creates a new API error with custom message
func NewAPIError(code int, message string) *APIError {
	return &APIError{
		Code:    code,
		Message: message,
	}
}

// NewAPIErrorWithDetails creates a new API error with custom message and details
func NewAPIErrorWithDetails(code int, message, details string) *APIError {
	return &APIError{
		Code:    code,
		Message: message,
		Details: details,
	}
}

// WriteErrorResponse writes error response to HTTP response
func WriteErrorResponse(w http.ResponseWriter, err *APIError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.Code)
	json.NewEncoder(w).Encode(err)
}
