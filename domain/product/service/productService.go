package service

import "go-crud/models"

type ProductService interface {
	GetProduct(productID int) (*models.ProductResponse, error)
	GetProducts() ([]*models.ProductResponse, error)
	CreateProduct(product *models.ProductRequest) (*models.ProductResponse, error)
	UpdateProduct(productID int, product *models.ProductRequest) (*models.ProductResponse, error)
	DeleteProduct(productID int) error
}
