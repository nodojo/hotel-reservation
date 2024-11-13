package api

import "github.com/gofiber/fiber/v2"

func HandleGetUsers(c *fiber.Ctx) error {
	return c.JSON("HandleGetUsers: John Doe, Jane Doe")
}

func HandleGetUser(c *fiber.Ctx) error {
	return c.JSON("HandleGetUser: John Doe")
}
