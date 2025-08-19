package main

import "fmt"

func main() {
	var initial_hour, initial_minute, final_hour, final_minute int
	fmt.Scan(&initial_hour, &initial_minute, &final_hour, &final_minute)

	minutesStart := initial_hour*60 + initial_minute
	minutesEnd := final_hour*60 + final_minute

	if minutesStart > minutesEnd {
		minutesEnd += 60 * 24
	}
	minutes := minutesEnd - minutesStart
	hours := minutes / 60
	min := minutes % 60
	if hours == 0 && min == 0 {
		hours += 24
	}

	fmt.Printf("O JOGO DUROU %d HORA(S) E %d MINUTO(S)\n", hours, min)
}
