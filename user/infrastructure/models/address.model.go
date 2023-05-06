package model

type Address struct {
	Id      string `json:"id"`
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zip_code"`
	UserId  string `json:"user_id"`
}
