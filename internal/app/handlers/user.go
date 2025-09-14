package handlers

import (
	"context"

	"github.com/IbadT/tutor_app_back.git/internal/domain/shared"
	"github.com/IbadT/tutor_app_back.git/internal/domain/user"
	web_users "github.com/IbadT/tutor_app_back.git/internal/web/users"
	"github.com/google/uuid"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// UserHandler handles user-related requests
type UserHandler struct {
	userService user.Service
}

// NewUserHandler creates a new user handler
func NewUserHandler(userService user.Service) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// GetUsersProfileId handles GET /users/profile/{id}
func (h *UserHandler) GetUsersProfileId(ctx context.Context, request web_users.GetUsersProfileIdRequestObject) (web_users.GetUsersProfileIdResponseObject, error) {
	userID := uuid.UUID(request.Id)

	userInfo, err := h.userService.GetUserInfo(userID)
	if err != nil {
		return h.handleGetUserProfileError(err)
	}

	return web_users.GetUsersProfileId200JSONResponse{
		Id:        (*openapi_types.UUID)(&userInfo.ID),
		FirstName: &userInfo.FirstName,
		LastName:  &userInfo.LastName,
		Avatar:    &userInfo.Avatar,
		Bio:       &userInfo.Bio,
		Location:  &userInfo.Location,
		Phone:     &userInfo.Phone,
	}, nil
}

// PatchUsersProfileId handles PATCH /users/profile/{id}
func (h *UserHandler) PatchUsersProfileId(ctx context.Context, request web_users.PatchUsersProfileIdRequestObject) (web_users.PatchUsersProfileIdResponseObject, error) {
	userID := uuid.UUID(request.Id)

	body := request.Body
	updateRequest := &user.UpdateUserInfoRequest{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Bio:       body.Bio,
		Location:  body.Location,
		Phone:     body.Phone,
	}

	err := h.userService.UpdateUserInfo(userID, updateRequest)
	if err != nil {
		return h.handleUpdateUserProfileError(err)
	}

	return web_users.PatchUsersProfileId200JSONResponse{
		Message: func() *string { msg := "User profile updated successfully"; return &msg }(),
	}, nil
}

// Error handling methods
func (h *UserHandler) handleGetUserProfileError(err error) (web_users.GetUsersProfileIdResponseObject, error) {
	if apiErr, ok := err.(*shared.APIError); ok {
		switch apiErr.Code {
		case 401:
			return web_users.GetUsersProfileId401JSONResponse{Code: &apiErr.Code, Message: &apiErr.Message}, nil
		default:
			return web_users.GetUsersProfileId500JSONResponse{Code: &apiErr.Code, Message: &apiErr.Message}, nil
		}
	}
	code := 500
	msg := "Internal server error"
	return web_users.GetUsersProfileId500JSONResponse{Code: &code, Message: &msg}, nil
}

func (h *UserHandler) handleUpdateUserProfileError(err error) (web_users.PatchUsersProfileIdResponseObject, error) {
	if apiErr, ok := err.(*shared.APIError); ok {
		switch apiErr.Code {
		case 400:
			return web_users.PatchUsersProfileId400JSONResponse{Code: &apiErr.Code, Message: &apiErr.Message}, nil
		default:
			return web_users.PatchUsersProfileId500JSONResponse{Code: &apiErr.Code, Message: &apiErr.Message}, nil
		}
	}
	code := 500
	msg := "Internal server error"
	return web_users.PatchUsersProfileId500JSONResponse{Code: &code, Message: &msg}, nil
}

// GetUsersIdStats handles GET /users/{id}/stats
func (h *UserHandler) GetUsersIdStats(ctx context.Context, request web_users.GetUsersIdStatsRequestObject) (web_users.GetUsersIdStatsResponseObject, error) {
	userID := uuid.UUID(request.Id)

	stats, err := h.userService.GetUserStats(userID)
	if err != nil {
		return h.handleGetUserStatsError(err)
	}

	return web_users.GetUsersIdStats200JSONResponse{
		Id:                (*openapi_types.UUID)(&stats.ID),
		UserId:            (*openapi_types.UUID)(&stats.UserID),
		CoursesCompleted:  &stats.CoursesCompleted,
		CoursesInProgress: &stats.CoursesInProgress,
		Followers:         &stats.Followers,
		Following:         &stats.Following,
		Level:             &stats.Level,
		Xp:                &stats.XP,
		NextLevelXp:       &stats.NextLevelXP,
	}, nil
}

// GetUsersIdAchievements handles GET /users/{id}/achievements
func (h *UserHandler) GetUsersIdAchievements(ctx context.Context, request web_users.GetUsersIdAchievementsRequestObject) (web_users.GetUsersIdAchievementsResponseObject, error) {
	userID := uuid.UUID(request.Id)

	achievements, err := h.userService.GetUserAchievements(userID)
	if err != nil {
		return h.handleGetUserAchievementsError(err)
	}

	// Convert to response format
	var responseAchievements []web_users.UserAchievement
	for _, achievement := range achievements {
		responseAchievements = append(responseAchievements, web_users.UserAchievement{
			Id:              (*openapi_types.UUID)(&achievement.ID),
			UserId:          (*openapi_types.UUID)(&achievement.UserID),
			AchievementName: &achievement.AchievementName,
		})
	}

	return web_users.GetUsersIdAchievements200JSONResponse(responseAchievements), nil
}

// GetUsersIdBadges handles GET /users/{id}/badges
func (h *UserHandler) GetUsersIdBadges(ctx context.Context, request web_users.GetUsersIdBadgesRequestObject) (web_users.GetUsersIdBadgesResponseObject, error) {
	userID := uuid.UUID(request.Id)

	badges, err := h.userService.GetUserBadges(userID)
	if err != nil {
		return h.handleGetUserBadgesError(err)
	}

	// Convert to response format
	var responseBadges []web_users.UserBadge
	for _, badge := range badges {
		responseBadges = append(responseBadges, web_users.UserBadge{
			Id:        (*openapi_types.UUID)(&badge.ID),
			UserId:    (*openapi_types.UUID)(&badge.UserID),
			BadgeName: &badge.BadgeName,
		})
	}

	return web_users.GetUsersIdBadges200JSONResponse(responseBadges), nil
}

// Error handling methods for new endpoints
func (h *UserHandler) handleGetUserStatsError(err error) (web_users.GetUsersIdStatsResponseObject, error) {
	if apiErr, ok := err.(*shared.APIError); ok {
		switch apiErr.Code {
		case 401:
			return web_users.GetUsersIdStats401JSONResponse{Code: &apiErr.Code, Message: &apiErr.Message}, nil
		default:
			return web_users.GetUsersIdStats500JSONResponse{Code: &apiErr.Code, Message: &apiErr.Message}, nil
		}
	}
	code := 500
	msg := "Internal server error"
	return web_users.GetUsersIdStats500JSONResponse{Code: &code, Message: &msg}, nil
}

func (h *UserHandler) handleGetUserAchievementsError(err error) (web_users.GetUsersIdAchievementsResponseObject, error) {
	if apiErr, ok := err.(*shared.APIError); ok {
		switch apiErr.Code {
		case 401:
			return web_users.GetUsersIdAchievements401JSONResponse{Code: &apiErr.Code, Message: &apiErr.Message}, nil
		default:
			return web_users.GetUsersIdAchievements500JSONResponse{Code: &apiErr.Code, Message: &apiErr.Message}, nil
		}
	}
	code := 500
	msg := "Internal server error"
	return web_users.GetUsersIdAchievements500JSONResponse{Code: &code, Message: &msg}, nil
}

func (h *UserHandler) handleGetUserBadgesError(err error) (web_users.GetUsersIdBadgesResponseObject, error) {
	if apiErr, ok := err.(*shared.APIError); ok {
		switch apiErr.Code {
		case 401:
			return web_users.GetUsersIdBadges401JSONResponse{Code: &apiErr.Code, Message: &apiErr.Message}, nil
		default:
			return web_users.GetUsersIdBadges500JSONResponse{Code: &apiErr.Code, Message: &apiErr.Message}, nil
		}
	}
	code := 500
	msg := "Internal server error"
	return web_users.GetUsersIdBadges500JSONResponse{Code: &code, Message: &msg}, nil
}
