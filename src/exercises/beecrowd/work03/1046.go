package main

import "fmt"

func main() {
	var initial_hour, end_hour, duration int

	fmt.Scan(&initial_hour, &end_hour)

	if initial_hour > end_hour {
		end_hour += 24
	}
	if initial_hour == end_hour {
		duration = 24
	} else {
		duration = end_hour - initial_hour
	}

	fmt.Printf("O JOGO DUROU %d HORA(S)\n", duration)
}
