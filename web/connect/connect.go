package connect

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var Db *sql.DB

func Connect() {
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

	db, errorDb := sql.Open("mysql", DB_URI)
	if errorDb != nil {
		panic(errorDb)
	}

	Db = db
}

func CloseConection() {
	Db.Close()
}
