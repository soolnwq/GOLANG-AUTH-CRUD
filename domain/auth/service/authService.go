package service

import "go-crud/models"

type AuthService interface {
	Login(loginRequest *models.LoginRequest) (*models.LoginResponse, error)
	Register(registerRequest *models.RegisterRequest) error
}
