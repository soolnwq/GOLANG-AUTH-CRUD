package mapper

import (
	"go-crud/entities"
	"go-crud/models"
)

func RegisterRequestToEntity(req *models.RegisterRequest) *entities.User {
	return &entities.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}
}
