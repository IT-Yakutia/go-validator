package validate_email

import (
	"fmt"
	"net/mail"

	rule_string "github.com/IT-Yakutia/go-validator/pkg/rules/string"
)

type Email struct {
	email string
}

type EmailValidator interface {
	Validate() bool
}

func NewEmail(email string) EmailValidator {
	return &Email{email: email}
}

func (e *Email) Validate() bool {
	_, error := mail.ParseAddress(e.email)

	var atPosition int = rule_string.GetSymbolPosition(e.email, '@')
	if error != nil {
		fmt.Printf("Email не правильный %s: %v\n", e.email, error == nil)
		return false
	}

	if atPosition <= 0 {
		fmt.Printf("Нет символа @, или @ стоит первым. В %s: %v\n", e.email, error == nil)
		return false
	}

	var atStartString string = rule_string.Substr(e.email, atPosition)
	var dotPosition int = rule_string.GetSymbolPosition(atStartString, '.')
	if atPosition <= 0 || dotPosition < 1 {
		fmt.Printf("Нет символа . после симовла @. В %s: %v\n", e.email, error == nil)
		return false
	}

	return true
}
