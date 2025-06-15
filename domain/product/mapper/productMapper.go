package mapper

import (
	"go-crud/entities"
	"go-crud/models"
)

func ProdcutRequestToEntity(req *models.ProductRequest) *entities.Product {
	return &entities.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
	}
}

func ProductEntityToResponse(p *entities.Product) *models.ProductResponse {
	return &models.ProductResponse{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
	}
}

func ProductEntityListToResponse(list []entities.Product) []*models.ProductResponse {
	res := make([]*models.ProductResponse, 0, len(list))
	for _, p := range list {
		res = append(res, ProductEntityToResponse(&p))
	}
	return res
}
