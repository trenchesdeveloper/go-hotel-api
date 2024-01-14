package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trenchesdeveloper/go-hotel/db"
	"github.com/trenchesdeveloper/go-hotel/types"
	"go.mongodb.org/mongo-driver/bson"
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

func (h *BookingHandler) HandleCancelBooking(c *fiber.Ctx) error {
	bookingID := c.Params("id")

	booking, err := h.store.BookingStore.GetBookingById(c.Context(), bookingID)

	if err != nil {
		return err
	}

	// get the user
	user := c.Locals("user").(*types.User)

	// check if the user is the owner of the booking or an admin
	if booking.UserID != user.ID && !user.IsAdmin {
		return fiber.ErrForbidden
	}

	resp, err := h.store.BookingStore.UpdateBooking(c.Context(), bookingID, bson.M{
		"$set": bson.M{
			"canceled": true,
		},
	})

	if err != nil {
		return err
	}

	return c.JSON(resp)
}
