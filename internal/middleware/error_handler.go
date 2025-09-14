package middleware

import (
	"net/http"

	"github.com/IbadT/tutor_app_back.git/internal/errors"
	"github.com/labstack/echo/v4"
)

// ErrorHandler is a custom error handler middleware
func ErrorHandler(err error, c echo.Context) {
	// Check if it's already an APIError
	if apiErr, ok := err.(*errors.APIError); ok {
		c.JSON(apiErr.Code, apiErr)
		return
	}

	// Handle different types of errors
	switch e := err.(type) {
	case *echo.HTTPError:
		// Echo HTTP errors
		apiErr := &errors.APIError{
			Code:    e.Code,
			Message: e.Message.(string),
		}
		c.JSON(e.Code, apiErr)
	case *errors.APIError:
		// Our custom API errors
		c.JSON(e.Code, e)
	default:
		// Unknown errors - return 500
		apiErr := &errors.APIError{
			Code:    http.StatusInternalServerError,
			Message: "Internal server error",
		}
		c.JSON(http.StatusInternalServerError, apiErr)
	}
}

// ValidationErrorHandler handles validation errors
func ValidationErrorHandler(err error, c echo.Context) {
	if validationErr, ok := err.(*echo.HTTPError); ok {
		apiErr := &errors.APIError{
			Code:    http.StatusBadRequest,
			Message: "Validation failed",
			Details: validationErr.Message.(string),
		}
		c.JSON(http.StatusBadRequest, apiErr)
		return
	}
	
	// Fallback
	ErrorHandler(err, c)
}
