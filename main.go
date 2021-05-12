package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		c.Send([]byte("Hello WORLD"))
		return nil
	})
	app.Listen(":3000")
}
