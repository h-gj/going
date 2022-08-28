package puzzles

import "strings"

func StringOrderedEqual(s1, s2 string) bool {
	strLen1 := len([]rune(s1))
	strLen2 := len([]rune(s2))
	if strLen1 > 5000 || strLen2 > 5000 || strLen1 != strLen2 {
		return false
	}

	for _, v := range s1 {
		if strings.Count(s1, string(v)) != strings.Count(s2, string(v)) {
			return false
		}
	}
	return true

}
