package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trenchesdeveloper/go-hotel/db"
)

type BookingHandler struct {
	store *db.Store
}

func NewBookingHandler(store *db.Store) *BookingHandler {
	return &BookingHandler{
		store: store,
	}
}

// TODO: this needs to be admin only
func (h *BookingHandler) HandleGetBookings(c *fiber.Ctx) error {
	bookings, err := h.store.BookingStore.GetBookings(c.Context(), nil)

	if err != nil {
		return err
	}

	return c.JSON(bookings)

}

func (h *BookingHandler) HandleGetBooking(c *fiber.Ctx) error {
	bookingID := c.Params("id")

	booking, err := h.store.BookingStore.GetBookingById(c.Context(), bookingID)

	if err != nil {
		return err
	}

	return c.JSON(booking)
}
