package main

import "fmt"

func main() {

	nums := []int{1, 3, 4}

	fmt.Println(resultsArray(nums, 2))

}

func resultsArray(nums []int, k int) []int {

	var result []int

	for i, e := range nums {

		control := e

		if len(nums)-i < k {
			break
		}

		for j := 1; j < k; j++ {
			if nums[i+j] <= control || nums[i+j] != control+1 {
				control = -1
				break
			}
			control = nums[i+j]
		}
		result = append(result, control)
	}

	return result

}
