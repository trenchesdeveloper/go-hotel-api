package db

import "os"

var (
	DBURI string
	DBNAME string
	TestDBName string
)

type Store struct {
	UserStore
	HotelStore
	RoomStore
	BookingStore
}


func init() {
	DBURI = os.Getenv("DBURI")
	DBNAME = os.Getenv("DBNAME")
	TestDBName = os.Getenv("TESTDBNAME")
}