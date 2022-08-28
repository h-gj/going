package puzzles

import "fmt"

// 1 2 3 4
// 0 1 2 3
//len - i - 1

// 1 2 3 4 5
// 0 1 2 3 4
// len - i -1

func ReverseString(s string) (string, bool) {
	str := []rune(s)
	if len(str) >= 5000 {
		return s, false
	}
	fmt.Println(str)

	for i := 0; i < len(str)/2; i++ {
		str[i], str[len(str)-i-1] = str[len(str)-i-1], str[i]
	}

	return string(str), true
}
