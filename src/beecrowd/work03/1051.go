package main

import "fmt"

func main() {

	var s, a, c float64

	fmt.Scan(&s)
	c = s

	if s > 2000 {
		c = s - 2000
		if c > 1000 {
			a += 1000 * 0.08
		} else {
			a += c * 0.08
		}
	}
	if s > 3000 {
		c = s - 3000
		if c > 1500 {
			a += 1500 * 0.18
		} else {
			a += c * 0.18
		}
	}
	if s > 4500 {
		c = s - 4500
		a += c * 0.28
	}

	if a > 0 {
		fmt.Printf("R$ %.2f\n", a)

	} else {
		fmt.Printf("Isento\n")

	}

}
