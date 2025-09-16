package repositories

import (
	"github.com/IbadT/tutor_app_back.git/internal/domain/auth"
	"github.com/IbadT/tutor_app_back.git/internal/domain/shared"
	"gorm.io/gorm"
)

// authRepository implements the auth.Repository interface
type authRepository struct {
	db *gorm.DB
}

// NewAuthRepository creates a new auth repository
func NewAuthRepository(db *gorm.DB) auth.Repository {
	return &authRepository{db: db}
}

// GetUserByEmail retrieves a user by email
func (r *authRepository) GetUserByEmail(email string) (*shared.User, error) {
	var u shared.User
	err := r.db.Where("email = ?", email).First(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

// CreateUser creates a new user
func (r *authRepository) CreateUser(u *shared.User) error {
	return r.db.Create(u).Error
}

// UserExists checks if a user exists by email
func (r *authRepository) UserExists(email string) (bool, error) {
	var count int64
	err := r.db.Model(&shared.User{}).Where("email = ?", email).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
