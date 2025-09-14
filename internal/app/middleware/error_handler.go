package middleware

import (
	"log"
	"net/http"

	"github.com/IbadT/tutor_app_back.git/internal/domain/shared"
	"github.com/labstack/echo/v4"
)

// ErrorHandler is a custom error handler for Echo
func ErrorHandler(err error, c echo.Context) {
	// Check if it's an HTTP error
	if he, ok := err.(*echo.HTTPError); ok {
		response := map[string]interface{}{
			"code":    he.Code,
			"message": he.Message,
		}
		if err := c.JSON(he.Code, response); err != nil {
			log.Printf("Failed to send error response: %v", err)
		}
		return
	}

	// Check if it's our custom API error
	if apiErr, ok := err.(*shared.APIError); ok {
		response := map[string]interface{}{
			"code":    apiErr.Code,
			"message": apiErr.Message,
		}
		if apiErr.Details != "" {
			response["details"] = apiErr.Details
		}
		if err := c.JSON(apiErr.Code, response); err != nil {
			log.Printf("Failed to send error response: %v", err)
		}
		return
	}

	// Default error response
	response := map[string]interface{}{
		"code":    http.StatusInternalServerError,
		"message": "Internal server error",
	}
	if err := c.JSON(http.StatusInternalServerError, response); err != nil {
		log.Printf("Failed to send error response: %v", err)
	}
}
