package main

import (
	"fmt"
	"math"
)

func main() {

	const DIMENSIONE = 10

	var arr1 [DIMENSIONE]int

	for i := 0; i < DIMENSIONE; i++ {
		arr1[i] = int(math.Pow(2, float64(i)))
	}

	for _, v := range arr1 {
		fmt.Print(v, " ")
	}

	fmt.Println()

	slice := arr1[4:7]

	for _, v := range slice {
		fmt.Print(v, " ")
	}

	fmt.Println()

	slice = slice[:cap(slice)]

	for _, v := range slice {
		fmt.Print(v, " ")
	}

	fmt.Println()

	slice2 := arr1[:]

	// Ad ogni iterazione decremento la lunghezza della fetta
	// spostando in avanti di una posizione l'inizio della fetta
	// nell'array
	for i := 0; i < 10; i++ {
		fmt.Print(slice2[0], " ")
		slice2 = slice2[1:]
	}

	fmt.Println()

	slice3 := arr1[:]

	// Ad ogni iterazione decremento la lunghezza della fetta
	// spostando indietro di una posizione la fine della fetta
	// nell'array
	for i := 0; i < 10; i++ {
		fmt.Print(slice3[len(slice3)-1], " ")
		slice3 = slice3[:len(slice3)-1]
	}

	fmt.Println()

}
