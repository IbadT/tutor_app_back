package user

import (
	"github.com/IbadT/tutor_app_back.git/internal/domain/shared"
	"github.com/google/uuid"
)

// service implements the user business logic
type service struct {
	userRepo Repository
}

// NewService creates a new user service
func NewService(userRepo Repository) Service {
	return &service{
		userRepo: userRepo,
	}
}

// GetUserInfo retrieves user information by user ID
func (s *service) GetUserInfo(userID uuid.UUID) (*UserInfo, error) {
	if userID == uuid.Nil {
		return nil, shared.ErrInvalidInput
	}

	userInfo, err := s.userRepo.GetUserInfo(userID)
	if err != nil {
		return nil, shared.ErrDatabaseError
	}

	return userInfo, nil
}

// UpdateUserInfo updates user information
func (s *service) UpdateUserInfo(userID uuid.UUID, userInfo *UpdateUserInfoRequest) error {
	if userID == uuid.Nil {
		return shared.ErrInvalidInput
	}

	if userInfo == nil {
		return shared.ErrMissingFields
	}

	// Validate required fields
	if userInfo.FirstName == "" || userInfo.LastName == "" || userInfo.Location == "" {
		return shared.ErrMissingFields
	}

	err := s.userRepo.UpdateUserInfo(userID, userInfo)
	if err != nil {
		return shared.ErrDatabaseError
	}

	return nil
}

// GetUserStats retrieves user statistics
func (s *service) GetUserStats(userID uuid.UUID) (*UserStats, error) {
	if userID == uuid.Nil {
		return nil, shared.ErrInvalidInput
	}

	stats, err := s.userRepo.GetUserStats(userID)
	if err != nil {
		return nil, shared.ErrDatabaseError
	}

	return stats, nil
}

// GetUserAchievements retrieves user achievements
func (s *service) GetUserAchievements(userID uuid.UUID) ([]UserAchievements, error) {
	if userID == uuid.Nil {
		return nil, shared.ErrInvalidInput
	}

	achievements, err := s.userRepo.GetUserAchievements(userID)
	if err != nil {
		return nil, shared.ErrDatabaseError
	}

	return achievements, nil
}

// GetUserBadges retrieves user badges
func (s *service) GetUserBadges(userID uuid.UUID) ([]UserBadges, error) {
	if userID == uuid.Nil {
		return nil, shared.ErrInvalidInput
	}

	badges, err := s.userRepo.GetUserBadges(userID)
	if err != nil {
		return nil, shared.ErrDatabaseError
	}

	return badges, nil
}
