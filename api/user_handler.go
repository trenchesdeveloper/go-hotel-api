package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trenchesdeveloper/go-hotel/types"
)

func HandleGetUsers(c *fiber.Ctx) error {
	user := types.User{
		FirstName: "John",
		LastName:  "Doe",
	}
	return c.JSON(user)
}

func HandleGetUser(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "a user",
	})
}
