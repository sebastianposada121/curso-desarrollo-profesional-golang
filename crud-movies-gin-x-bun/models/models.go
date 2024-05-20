package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Category struct {
	bun.BaseModel `bun:"table:categories"`
	Id            int64  `bun:",pk,autoincrement" json:"id"`
	Name          string `bun:"name,notnull" json:"name"`
}

type Categories []Category

type GenericResponse struct {
	Successful bool        `json:"successful"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

type Movie struct {
	bun.BaseModel `bun:"table:movies"`
	Id            int64     `bun:",pk,autoincrement" json:"id"`
	Name          string    `bun:"name,notnull" json:"name"`
	Description   string    `bun:"description,notnull" json:"description"`
	Year          time.Time `bun:"year,nullzero,notnull" json:"year"`
	CategoryId    int64     `json:"category_id"`
	Category      Category  `bun:"rel:belongs-to,join:category_id=id" json:"category"`
}

type Movies []Movie

type PhotoMovie struct {
	bun.BaseModel `bun:"photo_movies"`
	Id            int64  `bun:",pk,autoincrement" json:"id"`
	Name          string `bun:"name,notnull" json:"name"`
	MovieId       int64  `json:"movie_id"`
}

type PhotoMovies []PhotoMovie

type User struct {
	bun.BaseModel `bun:"users"`
	Id            int64  `bun:",pk,autoincrement" json:"id"`
	Email         string `bun:"email,notnull" json:"email"`
	Name          string `bun:"name,notnull" json:"name"`
	Phone         string `bun:"phone,notnull" json:"phone"`
	Password      string `bun:"password,notnull" json:"password"`
}
