package rule_string

import (
	"fmt"
)

func Equals(s string, length int) bool {
	fmt.Printf("строка: %s должен иметь длину %v символов.\n", s, length)
	asRunes := []rune(s)
	return length == len(asRunes)
}

func GetSymbolPosition(s string, symbol rune) int {
	asRunes := []rune(s)
	index := -1
	for i := 0; i < len(asRunes); i++ {
		if asRunes[i] == symbol {
			index = i
			break
		}
	}
	return index
}

func GetLength(s string) int {
	asRunes := []rune(s)
	return len(asRunes)
}

func Substr(input string, start int) string {
	asRunes := []rune(input)

	if start >= len(asRunes) {
		return ""
	}

	return string(asRunes[start:])
}
