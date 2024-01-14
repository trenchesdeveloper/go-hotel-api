package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trenchesdeveloper/go-hotel/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HotelHandler struct {
	store *db.Store
}

func NewHotelHandler(store *db.Store) *HotelHandler {
	return &HotelHandler{
		store: store,
	}
}

func (h *HotelHandler) HandleGetHotels(c *fiber.Ctx) error {
	// Read pagination parameters
	page := c.QueryInt("page", 1) // default page is 1
	if page < 1 {
		page = 1
	}

	limit := c.QueryInt("limit", 4)
	if limit < 1 {
		limit = 10
	}

	// Call the store method with pagination parameters
	hotels, err := h.store.HotelStore.GetHotels(c.Context(), nil, page, limit)
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
		return ErrInvalidID()
	}
	filter := bson.M{"hotelID": oid}

	rooms, err := h.store.RoomStore.GetRooms(c.Context(), filter)
	if err != nil {
		return err
	}
	return c.JSON(rooms)
}

func (h *HotelHandler) HandleGetHotel(c *fiber.Ctx) error {
	id := c.Params("id")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid hotel id",
		})
	}

	hotel, err := h.store.HotelStore.GetHotelById(c.Context(), oid)
	if err != nil {
		return err
	}

	return c.JSON(hotel)
}
