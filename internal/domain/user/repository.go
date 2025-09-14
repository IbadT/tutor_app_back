package user

import (
	"github.com/google/uuid"
)

// Repository defines the interface for user data operations
type Repository interface {
	// User operations
	GetByID(id uuid.UUID) (*User, error)
	GetByEmail(email string) (*User, error)
	Create(user *User) error
	Update(user *User) error
	Delete(id uuid.UUID) error

	// UserInfo operations
	GetUserInfo(userID uuid.UUID) (*UserInfo, error)
	CreateUserInfo(userInfo *UserInfo) error
	UpdateUserInfo(userID uuid.UUID, userInfo *UpdateUserInfoRequest) error
	DeleteUserInfo(userID uuid.UUID) error

	// UserStats operations
	GetUserStats(userID uuid.UUID) (*UserStats, error)
	CreateUserStats(stats *UserStats) error
	UpdateUserStats(stats *UserStats) error

	// UserAchievements operations
	GetUserAchievements(userID uuid.UUID) ([]UserAchievements, error)
	CreateUserAchievement(achievement *UserAchievements) error
	DeleteUserAchievement(id uuid.UUID) error

	// UserBadges operations
	GetUserBadges(userID uuid.UUID) ([]UserBadges, error)
	CreateUserBadge(badge *UserBadges) error
	DeleteUserBadge(id uuid.UUID) error
}

// Service defines the interface for user business logic
type Service interface {
	GetUserInfo(userID uuid.UUID) (*UserInfo, error)
	UpdateUserInfo(userID uuid.UUID, userInfo *UpdateUserInfoRequest) error
	GetUserStats(userID uuid.UUID) (*UserStats, error)
	GetUserAchievements(userID uuid.UUID) ([]UserAchievements, error)
	GetUserBadges(userID uuid.UUID) ([]UserBadges, error)
}
