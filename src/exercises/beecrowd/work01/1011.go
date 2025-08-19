package main

import "fmt"
import "math"

const pi = 3.14159

func main(){
    var radio float64
    fmt.Scan(&radio)
    var volume float64 = (4.0/3.0)*pi * math.Pow(radio, 3)
    fmt.Printf("VOLUME = %.3f\n", volume)


}
