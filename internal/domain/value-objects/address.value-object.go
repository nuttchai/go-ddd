package value_object

import (
	vprops "github.com/nuttchai/go-ddd/internal/domain/value-objects/props"
)

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zip_code"`
}

func NewAddress(props *vprops.AddressProps) *Address {
	return &Address{
		Street:  props.Street,
		City:    props.City,
		State:   props.State,
		ZipCode: props.ZipCode,
	}
}
