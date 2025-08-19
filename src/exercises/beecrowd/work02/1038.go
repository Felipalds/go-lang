package main
import "fmt"

func main() {
  var code, amount, price float32
  fmt.Scan(&code, &amount)

  switch code {
  case 1: price = 4
  case 2: price = 4.5
  case 3: price = 5
  case 4: price = 2
  case 5: price = 1.5
  }
  
  price *= amount
  fmt.Printf("Total: R$ %.2f\n", price)

}
