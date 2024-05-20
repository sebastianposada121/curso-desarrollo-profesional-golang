package main

import (
	"log"
	"net/http"
	"os"
	"web/connect"
	"web/middlewares"
	"web/routes"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

/*
	mux: libreria de enrutamiento
	-Enrutamiento flexible y potente
	-Manejo de parámetros de URL
	-Compatibilidad con middleware
	-Personalización de respuestas HTTP
	-Documentación y comunidad
*/

func main() {
	connect.Connect()
	mux := mux.NewRouter()
	// archivos estaticos
	assets := http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/")))
	mux.PathPrefix("/assets/").Handler(assets)
	// rutas
	mux.HandleFunc("/", middlewares.ProtectSession(routes.Home))
	mux.HandleFunc("/clients", routes.GetClients)
	// clients
	// mux.HandleFunc("/client/{id:.*}/{name:.*}", routes.Client)
	mux.HandleFunc("/delete-client/{id:.*}", routes.DeleteClient)
	mux.HandleFunc("/register-client", routes.RegisterClient)
	mux.HandleFunc("/register-client-post", routes.RegisterClientPost).Methods("POST")
	mux.HandleFunc("/edit-client/{id:.*}", routes.EditClient)
	mux.HandleFunc("/edit-client-put/{id:.*}", routes.EditClientPut)
	// users
	mux.HandleFunc("/login", routes.LoginPage)
	mux.HandleFunc("/signup", routes.SignupPage)
	mux.HandleFunc("/login-post", routes.Login).Methods("POST")
	mux.HandleFunc("/signup-post", routes.Signup).Methods("POST")
	//
	mux.HandleFunc("/resources", routes.Resources)
	mux.HandleFunc("/generate-pdf", routes.GeneratePdf)
	mux.HandleFunc("/generate-excel", routes.GenerateExcel)
	mux.HandleFunc("/generate-qr", routes.GenerateQr)
	mux.HandleFunc("/generate-email", routes.GenerateEmail)
	mux.HandleFunc("/rick-and-morty", routes.RickAndMorty)
	mux.HandleFunc("/webpay", routes.Webpay)
	mux.HandleFunc("/payment-gateway", routes.PaymentGateway)
	mux.HandleFunc("/payment-gateway-webpay-info", routes.WebpayInfo)
	mux.HandleFunc("/paypal", routes.Paypal)
	mux.HandleFunc("/payment-gateway-paypal-info", routes.PaypalInfo)
	mux.HandleFunc("/payment-gateway-paypal-cancel", routes.PaypalCancel)
	mux.HandleFunc("/logout", middlewares.ProtectSession(routes.Logout))
	mux.NotFoundHandler = mux.NewRoute().HandlerFunc(routes.NotFound).GetHandler()

	// cargar constantes
	errorGodotenv := godotenv.Load()

	if errorGodotenv != nil {
		panic(errorGodotenv)
	}
	server := &http.Server{
		Addr:    "localhost:" + os.Getenv("PORT"),
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}

/*
	TODO usar fresh para actualizar al guardarr

	go get  github.com/pilu/fresh
	go run github.com/pilu/fresh
*/

/*
Montar servidor http
*/

// func main() {
// 	PORT := "localhost:8081"
// 	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
// 		fmt.Fprintln(response, "holiii")
// 	})

// 	fmt.Println("start in ", PORT)
// 	log.Fatal(http.ListenAndServe(PORT, nil))
// }
