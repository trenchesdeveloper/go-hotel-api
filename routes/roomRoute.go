package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trenchesdeveloper/go-hotel/api"
	"github.com/trenchesdeveloper/go-hotel/middleware"
)

func RoomRoutes(r fiber.Router, roomHandler *api.RoomHandler) {
	r.Use(middleware.AuthRequired)
	r.Post("/room/:id/book",  roomHandler.HandleBookRoom)
	r.Get("/room",  roomHandler.HandleGetRooms)

}
