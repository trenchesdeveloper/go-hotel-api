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
	"github.com/joho/godotenv"
	"github.com/trenchesdeveloper/go-hotel/db"
	"github.com/trenchesdeveloper/go-hotel/types"
)

type authResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func InsertTestUser(t *testing.T, userStore db.UserStore) *types.User {

	user, err := types.NewUserFromParams(types.CreateUserParams{
		FirstName: "John",
		LastName:  "james",
		Email:     "john@example.com",
		Password:  "password",
	})

	if err != nil {
		t.Fatal(err)
	}

	_, err = userStore.CreateUser(context.TODO(), user)
	if err != nil {
		t.Fatal(err)
	}

	return user
}

func TestHandleAuthenticate(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal(err)
	}
	db := setup(t)
	defer db.teardown(t)

	InsertTestUser(t, db.UserStore)

	// create new app
	app := fiber.New()

	authHandler := NewAuthHandler(db.UserStore)

	app.Post("/", authHandler.HandleAuthenticate)

	params := types.AuthParams{
		Email:    "john@example.com",
		Password: "password",
	}

	body, err := json.Marshal(params)
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(body))
	req.Header.Add("Content-Type", "application/json")
	resp, err := app.Test(req)

	if err != nil {
		t.Error(err)
	}

	var response authResponse

	err = json.NewDecoder(resp.Body).Decode(&response)

	if err != nil {
		t.Fatal(err)
	}

	log.Println("response", response)

	if response.AccessToken == "" {
		t.Error("access token is empty")
	}

	if response.RefreshToken == "" {
		t.Error("refresh token is empty")
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status code 200, got %d", resp.StatusCode)
	}

}
