package main

import (
  "fmt"
)

type Rational struct {
  numerator, denominator int
}

func (r *Rational) sum (r1 Rational) {
  r.numerator = r.numerator * r1.denominator + r.denominator * r1.numerator
  r.denominator = r.denominator * r1.denominator
}

func (r *Rational) sub (r1 Rational) {
  r.numerator = r.numerator * r1.denominator - r.denominator * r1.numerator
  r.denominator = r.denominator * r1.denominator
}

func (r *Rational) mult (r1 Rational) {
  r.numerator = r.numerator * r1.numerator
  r.denominator = r.denominator * r1.denominator
}

func (r *Rational) div (r1 Rational) {
  r.numerator = r.numerator * r1.denominator
  r.denominator = r.denominator * r1.numerator
}

func main() {
  zoz := Rational{1, 2} 
  pedro := Rational{1, 2}
  pedro2 := Rational { 1, 3 }
  pedro3 := Rational { 2, 3 }
  pedro4 := Rational { 5, 3 }

  pedro.sum(zoz)
  pedro2.sub(zoz)
  pedro3.mult(zoz)
  pedro4.div(zoz)

  fmt.Println(pedro)
  fmt.Println(pedro2)
  fmt.Println(pedro3)
  fmt.Println(pedro4)

}
