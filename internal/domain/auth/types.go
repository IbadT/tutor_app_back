package auth

import "github.com/google/uuid"

// LoginRequest represents the login request
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

// RegisterRequest represents the registration request
type RegisterRequest struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=6"`
	Role      string `json:"role" validate:"required,oneof=student tutor admin"`
	Location  string `json:"location" validate:"required"`
}

// LoginResponse represents the authentication response
type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// RefreshTokenRequest represents the refresh token request
type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

// TokenGenerator defines the interface for token operations
type TokenGenerator interface {
	GenerateToken(userID uuid.UUID, role string) (*LoginResponse, error)
	ParseToken(tokenString string) (map[string]interface{}, error)
	ValidateToken(tokenString string) error
}

// PasswordHasher defines the interface for password operations
type PasswordHasher interface {
	HashPassword(password string) (string, error)
	ComparePassword(password, hashedPassword string) bool
}
