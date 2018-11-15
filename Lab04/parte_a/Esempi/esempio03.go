/*
Il package stampa una sequenza di numeri

Dato un numero n inserito da tastiera, il package stampa tutti i numeri compresi tra 1 e n (estremi inclusi).
La sequenza è prodotta su un'unica linea di output, con i numeri separati da uno spazio.
 */
package main

import (
	"fmt"
)


// Questo programma è errato, correggilo (3 errori: 1 di sintassi, 2 di logica del programma)
func main() {

	var n int

	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&n)

	for i := 1; i < n { // for i := 1; i <= n ; {
		fmt.Print(i) // fmt.Print(i, " ")
		i++
	}

	fmt.Println()

}
