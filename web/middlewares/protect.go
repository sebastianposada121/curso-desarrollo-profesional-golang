package middlewares

import (
	"net/http"
	"web/utils"
)

func ProtectSession(next http.HandlerFunc) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		session, _ := utils.Store.Get(request, "session-name")

		if session.Values["token"] != nil {
			next.ServeHTTP(response, request)
		} else {
			http.Redirect(response, request, "/login", http.StatusSeeOther)
		}

	}
}
