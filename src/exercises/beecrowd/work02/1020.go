package main

import "fmt"

func main() {
  var days int
  divisors := [] int {365, 30, 1}
  var results [3] int
  var k = 0
  fmt.Scan(&days)

  for k < len(divisors) {
    results[k] = days / divisors[k]
    days -= divisors[k] * results[k]
    k++
  }
  
  fmt.Printf("%d ano(s)\n", results[0])
  fmt.Printf("%d mes(es)\n", results[1])  
  fmt.Printf("%d dia(s)\n", results[2])
}
