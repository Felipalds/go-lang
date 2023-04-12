package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var quantity int
	var expression string

	file, ferror := os.Open("./calcula.in")

	if ferror != nil {
		fmt.Println("Erro: ", ferror)
	} else {
		reader := bufio.NewReader(file)
		for {
			fmt.Fscanf(reader, "%d\n", &quantity)
			if quantity == 0 {
				break
			}
			fmt.Fscanf(reader, "%s\n", &expression)
			fmt.Println(expression)
			calculate(expression)
		}
	}
}

func calculate(expression string) {
	var numbers []int
	var operators []byte
	len := len(expression)
	for c := 0; c < len; c++ {
		if expression[c] == '-' || expression[c] == '+' {
			append(operators, expression[c])
		}
	}
}
