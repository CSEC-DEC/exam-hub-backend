package helpers

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// ! HIDE THIS
var jwtKey = []byte("better_put_this_in_a_.env")

func GenerateJWT(username string) (string, error) {
	claims := jwt.RegisteredClaims{
		Subject:   username,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
