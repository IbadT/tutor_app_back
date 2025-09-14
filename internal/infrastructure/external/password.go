package external

import (
	"github.com/IbadT/tutor_app_back.git/internal/domain/auth"
	"golang.org/x/crypto/bcrypt"
)

// passwordService implements the auth.PasswordHasher interface
type passwordService struct {
	cost int
}

// NewPasswordService creates a new password service
func NewPasswordService() auth.PasswordHasher {
	return &passwordService{
		cost: bcrypt.DefaultCost,
	}
}

// HashPassword hashes a password using bcrypt
func (p *passwordService) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), p.cost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// ComparePassword compares a password with its hash
func (p *passwordService) ComparePassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
