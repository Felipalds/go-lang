package main

import "fmt"

func main() {
	var tests, x0, y0, x, y int
	tests = 1
	for {
		fmt.Scan(&tests)

		if tests == 0 {
			return
		} else {

			fmt.Scan(&x0, &y0)
			for i := 0; i < tests; i++ {
				fmt.Scan(&x, &y)
				if x == x0 || y == y0 {
					fmt.Println("divisa")
				} else if x > x0 && y > y0 {
					fmt.Println("NE")
				} else if x > x0 && y < y0 {
					fmt.Println("SE")
				} else if x < x0 && y > y0 {
					fmt.Println("NO")
				} else if x < x0 && y < y0 {
					fmt.Println("SO")
				}
			}
		}
	}
}
