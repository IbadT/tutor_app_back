package handlers

import (
	"context"
	"time"

	"github.com/IbadT/tutor_app_back.git/internal/domain/lessons"
	"github.com/IbadT/tutor_app_back.git/internal/domain/shared"
	web_lessons "github.com/IbadT/tutor_app_back.git/internal/web/lessons"
	"github.com/google/uuid"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

type LessonsHandler struct {
	lessonsService lessons.Service
}

func NewLessonsHandler(lessonsService lessons.Service) *LessonsHandler {
	return &LessonsHandler{lessonsService: lessonsService}
}

func (h *LessonsHandler) GetLessons(ctx context.Context, request web_lessons.GetLessonsRequestObject) (web_lessons.GetLessonsResponseObject, error) {
	lessons, err := h.lessonsService.GetLessons()
	if err != nil {
		return h.handleGetLessonsError(err)
	}

	var responseLessons []web_lessons.Lesson
	for _, lesson := range lessons {
		responseLessons = append(responseLessons, web_lessons.Lesson{
			Id:          (openapi_types.UUID)(lesson.ID),
			Title:       lesson.Title,
			Description: lesson.Description,
			VideoUrl:    lesson.VideoURL,
			Duration:    lesson.Duration,
			CreatedAt:   lesson.CreatedAt,
			UpdatedAt:   lesson.UpdatedAt,
		})
	}
	return web_lessons.GetLessons200JSONResponse(responseLessons), nil
}

func (h *LessonsHandler) PostLessonsCreaterId(ctx context.Context, request web_lessons.PostLessonsCreaterIdRequestObject) (web_lessons.PostLessonsCreaterIdResponseObject, error) {
	creater_id := uuid.UUID(request.CreaterId)
	lesson := lessons.Lesson{
		ID:          uuid.New(),
		Title:       request.Body.Title,
		Description: request.Body.Description,
		VideoURL:    request.Body.VideoUrl,
		Duration:    request.Body.Duration,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		CourseID:    request.Body.CourseId,
	}
	err := h.lessonsService.CreateLesson(creater_id, &lesson)
	if err != nil {
		return h.handleCreateLessonError(err)
	}
	return web_lessons.PostLessonsCreaterId200JSONResponse{
		Message: func() *string { msg := "Lesson created successfully"; return &msg }(),
	}, nil
}

func (h *LessonsHandler) handleCreateLessonError(err error) (web_lessons.PostLessonsCreaterIdResponseObject, error) {
	if apiErr, ok := err.(*shared.APIError); ok {
		return web_lessons.PostLessonsCreaterId401JSONResponse{Code: &apiErr.Code, Message: &apiErr.Message}, nil
	}
	code := 500
	msg := "Internal server error"
	return web_lessons.PostLessonsCreaterId500JSONResponse{Code: &code, Message: &msg}, nil
}

func (h *LessonsHandler) handleGetLessonsError(err error) (web_lessons.GetLessonsResponseObject, error) {
	if apiErr, ok := err.(*shared.APIError); ok {
		switch apiErr.Code {
		case 401:
			return web_lessons.GetLessons401JSONResponse{Code: &apiErr.Code, Message: &apiErr.Message}, nil
		default:
			return web_lessons.GetLessons500JSONResponse{Code: &apiErr.Code, Message: &apiErr.Message}, nil
		}
	}
	code := 500
	msg := "Internal server error"
	return web_lessons.GetLessons500JSONResponse{Code: &code, Message: &msg}, nil
}
