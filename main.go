package main

import (
	"flag"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// command to update listenAddr to :7000: ./bin/api --listenAddr :7000
	listenAddr := flag.String("listenAddr", ":5000", "listen address of api server...")
	flag.Parse()

	app := fiber.New()
	apiv1 := app.Group("/api/v1")

	app.Get("/test", handleTest)
	apiv1.Get("/user", handleUser)

	app.Listen(*listenAddr)
}

func handleTest(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"msg": "I am working..."})
}

func handleUser(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"user": "John Doe"})
}
