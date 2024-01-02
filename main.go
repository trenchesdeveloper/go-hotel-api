package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/trenchesdeveloper/go-hotel/api"
	"github.com/trenchesdeveloper/go-hotel/db"
	"github.com/trenchesdeveloper/go-hotel/routes"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var config = fiber.Config{
	// Override default error handler
	ErrorHandler: func(ctx *fiber.Ctx, err error) error {
		return ctx.JSON(fiber.Map{
			"error": err.Error(),
		})
	},
}

func main() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(db.DBURI))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	fmt.Println(client)

	PORT := flag.String("port", "4000", "port to listen on")

	flag.Parse()
	app := fiber.New(config)

	var (
		apiV1 = app.Group("/api/v1")

		// create a new mongo user store
		userStore = db.NewMongoUserStore(client)
		// create a new mongo hotel store
		hotelStore = db.NewMongoHotelStore(client)
		roomStore  = db.NewMongoRoomStore(client, hotelStore)

		store = &db.Store{
			UserStore:  userStore,
			HotelStore: hotelStore,
			RoomStore:  roomStore,
		}
		// send the user store to the api
		userHandler = api.NewUserHandler(userStore)

		authHandler = api.NewAuthHandler(userStore)

		hotelHandler = api.NewHotelHandler(store)
	)

	// user routes
	routes.UserRoutes(apiV1, userHandler)

	routes.AuthRoutes(apiV1, authHandler)

	// hotel routes
	routes.HotelRoutes(apiV1, hotelHandler)

	app.Listen(":" + *PORT)
}
