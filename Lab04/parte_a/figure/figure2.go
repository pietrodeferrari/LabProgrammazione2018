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

	// stampa n righe di caratteri ** e ++ alternati sia sulla stessa righa che sulla colonna
	for i := 0; i < numero; i++ {
		for j := 0; j < numero; j++ {
			offset := j
			if i%2 != 0 {
				offset += 2
			}

			if offset%4 <= 1 {
				fmt.Print("*")
			} else {
				fmt.Print("+")
			}
		}
		fmt.Println()
	}

}
