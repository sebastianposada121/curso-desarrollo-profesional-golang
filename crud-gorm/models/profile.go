package models

type Rol struct {
	Id   uint `json:"id"`
	Name uint `gorm:"type:varchar(100)" json:"name"`
}

type Rols []Rol

type User struct {
	Id       uint   `json:"id"`
	RolId    uint   `json:"rol_id"`
	Rol      Rol    `json:"rol"`
	Name     string `gorm:"type:varchar(100)" json:"name"`
	Email    string `gorm:"type:varchar(100)" json:"email"`
	Phone    string `gorm:"type:varchar(100)" json:"phone"`
	Password string `gorm:"type:varchar(100)" json:"password"`
}
