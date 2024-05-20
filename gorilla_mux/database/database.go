package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database = func() (db *gorm.DB) {

	errorGodotenv := godotenv.Load()

	if errorGodotenv != nil {
		panic(errorGodotenv)
	}

	DB_USER := os.Getenv("DB_USER")
	DB_NAME := os.Getenv("DB_NAME")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_SERVER := os.Getenv("DB_SERVER")
	DB_PORT := os.Getenv("DB_PORT")
	DB_URI := DB_USER + ":" + DB_PASSWORD + "@tcp(" + DB_SERVER + ":" + DB_PORT + ")/" + DB_NAME

	if db, errorDb := gorm.Open(mysql.Open(DB_URI), &gorm.Config{}); errorDb != nil {
		fmt.Println("error connection")
		panic(errorDb)
	} else {
		fmt.Println("sucess connection")
		return db
	}

}()
