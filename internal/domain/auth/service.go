package auth

import (
	"errors"
	"slices"

	"github.com/IbadT/tutor_app_back.git/internal/domain/shared"
	"github.com/IbadT/tutor_app_back.git/internal/domain/user"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Service defines the interface for authentication business logic
type Service interface {
	Login(req *LoginRequest) (*LoginResponse, error)
	Register(req *RegisterRequest) (*LoginResponse, error)
	RefreshToken(refreshToken string) (*LoginResponse, error)
	ValidateToken(tokenString string) (uuid.UUID, string, error)
}

// service implements the authentication business logic
type service struct {
	authRepo     Repository
	userRepo     user.Repository
	tokenGen     TokenGenerator
	passwordHash shared.PasswordHasher
}

// NewService creates a new authentication service
func NewService(
	authRepo Repository,
	userRepo user.Repository,
	tokenGen TokenGenerator,
	passwordHash shared.PasswordHasher,
) Service {
	return &service{
		authRepo:     authRepo,
		userRepo:     userRepo,
		tokenGen:     tokenGen,
		passwordHash: passwordHash,
	}
}

// Login authenticates a user and returns tokens
func (s *service) Login(req *LoginRequest) (*LoginResponse, error) {
	// Validate input
	if req.Email == "" || req.Password == "" {
		return nil, shared.ErrMissingFields
	}

	// Get user by email
	user, err := s.authRepo.GetUserByEmail(req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, shared.ErrInvalidCredentials
		}
		return nil, shared.ErrDatabaseError
	}

	// Verify password
	if !s.passwordHash.ComparePassword(req.Password, user.Password) {
		return nil, shared.ErrInvalidCredentials
	}

	// Generate tokens
	tokens, err := s.tokenGen.GenerateToken(user.ID, user.Role)
	if err != nil {
		return nil, shared.ErrTokenGeneration
	}

	return tokens, nil
}

// Register creates a new user and returns tokens
func (s *service) Register(req *RegisterRequest) (*LoginResponse, error) {
	// Validate input
	if req.Email == "" || req.Password == "" ||
		req.FirstName == "" || req.LastName == "" ||
		req.Role == "" || req.Location == "" {
		return nil, shared.ErrMissingFields
	}

	// Validate role
	validRoles := []string{"student", "tutor", "admin"}
	if !slices.Contains(validRoles, req.Role) {
		return nil, shared.NewAPIError(400, "Invalid role. Must be one of: student, tutor, admin")
	}

	// Check if user already exists
	exists, err := s.authRepo.UserExists(req.Email)
	if err != nil {
		return nil, shared.ErrDatabaseError
	}
	if exists {
		return nil, shared.ErrUserAlreadyExists
	}

	// Hash password
	hashedPassword, err := s.passwordHash.HashPassword(req.Password)
	if err != nil {
		return nil, shared.ErrInternalServer
	}

	// Create user
	newUser := &shared.User{
		ID:       uuid.New(),
		Email:    req.Email,
		Password: hashedPassword,
		Role:     req.Role,
		Location: req.Location,
	}

	// Save user to database
	if err := s.authRepo.CreateUser(newUser); err != nil {
		return nil, shared.ErrDatabaseError
	}

	// Create user info
	userInfo := &user.UserInfo{
		ID:        uuid.New(),
		UserID:    newUser.ID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Location:  req.Location,
		Avatar:    "", // Default empty avatar
		Bio:       "", // Default empty bio
		Phone:     "", // Default empty phone
	}

	if err := s.userRepo.CreateUserInfo(userInfo); err != nil {
		// Rollback user creation
		s.userRepo.Delete(newUser.ID)
		return nil, shared.ErrDatabaseError
	}

	// Generate tokens
	tokens, err := s.tokenGen.GenerateToken(newUser.ID, newUser.Role)
	if err != nil {
		return nil, shared.ErrTokenGeneration
	}

	return tokens, nil
}

// RefreshToken generates new tokens using refresh token
func (s *service) RefreshToken(refreshToken string) (*LoginResponse, error) {
	// Validate input
	if refreshToken == "" {
		return nil, shared.ErrMissingFields
	}

	// Parse and validate refresh token
	claims, err := s.tokenGen.ParseToken(refreshToken)
	if err != nil {
		return nil, shared.ErrInvalidCredentials
	}

	// Extract user ID and role from token
	userIDStr, ok := claims["sub"].(string)
	if !ok {
		return nil, shared.ErrInvalidCredentials
	}

	role, ok := claims["role"].(string)
	if !ok {
		return nil, shared.ErrInvalidCredentials
	}

	// Parse user ID
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return nil, shared.ErrInvalidCredentials
	}

	// Generate new tokens
	tokens, err := s.tokenGen.GenerateToken(userID, role)
	if err != nil {
		return nil, shared.ErrTokenGeneration
	}

	return tokens, nil
}

// ValidateToken validates a token and returns user ID and role
func (s *service) ValidateToken(tokenString string) (uuid.UUID, string, error) {
	if tokenString == "" {
		return uuid.Nil, "", shared.ErrMissingFields
	}

	claims, err := s.tokenGen.ParseToken(tokenString)
	if err != nil {
		return uuid.Nil, "", shared.ErrInvalidCredentials
	}

	// Extract user ID and role from token
	userIDStr, ok := claims["sub"].(string)
	if !ok {
		return uuid.Nil, "", shared.ErrInvalidCredentials
	}

	role, ok := claims["role"].(string)
	if !ok {
		return uuid.Nil, "", shared.ErrInvalidCredentials
	}

	// Parse user ID
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return uuid.Nil, "", shared.ErrInvalidCredentials
	}

	return userID, role, nil
}
