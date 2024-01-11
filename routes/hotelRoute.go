package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trenchesdeveloper/go-hotel/api"
	"github.com/trenchesdeveloper/go-hotel/middleware"
)

func HotelRoutes(r fiber.Router, hotelHandler *api.HotelHandler) {

	r.Get("/hotel", middleware.AuthRequired,  hotelHandler.HandleGetHotels)
	r.Get("/hotel/:id", hotelHandler.HandleGetHotel)

	r.Get("/hotel/:id/rooms", hotelHandler.HandleGetHotelRooms)

	// list routes that require authentication, and pass the auth middleware
	r.Post("/hotel", hotelHandler.HandleGetHotelRooms)

}
