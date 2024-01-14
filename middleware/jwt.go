package middleware

import (
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/trenchesdeveloper/go-hotel/db"
	"github.com/trenchesdeveloper/go-hotel/types"
)

func AuthRequired(store *db.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader, ok := c.GetReqHeaders()["Authorization"]
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "unauthorized",
			})
		}

		tokenString := authHeader[0]

		tokenParts := strings.Split(tokenString, " ")

		if len(tokenParts) != 2 {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "unauthorized",
			})
		}

		token := tokenParts[1]

		claims := &types.Claims{}

		_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			secret := os.Getenv("JWT_SECRET")
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unauthorized")
			}
			return []byte(secret), nil
		})

		if err != nil {
			if strings.HasPrefix(err.Error(), "token is expired") {
				return fiber.NewError(fiber.StatusUnauthorized, "token expired")
			}
			return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
		}

		jwtIssuer := os.Getenv("JWT_ISSUER")

		if jwtIssuer == "" {
			jwtIssuer = "go-hotel"
		}

		// check the token is valid
		if claims.Issuer != os.Getenv("JWT_ISSUER") {
			return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
		}

		// get the userId from the claims
		userId := claims.Subject

		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
		}

		// strip the ObjectId(") from the userId
		hexId := strings.TrimPrefix(userId, "ObjectID(\"")
		hexId = strings.TrimSuffix(hexId, "\")")
		// set the user in the context
		c.Locals("userID", hexId)

		// Use the store decorator to query user
		user, err := store.GetUserById(c.Context(), hexId)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "failed to query user")
		}

		// Set the user in the context
		c.Locals("user", user)

		return c.Next()
	}
}
