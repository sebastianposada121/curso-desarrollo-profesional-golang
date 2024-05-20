package routes

import (
	"net/http"
	"text/template"
	"web/utils"
)

func NotFound(response http.ResponseWriter, request *http.Request) {
	// al usar must no necesitamos validar si la plantilla existe
	template := template.Must(template.ParseFiles("templates/not-found.html", utils.PATH_TEMPLATE))

	template.Execute(response, nil)
}
