package routes

import (
	"marketplace/controllers"

	"github.com/gofiber/fiber/v2"
)

func Router(a *fiber.App) {
	api := a.Group("/api/v1")

	api.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello world")
	})

	products := api.Group("/products")
	products.Get("/", controllers.FindAllProducts)
	products.Get("/:id", controllers.FindProductById)
	products.Post("/", controllers.CreateProduct)
	products.Put("/:id", controllers.UpdateProduct)
	products.Delete("/:id", controllers.DeleteProduct)

	categories := api.Group("/categories")
	categories.Get("/", controllers.FindAllCategories)
	categories.Get("/:id", controllers.FindCategoryByID)
	categories.Post("/", controllers.CreateCategory)
	categories.Put("/:id", controllers.UpdateCategory)
	categories.Delete("/:id", controllers.DeleteCategory)
}
