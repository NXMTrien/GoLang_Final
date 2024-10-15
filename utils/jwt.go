package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("your_secret_key")

func GenerateToken(userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 72).Unix(),
	})

	return token.SignedString(jwtKey)
}
