package main

import "fmt"

func main() {
	var personAmount, jump, k int
	fmt.Scan(&personAmount, &jump)
	var persons []int

	for k < personAmount {
		persons = append(persons, 1)
		k++
	}

	p := personAmount
	k = -1

	for p != 1 {
		k += jump
		fmt.Println(k)

		if k >= personAmount {
			k -= personAmount
		}

		for persons[k] == 0 {
			k++
		}

		persons[k] = 0
		p--
		fmt.Println(persons)

	}

	fmt.Println(persons)
}
