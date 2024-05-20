package routes

import (
	"context"
	"crud-movies-gin-x-bun/database"
	"crud-movies-gin-x-bun/dto"
	"crud-movies-gin-x-bun/models"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func GetMovies(ctx *gin.Context) {
	ctx.Writer.Header().Set("seb", "www.seb.co")

	movies := models.Movies{}
	if err := database.Connect.NewSelect().Model(&movies).Relation("Category").Scan(context.TODO()); err != nil {
		ctx.JSON(http.StatusBadRequest, models.GenericResponse{
			Successful: false,
			Message:    "Error get",
		})
		return
	}
	ctx.JSON(http.StatusAccepted, models.GenericResponse{
		Successful: true,
		Message:    "Success movies",
		Data:       movies,
	})
}

func GetMovie(ctx *gin.Context) {
	// ---- obtemer params
	id := ctx.Param("id")
	// ---- obtener query
	// ctx.Query("id")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, models.GenericResponse{
			Successful: false,
			Message:    "Error id",
		})
		return
	}

	movie, err := MovieById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.GenericResponse{
			Successful: false,
			Message:    "no found",
		})
		return
	}

	ctx.JSON(http.StatusAccepted, models.GenericResponse{
		Successful: true,
		Message:    "Success movie",
		Data:       movie,
	})
}

func CreateMovie(ctx *gin.Context) {
	var movie dto.Movie

	// obtener header
	// auth := ctx.Request.Header.Get("auth")

	if err := ctx.ShouldBindJSON(&movie); err != nil {
		ctx.JSON(http.StatusBadRequest, models.GenericResponse{
			Successful: false,
			Message:    "Error create movie",
		})
		return
	}

	result, err := database.Connect.NewInsert().Model(&movie).Exec(context.TODO())

	id, errorId := result.LastInsertId()
	if err != nil || errorId != nil {
		ctx.JSON(http.StatusBadRequest, models.GenericResponse{
			Message:    "Error response",
			Successful: false,
		})
		return
	}

	movieData, err := MovieById(strconv.FormatInt(int64(int(id)), 10))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.GenericResponse{
			Message:    "Error response",
			Successful: false,
		})
		return
	}

	ctx.JSON(http.StatusOK, models.GenericResponse{
		Message:    "Success create",
		Successful: true,
		Data:       movieData,
	})
}

func UpdateMovie(ctx *gin.Context) {
	var movie dto.Movie
	id := ctx.Param("id")
	// obtener header
	// auth := ctx.Request.Header.Get("auth")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, models.GenericResponse{
			Successful: false,
			Message:    "Error id",
		})
		return
	}

	if err := ctx.ShouldBindJSON(&movie); err != nil {
		ctx.JSON(http.StatusBadRequest, models.GenericResponse{
			Successful: false,
			Message:    "Error update movie",
		})
		return
	}

	if _, err := database.Connect.NewUpdate().Where(`id=?`, id).Model(&movie).Exec(context.TODO()); err != nil {
		ctx.JSON(http.StatusBadRequest, models.GenericResponse{
			Successful: false,
			Message:    "Error update movie",
		})
		return
	}

	movieData, err := MovieById(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.GenericResponse{
			Message:    "Error response",
			Successful: false,
		})
		return
	}

	ctx.JSON(http.StatusOK, models.GenericResponse{
		Message:    "Success update",
		Successful: true,
		Data:       movieData,
	})
}

func DeleteMovie(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, models.GenericResponse{
			Successful: false,
			Message:    "Error id",
		})
		return
	}

	if _, err := database.Connect.NewDelete().Where(`id=?`, id).Model(&models.Movie{}).Exec(context.TODO()); err != nil {

		ctx.JSON(http.StatusBadRequest, models.GenericResponse{
			Successful: false,
			Message:    "Error delete",
		})
		return
	}

	ctx.JSON(http.StatusAccepted, models.GenericResponse{
		Successful: true,
		Message:    "Success delete",
	})
}

func UploadPhotoMovie(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.GenericResponse{
			Successful: false,
			Message:    "Error id",
		})
		return
	}
	name, err := UploadPhoto(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.GenericResponse{
			Successful: false,
			Message:    "Error upload",
		})
		return
	}

	photo := &dto.PhotoMovie{Name: name, MovieId: id}

	if _, err := database.Connect.NewInsert().Model(photo).Exec(context.TODO()); err != nil {
		ctx.JSON(http.StatusBadRequest, models.GenericResponse{
			Successful: false,
			Message:    "Error upload",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.GenericResponse{
		Successful: true,
		Message:    "Success upload",
	})
}

func GetPhotoByMovie(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, models.GenericResponse{
			Successful: false,
			Message:    "Error id",
		})
		return
	}

	if _, err := MovieById(id); err != nil {
		ctx.JSON(http.StatusBadRequest, models.GenericResponse{
			Successful: false,
			Message:    "No found movie",
		})
		return
	}

	photos := models.PhotoMovies{}

	if err := database.Connect.NewSelect().Model(&photos).Where("movie_id=?", id).Scan(context.TODO()); err != nil {
		ctx.JSON(http.StatusBadRequest, models.GenericResponse{
			Successful: false,
			Message:    "No photos",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.GenericResponse{
		Successful: true,
		Message:    "Success photos",
		Data:       photos,
	})
}

func DeletePhoto(ctx *gin.Context) {
	id := ctx.Param("id")
	var photo models.PhotoMovie

	if id == "" {
		ctx.JSON(http.StatusBadRequest, models.GenericResponse{
			Successful: false,
			Message:    "Error id",
		})
		return
	}

	if err := database.Connect.NewSelect().Where(`id=?`, id).Model(&photo).Scan(context.TODO()); err != nil {
		ctx.JSON(http.StatusBadRequest, models.GenericResponse{
			Successful: false,
			Message:    "No found photo",
		})
		return
	}

	if _, err := database.Connect.NewDelete().Model(&models.PhotoMovie{}).Where(`id=?`, id).Exec(context.TODO()); err != nil {
		ctx.JSON(http.StatusBadRequest, models.GenericResponse{
			Successful: false,
			Message:    "Error delete",
		})
	}

	delete := "public/photos/" + photo.Name

	if err := os.Remove(delete); err != nil {
		log.Fatal(err)
	}

	ctx.JSON(http.StatusOK, models.GenericResponse{
		Successful: true,
		Message:    "Success delete",
	})
}

func UploadPhoto(ctx *gin.Context) (string, error) {
	file, err := ctx.FormFile("photo")
	var ext = strings.Split(file.Filename, ".")[1]
	time := strings.Split(time.Now().String(), " ")
	name := string(time[4][6:14]) + "." + ext
	if err != nil {
		return "", err
	}

	ctx.SaveUploadedFile(file, "public/photos/"+name)
	return name, nil
}

func MovieById(id string) (models.Movie, error) {
	var movie models.Movie
	if err := database.Connect.NewSelect().Model(&movie).Where("id=?", id).Scan(context.TODO()); err != nil {
		return movie, fmt.Errorf("user with id %s not found", err)
	}
	category, _ := CategoryById(strconv.FormatInt(int64(int(movie.CategoryId)), 10))
	movie.Category = category
	return movie, nil
}
