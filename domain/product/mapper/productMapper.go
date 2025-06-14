package mapper

import (
	"go-crud/entities"
	"go-crud/models"
)

func ToEntity(req *models.ProductRequest) *entities.Product {
	return &entities.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
	}
}

func ToResponse(p *entities.Product) *models.ProductResponse {
	return &models.ProductResponse{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
	}
}

func ToResponseList(list []entities.Product) []*models.ProductResponse {
	res := make([]*models.ProductResponse, 0, len(list))
	for _, p := range list {
		res = append(res, ToResponse(&p))
	}
	return res
}
