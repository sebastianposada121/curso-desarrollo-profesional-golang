package models

type GenericResponse struct {
	Data       any    `json:"data"`
	Successful bool   `json:"status"`
	Message    string `json:"message"`
}
