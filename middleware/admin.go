package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trenchesdeveloper/go-hotel/types"
)

func AdminRequired(c *fiber.Ctx) error {
	user := c.Locals("user").(*types.User)
	if !user.IsAdmin {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "unauthorized",
		})
	}
	return c.Next()
}
