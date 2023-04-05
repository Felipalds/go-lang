package main

import "fmt"

func main() {
	var k int
	var t, m uint64

	fmt.Scan(&k)

	for i := 0; i < k; i++ {
		fmt.Scan(&t)
		if t == 64 {
			m = 1537228672809129
		} else {
			m = 1 << t
			m /= 12000
		}

		fmt.Printf("%d kg\n", m)

	}
}
