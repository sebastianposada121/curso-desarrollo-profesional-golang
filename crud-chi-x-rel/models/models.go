package models

type GenericResponse struct {
	Successful bool        `json:"successful"`
	Message    string      `json:"message"`
	Data       interface{} `jsn:"data"`
}

type Team struct {
	Id   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	// Player Player `ref:"team_id" fk:"id"`
	// Player Player `ref:"id" fk:"team_id"`
}

type Teams []Team

func (b Team) TeamsTable() string {
	return `teams`
}

type Player struct {
	Id          int    `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
	TeamId      int    `db:"team_id" json:"team_id"`
	Team        Team   `json:"team" rel:"belongs_to:team_id"`
}

type Players []Player

func (b Player) PlayerTables() string {
	return `players`
}

type PlayerPhoto struct {
	Id       int    `db:"id" json:"id"`
	Name     string `db:"name" json:"name"`
	PlayerId int    `db:"player_id" json:"player_id"`
}

type PlayerPhotos []PlayerPhoto

func (b PlayerPhoto) PlayerPhotoTable() string {
	return `player_photos`
}
