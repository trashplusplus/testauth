package validator

import (
	"net/mail"
)

func ValidateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func ValidatePassword(password string) bool {
	return len(password) >= 8
}

func ValidateUsername(username string) bool {
	return len(username) >= 3
}
