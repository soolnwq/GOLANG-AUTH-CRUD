package handler

import (
	"go-crud/domain/product/service"
	"go-crud/errs"
	"go-crud/models"

	"github.com/gofiber/fiber/v2"
)

type productHandler struct {
	productService service.ProductService
}

func NewProductHandler(productService service.ProductService) *productHandler {
	return &productHandler{productService: productService}
}

func (p *productHandler) GetProduct(c *fiber.Ctx) error {
	products, err := p.productService.GetProducts()
	if err != nil {
		return errs.HandleError(c, err)
	}

	return c.Status(200).JSON(products)
}

func (p *productHandler) GetProducts(c *fiber.Ctx) error {
	productID, err := c.ParamsInt("productID")
	if err != nil {
		return errs.HandleError(c, errs.NewAppError("invalid product id", 400))
	}

	product, err := p.productService.GetProduct(productID)
	if err != nil {
		return errs.HandleError(c, err)
	}

	return c.Status(200).JSON(product)
}

func (p *productHandler) CreateProduct(c *fiber.Ctx) error {
	var productRequest models.ProductRequest

	if err := c.BodyParser(&productRequest); err != nil {
		return errs.HandleError(c, errs.NewAppError(err.Error(), 400))
	}

	product, err := p.productService.CreateProduct(&productRequest)
	if err != nil {
		return errs.HandleError(c, err)
	}

	return c.Status(201).JSON(&fiber.Map{
		"message": "product created",
		"data":    product,
	})
}

func (p *productHandler) UpdateProduct(c *fiber.Ctx) error {
	productID, err := c.ParamsInt("productID")
	if err != nil {
		return errs.HandleError(c, errs.NewAppError("invalid product id", 400))
	}

	var productRequest models.ProductRequest

	if err := c.BodyParser(&productRequest); err != nil {
		return errs.HandleError(c, errs.NewAppError(err.Error(), 400))
	}

	product, err := p.productService.UpdateProduct(productID, &productRequest)
	if err != nil {
		return errs.HandleError(c, err)
	}

	return c.Status(200).JSON(&fiber.Map{
		"message": "product updated",
		"data":    product,
	})
}

func (p *productHandler) DeleteProduct(c *fiber.Ctx) error {
	productID, err := c.ParamsInt("productID")
	if err != nil {
		return errs.HandleError(c, errs.NewAppError("invalid product id", 400))
	}

	if err := p.productService.DeleteProduct(productID); err != nil {
		return errs.HandleError(c, err)
	}

	return c.Status(200).JSON(&fiber.Map{
		"message": "product deleted",
	})
}
