package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	in, inerr := os.Open("aero.in")
	out, outerr := os.Create("aero.out")

	if inerr != nil || outerr != nil {
		fmt.Println("Erro ao abrir/criar arquivos")
		panic("ERRO")
	}

	defer in.Close()
	defer out.Close()

	reader := bufio.NewReader(in)
	writer := bufio.NewWriter(out)
	defer writer.Flush()

	var aeros, flies int
	var from, destiny int
	tests := 1

	for {
		fmt.Fscanf(reader, "%d %d\n", &aeros, &flies)
		if aeros == 0 && flies == 0 {
			break
		}

		aerosVector := make([]int, aeros)

		fmt.Println(aeros)
		fmt.Println(aerosVector)

		for c := 0; c < flies; c++ {
			fmt.Fscanf(reader, "%d %d\n", &from, &destiny)
			aerosVector[from-1]++
			aerosVector[destiny-1]++
		}

		max := vecMax(aerosVector, aeros)
		maxAeros := maxIndexInVector(aerosVector, aeros, max)

		fmt.Fprintf(writer, "Teste %d\n", tests)
		maxAerosLen := len(maxAeros)

		for c := 0; c < maxAerosLen; c++ {
			fmt.Fprintf(writer, "%d ", maxAeros[c])
		}
		fmt.Fprintf(writer, "\n\n")

		tests++

	}

}

func vecMax(vec []int, l int) int {
	var max int
	for c := 0; c < l; c++ {
		if c == 0 {
			max = vec[c]
		}
		if vec[c] > max {
			max = vec[c]
		}
	}
	return max
}

func maxIndexInVector(vec []int, l int, max int) []int {
	var maxValues []int
	for c := 0; c < l; c++ {
		if vec[c] == max {
			maxValues = append(maxValues, c+1)
		}
	}
	return maxValues
}
