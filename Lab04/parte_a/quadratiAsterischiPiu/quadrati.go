/*
Stampa un quadrato di asterischi

Dato un numero n letto a tastiera, stampa un quadrato n x n con * sui bordi e + al centro
*/
package main

import "fmt"

func main() {

	var numero int

	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&numero)

	for i := 0; i < numero; i++ {
		for j := 0; j < numero; j++ {
			if i == 0 || j == 0 || i == numero-1 || j == numero-1 {
				fmt.Print("* ")
			} else {
				fmt.Print("+ ")
			}
		}
		fmt.Println()
	}

}
