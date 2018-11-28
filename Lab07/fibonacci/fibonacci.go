package main

import "fmt"

func main() {
	const DIMENSIONE = 50

	var numeri [DIMENSIONE]int

	// Primi due valori della serie inizializzati a 1 per fibonacci(0) e fibonacci(1)
	numeri[0], numeri[1] = 1, 1
	// Questo codice funziona perché so che i valori della serie richiesti son 50
	// Se il numero di valori fosse un input, questo codice potrebbe generare errori!
	// Infatti, se venisse richiesto solo il primo valore della serie di fibonacci,
	// la riga 11 genererebbe un errore perché cercherebbe di accedere ad una zona dell'array
	// non allocata

	for i:=2; i<DIMENSIONE; i++ {
		numeri[i] = numeri[i-1] + numeri[i-2]
	}

	fmt.Println(numeri)
}
