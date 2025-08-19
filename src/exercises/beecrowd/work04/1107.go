package main

import "fmt"

func main() {
	var x, y, h, act int

	for {
		act = 0
		fmt.Scan(&x, &y)
		if x == 0 && y == 0 {
			return
		}

		for i := 0; i < y; i++ {
			fmt.Scan(&h)
			if h < x {
				act += x - h
			}
			x = h
		}

		fmt.Println(act)

	}
}
