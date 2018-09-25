package kata

import (
	"strings"
)

func bandNameGenerator(word string) string {
	if word == "" {
		return ""
	} else if word[:1] == word[len(word)-1:] {
		return strings.Title(word + word[1:])
	}
	return strings.Title("the " + word)
}

func BestbandNameGenerator(word string) string {
	first := word[:1]
	last := word[len(word)-1:]

	if first != last {
		return "The " + strings.Title(word)
	}

	return strings.Title(word) + word[1:]

}
