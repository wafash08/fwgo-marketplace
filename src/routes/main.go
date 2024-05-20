package routes

import (
	"marketplace/src/controllers"
	"marketplace/src/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Router(a *fiber.App) {
	api := a.Group("/api/v1")

	api.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello world")
	})

	api.Post("/upload-local", controllers.UploadFileLocal)
	api.Post("/upload", controllers.UploadFileCloudinary)

	products := api.Group("/products")
	products.Get("/", controllers.FindAllProducts)
	products.Get("/:id", controllers.FindProductById)
	products.Post("/", middlewares.JwtMiddleware(), middlewares.RoleCheckMiddleware("seller"), controllers.CreateProduct)
	products.Put("/:id", controllers.UpdateProduct)
	products.Delete("/:id", controllers.DeleteProduct)

	categories := api.Group("/categories")
	categories.Get("/", controllers.FindAllCategories)
	categories.Get("/:id", controllers.FindCategoryByID)
	categories.Post("/", middlewares.JwtMiddleware(), controllers.CreateCategory)
	categories.Put("/:id", controllers.UpdateCategory)
	categories.Delete("/:id", controllers.DeleteCategory)

	auth := api.Group("/auth")
	auth.Post("/register/seller", controllers.RegisterSeller)
	auth.Post("/login/seller", controllers.LoginSeller)
	auth.Post("/register/customer", controllers.RegisterCustomer)
	auth.Post("/login/customer", controllers.LoginCustomer)
	auth.Post("/refresh-token", controllers.RefreshToken)

	seller := api.Group("/sellers")
	seller.Get("/", controllers.FindAllSellers)
	seller.Get("/:id", middlewares.JwtMiddleware(), controllers.FindSellerById)
	seller.Put("/:id", middlewares.JwtMiddleware(), controllers.UpdateSeller)
	seller.Delete("/:id", middlewares.JwtMiddleware(), controllers.DeleteSeller)

	customers := api.Group("/customers")
	customers.Get("/", controllers.FindAllCustomers)
	customers.Get("/:id", middlewares.JwtMiddleware(), controllers.FindCustomerById)
	customers.Put("/:id", middlewares.JwtMiddleware(), controllers.UpdateCustomer)
	customers.Delete("/:id", middlewares.JwtMiddleware(), controllers.DeleteCustomer)

	addresses := api.Group("/addresses")
	addresses.Get("/", controllers.FindAllAddresses)
	addresses.Get("/:id", controllers.FindAddressByID)
	addresses.Post("/", middlewares.JwtMiddleware(), controllers.CreateAddress)
	addresses.Put("/:id", controllers.UpdateAddress)
	addresses.Delete("/:id", controllers.DeleteAddress)

}
