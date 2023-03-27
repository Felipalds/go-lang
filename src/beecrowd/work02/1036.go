package main

import (
	"fmt"
	"math"
)

func main () {
  var a, b, c, delta, x1, x2 float64
  fmt.Scan(&a, &b, &c)
  delta = math.Pow(b, 2) - 4.0 * a * c
  if delta < 0 || a == 0 {
    fmt.Println("Impossivel calcular")
  } else {
    x1 = (b*-1 + math.Sqrt(delta)) / (2*a)
    x2 = (b*-1 - math.Sqrt(delta)) / (2*a)

    fmt.Printf("R1 = %.5f\nR2 = %.5f\n", x1, x2)
  }
}
