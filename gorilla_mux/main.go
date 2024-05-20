package main

import (
	"gorilla_mux/handlers"
	"gorilla_mux/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	// migrar la bd

	models.Migrations()

	mux := mux.NewRouter()
	prefix := "/api/gorilla-mux/"


	mux.HandleFunc(prefix+"detail/{id:.*}", handlers.Get).Methods("GET")

	// 
	mux.HandleFunc(prefix+"detail/{id:.*}", handlers.Get).Methods("GET")
	mux.HandleFunc(prefix, handlers.Post).Methods("POST")
	mux.HandleFunc(prefix+"upload", handlers.Upload).Methods("POST")
	mux.HandleFunc(prefix+"file", handlers.ViewFile).Methods("GET")
	mux.HandleFunc(prefix, handlers.Delete).Methods("DELETE")
	mux.HandleFunc(prefix, handlers.Put).Methods("PUT")

	// cors
	corsMux := cors.AllowAll().Handler(mux)
	log.Fatal(http.ListenAndServe(":8082", corsMux))
}
