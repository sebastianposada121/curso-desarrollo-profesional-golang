package main

import (
	"crud-echo-x-mongo-driver/routes"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var prefix = "/api/v1"

func main() {
	e := echo.New()
	// static files
	e.Static(prefix+"/public", "public")

	// Connect DB

	// photos
	e.DELETE(prefix+"/photos/:id", routes.DeletePhoto)

	// database.ConnectDB()

	// --- users
	e.POST(prefix+"/signup", routes.Signup)
	e.POST(prefix+"/login", routes.Login)
	// ----
	e.GET(prefix+"/products", routes.GetProducts)
	e.GET(prefix+"/products/:id", routes.GetProduct)
	e.POST(prefix+"/products", routes.CreateProduct)
	e.PUT(prefix+"/products/:id", routes.UpdateProduct)
	e.DELETE(prefix+"/products/:id", routes.DeleteProduct)
	e.POST(prefix+"/products/photos/:id", routes.UploadProductPhoto)

	// --- category
	e.POST(prefix+"/categories", routes.CreateCategory)
	e.GET(prefix+"/categories", routes.GetCategories)
	e.GET(prefix+"/categories/:id", routes.GetCategory)
	e.DELETE(prefix+"/categories/:id", routes.DeleteCategory)
	e.PUT(prefix+"/categories/:id", routes.UpdateCategory)

	errorVariables := godotenv.Load()

	if errorVariables != nil {
		panic(errorVariables)
	}

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
