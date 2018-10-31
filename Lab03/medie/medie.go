/*
Calcola valori medi di una sequenza di numeri

Vengono calcolate la media aritmetica, la media geometrica, la media quadratica e la media armonica di una sequenza di numeri positivi inserita da tastiera
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	var numero int
	numeriLetti := 0
	// Tutte le variabili che contengono le medie finali sono inizializzate a 0 tranne la media geometrica.
	// La media geometrica è l'unica che durante il ciclo viene MOLTIPLICATA per un valore, mentre a tutte le altre viene SOMMATO un valore.
	// Quando si calcola il risultato di una produttoria, il valore iniziale della variabile che conterrà la produttoria deve essere 1
	// Quando si calcola il risultato di una sommatoria, il valore iniziale della variabile che conterrà la sommatoria deve essere 0
	// Questo perché 1 e 0 sono gli elementi neutri della moltiplicazione e della somma
	mAritmetica, mGeometrica, mQuadratica, mArmonica := 0.0, 1.0, 0.0, 0.0

	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&numero)

	fmt.Print("Inserisci ", numero, " numeri: ")
	for i := 0; i < numero; i++ {
		var x float64
		fmt.Scan(&x)

		if x > 0 {
			numeriLetti++
			mAritmetica += x
			mGeometrica *= x
			mQuadratica += x * x
			mArmonica += 1 / x
		}
	}

	mAritmetica /= float64(numeriLetti)
	mGeometrica = math.Pow(mGeometrica, 1/float64(numeriLetti))
	mQuadratica = math.Sqrt(mQuadratica / float64(numeriLetti))
	mArmonica = float64(numeriLetti) / mArmonica

	fmt.Println("Media aritmetica:", mAritmetica,
		"\nMedia geometrica:", mGeometrica,
		"\nMedia quadratica:", mQuadratica,
		"\nMedia armonica:", mArmonica)

}
