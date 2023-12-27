package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trenchesdeveloper/go-hotel/db"
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

type HotelQueryParams struct {
	Rooms bool
	Rating float64
}

func (h *HotelHandler) HandleGetHotels(c *fiber.Ctx) error {
	var qparams HotelQueryParams

	if err := c.QueryParser(&qparams); err != nil {
		return err
	}
	hotels, err := h.hotelStore.GetHotels(c.Context(), nil)
	if err != nil {
		return err
	}
	return c.JSON(hotels)

}
