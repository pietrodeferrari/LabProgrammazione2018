/*
Il package restituisce informazioni su una sequenza di n numeri

Le informazioni restituite sono la somma, minimo e massimo, numero di interi positivi, negativi, e nulli.
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	// Tute le variabili intere sono inizializzate a 0 di default
	var (
		n,
		min, max,
		somma,
		positivi, negativi, nulli int
	)

	fmt.Scan(&n)
	fmt.Print("Inserisci ", n, " numeri\n")

	// Inizializzazione di minimo e massimo
	min, max = math.MaxInt64, math.MinInt64

	for i := 0; i < n; i++ {
		// Ad ogni iterazione leggo un numero e aggiorno i valori

		var numeroLetto int
		fmt.Scan(&numeroLetto)

		// Incremento somma
		somma += numeroLetto

		// Min e max
		if numeroLetto > max {
			max = numeroLetto
		}
		if numeroLetto < min {
			min = numeroLetto
		}

		// Test segno
		switch {
		case numeroLetto > 0:
			positivi++
		case numeroLetto < 0:
			negativi++
		default:
			nulli++
		}
	}

	fmt.Print("somma = ", somma, "\n")
	fmt.Print("valore minimo = ", min, "\n")
	fmt.Print("valore massimo = ", max, "\n")
	fmt.Print("interi > 0 = ", positivi, "\n")
	fmt.Print("interi < 0 = ", negativi, "\n")
	fmt.Print("interi = 0 = ", nulli, "\n")

}
