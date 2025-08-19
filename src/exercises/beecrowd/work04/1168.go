package main

import "fmt"

func main() {

	var k, s int
	var n string
	fmt.Scan(&k)
	for i := 0; i < k; i++ {
		s = 0
		fmt.Scan(&n)
		for k := range n {
			l := string(n[k])
			if l == "1" {
				s += 2
			}
			if l == "2" || l == "3" || l == "5" {
				s += 5
			}
			if l == "4" {
				s += 4
			}
			if l == "6" || l == "9" || l == "0" {
				s += 6
			}
			if l == "7" {
				s += 3
			}
			if l == "8" {
				s += 7
			}
		}
		fmt.Printf("%d leds\n", s)
	}
}
