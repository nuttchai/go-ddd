package value_object

import (
	props "github.com/nuttchai/go-ddd/internal/domain/props"
)

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zip_code"`
}

func NewAddress(props *props.AddressProps) *Address {
	return &Address{
		Street:  props.Street,
		City:    props.City,
		State:   props.State,
		ZipCode: props.ZipCode,
	}
}
