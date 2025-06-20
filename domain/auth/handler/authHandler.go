package handler

import (
	"go-crud/domain/auth/service"
	"go-crud/errs"
	"go-crud/models"
	"go-crud/utils"

	"github.com/gofiber/fiber/v2"
)

type authHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *authHandler {
	return &authHandler{authService: authService}
}

func (a *authHandler) Login(c *fiber.Ctx) error {
	var loginRequest models.LoginRequest
	if err := c.BodyParser(&loginRequest); err != nil {
		return errs.HandleError(c, errs.NewAppError("invalid body json request", 400))
	}

	loginResponse, err := a.authService.Login(&loginRequest)
	if err != nil {
		return errs.HandleError(c, err)
	}

	utils.SetCookie(c, "access_token", loginResponse.AccessToken, 86400)

	return c.Status(200).JSON(&fiber.Map{
		"message": "login success",
	})
}

func (a *authHandler) Register(c *fiber.Ctx) error {
	var registerRequest models.RegisterRequest
	if err := c.BodyParser(&registerRequest); err != nil {
		return errs.HandleError(c, errs.NewAppError("invalid body json request", 400))
	}

	if err := a.authService.Register(&registerRequest); err != nil {
		return errs.HandleError(c, err)
	}

	return c.Status(201).JSON(&fiber.Map{
		"message": "register success",
	})
}
