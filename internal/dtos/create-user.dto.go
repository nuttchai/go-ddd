package dto

type CreateUserDTO struct {
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Email     string     `json:"email"`
	Address   AddressDTO `json:"address"`
}
