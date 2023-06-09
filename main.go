package main

import (
	"mvc/controllers"
	"mvc/initializers"
	"mvc/middleware"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
	initializers.SyncDB()
}

func main() {

	//Load Templates
	engine := html.New("./views", ".tmpl")

	//Setup
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	//Configure
	app.Static("/", "./public")
	app.Use(middleware.RequireAuth)

	//Routes
	app.Get("/", controllers.PostsIndex)
	//app.Get("/about", controllers.PostsIndex)

	//Start
	app.Listen(":" + os.Getenv("PORT"))
}
