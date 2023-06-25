package router

import (
	"people/controllers"

	"github.com/gofiber/fiber/v2"
)

func AddRoutes(app *fiber.App, controller *controllers.PersonController) {

	//Define Group people
	people := app.Group("/people")

	//HTTP METHODS
	people.Delete("/:id", controller.Delete)
	people.Get("/", controller.FindAll)
	people.Get("/:id", controller.FindById)
	people.Patch("/", controller.Update)
	people.Post("/", controller.Create)

	//Define Group ping
	ping := app.Group("/ping")

	//HTTP METHODS
	ping.Get("/", controllers.Ping)

}
