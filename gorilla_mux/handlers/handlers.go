package handlers

import (
	"encoding/json"
	"fmt"
	"gorilla_mux/dto"
	"gorilla_mux/models"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func Get(response http.ResponseWriter, request *http.Request) {
	response = setHeaders(response)
	params := mux.Vars(request)

	// -------get params
	// params := mux.Vars(request)
	// params["id"] params["name"]
	// -------query params
	// request.URL.Query().Get("status");
	// status := request.URL.Query().Get("id")

	response.Header().Set("Content-Type", "application/json")
	// json.Marshal codificación y decodificación de JSON
	output, _ := json.Marshal(models.GenericResponse{
		Status:  true,
		Message: params["id"],
	})

	fmt.Fprintln(response, string(output))
}

func Post(response http.ResponseWriter, request *http.Request) {
	response = setHeaders(response)

	auth := request.Header.Get("Authorization")

	if auth == "" {
		response.WriteHeader(http.StatusForbidden)

		output, _ := json.Marshal(models.GenericResponse{
			Status:  false,
			Message: "Denied",
		})

		fmt.Fprintln(response, string(output))
		return
	}

	var category dto.Category

	errCategoryMarshal := json.NewDecoder(request.Body).Decode(&category)

	if errCategoryMarshal != nil {
		response.WriteHeader(http.StatusBadRequest)

		output, _ := json.Marshal(models.GenericResponse{
			Status:  false,
			Message: "Error :(",
		})

		fmt.Fprintln(response, string(output))
		return
	}

	output, _ := json.Marshal(models.GenericResponse{
		Status:  true,
		Message: "created with successful " + category.Name,
	})

	// crear estados
	response.WriteHeader(http.StatusCreated)
	fmt.Fprintln(response, string(output))
}

func Upload(response http.ResponseWriter, request *http.Request) {
	response = setHeaders(response)

	file, handler, errFile := request.FormFile("photo")
	var ext = strings.Split(handler.Filename, ".")[1]
	time := strings.Split(time.Now().String(), " ")
	photo := string(time[4][6:14]) + "." + ext
	route := "public/upload/" + photo
	f, err := os.OpenFile(route, os.O_WRONLY|os.O_CREATE, 0777)

	if _, errCopy := io.Copy(f, file); errCopy != nil || err != nil || errFile != nil {
		response.WriteHeader(http.StatusBadRequest)

		output, _ := json.Marshal(models.GenericResponse{
			Status:  false,
			Message: "Error :(",
		})

		fmt.Fprintln(response, string(output))
		return
	}

	output, _ := json.Marshal(models.GenericResponse{
		Status:  true,
		Message: "subida con exito",
	})

	// crear estados
	response.WriteHeader(http.StatusCreated)

	fmt.Fprintln(response, string(output))
}

func ViewFile(w http.ResponseWriter, r *http.Request) {
	fileName := r.URL.Query().Get("name")

	fmt.Println(fileName)

	if fileName == "" {
		w.WriteHeader(http.StatusBadRequest)

		output, _ := json.Marshal(models.GenericResponse{
			Status:  false,
			Message: "Error :(",
		})

		fmt.Fprintln(w, string(output))
		return
	}

	OpenFile, errOpen := os.Open("public/upload/" + fileName)

	if errOpen != nil {
		w.WriteHeader(http.StatusBadRequest)

		output, _ := json.Marshal(models.GenericResponse{
			Status:  false,
			Message: "no found :(",
		})

		fmt.Fprintln(w, string(output))
		return
	}

	if _, errCopy := io.Copy(w, OpenFile); errCopy != nil {
		w.WriteHeader(http.StatusBadRequest)

		output, _ := json.Marshal(models.GenericResponse{
			Status:  false,
			Message: "Error :(",
		})

		fmt.Fprintln(w, string(output))
		return
	}
}

func Put(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)

	fmt.Fprintln(response, "this put", params)
}

func Delete(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)

	fmt.Fprintln(response, "this delete", params)
}

func setHeaders(response http.ResponseWriter) http.ResponseWriter {
	response.Header().Set("Content-Type", "application/json")

	return response
}
