package auth

import (
	"github.com/IbadT/tutor_app_back.git/internal/domain/user"
	"github.com/google/uuid"
)

// Repository defines the interface for authentication data operations
type Repository interface {
	// User operations for authentication
	GetUserByEmail(email string) (*user.User, error)
	CreateUser(user *user.User) error
	UserExists(email string) (bool, error)
}

// Service defines the interface for authentication business logic
type Service interface {
	Login(req *LoginRequest) (*LoginResponse, error)
	Register(req *RegisterRequest) (*LoginResponse, error)
	RefreshToken(refreshToken string) (*LoginResponse, error)
	ValidateToken(tokenString string) (uuid.UUID, string, error)
}
