package main

import (
	"flag"

	"github.com/gofiber/fiber/v2"
	"github.com/trenchesdeveloper/go-hotel/api"
)

func main() {
	PORT := flag.String("port", "4000", "port to listen on")

	flag.Parse()
	app := fiber.New()

	apiV1 := app.Group("/api/v1")

	apiV1.Get("/user", api.HandleGetUsers)

	apiV1.Get("/user/:id", api.HandleGetUser)

	app.Listen(":" + *PORT)
}
