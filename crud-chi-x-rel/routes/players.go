package routes

import (
	"context"
	"crud-chi-x-rel/database"
	"crud-chi-x-rel/models"
	"crud-chi-x-rel/utils"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-rel/rel"
	"github.com/go-rel/rel/where"
)

var ctx = context.TODO()

func GetPlayers(w http.ResponseWriter, r *http.Request) {
	var players []models.Player

	// Seleccionar jugadores junto con sus equipos

	// Crear la consulta para seleccionar jugadores y unir con la tabla de equipos
	query := rel.Select("*", "team.*").JoinAssoc("team")

	// Realizar la consulta
	if err := database.Database().FindAll(ctx, &players, query); err != nil {
		utils.ReponseJson(w, http.StatusBadRequest, models.GenericResponse{
			Successful: false,
			Message:    "Error getting players",
		})
		return
	}

	utils.ReponseJson(w, http.StatusAccepted, models.GenericResponse{
		Successful: true,
		Message:    "Get players",
		Data:       players,
	})
}

func CreatePlayer(w http.ResponseWriter, r *http.Request) {
	player := models.Player{}

	// if r.Header.Get("Authorization") != "paco" {
	// 	utils.ReponseJson(w, http.StatusUnauthorized, models.GenericResponse{
	// 		Successful: false,
	// 		Message:    "Unauthorized",
	// 	})
	// 	return
	// }

	if err := json.NewDecoder(r.Body).Decode(&player); err != nil {
		fmt.Println(err)
		utils.ReponseJson(w, http.StatusBadRequest, models.GenericResponse{
			Successful: false,
			Message:    "Error creating player",
		})
		return
	}

	fmt.Println(player)
	if err := database.Database().Insert(context.TODO(), &player); err != nil {
		utils.ReponseJson(w, http.StatusBadRequest, models.GenericResponse{
			Successful: false,
			Message:    "Error creating player",
		})
		return
	}

	utils.ReponseJson(w, http.StatusAccepted, models.GenericResponse{
		Successful: true,
		Message:    "Success player",
		Data:       player,
	})
}

func GetPlayer(w http.ResponseWriter, r *http.Request) {
	id, errorId := strconv.Atoi(chi.URLParam(r, "id"))

	player := models.Player{}

	if errorId != nil {
		utils.ReponseJson(w, http.StatusAccepted, models.GenericResponse{
			Successful: true,
			Message:    "Error id",
		})
		return
	}

	if err := database.Database().Find(ctx, &player, where.Eq("id", id)); err != nil {
		utils.ReponseJson(w, http.StatusAccepted, models.GenericResponse{
			Successful: false,
			Message:    "no found player",
		})
		return
	}

	utils.ReponseJson(w, http.StatusAccepted, models.GenericResponse{
		Successful: true,
		Message:    "Success get",
		Data:       player,
	})

}

func UpdatePlayer(w http.ResponseWriter, r *http.Request) {
	var player models.Player
	id, errorId := strconv.Atoi(chi.URLParam(r, "id"))

	if errorId != nil {
		utils.ReponseJson(w, http.StatusAccepted, models.GenericResponse{
			Successful: true,
			Message:    "Error id",
		})
		return
	}
	// if r.Header.Get("Authorization") != "paco" {
	// 	utils.ReponseJson(w, http.StatusUnauthorized, models.GenericResponse{
	// 		Successful: false,
	// 		Message:    "Unauthorized",
	// 	})
	// 	return
	// }

	if err := json.NewDecoder(r.Body).Decode(&player); err != nil {
		utils.ReponseJson(w, http.StatusBadRequest, models.GenericResponse{
			Successful: false,
			Message:    "Error update player",
		})
		return
	}

	player.Id = id

	if err := database.Database().Update(context.TODO(), &player); err != nil {
		fmt.Println(err)
		utils.ReponseJson(w, http.StatusBadRequest, models.GenericResponse{
			Successful: false,
			Message:    "Error update player",
		})
		return
	}

	utils.ReponseJson(w, http.StatusAccepted, models.GenericResponse{
		Successful: true,
		Message:    "Success update player",
		Data:       player,
	})
}

