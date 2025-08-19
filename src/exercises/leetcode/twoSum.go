package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var m map[int]int
	result := make([]int, 2)
	m = make(map[int]int)

	for i, num := range nums {
		inverse := target - num
		if j, ok := m[num]; ok {
			result[0] = i
			result[1] = j
			break
		}
		m[inverse] = i
	}
	return result
}

func main() {
	a := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(a)
}
