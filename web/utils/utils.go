package utils

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var PATH_TEMPLATE string = "templates/index.html"

var Store = sessions.NewCookieStore([]byte("session-name"))

/*
obtener mensaje flash por cookies
*/
func GetFlashMessage(response http.ResponseWriter, request *http.Request) (message string, status string) {
	session, _ := Store.Get(request, "flash-session")
	_message := session.Flashes("message")
	_status := session.Flashes("status")

	if _message != nil {
		message = _message[0].(string)
	}
	if _status != nil {
		status = _status[0].(string)
	}
	session.Save(request, response)
	return message, status
}

/*
crear mensaje flash por cookies
*/
func CreateFlashMessage(response http.ResponseWriter, request *http.Request, message string, status string) {
	session, error := Store.Get(request, "flash-session")

	if error != nil {
		http.Error(response, error.Error(), http.StatusInternalServerError)
		return
	}
	session.AddFlash(message, "message")
	session.AddFlash(status, "status")
	session.Save(request, response)
}

func LoginReturn(request *http.Request) (token string, name string) {
	session, _ := Store.Get(request, "session-name")

	if session.Values["token"] != nil {
		token = session.Values["token"].(string)
	}
	if session.Values["name"] != nil {
		name = session.Values["name"].(string)
	}
	return token, name
}
