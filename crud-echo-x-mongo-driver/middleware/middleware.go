package middleware

import (
	"context"
	"crud-echo-x-mongo-driver/database"
	"crud-echo-x-mongo-driver/models"
	"fmt"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/joho/godotenv"
)

func ValidateJwt(c echo.Context) bool {
	errorGodotenv := godotenv.Load()
	if errorGodotenv != nil {
		panic("Error loading .env file")
	}
	secret := []byte(os.Getenv("SECRET_KEY"))

	authorizedToken := c.Request().Header.Get("Authorization")
	splitBearer := strings.Split(authorizedToken, " ")

	// validar que el token sea bearer y exista
	if len(authorizedToken) == 0 || len(splitBearer) != 2 {
		return false
	}

	// obtener token
	splitToken := strings.Split(splitBearer[1], ".")

	if len(splitToken) != 3 {
		return false
	}

	tk := strings.TrimSpace(splitBearer[1])

	token, err := jwt.Parse(tk, func(t *jwt.Token) (interface{}, error) {

		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {

			return nil, fmt.Errorf("Unexpect signing method: ")
		}

		return secret, nil
	})

	if err != nil {
		return false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		email := claims["email"].(string)
		_, err := FindUser(email)

		if err != nil {

			return false

		} else {
			return true
		}
	} else {
		return false
	}

}

func FindUser(email string) (models.User, error) {
	var user models.User
	if err := database.UserCollection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			return user, fmt.Errorf("user with email %s not found", email)
		}
		return user, err
	}
	return user, nil
}
