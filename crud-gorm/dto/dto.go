package dto

type Category struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Product struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	CategoryId  uint   `json:"categoryId"`
}

type User struct {
	// Id       uint   `json:"id"`
	RolId    uint   `json:"rol_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type Login struct {
	Email     string `json:"email"`
	Passsword string `json:"password"`
}
