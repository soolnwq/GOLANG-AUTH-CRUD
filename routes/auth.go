package routes

import (
	"go-crud/database"
	"go-crud/domain/auth/handler"
	"go-crud/domain/auth/repository"
	"go-crud/domain/auth/service"

	"github.com/gofiber/fiber/v2"
)

func RegisterAuthRoute(app *fiber.App) {
	authUserRepostiory := repository.NewAuthUserRepository(database.DB)
	authService := service.NewAuthService(authUserRepostiory)
	authHandler := handler.NewAuthHandler(authService)

	app.Route("/auth", func(r fiber.Router) {
		r.Post("/register", authHandler.Register)
		r.Post("/login", authHandler.Login)
	})
}
