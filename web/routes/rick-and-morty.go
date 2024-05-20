package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"text/template"
	"web/models"
	"web/utils"
)

/*
consumir rick and morty api
*/
func RickAndMorty(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("templates/rick-and-morty.html", utils.PATH_TEMPLATE))

	var characters, err = getCharacters()

	if err != nil {
		fmt.Println("error api")
	}

	// retornar json caracteres json.NewEncoder(response).Encode(characters)
	template.Execute(response, characters)
}

/*
Consumir api externa
*/
func getCharacters() (*models.Response, error) {
	var data models.Response
	response, error := http.Get("https://rickandmortyapi.com/api/character")

	if error != nil {
		fmt.Println(error)
	}

	// cerrar conexion al final del todo
	defer response.Body.Close()

	// Leer el cuerpo de la respuesta
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	// Decodificar el cuerpo de la respuesta
	// --> Unmarshal convierte a json
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println(err)
	}

	fmt.Println(&data)
	return &data, nil
}
