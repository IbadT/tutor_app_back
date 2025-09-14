package handlers

import (
	"context"

	"github.com/IbadT/tutor_app_back.git/internal/domain/auth"
	"github.com/IbadT/tutor_app_back.git/internal/domain/shared"
	web_auth "github.com/IbadT/tutor_app_back.git/internal/web/auth"
)

// AuthHandler handles authentication requests
type AuthHandler struct {
	authService auth.Service
}

// NewAuthHandler creates a new auth handler
func NewAuthHandler(authService auth.Service) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// PostAuthLogin handles POST /auth/login
func (h *AuthHandler) PostAuthLogin(ctx context.Context, request web_auth.PostAuthLoginRequestObject) (web_auth.PostAuthLoginResponseObject, error) {
	body := request.Body
	loginRequest := &auth.LoginRequest{
		Email:    string(body.Email),
		Password: body.Password,
	}

	response, err := h.authService.Login(loginRequest)
	if err != nil {
		return h.handleLoginError(err)
	}

	return web_auth.PostAuthLogin200JSONResponse{
		AccessToken:  &response.AccessToken,
		RefreshToken: &response.RefreshToken,
	}, nil
}

// PostAuthRegister handles POST /auth/register
func (h *AuthHandler) PostAuthRegister(ctx context.Context, request web_auth.PostAuthRegisterRequestObject) (web_auth.PostAuthRegisterResponseObject, error) {
	body := request.Body
	registerRequest := &auth.RegisterRequest{
		FirstName: body.FirstName,
		LastName:  string(body.LastName),
		Email:     string(body.Email),
		Password:  body.Password,
		Role:      string(body.Role),
		Location:  body.Location,
	}

	response, err := h.authService.Register(registerRequest)
	if err != nil {
		return h.handleRegisterError(err)
	}

	return web_auth.PostAuthRegister201JSONResponse{
		AccessToken:  &response.AccessToken,
		RefreshToken: &response.RefreshToken,
	}, nil
}

// PostAuthRefreshToken handles POST /auth/refresh-token
func (h *AuthHandler) PostAuthRefreshToken(ctx context.Context, request web_auth.PostAuthRefreshTokenRequestObject) (web_auth.PostAuthRefreshTokenResponseObject, error) {
	body := request.Body

	response, err := h.authService.RefreshToken(body.RefreshToken)
	if err != nil {
		return h.handleRefreshError(err)
	}

	return web_auth.PostAuthRefreshToken200JSONResponse{
		AccessToken:  &response.AccessToken,
		RefreshToken: &response.RefreshToken,
	}, nil
}

// handleLoginError converts service errors to appropriate HTTP responses
func (h *AuthHandler) handleLoginError(err error) (web_auth.PostAuthLoginResponseObject, error) {
	if apiErr, ok := err.(*shared.APIError); ok {
		switch apiErr.Code {
		case 400:
			return web_auth.PostAuthLogin400JSONResponse{Code: &apiErr.Code, Message: &apiErr.Message}, nil
		case 401:
			return web_auth.PostAuthLogin401JSONResponse{Code: &apiErr.Code, Message: &apiErr.Message}, nil
		default:
			return web_auth.PostAuthLogin500JSONResponse{Code: &apiErr.Code, Message: &apiErr.Message}, nil
		}
	}
	// Fallback for unexpected errors
	code := 500
	msg := "Internal server error"
	return web_auth.PostAuthLogin500JSONResponse{Code: &code, Message: &msg}, nil
}

// handleRegisterError converts service errors to appropriate HTTP responses
func (h *AuthHandler) handleRegisterError(err error) (web_auth.PostAuthRegisterResponseObject, error) {
	if apiErr, ok := err.(*shared.APIError); ok {
		switch apiErr.Code {
		case 400:
			return web_auth.PostAuthRegister400JSONResponse{Code: &apiErr.Code, Message: &apiErr.Message}, nil
		case 409:
			return web_auth.PostAuthRegister409JSONResponse{Code: &apiErr.Code, Message: &apiErr.Message}, nil
		default:
			return web_auth.PostAuthRegister500JSONResponse{Code: &apiErr.Code, Message: &apiErr.Message}, nil
		}
	}
	// Fallback for unexpected errors
	code := 500
	msg := "Internal server error"
	return web_auth.PostAuthRegister500JSONResponse{Code: &code, Message: &msg}, nil
}

// handleRefreshError converts service errors to appropriate HTTP responses
func (h *AuthHandler) handleRefreshError(err error) (web_auth.PostAuthRefreshTokenResponseObject, error) {
	if apiErr, ok := err.(*shared.APIError); ok {
		switch apiErr.Code {
		case 400:
			return web_auth.PostAuthRefreshToken400JSONResponse{Code: &apiErr.Code, Message: &apiErr.Message}, nil
		case 401:
			return web_auth.PostAuthRefreshToken401JSONResponse{Code: &apiErr.Code, Message: &apiErr.Message}, nil
		default:
			return web_auth.PostAuthRefreshToken500JSONResponse{Code: &apiErr.Code, Message: &apiErr.Message}, nil
		}
	}
	// Fallback for unexpected errors
	code := 500
	msg := "Internal server error"
	return web_auth.PostAuthRefreshToken500JSONResponse{Code: &code, Message: &msg}, nil
}
