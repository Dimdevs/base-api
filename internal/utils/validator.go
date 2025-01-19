package utils

import (
	"regexp"
)

func IsValidEmail(email string) bool {
	regex := `^[a-z0-9]+@[a-z0-9]+\.[a-z]{2,}$`
	re := regexp.MustCompile(regex)
	return re.MatchString(email)
}
