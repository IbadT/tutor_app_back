package courses

import "github.com/google/uuid"

type Service interface {
	GetCourses() ([]Course, error)
	GetCourseByID(id uuid.UUID) (*Course, error)
}

type service struct {
	courseRepo Repository
}

func NewService(courseRepo Repository) Service {
	return &service{courseRepo: courseRepo}
}

func (s *service) GetCourses() ([]Course, error) {
	return s.courseRepo.GetCourses()
}

func (s *service) GetCourseByID(id uuid.UUID) (*Course, error) {
	return s.courseRepo.GetCourseByID(id)
}
