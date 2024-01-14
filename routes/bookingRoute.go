package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trenchesdeveloper/go-hotel/api"
	"github.com/trenchesdeveloper/go-hotel/db"
	"github.com/trenchesdeveloper/go-hotel/middleware"
)

func BookingRoutes(r fiber.Router, bookingHandler *api.BookingHandler, store *db.Store) {
	r.Get("/booking",middleware.AuthRequired(store), middleware.AdminRequired, bookingHandler.HandleGetBookings)
	r.Get("/booking/:id", bookingHandler.HandleGetBooking)

	r.Put("/booking/:id/cancel", middleware.AuthRequired(store), bookingHandler.HandleCancelBooking)

}
