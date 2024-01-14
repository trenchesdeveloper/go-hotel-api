package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trenchesdeveloper/go-hotel/api"
	"github.com/trenchesdeveloper/go-hotel/db"
	"github.com/trenchesdeveloper/go-hotel/middleware"
)

func RoomRoutes(r fiber.Router, roomHandler *api.RoomHandler, store *db.Store) {
	r.Use(middleware.AuthRequired(store))
	r.Post("/room/:id/book", roomHandler.HandleBookRoom)
	r.Get("/room", roomHandler.HandleGetRooms)

}
