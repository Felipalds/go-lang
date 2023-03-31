package main

import "fmt"

func isTriangle(a float64, b float64, c float64) bool {
	if a < b+c && b < a+c && c < a+b {
		return true
	} else {
		return false
	}
}

func main() {
	var a, b, c float64

	fmt.Scan(&a, &b, &c)

	if isTriangle(a, b, c) {
		fmt.Printf("Perimetro = %.1f\n", a+b+c)
	} else {
		fmt.Printf("Area = %.1f\n", (a+b)*c/2)
	}
}
