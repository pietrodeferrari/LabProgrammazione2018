/*
Produce varie figure

Dato un intero in ingresso, produce varie figure composte da * e +
o++++
*o+++
**o++
***o+
****o
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
				// La diagonale la si stampa confrontando semplicemente se gli indici di riga e colonna sono uguali
				fmt.Print("o")
			case i < j:
				// nel caso l'indice di colonna sia maggiore dell'indice di riga, si sta prendendo in esame il triangolo
				// superiore
				fmt.Print("+")
			case i > j:
				// quando invece l'indice di riga è maggiore dell'indice di colonna siamo nel triangolo inferiore
				fmt.Print("*")
			}
		}
		fmt.Println()
	}

}
