package main

import "fmt"

func main() {
	var x int = 5
	var p *int
	p = &x
	fmt.Println(p)
}
