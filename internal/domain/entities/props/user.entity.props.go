package entity_props

import (
	props "github.com/nuttchai/go-ddd/internal/domain/value-objects/props"
)

type UserProps struct {
	FirstName string             `json:"first_name"`
	LastName  string             `json:"last_name"`
	Email     string             `json:"email"`
	Address   props.AddressProps `json:"address"`
}
