package api

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/trenchesdeveloper/go-hotel/db"
	"github.com/trenchesdeveloper/go-hotel/types"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	userStore db.UserStore
}

func NewAuthHandler(userStore db.UserStore) *AuthHandler {
	return &AuthHandler{
		userStore: userStore,
	}
}

func (a *AuthHandler) HandleAuthenticate(c *fiber.Ctx) error {
	fmt.Println("HandleAuthenticate")
	var params types.AuthParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}

	// get the user from the database
	user, err := a.userStore.GetUserByEmail(c.Context(), params.Email)

	// if the user doesn't exist, return an error
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return fiber.NewError(fiber.StatusUnauthorized, "invalid email or password")
		}
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password))

	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "invalid credentials")
	}

	return c.JSON(fiber.Map{
		"message": "success",
		"token":   params.Email,
	})
}
