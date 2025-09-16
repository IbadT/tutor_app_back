package user

import (
	"github.com/IbadT/tutor_app_back.git/internal/domain/shared"
	"github.com/google/uuid"
)

// Service defines the interface for user business logic
type Service interface {
	GetUserInfo(userID uuid.UUID) (*UserInfo, error)
	UpdateUserInfo(userID uuid.UUID, userInfo *UpdateUserInfoRequest) error
	GetUserStats(userID uuid.UUID) (*UserStats, error)
	GetUserAchievements(userID uuid.UUID) ([]UserAchievements, error)
	GetUserBadges(userID uuid.UUID) ([]UserBadges, error)
	UpdateUserPassword(userID uuid.UUID, passwordsRequest *UpdateUserPasswordRequest) error
	UpdateStudentStatus(replacerID, studentID uuid.UUID, status BooleanUpdateRequest) error
}

// service implements the user business logic
type service struct {
	userRepo     Repository
	passwordHash shared.PasswordHasher
}

// NewService creates a new user service
func NewService(userRepo Repository, passwordHash shared.PasswordHasher) Service {
	return &service{
		userRepo:     userRepo,
		passwordHash: passwordHash,
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

// UpdateUserPassword updates user password
func (s *service) UpdateUserPassword(userID uuid.UUID, passwordsRequest *UpdateUserPasswordRequest) error {
	if userID == uuid.Nil {
		return shared.ErrInvalidInput
	}

	if passwordsRequest == nil {
		return shared.ErrMissingFields
	}

	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		return shared.ErrDatabaseError
	}

	if s.passwordHash.ComparePassword(passwordsRequest.CurrentPassword, user.Password) {
		return shared.ErrInvalidCredentials
	}

	passwordHash, err := s.passwordHash.HashPassword(passwordsRequest.NewPassword)
	if err != nil {
		return shared.ErrDatabaseError
	}

	err = s.userRepo.UpdateUserPassword(userID, passwordHash)
	if err != nil {
		return shared.ErrDatabaseError
	}

	return nil
}

func (s *service) UpdateStudentStatus(replacerID, studentID uuid.UUID, status BooleanUpdateRequest) error {
	if replacerID == uuid.Nil || studentID == uuid.Nil {
		return shared.ErrInvalidInput
	}

	if status.IsActive == nil && status.IsVerified == nil {
		return shared.ErrMissingFields
	}

	err := s.userRepo.UpdateStudentStatus(replacerID, studentID, status)
	if err != nil {
		return shared.ErrDatabaseError
	}

	return nil
}
