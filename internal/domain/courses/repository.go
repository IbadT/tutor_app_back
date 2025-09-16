package courses

import "github.com/google/uuid"

type Repository interface {
	GetCourses() ([]Course, error)
	GetCourseByID(id uuid.UUID) (*Course, error)
}
