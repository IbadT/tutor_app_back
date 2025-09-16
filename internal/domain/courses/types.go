package courses

import (
	"time"

	"github.com/google/uuid"
)

type Pagination struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
	Total int `json:"total"`
}

type Course struct {
	ID               uuid.UUID `json:"id" gorm:"type:uuid;primary_key;"`
	StudentID        uuid.UUID `json:"student_id" gorm:"type:uuid;not null"`
	TutorID          uuid.UUID `json:"tutor_id" gorm:"type:uuid;not null"`
	Title            string    `json:"title" gorm:"not null"`
	Description      string    `json:"description" gorm:"not null"`
	Progress         int       `json:"progress" gorm:"not null"`
	TotalLessons     int       `json:"total_lessons" gorm:"not null"`
	CompletedLessons int       `json:"completed_lessons" gorm:"not null"`
	Duration         string    `json:"duration" gorm:"not null"`
	StudentsCount    int       `json:"students_count" gorm:"not null"`
	Rating           float32   `json:"rating" gorm:"not null"`
	CategoryID       uuid.UUID `json:"category_id" gorm:"type:uuid;not null"`
	CreatedAt        time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt        time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	// Related data
	Student  *User     `json:"student,omitempty" gorm:"foreignKey:StudentID;references:ID"`
	Tutor    *User     `json:"tutor,omitempty" gorm:"foreignKey:TutorID;references:ID"`
	Category *Category `json:"category,omitempty" gorm:"foreignKey:CategoryID;references:ID"`
}

type User struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;"`
	FirstName string    `json:"first_name" gorm:"not null"`
	LastName  string    `json:"last_name" gorm:"not null"`
	Email     string    `json:"email" gorm:"not null"`
	Avatar    string    `json:"avatar"`
}

type Category struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primary_key;"`
	Name        string    `json:"name" gorm:"not null"`
	Description string    `json:"description"`
}

type GetCoursesResponse struct {
	Pagination Pagination `json:"pagination"`
	Courses    []Course   `json:"courses"`
	Duration   string     `json:"duration"`
	Total      int        `json:"total"`
}
