package main

import "fmt"

func main() {

	var word1, word2, word3 string
	fmt.Scan(&word1, &word2, &word3)

	if word1 == "vertebrado" {
		if word2 == "ave" {
			if word3 == "carnivoro" {
				fmt.Println("aguia")
			} else {
				fmt.Println("pomba")
			}
		} else {
			if word3 == "onivoro" {
				fmt.Println("homem")
			} else {
				fmt.Println("vaca")
			}
		}
	} else {
		if word2 == "inseto" {
			if word3 == "hematofago" {
				fmt.Println("pulga")
			} else {
				fmt.Println("lagarta")
			}
		} else {
			if word3 == "hematofago" {
				fmt.Println("sanguessuga")
			} else {
				fmt.Println("minhoca")
			}
		}
	}
}
