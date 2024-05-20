package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type Category struct {
	Name string `json:"name"`
}

type Product struct {
	Name        string             `json:"name"`
	Price       int                `json:"price"`
	Stock       int                `json:"stock"`
	Description string             `json:"description"`
	CategoryId  primitive.ObjectID `json:"category_id" bson:"category_id"`
}

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Name  string `json:"name"`
	Token string `json:"token"`
}

type Photo struct {
	ItemId primitive.ObjectID `json:"item_id" bson:"item_id"`
	Name   string             `json:"name"`
	Path   string             `json:"path"`
}
