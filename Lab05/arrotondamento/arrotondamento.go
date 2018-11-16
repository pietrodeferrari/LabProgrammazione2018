package main

import (
	"fmt"
	"math"
)

func main() {

	var n int
	var numero float64

	fmt.Println("Inserisci il numero da arrotondare: ")
	fmt.Scan(&numero)
	fmt.Println("Inserisci il numero di cifre dopo la virgola: ")
	fmt.Scan(&n)

	potenza := math.Pow(10, float64(n))
	troncamento := float64(int(numero*potenza)) / potenza
	fmt.Println("Valore troncato =", troncamento)

	arrotondamento := math.Round(numero*potenza*10) / (potenza * 10)
	fmt.Println("Valore arrotondato =", arrotondamento)

}
