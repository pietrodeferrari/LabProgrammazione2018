/*
Stampa un quadrato di * e +

Dato un numero n letto a tastiera, stampa un quadrato composto da n righe di * e + alternate
*/
package main

import (
	"fmt"
)

func main() {

	var numero int
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&numero)

	for i := 0; i < numero; i++ {
		for j := 0; j < numero; j++ {
			if i%2 == 0 {
				fmt.Print("* ")
			} else {
				fmt.Print("+ ")
			}
		}
		fmt.Println()
	}

}
