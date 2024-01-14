package db

import "os"

var (
	TestDBName string
)

type Store struct {
	UserStore
	HotelStore
	RoomStore
	BookingStore
}


func init() {
	TestDBName = os.Getenv("TESTDBNAME")
}