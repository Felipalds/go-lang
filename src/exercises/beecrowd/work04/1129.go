package main

import "fmt"

func main() {
	var test, a, b, c, d, e, more int
	var letter string
	test = 1
	for {

		fmt.Scan(&test)
		if test == 0 {
			return
		}

		for k := 0; k < test; k++ {
			fmt.Scan(&a, &b, &c, &d, &e)
			more = 0
			if a <= 127 {
				more++
				letter = "A"
			}
			if b <= 127 {
				more++
				letter = "B"
			}
			if c <= 127 {
				more++
				letter = "C"

			}
			if d <= 127 {
				more++
				letter = "D"
			}
			if e <= 127 {
				more++
				letter = "E"
			}

			if more != 1 {
				fmt.Println("*")
			} else {
				fmt.Println(letter)
			}

		}

	}
}
