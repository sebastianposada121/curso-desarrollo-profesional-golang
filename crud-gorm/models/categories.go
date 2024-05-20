package models

type Category struct {
	Id          uint   `json:"id"`
	Name        string `gorm:"type:varchar(100)" json:"name"`
	Description string `gorm:"type:varchar(2500)" json:"description"`
}

type Categories []Category

// func Migrations() {
// 	database.Database.AutoMigrate(&User{}, &Rol{})
// }
