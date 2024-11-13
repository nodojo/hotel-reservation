package main

import (
	"flag"

	"github.com/gofiber/fiber/v2"
	"github.com/nodojo/hotel-reservation/api"
)

func main() {
	// command to update listenAddr to :7000: ./bin/api --listenAddr :7000
	listenAddr := flag.String("listenAddr", ":5000", "listen address of api server...")
	flag.Parse()

	app := fiber.New()
	apiv1 := app.Group("/api/v1")

	apiv1.Get("/user", api.HandleGetUser)
	apiv1.Get("/user/:id", api.HandleGetUsers)

	app.Listen(*listenAddr)
}
