package repositories

import (
	"github.com/IbadT/tutor_app_back.git/internal/domain/shared"
	"github.com/IbadT/tutor_app_back.git/internal/domain/user"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// userRepository implements the user.Repository interface
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new user repository
func NewUserRepository(db *gorm.DB) user.Repository {
	return &userRepository{db: db}
}

// GetByID retrieves a user by ID
func (r *userRepository) GetByID(id uuid.UUID) (*shared.User, error) {
	var u shared.User
	err := r.db.Where("id = ?", id).First(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

// GetByEmail retrieves a user by email
func (r *userRepository) GetByEmail(email string) (*shared.User, error) {
	var u shared.User
	err := r.db.Where("email = ?", email).First(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

// Create creates a new user
func (r *userRepository) Create(u *shared.User) error {
	return r.db.Create(u).Error
}

// Update updates an existing user
func (r *userRepository) Update(u *shared.User) error {
	return r.db.Save(u).Error
}

// Delete deletes a user by ID
func (r *userRepository) Delete(id uuid.UUID) error {
	return r.db.Where("id = ?", id).Delete(&shared.User{}).Error
}

// GetUserInfo retrieves user information by user ID
func (r *userRepository) GetUserInfo(userID uuid.UUID) (*user.UserInfo, error) {
	var userInfo user.UserInfo
	err := r.db.Where("user_id = ?", userID).First(&userInfo).Error
	if err != nil {
		return nil, err
	}
	return &userInfo, nil
}

// CreateUserInfo creates new user information
func (r *userRepository) CreateUserInfo(userInfo *user.UserInfo) error {
	return r.db.Create(userInfo).Error
}

// UpdateUserInfo updates user information
func (r *userRepository) UpdateUserInfo(userID uuid.UUID, userInfo *user.UpdateUserInfoRequest) error {
	return r.db.Model(&user.UserInfo{}).Where("user_id = ?", userID).Updates(userInfo).Error
}

// DeleteUserInfo deletes user information
func (r *userRepository) DeleteUserInfo(userID uuid.UUID) error {
	return r.db.Where("user_id = ?", userID).Delete(&user.UserInfo{}).Error
}

// GetUserStats retrieves user statistics
func (r *userRepository) GetUserStats(userID uuid.UUID) (*user.UserStats, error) {
	var stats user.UserStats
	err := r.db.Where("user_id = ?", userID).First(&stats).Error
	if err != nil {
		return nil, err
	}
	return &stats, nil
}

// CreateUserStats creates new user statistics
func (r *userRepository) CreateUserStats(stats *user.UserStats) error {
	return r.db.Create(stats).Error
}

// UpdateUserStats updates user statistics
func (r *userRepository) UpdateUserStats(stats *user.UserStats) error {
	return r.db.Save(stats).Error
}

// GetUserAchievements retrieves user achievements
func (r *userRepository) GetUserAchievements(userID uuid.UUID) ([]user.UserAchievements, error) {
	var achievements []user.UserAchievements
	err := r.db.Where("user_id = ?", userID).Find(&achievements).Error
	if err != nil {
		return nil, err
	}
	return achievements, nil
}

// CreateUserAchievement creates a new user achievement
func (r *userRepository) CreateUserAchievement(achievement *user.UserAchievements) error {
	return r.db.Create(achievement).Error
}

// DeleteUserAchievement deletes a user achievement
func (r *userRepository) DeleteUserAchievement(id uuid.UUID) error {
	return r.db.Where("id = ?", id).Delete(&user.UserAchievements{}).Error
}

// GetUserBadges retrieves user badges
func (r *userRepository) GetUserBadges(userID uuid.UUID) ([]user.UserBadges, error) {
	var badges []user.UserBadges
	err := r.db.Where("user_id = ?", userID).Find(&badges).Error
	if err != nil {
		return nil, err
	}
	return badges, nil
}

// CreateUserBadge creates a new user badge
func (r *userRepository) CreateUserBadge(badge *user.UserBadges) error {
	return r.db.Create(badge).Error
}

// DeleteUserBadge deletes a user badge
func (r *userRepository) DeleteUserBadge(id uuid.UUID) error {
	return r.db.Where("id = ?", id).Delete(&user.UserBadges{}).Error
}

func (r *userRepository) UpdateUserPassword(userID uuid.UUID, password string) error {
	return r.db.Model(&shared.User{}).Where("id = ?", userID).Update("password", password).Error
}

func (r *userRepository) UpdateStudentStatus(replacerID, studentID uuid.UUID, status user.BooleanUpdateRequest) error {
	updates := make(map[string]interface{})

	if status.IsActive != nil {
		updates["is_active"] = *status.IsActive
	}
	if status.IsVerified != nil {
		updates["is_verified"] = *status.IsVerified
	}

	if len(updates) == 0 {
		return nil // Нет полей для обновления
	}

	return r.db.Model(&shared.User{}).Where("id = ?", studentID).Updates(updates).Error
}
