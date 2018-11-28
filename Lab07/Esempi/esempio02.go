package main

import "fmt"

func main() {

	a := [...]int{1, 2, 3, 4}

	fmt.Println(a)

	for i:=len(a)-1; i>=0; i-- {
		fmt.Println("Indice", i, " - Valore:", a[i])
	}

	fmt.Println()

	for i, _ := range a {
		fmt.Println("Indice", len(a)-1-i, " - Valore:", a[len(a)-1-i])
	}

	// Entrambi i cicli for stampano il contenuto dell'array a ma al contrario

}