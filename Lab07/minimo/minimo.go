package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {

	const DIMENSIONE = 5

	var matrice [DIMENSIONE][DIMENSIONE]float64

	rand.Seed(int64(time.Now().Nanosecond()))

	for i := 0; i < DIMENSIONE; i++ {
		for j := 0; j < DIMENSIONE; j++ {
			matrice[i][j] = float64(rand.Intn((i*j + 1) * 100))
		}
	}

	for _, a := range matrice {
		for _, v := range a {
			fmt.Printf("%.0f ", v)
		}
		fmt.Println()
	}

	minimo := math.MaxFloat64
	rigaMinimo, colonnaMinimo := -1, -1
	for riga, a := range matrice {
		for colonna, v := range a {
			if minimo > v {
				minimo = v
				rigaMinimo = riga
				colonnaMinimo = colonna
			}
		}
	}

	fmt.Println("Minimo:", minimo)
	fmt.Println("Riga minimo")
	for i := range matrice {
		fmt.Printf("%.0f ", matrice[rigaMinimo][i])
	}
	fmt.Println("\nColonna minimo")
	for i := range matrice {
		fmt.Printf("%.0f\n", matrice[i][colonnaMinimo])
	}
}
