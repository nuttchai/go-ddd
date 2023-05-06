package validator

import "github.com/go-playground/validator"

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func IsValidStruct(s interface{}) (bool, error) {
	if err := validate.Struct(s); err != nil {
		return false, err
	}
	return true, nil
}
