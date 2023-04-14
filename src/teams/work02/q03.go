package main

import (
	"fmt"
	"os"
)

type BitVector struct {
	bits []uint
	size int
}

func NewBitVector(size int) *BitVector {
	bits := make([]uint, (size+31)/32)
	return &BitVector{bits, size}
}

func (v *BitVector) SetBit(index int) {
	if index >= v.size {
		return
	}
	word := index / 32
	bit := uint(index % 32)
	v.bits[word] |= (1 << bit)
}

func (v *BitVector) String() string {
	result := ""
	for i := 0; i < v.size; i++ {
		if v.GetBit(i) {
			result += "1"
		} else {
			result += "0"
		}
	}
	return result
}

func (v *BitVector) GetBit(index int) bool {
	if index >= v.size {
		return false
	}
	word := index / 32
	bit := uint(index % 32)
	return (v.bits[word] & (1 << bit)) != 0
}

func (v *BitVector) Increment() {
	for i := 0; i < v.size; i++ {
		if !v.GetBit(i) {
			v.SetBit(i)
			return
		}
		v.bits[i/32] &^= (1 << uint(i%32))
	}
}

func main() {

	in, inerr := os.Open("tesouro.in")
	out, outerr := os.Create("tesouro.out")

	if inerr != nil || outerr != nil {
		fmt.Println("Erro ao abrir/criar arquivos")
		panic("ERRO")
	}

	defer in.Close()
	defer out.Close()

	var valueA, valueB, itemAmount int
	tests := 1

	for {
		fmt.Fscanf(in, "%d %d %d\n", &valueA, &valueB, &itemAmount)
		if valueA == 0 && valueB == 0 && itemAmount == 0 {
			break
		}

		itemValues := make([]int, itemAmount)
		binaryPack := NewBitVector(itemAmount)

		for c := 0; c < itemAmount; c++ {
			fmt.Fscanf(in, "%d\n", &itemValues[c])
		}
		fmt.Fprintf(out, "Teste %d\n", tests)

		isValid := verifyPacking(itemValues, binaryPack, valueA, valueB)
		if isValid == true {
			fmt.Fprintf(out, "S\n\n")
		} else {
			fmt.Fprintf(out, "N\n\n")
		}

		tests++

	}

}

func verifyPacking(itemValues []int, binaryPack *BitVector, valueA int, valueB int) bool {
	//0 for A
	//1 for B

	// SUM
	itemAmount := len(itemValues)
	var stringOut string
	for c := 0; c < itemAmount; c++ {
		stringOut += "0"
	}

	currentIteration := 0

	for {
		newA := valueA
		newB := valueB
		for c := 0; c < itemAmount; c++ {
			isOn := binaryPack.GetBit(c)
			if isOn == true {
				newA += itemValues[c]
			} else {
				newB += itemValues[c]
			}
		}
		if newA == newB {
			return true
		}

		//GENERATE NEW
		binaryPack.Increment()
		if binaryPack.String() == stringOut {
			break
		}
		currentIteration++
	}
	return false
}
