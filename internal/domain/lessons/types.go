package lessons

import (
	"time"

	"github.com/google/uuid"
)

type Lesson struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primary_key;"`
	CourseID    uuid.UUID `json:"course_id" gorm:"type:uuid;not null;foreignKey:CourseID;references:ID"`
	Title       string    `json:"title" gorm:"type:varchar(255);not null"`
	Description string    `json:"description" gorm:"type:text;not null"`
	VideoURL    string    `json:"video_url" gorm:"type:varchar(255);not null"`
	Duration    string    `json:"duration" gorm:"type:varchar(255);not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"type:timestamp;not null"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"type:timestamp;not null"`
}
