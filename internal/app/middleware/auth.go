package middleware

import (
	"strings"

	"github.com/IbadT/tutor_app_back.git/internal/domain/auth"
	"github.com/labstack/echo/v4"
)

// AuthMiddleware creates authentication middleware
func AuthMiddleware(authService auth.Service) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Get Authorization header
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return echo.NewHTTPError(401, "Authorization header required")
			}

			// Check if it starts with "Bearer "
			if !strings.HasPrefix(authHeader, "Bearer ") {
				return echo.NewHTTPError(401, "Invalid authorization header format")
			}

			// Extract token
			token := strings.TrimPrefix(authHeader, "Bearer ")

			// Validate token
			userID, role, err := authService.ValidateToken(token)
			if err != nil {
				return echo.NewHTTPError(401, "Invalid token")
			}

			// Set user context
			c.Set("user_id", userID)
			c.Set("user_role", role)

			return next(c)
		}
	}
}

// RoleMiddleware creates role-based access control middleware
func RoleMiddleware(allowedRoles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userRole, ok := c.Get("user_role").(string)
			if !ok {
				return echo.NewHTTPError(401, "User role not found in context")
			}

			// Check if user role is in allowed roles
			for _, role := range allowedRoles {
				if userRole == role {
					return next(c)
				}
			}

			return echo.NewHTTPError(403, "Insufficient permissions")
		}
	}
}
