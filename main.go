package main

import (
	"mvc/controllers"
	"mvc/initializers"
	"os"

	"github.com/gofiber/fiber/v2"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
	initializers.SyncDB()
}

func main() {

	//Setup
	app := fiber.New()

	//Routes
	app.Get("/", controllers.PostsIndex)

	//Start
	app.Listen(":" + os.Getenv("PORT"))
}
