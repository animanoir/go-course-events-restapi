package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
)

const secretKey = "supersecret"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"email":  "",
		"userId": "",
		"exp":    time.Now().Add(time.Hour * 2).Unix(), // expiration token
	})

	return token.SignedString([]byte(secretKey))
}
