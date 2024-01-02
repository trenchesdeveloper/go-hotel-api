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
	log.Println("authHeader", authHeader)
	if !ok {
		return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
	}

	tokenString := authHeader[0]

	log.Println("tokenString", tokenString)

	tokenParts := strings.Split(tokenString, " ")

	if len(tokenParts) != 2 {
		return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
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
	log.Println("token2", token)

	if err != nil {
		if strings.HasPrefix(err.Error(), "token is expired") {
			log.Println("token expired")
			return fiber.NewError(fiber.StatusUnauthorized, "token expired")
		}
		log.Println("token error", err)
		return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
	}

	log.Println("token3", token)

	jwtIssuer := os.Getenv("JWT_ISSUER")

	if jwtIssuer == "" {
		jwtIssuer = "go-hotel"
	}

	// check the token is valid
	if claims.Issuer != os.Getenv("JWT_ISSUER"){
		log.Println("token4", token)
		return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
	}

	return c.Next()
}
