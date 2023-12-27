package api

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/trenchesdeveloper/go-hotel/db"
	"github.com/trenchesdeveloper/go-hotel/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const testdbURI = "mongodb://localhost:27017"
const dbName = "hotel-reservation-test"

type responseStruct struct {
	Message string     `json:"message"`
	User    types.User `json:"user"`
}

type testDb struct {
	db.UserStore
}

func (db *testDb) teardown(t *testing.T) {
	if err := db.UserStore.Drop(context.TODO()); err != nil {
		log.Fatal(err)
	}
}

func setup(t *testing.T) *testDb {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(testdbURI))

	if err != nil {
		log.Fatal(err)
	}

	return &testDb{
		UserStore: db.NewMongoUserStore(client, dbName),
	}
}

func TestCreateUser(t *testing.T) {
	db := setup(t)
	defer db.teardown(t)

	// create new app
	app := fiber.New()

	userHandler := NewUserHandler(db.UserStore)

	app.Post("/", userHandler.HandleCreateUser)

	params := types.CreateUserParams{
		Email:     "james@james.com",
		FirstName: "James",
		LastName:  "John",
		Password:  "newPassword",
	}

	b, _ := json.Marshal(params)

	req := httptest.NewRequest("POST", "/", bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")
	resp, err := app.Test(req)

	if err != nil {
		t.Error(err)
	}

	var response responseStruct

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		t.Fatalf("Error decoding response: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK; got %s", resp.Status)
	}

	if response.Message != "success" {
		t.Errorf("Expected message 'User created successfully'; got %s", response.Message)
	}

	if response.User.FirstName != params.FirstName {
		t.Errorf("Expected first name %s; got %s", params.FirstName, response.User.FirstName)
	}

	if response.User.LastName != params.LastName {
		t.Errorf("Expected last name %s; got %s", params.LastName, response.User.LastName)
	}

	if response.User.Email != params.Email {
		t.Errorf("Expected email %s; got %s", params.Email, response.User.Email)
	}

}
