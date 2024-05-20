package models

type Client struct {
	Id    int
	Name  string
	Email string
	Phone string
	Date  string
}

type Clients []Client

type FlashClient struct {
	status  string
	message string
	client  Client
}
