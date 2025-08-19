package main

import "fmt"

func main() {
	var k, q int
	var t string
	var coe, rat, sap int
	fmt.Scan(&k)
	for i := 0; i < k; i++ {
		fmt.Scan(&q, &t)
		if t == "C" {
			coe += q
		}
		if t == "R" {
			rat += q
		}
		if t == "S" {
			sap += q
		}
	}
	tot := coe + rat + sap
	fmt.Printf("Total: %d cobaias\n", tot)
	fmt.Printf("Total de coelhos: %d\n", coe)
	fmt.Printf("Total de ratos: %d\n", rat)
	fmt.Printf("Total de sapos: %d\n", sap)
	fmt.Printf("Percentual de coelhos: %.2f %%\n", float64(coe)/float64(tot)*100)
	fmt.Printf("Percentual de ratos: %.2f %%\n", float64(rat)/float64(tot)*100)
	fmt.Printf("Percentual de sapos: %.2f %%\n", float64(sap)/float64(tot)*100)
}
