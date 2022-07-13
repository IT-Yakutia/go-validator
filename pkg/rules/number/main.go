package rule_number

import (
	"regexp"
	"strconv"

	rule_string "github.com/IT-Yakutia/go-validator/pkg/rules/string"
)

func IsNumberWithLenght(n string, length int) bool {
	if n == "" {
		return false
	}
	if rule_string.GetLength(n) != length {
		return false
	}
	str := strconv.Itoa(length)
	matched, _ := regexp.MatchString(`\d{`+str+`}`, n)
	return matched
}
