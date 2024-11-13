package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/test", handleTest)
	app.Listen(":5000")
}

func handleTest(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"msg": "I am working..."})
}
