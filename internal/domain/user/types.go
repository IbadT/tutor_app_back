package user

import "github.com/google/uuid"

// User represents the core user entity
type User struct {
	ID       uuid.UUID `json:"id" gorm:"type:uuid;primary_key;"`
	Email    string    `json:"email" gorm:"type:varchar(255);unique;not null"`
	Password string    `json:"-" gorm:"type:varchar(255);not null"`
	Role     string    `json:"role" gorm:"type:varchar(255);not null"`
	Location string    `json:"location" gorm:"type:varchar(255);not null"`
}

// UserInfo represents additional user information
type UserInfo struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;not null;foreignKey:UserID;references:ID"`
	FirstName string    `json:"first_name" gorm:"type:varchar(255);not null"`
	LastName  string    `json:"last_name" gorm:"type:varchar(255);not null"`
	Avatar    string    `json:"avatar" gorm:"type:varchar(255);not null"`
	Bio       string    `json:"bio" gorm:"type:text;not null"`
	Location  string    `json:"location" gorm:"type:varchar(255);not null"`
	Phone     string    `json:"phone" gorm:"type:varchar(255);not null"`
}

// UpdateUserInfoRequest represents the request to update user info
type UpdateUserInfoRequest struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Bio       string `json:"bio"`
	Location  string `json:"location" validate:"required"`
	Phone     string `json:"phone"`
}

// UserStats represents user statistics
type UserStats struct {
	ID                uuid.UUID `json:"id" gorm:"type:uuid;primary_key;"`
	UserID            uuid.UUID `json:"user_id" gorm:"type:uuid;not null;foreignKey:UserID;references:ID"`
	CoursesCompleted  int       `json:"courses_completed" gorm:"type:int;not null"`
	CoursesInProgress int       `json:"courses_in_progress" gorm:"type:int;not null"`
	Followers         int       `json:"followers" gorm:"type:int;not null"`
	Following         int       `json:"following" gorm:"type:int;not null"`
	Level             int       `json:"level" gorm:"type:int;not null"`
	XP                int       `json:"xp" gorm:"type:int;not null"`
	NextLevelXP       int       `json:"next_level_xp" gorm:"type:int;not null"`
}

// UserAchievements represents user achievements
type UserAchievements struct {
	ID              uuid.UUID `json:"id" gorm:"type:uuid;primary_key;"`
	UserID          uuid.UUID `json:"user_id" gorm:"type:uuid;not null;foreignKey:UserID;references:ID"`
	AchievementName string    `json:"achievement_name" gorm:"type:varchar(255);not null"`
}

// UserBadges represents user badges
type UserBadges struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;not null;foreignKey:UserID;references:ID"`
	BadgeName string    `json:"badge_name" gorm:"type:varchar(255);not null"`
}
