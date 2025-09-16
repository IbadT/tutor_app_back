package repositories

import (
	"github.com/IbadT/tutor_app_back.git/internal/domain/courses"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type courseRepository struct {
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) courses.Repository {
	return &courseRepository{db: db}
}

func (r *courseRepository) GetCourses() ([]courses.Course, error) {
	var courses []courses.Course
	if err := r.db.
		Preload("Student").
		Preload("Tutor").
		Preload("Category").
		Find(&courses).Error; err != nil {
		return nil, err
	}
	return courses, nil
}

func (r *courseRepository) GetCourseByID(id uuid.UUID) (*courses.Course, error) {
	var course courses.Course
	if err := r.db.
		Preload("Student").
		Preload("Tutor").
		Preload("Category").
		Where("id = ?", id).
		First(&course).Error; err != nil {
		return nil, err
	}
	return &course, nil
}
