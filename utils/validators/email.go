package validator

import "regexp"

func IsEmailValid(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	emailPattern := regexp.MustCompile(emailRegex)
	return emailPattern.MatchString(email)
}
