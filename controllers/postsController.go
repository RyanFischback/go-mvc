package controllers

import "github.com/gofiber/fiber/v2"

func PostsIndex(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
