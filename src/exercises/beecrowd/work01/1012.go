package main

import ("fmt";"math")

const PI = 3.14159

func main(){
  var A, B, C, triangle, circle, trapezy, square, rectangle float64  

  fmt.Scan(&A, &B, &C)
  triangle = A * C / 2
  circle = PI * math.Pow(C, 2)
  trapezy = (A + B) * C / 2
  square = math.Pow(B, 2)
  rectangle = A * B

  fmt.Printf("TRIANGULO: %.3f\nCIRCULO: %.3f\nTRAPEZIO: %.3f\nQUADRADO: %.3f\nRETANGULO: %.3f\n",triangle, circle, trapezy, square, rectangle)
}
