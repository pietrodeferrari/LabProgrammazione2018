/*
Stampa un quadrato di *, + e o

Dato un numero n letto a tastiera, stampa un quadrato composto da n righe di *, + e o alternate
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

			// Essendoci 3 possibili casi (*, + e o)
			// basta controllare il resto della divisione tra l'indice e 3
			// per sapere in quale caso siamo
			switch i % 3 {
			case 0:
				fmt.Print("* ")
			case 1:
				fmt.Print("+ ")
			case 2:
				fmt.Print("o ")
			}
		}
		fmt.Println()
	}

}
