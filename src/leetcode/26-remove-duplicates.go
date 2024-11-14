package main

import "fmt"

func removeElement(vet []int, index int) []int {
	return append(vet[:index], vet[index+1:]...)
}

func main() {

	numbers := []int{1, 1, 2, 2}
	i := removeDuplicates(numbers)
	fmt.Println(i)
	fmt.Println(numbers)

}

func removeDuplicates(nums []int) int {
	// 0 1 2 2
	k := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}

vetor[0, 1, 2, 3]

void memama(&vet) {
	vet[0] = 2
}

void mama(vet) {
	vet_novo = vet
}
