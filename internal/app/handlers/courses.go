package handlers

import (
	"context"

	"github.com/IbadT/tutor_app_back.git/internal/domain/courses"
	"github.com/IbadT/tutor_app_back.git/internal/domain/shared"
	web_courses "github.com/IbadT/tutor_app_back.git/internal/web/courses"
	"github.com/google/uuid"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

type CourseHandler struct {
	courseService courses.Service
}

func NewCourseHandler(courseService courses.Service) *CourseHandler {
	return &CourseHandler{courseService: courseService}
}

// GetCourses handles GET /courses
func (h *CourseHandler) GetCourses(ctx context.Context, request web_courses.GetCoursesRequestObject) (web_courses.GetCoursesResponseObject, error) {
	courses, err := h.courseService.GetCourses()
	if err != nil {
		return h.handleGetCoursesError(err)
	}

	// Convert domain courses to web response format
	var responseCourses []web_courses.Course
	for _, course := range courses {
		responseCourse := web_courses.Course{
			Id:               (*openapi_types.UUID)(&course.ID),
			Title:            &course.Title,
			Description:      &course.Description,
			Progress:         &course.Progress,
			TotalLessons:     &course.TotalLessons,
			CompletedLessons: &course.CompletedLessons,
			Duration:         &course.Duration,
			StudentsCount:    &course.StudentsCount,
			Rating:           &course.Rating,
			CategoryId:       (*openapi_types.UUID)(&course.CategoryID),
			CreatedAt:        &course.CreatedAt,
			UpdatedAt:        &course.UpdatedAt,
		}

		// Add related data if available
		if course.Student != nil {
			responseCourse.StudentId = (*openapi_types.UUID)(&course.Student.ID)
		}
		if course.Tutor != nil {
			responseCourse.TutorId = (*openapi_types.UUID)(&course.Tutor.ID)
		}

		responseCourses = append(responseCourses, responseCourse)
	}

	return web_courses.GetCourses200JSONResponse(responseCourses), nil
}

// GetCourseByID handles GET /courses/{course_id}
func (h *CourseHandler) GetCoursesCourseId(ctx context.Context, request web_courses.GetCoursesCourseIdRequestObject) (web_courses.GetCoursesCourseIdResponseObject, error) {
	courseID := uuid.UUID(request.CourseId)

	course, err := h.courseService.GetCourseByID(courseID)
	if err != nil {
		return h.handleGetCourseByIDError(err)
	}

	responseCourse := web_courses.Course{
		Id:               (*openapi_types.UUID)(&course.ID),
		Title:            &course.Title,
		Description:      &course.Description,
		Progress:         &course.Progress,
		TotalLessons:     &course.TotalLessons,
		CompletedLessons: &course.CompletedLessons,
		Duration:         &course.Duration,
		StudentsCount:    &course.StudentsCount,
		Rating:           &course.Rating,
		CategoryId:       (*openapi_types.UUID)(&course.CategoryID),
		CreatedAt:        &course.CreatedAt,
		UpdatedAt:        &course.UpdatedAt,
	}

	// Add related data if available
	if course.Student != nil {
		responseCourse.StudentId = (*openapi_types.UUID)(&course.Student.ID)
	}
	if course.Tutor != nil {
		responseCourse.TutorId = (*openapi_types.UUID)(&course.Tutor.ID)
	}

	return web_courses.GetCoursesCourseId200JSONResponse(responseCourse), nil
}

func (h *CourseHandler) handleGetCourseByIDError(err error) (web_courses.GetCoursesCourseIdResponseObject, error) {
	if apiErr, ok := err.(*shared.APIError); ok {
		switch apiErr.Code {
		case 401:
			return web_courses.GetCoursesCourseId401JSONResponse{Code: &apiErr.Code, Message: &apiErr.Message}, nil
		default:
			return web_courses.GetCoursesCourseId500JSONResponse{Code: &apiErr.Code, Message: &apiErr.Message}, nil
		}
	}
	code := 500
	msg := "Internal server error"
	return web_courses.GetCoursesCourseId500JSONResponse{Code: &code, Message: &msg}, nil
}

func (h *CourseHandler) handleGetCoursesError(err error) (web_courses.GetCoursesResponseObject, error) {
	if apiErr, ok := err.(*shared.APIError); ok {
		switch apiErr.Code {
		case 401:
			return web_courses.GetCourses401JSONResponse{Code: &apiErr.Code, Message: &apiErr.Message}, nil
		default:
			return web_courses.GetCourses500JSONResponse{Code: &apiErr.Code, Message: &apiErr.Message}, nil
		}
	}
	code := 500
	msg := "Internal server error"
	return web_courses.GetCourses500JSONResponse{Code: &code, Message: &msg}, nil
}
