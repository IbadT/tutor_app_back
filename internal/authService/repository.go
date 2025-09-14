package authservice

import (
	userservice "github.com/IbadT/tutor_app_back.git/internal/userService"
	"gorm.io/gorm"
)

type AuthRepository interface {
	CheckUserExists(email string) (userservice.User, error)
	Register(user userservice.User) error
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) CheckUserExists(email string) (userservice.User, error) {
	var user userservice.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return userservice.User{}, err
	}
	return user, nil
}

func (r *authRepository) Register(user userservice.User) error {
	if err := r.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}
