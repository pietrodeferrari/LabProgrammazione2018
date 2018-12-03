package main

import "fmt"

func main() {

	var n int
	fmt.Println("Inserisci la dimensione n:")
	fmt.Scan(&n)

	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	for i, riga := range matrix {
		for j := range riga {
			riga[j] = (i+1) * (j+1)
		}
	}

	fmt.Println("Tavola pitagorica", n,"x", n)
	fmt.Printf("  ")
	for i := range matrix {
		fmt.Printf("%6d", i+1)
	}
	fmt.Println()

	for i, riga := range matrix {
		fmt.Printf("%2d", i+1)
		for _, valore := range riga {
			fmt.Printf("%6d", valore)
		}
		fmt.Println()
	}

}
