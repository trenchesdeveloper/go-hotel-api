package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/trenchesdeveloper/go-hotel/api"
	"github.com/trenchesdeveloper/go-hotel/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbURI = "mongodb://localhost:27017"
const DBNAME = "hotel-reservation"

var config = fiber.Config{
	// Override default error handler
	ErrorHandler: func(ctx *fiber.Ctx, err error) error {
		return ctx.JSON(fiber.Map{
			"error": err.Error(),
		})
	},
}

func main() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(dbURI))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	fmt.Println(client)

	PORT := flag.String("port", "4000", "port to listen on")

	flag.Parse()
	app := fiber.New(config)

	apiV1 := app.Group("/api/v1")

	// create a new mongo user store
	userStore := db.NewMongoUserStore(client, DBNAME)

	// send the user store to the api
	userHandler := api.NewUserHandler(userStore)

	apiV1.Get("/users", userHandler.HandleGetUsers)

	apiV1.Get("/user/:id", userHandler.HandleGetUser)

	apiV1.Post("/user", userHandler.HandleCreateUser)
	apiV1.Delete("/user/:id", userHandler.HandleDeleteUser)

	apiV1.Put("/user/:id", userHandler.HandleUpdateUser)

	app.Listen(":" + *PORT)
}
