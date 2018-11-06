/*
Produce varie figure

Dato un intero in ingresso, produce varie figure composte da * e +
*/
package main

import (
	"fmt"
)

func main() {

	var numero int
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&numero)

	// stampa un quadrato con diagonale fatta di o, triangolo superiore di + e triangolo inferiore di *
	for i := 0; i < numero; i++ {
		for j := 0; j < numero; j++ {
			switch {
			case i == j:
				fmt.Print("o")
			case i < j:
				fmt.Print("+")
			case i > j:
				fmt.Print("*")
			}
		}
		fmt.Println()
	}

}
