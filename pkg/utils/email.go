package utils

import "net/mail"

func ValidateEmail(email string) (isValid bool) {
	_, err := mail.ParseAddress(email)
	return err == nil
}
