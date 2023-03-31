package main

import "fmt"

func main() {

	var s, a, p float64

	fmt.Scan(&s)

	if s >= 0 && s <= 400 {
		p = 0.15
	} else if s > 400 && s <= 800 {
		p = 0.12
	} else if s > 800 && s <= 1200 {
		p = 0.10
	} else if s > 1200 && s <= 2000 {
		p = 0.07
	} else {
		p = 0.04
	}

	a = p * s
	s += a

	fmt.Printf("Novo salario: %.2f\nReajuste ganho: %.2f\nEm percentual: %.0f %%\n", s, a, p*100)
}
