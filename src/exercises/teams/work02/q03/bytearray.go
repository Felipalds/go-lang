package main

import (
	"fmt"
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

	v := NewBitVector(3)
	fmt.Println(v)
	v.Increment()
	fmt.Println(v)
	v.Increment()
	fmt.Println(v)
	v.Increment()
	fmt.Println(v)

	v.Increment()
	fmt.Println(v)
	v.Increment()
	fmt.Println(v)

	v.Increment()
	fmt.Println(v)
	v.Increment()
	fmt.Println(v)
	v.Increment()
	fmt.Println(v)

}
