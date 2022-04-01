package main

import (
	"fiber-mongo-api/configs"
	"fiber-mongo-api/routes"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	//run database
	configs.ConnectDB()

	//routes
	routes.CustomerRoute(app)

	port := os.Getenv("PORT")

	app.Listen(":" + port)
}
