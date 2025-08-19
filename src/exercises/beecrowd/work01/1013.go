package main

import (
	"fmt"
	"math"
)

const N = 3

func main(){

  var nums[N] int
  for i := 0; i < N; i++ {
   fmt.Scan(&nums[i]) 
  }

  var bigger int
  for i:= 0; i < N; i++ {
    bigger = (bigger + nums[i] + int(math.Abs(float64(bigger - nums[i])))) / 2
  }

  fmt.Print(bigger, " eh o maior\n")

}
