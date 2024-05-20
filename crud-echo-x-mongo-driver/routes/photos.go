package routes

import (
	"context"
	"crud-echo-x-mongo-driver/database"
	"crud-echo-x-mongo-driver/models"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type photo struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

func UploadPhoto(c echo.Context, route string) (photo, error) {
	file, err := c.FormFile("photo")
	var photo photo

	if err != nil {
		return photo, err
	}

	src, err := file.Open()
	if err != nil {
		return photo, err
	}

	defer src.Close()

	var ext = strings.Split(file.Filename, ".")[1]
	time := strings.Split(time.Now().String(), " ")
	photo.Name = string(time[4][6:14]) + "." + ext
	photo.Path = "public/photos/" + route + "/" + photo.Name

	dst, err := os.Create(photo.Path)

	if err != nil {
		return photo, err
	}

	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return photo, err
	}

	return photo, nil
}

func DeletePhoto(c echo.Context) error {
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))

	if result, err := database.PhotoCollection.DeleteOne(context.TODO(), bson.M{"_id": id}); err != nil || result.DeletedCount == 0 {
		return c.JSON(http.StatusBadRequest, models.GenericResponse{
			Successful: false,
			Message:    "photo not found",
			Data:       nil,
		})
	}

	return c.JSON(http.StatusAccepted, models.GenericResponse{
		Successful: true,
		Message:    "Success delete photo",
		Data:       nil,
	})
}
