package main

import (
	"crud-mysql/connect"
	"crud-mysql/handlers"
)

func main() {
	connect.Connect()
	// handlers.GetClients()
	// handlers.GetClientById(1)

	// client := models.Client{
	// 	Id:    8,
	// 	Name:  "lulu",
	// 	Email: "gata@dada.co",
	// 	Phone: "1",
	// 	Date:  "2024/04/27",
	// }
	// handlers.CreateClient(client)

	// handlers.UpdateClient(client)
	// handlers.DeleteClient(client.Id)
	// handlers.GetClients()

	handlers.Exec()
}
