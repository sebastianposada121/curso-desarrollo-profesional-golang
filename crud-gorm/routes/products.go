package routes

import (
	"crud-gorm/database"
	"crud-gorm/models"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(models.GenericResponse{
			Successful: false,
			Message:    "Could not create",
		})
		return
	}

	if database.Database.Save(&product).Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.GenericResponse{
			Successful: true,
			Message:    "Error created",
		})
	} else {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(models.GenericResponse{
			Successful: true,
			Message:    "successfully created",
		})
	}

}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data := models.Products{}
	database.Database.Order("id asc").Preload("Category").Find(&data)
	json.Marshal(&data)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	w.Header().Set("Content-Type", "application/json")
	data := models.Product{}
	if database.Database.Where("id = ?", id).Preload("Category").First(&data).Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.GenericResponse{
			Successful: true,
			Message:    "Not found product",
		})
	} else {
		json.Marshal(&data)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data)
	}
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var data models.Product
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(models.GenericResponse{
			Successful: false,
			Message:    "Could not create",
		})
		return
	}

	if database.Database.Where("id = ?", id).Updates(&data).Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.GenericResponse{
			Successful: true,
			Message:    "Error update",
		})
	} else {
		json.Marshal(&data)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(models.GenericResponse{
			Successful: true,
			Message:    "Successful update",
		})
	}

}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	if database.Database.Where("id = ?", id).Delete(&models.Product{}).Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.GenericResponse{
			Successful: true,
			Message:    "Error delete",
		})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(models.GenericResponse{
			Successful: true,
			Message:    "Successful delete",
		})
	}
}

func UploadProductPhoto(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 10, 32)

	file, handler, errFile := r.FormFile("photo")
	var ext = strings.Split(handler.Filename, ".")[1]
	time := strings.Split(time.Now().String(), " ")
	name := string(time[4][6:14]) + "." + ext
	route := "public/products/" + name
	f, err := os.OpenFile(route, os.O_WRONLY|os.O_CREATE, 0777)

	if _, errCopy := io.Copy(f, file); errCopy != nil || err != nil || errFile != nil {
		w.WriteHeader(http.StatusBadRequest)

		output, _ := json.Marshal(models.GenericResponse{
			Successful: false,
			Message:    "Error :(",
		})

		fmt.Fprintln(w, string(output))
		return
	}

	data := models.ProductPhoto{
		Name:      name,
		ProductId: uint(id),
	}

	if database.Database.Save(&data).Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.GenericResponse{
			Successful: false,
			Message:    "Error created",
		})
	} else {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(models.GenericResponse{
			Successful: true,
			Message:    "successfully created",
		})
	}
}

func ProdoctPhotos(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	data := models.ProductPhotos{}
	database.Database.Where("product_id", id).Preload("Product").Find(&data)
	json.Marshal(&data)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func DeleteProductPhoto(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	data := models.ProductPhoto{}
	if database.Database.First(&data, id).Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.GenericResponse{
			Successful: false,
			Message:    "Error delete",
		})
	} else {
		fmt.Println(data.Name)
		if e := os.Remove("public/products/" + data.Name); e != nil {
			log.Fatal(e)
			return
		}

		database.Database.Where("id = ?", id).Delete(&data)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(models.GenericResponse{
			Successful: true,
			Message:    "Successful delete",
		})
	}
}

func ViewFile(w http.ResponseWriter, r *http.Request) {
	fileName := r.URL.Query().Get("name")

	if fileName == "" {
		w.WriteHeader(http.StatusBadRequest)

		output, _ := json.Marshal(models.GenericResponse{
			Successful: false,
			Message:    "Error :(",
		})

		fmt.Fprintln(w, string(output))
		return
	}

	OpenFile, errOpen := os.Open("public/products/" + fileName)

	if errOpen != nil {
		w.WriteHeader(http.StatusBadRequest)

		output, _ := json.Marshal(models.GenericResponse{
			Successful: false,
			Message:    "no found :(",
		})

		fmt.Fprintln(w, string(output))
		return
	}

	if _, errCopy := io.Copy(w, OpenFile); errCopy != nil {
		w.WriteHeader(http.StatusBadRequest)

		output, _ := json.Marshal(models.GenericResponse{
			Successful: false,
			Message:    "Error :(",
		})

		fmt.Fprintln(w, string(output))
		return
	}
}
