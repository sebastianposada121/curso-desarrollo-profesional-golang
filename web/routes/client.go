package routes

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"text/template"
	"time"
	"web/connect"
	"web/models"
	"web/utils"
	"web/validate"

	"github.com/gorilla/mux"
)

/*
Obtener clientes
*/
func GetClients(response http.ResponseWriter, request *http.Request) {

	sql := "select id, name, email, phone, date from clients order by id desc"
	data, err := connect.Db.Query(sql)
	if err != nil {
		panic(err)
	}

	clients := models.Clients{}
	for data.Next() {
		client := models.Client{}
		data.Scan(&client.Id, &client.Name, &client.Email, &client.Phone, &client.Date)
		clients = append(clients, client)
	}
	fmt.Println(clients)

	template, errorTemplate := template.ParseFiles("templates/clients.html", utils.PATH_TEMPLATE)
	if errorTemplate != nil {
		panic(errorTemplate)
	}

	template.Execute(response, &clients)
	// defer connect.CloseConection()
}

/*
editar cliente
*/
func EditClientPut(response http.ResponseWriter, request *http.Request) {

	params := mux.Vars(request)
	paramsId := params["id"]
	id, errParseId := strconv.Atoi(paramsId)
	if errParseId != nil {
		// Manejar el error si la conversión falla
		fmt.Fprintf(response, "ID inválido: %v", errParseId)
		return
	}

	name := request.FormValue("name")
	email := request.FormValue("email")
	phone := request.FormValue("phone")
	message := ""

	if len(name) == 0 {
		message = message + "name required,"
	}
	if !validate.Email(email) {
		message = message + "email required,"
	}

	if message != "" {
		utils.CreateFlashMessage(response, request, message, "danger")
		http.Redirect(response, request, "/edit-client/"+strconv.Itoa(id), http.StatusSeeOther)
	}

	client := models.Client{
		Name:  name,
		Email: email,
		Phone: phone,
	}

	connect.Connect()
	sql := "update clients set name=?, email=?, phone=? where id=?"

	_, err := connect.Db.Exec(sql, client.Name, client.Email, client.Phone, id)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	http.Redirect(response, request, "/clients", http.StatusSeeOther)
}

/*
eliminar cliente
*/
func DeleteClient(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	paramsId := params["id"]
	id, errParseId := strconv.Atoi(paramsId)
	if errParseId != nil {
		// Manejar el error si la conversión falla
		fmt.Fprintf(response, "ID inválido: %v", errParseId)
		return
	}

	connect.Connect()
	sql := "delete from clients where id=?"
	_, err := connect.Db.Exec(sql, id)

	if err != nil {
		panic(err)
	}

	http.Redirect(response, request, "/clients", http.StatusSeeOther)
}

/*
mostrar vista para actualizar clientes
*/
func EditClient(response http.ResponseWriter, request *http.Request) {
	// obtener params
	// ----path
	// params["id"] params["name"]
	// ----query params
	//  request.URL.Query().Get("status"); request.URL.Query().Get("address")
	// status := request.URL.Query().Get("id")
	// address := request.URL.Query().Get("address")

	// data := map[string]string{
	// 	"id":      params["id"],
	// 	"name":    params["name"],
	// 	"s":       "hellow",
	// 	"status":  status,
	// 	"address": address,
	// }

	params := mux.Vars(request)
	paramsId := params["id"]
	if params["id"] == "" {
		http.Redirect(response, request, "/clients", http.StatusSeeOther)
		return
	}

	client := models.Client{}
	query := "select id, name, email, phone, date FROM clients where id=?"
	// Convierte el id de string a int
	id, errParseId := strconv.Atoi(paramsId)
	if errParseId != nil {
		// Manejar el error si la conversión falla
		fmt.Fprintf(response, "ID inválido: %v", errParseId)
		return
	}

	errorQuery := connect.Db.QueryRow(query, id).Scan(&client.Id, &client.Name, &client.Email, &client.Phone, &client.Date)

	if errorQuery != nil {
		if errorQuery == sql.ErrNoRows {
			// No se encontró el cliente
			http.Redirect(response, request, "/clients", http.StatusSeeOther)
			return
		}
		panic(errorQuery) // O maneja el error de manera más adecuada
	}

	template, errorTemplate := template.ParseFiles("templates/edit-client.html", utils.PATH_TEMPLATE)
	if errorTemplate != nil {
		panic(errorTemplate)
	}

	fmt.Println(client)
	template.Execute(response, &client)
}

/*
mostar vista para registrar clientes
*/
func RegisterClient(response http.ResponseWriter, request *http.Request) {
	// al usar must no necesitamos validar si la plantilla existe
	template := template.Must(template.ParseFiles("templates/register-client.html", utils.PATH_TEMPLATE))
	message, status := utils.GetFlashMessage(response, request)
	flash := map[string]string{
		"message": message,
		"status":  status,
	}
	template.Execute(response, flash)
}

func RegisterClientPost(response http.ResponseWriter, request *http.Request) {

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

	uploadPhoto := Upload(response, request, "photo")

	if uploadPhoto != "" {
		message = message + uploadPhoto
	}

	if message != "" {
		utils.CreateFlashMessage(response, request, message, "danger")
		http.Redirect(response, request, "/register-client", http.StatusSeeOther)
	}

	connect.Connect()
	sql := "insert into clients (name, email, phone) values(?,?,?)"
	client := models.Client{
		Name:  name,
		Email: email,
		Phone: phone,
	}

	_, err := connect.Db.Exec(sql, client.Name, client.Email, client.Phone)
	if err != nil {
		panic(err)
	}

	utils.CreateFlashMessage(response, request, "the client was successfully registered", "success")
	http.Redirect(response, request, "/clients", http.StatusSeeOther)
}

/*
Subir archivos a la aplicacion
*/
func Upload(response http.ResponseWriter, request *http.Request, name string) string {
	// archivo
	message := ""
	file, handler, errorFile := request.FormFile(name)

	if errorFile != nil {
		message = name + " required"
	}

	var ext = strings.Split(handler.Filename, ".")[1]
	encryptedName := strings.Split(time.Now().String(), " ")
	fmt.Println(encryptedName)
	nameFile := string(encryptedName[4][6:14]) + "." + ext
	var path_photo string = "assets/photos/" + nameFile
	f, errorOpenFile := os.OpenFile(path_photo, os.O_WRONLY|os.O_CREATE, 0777)
	_, errCopy := io.Copy(f, file)

	if errCopy != nil || errorOpenFile != nil {
		message = "We couldn't load the file"
	}

	return message
}
