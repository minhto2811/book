package utils

import (
	"regexp"
	"strings"
)

func ValidateEmail(email string) bool {
	regex := `^[a-zA-Z0-9._%+-]+@gmail\.com$`
	re := regexp.MustCompile(regex)
	return re.MatchString(email)
}

func IsEmpty(text string) bool {
	return strings.TrimSpace(text) == ""
}
