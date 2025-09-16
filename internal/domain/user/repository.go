package user

import (
	"github.com/IbadT/tutor_app_back.git/internal/domain/shared"
	"github.com/google/uuid"
)

// Repository defines the interface for user data operations
type Repository interface {
	// User operations
	GetByID(id uuid.UUID) (*shared.User, error)
	GetByEmail(email string) (*shared.User, error)
	Create(user *shared.User) error
	Update(user *shared.User) error
	Delete(id uuid.UUID) error
	UpdateUserPassword(userID uuid.UUID, password string) error

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

	// Student status operations
	UpdateStudentStatus(replacerID, studentID uuid.UUID, status BooleanUpdateRequest) error
}
