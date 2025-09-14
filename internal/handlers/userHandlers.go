package handlers

import (
	"context"

	userservice "github.com/IbadT/tutor_app_back.git/internal/userService"
	"github.com/IbadT/tutor_app_back.git/internal/web/users"
	"github.com/google/uuid"
)

type UserHandlers struct {
	service userservice.UserService
}

func NewUserHandlers(service userservice.UserService) *UserHandlers {
	return &UserHandlers{service: service}
}

// GET /users/profile
func (h *UserHandlers) GetUsersProfileId(ctx context.Context, request users.GetUsersProfileIdRequestObject) (users.GetUsersProfileIdResponseObject, error) {
	userIDStr := request.Id
	userID, err := uuid.Parse(userIDStr.String())
	if err != nil {
		return nil, err
	}
	userInfo, err := h.service.GetUsersInfo(userID)
	if err != nil {
		return nil, err
	}
	return users.GetUsersProfileId200JSONResponse{
		Id:        &userInfo.ID,
		FirstName: &userInfo.FirstName,
		LastName:  &userInfo.LastName,
		Bio:       &userInfo.Bio,
		Location:  &userInfo.Location,
		Phone:     &userInfo.Phone,
	}, nil
}

func (h *UserHandlers) PatchUsersProfileId(ctx context.Context, request users.PatchUsersProfileIdRequestObject) (users.PatchUsersProfileIdResponseObject, error) {
	userIDStr := request.Id
	userID, err := uuid.Parse(userIDStr.String())
	if err != nil {
		return nil, err
	}
	body := request.Body

	userInfo := userservice.UpdateUserInfo{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Bio:       body.Bio,
		Location:  body.Location,
		Phone:     body.Phone,
	}

	err = h.service.UpdateUserInfo(userID, userInfo)
	if err != nil {
		return nil, err
	}

	return users.PatchUsersProfileId200JSONResponse{
		Code:    &[]int{200}[0],
		Message: &[]string{"User profile updated successfully"}[0],
	}, nil
}
