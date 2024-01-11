package middleware

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/trenchesdeveloper/go-hotel/types"
)

func AuthRequired(c *fiber.Ctx) error {
	fmt.Println("jwt middleware", c.GetReqHeaders())
	authHeader, ok := c.GetReqHeaders()["Authorization"]
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "unauthorized",
		})
	}

	tokenString := authHeader[0]

	log.Println("tokenString", tokenString)

	tokenParts := strings.Split(tokenString, " ")

	if len(tokenParts) != 2 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "unauthorized",
		})
	}

	token := tokenParts[1]

	log.Println("token1", token)

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
			log.Println("token expired")
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
		log.Println("token5", token)
		return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
	}

	// strip the ObjectId(") from the userId
	hexId := strings.TrimPrefix(userId, "ObjectID(\"")
	hexId = strings.TrimSuffix(hexId, "\")")
	// set the user in the context
	c.Locals("userID", hexId)

	return c.Next()
}
