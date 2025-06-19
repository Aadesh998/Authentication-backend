package utils

import (
	"time"

	"main/config"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(email, name string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"name":  name,
		"exp":   jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
	})
	return token.SignedString(config.JWTSecret)
}
