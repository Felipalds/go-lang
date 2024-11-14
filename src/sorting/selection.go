package main

import (
	"fmt"
)

func main() {

	vet := []int{3, 5, 1, 4, 2, 9, 5, 8}
	n := len(vet)

	for i := 0; i < n; i++ {
		minn := i

		for j := i + 1; j < n; j++ {
			if vet[j] < vet[minn] {
				minn = j
			}
		}
		if vet[i] != vet[minn] {
			aux := vet[i]
			vet[i] = vet[minn]
			vet[minn] = aux
		}
	}
	fmt.Print(vet)
}
