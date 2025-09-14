package main

import (
	"log"

	authservice "github.com/IbadT/tutor_app_back.git/internal/authService"
	"github.com/IbadT/tutor_app_back.git/internal/db"
	"github.com/IbadT/tutor_app_back.git/internal/handlers"
	custom_middleware "github.com/IbadT/tutor_app_back.git/internal/middleware"
	userservice "github.com/IbadT/tutor_app_back.git/internal/userService"
	"github.com/IbadT/tutor_app_back.git/internal/web/auth"
	"github.com/IbadT/tutor_app_back.git/internal/web/users"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Println("Error initializing database: ", err)
		log.Fatalf("Error initializing database: %v", err)
		panic(err)
	}

	e := echo.New()

	// Set custom error handler
	e.HTTPErrorHandler = custom_middleware.ErrorHandler

	userRepo := userservice.NewUserRepository(database)
	userService := userservice.NewUserService(userRepo)
	userHandlers := handlers.NewUserHandlers(userService)
	userStrictHandler := users.NewStrictHandler(userHandlers, nil)
	users.RegisterHandlers(e, userStrictHandler)

	authRepo := authservice.NewAuthRepository(database)
	authService := authservice.NewAuthService(authRepo, userRepo)
	authHandlers := handlers.NewAuthHandlers(authService)
	authStrictHandler := auth.NewStrictHandler(authHandlers, nil)
	auth.RegisterHandlers(e, authStrictHandler)

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	log.Fatal(e.Start(":8080"))
}
