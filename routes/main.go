package routes

import "github.com/gofiber/fiber/v2"

func Router(a *fiber.App) {
	api := a.Group("/api/v1")

	api.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello world")
	})
}
