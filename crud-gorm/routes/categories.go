package routes

import (
	"crud-gorm/database"
	"crud-gorm/dto"
	"crud-gorm/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data := models.Categories{}

	database.Database.Order("id asc").Find(&data)

	json.Marshal(&data)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(data)
}

func GetCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	data := models.Category{}

	if err := database.Database.First(&data, id); err.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(models.GenericResponse{
			Successful: false,
			Message:    "No found category",
		})
	} else {
		json.Marshal(&data)
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(data)
	}
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category dto.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(models.GenericResponse{
			Successful: false,
			Message:    "Could not create",
		})
		return
	}

	data := models.Category{
		Name:        category.Name,
		Description: category.Description,
	}
	database.Database.Save(&data)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.GenericResponse{
		Successful: true,
		Message:    "successfully created",
	})

}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var category dto.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(models.GenericResponse{
			Successful: false,
			Message:    "Could not update",
		})
		return
	}

	data := models.Category{
		Name:        category.Name,
		Description: category.Description,
	}

	if database.Database.Where("id = ?", id).Updates(&data).Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.GenericResponse{
			Successful: false,
			Message:    "Could not update",
		})
	} else {
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(models.GenericResponse{
			Successful: true,
			Message:    "successfully updated",
		})
	}

}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	if database.Database.Where("id = ?", id).Delete(&models.Category{}).Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.GenericResponse{
			Successful: false,
			Message:    "Could not delete",
		})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(models.GenericResponse{
			Successful: false,
			Message:    "Successfuly delete",
		})
	}

}
