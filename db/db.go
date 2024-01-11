package db

const (
	DBURI      = "mongodb://localhost:27017"
	DBNAME     = "hotel-reservation"
	TestDBName = "hotel-reservation-test"
)

type Store struct {
	UserStore
	HotelStore
	RoomStore
	BookingStore
}
