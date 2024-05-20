package main

import (
	"crud-gorm/middleware"
	"crud-gorm/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	prefix := "/api/crud-gorm/"

	mux := mux.NewRouter()

	// categories
	mux.HandleFunc(prefix+"categories", routes.GetCategories).Methods("GET")
	mux.HandleFunc(prefix+"category/{id:[0-9]+}", routes.GetCategory).Methods("GET")
	mux.HandleFunc(prefix+"category", routes.CreateCategory).Methods("POST")
	mux.HandleFunc(prefix+"category/{id:[0-9]+}", routes.UpdateCategory).Methods("PUT")
	mux.HandleFunc(prefix+"category/{id:[0-9]+}", routes.DeleteCategory).Methods("DELETE")

	// products
	mux.HandleFunc(prefix+"products", routes.CreateProduct).Methods("POST")
	mux.HandleFunc(prefix+"products", middleware.ValidateJwt(routes.GetProducts)).Methods("GET")
	mux.HandleFunc(prefix+"products/{id:[0-9]+}", routes.GetProduct).Methods("GET")
	mux.HandleFunc(prefix+"products/{id:[0-9]+}", routes.UpdateProduct).Methods("PUT")
	mux.HandleFunc(prefix+"products/{id:[0-9]+}", routes.DeleteProduct).Methods("DELETE")
	mux.HandleFunc(prefix+"products-photo/{id:[0-9]+}", routes.UploadProductPhoto).Methods("POST")
	mux.HandleFunc(prefix+"products-photo/{id:[0-9]+}", routes.ProdoctPhotos).Methods("GET")
	mux.HandleFunc(prefix+"products-photo/{id:[0-9]+}", routes.DeleteProductPhoto).Methods("DELETE")
	mux.HandleFunc(prefix+"view-photo/{name:*}", routes.ProdoctPhotos).Methods("GET")

	// auth
	mux.HandleFunc(prefix+"signup", routes.Signup).Methods("POST")
	mux.HandleFunc(prefix+"login", routes.Login).Methods("POST")

	// cors
	corsMux := cors.AllowAll().Handler(mux)
	log.Fatal(http.ListenAndServe(":8082", corsMux))
}
