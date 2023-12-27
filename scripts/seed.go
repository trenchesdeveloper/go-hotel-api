package main

import (
	"context"
	"fmt"
	"log"

	"github.com/trenchesdeveloper/go-hotel/db"
	"github.com/trenchesdeveloper/go-hotel/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx := context.Background()
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))

	if err != nil {
		log.Fatal(err)
	}

	hotelStore := db.NewMongoHotelStore(client, db.DBNAME)
	roomStore := db.NewMongoRoomStore(client, db.DBNAME)

	hotel := types.Hotel{
		Name:     "Rock",
		Location: "Rick",
	}

	rooms := []types.Room{
		{Type: types.SingleRoomType,
			BasePrice: 99.9,
		},
		{Type: types.DoubleRoomType,
			BasePrice: 199.9,
		},
		{Type: types.DeluxeRoomType,
			BasePrice: 299.9,
		},
	}
	insertedHotel, err := hotelStore.CreateHotel(ctx, &hotel)
	for _, room := range rooms {
		room.HotelID = insertedHotel.ID

		if err != nil {
			log.Fatal(err)
		}

		insertedRoom, err := roomStore.CreateRoom(ctx, &room)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(insertedRoom)
	}

}
