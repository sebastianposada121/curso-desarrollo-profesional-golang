package routes

import (
	"context"
	"crud-echo-x-mongo-driver/database"
	"crud-echo-x-mongo-driver/dto"
	"crud-echo-x-mongo-driver/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateCategory(c echo.Context) error {
	category := new(dto.Category)

	if err := c.Bind(category); err != nil {
		return c.JSON(http.StatusBadRequest, models.GenericResponse{
			Successful: false,
			Message:    "Error creating category",
			Data:       nil,
		})
	}

	if len(category.Name) == 0 {
		return c.JSON(http.StatusBadRequest, models.GenericResponse{
			Successful: false,
			Message:    "Error creating category",
			Data:       nil,
		})
	}

	_, err := database.CategoryCollection.InsertOne(context.TODO(), category)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.GenericResponse{
			Successful: false,
			Message:    "Error creating category",
			Data:       nil,
		})
	}

	return c.JSON(http.StatusCreated, models.GenericResponse{
		Successful: true,
		Message:    "Category create with success",
		Data:       nil,
	})
}

func GetCategories(c echo.Context) error {
	var categories models.Categories
	coll, err := database.CategoryCollection.Find(context.TODO(), bson.D{})

	if err != nil {
		panic(err)
	}

	if err = coll.All(context.TODO(), &categories); err != nil {
		panic(err)
	}

	return c.JSON(http.StatusAccepted, models.GenericResponse{
		Successful: true,
		Message:    "Success categories",
		Data:       categories,
	})

}

func GetCategory(c echo.Context) error {
	var category models.Category
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	if err := database.CategoryCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&category); err != nil {
		return c.JSON(http.StatusBadRequest, models.GenericResponse{
			Successful: false,
			Message:    "category not found",
			Data:       nil,
		})
	}
	return c.JSON(http.StatusAccepted, models.GenericResponse{
		Successful: true,
		Message:    "Success category",
		Data:       category,
	})
}

func UpdateCategory(c echo.Context) error {
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	var category dto.Category

	if err := c.Bind(&category); err != nil {
		return c.JSON(http.StatusBadRequest, models.GenericResponse{
			Successful: false,
			Message:    "Error update category",
			Data:       nil,
		})
	}

	if len(category.Name) == 0 {
		return c.JSON(http.StatusBadRequest, models.GenericResponse{
			Successful: false,
			Message:    "Error update category",
			Data:       nil,
		})
	}

	if result, err := database.CategoryCollection.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": category}); err != nil || result.MatchedCount == 0 {
		return c.JSON(http.StatusBadRequest, models.GenericResponse{
			Successful: false,
			Message:    "category not found",
			Data:       nil,
		})
	}

	return c.JSON(http.StatusAccepted, models.GenericResponse{
		Successful: true,
		Message:    "Success update category",
		Data: models.Category{
			Id:   id.Hex(),
			Name: category.Name,
		},
	})
}

func DeleteCategory(c echo.Context) error {
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	if result, err := database.CategoryCollection.DeleteOne(context.TODO(), bson.M{"_id": id}); err != nil || result.DeletedCount == 0 {
		return c.JSON(http.StatusBadRequest, models.GenericResponse{
			Successful: false,
			Message:    "category not found",
			Data:       nil,
		})
	}
	return c.JSON(http.StatusAccepted, models.GenericResponse{
		Successful: true,
		Message:    "Success delete category",
		Data:       nil,
	})
}
