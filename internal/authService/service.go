package authservice

import (
	"errors"
	"slices"

	apierrors "github.com/IbadT/tutor_app_back.git/internal/errors"
	userservice "github.com/IbadT/tutor_app_back.git/internal/userService"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthService interface {
	Login(loginRequest LoginRequest) (LoginResponse, error)
	Register(user RegisterRequest) (LoginResponse, error)
	RefreshToken(refreshToken string) (LoginResponse, error)
}

type authService struct {
	authRepo AuthRepository
	userRepo userservice.UserRepository
}

func NewAuthService(authRepo AuthRepository, userRepo userservice.UserRepository) AuthService {
	return &authService{authRepo: authRepo, userRepo: userRepo}
}

func (s *authService) Login(loginRequest LoginRequest) (LoginResponse, error) {
	// Validate input
	if loginRequest.Email == "" || loginRequest.Password == "" {
		return LoginResponse{}, apierrors.ErrMissingFields
	}

	// Check if user exists
	userExists, err := s.authRepo.CheckUserExists(loginRequest.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return LoginResponse{}, apierrors.ErrInvalidCredentials
		}
		return LoginResponse{}, apierrors.ErrDatabaseError
	}

	// Verify password
	if !ComparePassword(loginRequest.Password, userExists.Password) {
		return LoginResponse{}, apierrors.ErrInvalidCredentials
	}

	// Generate tokens
	tokens, err := GenerateToken(userExists.ID, userExists.Role)
	if err != nil {
		return LoginResponse{}, apierrors.ErrTokenGeneration
	}

	return tokens, nil
}

func (s *authService) Register(userRequest RegisterRequest) (LoginResponse, error) {
	// Validate input
	if userRequest.Email == "" || userRequest.Password == "" ||
		userRequest.FirstName == "" || userRequest.LastName == "" ||
		userRequest.Role == "" || userRequest.Location == "" {
		return LoginResponse{}, apierrors.ErrMissingFields
	}

	// Validate role
	validRoles := []string{"student", "tutor", "admin"}
	if !slices.Contains(validRoles, userRequest.Role) {
		return LoginResponse{}, apierrors.NewAPIError(400, "Invalid role. Must be one of: student, tutor, admin")
	}

	// Check if user already exists
	_, err := s.authRepo.CheckUserExists(userRequest.Email)
	if err == nil {
		return LoginResponse{}, apierrors.ErrUserAlreadyExists
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return LoginResponse{}, apierrors.ErrDatabaseError
	}

	// Hash password
	hashedPassword, err := HashPassword(userRequest.Password)
	if err != nil {
		return LoginResponse{}, apierrors.ErrInternalServer
	}

	// Create user
	user := userservice.User{
		ID:       uuid.New(),
		Email:    userRequest.Email,
		Password: hashedPassword,
		Role:     userRequest.Role,
		Location: userRequest.Location,
	}

	// Save user to database
	if err := s.authRepo.Register(user); err != nil {
		return LoginResponse{}, apierrors.ErrDatabaseError
	}

	// Create user info
	userInfo := userservice.UserInfo{
		ID:        uuid.New(),
		UserID:    user.ID,
		FirstName: userRequest.FirstName,
		LastName:  userRequest.LastName,
		Location:  userRequest.Location,
	}

	if err := s.userRepo.CreateUserInfo(userInfo); err != nil {
		s.userRepo.DeleteUser(user.ID)
		return LoginResponse{}, apierrors.ErrDatabaseError
	}

	// Generate tokens
	tokens, err := GenerateToken(user.ID, user.Role)
	if err != nil {
		return LoginResponse{}, apierrors.ErrTokenGeneration
	}

	return tokens, nil
}

func (s *authService) RefreshToken(refreshToken string) (LoginResponse, error) {
	// Validate input
	if refreshToken == "" {
		return LoginResponse{}, apierrors.ErrMissingFields
	}

	// Parse and validate refresh token
	claims, err := ParseToken(refreshToken)
	if err != nil {
		return LoginResponse{}, apierrors.ErrInvalidCredentials
	}

	// Extract user ID and role from token
	userIDStr, ok := claims["sub"].(string)
	if !ok {
		return LoginResponse{}, apierrors.ErrInvalidCredentials
	}

	role, ok := claims["role"].(string)
	if !ok {
		return LoginResponse{}, apierrors.ErrInvalidCredentials
	}

	// Parse user ID
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return LoginResponse{}, apierrors.ErrInvalidCredentials
	}

	// Generate new tokens
	tokens, err := GenerateToken(userID, role)
	if err != nil {
		return LoginResponse{}, apierrors.ErrTokenGeneration
	}

	return tokens, nil
}
