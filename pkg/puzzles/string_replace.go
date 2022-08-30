package puzzles

import (
	"strings"
	"unicode"
)

func StringReplace(s string) (string, bool) {
	str := []rune(s)
	if len(str) > 1000 {
		return s, false
	}

	for _, v := range str {
		if string(v) != " " && unicode.IsLetter(v) == false {
			return s, false
		}
	}

	return strings.Replace(s, " ", "%20", -1), true
}
