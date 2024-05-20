package routes

import (
	"context"
	"crud-chi-x-rel/database"
	"crud-chi-x-rel/models"
	"crud-chi-x-rel/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-rel/rel/where"
)

func GetTeams(w http.ResponseWriter, r *http.Request) {

	teams := models.Teams{}

	if err := database.Database().FindAll(context.TODO(), &teams); err != nil {
		utils.ReponseJson(w, http.StatusBadRequest, models.GenericResponse{
			Successful: false,
			Message:    "Error get",
		})
		return
	}

	utils.ReponseJson(w, http.StatusBadRequest, models.GenericResponse{
		Successful: true,
		Data:       teams,
	})

}

func GetTeam(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	team := models.Team{}
	query := where.Eq("id", id)

	if id == "" {
		utils.ReponseJson(w, http.StatusBadRequest, models.GenericResponse{
			Successful: false,
			Message:    "Error id",
		})
		return
	}

	if err := database.Database().Find(context.TODO(), &team, query); err != nil {
		utils.ReponseJson(w, http.StatusBadRequest, models.GenericResponse{
			Successful: false,
			Message:    "Not found",
		})
		return
	}

	utils.ReponseJson(w, http.StatusBadRequest, models.GenericResponse{
		Successful: true,
		Data:       team,
	})

}

func CreateTeam(w http.ResponseWriter, r *http.Request) {
	// id := chi.URLParam(r, "id")
	team := models.Team{}

	if err := json.NewDecoder(r.Body).Decode(&team); err != nil {
		fmt.Println(err)
		utils.ReponseJson(w, http.StatusBadRequest, models.GenericResponse{
			Successful: false,
			Message:    "Error create",
		})
		return
	}

	if err := database.Database().Insert(context.TODO(), &team); err != nil {
		fmt.Println(err)
		utils.ReponseJson(w, http.StatusBadRequest, models.GenericResponse{
			Successful: false,
			Message:    "Error create",
		})
		return
	}

	utils.ReponseJson(w, http.StatusAccepted, models.GenericResponse{
		Successful: true,
		Message:    "Success create",
	})

}

func DeleteTeam(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	team := models.Team{}
	if err != nil {
		utils.ReponseJson(w, http.StatusBadRequest, models.GenericResponse{
			Successful: false,
			Message:    "Error id",
		})
		return
	}
	team.Id = id
	if err := database.Database().Delete(context.TODO(), &team); err != nil {
		utils.ReponseJson(w, http.StatusBadRequest, models.GenericResponse{
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

func UpdateTeam(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	team := models.Team{}

	if err != nil {
		utils.ReponseJson(w, http.StatusBadRequest, models.GenericResponse{
			Successful: false,
			Message:    "Error id",
		})
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&team); err != nil {
		utils.ReponseJson(w, http.StatusBadRequest, models.GenericResponse{
			Successful: false,
			Message:    "Error update",
		})
		return
	}

	team.Id = id

	if err := database.Database().Update(context.TODO(), &team); err != nil {
		utils.ReponseJson(w, http.StatusBadRequest, models.GenericResponse{
			Successful: false,
			Message:    "Error update",
		})
		return
	}

	utils.ReponseJson(w, http.StatusAccepted, models.GenericResponse{
		Successful: true,
		Message:    "Success update",
	})
}