func DeletePlayer(w http.ResponseWriter, r *http.Request) {
	id, errorId := strconv.Atoi(chi.URLParam(r, "id"))

	if errorId != nil {
		utils.ReponseJson(w, http.StatusAccepted, models.GenericResponse{
			Successful: true,
			Message:    "Error id",
		})
		return
	}

	if err := database.Database().Delete(ctx, &models.Player{Id: id}); err != nil {
		utils.ReponseJson(w, http.StatusAccepted, models.GenericResponse{
			Successful: false,
			Message:    "Error delete",
		})
		return
	}

	utils.ReponseJson(w, http.StatusAccepted, models.GenericResponse{
		Successful: true,
		Message:    "Success delete",
	})
}

func GetPhotoByPlayer(w http.ResponseWriter, r *http.Request) {
	id, errorId := strconv.Atoi(chi.URLParam(r, "id"))
	photos := models.PlayerPhotos{}

	if errorId != nil {
		utils.ReponseJson(w, http.StatusAccepted, models.GenericResponse{
			Successful: false,
			Message:    "Error id",
		})
		return
	}

	if err := database.Database().FindAll(ctx, &photos, where.Eq("player_id", id)); err != nil {
		utils.ReponseJson(w, http.StatusAccepted, models.GenericResponse{
			Successful: false,
			Message:    "No found",
		})
		return
	}

	utils.ReponseJson(w, http.StatusAccepted, models.GenericResponse{
		Successful: true,
		Message:    "Success photos",
		Data:       photos,
	})
}

func DeletePhotPlayer(w http.ResponseWriter, r *http.Request) {
	id, errorId := strconv.Atoi(chi.URLParam(r, "id"))
	var photo models.PlayerPhoto
	// Crear la consulta para seleccionar jugadores y unir con la tabla de equipos

	if errorId != nil {
		utils.ReponseJson(w, http.StatusAccepted, models.GenericResponse{
			Successful: false,
			Message:    "Error id",
		})
		return
	}

	if err := database.Database().Find(ctx, &photo, where.Eq("id", id)); err != nil {
		utils.ReponseJson(w, http.StatusAccepted, models.GenericResponse{
			Successful: false,
			Message:    "Error delete",
		})
		return
	}

	if err := database.Database().Delete(ctx, &photo); err != nil {
		utils.ReponseJson(w, http.StatusAccepted, models.GenericResponse{
			Successful: false,
			Message:    "Error delete",
		})
		return
	}

	if err := os.Remove("public/photos/" + photo.Name); err != nil {
		log.Fatal(err)
	}

	utils.ReponseJson(w, http.StatusAccepted, models.GenericResponse{
		Successful: true,
		Message:    "Success delete",
	})
}

func UploadPlayerPhoto(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("photo")

	id, errorId := strconv.Atoi(chi.URLParam(r, "id"))

	if errorId != nil {
		utils.ReponseJson(w, http.StatusAccepted, models.GenericResponse{
			Successful: true,
			Message:    "Error id",
		})
		return
	}

	if err != nil {
		utils.ReponseJson(w, http.StatusUnauthorized, models.GenericResponse{
			Successful: false,
			Message:    "Unauthorized",
		})
		return
	}

	ext := strings.Split(handler.Filename, ".")[1]
	time := strings.Split(time.Now().String(), " ")
	name := string(time[4][6:14]) + "." + ext

	path := "public/photos/" + name

	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0777)

	if err != nil {
		utils.ReponseJson(w, http.StatusBadRequest, models.GenericResponse{
			Successful: false,
			Message:    "Error upload",
		})
		return
	}

	if _, err := io.Copy(f, file); err != nil {
		utils.ReponseJson(w, http.StatusBadRequest, models.GenericResponse{
			Successful: false,
			Message:    "Error upload",
		})
		return
	}

	photo := models.PlayerPhoto{
		Name:     name,
		PlayerId: id,
	}

	if err := database.Database().Insert(ctx, &photo); err != nil {
		utils.ReponseJson(w, http.StatusBadRequest, models.GenericResponse{
			Successful: false,
			Message:    "Error upload",
		})
		return
	}

	utils.ReponseJson(w, http.StatusOK, models.GenericResponse{
		Successful: true,
		Message:    "Success upload",
	})
}
