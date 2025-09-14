package userservice

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUsersInfo(userID uuid.UUID) (UserInfo, error)
	CreateUserInfo(userInfo UserInfo) error
	UpdateUserInfo(userID uuid.UUID, userInfo UpdateUserInfo) error
	DeleteUser(userID uuid.UUID) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetUsersInfo(userID uuid.UUID) (UserInfo, error) {
	var user UserInfo
	err := r.db.Where("user_id = ?", userID).First(&user).Error
	if err != nil {
		return UserInfo{}, err
	}
	return user, nil
}

func (r *userRepository) CreateUserInfo(userInfo UserInfo) error {
	if err := r.db.Create(&userInfo).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) UpdateUserInfo(userID uuid.UUID, userInfo UpdateUserInfo) error {
	if err := r.db.Model(&UserInfo{}).Where("user_id = ?", userID).Updates(&userInfo).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) DeleteUser(userID uuid.UUID) error {
	if err := r.db.Where("id = ?", userID).Delete(&User{}).Error; err != nil {
		return err
	}
	return nil
}
