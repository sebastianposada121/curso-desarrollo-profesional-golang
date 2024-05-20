package middleware

import (
	"context"
	"crud-movies-gin-x-bun/database"
	"crud-movies-gin-x-bun/models"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
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

func GuardAuth() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		errorGodotenv := godotenv.Load()
		if errorGodotenv != nil {
			ctx.JSON(http.StatusUnauthorized, models.GenericResponse{
				Successful: false,
				Message:    "Unauthorized call admin",
			})
			return
		}
		secret := []byte(os.Getenv("SECRET_KEY"))

		authorizedToken := ctx.Request.Header.Get("Authorization")
		splitBearer := strings.Split(authorizedToken, " ")

		fmt.Println(authorizedToken)

		// validar que el token sea bearer y exista
		if len(authorizedToken) == 0 || len(splitBearer) != 2 {
			ctx.JSON(http.StatusUnauthorized, models.GenericResponse{
				Successful: false,
				Message:    "Unauthorized call admin",
			})
			return
		}

		// obtener token
		splitToken := strings.Split(splitBearer[1], ".")

		if len(splitToken) != 3 {
			ctx.JSON(http.StatusUnauthorized, models.GenericResponse{
				Successful: false,
				Message:    "Unauthorized call admin",
			})
			return
		}

		tk := strings.TrimSpace(splitBearer[1])

		token, err := jwt.Parse(tk, func(t *jwt.Token) (interface{}, error) {

			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpect signing method: ")
			}

			return secret, nil
		})

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, models.GenericResponse{
				Successful: false,
				Message:    "Unauthorized call admin",
			})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			email := claims["email"].(string)
			_, err := userByEmail(email)

			if err != nil {
				ctx.JSON(http.StatusUnauthorized, models.GenericResponse{
					Successful: false,
					Message:    "Unauthorized call admin",
				})

			} else {
				ctx.Next()
			}
		} else {
			ctx.JSON(http.StatusUnauthorized, models.GenericResponse{
				Successful: false,
				Message:    "Unauthorized call admin",
			})
		}
	}

}

func userByEmail(email string) (user models.User, err error) {
	if err := database.Connect.NewSelect().Model(&user).Where(`email=?`, email).Scan(context.TODO()); err != nil {
		return user, errors.New(err.Error())
	}
	return user, nil
}
