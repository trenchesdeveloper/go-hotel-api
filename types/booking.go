package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Booking struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	RoomID primitive.ObjectID `bson:"roomID" json:"roomID"`
	FromDate   time.Time          `bson:"fromDate" json:"fromDate"`
	ToDate time.Time          `bson:"toDate" json:"toDate"`
	UserID   primitive.ObjectID `bson:"userID" json:"userID"`
	NumPersons int `bson:"numPersons" json:"numPersons"`
}

type BookRoomParams struct {
	FromDate   time.Time          `json:"fromDate"`
	ToDate time.Time          `json:"toDate"`
	NumPersons int `json:"numPersons"`
}
