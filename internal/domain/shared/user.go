package shared

import (
	"time"

	"github.com/google/uuid"
)

// User represents a user entity shared across domain layers
type User struct {
	ID         uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Email      string    `json:"email" gorm:"uniqueIndex;not null"`
	Password   string    `json:"-" gorm:"not null"`
	Role       string    `json:"role" gorm:"not null"`
	Location   string    `json:"location" gorm:"not null"`
	IsVerified bool      `json:"is_verified" gorm:"column:is_verified;default:false"`
	IsActive   bool      `json:"is_active" gorm:"column:is_active;default:true"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
