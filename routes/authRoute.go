package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trenchesdeveloper/go-hotel/api"
)

func AuthRoutes(r fiber.Router, authHandler *api.AuthHandler) {

	r.Post("/auth/login", authHandler.HandleAuthenticate)
}
