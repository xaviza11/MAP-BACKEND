package validators

import (
	"errors"
	"regexp"
)

func ValidateName(name string) error {
	if len(name) < 2 || len(name) > 22 {
		return errors.New("name must be between 2 and 22 characters")
	}
	return nil
}

func ValidateEmail(email string) error {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(emailRegex, email)
	if !match {
		return errors.New("invalid email format")
	}
	return nil
}

func ValidatePassword(password string) error {
	passwordRegex := `^[a-zA-Z0-9]{9,}$`
	match, _ := regexp.MatchString(passwordRegex, password)
	if !match {
		return errors.New("password must be at least 9 characters long and contain only letters and numbers")
	}
	return nil
}
