package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trenchesdeveloper/go-hotel/api"
	"github.com/trenchesdeveloper/go-hotel/middleware"
)

func BookingRoutes(r fiber.Router, bookingHandler *api.BookingHandler) {
	r.Use(middleware.AuthRequired)
	r.Get("/booking",  bookingHandler.HandleGetBookings)
	r.Get("/booking/:id",  bookingHandler.HandleGetBooking)

}
