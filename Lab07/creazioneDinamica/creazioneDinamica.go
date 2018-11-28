package main

import "fmt"

func main() {

	var n int
	fmt.Scan(&n)

	// Crea una fetta con un array sottostante di dimensione 10
	arr := make([]int, n)

	// Leggo tutti gli n elementi, ognuno memorizzato in una locazione diversa dell'array
	fmt.Println("Inserisci", n, "interi:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	// Stampo tutti i numeri dell'array partendo dall'ultimo elemento
	fmt.Println("Numeri stampati in ordine inverso:")
	for i := n - 1; i >= 0; i-- {
		fmt.Print(arr[i], " ")
	}

	fmt.Println()
}
