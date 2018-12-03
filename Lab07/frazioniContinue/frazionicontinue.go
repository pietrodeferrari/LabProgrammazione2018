package main

import (
	"fmt"
)

func main() {

	var numero int
	fmt.Println("Inserisci un numero:")
	fmt.Scan(&numero)

	values := make([]int, numero)
	fmt.Println("Inserisci i", numero, "numeri interi:")
	for i := range values {
		fmt.Scan(&values[i])
	}

	if len(values) > 0 {
		fmt.Println("Frazione continua:", frazioneContinua(values))
	} else {
		fmt.Println("Nessun numero in input")
	}

}

func frazioneContinua(values []int) float64 {
	sum := 0.0
	for i := len(values) - 1; i > 0; i-- {
		sum = 1/(float64(values[i]) + sum)
	}
	return float64(values[0]) + sum
}