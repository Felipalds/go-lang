/**
 * Exemplo de manipulação de arquivos em Go:
 * Abre o arquivo e lê uma linha de cada vez como uma string
 */
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	// Tenta abrir o arquivo
	file, ferror := os.Open("1001.go") // abre o arquivo com buffer em file

	// testa o erro ocorrido
	if ferror != nil {
		fmt.Println("Erro:", ferror) // Imprime o erro ocorrido
	} else {
		reader := bufio.NewReader(file) // cria um objeto reader para o buffer file
		for {
			line, rerror := reader.ReadString('\n')
			fmt.Print(line) // veja que o '\n' permanece na string
			if rerror == io.EOF {
				fmt.Println("\n--- Fim do arquivo ---")
				break
			}
		}
	}
	if err := file.Close(); err != nil {
		panic(err)
	}
}
