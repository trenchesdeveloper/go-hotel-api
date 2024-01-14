package main

import (
	"context"
	"fmt"
	"log"

	"github.com/trenchesdeveloper/go-hotel/db"
	"github.com/trenchesdeveloper/go-hotel/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client     *mongo.Client
	hotelStore db.HotelStore
	roomStore  db.RoomStore
	userStore  db.UserStore
	ctx        = context.Background()
)

func seedUser(fname, lname, email, password string) {

	user, err := types.NewUserFromParams(types.CreateUserParams{
		FirstName: fname,
		LastName:  lname,
		Email:     email,
		Password:  password,
	})

	if err != nil {
		log.Fatal(err)
	}

	_, err = userStore.CreateUser(ctx, user)
	if err != nil {
		log.Fatal(err)
	}
}

func seedHotel(rating float64, name, location string) {
	hotel := types.Hotel{
		Name:     name,
		Location: location,
		Rooms:    []primitive.ObjectID{},
		Rating:   rating,
	}

	rooms := []types.Room{
		{
			Size:  "small",
			Price: 99.9,
		},
		{
			Size:  "normal",
			Price: 199.9,
		},
		{
			Size:  "kingsize",
			Price: 299.9,
		},
	}
	insertedHotel, err := hotelStore.CreateHotel(ctx, &hotel)
	if err != nil {
		log.Fatal(err)
	}
	for _, room := range rooms {
		room.HotelID = insertedHotel.ID

		insertedRoom, err := roomStore.CreateRoom(ctx, &room)

		if err != nil {
			log.Fatal(err)
		}

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(insertedRoom)
	}
}

func main() {
	seedHotel(4.5, "Hilton", "New York")
	seedHotel(4.0, "The commonwealth", "Boston")

	for i := 0; i< 10; i++ {
		seedHotel(4.0, fmt.Sprintf("Hotel %d", i), fmt.Sprintf("Location %d", i))
	}

	seedUser("John", "Doe", "john@john.com", "123456")
}

func init() {
	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))

	if err != nil {
		log.Fatal(err)
	}

	// drop database
	if err = client.Database(db.DBNAME).Drop(ctx); err != nil {
		log.Fatal(err)
	}

	hotelStore = db.NewMongoHotelStore(client)
	roomStore = db.NewMongoRoomStore(client, hotelStore)
	userStore = db.NewMongoUserStore(client)
}
