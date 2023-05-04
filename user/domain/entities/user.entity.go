package entity

import (
	"regexp"

	value_object "github.com/nuttchai/go-ddd/user/domain/value-objects"
)

type User struct {
	Id        string               `json:"id"`
	FirstName string               `json:"first_name"`
	LastName  string               `json:"last_name"`
	Email     string               `json:"email"`
	Address   value_object.Address `json:"address"`
}

func (u *User) isEmailValid() bool {
	emailRegex := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	emailPattern := regexp.MustCompile(emailRegex)
	return emailPattern.MatchString(u.Email)
}

func (u *User) IsUserValid() bool {
	if u.FirstName == "" || u.LastName == "" || u.isEmailValid() {
		return false
	}
	return true
}
