package kata

import (
	"strings"
)

func ToCamelCase(s string) string {
	if s == "" {
		return ""
	}
	result := strings.Title(strings.Replace(strings.Replace(s, "-", " ", -1), "_", " ", -1))
	result = s[:1] + result[1:]
	result = strings.Replace(result, " ", "", -1)
	return result
}
