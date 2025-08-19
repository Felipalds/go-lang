package main

import "fmt"

func isPerfect(n int) bool {
	div := n / 2
	num := 0
	for i := div; i > 0; i-- {
		if n%i == 0 {
			num += i
		}
	}

	if num == n {
		return true
	} else {
		return false
	}
}

func main() {
	var k, n int
	fmt.Scan(&k)
	for i := 0; i < k; i++ {
		fmt.Scan(&n)

		isP := isPerfect(n)
		if isP {
			fmt.Printf("%d eh perfeito\n", n)
		} else {
			fmt.Printf("%d nao eh perfeito\n", n)
		}
	}
}
