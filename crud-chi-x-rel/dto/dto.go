package dto

type Player struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	TeamId      int    `json:"teamId"`
}

type Team struct {
	Name string `json:"name"`
}
