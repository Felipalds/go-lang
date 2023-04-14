package main

import (
	"fmt"
	"os"
)

func main() {
	var (
		teste    int16
		stringin string
		vetnum   [100]int16
		soma     int16
		i        int16
		cont     int16 = 1
	)

	in, ferror := os.Open("calcula.in")
	out, ferror2 := os.Create("calcula.out")

	if ferror != nil || ferror2 != nil {
		fmt.Println("Erro ao abrir arquivo!")
	} else {

		for {
			fmt.Fscanf(in, "%d\n", &teste)

			if teste == 0 {
				return
			}
			soma = 0
			fmt.Fscanf(in, "%s\n", &stringin)

			xnums := make([]interface{}, len(vetnum))

			for n := range vetnum {
				xnums[n] = &vetnum[n]
			}

			fmt.Sscan(stringin, xnums...)

			for i = 0; i < teste; i++ {
				soma += vetnum[i]
			}

			fmt.Fprintf(out, "Teste %d\n%d\n\n", cont, soma)
			cont++
		}
	}
}
