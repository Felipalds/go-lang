package main
import "fmt"

func main(){
  const q = 7
  notes := []int {100, 50, 20, 10, 5, 2, 1}
  var value, iterator int
  fmt.Scan(&value) 
  fmt.Println(value)


  for iterator < q {
    fmt.Printf("%d nota(s) de R$ %d,00\n", value / notes[iterator], notes[iterator])
    value -= notes[iterator] * (value / notes[iterator])
    iterator++
  } 
}

