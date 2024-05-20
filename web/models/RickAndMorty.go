package models

type Response struct {
	Info    Info
	Results Results
}

type Info struct {
	Count int
	Pages int
	Next  string
	Prev  int
}

type Character struct {
	Id      int
	Name    string
	status  string
	Image   string
	Gender  string
	Species string
}

type Results []Character
