package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 3, 1, -1, 3}
	k := singleNumber(nums)
	fmt.Println(k)
}

func singleNumber(nums []int) int {

	result := 0

	for _, num := range nums {
		result ^= num
	}

	return result
}
