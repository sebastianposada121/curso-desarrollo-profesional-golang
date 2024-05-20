package routes

import (
	"context"
	"crud-movies-gin-x-bun/database"
	"crud-movies-gin-x-bun/dto"
	"crud-movies-gin-x-bun/middleware"
	"crud-movies-gin-x-bun/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Signup(ctx *gin.Context) {

	var signup dto.Signup

	if err := ctx.ShouldBindJSON(&signup); err != nil {
		ctx.JSON(http.StatusBadRequest, models.GenericResponse{
			Successful: false,
			Message:    "Error signup",
		})
		return
	}

	if err := database.Connect.NewSelect().Model(&models.User{}).Where(`email=?`, signup.Email).Scan(context.TODO()); err == nil {
		ctx.JSON(http.StatusBadRequest, models.GenericResponse{
			Successful: false,
			Message:    "Error signup",
		})
		return
	}

	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(signup.Password), 8)
	signup.Password = string(encryptedPassword)

	if _, err := database.Connect.NewInsert().Model(&signup).Exec(context.TODO()); err != nil {
		ctx.JSON(http.StatusBadRequest, models.GenericResponse{
			Successful: false,
			Message:    "Error signup",
		})
		return
	}

	ctx.JSON(http.StatusAccepted, models.GenericResponse{
		Successful: true,
		Message:    "Signup successful",
	})
}

func Login(ctx *gin.Context) {
	var login dto.Login

	if err := ctx.ShouldBindJSON(&login); err != nil {
		ctx.JSON(http.StatusBadRequest, models.GenericResponse{
			Successful: false,
			Message:    "Error login",
		})
		return
	}

	user, err := UserByEmail(login.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.GenericResponse{
			Successful: false,
			Message:    "Error login",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password)); err != nil {
		ctx.JSON(http.StatusBadRequest, models.GenericResponse{
			Successful: false,
			Message:    "Error login",
		})
		return
	}

	token, err := middleware.GenerateJwt(user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.GenericResponse{
			Successful: false,
			Message:    "Error login",
		})
		return
	}

	ctx.JSON(http.StatusAccepted, models.GenericResponse{
		Successful: true,
		Message:    "Success login",
		Data:       token,
	})
}

func UserByEmail(email string) (user models.User, err error) {
	if err := database.Connect.NewSelect().Model(&user).Where(`email=?`, email).Scan(context.TODO()); err != nil {
		return user, errors.New(err.Error())
	}
	return user, nil
}
