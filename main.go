package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()
	app.Get("/indexgi", func(c *fiber.Ctx) error {
		return c.SendString("API is working")
	})

	app.Listen(":3000")

}
