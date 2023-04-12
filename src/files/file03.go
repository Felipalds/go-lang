/**
 * Exemplo de leitura e gravação de arquivos em Go
 */
package main

import (
	"io"
	"os"
)

func main() {
	// Abre o arquivo de entrada
	infile, ferror := os.Open("input.txt")
	if ferror != nil {
		panic(ferror)
	}
	// Fecha o arquivo de entrada
	// defer transfere a execução da função para o final
	defer func() { // usa uma função anônima aqui
		if ferror := infile.Close(); ferror != nil {
			panic(ferror)
		}
	}()

	// Abre o arquivo de saída para gravação
	outfile, ferror := os.Create("output.txt")
	if ferror != nil {
		panic(ferror)
	}
	// Fecha o arquivo de saída
	defer func() {
		if ferror := outfile.Close(); ferror != nil {
			panic(ferror)
		}
	}()

	// Cria um buffer para conter o que será lido
	buf := make([]byte, 1024) // cria um slice vazio de 1024 bytes
	for {
		// lê a entrada
		n, ferror := infile.Read(buf)
		if ferror != nil && ferror != io.EOF {
			panic(ferror)
		}
		if n == 0 {
			break
		}

		// escreve o buffer no arquivo
		if _, ferror := outfile.Write(buf[:n]); ferror != nil {
			panic(ferror)
		}
	}
}
