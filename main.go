package main

import (
	"fmt"
	"log"
	"marketplace/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()
	app.Use(helmet.New())

	// cors
	app.Use(cors.New(cors.Config{
		AllowOrigins:  "*",
		AllowMethods:  "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:  "*",
		ExposeHeaders: "Content-Length",
	}))

	routes.Router(app)

	err = app.Listen(":3000")
	fmt.Println("app is running in port 3000")
	if err != nil {
		panic(err)
	}
}
