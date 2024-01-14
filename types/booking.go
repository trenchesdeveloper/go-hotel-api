package types

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Booking struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	RoomID     primitive.ObjectID `bson:"roomID" json:"roomID"`
	FromDate   time.Time          `bson:"fromDate" json:"fromDate"`
	ToDate     time.Time          `bson:"toDate" json:"toDate"`
	UserID     primitive.ObjectID `bson:"userID" json:"userID"`
	NumPersons int                `bson:"numPersons" json:"numPersons"`
	Canceled   bool               `bson:"canceled" json:"canceled"`
}

type BookRoomParams struct {
	FromDate   time.Time `json:"fromDate"`
	ToDate     time.Time `json:"toDate"`
	NumPersons int       `json:"numPersons"`
}

func (p BookRoomParams) Validate() error {
	now := time.Now()
	if p.FromDate.Before(now) {
		return fmt.Errorf("check in date must be after today")
	}
	if p.ToDate.Before(p.FromDate) {
		return fmt.Errorf("check out date must be after check in date")
	}
	if p.NumPersons < 1 {
		return fmt.Errorf("number of persons must be greater than 0")
	}
	return nil
}
