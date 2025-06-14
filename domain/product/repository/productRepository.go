package repository

import (
	"go-crud/entities"
)

type ProductRepository interface {
	FindAll() (*[]entities.Product, error)
	FindByID(productID int) (*entities.Product, error)
	Insert(product *entities.Product) (*entities.Product, error)
	UpdateByID(productID int, product *entities.Product) (*entities.Product, error)
	DeleteByID(productID int) error
}
