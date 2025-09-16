package repositories

import (
	"github.com/IbadT/tutor_app_back.git/internal/domain/lessons"
	"gorm.io/gorm"
)

type LessonsRepository struct {
	db *gorm.DB
}

func NewLessonsRepository(db *gorm.DB) *LessonsRepository {
	return &LessonsRepository{db: db}
}

func (r *LessonsRepository) GetLessons() ([]lessons.Lesson, error) {
	var lessons []lessons.Lesson
	if err := r.db.Find(&lessons).Error; err != nil {
		return nil, err
	}
	return lessons, nil
}

func (r *LessonsRepository) CreateLesson(lesson *lessons.Lesson) error {
	return r.db.Create(lesson).Error
}
