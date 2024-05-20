package routes

import (
	"context"
	"crud-movies-gin-x-bun/database"
	"crud-movies-gin-x-bun/dto"
	"crud-movies-gin-x-bun/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCategories(ctx *gin.Context) {
	// set header
	// ctx.Writer.Header().Set("seb", "www.seb.co")

	categories := models.Categories{}

	if err := database.Connect.NewSelect().Model(&categories).Scan(context.TODO()); err != nil {
		ctx.JSON(http.StatusBadRequest, models.GenericResponse{
			Message:    "Error response",
			Successful: false,
		})
		return
	}

	ctx.JSON(http.StatusOK, models.GenericResponse{
		Message:    "Success response",
		Successful: true,
		Data:       categories,
	})
}

func GetCategory(ctx *gin.Context) {
	// ---- obtemer params
	// ctx.Param("id")
	// ---- obtener query
	// ctx.Query("id")
	id := ctx.Param("id")
	category, err := CategoryById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.GenericResponse{
			Message:    "No found",
			Successful: false,
		})
		return
	}

	ctx.JSON(http.StatusOK, models.GenericResponse{
		Message:    "Success response",
		Successful: true,
		Data:       category,
	})
}

func CreateCategory(ctx *gin.Context) {
	var category dto.Category

	// obtener header
	// auth := ctx.Request.Header.Get("auth")

	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, models.GenericResponse{
			Successful: false,
			Message:    "Error create",
		})
		return
	}

	result, err := database.Connect.NewInsert().Model(&category).Exec(context.TODO())

	id, errorId := result.LastInsertId()
	if err != nil || errorId != nil {
		ctx.JSON(http.StatusBadRequest, models.GenericResponse{
			Message:    "Error response",
			Successful: false,
		})
		return
	}

	categoryData, err := CategoryById(strconv.FormatInt(int64(int(id)), 10))

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
		Data:       categoryData,
	})
}

func UpdateCategory(ctx *gin.Context) {
	var category dto.Category
	id := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&category); err != nil || id == "" {
		ctx.JSON(http.StatusBadRequest, models.GenericResponse{
			Successful: false,
			Message:    "Error update",
		})
		return
	}

	if _, err := database.Connect.NewUpdate().Where(`id=?`, id).Model(&category).Exec(context.TODO()); err != nil {
		ctx.JSON(http.StatusBadRequest, models.GenericResponse{
			Successful: false,
			Message:    "Error update",
		})
		return
	}

	categoryData, err := CategoryById(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.GenericResponse{
			Message:    "Error update",
			Successful: false,
		})
		return
	}

	ctx.JSON(http.StatusOK, models.GenericResponse{
		Message:    "Success create",
		Successful: true,
		Data:       categoryData,
	})
}

func DeleteCategory(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, models.GenericResponse{
			Message:    "Error delete",
			Successful: false,
		})
		return
	}

	if _, err := database.Connect.NewDelete().Where(`id=?`, id).Model(&models.Category{}).Exec(context.TODO()); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, models.GenericResponse{
			Message:    "Error delete",
			Successful: false,
		})
		return
	}

	ctx.JSON(http.StatusAccepted, models.GenericResponse{
		Successful: true,
		Message:    "Success delete",
	})
}

func CategoryById(id string) (models.Category, error) {
	var category models.Category

	if err := database.Connect.NewSelect().Where(`id=?`, id).Model(&category).Scan(context.TODO()); err != nil {
		return category, fmt.Errorf("user with id %s not found", err)
	}

	return category, nil
}
