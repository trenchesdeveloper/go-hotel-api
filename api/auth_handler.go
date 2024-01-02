package api

import (
	"errors"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/trenchesdeveloper/go-hotel/db"
	"github.com/trenchesdeveloper/go-hotel/types"
	"go.mongodb.org/mongo-driver/mongo"
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

	if !types.IsValidPassword(user.Password, params.Password) {
		return fiber.NewError(fiber.StatusUnauthorized, "invalid email or password")
	}

	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "invalid credentials")
	}

	tokenPairs, err := types.GenerateTokenPair(user)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Something went wrong")
	}

	// set the refresh token as a cookie
	refreshCookie := types.GenerateRefreshCookie(tokenPairs.RefreshToken)

	// set the refreshCookie on the response
	c.Cookie(refreshCookie)

	log.Println("refresh cookie set")

	return c.Status(fiber.StatusOK).JSON(tokenPairs)
}
