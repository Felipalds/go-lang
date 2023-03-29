package main

import "fmt"

func main() {
	var x, base, ca, before int
	ca++
	for {
		fmt.Scan(&x, &base)
		if x == 0 && base == 0 {
			break
		}
		before = x

		a := 0
		b := 1
		c := 1

		for x > 0 {
			a = 1 + b + c
			if a > base {
				a -= base
			}

			c = b
			b = a
			x--
		}

		fmt.Printf("Case %d: %d %d %d\n", ca, before, base, c)
		ca++
	}

}
