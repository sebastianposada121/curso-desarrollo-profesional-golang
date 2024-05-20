package handlers

import (
	"bufio"
	"crud-mysql/connect"
	"crud-mysql/models"
	"fmt"
	"log"
	"os"
	"strconv"
)

func GetClients() {
	sql := "select id, name, email, phone, date from clients order by id desc"
	data, err := connect.Db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer connect.CloseConection()
	clients := models.Clients{}
	for data.Next() {
		client := models.Client{}
		data.Scan(&client.Id, &client.Name, &client.Email, &client.Phone, &client.Date)
		clients = append(clients, client)
	}
	fmt.Println(clients)
}

func GetClientById(id int) {
	sql := "select id, name, email, phone, date from clients where id=?"
	data, err := connect.Db.Query(sql, id)
	if err != nil {
		panic(err)
	}
	defer connect.CloseConection()

	for data.Next() {
		var client models.Client
		err := data.Scan(&client.Id, &client.Name, &client.Email, &client.Phone, &client.Date)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(client)
	}
}

func CreateClient(client models.Client) {
	connect.Connect()
	sql := "insert into clients (name, email, phone) values(?,?,?)"

	fmt.Println(client)

	result, err := connect.Db.Exec(sql, client.Name, client.Email, client.Phone)
	if err != nil {
		panic(err)
	}

	fmt.Println("Se registro un cliente exitosamente", result)
}

func UpdateClient(client models.Client) {
	connect.Connect()
	sql := "update clients set name=?, email=?, phone=? where id=?"

	result, err := connect.Db.Exec(sql, client.Name, client.Email, client.Phone, client.Id)
	fmt.Println(client)
	if err != nil {
		panic(err)
	}

	fmt.Println("Se actualizo un cliente exitosamente", result)
}

func DeleteClient(id int) {
	connect.Connect()
	sql := "delete from clients where id=?"
	_, err := connect.Db.Exec(sql, id)

	if err != nil {
		panic(err)
	}

	fmt.Println("Se elimino el registro")
}

func Exec() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Select option:")
	fmt.Println("1 - List clients")
	fmt.Println("2 - List client detail")
	fmt.Println("3 - Create client")
	fmt.Println("4 - Update client")
	fmt.Println("5 - Delete client")

	if scanner.Scan() {
		switch scanner.Text() {
		case "1":
			GetClients()
		case "2":
			fmt.Println("Id")
			var id int
			if scanner.Scan() {
				id, _ = strconv.Atoi(scanner.Text())
			}
			GetClientById(id)
		case "3":
			var client models.Client
			fmt.Println("Name")
			if scanner.Scan() {
				client.Name = scanner.Text()
			}
			fmt.Println("Email")
			if scanner.Scan() {
				client.Email = scanner.Text()
			}
			fmt.Println("Phone")
			if scanner.Scan() {
				client.Phone = scanner.Text()
			}
			CreateClient(client)
		case "4":
			var client models.Client
			fmt.Println("Id")
			if scanner.Scan() {
				client.Id, _ = strconv.Atoi(scanner.Text())
			}
			fmt.Println("Name")
			if scanner.Scan() {
				client.Name = scanner.Text()
			}
			fmt.Println("Email")
			if scanner.Scan() {
				client.Email = scanner.Text()
			}
			fmt.Println("Phone")
			if scanner.Scan() {
				client.Phone = scanner.Text()
			}
			UpdateClient(client)
		case "5":
			fmt.Println("Id")
			var id int
			if scanner.Scan() {
				id, _ = strconv.Atoi(scanner.Text())
			}
			DeleteClient(id)
		}

	}
}
