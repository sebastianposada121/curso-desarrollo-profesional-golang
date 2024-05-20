package jwt

import (
	"crud-echo-x-mongo-driver/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func GenerateJwt(user models.User) (string, error) {
	errorGodotenv := godotenv.Load()
	if errorGodotenv != nil {
		panic("Error loading .env file")
	}
	secret := []byte(os.Getenv("SECRET_KEY"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"name":  user.Name,
		"id":    user.Id,
		"iat":   time.Now().Unix(),
		"exp":   time.Now().Add(time.Hour * 25).Unix(),
	})

	tokenString, err := token.SignedString(secret)

	return tokenString, err

}
