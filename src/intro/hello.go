package main

/*
nome do pacote que está sendo criado, equivale ao nome do programa
sempre que for um exe, será main, se for lib ou incluído à outro
programa, pode ser outro nome
*/

import "fmt"

func main() {
	fmt.Println("Hello world!")
	//funções ou métodos com letras maiúsculas são exportadas
	//e visíveis fora do seu pacote
}
