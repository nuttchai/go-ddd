package entity

import (
	"github.com/google/uuid"
	value_object "github.com/nuttchai/go-ddd/internal/domain/value-objects"
	validator "github.com/nuttchai/go-ddd/utils/validators"
)

type User struct {
	Id        string               `json:"id"`
	FirstName string               `json:"first_name"`
	LastName  string               `json:"last_name"`
	Email     string               `json:"email"`
	Address   value_object.Address `json:"address"`
}

func (u *User) IsUserValid() bool {
	isNameIncluded := u.FirstName != "" && u.LastName != ""
	isEmailValid := validator.IsEmailValid(u.Email)
	if !isNameIncluded || !isEmailValid {
		return false
	}
	return true
}

type UserProps struct {
	FirstName string               `json:"first_name"`
	LastName  string               `json:"last_name"`
	Email     string               `json:"email"`
	Address   value_object.Address `json:"address"`
}

func NewUser(props *UserProps, id ...string) *User {
	userId := uuid.New().String()
	if len(id) > 0 {
		userId = id[0]
	}

	return &User{
		Id:        userId,
		FirstName: props.FirstName,
		LastName:  props.LastName,
		Email:     props.Email,
		Address:   props.Address,
	}
}
