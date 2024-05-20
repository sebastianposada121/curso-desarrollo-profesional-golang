package models

type Product struct {
	Id          uint     `json:"id"`
	Name        string   `gorm:"type:varchar(100)" json:"name"`
	Price       int      `json:"price"`
	Stock       int      `json:"stock"`
	Description string   `json:"description"`
	CategoryId  uint     `json:"categoryId"`
	Category    Category `json:"category"`
}

type Products []Product

type ProductPhoto struct {
	Id        int     `json:"id"`
	Name      string  `gorm:"type:varchar(100)" json:"name"`
	ProductId uint    `json:"productId"`
	Product   Product `json:"product"`
}

type ProductPhotos []ProductPhoto
