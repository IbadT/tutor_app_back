package userservice

import "github.com/google/uuid"

type UserService interface {
	GetUsersInfo(userID uuid.UUID) (UserInfo, error)
	UpdateUserInfo(userID uuid.UUID, userInfo UpdateUserInfo) error
}

type userService struct {
	userRepo UserRepository
}

func NewUserService(userRepo UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) GetUsersInfo(userID uuid.UUID) (UserInfo, error) {
	return s.userRepo.GetUsersInfo(userID)
}

func (s *userService) UpdateUserInfo(userID uuid.UUID, userInfo UpdateUserInfo) error {
	return s.userRepo.UpdateUserInfo(userID, userInfo)
}
