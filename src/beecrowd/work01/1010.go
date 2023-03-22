package main

import "fmt"

func main(){

	var(
		code1 int
		code2 int
		amount1 int
		amount2 int
		price1 float32
		price2 float32
	)

	fmt.Scan(&code1, &amount1, &price1,
			&code2, &amount2, &price2)

	toPay := float32(amount1) * price1 + float32(amount2) * price2
	fmt.Printf("VALOR A PAGAR: R$ %.2f\n", toPay)

}