package repository

import (
	"go-crud/entities"
)

type AuthUserRepository interface {
	Insert(user *entities.User) (*entities.User, error)
	FindByUsernameOrEmail(username string, email string) (*entities.User, error)
	FindByUsername(username string) (*entities.User, error)
}
