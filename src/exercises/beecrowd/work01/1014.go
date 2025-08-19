package main

import "fmt"

func main(){
  var x, y float32
  fmt.Scan(&x, &y)
  m := x/y

  fmt.Printf("%.3f km/l\n", m)
}


