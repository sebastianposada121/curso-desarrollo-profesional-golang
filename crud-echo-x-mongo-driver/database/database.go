package database

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ClientMongo = ConnectDB()
var MongoDB = "echo_golang"
var CategoryCollection = ClientMongo.Database(MongoDB).Collection("categories")
var ProductCollection = ClientMongo.Database(MongoDB).Collection("products")
var PhotoCollection = ClientMongo.Database(MongoDB).Collection("photos")
var UserCollection = ClientMongo.Database(MongoDB).Collection("users")

func ConnectDB() *mongo.Client {
	errorVariables := godotenv.Load()

	if errorVariables != nil {
		panic(errorVariables)
	}

	var clientOptions = options.Client().ApplyURI(os.Getenv("DATABASE_URI"))

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("successful database connection 😊")

	return client
}
