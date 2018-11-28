package main

import (
	"fmt"
	"math"
)

func main() {

	// Dichiaro la dimensione dell'array come costante
	// in questo modo posso riutilizzare DIMENSIONE
	// all'interno del codice, senza trascinarmi il valore 10 ovunque
	// Questo rende il codice più facile da leggere e manutenibile
	const DIMENSIONE = 10

	var arr1 [DIMENSIONE]int
	var arr2 [DIMENSIONE]int
	var arr3 [DIMENSIONE]int

	for i := 0; i < DIMENSIONE; i++ {
		arr1[i] = int(math.Pow(2, float64(i)))
	}

	for _, v := range arr1 {
		fmt.Print(v, " ")
	}

	fmt.Println()

	for i := 0; i < DIMENSIONE; i++ {
		arr2[i] = arr1[(i+1)%DIMENSIONE] - arr1[i]
	}

	for _, v := range arr2 {
		fmt.Print(v, " ")
	}

	fmt.Println()

	for i := 0; i < DIMENSIONE; i++ {
		arr3[i] = arr2[DIMENSIONE-1-i]
	}

	for _, v := range arr3 {
		fmt.Print(v, " ")
	}

}
