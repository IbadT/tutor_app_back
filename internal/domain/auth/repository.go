package auth

import (
	"github.com/IbadT/tutor_app_back.git/internal/domain/shared"
)

// Repository defines the interface for authentication data operations
type Repository interface {
	// User operations for authentication
	GetUserByEmail(email string) (*shared.User, error)
	CreateUser(user *shared.User) error
	UserExists(email string) (bool, error)
}
