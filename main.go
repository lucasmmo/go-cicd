package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		c.Send([]byte("Testing webhook to automaticated build"))
		return nil
	})
	app.Listen(":3000")
}
