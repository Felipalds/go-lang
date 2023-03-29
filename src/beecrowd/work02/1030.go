package main

import "fmt"

func main() {
	var initialPersonAmount, personAmount, jump, selectedPerson int
	fmt.Scan(&personAmount, &jump)
	initialPersonAmount = personAmount
	for personAmount > 1 {
		fmt.Printf("Pessoas restantes: %d\n", personAmount)
		selectedPerson += jump
		if selectedPerson > initialPersonAmount {
			selectedPerson -= initialPersonAmount
		}

		fmt.Printf("Pessoa selecionada: %d\n\n", selectedPerson)
		personAmount--
	}
	fmt.Println(selectedPerson)
}
