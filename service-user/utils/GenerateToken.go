package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("your-secret-key")

func RefreshToken(id string, name string) (string, error) {

	return GenerateToken(id, name, true)
}

func AccessToken(id string, name string) (string, error) {

	return GenerateToken(id, name, false)
}

func GenerateToken(id string, name string, isRefreshToken bool) (string, error) {

	claims := jwt.MapClaims{}

	claims["authorized"] = true
	claims["userid"] = id
	claims["user_name"] = name

	duration := time.Minute * 15 // Access token in 15 minutes

	if isRefreshToken {
		duration = time.Hour * 72 // Refresh token expires in 3 days
	}

	claims["exp"] = time.Now().Add(duration).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtKey)
}
