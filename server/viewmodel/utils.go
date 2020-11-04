package viewmodel

import "strings"

func emptyString(str string) bool {
	return strings.TrimSpace(str) == ""
}
