package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	n := leggiN()

	if n < 0 {
		fmt.Println("Altezza del rombo minore di zero.")
	} else {
		stampaRomboAlt(n)
	}

}

func leggiN() (n int) {
	if len(os.Args) < 2 {
		return
	}
	n, _ = strconv.Atoi(os.Args[1])
	return
}

func stampaRombo(n int) {
	const BASE = 'A'

	// Una iterazione per ogni riga
	for i:=0; i<2*n - 1; i++ {

		// Una iterazione per ogni colonna a destra della colonna di sole 'A'
		for j:=0; j<n-1; j++ {
			if (i < n && j + i >= n-1) || (i >= n && j > i-n ) {
				fmt.Printf("%c ", BASE+(n-1-j))
			} else {
				fmt.Printf("  ")
			}
		}

		// Stampo una singola 'A'
		fmt.Printf("%c", BASE)

		// Una iterazione per ogni colonna a sinistra della colonna di sole 'A'
		for j:=0; j<n-1; j++ {
			if (i < n && j < i) || (i>=n && j + (i-n) < n-2) {
				fmt.Printf(" %c", BASE+(j+1))
			}
		}

		fmt.Println()
	}

}

func stampaRomboAlt(n int) {
	const BASE = 'A'
	lettere := "A"
	for i:=1; i<n; i++ {
		lettere = string(BASE + i) + lettere + string(BASE + i)
	}
	spazi := strings.Repeat(" ", n)

	for i:=0; i<n; i++ {
		indice := n-1-i
		riga := spazi[:indice] + lettere[indice:2*n-1-indice]
		fmt.Println(strings.Join(strings.Split(riga, ""), " "))
	}

	for i:=n-2; i>=0; i-- {
		indice := n-1-i
		riga := spazi[:indice] + lettere[indice:2*n-1-indice]
		fmt.Println(strings.Join(strings.Split(riga, ""), " "))
	}
}