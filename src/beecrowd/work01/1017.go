package main

import "fmt"

const spent = 12.0
func main(){
  var time, speed int
  fmt.Scan(&time, &speed)
  liters := float32(time) * float32(speed) / spent
  fmt.Printf("%.3f\n", liters)
  
}
