package entity

import (
	value_object "github.com/nuttchai/go-ddd/user/domain/value-objects"
)

type User struct {
	Id        int                  `json:"id"`
	FirstName string               `json:"first_name"`
	LastName  string               `json:"last_name"`
	Email     string               `json:"email"`
	Address   value_object.Address `json:"address"`
}
