package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	apiv1 := app.Group("/api/v1")

	app.Get("/test", handleTest)
	apiv1.Get("/user", handleUser)

	app.Listen(":5000")
}

func handleTest(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"msg": "I am working..."})
}

func handleUser(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"user": "John Doe"})
}
