package routes

import (
	"go-crud/database"
	"go-crud/domain/product/handler"
	"go-crud/domain/product/repository"
	"go-crud/domain/product/service"

	"github.com/gofiber/fiber/v2"
)

func RegisterProductRoute(app *fiber.App) {
	productRepository := repository.NewProductRepositoryDB(database.DB)
	productService := service.NewProductServiceImpl(productRepository)
	productHandler := handler.NewProductHandler(productService)

	app.Route("/products", func(r fiber.Router) {
		r.Get("/", productHandler.GetProduct)
		r.Get("/:productID", productHandler.GetProducts)
		r.Post("/", productHandler.CreateProduct)
		r.Put("/:productID", productHandler.UpdateProduct)
		r.Delete("/:productID", productHandler.DeleteProduct)
	})
}
