package main

import (
	"fmt"
)

func main() {
	var x1, x2 int64
	for {
		n, err := fmt.Scan(&x1, &x2)
		if n != 2 && err != nil {
			return
		}

		s := x1 ^ x2
		fmt.Println(s)
	}
}
