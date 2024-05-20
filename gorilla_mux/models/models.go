package models

import "gorilla_mux/database"

type GenericResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type Category struct {
	Id           uint   `json:"id"`
	Name         string `gorm:"type:varchar(100)" json:"name"`
	Desctription string `gorm:"type:varchar(100)" json:"description"`
}

type Categories []Category

func Migrations() {
	database.Database.AutoMigrate(&Category{})
}
