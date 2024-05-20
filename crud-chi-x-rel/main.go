package main

import (
	"crud-chi-x-rel/routes"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	prefix := "/api/v1"

	fs := http.FileServer(http.Dir("public"))

	// router
	router := chi.NewRouter()

	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	router.Use(middleware.Logger)

	router.Handle("/public/*", http.StripPrefix("/public/", fs))
	router.Get(prefix+"/players", routes.GetPlayers)
	router.Post(prefix+"/players", routes.CreatePlayer)
	router.Get(prefix+"/players/{id}", routes.GetPlayer)
	router.Put(prefix+"/players/{id}", routes.UpdatePlayer)
	router.Delete(prefix+"/players/{id}", routes.DeletePlayer)
	router.Post(prefix+"/players/photo/{id}", routes.UploadPlayerPhoto)
	router.Get(prefix+"/players/photo/{id}", routes.GetPhotoByPlayer)
	router.Delete(prefix+"/players/photo/{id}", routes.DeletePhotPlayer)

	// teams
	router.Get(prefix+"/teams", routes.GetTeams)
	router.Post(prefix+"/teams", routes.CreateTeam)
	router.Put(prefix+"/teams/{id}", routes.UpdateTeam)
	router.Delete(prefix+"/teams/{id}", routes.DeleteTeam)

	http.ListenAndServe(":"+os.Getenv("PORT"), router)
}
