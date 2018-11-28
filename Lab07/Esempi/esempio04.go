package main

import (
	"fmt"
	"strings"
)

func main() {

	var a = [6]int{1, 2, 3, 4, 5, 6}

	var b []int

	// Passando un array ad una funzione viene fatta una copia di tale array
	// e quindi le modifiche eventualmente fatte all'interno della funzione
	// si perdono e non hanno effetto sull'array originale
	stampaArray(a)
	modificaArray(a)
	stampaArray(a)

	fmt.Println(strings.Repeat("=", 10))

	b = a[:]
	// Anche una fetta viene passata per copia, ma ogni fetta contiene un puntatore
	// all'array e le modifiche fatte sulla fetta si riflettono quindi sull'array
	stampaSlice(b)
	modificaSlice(b)
	stampaSlice(b)
	stampaArray(a)
}

func modificaArray(a [6]int) {
	a[0] = 10
}

func modificaSlice(a []int) {
	a[0] = 10
}

func stampaArray(a [6]int) {
	for _, v := range a {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func stampaSlice(a []int) {
	for _, v := range a {
		fmt.Print(v, " ")
	}
	fmt.Println()
}
