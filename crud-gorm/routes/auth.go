package routes

import (
	"crud-gorm/database"
	"crud-gorm/dto"
	"crud-gorm/jwt"
	"crud-gorm/models"
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	var body dto.User

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.GenericResponse{
			Successful: false,
			Message:    "Error user",
		})
		return
	}

	var count int64
	database.Database.Where("email = ?", body.Email).Find(&models.User{}).Limit(1).Count(&count)

	if count > 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.GenericResponse{
			Successful: false,
			Message:    "Error user",
		})
		return
	}

	cryptoPass, _ := bcrypt.GenerateFromPassword([]byte(body.Password), 8)

	user := models.User{
		RolId:    body.RolId,
		Name:     body.Name,
		Email:    body.Email,
		Phone:    body.Phone,
		Password: string(cryptoPass),
	}
	if database.Database.Save(&user).Error != nil {
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

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var body dto.Login
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.GenericResponse{
			Successful: false,
			Message:    "Not authorized",
		})
		return
	}

	var count int64
	var user models.User
	database.Database.Where("email = ?", body.Email).First(&user).Limit(1).Count(&count)

	if count == 0 {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.GenericResponse{
			Successful: false,
			Message:    "Not authorized",
		})
		return
	}

	passwordBody := []byte(body.Passsword)
	passwordUser := []byte(user.Password)

	err := bcrypt.CompareHashAndPassword(passwordUser, passwordBody)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.GenericResponse{
			Successful: false,
			Message:    "Not authorized",
		})
	} else {
		tokenJwt, err := jwt.GenerateJwt(user)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(models.GenericResponse{
				Successful: false,
				Message:    "Not authorized",
			})
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(models.GenericResponse{
			Successful: true,
			Message:    "Auth success",
			Data:       tokenJwt,
		})
	}
}
