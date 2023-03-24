package main

import ("fmt"

"math")

func main() {
    
  var money float64
  notes := []float64{100, 50, 20, 10, 5, 2, 1, 50, 25, 10, 5, 1}
  var results [12] float64
  var k int

  fmt.Scan(&money)
  for k < len(notes) {
    if k == 7 {
      money *= 100
    }
    results[k] = math.Floor(money/notes[k])
    money -= results[k] * notes[k]
    k++
  }

  k = 0
  fmt.Println("NOTAS: ")
  for k < 6 {
    fmt.Printf("%.0f nota(s) de R$ %.0f.00\n", results[k], notes[k])
    k++
  } 

  fmt.Println("MOEDAS: ")
  for k < 12 {
    fmt.Printf("%.0f moeda(s) de R$ ", results[k])
    if(k == 6){
      fmt.Printf("%.0f.00\n", notes[k])
    }
    if k > 6 && k < 10 {
      fmt.Printf("0.%.0f\n", notes[k])
    }
    if k >= 10 {
      fmt.Printf("0.0%.0f\n", notes[k])
    }
    k++
  }

}
