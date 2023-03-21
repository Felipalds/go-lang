package main

import "fmt"

func main() {
	var a, b int16
	fmt.Scan(&a, &b)

	if a >= b {
		fmt.Println("A é maior ou igual a B")
	} else {
		fmt.Println("B é maior que A")
	}

	switch a {
	case 1:
		fmt.Print(1)
		fmt.Print(2)
	case 2:
		fmt.Print(3)
		fmt.Print(4)
		fmt.Print(1)
		fmt.Print(2)
	}
}
