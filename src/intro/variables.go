package main

import "fmt"

var global uint8 = 1

func main() {
	var a = "Hello August!!!"
	fmt.Println(a, global)

	var b, c = 1, 2
	string1 := "AAAAAAA"
	fmt.Println(b, c)
	fmt.Println("b + c =", b+c)
	fmt.Printf("%d", b)
	fmt.Printf("%v", string1)
	//Go will infer the type of initialized varabe.
	// IMPLEMENTAÇÃO EM SOFTWARE PARA RODAR EM QUALQUER ARQUITETURA!!!
	var d int
	var e uint8 = 255    //0-255
	var f uint16 = 65535 //0-65535
	// ainda temos o uint32 e uint64
	//int8 -128 - 127
	//int16
	//int32
	//int64
	//float32
	//float64
	//complex64
	//comples128
	fmt.Println(d, e, f)
	/*
		Variables declared without a corresponding
		initialization are zero-valued.
		For example, the zero value for an int is 0.
	*/

	apple := "Steve Jobs"
	fmt.Println(apple)

	//IMPLEMENTAÇÃO EM HARDWARE:
	/*
		uint - pode ser 32 ou 64, dependendo da máquina
		int - pode ser 32 ou 64
		uintptr - unsinged para ponteiros, grande o bastante para
			caber o ponteiro na máquina
		byte - alias para uint8
		rune alias para int32 (UNICODE UTF-8)
	*/

	l1 := 1
	//l1 = 1 assim nao funfa
	// ele sempre coloca o maior tipo como tipo
	l2 := "A" //string
	l3 := 'A' //rune, printando sai o numero
	// := nao e atribuiçao, é declaração short hand. só pode na declaraçao
	fmt.Println(l1, l2, l3)

}
