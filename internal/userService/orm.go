package userservice

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key;"`
	Email    string    `gorm:"type:varchar(255);unique;not null"`
	Password string    `gorm:"type:varchar(255);not null"`
	Role     string    `gorm:"type:varchar(255);not null"`
	Location string    `gorm:"type:varchar(255);not null"`
}

type UpdateUserInfo struct {
	FirstName string `gorm:"type:varchar(255);not null"`
	LastName  string `gorm:"type:varchar(255);not null"`
	Bio       string `gorm:"type:text;not null"`
	Location  string `gorm:"type:varchar(255);not null"`
	Phone     string `gorm:"type:varchar(255);not null"`
}

type UserInfo struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key"`
	UserID    uuid.UUID `gorm:"type:uuid;not null;foreignKey:UserID;references:ID"`
	FirstName string    `gorm:"type:varchar(255);not null"`
	LastName  string    `gorm:"type:varchar(255);not null"`
	Avatar    string    `gorm:"type:varchar(255);not null"`
	Bio       string    `gorm:"type:text;not null"`
	Location  string    `gorm:"type:varchar(255);not null"`
	Phone     string    `gorm:"type:varchar(255);not null"`
}

type UserStats struct {
	ID                uuid.UUID `gorm:"type:uuid;primary_key;"`
	UserID            uuid.UUID `gorm:"type:uuid;not null;foreignKey:UserID;references:ID"`
	CoursesCompleted  int       `gorm:"type:int;not null"`
	CoursesInProgress int       `gorm:"type:int;not null"`
	Followers         int       `gorm:"type:int;not null"`
	Following         int       `gorm:"type:int;not null"`
	Level             int       `gorm:"type:int;not null"`
	XP                int       `gorm:"type:int;not null"`
	NextLevelXP       int       `gorm:"type:int;not null"`
}

type UserAchievements struct {
	ID              uuid.UUID `gorm:"type:uuid;primary_key;"`
	UserID          uuid.UUID `gorm:"type:uuid;not null;foreignKey:UserID;references:ID"`
	AchievementName string    `gorm:"type:varchar(255);not null"`
}

type UserBadges struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	UserID    uuid.UUID `gorm:"type:uuid;not null;foreignKey:UserID;references:ID"`
	BadgeName string    `gorm:"type:varchar(255);not null"`
}
