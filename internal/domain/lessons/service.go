package lessons

import (
	"github.com/IbadT/tutor_app_back.git/internal/domain/shared"
	"github.com/IbadT/tutor_app_back.git/internal/domain/user"
	"github.com/google/uuid"
)

type Service interface {
	GetLessons() ([]Lesson, error)
	CreateLesson(creater_id uuid.UUID, lesson *Lesson) error
}

type service struct {
	lessonsRepo Repository
	userRepo    user.Repository
}

func NewService(lessonsRepo Repository, userRepo user.Repository) Service {
	return &service{
		lessonsRepo: lessonsRepo,
		userRepo:    userRepo,
	}
}

func (s *service) GetLessons() ([]Lesson, error) {
	return s.lessonsRepo.GetLessons()
}

func (s *service) CreateLesson(creater_id uuid.UUID, lesson *Lesson) error {
	// проверяем, что создатель - это админ
	user, err := s.userRepo.GetByID(creater_id)
	if err != nil {
		return err
	}
	if user.Role != "admin" {
		return shared.ErrForbidden
	}

	return s.lessonsRepo.CreateLesson(lesson)
}
