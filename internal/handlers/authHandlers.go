package handlers

import (
	"context"

	authservice "github.com/IbadT/tutor_app_back.git/internal/authService"
	"github.com/IbadT/tutor_app_back.git/internal/errors"
	"github.com/IbadT/tutor_app_back.git/internal/web/auth"
)

type AuthHandlers struct {
	service authservice.AuthService
}

func NewAuthHandlers(service authservice.AuthService) *AuthHandlers {
	return &AuthHandlers{service: service}
}

// POST /auth/login
func (h *AuthHandlers) PostAuthLogin(ctx context.Context, request auth.PostAuthLoginRequestObject) (auth.PostAuthLoginResponseObject, error) {
	body := request.Body
	authLoginRequest := authservice.LoginRequest{
		Email:    string(body.Email),
		Password: body.Password,
	}

	r, err := h.service.Login(authLoginRequest)
	if err != nil {
		return h.handleLoginError(err)
	}

	response := auth.PostAuthLogin200JSONResponse{
		AccessToken:  &r.AccessToken,
		RefreshToken: &r.RefreshToken,
	}
	return response, nil
}

// handleLoginError converts service errors to appropriate HTTP responses
func (h *AuthHandlers) handleLoginError(err error) (auth.PostAuthLoginResponseObject, error) {
	if apiErr, ok := err.(*errors.APIError); ok {
		switch apiErr.Code {
		case 400:
			return auth.PostAuthLogin400JSONResponse{Code: &apiErr.Code, Message: &apiErr.Message}, nil
		case 401:
			return auth.PostAuthLogin401JSONResponse{Code: &apiErr.Code, Message: &apiErr.Message}, nil
		default:
			return auth.PostAuthLogin500JSONResponse{Code: &apiErr.Code, Message: &apiErr.Message}, nil
		}
	}
	// Fallback for unexpected errors
	code := 500
	msg := "Internal server error"
	return auth.PostAuthLogin500JSONResponse{Code: &code, Message: &msg}, nil
}

// POST /auth/register
func (h *AuthHandlers) PostAuthRegister(ctx context.Context, request auth.PostAuthRegisterRequestObject) (auth.PostAuthRegisterResponseObject, error) {
	body := request.Body
	authRegisterRequest := authservice.RegisterRequest{
		FirstName: body.FirstName,
		LastName:  string(body.LastName),
		Email:     string(body.Email),
		Password:  body.Password,
		Role:      string(body.Role),
		Location:  body.Location,
	}

	r, err := h.service.Register(authRegisterRequest)
	if err != nil {
		return h.handleRegisterError(err)
	}

	response := auth.PostAuthRegister201JSONResponse{
		AccessToken:  &r.AccessToken,
		RefreshToken: &r.RefreshToken,
	}
	return response, nil
}

// handleRegisterError converts service errors to appropriate HTTP responses
func (h *AuthHandlers) handleRegisterError(err error) (auth.PostAuthRegisterResponseObject, error) {
	if apiErr, ok := err.(*errors.APIError); ok {
		switch apiErr.Code {
		case 400:
			return auth.PostAuthRegister400JSONResponse{Code: &apiErr.Code, Message: &apiErr.Message}, nil
		case 409:
			return auth.PostAuthRegister409JSONResponse{Code: &apiErr.Code, Message: &apiErr.Message}, nil
		default:
			return auth.PostAuthRegister500JSONResponse{Code: &apiErr.Code, Message: &apiErr.Message}, nil
		}
	}
	// Fallback for unexpected errors
	code := 500
	msg := "Internal server error"
	return auth.PostAuthRegister500JSONResponse{Code: &code, Message: &msg}, nil
}

// POST /auth/refresh-token
func (h *AuthHandlers) PostAuthRefreshToken(ctx context.Context, request auth.PostAuthRefreshTokenRequestObject) (auth.PostAuthRefreshTokenResponseObject, error) {
	body := request.Body

	r, err := h.service.RefreshToken(body.RefreshToken)
	if err != nil {
		return h.handleRefreshError(err)
	}

	response := auth.PostAuthRefreshToken200JSONResponse{
		AccessToken:  &r.AccessToken,
		RefreshToken: &r.RefreshToken,
	}
	return response, nil
}

// handleRefreshError converts service errors to appropriate HTTP responses
func (h *AuthHandlers) handleRefreshError(err error) (auth.PostAuthRefreshTokenResponseObject, error) {
	if apiErr, ok := err.(*errors.APIError); ok {
		switch apiErr.Code {
		case 400:
			return auth.PostAuthRefreshToken400JSONResponse{Code: &apiErr.Code, Message: &apiErr.Message}, nil
		case 401:
			return auth.PostAuthRefreshToken401JSONResponse{Code: &apiErr.Code, Message: &apiErr.Message}, nil
		default:
			return auth.PostAuthRefreshToken500JSONResponse{Code: &apiErr.Code, Message: &apiErr.Message}, nil
		}
	}
	// Fallback for unexpected errors
	code := 500
	msg := "Internal server error"
	return auth.PostAuthRefreshToken500JSONResponse{Code: &code, Message: &msg}, nil
}
