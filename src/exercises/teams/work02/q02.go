package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {

	in, ferror := os.Open("lua.in")
	out, ferror2 := os.Create("lua.out")
	test := 1
	defer out.Close()

	if ferror != nil || ferror2 != nil {
		fmt.Println("Erro ao abrir arquivo!")
	} else {
		var testAmount, minuteInterval int
		var min, max int

		reader := bufio.NewReader(in) //buffer
		readerOut := bufio.NewWriter(out)

		for {
			//reading from file
			fmt.Fscanf(reader, "%d %d\n", &testAmount, &minuteInterval)
			if testAmount == 0 && minuteInterval == 0 {
				break
			}
			temperatures := make([]int, testAmount)

			for i := 0; i < testAmount; i++ {
				var current int
				fmt.Fscanf(reader, "%d\n", &current)
				temperatures[i] = current
			}

			min, max = calculateAverage(temperatures, testAmount, minuteInterval)

			fmt.Fprintf(readerOut, "Teste %d\n", test)
			fmt.Fprintf(readerOut, "%d %d\n\n", min, max)
			test++
		}
		readerOut.Flush()
	}

}

func calculateAverage(
	temperatures []int,
	testAmount int,
	minuteInterval int) (int, int) {
	var min, max, sum int
	for c := 0; c < testAmount; c++ {
		if c == testAmount-minuteInterval+1 {
			break
		}
		for j := 0; j < minuteInterval; j++ {
			sum += temperatures[c+j]
		}
		average := float64(sum) / float64(minuteInterval)
		average = math.Trunc(average)
		averageInt := int(average)
		if c == 0 {
			min = averageInt
			max = averageInt
		}
		if averageInt < min {
			min = averageInt
		}
		if averageInt > max {
			max = averageInt
		}
		sum = 0
	}
	return min, max
}
