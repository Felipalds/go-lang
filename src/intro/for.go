package main

import "fmt"

func main() {
	var k int = 0
	for {
		k++
		fmt.Println(k)
		if k == 5 {
			break
		}
	}
}
