package types

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

const (
	cookieName = "refresh_token"
	cookiePath = "/"
)

type TokenPairs struct {
	Token        string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Auth struct {
	Issuer        string
	Audience      string
	TokenExpiry   time.Duration
	RefreshExpiry time.Duration
	CookieDomain  string
	CookiePath    string
	CookieName    string
}

type Claims struct {
	jwt.RegisteredClaims
}

func GenerateTokenPair(user *User) (TokenPairs, error) {

	claims := jwt.MapClaims{

		"sub":  fmt.Sprint(user.ID),
		"name": fmt.Sprintf("%s %s", user.FirstName, user.LastName),
		"aud":  os.Getenv("JWT_AUDIENCE"),
		"iss":  os.Getenv("JWT_ISSUER"),
		"iat":  time.Now().UTC().Unix(),
		"exp":  time.Now().UTC().Add(time.Hour * 4).Unix(),
		"typ":  "JWT",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := os.Getenv("JWT_SECRET")

	log.Println("secret: ", secret)

	signedAccessToken, err := token.SignedString([]byte(secret))

	if err != nil {
		return TokenPairs{}, err
	}

	refreshClaims := jwt.MapClaims{
		"sub": fmt.Sprint(user.ID),
		"iat": time.Now().UTC().Unix(),
		"exp": time.Now().UTC().Add(time.Hour * 24 * 7).Unix(),
	}

	signedRefreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(secret))

	if err != nil {
		return TokenPairs{}, err
	}

	return TokenPairs{
		Token:        signedAccessToken,
		RefreshToken: signedRefreshToken,
	}, nil
}

func GenerateRefreshCookie(refreshToken string) *fiber.Cookie {
	return &fiber.Cookie{
		Name:     cookieName,
		Value:    refreshToken,
		Path:     "/",
		Domain:   os.Getenv("JWT_COOKIE_DOMAIN"),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict",
		MaxAge:   int(time.Hour * 24),
	}
}

func GetExpiredRefreshCookie() *fiber.Cookie {
	return &fiber.Cookie{
		Name:     cookieName,
		Value:    "",
		Path:     cookiePath,
		Domain:   os.Getenv("JWT_COOKIE_DOMAIN"),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict",
		Expires:  time.Now().UTC().Add(time.Hour * -24),
		MaxAge:   -1,
	}
}
