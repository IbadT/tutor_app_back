package app

import (
	"log"

	"github.com/IbadT/tutor_app_back.git/internal/app/handlers"
	"github.com/IbadT/tutor_app_back.git/internal/app/middleware"
	"github.com/IbadT/tutor_app_back.git/internal/domain/auth"
	"github.com/IbadT/tutor_app_back.git/internal/domain/user"
	"github.com/IbadT/tutor_app_back.git/internal/infrastructure/database"
	"github.com/IbadT/tutor_app_back.git/internal/infrastructure/external"
	"github.com/IbadT/tutor_app_back.git/internal/infrastructure/repositories"
	web_auth "github.com/IbadT/tutor_app_back.git/internal/web/auth"
	web_users "github.com/IbadT/tutor_app_back.git/internal/web/users"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

// Server represents the application server
type Server struct {
	echo *echo.Echo
}

// NewServer creates a new server instance
func NewServer() (*Server, error) {
	// Initialize database
	db, err := database.InitDB()
	if err != nil {
		return nil, err
	}

	// Initialize Echo
	e := echo.New()

	// Set custom error handler
	e.HTTPErrorHandler = middleware.ErrorHandler

	// Initialize repositories
	userRepo := repositories.NewUserRepository(db)
	authRepo := repositories.NewAuthRepository(db)

	// Initialize external services
	jwtService := external.NewJWTService()
	passwordService := external.NewPasswordService()

	// Initialize domain services
	userService := user.NewService(userRepo)
	authService := auth.NewService(authRepo, userRepo, jwtService, passwordService)

	// Initialize handlers
	userHandler := handlers.NewUserHandler(userService)
	authHandler := handlers.NewAuthHandler(authService)

	// Create strict handlers for OpenAPI
	userStrictHandler := web_users.NewStrictHandler(userHandler, nil)
	authStrictHandler := web_auth.NewStrictHandler(authHandler, nil)

	// Register routes
	registerRoutes(e, userStrictHandler, authStrictHandler, authService)

	// Setup middleware
	setupMiddleware(e)

	return &Server{echo: e}, nil
}

// registerRoutes registers all application routes
func registerRoutes(e *echo.Echo, userHandler web_users.ServerInterface, authHandler web_auth.ServerInterface, authService auth.Service) {
	// Static files for Swagger UI
	e.Static("/swagger", "static/swagger")

	// OpenAPI specification file
	e.File("/openapi.yaml", "openapi/openapi.yaml")

	// Auth routes (no authentication required)
	web_auth.RegisterHandlers(e, authHandler)

	// User routes (authentication required)
	userGroup := e.Group("/users")
	userGroup.Use(middleware.AuthMiddleware(authService))
	web_users.RegisterHandlers(e, userHandler)
}

// setupMiddleware configures Echo middleware
func setupMiddleware(e *echo.Echo) {
	// Logger middleware
	e.Use(echoMiddleware.Logger())

	// Recover middleware
	e.Use(echoMiddleware.Recover())

	// CORS middleware
	e.Use(echoMiddleware.CORS())

	// Request ID middleware
	e.Use(echoMiddleware.RequestID())
}

// Start starts the server
func (s *Server) Start(port string) error {
	log.Printf("Server starting on port %s", port)
	return s.echo.Start(":" + port)
}

// Shutdown gracefully shuts down the server
func (s *Server) Shutdown() error {
	return s.echo.Close()
}
