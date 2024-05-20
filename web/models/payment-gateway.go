package models

type WebpayOrder struct {
	Buy_order  string
	Session_id string
	Amount     int
	Return_url string
}

type WebpayModel struct {
	Url   string
	Token string
}
