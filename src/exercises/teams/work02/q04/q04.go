package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {

	in, inerr := os.Open("dobra.in")
	out, outerr := os.Create("dobra.out")

	if inerr != nil || outerr != nil {
		fmt.Println("Erro ao abrir/criar arquivos")
		panic("ERRO")
	}

	defer in.Close()
	defer out.Close()

	reader := bufio.NewReader(in)
	writer := bufio.NewWriter(out)
	defer writer.Flush()

	var cuts int
	tests := 1

	for {
		fmt.Fscanf(reader, "%d\n", &cuts)
		if cuts == -1 {
			break
		}

		result := int64(math.Pow(4.0, float64(cuts)) + math.Pow(2, float64(cuts+1)) + 1)

		fmt.Fprintf(writer, "Teste %d\n", tests)
		fmt.Fprintf(writer, "%d\n\n", result)

		tests++

	}

}
