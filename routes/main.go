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

	seller := api.Group("sellers")
	seller.Get("/", controllers.FindAllSellers)
	seller.Get("/:id", controllers.FindSellerById)
	seller.Put("/:id", controllers.UpdateSeller)
	seller.Delete("/:id", controllers.DeleteSeller)

	auth := api.Group("auth")
	auth.Post("/register/seller", controllers.RegisterSeller)
	auth.Post("/login/seller", controllers.LoginSeller)

	addresses := api.Group("addresses")
	addresses.Get("/", controllers.FindAllAddresses)
	addresses.Get("/:id", controllers.FindAddressByID)
	addresses.Post("/", controllers.CreateAddress)
	addresses.Post("/:id", controllers.UpdateAddress)
	addresses.Delete("/:id", controllers.DeleteAddress)
}
