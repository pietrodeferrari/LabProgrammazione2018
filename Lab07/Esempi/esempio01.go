package main

import "fmt"

func main() {

	var a = [...]int{1, 2, 3, 4}

	fmt.Println(a)

	for i := 0; i < len(a); i++ {
		fmt.Println("Indice", i, " - Valore:", a[i])
	}

	fmt.Println()

	for i, v := range a {
		fmt.Println("Indice", i, " - Valore:", v)
	}

	// Entrambi i cicli for stampano il contenuto dell'array a
}
