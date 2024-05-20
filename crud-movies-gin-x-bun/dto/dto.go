package dto

import "github.com/uptrace/bun"

type Movie struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Year        string `json:"year"`
	CategoryId  int64  `json:"category_id"`
}

type Category struct {
	Name string `json:"name"`
}

type PhotoMovie struct {
	Name    string `json:"name"`
	MovieId int64  `json:"movie_id"`
}

type Signup struct {
	bun.BaseModel `bun:"users"`
	Email         string `json:"email"`
	Name          string `json:"name"`
	Phone         string `json:"phone"`
	Password      string `json:"password"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
