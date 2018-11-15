/*
Inverte una lista di numeri

Data una lista di interi tra [1,9], stampa la lista di ordine inverso e calcola la somma dei numeri per entrambe le liste.
Le due somme sono calcolate in modo tale che i numeri di posizione dispari in entrambe le lista siano moltiplicati per -1.
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Inserisci i numeri:")

	numeriLetti := 0
	stringaInput, stringaInversa := "", ""
	somma, sommaInversa := 0, 0

	for numero := 1; numero != 0; {
		fmt.Scan(&numero)
		if numero > 0 && numero < 10 {
			stringaInput += strconv.Itoa(numero) + " "
			stringaInversa = strconv.Itoa(numero) + " " + stringaInversa

			if numeriLetti%2 == 0 {
				somma += numero
			} else {
				somma -= numero
			}

			numeriLetti++
		}
	}

	if sommaInversa = somma; numeriLetti%2 == 0 {
		sommaInversa = -somma
	}

	fmt.Println("Termine inserimento.")
	fmt.Println("Sequenza letta:")
	fmt.Println(stringaInput)
	fmt.Println("Sequenza in ordine inverso:")
	fmt.Println(stringaInversa)
	fmt.Println("Valori da calcolare:")
	fmt.Println(somma)
	fmt.Println(sommaInversa)

}
