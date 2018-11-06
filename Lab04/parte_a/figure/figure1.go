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

	// stampa n righe di n caratteri * e + alternati
	for i := 0; i < numero; i++ {
		for j := 0; j < numero; j++ {
			if j%2 == 0 {
				fmt.Print("*")
			} else {
				fmt.Print("+")
			}
		}
		fmt.Println()
	}

}
