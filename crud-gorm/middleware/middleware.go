package middleware

import (
	"crud-gorm/database"
	"crud-gorm/models"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"

	"github.com/joho/godotenv"
)

func ValidateJwt(next http.HandlerFunc) http.HandlerFunc {
	errorGodotenv := godotenv.Load()
	if errorGodotenv != nil {
		panic("Error loading .env file")
	}
	secret := []byte(os.Getenv("SECRET_KEY"))

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		authorizedToken := r.Header.Get("Authorization")
		splitBearer := strings.Split(authorizedToken, " ")
		

		// validar que el token sea bearer y exista
		if len(authorizedToken) == 0 || len(splitBearer) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(models.GenericResponse{
				Successful: false,
				Message:    "Not authorized",
			})
			return
		}

		// obtener token
		splitToken := strings.Split(splitBearer[1], ".")

		if len(splitToken) != 3 {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(models.GenericResponse{
				Successful: false,
				Message:    "Not authorized",
			})
			return
		}

		tk := strings.TrimSpace(splitBearer[1])

		token, err := jwt.Parse(tk, func(t *jwt.Token) (interface{}, error) {

			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {

				return nil, fmt.Errorf("Unexpect signing method: ")
			}

			return secret, nil
		})

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(models.GenericResponse{
				Successful: false,
				Message:    "Not authorized",
			})
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			user := models.User{}

			if err := database.Database.Where("email = ?", claims["email"]).First(&user); err.Error != nil {

				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(models.GenericResponse{
					Successful: false,
					Message:    "Not authorized",
				})

			} else {
				next.ServeHTTP(w, r)
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(models.GenericResponse{
				Successful: false,
				Message:    "Not authorized",
			})
		}
	}

}
