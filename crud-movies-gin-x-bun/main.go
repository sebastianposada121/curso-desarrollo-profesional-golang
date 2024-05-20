package main

import (
	"crud-movies-gin-x-bun/middleware"
	"crud-movies-gin-x-bun/routes"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	baseUrl := os.Getenv("BASE_URL")
	port := os.Getenv("PORT")

	router := gin.Default()
	router.Static("/public", "./public")

	// login
	router.POST(baseUrl+"signup", routes.Signup)
	router.POST(baseUrl+"login", routes.Login)

	// movies
	router.GET(baseUrl+"movies", middleware.GuardAuth(), routes.GetMovies)
	router.GET(baseUrl+"movies/:id", routes.GetMovie)
	router.POST(baseUrl+"movies", routes.CreateMovie)
	router.PUT(baseUrl+"movies/:id", routes.UpdateMovie)
	router.DELETE(baseUrl+"movies/:id", routes.DeleteMovie)
	router.POST(baseUrl+"movies/photo/:id", routes.UploadPhotoMovie)
	router.GET(baseUrl+"movies/photo/:id", routes.GetPhotoByMovie)
	router.DELETE(baseUrl+"movies/photo/:id", routes.DeletePhoto)

	// categories
	router.GET(baseUrl+"categories", routes.GetCategories)
	router.GET(baseUrl+"categories/:id", routes.GetCategory)
	router.POST(baseUrl+"categories", routes.CreateCategory)
	router.PUT(baseUrl+"categories/:id", routes.UpdateCategory)
	router.DELETE(baseUrl+"categories/:id", routes.DeleteCategory)

	router.Run(":" + port)

}
