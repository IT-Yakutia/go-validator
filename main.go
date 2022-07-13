package main

import (
	validate_email "github.com/IT-Yakutia/go-validator/pkg/validators/email"
	validate_phone "github.com/IT-Yakutia/go-validator/pkg/validators/phone"
)

func main() {
	var email_validator validate_email.EmailValidator = validate_email.NewEmail("sadasdasd@asdad.asda")
	email_validator.Validate()

	var validate_phone validate_phone.PhoneValidator = validate_phone.NewPhone("+7", "9142736836")
	validate_phone.ValidateRussian()
}
