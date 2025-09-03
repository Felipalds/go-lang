package main

import (
	"fmt"
	"time"
)

func main() {
	timeParsed, err := time.ParseDuration("-2s")
	if err != nil {
		fmt.Errorf("%v", err)
	}
	otherTime, _ := time.ParseDuration("2s")
	fmt.Println(timeParsed < otherTime)
	fmt.Println(otherTime - timeParsed)
	fmt.Println(timeParsed < 0)

}
