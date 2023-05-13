package entity

import (
	"regexp"

	"github.com/google/uuid"
	value_object "github.com/nuttchai/go-ddd/internal/domain/value-objects"
)

type UserProps struct {
	FirstName string               `json:"first_name"`
	LastName  string               `json:"last_name"`
	Email     string               `json:"email"`
	Address   value_object.Address `json:"address"`
}

type User struct {
	Id        string               `json:"id"`
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

func (u *User) IsEmailValid() bool {
	emailRegex := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	emailPattern := regexp.MustCompile(emailRegex)
	return emailPattern.MatchString(u.Email)
}

func (u *User) IsUserValid() bool {
	if u.FirstName == "" || u.LastName == "" || u.IsEmailValid() {
		return false
	}
	return true
}
