package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	Name        string             `json:"name"`
	Price       int                `json:"price"`
	Stock       int                `json:"stock"`
	Description string             `json:"description"`
	Category    Category           `json:"category"`
	CategoryId  primitive.ObjectID `json:"category_id" bson:"category_id"`
	Photos      []Photo            `json:"photos"`
}

type GenericResponse struct {
	Successful bool   `json:"successful"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
}

type Category struct {
	Id   string `json:"id" bson:"_id"`
	Name string `json:"name"`
}

type Categories []Category

type Products []Product

type Photo struct {
	Id   primitive.ObjectID `json:"id" bson:"_id"`
	Name string             `json:"name"`
}

type User struct {
	Id       primitive.ObjectID `json:"id" bson:"_id"`
	Name     string             `json:"name"`
	Phone    string             `json:"phone"`
	Email    string             `json:"email"`
	Password string             `json:"password"`
}
