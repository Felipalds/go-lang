package main

import (
	"fmt"
	"math"
)

func main() {
	for {
		var r1, x1, y1, r2, x2, y2 float64
		n, err := fmt.Scan(&r1, &x1, &y1, &r2, &x2, &y2)
		if n != 6 && err != nil {
			return
		}

		distCenters := math.Sqrt(math.Pow(x1-x2, 2) + math.Pow(y1-y2, 2))

		if distCenters+r2 <= r1 {
			fmt.Println("RICO")
		} else {
			fmt.Println("MORTO")
		}
	}
}
