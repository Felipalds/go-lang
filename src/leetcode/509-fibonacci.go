package main

import "fmt"

func main() {

	fmt.Println(fib(4))

}

func fib(n int) int {
	// 5
	// 0 1 1 2 3

	a := 0
	b := 1
	c := 0

	if n == 0 {
		return 0
	}

	if n == 1 || n == 2 {
		return 1
	}

	for i := 1; i < n; i++ {
		c = a + b
		a = b
		b = c
	}

	return c
}
