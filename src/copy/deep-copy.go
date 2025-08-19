package main

import "fmt"

func main() {

	// slices are value types
	// they are shallow copied
	slice1 := []int{1, 2, 3}
	fmt.Println(slice1)

	slice2 := slice1
	slice2[0] = 10
	fmt.Println(slice2)
	fmt.Println(slice1)

	// this is how we deep copy slices in go! :)
	slice3 := make([]int, len(slice1))
	copy(slice3, slice1)
	slice3[0] = 20
	fmt.Println(slice3)
	fmt.Println(slice1)

	// this is a pointer
	var x int = 5
	var p *int = &x
	fmt.Println(p) // will print the location in memory!
}
