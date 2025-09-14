package external

import (
	"fmt"
	"os"
	"time"

	"github.com/IbadT/tutor_app_back.git/internal/domain/auth"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// jwtService implements the auth.TokenGenerator interface
type jwtService struct {
	secretKey string
}

// NewJWTService creates a new JWT service
func NewJWTService() auth.TokenGenerator {
	return &jwtService{
		secretKey: getJWTSecret(),
	}
}

// GenerateToken generates access and refresh tokens
func (j *jwtService) GenerateToken(userID uuid.UUID, role string) (*auth.LoginResponse, error) {
	// Generate access token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  userID.String(),
		"role": role,
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
		"type": "access",
	})

	accessTokenString, err := accessToken.SignedString([]byte(j.secretKey))
	if err != nil {
		return nil, fmt.Errorf("failed to sign access token: %w", err)
	}

	// Generate refresh token
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  userID.String(),
		"role": role,
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(time.Hour * 24 * 7).Unix(),
		"type": "refresh",
	})

	refreshTokenString, err := refreshToken.SignedString([]byte(j.secretKey))
	if err != nil {
		return nil, fmt.Errorf("failed to sign refresh token: %w", err)
	}

	return &auth.LoginResponse{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}

// ParseToken parses and validates a JWT token
func (j *jwtService) ParseToken(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

// ValidateToken validates a token and returns user ID and role
func (j *jwtService) ValidateToken(tokenString string) error {
	_, err := j.ParseToken(tokenString)
	return err
}

// getJWTSecret retrieves JWT secret from environment
func getJWTSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		// In production, this should be a proper secret
		secret = "default-secret-key-change-in-production"
	}
	return secret
}
