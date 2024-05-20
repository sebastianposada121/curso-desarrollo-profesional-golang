package routes

import (
	"context"
	"crud-echo-x-mongo-driver/database"
	"crud-echo-x-mongo-driver/dto"
	"crud-echo-x-mongo-driver/jwt"
	"crud-echo-x-mongo-driver/models"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c echo.Context) error {
	user := new(dto.User)

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, models.GenericResponse{
			Successful: false,
			Message:    "Error register user",
			Data:       nil,
		})
	}

	if _, err := FindUser(user.Email); err == nil {
		return c.JSON(http.StatusBadRequest, models.GenericResponse{
			Successful: false,
			Message:    "Error register user",
			Data:       nil,
		})
	}

	cryptoPass, errGenerateHash := bcrypt.GenerateFromPassword([]byte(user.Password), 8)

	if errGenerateHash != nil {
		return c.JSON(http.StatusBadRequest, models.GenericResponse{
			Successful: false,
			Message:    "Error register user",
			Data:       nil,
		})
	}

	user.Password = string(cryptoPass)

	if _, err := database.UserCollection.InsertOne(context.TODO(), user); err != nil {
		return c.JSON(http.StatusBadRequest, models.GenericResponse{
			Successful: false,
			Message:    "Error register user",
			Data:       nil,
		})
	}

	return c.JSON(http.StatusAccepted, models.GenericResponse{
		Successful: true,
		Message:    "Success register user",
	})

}

func Login(c echo.Context) error {

	login := new(dto.Login)

	if err := c.Bind(&login); err != nil {
		return c.JSON(http.StatusBadRequest, models.GenericResponse{
			Successful: false,
			Message:    "error login",
		})
	}

	user, err := FindUser(login.Email)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, models.GenericResponse{
			Successful: false,
			Message:    "error login",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, models.GenericResponse{
			Successful: false,
			Message:    "error login",
		})
	}

	token, errGenerateToken := jwt.GenerateJwt(user)

	if errGenerateToken != nil {
		return c.JSON(http.StatusUnauthorized, models.GenericResponse{
			Successful: false,
			Message:    "error login",
		})
	}

	return c.JSON(http.StatusAccepted, models.GenericResponse{
		Successful: true,
		Message:    token,
	})
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
