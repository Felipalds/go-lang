package reverseinteger

import (
	"math"
	"strconv"
)

func Reverse(s int) int {
	return reverse(s)
}

func reverse(i int) int {

	s := strconv.Itoa(i)
	multiplier := 1
	if s[0] == '-' {
		s = s[1:]
		multiplier = -1
	}

	string_size := len(s)
	buff := make([]byte, string_size)

	for i := string_size - 1; i >= 0; i-- {
		buff[string_size-i-1] = s[i]
	}

	string_value := string(buff)

	int_value, _ := strconv.Atoi(string_value)

	if int_value >= int(math.Pow(2, 31)-1) || int_value <= -int(math.Pow(2, 31)) {
		return 0
	}

	return int_value * multiplier
}
