package validate_phone

import (
	"fmt"

	rule_number "github.com/IT-Yakutia/go-validator/pkg/rules/number"
)

type Phone struct {
	code   string
	number string
}
type PhoneValidator interface {
	ValidateInternational() bool
	ValidateRussian() bool
	GetNumber() string
}

func NewPhone(code string, number string) PhoneValidator {
	return &Phone{code: code, number: number}
}

func (p Phone) ValidateInternational() bool {
	return false
}

func (p Phone) ValidateRussian() bool {
	if p.code == "+7" || p.code == "8" || p.code == "7" {
		if rule_number.IsNumberWithLenght(p.number, 10) {
			fmt.Printf("Валидация телефона %s прошла успешно!\n", p.GetNumber())
			return true
		}
	}
	fmt.Println("Телефон не правильный!")
	return false
}

func (p Phone) GetNumber() string {
	return p.code + p.number
}
