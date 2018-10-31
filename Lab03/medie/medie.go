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
	mAritmetica, mGeometrica, mQuadratica, mArmonica := 0.0, 1.0, 0.0, 0.0

	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&numero)

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
	mGeometrica = math.Pow(mGeometrica, 1.0/float64(numeriLetti))
	mQuadratica = math.Sqrt(mQuadratica / float64(numeriLetti))
	mArmonica = float64(numeriLetti) / mArmonica

	fmt.Println("Media aritmetica:", mAritmetica,
		"\nMedia geometrica:", mGeometrica,
		"\nMedia quadratica:", mQuadratica,
		"\nMedia armonica:", mArmonica)

}
