package main

import "fmt"

func createSlice(n int) []int {
	s := []int{}

	for i := 0; i < n; i++ {
		s = append(s, i+1)
	}

	return s
}

func pop(persons []int, k int) []int {
	copy(persons[k:], persons[k+1:])
	persons = persons[:len(persons)-1]
	return persons
}

func main() {
	var (
		nc, n, k int
	)

	fmt.Scan(&nc)
	for i := 0; i < nc; i++ {
		fmt.Scan(&n, &k)
		s := []int{}
		s = createSlice(n)
		j := 0

		for len(s) > 1 {
			j = (j + k - 1) % len(s)
			s = pop(s, j)
		}

		fmt.Printf("Case %d: %d\n", i+1, s[0])
	}
}
