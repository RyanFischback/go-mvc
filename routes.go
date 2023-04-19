package main

import (
	"mvc/controllers"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	//Routes
	app.Get("/", controllers.PostsIndex)
}
