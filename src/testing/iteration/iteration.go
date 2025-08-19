package iteration

import "strings"

func RepeatRefactored(char string) string {
	var res strings.Builder

	for range 1000 {
		res.WriteString(char)
	}
	return res.String()
}

func RepeatDumb(char string) string {
	var res string

	for range 1000 {
		res += char
	}
	return res
}
