package service

import (
	"database/sql"
	"errors"
	"go-crud/domain/product/mapper"
	"go-crud/domain/product/repository"
	"go-crud/errs"
	"go-crud/models"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type productServiceImpl struct {
	productRepository repository.ProductRepository
	validate          *validator.Validate
}

func NewProductServiceImpl(productRepository repository.ProductRepository) ProductService {
	return &productServiceImpl{
		productRepository: productRepository,
		validate:          validator.New(validator.WithRequiredStructEnabled()),
	}
}

func (p *productServiceImpl) GetProduct(productID int) (*models.ProductResponse, error) {
	product, err := p.productRepository.FindByID(productID)
	if err != nil {
		zap.L().Error(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NewNotFoundError("product not found")
		}
		return nil, errs.NewUnexpectedError("unexpected error")
	}
	return mapper.ToResponse(product), nil
}

func (p *productServiceImpl) GetProducts() ([]*models.ProductResponse, error) {
	products, err := p.productRepository.FindAll()
	if err != nil {
		zap.L().Error(err.Error())
		return nil, errs.NewUnexpectedError("unexpected error")
	}
	return mapper.ToResponseList(*products), nil
}

func (p *productServiceImpl) CreateProduct(productRequest *models.ProductRequest) (*models.ProductResponse, error) {

	if err := p.validate.Struct(productRequest); err != nil {
		return nil, errs.ParseValidationErrors(err)
	}

	product, err := p.productRepository.Insert(mapper.ToEntity(productRequest))
	if err != nil {
		zap.L().Error(err.Error())
		return nil, errs.NewUnexpectedError("unexpected error")
	}
	return mapper.ToResponse(product), nil
}

func (p *productServiceImpl) UpdateProduct(productID int, productRequest *models.ProductRequest) (*models.ProductResponse, error) {
	if err := p.validate.Struct(productRequest); err != nil {
		return nil, errs.ParseValidationErrors(err)
	}

	product, err := p.productRepository.UpdateByID(productID, mapper.ToEntity(productRequest))
	if err != nil {
		zap.L().Error(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NewNotFoundError("product not found")
		}
		return nil, errs.NewUnexpectedError("unexpected error")
	}
	return mapper.ToResponse(product), nil
}

func (p *productServiceImpl) DeleteProduct(productID int) error {
	if err := p.productRepository.DeleteByID(productID); err != nil {
		zap.L().Error(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return errs.NewNotFoundError("product not found")
		}
		return errs.NewUnexpectedError("unexpected error")
	}
	return nil
}
