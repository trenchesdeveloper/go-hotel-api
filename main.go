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
		if apiError, ok := err.(*api.Error); ok {
			return ctx.Status(apiError.Code).JSON(apiError)
		}
		return api.NewError(fiber.StatusInternalServerError, "Something went wrong")
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
		bookingStore = db.NewMongoBookingStore(client)

		store = &db.Store{
			UserStore:  userStore,
			HotelStore: hotelStore,
			RoomStore:  roomStore,
			BookingStore: bookingStore,
		}
		// send the user store to the api
		userHandler = api.NewUserHandler(userStore)

		authHandler = api.NewAuthHandler(userStore)

		hotelHandler = api.NewHotelHandler(store)

		roomHandler = api.NewRoomHandler(store)

		bookingHandler = api.NewBookingHandler(store)
	)

	// user routes
	routes.UserRoutes(apiV1, userHandler)

	routes.AuthRoutes(apiV1, authHandler)

	// hotel routes
	routes.HotelRoutes(apiV1, hotelHandler, store)

	// room routes
	routes.RoomRoutes(apiV1, roomHandler, store)

	// booking routes
	routes.BookingRoutes(apiV1, bookingHandler, store)

	app.Listen(":" + *PORT)
}
