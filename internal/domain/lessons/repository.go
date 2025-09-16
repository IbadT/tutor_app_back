package lessons

type Repository interface {
	GetLessons() ([]Lesson, error)
	CreateLesson(lesson *Lesson) error
}
