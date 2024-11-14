package main

import "fmt"

func main() {

	vet := []int{5, 3, 2, 4, 7, 1, 0, 6}
	n := len(vet)

	for i := 1; i < n; i++ {
		aux := vet[i]
		j := i - 1

		for j >= 0 && vet[j] > aux {
			vet[j+1] = vet[j]
			j = j - 1
		}

		vet[j+1] = aux
	}
	fmt.Println(vet)

}
