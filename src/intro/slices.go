package main

import "fmt"

func read_array(ar [4]int) {
	fmt.Println("Received ", ar)
}

func read_pointer(ar *[4]int) {
	fmt.Println("Received ", ar)
}

func main() {
	fmt.Println("Hello!")

	// arrays
	var array1 [4]int
	array1[2] = 54

	fmt.Println(array1)
	read_array(array1)
	read_pointer(&array1)

	var array2 [5]int
	array2[3] = 32
	// read_array(array2) -> error
	// read_pointer(&array2) -> error

	array3 := [...]string{"Name 1", "Name 2"}
	// initializing an array and makig the compiler count for you
	fmt.Println(array3)

	// slices
	slice1 := []int{1, 2, 3}
	fmt.Println(slice1)

	var slice2 []int
	slice2 = append(slice2, 3)
	fmt.Println(slice2)

	slice3 := make([]int, 3, 3)
	fmt.Println(slice3)

	b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	b2 := b[1:4]
	fmt.Println(b)
	b2[2] = '3'
	fmt.Println(b)

}
