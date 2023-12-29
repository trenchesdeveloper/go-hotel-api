package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trenchesdeveloper/go-hotel/api"
)

func UserRoutes(r fiber.Router, userHandler *api.UserHandler) {
	
	r.Get("/users", userHandler.HandleGetUsers)

	r.Get("/user/:id", userHandler.HandleGetUser)

	r.Post("/user", userHandler.HandleCreateUser)
	r.Delete("/user/:id", userHandler.HandleDeleteUser)

	r.Put("/user/:id", userHandler.HandleUpdateUser)
}
