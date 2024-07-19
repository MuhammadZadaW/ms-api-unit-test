package utility

import (
	"regexp"
)

func CheckAlphaSpace(s string) bool {
	return regexp.MustCompile(`^[a-zA-Z \s]+$`).MatchString(s)
}