package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

func GenerateJwtToken(email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
