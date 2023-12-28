package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trenchesdeveloper/go-hotel/api"
)

func HotelRoutes(r fiber.Router, hotelHandler *api.HotelHandler) {
	r.Get("/hotel", hotelHandler.HandleGetHotels)
	r.Get("/hotel/:id/rooms", hotelHandler.HandleGetHotelRooms)
}
