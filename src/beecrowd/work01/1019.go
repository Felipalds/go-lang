package main

import "fmt"

func main() {
  var seconds, i int
  timestamp := [3] int { 3600, 60, 1}
  fmt.Scan(&seconds)
  
  for i < len(timestamp) {
    x := seconds / timestamp[i]
    seconds -= timestamp[i] * (seconds / timestamp[i])
    if i == len(timestamp) - 1 {
      fmt.Println(x)
    } else {  
      fmt.Print(x, ":")
    } 
    i++
  }
}
