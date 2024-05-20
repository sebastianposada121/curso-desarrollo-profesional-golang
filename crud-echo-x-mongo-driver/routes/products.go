package routes

import (
	"context"
	"crud-echo-x-mongo-driver/database"
	"crud-echo-x-mongo-driver/dto"
	"crud-echo-x-mongo-driver/middleware"
	"crud-echo-x-mongo-driver/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetProducts(c echo.Context) error {
	// page := c.QueryParam("page")
	// limit := c.QueryParam("limit")
	if !middleware.ValidateJwt(c) {
		return c.JSON(http.StatusUnauthorized, models.GenericResponse{
			Successful: false,
			Message:    "Unauthorized user",
		})
	}
	var products models.Products

	pipe := []bson.M{
		{"$lookup": bson.M{"from": "categories", "localField": "category_id", "foreignField": "_id", "as": "category"}},
		{"$addFields": bson.M{"category": bson.M{"$arrayElemAt": []interface{}{"$category", 0}}}},
		{"$lookup": bson.M{"from": "photos", "localField": "_id", "foreignField": "item_id", "as": "photos"}},
		{"$sort": bson.M{"_id": -1}},
	}

	coll, err := database.ProductCollection.Aggregate(context.TODO(), pipe)

	if err != nil {
		panic(err)
	}

	if err := coll.All(context.TODO(), &products); err != nil {
		return c.JSON(http.StatusBadRequest, models.GenericResponse{
			Successful: true,
			Message:    "Error get products",
		})
	}

	return c.JSON(http.StatusAccepted, models.GenericResponse{
		Successful: true,
		Message:    "Success products",
		Data:       products,
	})
}

func GetProduct(c echo.Context) error {
	// page := c.QueryParam("page")
	// limit := c.QueryParam("limit")
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))

	var products models.Products

	pipe := []bson.M{
		{"$match": bson.M{"_id": id}},
		{"$lookup": bson.M{"from": "categories", "localField": "category_id", "foreignField": "_id", "as": "category"}},
		{"$addFields": bson.M{"category": bson.M{"$arrayElemAt": []interface{}{"$category", 0}}}},
		{"$lookup": bson.M{"from": "photos", "localField": "_id", "foreignField": "item_id", "as": "photos"}},
		{"$sort": bson.M{"_id": -1}},
	}

	coll, err := database.ProductCollection.Aggregate(context.TODO(), pipe)

	if err != nil {
		panic(err)
	}

	if err := coll.All(context.TODO(), &products); err != nil {
		return c.JSON(http.StatusBadRequest, models.GenericResponse{
			Successful: true,
			Message:    "Product not found",
		})
	}

	return c.JSON(http.StatusAccepted, models.GenericResponse{
		Successful: true,
		Message:    "Success product",
		Data:       products[0],
	})
}

func CreateProduct(c echo.Context) error {

	// auth := c.Request().Header.Get("Authorization")

	product := new(dto.Product)
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, models.GenericResponse{
			Successful: false,
			Message:    "Error creating product",
			Data:       nil,
		})
	}

	if _, err := database.ProductCollection.InsertOne(context.TODO(), product); err != nil {
		return c.JSON(http.StatusBadRequest, models.GenericResponse{
			Successful: false,
			Message:    "Error creating product",
			Data:       nil,
		})
	}

	return c.JSON(http.StatusCreated, models.GenericResponse{
		Successful: true,
		Message:    "Created product success",
		Data:       product,
	})
}

func UpdateProduct(c echo.Context) error {
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))

	var product dto.Product

	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, models.GenericResponse{
			Successful: false,
			Message:    "Error update product",
			Data:       nil,
		})
	}

	if result, err := database.ProductCollection.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": product}); err != nil || result.MatchedCount == 0 {
		return c.JSON(http.StatusBadRequest, models.GenericResponse{
			Successful: false,
			Message:    "product not found",
			Data:       nil,
		})
	}

	return c.JSON(http.StatusAccepted, models.GenericResponse{
		Successful: true,
		Message:    "Success update product",
		Data:       nil,
	})
}

func DeleteProduct(c echo.Context) error {
	// obtener params
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))

	if result, err := database.ProductCollection.DeleteOne(context.TODO(), bson.M{"_id": id}); err != nil || result.DeletedCount == 0 {
		return c.JSON(http.StatusBadRequest, models.GenericResponse{
			Successful: false,
			Message:    "product not found",
			Data:       nil,
		})
	}

	return c.JSON(http.StatusAccepted, models.GenericResponse{
		Successful: true,
		Message:    "Success delete product",
		Data:       nil,
	})

}

func UploadProductPhoto(c echo.Context) error {
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))

	photo, err := UploadPhoto(c, "products")
	if err != nil || c.Param("id") == "" {
		return c.JSON(http.StatusBadRequest, models.GenericResponse{
			Successful: false,
			Message:    "Error upload photo",
		})
	}

	if _, err := database.PhotoCollection.InsertOne(context.TODO(), dto.Photo{
		ItemId: id,
		Name:   photo.Name,
		Path:   photo.Path,
	}); err != nil {
		return c.JSON(http.StatusBadRequest, models.GenericResponse{
			Successful: false,
			Message:    "Error upload photo",
		})
	}

	return c.JSON(http.StatusAccepted, models.GenericResponse{
		Successful: true,
		Message:    "Photo",
	})
}
