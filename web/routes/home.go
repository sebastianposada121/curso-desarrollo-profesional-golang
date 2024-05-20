package routes

import (
	"fmt"
	"net/http"
	"text/template"
	"web/utils"
)

type Skill struct {
	Name string
}

type Profile struct {
	Name   string
	Age    int
	Skills []Skill
}

func Home(response http.ResponseWriter, request *http.Request) {
	// crear plantilla html
	template, errorTemplate := template.ParseFiles("templates/home.html", utils.PATH_TEMPLATE)
	if errorTemplate != nil {
		panic(errorTemplate)
	}

	token, name := utils.LoginReturn(request)
	fmt.Println(token, name)

	var skills = make([]Skill, 2)
	skills[0] = Skill{Name: "strong"}
	skills[1] = Skill{Name: "fast"}
	template.Execute(response, Profile{"Cesar", 3, skills})
}
