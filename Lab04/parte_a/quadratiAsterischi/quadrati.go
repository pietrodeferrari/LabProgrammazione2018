/*
Stampa un quadrato di asterischi

Dato un numero n letto a tastiera, stampa un quadrato n x n di * intervallati da spazi
*/
package main

import "fmt"

func main() {

	var numero int

	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&numero)

	for i := 0; i < numero; i++ {
		for j := 0; j < numero; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}

}
