// pasta deve ter o mesmo nome do que a biblioteca!
// biblioteca nao deve ter main
package main

import (
	"fmt"

	r "packages/rational"
)

func main() {
	a := r.Rational{2, 3}
	b := r.Rational{2, 3}
	b.Sum(a)
	fmt.Println(b)
}
