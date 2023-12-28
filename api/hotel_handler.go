package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trenchesdeveloper/go-hotel/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HotelHandler struct {
	hotelStore db.HotelStore
	roomStore  db.RoomStore
}

func NewHotelHandler(hotelStore db.HotelStore, roomStore db.RoomStore) *HotelHandler {
	return &HotelHandler{
		hotelStore: hotelStore,
		roomStore:  roomStore,
	}
}


func (h *HotelHandler) HandleGetHotels(c *fiber.Ctx) error {
	hotels, err := h.hotelStore.GetHotels(c.Context(), nil)
	if err != nil {
		return err
	}
	return c.JSON(hotels)

}

func (h *HotelHandler) HandleGetHotelRooms(c *fiber.Ctx) error {
	var (
		hotelID = c.Params("id")
	)
	// check if it is a valid object id
	oid, err := primitive.ObjectIDFromHex(hotelID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid room id",
		})
	}
	filter := bson.M{"hotelID": oid}

	rooms, err := h.roomStore.GetRooms(c.Context(), filter)
	if err != nil {
		return err
	}
	return c.JSON(rooms)
}
