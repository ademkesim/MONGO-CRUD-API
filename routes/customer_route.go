package routes

import (
	"fiber-mongo-api/controllers"

	"github.com/gofiber/fiber/v2"
)

func CustomerRoute(app *fiber.App) {
	app.Post("/customer", controllers.CreateCustomer)
	app.Get("/customer/:customerId", controllers.GetACustomer)
	app.Delete("/customer/:customerId", controllers.DeleteACustomer)
	app.Get("/customer", controllers.GetAllCustomers)
}
