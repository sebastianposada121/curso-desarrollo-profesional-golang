package routes

import (
	"fmt"
	"net/http"
	"text/template"
	"web/connect"
	"web/models"
	"web/utils"
	"web/validate"

	"golang.org/x/crypto/bcrypt"
)

func Login(response http.ResponseWriter, request *http.Request) {

	email := request.FormValue("email")
	password := request.FormValue("password")

	message := ""

	if !validate.Email(email) {
		message = message + "email required,"
	}

	if !validate.Password(password) {
		message = message + "password required,"
	}

	if message != "" {
		utils.CreateFlashMessage(response, request, message, "danger")
		http.Redirect(response, request, "/login", http.StatusSeeOther)
	}
	connect.Connect()

	var hashPassword string
	var name string

	if errL := connect.Db.QueryRow("select password, name from users where email=?", email).Scan(&hashPassword, &name); errL != nil {
		http.Redirect(response, request, "/login", http.StatusBadRequest)
	}

	errorPassword := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))

	// sebasTian432*
	if errorPassword != nil {
		fmt.Println("error---------------")
		http.Redirect(response, request, "/login", http.StatusSeeOther)
	} else {
		session, _ := utils.Store.Get(request, "session-name")
		session.Values["token"] = "dasda123das32ddasxads"
		session.Values["name"] = name

		errorSession := session.Save(request, response)
		if errorSession != nil {
			http.Redirect(response, request, "/login", http.StatusSeeOther)
		}
		http.Redirect(response, request, "/", http.StatusSeeOther)
	}

}

func Signup(response http.ResponseWriter, request *http.Request) {

	name := request.FormValue("name")
	email := request.FormValue("email")
	password := request.FormValue("password")
	phone := request.FormValue("phone")

	message := ""

	if len(name) == 0 {
		message = message + "name required,"
	}

	if !validate.Email(email) {
		message = message + "email required,"
	}

	if len(phone) == 0 {
		message = message + "phone required,"
	}

	if !validate.Password(password) {
		message = message + "password required,"
	}

	if message != "" {
		utils.CreateFlashMessage(response, request, message, "danger")
		http.Redirect(response, request, "/register-client", http.StatusSeeOther)
	}
	connect.Connect()
	sql := "insert into users (name, email, phone, password) values(?,?,?,?)"
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 8)

	user := models.User{
		Name:     name,
		Email:    email,
		Phone:    phone,
		Password: string(hash),
	}

	_, err := connect.Db.Exec(sql, user.Name, user.Email, user.Phone, user.Password)
	if err != nil {
		panic(err)
	}

	http.Redirect(response, request, "/login", http.StatusSeeOther)

}

func LoginPage(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("templates/login.html"))
	template.Execute(response, nil)
}

func SignupPage(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("templates/signup.html"))

	template.Execute(response, nil)
}

func Logout(response http.ResponseWriter, request *http.Request) {
	session, _ := utils.Store.Get(request, "session")
	session.Values["token"] = nil
	session.Values["name"] = nil

	if errorSession := session.Save(request, response); errorSession != nil {
		http.Error(response, errorSession.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(response, request, "/login", http.StatusSeeOther)
}
