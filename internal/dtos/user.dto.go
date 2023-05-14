package dto

type UserDTO struct {
	Id        string     `json:"id"`
	FirstName string     `json:"firstName"`
	LastName  string     `json:"lastName"`
	Email     string     `json:"email"`
	Address   AddressDTO `json:"address"`
}
