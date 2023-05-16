package entity

import (
	"github.com/google/uuid"
	props "github.com/nuttchai/go-ddd/internal/domain/entities/props"
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

func NewUser(props *props.UserProps, id ...string) *User {
	userId := uuid.New().String()
	if len(id) > 0 && id[0] != "" {
		userId = id[0]
	}

	address := value_object.NewAddress(&props.Address)
	return &User{
		Id:        userId,
		FirstName: props.FirstName,
		LastName:  props.LastName,
		Email:     props.Email,
		Address:   *address,
	}
}
