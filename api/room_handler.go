package api

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/trenchesdeveloper/go-hotel/db"
	"github.com/trenchesdeveloper/go-hotel/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RoomHandler struct {
	store *db.Store
}

func NewRoomHandler(store *db.Store) *RoomHandler {
	return &RoomHandler{
		store: store,
	}
}

func (h *RoomHandler) HandleBookRoom(c *fiber.Ctx) error {
	var params types.BookRoomParams

	if err := c.BodyParser(&params); err != nil {
		return err
	}

	if err := params.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	roomIDParams := c.Params("id")
	// check if it is a valid object id
	roomID, err := primitive.ObjectIDFromHex(roomIDParams)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid room id",
		})
	}
	userID := c.Locals("userID").(string)

	// check if the user exists
	user, err := h.store.UserStore.GetUserById(c.Context(), userID)
	if err != nil {
		log.Println("error", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "unauthorized",
		})
	}

	// check if the room is not booked
	bookings, err := h.store.BookingStore.GetBookings(c.Context(), bson.M{
		"roomID": roomID,
		"fromDate": bson.M{
			"$lte": params.ToDate,
		},
		"toDate": bson.M{
			"$gte": params.FromDate,
		},
	})

	if err != nil {
		return err
	}

	if len(bookings) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("room is not available from %s to %s", params.FromDate, params.ToDate),
		})
	}

	booking := types.Booking{
		UserID:     user.ID,
		RoomID:     roomID,
		FromDate:   params.FromDate,
		ToDate:     params.ToDate,
		NumPersons: params.NumPersons,
	}

	log.Println("booking", booking)
	inserted, err := h.store.BookingStore.CreateBooking(c.Context(), &booking)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusCreated).JSON(inserted)
}
